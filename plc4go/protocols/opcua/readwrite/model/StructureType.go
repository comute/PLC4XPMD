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

// StructureType is an enum
type StructureType uint32

type IStructureType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	StructureType_structureTypeStructure                   StructureType = 0
	StructureType_structureTypeStructureWithOptionalFields StructureType = 1
	StructureType_structureTypeUnion                       StructureType = 2
	StructureType_structureTypeStructureWithSubtypedValues StructureType = 3
	StructureType_structureTypeUnionWithSubtypedValues     StructureType = 4
)

var StructureTypeValues []StructureType

func init() {
	_ = errors.New
	StructureTypeValues = []StructureType{
		StructureType_structureTypeStructure,
		StructureType_structureTypeStructureWithOptionalFields,
		StructureType_structureTypeUnion,
		StructureType_structureTypeStructureWithSubtypedValues,
		StructureType_structureTypeUnionWithSubtypedValues,
	}
}

func StructureTypeByValue(value uint32) (enum StructureType, ok bool) {
	switch value {
	case 0:
		return StructureType_structureTypeStructure, true
	case 1:
		return StructureType_structureTypeStructureWithOptionalFields, true
	case 2:
		return StructureType_structureTypeUnion, true
	case 3:
		return StructureType_structureTypeStructureWithSubtypedValues, true
	case 4:
		return StructureType_structureTypeUnionWithSubtypedValues, true
	}
	return 0, false
}

func StructureTypeByName(value string) (enum StructureType, ok bool) {
	switch value {
	case "structureTypeStructure":
		return StructureType_structureTypeStructure, true
	case "structureTypeStructureWithOptionalFields":
		return StructureType_structureTypeStructureWithOptionalFields, true
	case "structureTypeUnion":
		return StructureType_structureTypeUnion, true
	case "structureTypeStructureWithSubtypedValues":
		return StructureType_structureTypeStructureWithSubtypedValues, true
	case "structureTypeUnionWithSubtypedValues":
		return StructureType_structureTypeUnionWithSubtypedValues, true
	}
	return 0, false
}

func StructureTypeKnows(value uint32) bool {
	for _, typeValue := range StructureTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastStructureType(structType any) StructureType {
	castFunc := func(typ any) StructureType {
		if sStructureType, ok := typ.(StructureType); ok {
			return sStructureType
		}
		return 0
	}
	return castFunc(structType)
}

func (m StructureType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m StructureType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StructureTypeParse(ctx context.Context, theBytes []byte) (StructureType, error) {
	return StructureTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StructureTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StructureType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("StructureType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading StructureType")
	}
	if enum, ok := StructureTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for StructureType")
		return StructureType(val), nil
	} else {
		return enum, nil
	}
}

func (e StructureType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e StructureType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("StructureType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e StructureType) PLC4XEnumName() string {
	switch e {
	case StructureType_structureTypeStructure:
		return "structureTypeStructure"
	case StructureType_structureTypeStructureWithOptionalFields:
		return "structureTypeStructureWithOptionalFields"
	case StructureType_structureTypeUnion:
		return "structureTypeUnion"
	case StructureType_structureTypeStructureWithSubtypedValues:
		return "structureTypeStructureWithSubtypedValues"
	case StructureType_structureTypeUnionWithSubtypedValues:
		return "structureTypeUnionWithSubtypedValues"
	}
	return ""
}

func (e StructureType) String() string {
	return e.PLC4XEnumName()
}
