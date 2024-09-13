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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableOpen is an enum
type OpcuaNodeIdServicesVariableOpen int32

type IOpcuaNodeIdServicesVariableOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableOpen_OpenMethodType_InputArguments           OpcuaNodeIdServicesVariableOpen = 11739
	OpcuaNodeIdServicesVariableOpen_OpenMethodType_OutputArguments          OpcuaNodeIdServicesVariableOpen = 11740
	OpcuaNodeIdServicesVariableOpen_OpenFileMode_EnumValues                 OpcuaNodeIdServicesVariableOpen = 11940
	OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_InputArguments  OpcuaNodeIdServicesVariableOpen = 12514
	OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_OutputArguments OpcuaNodeIdServicesVariableOpen = 12515
)

var OpcuaNodeIdServicesVariableOpenValues []OpcuaNodeIdServicesVariableOpen

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOpenValues = []OpcuaNodeIdServicesVariableOpen{
		OpcuaNodeIdServicesVariableOpen_OpenMethodType_InputArguments,
		OpcuaNodeIdServicesVariableOpen_OpenMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableOpen_OpenFileMode_EnumValues,
		OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_InputArguments,
		OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableOpenByValue(value int32) (enum OpcuaNodeIdServicesVariableOpen, ok bool) {
	switch value {
	case 11739:
		return OpcuaNodeIdServicesVariableOpen_OpenMethodType_InputArguments, true
	case 11740:
		return OpcuaNodeIdServicesVariableOpen_OpenMethodType_OutputArguments, true
	case 11940:
		return OpcuaNodeIdServicesVariableOpen_OpenFileMode_EnumValues, true
	case 12514:
		return OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_InputArguments, true
	case 12515:
		return OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOpenByName(value string) (enum OpcuaNodeIdServicesVariableOpen, ok bool) {
	switch value {
	case "OpenMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableOpen_OpenMethodType_InputArguments, true
	case "OpenMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableOpen_OpenMethodType_OutputArguments, true
	case "OpenFileMode_EnumValues":
		return OpcuaNodeIdServicesVariableOpen_OpenFileMode_EnumValues, true
	case "OpenWithMasksMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_InputArguments, true
	case "OpenWithMasksMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOpenKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOpenValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableOpen(structType any) OpcuaNodeIdServicesVariableOpen {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOpen {
		if sOpcuaNodeIdServicesVariableOpen, ok := typ.(OpcuaNodeIdServicesVariableOpen); ok {
			return sOpcuaNodeIdServicesVariableOpen
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOpen) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOpenParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOpen, error) {
	return OpcuaNodeIdServicesVariableOpenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOpenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOpen, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOpen", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOpen")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOpenByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOpen")
		return OpcuaNodeIdServicesVariableOpen(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOpen", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableOpen) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOpen) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOpen_OpenMethodType_InputArguments:
		return "OpenMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableOpen_OpenMethodType_OutputArguments:
		return "OpenMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableOpen_OpenFileMode_EnumValues:
		return "OpenFileMode_EnumValues"
	case OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_InputArguments:
		return "OpenWithMasksMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableOpen_OpenWithMasksMethodType_OutputArguments:
		return "OpenWithMasksMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOpen) String() string {
	return e.PLC4XEnumName()
}
