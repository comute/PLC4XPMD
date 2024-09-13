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

// OpcuaNodeIdServicesVariableIdentity is an enum
type OpcuaNodeIdServicesVariableIdentity int32

type IOpcuaNodeIdServicesVariableIdentity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableIdentity_IdentityCriteriaType_EnumValues OpcuaNodeIdServicesVariableIdentity = 15633
)

var OpcuaNodeIdServicesVariableIdentityValues []OpcuaNodeIdServicesVariableIdentity

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableIdentityValues = []OpcuaNodeIdServicesVariableIdentity{
		OpcuaNodeIdServicesVariableIdentity_IdentityCriteriaType_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableIdentityByValue(value int32) (enum OpcuaNodeIdServicesVariableIdentity, ok bool) {
	switch value {
	case 15633:
		return OpcuaNodeIdServicesVariableIdentity_IdentityCriteriaType_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIdentityByName(value string) (enum OpcuaNodeIdServicesVariableIdentity, ok bool) {
	switch value {
	case "IdentityCriteriaType_EnumValues":
		return OpcuaNodeIdServicesVariableIdentity_IdentityCriteriaType_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIdentityKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableIdentityValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableIdentity(structType any) OpcuaNodeIdServicesVariableIdentity {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableIdentity {
		if sOpcuaNodeIdServicesVariableIdentity, ok := typ.(OpcuaNodeIdServicesVariableIdentity); ok {
			return sOpcuaNodeIdServicesVariableIdentity
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableIdentity) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableIdentity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableIdentityParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableIdentity, error) {
	return OpcuaNodeIdServicesVariableIdentityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableIdentityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableIdentity, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableIdentity", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableIdentity")
	}
	if enum, ok := OpcuaNodeIdServicesVariableIdentityByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableIdentity")
		return OpcuaNodeIdServicesVariableIdentity(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableIdentity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableIdentity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableIdentity", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableIdentity) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableIdentity) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableIdentity_IdentityCriteriaType_EnumValues:
		return "IdentityCriteriaType_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableIdentity) String() string {
	return e.PLC4XEnumName()
}
