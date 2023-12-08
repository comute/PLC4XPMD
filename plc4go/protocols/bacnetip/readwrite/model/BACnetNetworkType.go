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

// BACnetNetworkType is an enum
type BACnetNetworkType uint8

type IBACnetNetworkType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetNetworkType_ETHERNET                 BACnetNetworkType = 0x0
	BACnetNetworkType_ARCNET                   BACnetNetworkType = 0x1
	BACnetNetworkType_MSTP                     BACnetNetworkType = 0x2
	BACnetNetworkType_PTP                      BACnetNetworkType = 0x3
	BACnetNetworkType_LONTALK                  BACnetNetworkType = 0x4
	BACnetNetworkType_IPV4                     BACnetNetworkType = 0x5
	BACnetNetworkType_ZIGBEE                   BACnetNetworkType = 0x6
	BACnetNetworkType_VIRTUAL                  BACnetNetworkType = 0x7
	BACnetNetworkType_REMOVED_NON_BACNET       BACnetNetworkType = 0x8
	BACnetNetworkType_IPV6                     BACnetNetworkType = 0x9
	BACnetNetworkType_SERIAL                   BACnetNetworkType = 0xA
	BACnetNetworkType_VENDOR_PROPRIETARY_VALUE BACnetNetworkType = 0xFF
)

var BACnetNetworkTypeValues []BACnetNetworkType

func init() {
	_ = errors.New
	BACnetNetworkTypeValues = []BACnetNetworkType{
		BACnetNetworkType_ETHERNET,
		BACnetNetworkType_ARCNET,
		BACnetNetworkType_MSTP,
		BACnetNetworkType_PTP,
		BACnetNetworkType_LONTALK,
		BACnetNetworkType_IPV4,
		BACnetNetworkType_ZIGBEE,
		BACnetNetworkType_VIRTUAL,
		BACnetNetworkType_REMOVED_NON_BACNET,
		BACnetNetworkType_IPV6,
		BACnetNetworkType_SERIAL,
		BACnetNetworkType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetNetworkTypeByValue(value uint8) (enum BACnetNetworkType, ok bool) {
	switch value {
	case 0xFF:
		return BACnetNetworkType_VENDOR_PROPRIETARY_VALUE, true
	case 0x0:
		return BACnetNetworkType_ETHERNET, true
	case 0x1:
		return BACnetNetworkType_ARCNET, true
	case 0x2:
		return BACnetNetworkType_MSTP, true
	case 0x3:
		return BACnetNetworkType_PTP, true
	case 0x4:
		return BACnetNetworkType_LONTALK, true
	case 0x5:
		return BACnetNetworkType_IPV4, true
	case 0x6:
		return BACnetNetworkType_ZIGBEE, true
	case 0x7:
		return BACnetNetworkType_VIRTUAL, true
	case 0x8:
		return BACnetNetworkType_REMOVED_NON_BACNET, true
	case 0x9:
		return BACnetNetworkType_IPV6, true
	case 0xA:
		return BACnetNetworkType_SERIAL, true
	}
	return 0, false
}

func BACnetNetworkTypeByName(value string) (enum BACnetNetworkType, ok bool) {
	switch value {
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetNetworkType_VENDOR_PROPRIETARY_VALUE, true
	case "ETHERNET":
		return BACnetNetworkType_ETHERNET, true
	case "ARCNET":
		return BACnetNetworkType_ARCNET, true
	case "MSTP":
		return BACnetNetworkType_MSTP, true
	case "PTP":
		return BACnetNetworkType_PTP, true
	case "LONTALK":
		return BACnetNetworkType_LONTALK, true
	case "IPV4":
		return BACnetNetworkType_IPV4, true
	case "ZIGBEE":
		return BACnetNetworkType_ZIGBEE, true
	case "VIRTUAL":
		return BACnetNetworkType_VIRTUAL, true
	case "REMOVED_NON_BACNET":
		return BACnetNetworkType_REMOVED_NON_BACNET, true
	case "IPV6":
		return BACnetNetworkType_IPV6, true
	case "SERIAL":
		return BACnetNetworkType_SERIAL, true
	}
	return 0, false
}

func BACnetNetworkTypeKnows(value uint8) bool {
	for _, typeValue := range BACnetNetworkTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetNetworkType(structType any) BACnetNetworkType {
	castFunc := func(typ any) BACnetNetworkType {
		if sBACnetNetworkType, ok := typ.(BACnetNetworkType); ok {
			return sBACnetNetworkType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNetworkType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetNetworkType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNetworkTypeParse(ctx context.Context, theBytes []byte) (BACnetNetworkType, error) {
	return BACnetNetworkTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetNetworkTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetNetworkType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetNetworkType")
	}
	if enum, ok := BACnetNetworkTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetNetworkType")
		return BACnetNetworkType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetNetworkType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetNetworkType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetNetworkType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetNetworkType) PLC4XEnumName() string {
	switch e {
	case BACnetNetworkType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetNetworkType_ETHERNET:
		return "ETHERNET"
	case BACnetNetworkType_ARCNET:
		return "ARCNET"
	case BACnetNetworkType_MSTP:
		return "MSTP"
	case BACnetNetworkType_PTP:
		return "PTP"
	case BACnetNetworkType_LONTALK:
		return "LONTALK"
	case BACnetNetworkType_IPV4:
		return "IPV4"
	case BACnetNetworkType_ZIGBEE:
		return "ZIGBEE"
	case BACnetNetworkType_VIRTUAL:
		return "VIRTUAL"
	case BACnetNetworkType_REMOVED_NON_BACNET:
		return "REMOVED_NON_BACNET"
	case BACnetNetworkType_IPV6:
		return "IPV6"
	case BACnetNetworkType_SERIAL:
		return "SERIAL"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetNetworkType) String() string {
	return e.PLC4XEnumName()
}
