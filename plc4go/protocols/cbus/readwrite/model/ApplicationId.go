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

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ApplicationId is an enum
type ApplicationId uint8

type IApplicationId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ApplicationId_RESERVED                           ApplicationId = 0x00
	ApplicationId_FREE_USAGE                         ApplicationId = 0x01
	ApplicationId_TEMPERATURE_BROADCAST              ApplicationId = 0x02
	ApplicationId_ROOM_CONTROL_SYSTEM                ApplicationId = 0x03
	ApplicationId_LIGHTING                           ApplicationId = 0x04
	ApplicationId_VENTILATION                        ApplicationId = 0x05
	ApplicationId_IRRIGATION_CONTROL                 ApplicationId = 0x06
	ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL ApplicationId = 0x07
	ApplicationId_HEATING                            ApplicationId = 0x08
	ApplicationId_AIR_CONDITIONING                   ApplicationId = 0x09
	ApplicationId_TRIGGER_CONTROL                    ApplicationId = 0x0A
	ApplicationId_ENABLE_CONTROL                     ApplicationId = 0x0B
	ApplicationId_AUDIO_AND_VIDEO                    ApplicationId = 0x0C
	ApplicationId_SECURITY                           ApplicationId = 0x0D
	ApplicationId_METERING                           ApplicationId = 0x0E
	ApplicationId_ACCESS_CONTROL                     ApplicationId = 0x0F
	ApplicationId_CLOCK_AND_TIMEKEEPING              ApplicationId = 0x10
	ApplicationId_TELEPHONY_STATUS_AND_CONTROL       ApplicationId = 0x11
	ApplicationId_MEASUREMENT                        ApplicationId = 0x12
	ApplicationId_TESTING                            ApplicationId = 0x13
	ApplicationId_MEDIA_TRANSPORT_CONTROL            ApplicationId = 0x14
	ApplicationId_ERROR_REPORTING                    ApplicationId = 0x15
	ApplicationId_HVAC_ACTUATOR                      ApplicationId = 0x16
	ApplicationId_INFO_MESSAGES                      ApplicationId = 0x17
	ApplicationId_NETWORK_CONTROL                    ApplicationId = 0x18
)

var ApplicationIdValues []ApplicationId

func init() {
	_ = errors.New
	ApplicationIdValues = []ApplicationId{
		ApplicationId_RESERVED,
		ApplicationId_FREE_USAGE,
		ApplicationId_TEMPERATURE_BROADCAST,
		ApplicationId_ROOM_CONTROL_SYSTEM,
		ApplicationId_LIGHTING,
		ApplicationId_VENTILATION,
		ApplicationId_IRRIGATION_CONTROL,
		ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL,
		ApplicationId_HEATING,
		ApplicationId_AIR_CONDITIONING,
		ApplicationId_TRIGGER_CONTROL,
		ApplicationId_ENABLE_CONTROL,
		ApplicationId_AUDIO_AND_VIDEO,
		ApplicationId_SECURITY,
		ApplicationId_METERING,
		ApplicationId_ACCESS_CONTROL,
		ApplicationId_CLOCK_AND_TIMEKEEPING,
		ApplicationId_TELEPHONY_STATUS_AND_CONTROL,
		ApplicationId_MEASUREMENT,
		ApplicationId_TESTING,
		ApplicationId_MEDIA_TRANSPORT_CONTROL,
		ApplicationId_ERROR_REPORTING,
		ApplicationId_HVAC_ACTUATOR,
		ApplicationId_INFO_MESSAGES,
		ApplicationId_NETWORK_CONTROL,
	}
}

func ApplicationIdByValue(value uint8) (enum ApplicationId, ok bool) {
	switch value {
	case 0x00:
		return ApplicationId_RESERVED, true
	case 0x01:
		return ApplicationId_FREE_USAGE, true
	case 0x02:
		return ApplicationId_TEMPERATURE_BROADCAST, true
	case 0x03:
		return ApplicationId_ROOM_CONTROL_SYSTEM, true
	case 0x04:
		return ApplicationId_LIGHTING, true
	case 0x05:
		return ApplicationId_VENTILATION, true
	case 0x06:
		return ApplicationId_IRRIGATION_CONTROL, true
	case 0x07:
		return ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL, true
	case 0x08:
		return ApplicationId_HEATING, true
	case 0x09:
		return ApplicationId_AIR_CONDITIONING, true
	case 0x0A:
		return ApplicationId_TRIGGER_CONTROL, true
	case 0x0B:
		return ApplicationId_ENABLE_CONTROL, true
	case 0x0C:
		return ApplicationId_AUDIO_AND_VIDEO, true
	case 0x0D:
		return ApplicationId_SECURITY, true
	case 0x0E:
		return ApplicationId_METERING, true
	case 0x0F:
		return ApplicationId_ACCESS_CONTROL, true
	case 0x10:
		return ApplicationId_CLOCK_AND_TIMEKEEPING, true
	case 0x11:
		return ApplicationId_TELEPHONY_STATUS_AND_CONTROL, true
	case 0x12:
		return ApplicationId_MEASUREMENT, true
	case 0x13:
		return ApplicationId_TESTING, true
	case 0x14:
		return ApplicationId_MEDIA_TRANSPORT_CONTROL, true
	case 0x15:
		return ApplicationId_ERROR_REPORTING, true
	case 0x16:
		return ApplicationId_HVAC_ACTUATOR, true
	case 0x17:
		return ApplicationId_INFO_MESSAGES, true
	case 0x18:
		return ApplicationId_NETWORK_CONTROL, true
	}
	return 0, false
}

func ApplicationIdByName(value string) (enum ApplicationId, ok bool) {
	switch value {
	case "RESERVED":
		return ApplicationId_RESERVED, true
	case "FREE_USAGE":
		return ApplicationId_FREE_USAGE, true
	case "TEMPERATURE_BROADCAST":
		return ApplicationId_TEMPERATURE_BROADCAST, true
	case "ROOM_CONTROL_SYSTEM":
		return ApplicationId_ROOM_CONTROL_SYSTEM, true
	case "LIGHTING":
		return ApplicationId_LIGHTING, true
	case "VENTILATION":
		return ApplicationId_VENTILATION, true
	case "IRRIGATION_CONTROL":
		return ApplicationId_IRRIGATION_CONTROL, true
	case "POOLS_SPAS_PONDS_FOUNTAINS_CONTROL":
		return ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL, true
	case "HEATING":
		return ApplicationId_HEATING, true
	case "AIR_CONDITIONING":
		return ApplicationId_AIR_CONDITIONING, true
	case "TRIGGER_CONTROL":
		return ApplicationId_TRIGGER_CONTROL, true
	case "ENABLE_CONTROL":
		return ApplicationId_ENABLE_CONTROL, true
	case "AUDIO_AND_VIDEO":
		return ApplicationId_AUDIO_AND_VIDEO, true
	case "SECURITY":
		return ApplicationId_SECURITY, true
	case "METERING":
		return ApplicationId_METERING, true
	case "ACCESS_CONTROL":
		return ApplicationId_ACCESS_CONTROL, true
	case "CLOCK_AND_TIMEKEEPING":
		return ApplicationId_CLOCK_AND_TIMEKEEPING, true
	case "TELEPHONY_STATUS_AND_CONTROL":
		return ApplicationId_TELEPHONY_STATUS_AND_CONTROL, true
	case "MEASUREMENT":
		return ApplicationId_MEASUREMENT, true
	case "TESTING":
		return ApplicationId_TESTING, true
	case "MEDIA_TRANSPORT_CONTROL":
		return ApplicationId_MEDIA_TRANSPORT_CONTROL, true
	case "ERROR_REPORTING":
		return ApplicationId_ERROR_REPORTING, true
	case "HVAC_ACTUATOR":
		return ApplicationId_HVAC_ACTUATOR, true
	case "INFO_MESSAGES":
		return ApplicationId_INFO_MESSAGES, true
	case "NETWORK_CONTROL":
		return ApplicationId_NETWORK_CONTROL, true
	}
	return 0, false
}

func ApplicationIdKnows(value uint8) bool {
	for _, typeValue := range ApplicationIdValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastApplicationId(structType any) ApplicationId {
	castFunc := func(typ any) ApplicationId {
		if sApplicationId, ok := typ.(ApplicationId); ok {
			return sApplicationId
		}
		return 0
	}
	return castFunc(structType)
}

func (m ApplicationId) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ApplicationId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApplicationIdParse(ctx context.Context, theBytes []byte) (ApplicationId, error) {
	return ApplicationIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ApplicationIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationId, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ApplicationId", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ApplicationId")
	}
	if enum, ok := ApplicationIdByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ApplicationId")
		return ApplicationId(val), nil
	} else {
		return enum, nil
	}
}

func (e ApplicationId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ApplicationId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ApplicationId", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ApplicationId) PLC4XEnumName() string {
	switch e {
	case ApplicationId_RESERVED:
		return "RESERVED"
	case ApplicationId_FREE_USAGE:
		return "FREE_USAGE"
	case ApplicationId_TEMPERATURE_BROADCAST:
		return "TEMPERATURE_BROADCAST"
	case ApplicationId_ROOM_CONTROL_SYSTEM:
		return "ROOM_CONTROL_SYSTEM"
	case ApplicationId_LIGHTING:
		return "LIGHTING"
	case ApplicationId_VENTILATION:
		return "VENTILATION"
	case ApplicationId_IRRIGATION_CONTROL:
		return "IRRIGATION_CONTROL"
	case ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL:
		return "POOLS_SPAS_PONDS_FOUNTAINS_CONTROL"
	case ApplicationId_HEATING:
		return "HEATING"
	case ApplicationId_AIR_CONDITIONING:
		return "AIR_CONDITIONING"
	case ApplicationId_TRIGGER_CONTROL:
		return "TRIGGER_CONTROL"
	case ApplicationId_ENABLE_CONTROL:
		return "ENABLE_CONTROL"
	case ApplicationId_AUDIO_AND_VIDEO:
		return "AUDIO_AND_VIDEO"
	case ApplicationId_SECURITY:
		return "SECURITY"
	case ApplicationId_METERING:
		return "METERING"
	case ApplicationId_ACCESS_CONTROL:
		return "ACCESS_CONTROL"
	case ApplicationId_CLOCK_AND_TIMEKEEPING:
		return "CLOCK_AND_TIMEKEEPING"
	case ApplicationId_TELEPHONY_STATUS_AND_CONTROL:
		return "TELEPHONY_STATUS_AND_CONTROL"
	case ApplicationId_MEASUREMENT:
		return "MEASUREMENT"
	case ApplicationId_TESTING:
		return "TESTING"
	case ApplicationId_MEDIA_TRANSPORT_CONTROL:
		return "MEDIA_TRANSPORT_CONTROL"
	case ApplicationId_ERROR_REPORTING:
		return "ERROR_REPORTING"
	case ApplicationId_HVAC_ACTUATOR:
		return "HVAC_ACTUATOR"
	case ApplicationId_INFO_MESSAGES:
		return "INFO_MESSAGES"
	case ApplicationId_NETWORK_CONTROL:
		return "NETWORK_CONTROL"
	}
	return ""
}

func (e ApplicationId) String() string {
	return e.PLC4XEnumName()
}
