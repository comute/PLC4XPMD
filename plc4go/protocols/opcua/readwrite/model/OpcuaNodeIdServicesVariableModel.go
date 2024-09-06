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

// OpcuaNodeIdServicesVariableModel is an enum
type OpcuaNodeIdServicesVariableModel int32

type IOpcuaNodeIdServicesVariableModel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableModel_ModelChangeStructureVerbMask_EnumValues OpcuaNodeIdServicesVariableModel = 11942
)

var OpcuaNodeIdServicesVariableModelValues []OpcuaNodeIdServicesVariableModel

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableModelValues = []OpcuaNodeIdServicesVariableModel{
		OpcuaNodeIdServicesVariableModel_ModelChangeStructureVerbMask_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableModelByValue(value int32) (enum OpcuaNodeIdServicesVariableModel, ok bool) {
	switch value {
	case 11942:
		return OpcuaNodeIdServicesVariableModel_ModelChangeStructureVerbMask_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableModelByName(value string) (enum OpcuaNodeIdServicesVariableModel, ok bool) {
	switch value {
	case "ModelChangeStructureVerbMask_EnumValues":
		return OpcuaNodeIdServicesVariableModel_ModelChangeStructureVerbMask_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableModelKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableModelValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableModel(structType any) OpcuaNodeIdServicesVariableModel {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableModel {
		if sOpcuaNodeIdServicesVariableModel, ok := typ.(OpcuaNodeIdServicesVariableModel); ok {
			return sOpcuaNodeIdServicesVariableModel
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableModel) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableModel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableModelParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableModel, error) {
	return OpcuaNodeIdServicesVariableModelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableModelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableModel, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableModel", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableModel")
	}
	if enum, ok := OpcuaNodeIdServicesVariableModelByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableModel")
		return OpcuaNodeIdServicesVariableModel(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableModel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableModel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableModel", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableModel) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableModel) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableModel_ModelChangeStructureVerbMask_EnumValues:
		return "ModelChangeStructureVerbMask_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableModel) String() string {
	return e.PLC4XEnumName()
}
