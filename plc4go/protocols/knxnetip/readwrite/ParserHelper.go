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

package readwrite

import (
	"context"

	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/protocols/knxnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

type KnxnetipParserHelper struct {
}

func (m KnxnetipParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (any, error) {
	switch typeName {
	case "KnxProperty":
		propertyType, _ := KnxPropertyDataTypeByName(arguments[0])
		dataLengthInBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return KnxPropertyParseWithBuffer(context.Background(), io, propertyType, dataLengthInBytes)
	case "HPAIControlEndpoint":
		return HPAIControlEndpointParseWithBuffer(context.Background(), io)
	case "TunnelingResponseDataBlock":
		return TunnelingResponseDataBlockParseWithBuffer(context.Background(), io)
	case "DeviceDescriptorType2":
		return DeviceDescriptorType2ParseWithBuffer(context.Background(), io)
	case "ChannelInformation":
		return ChannelInformationParseWithBuffer(context.Background(), io)
	case "KnxDatapoint":
		datapointType, _ := KnxDatapointTypeByName(arguments[0])
		return KnxDatapointParseWithBuffer(context.Background(), io, datapointType)
	case "DeviceConfigurationAckDataBlock":
		return DeviceConfigurationAckDataBlockParseWithBuffer(context.Background(), io)
	case "ConnectionRequestInformation":
		return ConnectionRequestInformationParseWithBuffer[ConnectionRequestInformation](context.Background(), io)
	case "Apdu":
		dataLength, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return ApduParseWithBuffer[Apdu](context.Background(), io, dataLength)
	case "HPAIDiscoveryEndpoint":
		return HPAIDiscoveryEndpointParseWithBuffer(context.Background(), io)
	case "ProjectInstallationIdentifier":
		return ProjectInstallationIdentifierParseWithBuffer(context.Background(), io)
	case "ServiceId":
		return ServiceIdParseWithBuffer[ServiceId](context.Background(), io)
	case "HPAIDataEndpoint":
		return HPAIDataEndpointParseWithBuffer(context.Background(), io)
	case "RelativeTimestamp":
		return RelativeTimestampParseWithBuffer(context.Background(), io)
	case "CEMI":
		size, err := utils.StrToUint16(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return CEMIParseWithBuffer[CEMI](context.Background(), io, size)
	case "KnxNetIpMessage":
		return KnxNetIpMessageParseWithBuffer[KnxNetIpMessage](context.Background(), io)
	case "DeviceStatus":
		return DeviceStatusParseWithBuffer(context.Background(), io)
	case "IPAddress":
		return IPAddressParseWithBuffer(context.Background(), io)
	case "GroupObjectDescriptorRealisationTypeB":
		return GroupObjectDescriptorRealisationTypeBParseWithBuffer(context.Background(), io)
	case "CEMIAdditionalInformation":
		return CEMIAdditionalInformationParseWithBuffer[CEMIAdditionalInformation](context.Background(), io)
	case "ComObjectTable":
		firmwareType, _ := FirmwareTypeByName(arguments[0])
		return ComObjectTableParseWithBuffer[ComObjectTable](context.Background(), io, firmwareType)
	case "KnxAddress":
		return KnxAddressParseWithBuffer(context.Background(), io)
	case "ConnectionResponseDataBlock":
		return ConnectionResponseDataBlockParseWithBuffer[ConnectionResponseDataBlock](context.Background(), io)
	case "TunnelingRequestDataBlock":
		return TunnelingRequestDataBlockParseWithBuffer(context.Background(), io)
	case "DIBDeviceInfo":
		return DIBDeviceInfoParseWithBuffer(context.Background(), io)
	case "DeviceConfigurationRequestDataBlock":
		return DeviceConfigurationRequestDataBlockParseWithBuffer(context.Background(), io)
	case "DIBSuppSvcFamilies":
		return DIBSuppSvcFamiliesParseWithBuffer(context.Background(), io)
	case "LDataFrame":
		return LDataFrameParseWithBuffer[LDataFrame](context.Background(), io)
	case "ApduDataExt":
		length, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return ApduDataExtParseWithBuffer[ApduDataExt](context.Background(), io, length)
	case "ApduControl":
		return ApduControlParseWithBuffer[ApduControl](context.Background(), io)
	case "KnxGroupAddress":
		numLevels, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return KnxGroupAddressParseWithBuffer[KnxGroupAddress](context.Background(), io, numLevels)
	case "GroupObjectDescriptorRealisationType6":
		return GroupObjectDescriptorRealisationType6ParseWithBuffer(context.Background(), io)
	case "GroupObjectDescriptorRealisationType7":
		return GroupObjectDescriptorRealisationType7ParseWithBuffer(context.Background(), io)
	case "MACAddress":
		return MACAddressParseWithBuffer(context.Background(), io)
	case "GroupObjectDescriptorRealisationType2":
		return GroupObjectDescriptorRealisationType2ParseWithBuffer(context.Background(), io)
	case "ApduData":
		dataLength, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return ApduDataParseWithBuffer[ApduData](context.Background(), io, dataLength)
	case "GroupObjectDescriptorRealisationType1":
		return GroupObjectDescriptorRealisationType1ParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
