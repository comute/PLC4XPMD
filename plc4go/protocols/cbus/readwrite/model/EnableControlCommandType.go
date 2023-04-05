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

// EnableControlCommandType is an enum
type EnableControlCommandType uint8

type IEnableControlCommandType interface {
	fmt.Stringer
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	EnableControlCommandType_SET_NETWORK_VARIABLE EnableControlCommandType = 0x00
)

var EnableControlCommandTypeValues []EnableControlCommandType

func init() {
	_ = errors.New
	EnableControlCommandTypeValues = []EnableControlCommandType{
		EnableControlCommandType_SET_NETWORK_VARIABLE,
	}
}

func (e EnableControlCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 1
		}
	default:
		{
			return 0
		}
	}
}

func EnableControlCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (EnableControlCommandType, error) {
	for _, sizeValue := range EnableControlCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func EnableControlCommandTypeByValue(value uint8) (enum EnableControlCommandType, ok bool) {
	switch value {
	case 0x00:
		return EnableControlCommandType_SET_NETWORK_VARIABLE, true
	}
	return 0, false
}

func EnableControlCommandTypeByName(value string) (enum EnableControlCommandType, ok bool) {
	switch value {
	case "SET_NETWORK_VARIABLE":
		return EnableControlCommandType_SET_NETWORK_VARIABLE, true
	}
	return 0, false
}

func EnableControlCommandTypeKnows(value uint8) bool {
	for _, typeValue := range EnableControlCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastEnableControlCommandType(structType interface{}) EnableControlCommandType {
	castFunc := func(typ interface{}) EnableControlCommandType {
		if sEnableControlCommandType, ok := typ.(EnableControlCommandType); ok {
			return sEnableControlCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m EnableControlCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m EnableControlCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EnableControlCommandTypeParse(ctx context.Context, theBytes []byte) (EnableControlCommandType, error) {
	return EnableControlCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EnableControlCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EnableControlCommandType, error) {
	val, err := readBuffer.ReadUint8("EnableControlCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading EnableControlCommandType")
	}
	if enum, ok := EnableControlCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return EnableControlCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e EnableControlCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e EnableControlCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("EnableControlCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e EnableControlCommandType) PLC4XEnumName() string {
	switch e {
	case EnableControlCommandType_SET_NETWORK_VARIABLE:
		return "SET_NETWORK_VARIABLE"
	}
	return ""
}

func (e EnableControlCommandType) String() string {
	return e.PLC4XEnumName()
}
