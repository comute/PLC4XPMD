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

// AdsDataType is an enum
type AdsDataType int8

type IAdsDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint16
	PlcValueType() PlcValueType
}

const (
	AdsDataType_BOOL          AdsDataType = 0x01
	AdsDataType_BIT           AdsDataType = 0x02
	AdsDataType_BIT8          AdsDataType = 0x03
	AdsDataType_BYTE          AdsDataType = 0x04
	AdsDataType_BITARR8       AdsDataType = 0x05
	AdsDataType_WORD          AdsDataType = 0x06
	AdsDataType_BITARR16      AdsDataType = 0x07
	AdsDataType_DWORD         AdsDataType = 0x08
	AdsDataType_BITARR32      AdsDataType = 0x09
	AdsDataType_SINT          AdsDataType = 0x0A
	AdsDataType_INT8          AdsDataType = 0x0B
	AdsDataType_USINT         AdsDataType = 0x0C
	AdsDataType_UINT8         AdsDataType = 0x0D
	AdsDataType_INT           AdsDataType = 0x0E
	AdsDataType_INT16         AdsDataType = 0x0F
	AdsDataType_UINT          AdsDataType = 0x10
	AdsDataType_UINT16        AdsDataType = 0x11
	AdsDataType_DINT          AdsDataType = 0x12
	AdsDataType_INT32         AdsDataType = 0x13
	AdsDataType_UDINT         AdsDataType = 0x14
	AdsDataType_UINT32        AdsDataType = 0x15
	AdsDataType_LINT          AdsDataType = 0x16
	AdsDataType_INT64         AdsDataType = 0x17
	AdsDataType_ULINT         AdsDataType = 0x18
	AdsDataType_UINT64        AdsDataType = 0x19
	AdsDataType_REAL          AdsDataType = 0x1A
	AdsDataType_FLOAT         AdsDataType = 0x1B
	AdsDataType_LREAL         AdsDataType = 0x1C
	AdsDataType_DOUBLE        AdsDataType = 0x1D
	AdsDataType_CHAR          AdsDataType = 0x1E
	AdsDataType_WCHAR         AdsDataType = 0x1F
	AdsDataType_STRING        AdsDataType = 0x20
	AdsDataType_WSTRING       AdsDataType = 0x21
	AdsDataType_TIME          AdsDataType = 0x22
	AdsDataType_LTIME         AdsDataType = 0x23
	AdsDataType_DATE          AdsDataType = 0x24
	AdsDataType_TIME_OF_DAY   AdsDataType = 0x25
	AdsDataType_TOD           AdsDataType = 0x26
	AdsDataType_DATE_AND_TIME AdsDataType = 0x27
	AdsDataType_DT            AdsDataType = 0x28
)

var AdsDataTypeValues []AdsDataType

func init() {
	_ = errors.New
	AdsDataTypeValues = []AdsDataType{
		AdsDataType_BOOL,
		AdsDataType_BIT,
		AdsDataType_BIT8,
		AdsDataType_BYTE,
		AdsDataType_BITARR8,
		AdsDataType_WORD,
		AdsDataType_BITARR16,
		AdsDataType_DWORD,
		AdsDataType_BITARR32,
		AdsDataType_SINT,
		AdsDataType_INT8,
		AdsDataType_USINT,
		AdsDataType_UINT8,
		AdsDataType_INT,
		AdsDataType_INT16,
		AdsDataType_UINT,
		AdsDataType_UINT16,
		AdsDataType_DINT,
		AdsDataType_INT32,
		AdsDataType_UDINT,
		AdsDataType_UINT32,
		AdsDataType_LINT,
		AdsDataType_INT64,
		AdsDataType_ULINT,
		AdsDataType_UINT64,
		AdsDataType_REAL,
		AdsDataType_FLOAT,
		AdsDataType_LREAL,
		AdsDataType_DOUBLE,
		AdsDataType_CHAR,
		AdsDataType_WCHAR,
		AdsDataType_STRING,
		AdsDataType_WSTRING,
		AdsDataType_TIME,
		AdsDataType_LTIME,
		AdsDataType_DATE,
		AdsDataType_TIME_OF_DAY,
		AdsDataType_TOD,
		AdsDataType_DATE_AND_TIME,
		AdsDataType_DT,
	}
}

func (e AdsDataType) NumBytes() uint16 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 1
		}
	case 0x03:
		{ /* '0x03' */
			return 1
		}
	case 0x04:
		{ /* '0x04' */
			return 1
		}
	case 0x05:
		{ /* '0x05' */
			return 1
		}
	case 0x06:
		{ /* '0x06' */
			return 2
		}
	case 0x07:
		{ /* '0x07' */
			return 2
		}
	case 0x08:
		{ /* '0x08' */
			return 4
		}
	case 0x09:
		{ /* '0x09' */
			return 4
		}
	case 0x0A:
		{ /* '0x0A' */
			return 1
		}
	case 0x0B:
		{ /* '0x0B' */
			return 1
		}
	case 0x0C:
		{ /* '0x0C' */
			return 1
		}
	case 0x0D:
		{ /* '0x0D' */
			return 1
		}
	case 0x0E:
		{ /* '0x0E' */
			return 2
		}
	case 0x0F:
		{ /* '0x0F' */
			return 2
		}
	case 0x10:
		{ /* '0x10' */
			return 2
		}
	case 0x11:
		{ /* '0x11' */
			return 2
		}
	case 0x12:
		{ /* '0x12' */
			return 4
		}
	case 0x13:
		{ /* '0x13' */
			return 4
		}
	case 0x14:
		{ /* '0x14' */
			return 4
		}
	case 0x15:
		{ /* '0x15' */
			return 4
		}
	case 0x16:
		{ /* '0x16' */
			return 8
		}
	case 0x17:
		{ /* '0x17' */
			return 8
		}
	case 0x18:
		{ /* '0x18' */
			return 8
		}
	case 0x19:
		{ /* '0x19' */
			return 8
		}
	case 0x1A:
		{ /* '0x1A' */
			return 4
		}
	case 0x1B:
		{ /* '0x1B' */
			return 4
		}
	case 0x1C:
		{ /* '0x1C' */
			return 8
		}
	case 0x1D:
		{ /* '0x1D' */
			return 8
		}
	case 0x1E:
		{ /* '0x1E' */
			return 1
		}
	case 0x1F:
		{ /* '0x1F' */
			return 2
		}
	case 0x20:
		{ /* '0x20' */
			return 256
		}
	case 0x21:
		{ /* '0x21' */
			return 512
		}
	case 0x22:
		{ /* '0x22' */
			return 4
		}
	case 0x23:
		{ /* '0x23' */
			return 8
		}
	case 0x24:
		{ /* '0x24' */
			return 4
		}
	case 0x25:
		{ /* '0x25' */
			return 4
		}
	case 0x26:
		{ /* '0x26' */
			return 4
		}
	case 0x27:
		{ /* '0x27' */
			return 4
		}
	case 0x28:
		{ /* '0x28' */
			return 4
		}
	default:
		{
			return 0
		}
	}
}

func AdsDataTypeFirstEnumForFieldNumBytes(value uint16) (AdsDataType, error) {
	for _, sizeValue := range AdsDataTypeValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e AdsDataType) PlcValueType() PlcValueType {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return PlcValueType_BOOL
		}
	case 0x02:
		{ /* '0x02' */
			return PlcValueType_BOOL
		}
	case 0x03:
		{ /* '0x03' */
			return PlcValueType_BYTE
		}
	case 0x04:
		{ /* '0x04' */
			return PlcValueType_BYTE
		}
	case 0x05:
		{ /* '0x05' */
			return PlcValueType_BYTE
		}
	case 0x06:
		{ /* '0x06' */
			return PlcValueType_WORD
		}
	case 0x07:
		{ /* '0x07' */
			return PlcValueType_WORD
		}
	case 0x08:
		{ /* '0x08' */
			return PlcValueType_DWORD
		}
	case 0x09:
		{ /* '0x09' */
			return PlcValueType_DWORD
		}
	case 0x0A:
		{ /* '0x0A' */
			return PlcValueType_SINT
		}
	case 0x0B:
		{ /* '0x0B' */
			return PlcValueType_SINT
		}
	case 0x0C:
		{ /* '0x0C' */
			return PlcValueType_USINT
		}
	case 0x0D:
		{ /* '0x0D' */
			return PlcValueType_USINT
		}
	case 0x0E:
		{ /* '0x0E' */
			return PlcValueType_INT
		}
	case 0x0F:
		{ /* '0x0F' */
			return PlcValueType_INT
		}
	case 0x10:
		{ /* '0x10' */
			return PlcValueType_UINT
		}
	case 0x11:
		{ /* '0x11' */
			return PlcValueType_UINT
		}
	case 0x12:
		{ /* '0x12' */
			return PlcValueType_DINT
		}
	case 0x13:
		{ /* '0x13' */
			return PlcValueType_DINT
		}
	case 0x14:
		{ /* '0x14' */
			return PlcValueType_UDINT
		}
	case 0x15:
		{ /* '0x15' */
			return PlcValueType_UDINT
		}
	case 0x16:
		{ /* '0x16' */
			return PlcValueType_LINT
		}
	case 0x17:
		{ /* '0x17' */
			return PlcValueType_LINT
		}
	case 0x18:
		{ /* '0x18' */
			return PlcValueType_ULINT
		}
	case 0x19:
		{ /* '0x19' */
			return PlcValueType_ULINT
		}
	case 0x1A:
		{ /* '0x1A' */
			return PlcValueType_REAL
		}
	case 0x1B:
		{ /* '0x1B' */
			return PlcValueType_REAL
		}
	case 0x1C:
		{ /* '0x1C' */
			return PlcValueType_LREAL
		}
	case 0x1D:
		{ /* '0x1D' */
			return PlcValueType_LREAL
		}
	case 0x1E:
		{ /* '0x1E' */
			return PlcValueType_CHAR
		}
	case 0x1F:
		{ /* '0x1F' */
			return PlcValueType_WCHAR
		}
	case 0x20:
		{ /* '0x20' */
			return PlcValueType_STRING
		}
	case 0x21:
		{ /* '0x21' */
			return PlcValueType_WSTRING
		}
	case 0x22:
		{ /* '0x22' */
			return PlcValueType_TIME
		}
	case 0x23:
		{ /* '0x23' */
			return PlcValueType_LTIME
		}
	case 0x24:
		{ /* '0x24' */
			return PlcValueType_DATE
		}
	case 0x25:
		{ /* '0x25' */
			return PlcValueType_TIME_OF_DAY
		}
	case 0x26:
		{ /* '0x26' */
			return PlcValueType_TIME_OF_DAY
		}
	case 0x27:
		{ /* '0x27' */
			return PlcValueType_DATE_AND_TIME
		}
	case 0x28:
		{ /* '0x28' */
			return PlcValueType_DATE_AND_TIME
		}
	default:
		{
			return 0
		}
	}
}

func AdsDataTypeFirstEnumForFieldPlcValueType(value PlcValueType) (AdsDataType, error) {
	for _, sizeValue := range AdsDataTypeValues {
		if sizeValue.PlcValueType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing PlcValueType not found", value)
}
func AdsDataTypeByValue(value int8) (enum AdsDataType, ok bool) {
	switch value {
	case 0x01:
		return AdsDataType_BOOL, true
	case 0x02:
		return AdsDataType_BIT, true
	case 0x03:
		return AdsDataType_BIT8, true
	case 0x04:
		return AdsDataType_BYTE, true
	case 0x05:
		return AdsDataType_BITARR8, true
	case 0x06:
		return AdsDataType_WORD, true
	case 0x07:
		return AdsDataType_BITARR16, true
	case 0x08:
		return AdsDataType_DWORD, true
	case 0x09:
		return AdsDataType_BITARR32, true
	case 0x0A:
		return AdsDataType_SINT, true
	case 0x0B:
		return AdsDataType_INT8, true
	case 0x0C:
		return AdsDataType_USINT, true
	case 0x0D:
		return AdsDataType_UINT8, true
	case 0x0E:
		return AdsDataType_INT, true
	case 0x0F:
		return AdsDataType_INT16, true
	case 0x10:
		return AdsDataType_UINT, true
	case 0x11:
		return AdsDataType_UINT16, true
	case 0x12:
		return AdsDataType_DINT, true
	case 0x13:
		return AdsDataType_INT32, true
	case 0x14:
		return AdsDataType_UDINT, true
	case 0x15:
		return AdsDataType_UINT32, true
	case 0x16:
		return AdsDataType_LINT, true
	case 0x17:
		return AdsDataType_INT64, true
	case 0x18:
		return AdsDataType_ULINT, true
	case 0x19:
		return AdsDataType_UINT64, true
	case 0x1A:
		return AdsDataType_REAL, true
	case 0x1B:
		return AdsDataType_FLOAT, true
	case 0x1C:
		return AdsDataType_LREAL, true
	case 0x1D:
		return AdsDataType_DOUBLE, true
	case 0x1E:
		return AdsDataType_CHAR, true
	case 0x1F:
		return AdsDataType_WCHAR, true
	case 0x20:
		return AdsDataType_STRING, true
	case 0x21:
		return AdsDataType_WSTRING, true
	case 0x22:
		return AdsDataType_TIME, true
	case 0x23:
		return AdsDataType_LTIME, true
	case 0x24:
		return AdsDataType_DATE, true
	case 0x25:
		return AdsDataType_TIME_OF_DAY, true
	case 0x26:
		return AdsDataType_TOD, true
	case 0x27:
		return AdsDataType_DATE_AND_TIME, true
	case 0x28:
		return AdsDataType_DT, true
	}
	return 0, false
}

func AdsDataTypeByName(value string) (enum AdsDataType, ok bool) {
	switch value {
	case "BOOL":
		return AdsDataType_BOOL, true
	case "BIT":
		return AdsDataType_BIT, true
	case "BIT8":
		return AdsDataType_BIT8, true
	case "BYTE":
		return AdsDataType_BYTE, true
	case "BITARR8":
		return AdsDataType_BITARR8, true
	case "WORD":
		return AdsDataType_WORD, true
	case "BITARR16":
		return AdsDataType_BITARR16, true
	case "DWORD":
		return AdsDataType_DWORD, true
	case "BITARR32":
		return AdsDataType_BITARR32, true
	case "SINT":
		return AdsDataType_SINT, true
	case "INT8":
		return AdsDataType_INT8, true
	case "USINT":
		return AdsDataType_USINT, true
	case "UINT8":
		return AdsDataType_UINT8, true
	case "INT":
		return AdsDataType_INT, true
	case "INT16":
		return AdsDataType_INT16, true
	case "UINT":
		return AdsDataType_UINT, true
	case "UINT16":
		return AdsDataType_UINT16, true
	case "DINT":
		return AdsDataType_DINT, true
	case "INT32":
		return AdsDataType_INT32, true
	case "UDINT":
		return AdsDataType_UDINT, true
	case "UINT32":
		return AdsDataType_UINT32, true
	case "LINT":
		return AdsDataType_LINT, true
	case "INT64":
		return AdsDataType_INT64, true
	case "ULINT":
		return AdsDataType_ULINT, true
	case "UINT64":
		return AdsDataType_UINT64, true
	case "REAL":
		return AdsDataType_REAL, true
	case "FLOAT":
		return AdsDataType_FLOAT, true
	case "LREAL":
		return AdsDataType_LREAL, true
	case "DOUBLE":
		return AdsDataType_DOUBLE, true
	case "CHAR":
		return AdsDataType_CHAR, true
	case "WCHAR":
		return AdsDataType_WCHAR, true
	case "STRING":
		return AdsDataType_STRING, true
	case "WSTRING":
		return AdsDataType_WSTRING, true
	case "TIME":
		return AdsDataType_TIME, true
	case "LTIME":
		return AdsDataType_LTIME, true
	case "DATE":
		return AdsDataType_DATE, true
	case "TIME_OF_DAY":
		return AdsDataType_TIME_OF_DAY, true
	case "TOD":
		return AdsDataType_TOD, true
	case "DATE_AND_TIME":
		return AdsDataType_DATE_AND_TIME, true
	case "DT":
		return AdsDataType_DT, true
	}
	return 0, false
}

func AdsDataTypeKnows(value int8) bool {
	for _, typeValue := range AdsDataTypeValues {
		if int8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAdsDataType(structType any) AdsDataType {
	castFunc := func(typ any) AdsDataType {
		if sAdsDataType, ok := typ.(AdsDataType); ok {
			return sAdsDataType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AdsDataType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m AdsDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDataTypeParse(ctx context.Context, theBytes []byte) (AdsDataType, error) {
	return AdsDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt8("AdsDataType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AdsDataType")
	}
	if enum, ok := AdsDataTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AdsDataType")
		return AdsDataType(val), nil
	} else {
		return enum, nil
	}
}

func (e AdsDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AdsDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt8("AdsDataType", 8, int8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AdsDataType) PLC4XEnumName() string {
	switch e {
	case AdsDataType_BOOL:
		return "BOOL"
	case AdsDataType_BIT:
		return "BIT"
	case AdsDataType_BIT8:
		return "BIT8"
	case AdsDataType_BYTE:
		return "BYTE"
	case AdsDataType_BITARR8:
		return "BITARR8"
	case AdsDataType_WORD:
		return "WORD"
	case AdsDataType_BITARR16:
		return "BITARR16"
	case AdsDataType_DWORD:
		return "DWORD"
	case AdsDataType_BITARR32:
		return "BITARR32"
	case AdsDataType_SINT:
		return "SINT"
	case AdsDataType_INT8:
		return "INT8"
	case AdsDataType_USINT:
		return "USINT"
	case AdsDataType_UINT8:
		return "UINT8"
	case AdsDataType_INT:
		return "INT"
	case AdsDataType_INT16:
		return "INT16"
	case AdsDataType_UINT:
		return "UINT"
	case AdsDataType_UINT16:
		return "UINT16"
	case AdsDataType_DINT:
		return "DINT"
	case AdsDataType_INT32:
		return "INT32"
	case AdsDataType_UDINT:
		return "UDINT"
	case AdsDataType_UINT32:
		return "UINT32"
	case AdsDataType_LINT:
		return "LINT"
	case AdsDataType_INT64:
		return "INT64"
	case AdsDataType_ULINT:
		return "ULINT"
	case AdsDataType_UINT64:
		return "UINT64"
	case AdsDataType_REAL:
		return "REAL"
	case AdsDataType_FLOAT:
		return "FLOAT"
	case AdsDataType_LREAL:
		return "LREAL"
	case AdsDataType_DOUBLE:
		return "DOUBLE"
	case AdsDataType_CHAR:
		return "CHAR"
	case AdsDataType_WCHAR:
		return "WCHAR"
	case AdsDataType_STRING:
		return "STRING"
	case AdsDataType_WSTRING:
		return "WSTRING"
	case AdsDataType_TIME:
		return "TIME"
	case AdsDataType_LTIME:
		return "LTIME"
	case AdsDataType_DATE:
		return "DATE"
	case AdsDataType_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case AdsDataType_TOD:
		return "TOD"
	case AdsDataType_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case AdsDataType_DT:
		return "DT"
	}
	return ""
}

func (e AdsDataType) String() string {
	return e.PLC4XEnumName()
}
