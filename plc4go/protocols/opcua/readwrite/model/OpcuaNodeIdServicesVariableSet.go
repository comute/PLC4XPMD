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

// OpcuaNodeIdServicesVariableSet is an enum
type OpcuaNodeIdServicesVariableSet int32

type IOpcuaNodeIdServicesVariableSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableSet_SetPositionMethodType_InputArguments             OpcuaNodeIdServicesVariableSet = 11752
	OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_InputArguments  OpcuaNodeIdServicesVariableSet = 12753
	OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_OutputArguments OpcuaNodeIdServicesVariableSet = 12754
	OpcuaNodeIdServicesVariableSet_SetSecurityKeysMethodType_InputArguments         OpcuaNodeIdServicesVariableSet = 17299
	OpcuaNodeIdServicesVariableSet_SetRegistrarEndpointsMethodType_InputArguments   OpcuaNodeIdServicesVariableSet = 25730
)

var OpcuaNodeIdServicesVariableSetValues []OpcuaNodeIdServicesVariableSet

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableSetValues = []OpcuaNodeIdServicesVariableSet{
		OpcuaNodeIdServicesVariableSet_SetPositionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_InputArguments,
		OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableSet_SetSecurityKeysMethodType_InputArguments,
		OpcuaNodeIdServicesVariableSet_SetRegistrarEndpointsMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableSetByValue(value int32) (enum OpcuaNodeIdServicesVariableSet, ok bool) {
	switch value {
	case 11752:
		return OpcuaNodeIdServicesVariableSet_SetPositionMethodType_InputArguments, true
	case 12753:
		return OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_InputArguments, true
	case 12754:
		return OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_OutputArguments, true
	case 17299:
		return OpcuaNodeIdServicesVariableSet_SetSecurityKeysMethodType_InputArguments, true
	case 25730:
		return OpcuaNodeIdServicesVariableSet_SetRegistrarEndpointsMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSetByName(value string) (enum OpcuaNodeIdServicesVariableSet, ok bool) {
	switch value {
	case "SetPositionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableSet_SetPositionMethodType_InputArguments, true
	case "SetSubscriptionDurableMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_InputArguments, true
	case "SetSubscriptionDurableMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_OutputArguments, true
	case "SetSecurityKeysMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableSet_SetSecurityKeysMethodType_InputArguments, true
	case "SetRegistrarEndpointsMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableSet_SetRegistrarEndpointsMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableSetKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableSetValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableSet(structType any) OpcuaNodeIdServicesVariableSet {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableSet {
		if sOpcuaNodeIdServicesVariableSet, ok := typ.(OpcuaNodeIdServicesVariableSet); ok {
			return sOpcuaNodeIdServicesVariableSet
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableSet) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableSetParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableSet, error) {
	return OpcuaNodeIdServicesVariableSetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableSetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableSet, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableSet", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableSet")
	}
	if enum, ok := OpcuaNodeIdServicesVariableSetByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableSet")
		return OpcuaNodeIdServicesVariableSet(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableSet", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableSet) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableSet_SetPositionMethodType_InputArguments:
		return "SetPositionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_InputArguments:
		return "SetSubscriptionDurableMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableSet_SetSubscriptionDurableMethodType_OutputArguments:
		return "SetSubscriptionDurableMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableSet_SetSecurityKeysMethodType_InputArguments:
		return "SetSecurityKeysMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableSet_SetRegistrarEndpointsMethodType_InputArguments:
		return "SetRegistrarEndpointsMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableSet) String() string {
	return e.PLC4XEnumName()
}
