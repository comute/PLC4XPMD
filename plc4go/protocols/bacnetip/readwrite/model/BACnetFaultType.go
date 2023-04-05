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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetFaultType is an enum
type BACnetFaultType uint8

type IBACnetFaultType interface {
	fmt.Stringer
	utils.Serializable
}

const (
	BACnetFaultType_NONE                  BACnetFaultType = 0
	BACnetFaultType_FAULT_CHARACTERSTRING BACnetFaultType = 1
	BACnetFaultType_FAULT_EXTENDED        BACnetFaultType = 2
	BACnetFaultType_FAULT_LIFE_SAFETY     BACnetFaultType = 3
	BACnetFaultType_FAULT_STATE           BACnetFaultType = 4
	BACnetFaultType_FAULT_STATUS_FLAGS    BACnetFaultType = 5
	BACnetFaultType_FAULT_OUT_OF_RANGE    BACnetFaultType = 6
	BACnetFaultType_FAULT_LISTED          BACnetFaultType = 7
)

var BACnetFaultTypeValues []BACnetFaultType

func init() {
	_ = errors.New
	BACnetFaultTypeValues = []BACnetFaultType{
		BACnetFaultType_NONE,
		BACnetFaultType_FAULT_CHARACTERSTRING,
		BACnetFaultType_FAULT_EXTENDED,
		BACnetFaultType_FAULT_LIFE_SAFETY,
		BACnetFaultType_FAULT_STATE,
		BACnetFaultType_FAULT_STATUS_FLAGS,
		BACnetFaultType_FAULT_OUT_OF_RANGE,
		BACnetFaultType_FAULT_LISTED,
	}
}

func BACnetFaultTypeByValue(value uint8) (enum BACnetFaultType, ok bool) {
	switch value {
	case 0:
		return BACnetFaultType_NONE, true
	case 1:
		return BACnetFaultType_FAULT_CHARACTERSTRING, true
	case 2:
		return BACnetFaultType_FAULT_EXTENDED, true
	case 3:
		return BACnetFaultType_FAULT_LIFE_SAFETY, true
	case 4:
		return BACnetFaultType_FAULT_STATE, true
	case 5:
		return BACnetFaultType_FAULT_STATUS_FLAGS, true
	case 6:
		return BACnetFaultType_FAULT_OUT_OF_RANGE, true
	case 7:
		return BACnetFaultType_FAULT_LISTED, true
	}
	return 0, false
}

func BACnetFaultTypeByName(value string) (enum BACnetFaultType, ok bool) {
	switch value {
	case "NONE":
		return BACnetFaultType_NONE, true
	case "FAULT_CHARACTERSTRING":
		return BACnetFaultType_FAULT_CHARACTERSTRING, true
	case "FAULT_EXTENDED":
		return BACnetFaultType_FAULT_EXTENDED, true
	case "FAULT_LIFE_SAFETY":
		return BACnetFaultType_FAULT_LIFE_SAFETY, true
	case "FAULT_STATE":
		return BACnetFaultType_FAULT_STATE, true
	case "FAULT_STATUS_FLAGS":
		return BACnetFaultType_FAULT_STATUS_FLAGS, true
	case "FAULT_OUT_OF_RANGE":
		return BACnetFaultType_FAULT_OUT_OF_RANGE, true
	case "FAULT_LISTED":
		return BACnetFaultType_FAULT_LISTED, true
	}
	return 0, false
}

func BACnetFaultTypeKnows(value uint8) bool {
	for _, typeValue := range BACnetFaultTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetFaultType(structType interface{}) BACnetFaultType {
	castFunc := func(typ interface{}) BACnetFaultType {
		if sBACnetFaultType, ok := typ.(BACnetFaultType); ok {
			return sBACnetFaultType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetFaultType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetFaultType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultTypeParse(ctx context.Context, theBytes []byte) (BACnetFaultType, error) {
	return BACnetFaultTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultType, error) {
	val, err := readBuffer.ReadUint8("BACnetFaultType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetFaultType")
	}
	if enum, ok := BACnetFaultTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetFaultType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetFaultType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetFaultType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetFaultType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetFaultType) PLC4XEnumName() string {
	switch e {
	case BACnetFaultType_NONE:
		return "NONE"
	case BACnetFaultType_FAULT_CHARACTERSTRING:
		return "FAULT_CHARACTERSTRING"
	case BACnetFaultType_FAULT_EXTENDED:
		return "FAULT_EXTENDED"
	case BACnetFaultType_FAULT_LIFE_SAFETY:
		return "FAULT_LIFE_SAFETY"
	case BACnetFaultType_FAULT_STATE:
		return "FAULT_STATE"
	case BACnetFaultType_FAULT_STATUS_FLAGS:
		return "FAULT_STATUS_FLAGS"
	case BACnetFaultType_FAULT_OUT_OF_RANGE:
		return "FAULT_OUT_OF_RANGE"
	case BACnetFaultType_FAULT_LISTED:
		return "FAULT_LISTED"
	}
	return ""
}

func (e BACnetFaultType) String() string {
	return e.PLC4XEnumName()
}
