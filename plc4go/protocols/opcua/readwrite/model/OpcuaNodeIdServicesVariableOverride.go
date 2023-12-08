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

// OpcuaNodeIdServicesVariableOverride is an enum
type OpcuaNodeIdServicesVariableOverride int32

type IOpcuaNodeIdServicesVariableOverride interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableOverride_OverrideValueHandling_EnumStrings OpcuaNodeIdServicesVariableOverride = 15875
)

var OpcuaNodeIdServicesVariableOverrideValues []OpcuaNodeIdServicesVariableOverride

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOverrideValues = []OpcuaNodeIdServicesVariableOverride{
		OpcuaNodeIdServicesVariableOverride_OverrideValueHandling_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableOverrideByValue(value int32) (enum OpcuaNodeIdServicesVariableOverride, ok bool) {
	switch value {
	case 15875:
		return OpcuaNodeIdServicesVariableOverride_OverrideValueHandling_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOverrideByName(value string) (enum OpcuaNodeIdServicesVariableOverride, ok bool) {
	switch value {
	case "OverrideValueHandling_EnumStrings":
		return OpcuaNodeIdServicesVariableOverride_OverrideValueHandling_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOverrideKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOverrideValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableOverride(structType any) OpcuaNodeIdServicesVariableOverride {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOverride {
		if sOpcuaNodeIdServicesVariableOverride, ok := typ.(OpcuaNodeIdServicesVariableOverride); ok {
			return sOpcuaNodeIdServicesVariableOverride
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOverride) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOverride) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOverrideParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOverride, error) {
	return OpcuaNodeIdServicesVariableOverrideParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOverrideParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOverride, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOverride", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOverride")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOverrideByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOverride")
		return OpcuaNodeIdServicesVariableOverride(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOverride) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOverride) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOverride", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOverride) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOverride_OverrideValueHandling_EnumStrings:
		return "OverrideValueHandling_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOverride) String() string {
	return e.PLC4XEnumName()
}
