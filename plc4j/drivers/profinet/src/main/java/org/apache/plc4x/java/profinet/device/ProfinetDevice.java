/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.profinet.device;

import org.apache.commons.codec.DecoderException;
import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.api.exceptions.PlcException;
import org.apache.plc4x.java.api.messages.PlcBrowseItem;
import org.apache.plc4x.java.api.messages.PlcDiscoveryItem;
import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.profinet.context.ProfinetDeviceContext;
import org.apache.plc4x.java.profinet.gsdml.*;
import org.apache.plc4x.java.profinet.protocol.ProfinetProtocolLogic;
import org.apache.plc4x.java.profinet.readwrite.*;
import org.apache.plc4x.java.profinet.tag.ProfinetTag;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.DefaultPlcBrowseItem;
import org.apache.plc4x.java.spi.messages.DefaultPlcSubscriptionEvent;
import org.apache.plc4x.java.spi.messages.PlcSubscriber;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.model.DefaultPlcConsumerRegistration;
import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionHandle;
import org.apache.plc4x.java.spi.values.PlcSTRING;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.*;
import java.nio.ByteBuffer;
import java.time.Instant;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.function.BiFunction;
import java.util.function.Consumer;
import java.util.function.Function;

public class ProfinetDevice {

    private final Logger logger = LoggerFactory.getLogger(ProfinetDevice.class);
    private static final int DEFAULT_NUMBER_OF_PORTS_TO_SCAN = 100;
    private static final int MIN_CYCLE_NANO_SEC = 31250;
    private final BiFunction<String, String, ProfinetISO15745Profile> gsdHandler;
    private ProfinetDeviceContext deviceContext = new ProfinetDeviceContext();
    DatagramSocket socket = null;
    private String vendorId;
    private String deviceId;
    private List<Thread> threads = new ArrayList<>();

    public ProfinetDevice(String deviceName, String deviceAccess, String subModules, BiFunction<String, String, ProfinetISO15745Profile> gsdHandler)  {
        this.gsdHandler = gsdHandler;
        deviceContext.setDeviceAccess(deviceAccess);
        deviceContext.setSubModules(subModules);
        deviceContext.setDeviceName(deviceName);
        openDeviceUdpPort();
    }

    public BiFunction<String, String, ProfinetISO15745Profile> getGsdHandler() {
        return gsdHandler;
    }

    private void openDeviceUdpPort() {
        // Open the receiving UDP port.
        int count = 0;
        int port = ProfinetDeviceContext.DEFAULT_SEND_UDP_PORT;
        boolean portFound = false;
        while (!portFound && count < DEFAULT_NUMBER_OF_PORTS_TO_SCAN) {
            try {
                socket = new DatagramSocket(port + count);
                portFound = true;
            } catch (SocketException e) {
                count += 1;
                port += 1;
            }
        }
        if (!portFound) {
            throw new RuntimeException("Unable to find free open port");
        }
    }

    private long getObjectId() {
        return deviceContext.getAndIncrementIdentification();
    }

    public String getVendorId() {
        return vendorId;
    }

    public String getDeviceId() {
        return deviceId;
    }

    public void setVendorDeviceId(String vendorId, String deviceId) {
        try {
            this.vendorId = vendorId;
            this.deviceId = deviceId;
            if (deviceContext.getGsdFile() == null) {
                deviceContext.setGsdFile(gsdHandler.apply(vendorId, deviceId));
            }
        } catch (PlcException e) {
            throw new RuntimeException(e);
        }
    }

    private void recordIdAndSend(ProfinetCallable<DceRpc_Packet> callable) {
        deviceContext.getQueue().put(callable.getId(), callable);
        ProfinetMessageWrapper.sendUdpMessage(
            callable,
            this
        );
    }

    public boolean onConnect(PlcSubscriber subscriber) throws ExecutionException, InterruptedException, TimeoutException {
        CreateConnection createConnection = new CreateConnection();
        recordIdAndSend(createConnection);

        startSubscription(subscriber);
        createConnection.getResponseHandled().get(1000L, TimeUnit.MILLISECONDS);

        WriteParameters writeParameters = new WriteParameters();
        recordIdAndSend(writeParameters);
        writeParameters.getResponseHandled().get(1000L, TimeUnit.MILLISECONDS);

        WriteParametersEnd writeParametersEnd = new WriteParametersEnd();
        recordIdAndSend(writeParametersEnd);
        writeParametersEnd.getResponseHandled().get(1000L, TimeUnit.MILLISECONDS);

        return true;
    }

    /*
        Starts the subscription, sending data from controller to device.
     */
    public void startSubscription(PlcSubscriber subscriber) {
        if (deviceContext.getSubscriptionHandle() == null) {
            deviceContext.setSubscriptionHandle(new ProfinetSubscriptionHandle(subscriber));
        }
        Function<Object, Boolean> subscription =
            message -> {
                boolean escape = false;
                long startTime = System.nanoTime();
                while (!escape) {
                    try {
                        CyclicData cyclicData = new CyclicData(startTime);
                        ProfinetMessageWrapper.sendPnioMessage(cyclicData, this);

                        int sleepTime = (int) (deviceContext.getConfiguration().getSendClockFactor() * deviceContext.getConfiguration().getReductionRatio() * (MIN_CYCLE_NANO_SEC/1000000.0));
                        Thread.sleep(sleepTime);
                    } catch (InterruptedException e) {
                        escape = true;
                    }
                }
                return null;
            };

        Thread thread = new Thread(new ProfinetRunnable(null, subscription));
        threads.add(thread);
        thread.start();
    }

    /*
        Attempts to clean up threads after a subscription has been interrupted
     */
    public void stopSubscription() {
        for (Thread thread : threads) {
            thread.interrupt();
        }
        threads.clear();
    }

    /*
        Return metadata about the device. This is sourced from the connection string as well as GSD file.
     */
    public Map<String, PlcValue> getDeviceInfo() {
        Map<String, PlcValue> options = new HashMap<>();
        ProfinetDeviceIdentity deviceIdentity = this.deviceContext.getGsdFile().getProfileBody().getDeviceIdentity();
        options.put("device_id", new PlcSTRING(deviceIdentity.getDeviceID()));
        options.put("vendor_id", new PlcSTRING(deviceIdentity.getVendorId()));
        options.put("vendor_name", new PlcSTRING(deviceIdentity.getVendorName().getValue()));
        if (deviceIdentity.getInfoText() != null && deviceIdentity.getInfoText().getTextId() != null) {
            String key = deviceIdentity.getInfoText().getTextId();
            ProfinetExternalTextList externaltextList = this.deviceContext.getGsdFile().getProfileBody().getApplicationProcess().getExternalTextList();
            for (ProfinetTextIdValue s : externaltextList.getPrimaryLanguage().getText()) {
                if (key.equals(s.getTextId())) {
                    options.put("info_text", new PlcSTRING(s.getValue()));
                    break;
                }
            }

        }

        return options;
    }

    /*
        Create a structure including all the devices tags and child tags.
        The options include metadata about the device.
        The children are a list of configured submodules, with the same format as the parent.
        Each address of the children is formatted with the format i.e. parent.submodule.chiildtag
     */
    public Map<String, List<PlcBrowseItem>> browseTags(Map<String, List<PlcBrowseItem>> browseItems) {
        Map<String, PlcValue> options = getDeviceInfo();
        for (ProfinetModule module : deviceContext.getModules()) {
            browseItems = module.browseTags(browseItems, deviceContext.getDeviceName(), options);
        }

        return browseItems;
    }

    private int generateSessionKey() {
        // Generate a new session key.
        Integer sessionKey = deviceContext.getSessionKeyGenerator().getAndIncrement();
        // Reset the session key as soon as it reaches the max for a 16 bit uint
        if (deviceContext.getSessionKeyGenerator().get() == 0xFFFF) {
            deviceContext.getSessionKeyGenerator().set(1);
        }
        return sessionKey;
    }

    public boolean hasLldpPdu() {
        return deviceContext.isLldpReceived();
    }

    public boolean hasDcpPdu() {
        return deviceContext.isDcpReceived();
    }

    public void handleResponse(Ethernet_FramePayload_IPv4 packet) {
        logger.debug("Received packet for {}", packet.getPayload().getObjectUuid());
        try {
            long objectId = packet.getPayload().getSequenceNumber();
            if (deviceContext.getQueue().containsKey(objectId)) {
                deviceContext.getQueue().get(objectId).handle(packet.getPayload());
            } else {
                PnIoCm_Packet payloadPacket = packet.getPayload().getPayload();
                DceRpc_ActivityUuid activityUuid = packet.getPayload().getActivityUuid();
                long seqNumber = packet.getPayload().getSequenceNumber();
                if (payloadPacket instanceof PnIoCm_Packet_Req) {
                    PnIoCm_Packet_Req req = (PnIoCm_Packet_Req) payloadPacket;
                    for (PnIoCm_Block block : req.getBlocks()) {
                        if (block instanceof PnIoCM_Block_Request) {
                            deviceContext.setState(ProfinetDeviceState.APPLRDY);
                            ApplicationReadyResponse applicationReadyResponse = new ApplicationReadyResponse(activityUuid, seqNumber);
                            recordIdAndSend(applicationReadyResponse);
                            deviceContext.getContext().fireConnected();
                        }
                    }
                } else {
                    throw new RuntimeException("Unable to match Response with Requested Profinet packet");
                }
            }

        } catch (PlcException e) {
            throw new RuntimeException(e);
        }

    }

    public void handle(PlcDiscoveryItem item) {
        logger.debug("Received Discovered item at device");
        if (item.getOptions().containsKey("ipAddress")) {
            deviceContext.setIpAddress(item.getOptions().get("ipAddress"));
        }
        if (item.getOptions().containsKey("portId")) {
            deviceContext.setPortId(item.getOptions().get("portId"));
        }
        if (item.getOptions().containsKey("deviceTypeName")) {
            deviceContext.setDeviceTypeName(item.getOptions().get("deviceTypeName"));
        }
        if (item.getOptions().containsKey("vendorId") && item.getOptions().containsKey("deviceId")) {
            setVendorDeviceId(item.getOptions().get("vendorId"), item.getOptions().get("deviceId"));
        }
        if (item.getOptions().containsKey("deviceName")) {
            deviceContext.setDeviceName(item.getOptions().get("deviceName"));
        }
        if (item.getOptions().containsKey("localMacAddress")) {
            String macString = item.getOptions().get("localMacAddress").replace(":", "");
            try {
                deviceContext.setLocalMacAddress(new MacAddress(Hex.decodeHex(macString)));
            } catch (DecoderException e) {
                throw new RuntimeException(e);
            }
        }
        if (item.getOptions().containsKey("macAddress")) {
            String macString = item.getOptions().get("macAddress").replace(":", "");
            try {
                deviceContext.setMacAddress(new MacAddress(Hex.decodeHex(macString)));
            } catch (DecoderException e) {
                throw new RuntimeException(e);
            }
        }
        if (item.getOptions().containsKey("packetType")) {
            if (item.getOptions().get("packetType").equals("lldp")) {
                deviceContext.setLldpReceived(true);
            }
            if (item.getOptions().get("packetType").equals("dcp")) {
                deviceContext.setDcpReceived(true);
            }
        }
    }

    public void setContext(ConversationContext<Ethernet_Frame> context, ProfinetChannel channel) {
        deviceContext.setContext(context);
        deviceContext.setChannel(channel);
    }

    public ProfinetDeviceContext getDeviceContext() {
        return deviceContext;
    }

    public void handleRealTimeResponse(PnDcp_Pdu_RealTimeCyclic cyclicPdu) {
        logger.debug("Received Real Time Cyclic Data");
        Map<String, ResponseItem<PlcValue>> tags = new HashMap<>();
        ReadBuffer buffer = new ReadBufferByteBased(cyclicPdu.getDataUnit().getData());

        try {
            for (ProfinetModule module : deviceContext.getModules()) {
                module.parseTags(tags, deviceContext.getDeviceName(), buffer);
            }

            Set<Map.Entry<String, ResponseItem<PlcValue>>> entries = tags.entrySet();
            Map<String, ResponseItem<PlcValue>> publishedTags = new HashMap<>();
            for (Map.Entry<String, ResponseItem<PlcValue>> entry : tags.entrySet()) {
                Map<String, String> consumerTags = deviceContext.getSubscriptionHandle().getTags();
                if (consumerTags.containsKey(entry.getKey())) {
                    publishedTags.put(consumerTags.get(entry.getKey()), entry.getValue());
                }
            }

            for (Consumer<PlcSubscriptionEvent> consumer : deviceContext.getSubscriptionHandle().getConsumers()) {
                consumer.accept(new DefaultPlcSubscriptionEvent(Instant.now(), publishedTags));
            }
        } catch (ParseException e) {
            throw new RuntimeException(e);
        }
    }

    public void handleAlarmResponse(PnDcp_Pdu_AlarmLow alarmPdu) {
        stopSubscription();
        deviceContext.setFrameId(0x8001);
        logger.error("Received Alarm Low packet, attempting to re-connect");
        Thread thread = new Thread(new ProfinetRunnable(null, (subscriber) -> {
            try {
                return onConnect(deviceContext.getSubscriptionHandle().getPlcSubscriber());
            } catch (ExecutionException | InterruptedException | TimeoutException e) {
                throw new RuntimeException(e);
            }
        }));
        threads.add(thread);
        thread.start();
    }

    public class CreateConnection implements ProfinetCallable<DceRpc_Packet> {

        CompletableFuture<Boolean> responseHandled = new CompletableFuture<>();
        private long id = getObjectId();

        public CompletableFuture<Boolean> getResponseHandled() {
            return responseHandled;
        }

        public long getId() {
            return id;
        }

        public void setId(long id) {
            this.id = id;
        }

        public DceRpc_Packet create() throws PlcException {
            deviceContext.setSessionKey(ProfinetDevice.this.generateSessionKey());

            List<PnIoCm_Block> blocks = new ArrayList<>();
            blocks.add(new PnIoCm_Block_ArReq(
                    ProfinetDeviceContext.BLOCK_VERSION_HIGH,
                    ProfinetDeviceContext.BLOCK_VERSION_LOW,
                    PnIoCm_ArType.IO_CONTROLLER,
                    ProfinetDeviceContext.ARUUID,
                    deviceContext.getSessionKey(),
                    deviceContext.getLocalMacAddress(),
                    new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                    false,
                    deviceContext.isStartupMode(),
                    false,
                    false,
                    PnIoCm_CompanionArType.SINGLE_AR,
                    false,
                    true,
                    false,
                    PnIoCm_State.ACTIVE,
                    ProfinetDeviceContext.DEFAULT_ACTIVITY_TIMEOUT,
                    ProfinetDeviceContext.UDP_RT_PORT,
                    ProfinetDeviceContext.DEFAULT_PLC4X_STATION_NAME
                )
            );

            blocks.add(
                new PnIoCm_Block_AlarmCrReq(
                    (short) 1,
                    (short) 0,
                    PnIoCm_AlarmCrType.ALARM_CR,
                    0x8892,
                    false,
                    false,
                    1,
                    3,
                    0x0000,
                    200,
                    0xC000,
                    0xA000
                )
            );

            List<PnIoCm_IoCrBlockReqApi> inputApis = Collections.singletonList(
                new PnIoCm_IoCrBlockReqApi(
                    deviceContext.getInputIoPsApiBlocks(),
                    deviceContext.getInputIoCsApiBlocks())
            );

            List<PnIoCm_IoCrBlockReqApi> outputApis = Collections.singletonList(
                new PnIoCm_IoCrBlockReqApi(
                    deviceContext.getOutputIoPsApiBlocks(),
                    deviceContext.getOutputIoCsApiBlocks()
                )
            );

            deviceContext.setInputReq(new PnIoCm_Block_IoCrReq(
                (short) 1,
                (short) 0,
                PnIoCm_IoCrType.INPUT_CR,
                0x0001,
                ProfinetDeviceContext.UDP_RT_PORT,
                false,
                false,
                false,
                false,
                PnIoCm_RtClass.RT_CLASS_2,
                ProfinetDeviceContext.DEFAULT_IO_DATA_SIZE,
                deviceContext.getIncrementAndGetFrameId(),
                deviceContext.getConfiguration().getSendClockFactor(),
                deviceContext.getConfiguration().getReductionRatio(),
                1,
                0,
                0xffffffff,
                deviceContext.getConfiguration().getWatchdogFactor(),
                deviceContext.getConfiguration().getDataHoldFactor(),
                0xC000,
                ProfinetDeviceContext.DEFAULT_EMPTY_MAC_ADDRESS,
                inputApis

            ));

            blocks.add(deviceContext.getInputReq());

            deviceContext.setOutputReq(new PnIoCm_Block_IoCrReq(
                (short) 1,
                (short) 0,
                PnIoCm_IoCrType.OUTPUT_CR,
                0x0002,
                ProfinetDeviceContext.UDP_RT_PORT,
                false,
                false,
                false,
                false,
                PnIoCm_RtClass.RT_CLASS_2,
                ProfinetDeviceContext.DEFAULT_IO_DATA_SIZE,
                deviceContext.getIncrementAndGetFrameId(),
                deviceContext.getConfiguration().getSendClockFactor(),
                deviceContext.getConfiguration().getReductionRatio(),
                1,
                0,
                0xffffffff,
                deviceContext.getConfiguration().getWatchdogFactor(),
                deviceContext.getConfiguration().getDataHoldFactor(),
                0xC000,
                ProfinetDeviceContext.DEFAULT_EMPTY_MAC_ADDRESS,
                outputApis
            ));

            blocks.add(deviceContext.getOutputReq());

            for (PnIoCm_Block_ExpectedSubmoduleReq expectedSubModuleApiBlocksReq : deviceContext.getExpectedSubmoduleReq()) {
                blocks.add(expectedSubModuleApiBlocksReq);
            }

            long arrayLength = 0;
            for (PnIoCm_Block block : blocks) {
                arrayLength += block.getLengthInBytes();
            }

            return new DceRpc_Packet(
                DceRpc_PacketType.REQUEST,
                true,
                false,
                false,
                IntegerEncoding.BIG_ENDIAN,
                CharacterEncoding.ASCII,
                FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                new DceRpc_InterfaceUuid_DeviceInterface(),
                deviceContext.getUuid(),
                0,
                id,
                DceRpc_Operation.CONNECT,
                new PnIoCm_Packet_Req(ProfinetDeviceContext.DEFAULT_ARGS_MAXIMUM, ProfinetDeviceContext.DEFAULT_MAX_ARRAY_COUNT, 0, blocks)
            );
        }

        public void handle(DceRpc_Packet dceRpc_packet) throws PlcException {
            try {
                if ((dceRpc_packet.getOperation() == DceRpc_Operation.CONNECT) && (dceRpc_packet.getPacketType() == DceRpc_PacketType.RESPONSE)) {
                    if (dceRpc_packet.getPayload().getPacketType() == DceRpc_PacketType.RESPONSE) {

                        // Get the remote MAC address and store it in the context.
                        final PnIoCm_Packet_Res connectResponse = (PnIoCm_Packet_Res) dceRpc_packet.getPayload();
                        if ((connectResponse.getBlocks().size() > 0) && (connectResponse.getBlocks().get(0) instanceof PnIoCm_Block_ArRes)) {
                            final PnIoCm_Block_ArRes pnIoCm_block_arRes = (PnIoCm_Block_ArRes) connectResponse.getBlocks().get(0);
                            responseHandled.complete(true);
                            // Update the raw-socket transports filter expression.
                            //((RawSocketChannel) channel).setRemoteMacAddress(org.pcap4j.util.MacAddress.getByAddress(macAddress.getAddress()));
                        } else {
                            responseHandled.complete(true);
                            handleAlarmResponse(null);
                        }
                    } else {
                        throw new PlcException("Unexpected response");
                    }
                } else if (dceRpc_packet.getPacketType() == DceRpc_PacketType.REJECT) {
                    throw new PlcException("Device rejected connection request");
                } else {
                    throw new PlcException("Unexpected response");
                }
            } catch (Exception e) {
                responseHandled.complete(true);
                handleAlarmResponse(null);
            }

        }
    }

    public class WriteParameters implements ProfinetCallable<DceRpc_Packet> {

        CompletableFuture<Boolean> responseHandled = new CompletableFuture<>();
        private long id = getObjectId();

        public CompletableFuture<Boolean> getResponseHandled() {
            return responseHandled;
        }

        public long getId() {
            return id;
        }

        public void setId(long id) {
            this.id = id;
        }

        public DceRpc_Packet create() {

            List<PnIoCm_Block> requests = new ArrayList<>();
            requests.add(
                new IODWriteRequestHeader(
                    (short) 1,
                    (short) 0,
                    0,
                    ProfinetDeviceContext.ARUUID,
                    0x00000000,
                    0x0000,
                    0x0000,
                    0xe040,
                    180,
                    null

                ));
            requests.add(
                new IODWriteRequestHeader(
                    (short) 1,
                    (short) 0,
                    1,
                    ProfinetDeviceContext.ARUUID,
                    0x00000000,
                    0x0000,
                    0x8000,
                    0x8071,
                    12,
                    null

                ));
            requests.add(
                new PDInterfaceAdjust(
                    (short) 1,
                    (short) 0,
                    MultipleInterfaceModeNameOfDevice.NAME_PROVIDED_BY_LLDP
                )
            );
            int seqNumber = 2;
            int index = 1;
            int indexPacket = 0x007B;
            for (String submodule : deviceContext.getSubModules()) {
                ProfinetModuleItem foundModule = null;
                for (ProfinetModuleItem module : deviceContext.getGsdFile().getProfileBody().getApplicationProcess().getModuleList()) {
                    if (module.getId().equals(submodule)) {
                        foundModule = module;
                        break;
                    }
                }

                Integer identNumber = Integer.decode(foundModule.getModuleIdentNumber());
                if (foundModule.getVirtualSubmoduleList().get(0).getRecordDataList() != null) {
                    for (ProfinetParameterRecordDataItem record : foundModule.getVirtualSubmoduleList().get(0).getRecordDataList()) {
                        requests.add(
                            new IODWriteRequestHeader(
                                (short) 1,
                                (short) 0,
                                seqNumber,
                                ProfinetDeviceContext.ARUUID,
                                0x00000000,
                                index,
                                0x0001,
                                record.getIndex(),
                                record.getLength(),
                                new UserData(ByteBuffer.allocate(4).putInt(Integer.valueOf(record.getRef().getDefaultValue())).array(), (long) record.getLength())
                            ));
                        seqNumber += 1;
                    }
                }
                index += 1;
            }

            return new DceRpc_Packet(
                DceRpc_PacketType.REQUEST, true, false, false,
                IntegerEncoding.BIG_ENDIAN, CharacterEncoding.ASCII, FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                new DceRpc_InterfaceUuid_DeviceInterface(),
                deviceContext.getUuid(),
                0,
                id,
                DceRpc_Operation.WRITE,
                new PnIoCm_Packet_Req(16696, 16696, 0,
                    requests)
            );
        }

        @Override
        public void handle(DceRpc_Packet packet) throws PlcException {
            logger.debug("Received a Write Parameter Response");
            responseHandled.complete(true);
        }
    }

    public class WriteParametersEnd implements ProfinetCallable<DceRpc_Packet> {

        CompletableFuture<Boolean> responseHandled = new CompletableFuture<>();
        private long id = getObjectId();

        public CompletableFuture<Boolean> getResponseHandled() {
            return responseHandled;
        }

        public long getId() {
            return id;
        }

        public void setId(long id) {
            this.id = id;
        }

        public DceRpc_Packet create() {
            return new DceRpc_Packet(
                DceRpc_PacketType.REQUEST, true, false, false,
                IntegerEncoding.BIG_ENDIAN, CharacterEncoding.ASCII, FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                new DceRpc_InterfaceUuid_DeviceInterface(),
                deviceContext.getUuid(),
                0,
                id,
                DceRpc_Operation.CONTROL,
                new PnIoCm_Packet_Req(16696, 16696, 0,
                    Arrays.asList(
                        new PnIoCm_Control_Request(
                            (short) 1,
                            (short) 0,
                            ProfinetDeviceContext.ARUUID,
                            deviceContext.getSessionKey(),
                            0x0001

                        )
                    ))
            );
        }

        @Override
        public void handle(DceRpc_Packet packet) throws PlcException {
            logger.debug("Received a Write Parameter End Response");
            responseHandled.complete(true);
        }
    }

    public class ApplicationReadyResponse implements ProfinetCallable<DceRpc_Packet> {

        private final DceRpc_ActivityUuid activityUuid;
        private long id;

        public ApplicationReadyResponse(DceRpc_ActivityUuid activityUuid, long seqNumber) {
            this.activityUuid = activityUuid;
            this.id = seqNumber;
        }

        public CompletableFuture<Boolean> getResponseHandled() {
            return null;
        }

        public long getId() {
            return id;
        }

        public void setId(long id) {
            this.id = id;
        }

        public DceRpc_Packet create() {
            return new DceRpc_Packet(
                DceRpc_PacketType.RESPONSE,
                false,
                true,
                true,
                IntegerEncoding.BIG_ENDIAN,
                CharacterEncoding.ASCII,
                FloatingPointEncoding.IEEE,
                new DceRpc_ObjectUuid((byte) 0x00, 0x0001, Integer.valueOf(deviceId), Integer.valueOf(vendorId)),
                new DceRpc_InterfaceUuid_ControllerInterface(),
                activityUuid,
                0,
                id,
                DceRpc_Operation.CONTROL,
                new PnIoCm_Packet_Res(
                    (short) 0,
                    (short) 0,
                    (short) 0,
                    (short) 0,
                    ProfinetDeviceContext.DEFAULT_MAX_ARRAY_COUNT,
                    0,
                    Arrays.asList(
                        new PnIoCM_Block_Response(
                            (short) 1,
                            (short) 0,
                            ProfinetDeviceContext.ARUUID,
                            deviceContext.getSessionKey(),
                            0x0008,
                            0x0000
                        )
                    ))
            );
        }

        @Override
        public void handle(DceRpc_Packet packet) throws PlcException {
            logger.debug("Received an unintented packet - We were expecting a response for an Application Ready Response");
        }
    }

    public class CyclicData implements ProfinetCallable<Ethernet_Frame> {

        private final long startTime;
        private long id = getObjectId();

        public CyclicData(long startTime) {
            this.startTime = startTime;
        }

        public long getId() {
            return id;
        }

        public void setId(long id) {
            this.id = id;
        }

        public Ethernet_Frame create() {

            WriteBufferByteBased buffer = new WriteBufferByteBased(deviceContext.getOutputReq().getDataLength());
            PnIoCm_IoCrBlockReqApi api = deviceContext.getOutputReq().getApis().get(0);
            for (PnIoCm_IoCs iocs : api.getIoCss()) {
                PnIoCm_DataUnitIoCs ioc = new PnIoCm_DataUnitIoCs(false, (byte) 0x03, false);
                try {
                    ioc.serialize(buffer);
                } catch (SerializationException e) {
                    throw new RuntimeException(e);
                }
            }

            for (PnIoCm_IoDataObject dataObject : api.getIoDataObjects()) {
                // TODO: Need to specify the datatype length based on the gsd file
                PnIoCm_DataUnitDataObject ioc = new PnIoCm_DataUnitDataObject(
                    new byte[1],
                    new PnIoCm_DataUnitIoCs(false, (byte) 0x03, false),
                    1
                );
                try {
                    ioc.serialize(buffer);
                } catch (SerializationException e) {
                    throw new RuntimeException(e);
                }
            }

            while (buffer.getPos() < deviceContext.getOutputReq().getDataLength()) {
                try {
                    buffer.writeByte((byte) 0x00);
                } catch (SerializationException e) {
                    throw new RuntimeException(e);
                }
            }

            int elapsedTime = (int) ((((System.nanoTime() - startTime)/(MIN_CYCLE_NANO_SEC))) % 65536);

            Ethernet_Frame test = new Ethernet_Frame(
                deviceContext.getMacAddress(),
                deviceContext.getLocalMacAddress(),
                new Ethernet_FramePayload_VirtualLan(
                    VirtualLanPriority.INTERNETWORK_CONTROL,
                    false,
                    0,
                    new Ethernet_FramePayload_PnDcp(
                        new PnDcp_Pdu_RealTimeCyclic(
                            deviceContext.getOutputReq().getFrameId(),
                            new PnIo_CyclicServiceDataUnit(buffer.getBytes(), (short) deviceContext.getOutputReq().getDataLength()),
                            elapsedTime,
                            false,
                            true,
                            true,
                            true,
                            false,
                            true))
                ));
            return test;
        }

        @Override
        public void handle(Ethernet_Frame packet) throws PlcException {
            throw new PlcException("There is no returned packet for cyclic data");
        }
    }
}
