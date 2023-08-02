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

// DataSetFieldFlags is an enum
type DataSetFieldFlags uint16

type IDataSetFieldFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DataSetFieldFlags_dataSetFieldFlagsNone          DataSetFieldFlags = 0
	DataSetFieldFlags_dataSetFieldFlagsPromotedField DataSetFieldFlags = 1
)

var DataSetFieldFlagsValues []DataSetFieldFlags

func init() {
	_ = errors.New
	DataSetFieldFlagsValues = []DataSetFieldFlags{
		DataSetFieldFlags_dataSetFieldFlagsNone,
		DataSetFieldFlags_dataSetFieldFlagsPromotedField,
	}
}

func DataSetFieldFlagsByValue(value uint16) (enum DataSetFieldFlags, ok bool) {
	switch value {
	case 0:
		return DataSetFieldFlags_dataSetFieldFlagsNone, true
	case 1:
		return DataSetFieldFlags_dataSetFieldFlagsPromotedField, true
	}
	return 0, false
}

func DataSetFieldFlagsByName(value string) (enum DataSetFieldFlags, ok bool) {
	switch value {
	case "dataSetFieldFlagsNone":
		return DataSetFieldFlags_dataSetFieldFlagsNone, true
	case "dataSetFieldFlagsPromotedField":
		return DataSetFieldFlags_dataSetFieldFlagsPromotedField, true
	}
	return 0, false
}

func DataSetFieldFlagsKnows(value uint16) bool {
	for _, typeValue := range DataSetFieldFlagsValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataSetFieldFlags(structType any) DataSetFieldFlags {
	castFunc := func(typ any) DataSetFieldFlags {
		if sDataSetFieldFlags, ok := typ.(DataSetFieldFlags); ok {
			return sDataSetFieldFlags
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataSetFieldFlags) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m DataSetFieldFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSetFieldFlagsParse(ctx context.Context, theBytes []byte) (DataSetFieldFlags, error) {
	return DataSetFieldFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSetFieldFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataSetFieldFlags, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("DataSetFieldFlags", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DataSetFieldFlags")
	}
	if enum, ok := DataSetFieldFlagsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DataSetFieldFlags")
		return DataSetFieldFlags(val), nil
	} else {
		return enum, nil
	}
}

func (e DataSetFieldFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DataSetFieldFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("DataSetFieldFlags", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DataSetFieldFlags) PLC4XEnumName() string {
	switch e {
	case DataSetFieldFlags_dataSetFieldFlagsNone:
		return "dataSetFieldFlagsNone"
	case DataSetFieldFlags_dataSetFieldFlagsPromotedField:
		return "dataSetFieldFlagsPromotedField"
	}
	return ""
}

func (e DataSetFieldFlags) String() string {
	return e.PLC4XEnumName()
}
