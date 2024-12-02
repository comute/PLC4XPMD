/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.modbus.base.optimizer;

import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.modbus.base.tag.ModbusTagCoil;
import org.apache.plc4x.java.modbus.base.tag.ModbusTagHandler;
import org.apache.plc4x.java.modbus.base.tag.ModbusTagHoldingRegister;
import org.apache.plc4x.java.modbus.tcp.context.ModbusTcpContext;
import org.apache.plc4x.java.modbus.types.ModbusByteOrder;
import org.apache.plc4x.java.spi.connection.PlcTagHandler;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadRequest;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadResponse;
import org.apache.plc4x.java.spi.messages.PlcReader;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcResponseItem;
import org.apache.plc4x.java.spi.optimizer.BaseOptimizer;
import org.apache.plc4x.java.spi.values.PlcRawByteArray;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
import org.apache.plc4x.java.spi.values.PlcBOOL;

/*
* The objective of this test is to evaluate the response order for an array
* of coils that are continuously requested.
* 
* The response pattern is previously defined in the devices under test 
* in order to evaluate their standard response using the frame capture tool.
*
*  Register           Value
*   0x00001           true
*   0x00002           false
*   0x00003           true
*   0x00004           false
*   0x00005           true
*   0x00006           false
*   0x00007           true
*   0x00008           false
*   0x00009           false
*   0x00010           true
*   0x00011           false
*   0x00012           true
*   0x00013           false
*   0x00014           true
*   0x00015           false
*   0x00016           true
*   0x00017           false
*   0x00018           true
*   0x00019           false
*   0x00020           true
*
* This test was generated against the response patterns of the following devices.
*
* . OpenModSim
* . ModbusTools
* . Applicom 
*
*  
*  Rx: 00 48 00 00 00 06 01 01 00 00 00 14
*  Tx: 00 48 00 00 00 06 01 01 03 55 AA 0A  
*                                 ^^ ^^ ^^
* Test parser:
* https://rapidscada.net/modbus/
*
*/
public class ModbusCoilArrayTest {

    @Test
    public void testLittleEndianByteSwap() throws Exception {
        Map<String, String[]> input = new LinkedHashMap<>();
        input.put("variable01",    new String[]{"coil:1:BOOL[20]",               "false"});
        String[] response = {"true","false","true","false","true","false","true","false","false","true","false","true","false","true","false","true","false","true","false","true"}; 

        PlcReader mockPlcReader = Mockito.mock(PlcReader.class);
        PlcTagHandler modbusTagHandler = new ModbusTagHandler();
        PlcReadRequest.Builder builder = new DefaultPlcReadRequest.Builder(mockPlcReader, modbusTagHandler);
        for (String name : input.keySet()) {
            String[] data = input.get(name);
            builder.addTagAddress(name, data[0]);
        }
        PlcReadRequest readRequest = builder.build();

        ModbusOptimizer sut = new ModbusOptimizer();

        ModbusTcpContext mockContext = Mockito.mock(ModbusTcpContext.class);
        
        //TODO: In the optimizer it reads the byte array as BIG_ENDIAN, 
        //      regardless of what is placed here.
        Mockito.when(mockContext.getByteOrder()).thenReturn(ModbusByteOrder.LITTLE_ENDIAN);
        
        Mockito.when(mockContext.getMaxCoilsPerRequest()).thenReturn(2000);
        Mockito.when(mockContext.getMaxRegistersPerRequest()).thenReturn(125);

        ////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        // Do the first part of the optimizer ... split up into multiple requests ...
        ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

        List<PlcReadRequest> optimizedReadRequests = sut.processReadRequest(readRequest, mockContext);

        // Validate the results of this first step.

        Assertions.assertNotNull(optimizedReadRequests);
        Assertions.assertEquals(1, optimizedReadRequests.size());

        // Check the expected first request (Coil)
        PlcReadRequest firstRequest = optimizedReadRequests.get(0);
        Assertions.assertEquals(1, firstRequest.getNumberOfTags());
        Assertions.assertEquals("coils0", firstRequest.getTagNames().stream().findFirst().orElseThrow());
        PlcTag coilsTag = firstRequest.getTag("coils0");
        Assertions.assertInstanceOf(ModbusTagCoil.class, coilsTag);
        ModbusTagCoil coil = (ModbusTagCoil) coilsTag;
        Assertions.assertEquals(0, coil.getAddress());
        Assertions.assertEquals(20, coil.getNumberOfElements());

        // Prepare the results as we got them on the wire.

        Map<PlcReadRequest, BaseOptimizer.SubResponse<PlcReadResponse>> readResponses = new HashMap<>();
        readResponses.put(firstRequest, new BaseOptimizer.SubResponse<>(
            new DefaultPlcReadResponse(firstRequest, Map.of(
                "coils0", new DefaultPlcResponseItem<>(PlcResponseCode.OK, new PlcRawByteArray(Hex.decodeHex("55AA0A")))))));

        ////////////////////////////////////////////////////////////////////////////////////////////////////////////////
        // Process the responses
        ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

        PlcReadResponse readResponse = sut.processReadResponses(readRequest, readResponses, mockContext);

        // Check if there were no invalid items
        List<String> failedTags = readResponse.getTagNames().stream().filter(tagName -> readResponse.getResponseCode(tagName) != PlcResponseCode.OK).collect(Collectors.toList());
        failedTags.forEach(failedTag -> Assertions.fail("Field " + failedTag + "failed."));

        
        var plcResValues = readResponse.getObject("variable01");
        
        List<PlcBOOL> plcValues = (List<PlcBOOL>) plcResValues ;
        // Check if the returned values match the expected ones

        int i = 0;
        for (PlcBOOL plcBool:plcValues) {
            Assertions.assertEquals(response[i], plcBool.getString());
            i++;
        }
    }
    
}
