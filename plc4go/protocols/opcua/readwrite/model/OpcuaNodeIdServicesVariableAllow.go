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

// OpcuaNodeIdServicesVariableAllow is an enum
type OpcuaNodeIdServicesVariableAllow int32

type IOpcuaNodeIdServicesVariableAllow interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAllow_AllowNulls OpcuaNodeIdServicesVariableAllow = 3070
)

var OpcuaNodeIdServicesVariableAllowValues []OpcuaNodeIdServicesVariableAllow

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAllowValues = []OpcuaNodeIdServicesVariableAllow{
		OpcuaNodeIdServicesVariableAllow_AllowNulls,
	}
}

func OpcuaNodeIdServicesVariableAllowByValue(value int32) (enum OpcuaNodeIdServicesVariableAllow, ok bool) {
	switch value {
	case 3070:
		return OpcuaNodeIdServicesVariableAllow_AllowNulls, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAllowByName(value string) (enum OpcuaNodeIdServicesVariableAllow, ok bool) {
	switch value {
	case "AllowNulls":
		return OpcuaNodeIdServicesVariableAllow_AllowNulls, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAllowKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAllowValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAllow(structType any) OpcuaNodeIdServicesVariableAllow {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAllow {
		if sOpcuaNodeIdServicesVariableAllow, ok := typ.(OpcuaNodeIdServicesVariableAllow); ok {
			return sOpcuaNodeIdServicesVariableAllow
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAllow) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAllow) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAllowParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAllow, error) {
	return OpcuaNodeIdServicesVariableAllowParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAllowParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAllow, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAllow", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAllow")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAllowByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAllow")
		return OpcuaNodeIdServicesVariableAllow(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAllow) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAllow) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAllow", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableAllow) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAllow) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAllow_AllowNulls:
		return "AllowNulls"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAllow) String() string {
	return e.PLC4XEnumName()
}
