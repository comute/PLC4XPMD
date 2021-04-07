//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type KnxPropertyDataType uint8

type IKnxPropertyDataType interface {
	Number() uint8
	SizeInBytes() uint8
	Name() string
	Serialize(io utils.WriteBuffer) error
}

const (
	KnxPropertyDataType_PDT_UNKNOWN             KnxPropertyDataType = 0
	KnxPropertyDataType_PDT_CONTROL             KnxPropertyDataType = 1
	KnxPropertyDataType_PDT_CHAR                KnxPropertyDataType = 2
	KnxPropertyDataType_PDT_UNSIGNED_CHAR       KnxPropertyDataType = 3
	KnxPropertyDataType_PDT_INT                 KnxPropertyDataType = 4
	KnxPropertyDataType_PDT_UNSIGNED_INT        KnxPropertyDataType = 5
	KnxPropertyDataType_PDT_KNX_FLOAT           KnxPropertyDataType = 6
	KnxPropertyDataType_PDT_DATE                KnxPropertyDataType = 7
	KnxPropertyDataType_PDT_TIME                KnxPropertyDataType = 8
	KnxPropertyDataType_PDT_LONG                KnxPropertyDataType = 9
	KnxPropertyDataType_PDT_UNSIGNED_LONG       KnxPropertyDataType = 10
	KnxPropertyDataType_PDT_FLOAT               KnxPropertyDataType = 11
	KnxPropertyDataType_PDT_DOUBLE              KnxPropertyDataType = 12
	KnxPropertyDataType_PDT_CHAR_BLOCK          KnxPropertyDataType = 13
	KnxPropertyDataType_PDT_POLL_GROUP_SETTINGS KnxPropertyDataType = 14
	KnxPropertyDataType_PDT_SHORT_CHAR_BLOCK    KnxPropertyDataType = 15
	KnxPropertyDataType_PDT_DATE_TIME           KnxPropertyDataType = 16
	KnxPropertyDataType_PDT_VARIABLE_LENGTH     KnxPropertyDataType = 17
	KnxPropertyDataType_PDT_GENERIC_01          KnxPropertyDataType = 18
	KnxPropertyDataType_PDT_GENERIC_02          KnxPropertyDataType = 19
	KnxPropertyDataType_PDT_GENERIC_03          KnxPropertyDataType = 20
	KnxPropertyDataType_PDT_GENERIC_04          KnxPropertyDataType = 21
	KnxPropertyDataType_PDT_GENERIC_05          KnxPropertyDataType = 22
	KnxPropertyDataType_PDT_GENERIC_06          KnxPropertyDataType = 23
	KnxPropertyDataType_PDT_GENERIC_07          KnxPropertyDataType = 24
	KnxPropertyDataType_PDT_GENERIC_08          KnxPropertyDataType = 25
	KnxPropertyDataType_PDT_GENERIC_09          KnxPropertyDataType = 26
	KnxPropertyDataType_PDT_GENERIC_10          KnxPropertyDataType = 27
	KnxPropertyDataType_PDT_GENERIC_11          KnxPropertyDataType = 28
	KnxPropertyDataType_PDT_GENERIC_12          KnxPropertyDataType = 29
	KnxPropertyDataType_PDT_GENERIC_13          KnxPropertyDataType = 30
	KnxPropertyDataType_PDT_GENERIC_14          KnxPropertyDataType = 31
	KnxPropertyDataType_PDT_GENERIC_15          KnxPropertyDataType = 32
	KnxPropertyDataType_PDT_GENERIC_16          KnxPropertyDataType = 33
	KnxPropertyDataType_PDT_GENERIC_17          KnxPropertyDataType = 34
	KnxPropertyDataType_PDT_GENERIC_18          KnxPropertyDataType = 35
	KnxPropertyDataType_PDT_GENERIC_19          KnxPropertyDataType = 36
	KnxPropertyDataType_PDT_GENERIC_20          KnxPropertyDataType = 37
	KnxPropertyDataType_PDT_UTF_8               KnxPropertyDataType = 38
	KnxPropertyDataType_PDT_VERSION             KnxPropertyDataType = 39
	KnxPropertyDataType_PDT_ALARM_INFO          KnxPropertyDataType = 40
	KnxPropertyDataType_PDT_BINARY_INFORMATION  KnxPropertyDataType = 41
	KnxPropertyDataType_PDT_BITSET8             KnxPropertyDataType = 42
	KnxPropertyDataType_PDT_BITSET16            KnxPropertyDataType = 43
	KnxPropertyDataType_PDT_ENUM8               KnxPropertyDataType = 44
	KnxPropertyDataType_PDT_SCALING             KnxPropertyDataType = 45
	KnxPropertyDataType_PDT_NE_VL               KnxPropertyDataType = 46
	KnxPropertyDataType_PDT_NE_FL               KnxPropertyDataType = 47
	KnxPropertyDataType_PDT_FUNCTION            KnxPropertyDataType = 48
	KnxPropertyDataType_PDT_ESCAPE              KnxPropertyDataType = 49
)

func (e KnxPropertyDataType) Number() uint8 {
	switch e {
	case 0:
		{ /* '0' */
			return 0
		}
	case 1:
		{ /* '1' */
			return 0
		}
	case 10:
		{ /* '10' */
			return 9
		}
	case 11:
		{ /* '11' */
			return 10
		}
	case 12:
		{ /* '12' */
			return 11
		}
	case 13:
		{ /* '13' */
			return 12
		}
	case 14:
		{ /* '14' */
			return 13
		}
	case 15:
		{ /* '15' */
			return 14
		}
	case 16:
		{ /* '16' */
			return 15
		}
	case 17:
		{ /* '17' */
			return 16
		}
	case 18:
		{ /* '18' */
			return 17
		}
	case 19:
		{ /* '19' */
			return 18
		}
	case 2:
		{ /* '2' */
			return 1
		}
	case 20:
		{ /* '20' */
			return 19
		}
	case 21:
		{ /* '21' */
			return 20
		}
	case 22:
		{ /* '22' */
			return 21
		}
	case 23:
		{ /* '23' */
			return 22
		}
	case 24:
		{ /* '24' */
			return 23
		}
	case 25:
		{ /* '25' */
			return 24
		}
	case 26:
		{ /* '26' */
			return 25
		}
	case 27:
		{ /* '27' */
			return 26
		}
	case 28:
		{ /* '28' */
			return 27
		}
	case 29:
		{ /* '29' */
			return 28
		}
	case 3:
		{ /* '3' */
			return 2
		}
	case 30:
		{ /* '30' */
			return 29
		}
	case 31:
		{ /* '31' */
			return 30
		}
	case 32:
		{ /* '32' */
			return 31
		}
	case 33:
		{ /* '33' */
			return 32
		}
	case 34:
		{ /* '34' */
			return 33
		}
	case 35:
		{ /* '35' */
			return 34
		}
	case 36:
		{ /* '36' */
			return 35
		}
	case 37:
		{ /* '37' */
			return 36
		}
	case 38:
		{ /* '38' */
			return 47
		}
	case 39:
		{ /* '39' */
			return 48
		}
	case 4:
		{ /* '4' */
			return 3
		}
	case 40:
		{ /* '40' */
			return 49
		}
	case 41:
		{ /* '41' */
			return 50
		}
	case 42:
		{ /* '42' */
			return 51
		}
	case 43:
		{ /* '43' */
			return 52
		}
	case 44:
		{ /* '44' */
			return 53
		}
	case 45:
		{ /* '45' */
			return 54
		}
	case 46:
		{ /* '46' */
			return 60
		}
	case 47:
		{ /* '47' */
			return 61
		}
	case 48:
		{ /* '48' */
			return 62
		}
	case 49:
		{ /* '49' */
			return 63
		}
	case 5:
		{ /* '5' */
			return 4
		}
	case 6:
		{ /* '6' */
			return 5
		}
	case 7:
		{ /* '7' */
			return 6
		}
	case 8:
		{ /* '8' */
			return 7
		}
	case 9:
		{ /* '9' */
			return 8
		}
	default:
		{
			return 0
		}
	}
}

func (e KnxPropertyDataType) SizeInBytes() uint8 {
	switch e {
	case 0:
		{ /* '0' */
			return 0
		}
	case 1:
		{ /* '1' */
			return 10
		}
	case 10:
		{ /* '10' */
			return 4
		}
	case 11:
		{ /* '11' */
			return 4
		}
	case 12:
		{ /* '12' */
			return 8
		}
	case 13:
		{ /* '13' */
			return 10
		}
	case 14:
		{ /* '14' */
			return 3
		}
	case 15:
		{ /* '15' */
			return 5
		}
	case 16:
		{ /* '16' */
			return 8
		}
	case 17:
		{ /* '17' */
			return 0
		}
	case 18:
		{ /* '18' */
			return 1
		}
	case 19:
		{ /* '19' */
			return 2
		}
	case 2:
		{ /* '2' */
			return 1
		}
	case 20:
		{ /* '20' */
			return 3
		}
	case 21:
		{ /* '21' */
			return 4
		}
	case 22:
		{ /* '22' */
			return 5
		}
	case 23:
		{ /* '23' */
			return 6
		}
	case 24:
		{ /* '24' */
			return 7
		}
	case 25:
		{ /* '25' */
			return 8
		}
	case 26:
		{ /* '26' */
			return 9
		}
	case 27:
		{ /* '27' */
			return 10
		}
	case 28:
		{ /* '28' */
			return 11
		}
	case 29:
		{ /* '29' */
			return 12
		}
	case 3:
		{ /* '3' */
			return 1
		}
	case 30:
		{ /* '30' */
			return 13
		}
	case 31:
		{ /* '31' */
			return 14
		}
	case 32:
		{ /* '32' */
			return 15
		}
	case 33:
		{ /* '33' */
			return 16
		}
	case 34:
		{ /* '34' */
			return 17
		}
	case 35:
		{ /* '35' */
			return 18
		}
	case 36:
		{ /* '36' */
			return 19
		}
	case 37:
		{ /* '37' */
			return 20
		}
	case 38:
		{ /* '38' */
			return 0
		}
	case 39:
		{ /* '39' */
			return 2
		}
	case 4:
		{ /* '4' */
			return 2
		}
	case 40:
		{ /* '40' */
			return 6
		}
	case 41:
		{ /* '41' */
			return 1
		}
	case 42:
		{ /* '42' */
			return 1
		}
	case 43:
		{ /* '43' */
			return 2
		}
	case 44:
		{ /* '44' */
			return 1
		}
	case 45:
		{ /* '45' */
			return 1
		}
	case 46:
		{ /* '46' */
			return 0
		}
	case 47:
		{ /* '47' */
			return 0
		}
	case 48:
		{ /* '48' */
			return 0
		}
	case 49:
		{ /* '49' */
			return 0
		}
	case 5:
		{ /* '5' */
			return 2
		}
	case 6:
		{ /* '6' */
			return 2
		}
	case 7:
		{ /* '7' */
			return 3
		}
	case 8:
		{ /* '8' */
			return 3
		}
	case 9:
		{ /* '9' */
			return 4
		}
	default:
		{
			return 0
		}
	}
}

func (e KnxPropertyDataType) Name() string {
	switch e {
	case 0:
		{ /* '0' */
			return "Unknown Property Data Type"
		}
	case 1:
		{ /* '1' */
			return "PDT_CONTROL"
		}
	case 10:
		{ /* '10' */
			return "PDT_UNSIGNED_LONG"
		}
	case 11:
		{ /* '11' */
			return "PDT_FLOAT"
		}
	case 12:
		{ /* '12' */
			return "PDT_DOUBLE"
		}
	case 13:
		{ /* '13' */
			return "PDT_CHAR_BLOCK"
		}
	case 14:
		{ /* '14' */
			return "PDT_POLL_GROUP_SETTINGS"
		}
	case 15:
		{ /* '15' */
			return "PDT_SHORT_CHAR_BLOCK"
		}
	case 16:
		{ /* '16' */
			return "PDT_DATE_TIME"
		}
	case 17:
		{ /* '17' */
			return "PDT_VARIABLE_LENGTH"
		}
	case 18:
		{ /* '18' */
			return "PDT_GENERIC_01"
		}
	case 19:
		{ /* '19' */
			return "PDT_GENERIC_02"
		}
	case 2:
		{ /* '2' */
			return "PDT_CHAR"
		}
	case 20:
		{ /* '20' */
			return "PDT_GENERIC_03"
		}
	case 21:
		{ /* '21' */
			return "PDT_GENERIC_04"
		}
	case 22:
		{ /* '22' */
			return "PDT_GENERIC_05"
		}
	case 23:
		{ /* '23' */
			return "PDT_GENERIC_06"
		}
	case 24:
		{ /* '24' */
			return "PDT_GENERIC_07"
		}
	case 25:
		{ /* '25' */
			return "PDT_GENERIC_08"
		}
	case 26:
		{ /* '26' */
			return "PDT_GENERIC_09"
		}
	case 27:
		{ /* '27' */
			return "PDT_GENERIC_10"
		}
	case 28:
		{ /* '28' */
			return "PDT_GENERIC_11"
		}
	case 29:
		{ /* '29' */
			return "PDT_GENERIC_12"
		}
	case 3:
		{ /* '3' */
			return "PDT_UNSIGNED_CHAR"
		}
	case 30:
		{ /* '30' */
			return "PDT_GENERIC_13"
		}
	case 31:
		{ /* '31' */
			return "PDT_GENERIC_14"
		}
	case 32:
		{ /* '32' */
			return "PDT_GENERIC_15"
		}
	case 33:
		{ /* '33' */
			return "PDT_GENERIC_16"
		}
	case 34:
		{ /* '34' */
			return "PDT_GENERIC_17"
		}
	case 35:
		{ /* '35' */
			return "PDT_GENERIC_18"
		}
	case 36:
		{ /* '36' */
			return "PDT_GENERIC_19"
		}
	case 37:
		{ /* '37' */
			return "PDT_GENERIC_20"
		}
	case 38:
		{ /* '38' */
			return "PDT_UTF-8"
		}
	case 39:
		{ /* '39' */
			return "PDT_VERSION"
		}
	case 4:
		{ /* '4' */
			return "PDT_INT"
		}
	case 40:
		{ /* '40' */
			return "PDT_ALARM_INFO"
		}
	case 41:
		{ /* '41' */
			return "PDT_BINARY_INFORMATION"
		}
	case 42:
		{ /* '42' */
			return "PDT_BITSET8"
		}
	case 43:
		{ /* '43' */
			return "PDT_BITSET16"
		}
	case 44:
		{ /* '44' */
			return "PDT_ENUM8"
		}
	case 45:
		{ /* '45' */
			return "PDT_SCALING"
		}
	case 46:
		{ /* '46' */
			return "PDT_NE_VL"
		}
	case 47:
		{ /* '47' */
			return "PDT_NE_FL"
		}
	case 48:
		{ /* '48' */
			return "PDT_FUNCTION"
		}
	case 49:
		{ /* '49' */
			return "PDT_ESCAPE"
		}
	case 5:
		{ /* '5' */
			return "PDT_UNSIGNED_INT"
		}
	case 6:
		{ /* '6' */
			return "PDT_KNX_FLOAT"
		}
	case 7:
		{ /* '7' */
			return "PDT_DATE"
		}
	case 8:
		{ /* '8' */
			return "PDT_TIME"
		}
	case 9:
		{ /* '9' */
			return "PDT_LONG"
		}
	default:
		{
			return ""
		}
	}
}
func KnxPropertyDataTypeByValue(value uint8) KnxPropertyDataType {
	switch value {
	case 0:
		return KnxPropertyDataType_PDT_UNKNOWN
	case 1:
		return KnxPropertyDataType_PDT_CONTROL
	case 10:
		return KnxPropertyDataType_PDT_UNSIGNED_LONG
	case 11:
		return KnxPropertyDataType_PDT_FLOAT
	case 12:
		return KnxPropertyDataType_PDT_DOUBLE
	case 13:
		return KnxPropertyDataType_PDT_CHAR_BLOCK
	case 14:
		return KnxPropertyDataType_PDT_POLL_GROUP_SETTINGS
	case 15:
		return KnxPropertyDataType_PDT_SHORT_CHAR_BLOCK
	case 16:
		return KnxPropertyDataType_PDT_DATE_TIME
	case 17:
		return KnxPropertyDataType_PDT_VARIABLE_LENGTH
	case 18:
		return KnxPropertyDataType_PDT_GENERIC_01
	case 19:
		return KnxPropertyDataType_PDT_GENERIC_02
	case 2:
		return KnxPropertyDataType_PDT_CHAR
	case 20:
		return KnxPropertyDataType_PDT_GENERIC_03
	case 21:
		return KnxPropertyDataType_PDT_GENERIC_04
	case 22:
		return KnxPropertyDataType_PDT_GENERIC_05
	case 23:
		return KnxPropertyDataType_PDT_GENERIC_06
	case 24:
		return KnxPropertyDataType_PDT_GENERIC_07
	case 25:
		return KnxPropertyDataType_PDT_GENERIC_08
	case 26:
		return KnxPropertyDataType_PDT_GENERIC_09
	case 27:
		return KnxPropertyDataType_PDT_GENERIC_10
	case 28:
		return KnxPropertyDataType_PDT_GENERIC_11
	case 29:
		return KnxPropertyDataType_PDT_GENERIC_12
	case 3:
		return KnxPropertyDataType_PDT_UNSIGNED_CHAR
	case 30:
		return KnxPropertyDataType_PDT_GENERIC_13
	case 31:
		return KnxPropertyDataType_PDT_GENERIC_14
	case 32:
		return KnxPropertyDataType_PDT_GENERIC_15
	case 33:
		return KnxPropertyDataType_PDT_GENERIC_16
	case 34:
		return KnxPropertyDataType_PDT_GENERIC_17
	case 35:
		return KnxPropertyDataType_PDT_GENERIC_18
	case 36:
		return KnxPropertyDataType_PDT_GENERIC_19
	case 37:
		return KnxPropertyDataType_PDT_GENERIC_20
	case 38:
		return KnxPropertyDataType_PDT_UTF_8
	case 39:
		return KnxPropertyDataType_PDT_VERSION
	case 4:
		return KnxPropertyDataType_PDT_INT
	case 40:
		return KnxPropertyDataType_PDT_ALARM_INFO
	case 41:
		return KnxPropertyDataType_PDT_BINARY_INFORMATION
	case 42:
		return KnxPropertyDataType_PDT_BITSET8
	case 43:
		return KnxPropertyDataType_PDT_BITSET16
	case 44:
		return KnxPropertyDataType_PDT_ENUM8
	case 45:
		return KnxPropertyDataType_PDT_SCALING
	case 46:
		return KnxPropertyDataType_PDT_NE_VL
	case 47:
		return KnxPropertyDataType_PDT_NE_FL
	case 48:
		return KnxPropertyDataType_PDT_FUNCTION
	case 49:
		return KnxPropertyDataType_PDT_ESCAPE
	case 5:
		return KnxPropertyDataType_PDT_UNSIGNED_INT
	case 6:
		return KnxPropertyDataType_PDT_KNX_FLOAT
	case 7:
		return KnxPropertyDataType_PDT_DATE
	case 8:
		return KnxPropertyDataType_PDT_TIME
	case 9:
		return KnxPropertyDataType_PDT_LONG
	}
	return 0
}

func KnxPropertyDataTypeByName(value string) KnxPropertyDataType {
	switch value {
	case "PDT_UNKNOWN":
		return KnxPropertyDataType_PDT_UNKNOWN
	case "PDT_CONTROL":
		return KnxPropertyDataType_PDT_CONTROL
	case "PDT_UNSIGNED_LONG":
		return KnxPropertyDataType_PDT_UNSIGNED_LONG
	case "PDT_FLOAT":
		return KnxPropertyDataType_PDT_FLOAT
	case "PDT_DOUBLE":
		return KnxPropertyDataType_PDT_DOUBLE
	case "PDT_CHAR_BLOCK":
		return KnxPropertyDataType_PDT_CHAR_BLOCK
	case "PDT_POLL_GROUP_SETTINGS":
		return KnxPropertyDataType_PDT_POLL_GROUP_SETTINGS
	case "PDT_SHORT_CHAR_BLOCK":
		return KnxPropertyDataType_PDT_SHORT_CHAR_BLOCK
	case "PDT_DATE_TIME":
		return KnxPropertyDataType_PDT_DATE_TIME
	case "PDT_VARIABLE_LENGTH":
		return KnxPropertyDataType_PDT_VARIABLE_LENGTH
	case "PDT_GENERIC_01":
		return KnxPropertyDataType_PDT_GENERIC_01
	case "PDT_GENERIC_02":
		return KnxPropertyDataType_PDT_GENERIC_02
	case "PDT_CHAR":
		return KnxPropertyDataType_PDT_CHAR
	case "PDT_GENERIC_03":
		return KnxPropertyDataType_PDT_GENERIC_03
	case "PDT_GENERIC_04":
		return KnxPropertyDataType_PDT_GENERIC_04
	case "PDT_GENERIC_05":
		return KnxPropertyDataType_PDT_GENERIC_05
	case "PDT_GENERIC_06":
		return KnxPropertyDataType_PDT_GENERIC_06
	case "PDT_GENERIC_07":
		return KnxPropertyDataType_PDT_GENERIC_07
	case "PDT_GENERIC_08":
		return KnxPropertyDataType_PDT_GENERIC_08
	case "PDT_GENERIC_09":
		return KnxPropertyDataType_PDT_GENERIC_09
	case "PDT_GENERIC_10":
		return KnxPropertyDataType_PDT_GENERIC_10
	case "PDT_GENERIC_11":
		return KnxPropertyDataType_PDT_GENERIC_11
	case "PDT_GENERIC_12":
		return KnxPropertyDataType_PDT_GENERIC_12
	case "PDT_UNSIGNED_CHAR":
		return KnxPropertyDataType_PDT_UNSIGNED_CHAR
	case "PDT_GENERIC_13":
		return KnxPropertyDataType_PDT_GENERIC_13
	case "PDT_GENERIC_14":
		return KnxPropertyDataType_PDT_GENERIC_14
	case "PDT_GENERIC_15":
		return KnxPropertyDataType_PDT_GENERIC_15
	case "PDT_GENERIC_16":
		return KnxPropertyDataType_PDT_GENERIC_16
	case "PDT_GENERIC_17":
		return KnxPropertyDataType_PDT_GENERIC_17
	case "PDT_GENERIC_18":
		return KnxPropertyDataType_PDT_GENERIC_18
	case "PDT_GENERIC_19":
		return KnxPropertyDataType_PDT_GENERIC_19
	case "PDT_GENERIC_20":
		return KnxPropertyDataType_PDT_GENERIC_20
	case "PDT_UTF_8":
		return KnxPropertyDataType_PDT_UTF_8
	case "PDT_VERSION":
		return KnxPropertyDataType_PDT_VERSION
	case "PDT_INT":
		return KnxPropertyDataType_PDT_INT
	case "PDT_ALARM_INFO":
		return KnxPropertyDataType_PDT_ALARM_INFO
	case "PDT_BINARY_INFORMATION":
		return KnxPropertyDataType_PDT_BINARY_INFORMATION
	case "PDT_BITSET8":
		return KnxPropertyDataType_PDT_BITSET8
	case "PDT_BITSET16":
		return KnxPropertyDataType_PDT_BITSET16
	case "PDT_ENUM8":
		return KnxPropertyDataType_PDT_ENUM8
	case "PDT_SCALING":
		return KnxPropertyDataType_PDT_SCALING
	case "PDT_NE_VL":
		return KnxPropertyDataType_PDT_NE_VL
	case "PDT_NE_FL":
		return KnxPropertyDataType_PDT_NE_FL
	case "PDT_FUNCTION":
		return KnxPropertyDataType_PDT_FUNCTION
	case "PDT_ESCAPE":
		return KnxPropertyDataType_PDT_ESCAPE
	case "PDT_UNSIGNED_INT":
		return KnxPropertyDataType_PDT_UNSIGNED_INT
	case "PDT_KNX_FLOAT":
		return KnxPropertyDataType_PDT_KNX_FLOAT
	case "PDT_DATE":
		return KnxPropertyDataType_PDT_DATE
	case "PDT_TIME":
		return KnxPropertyDataType_PDT_TIME
	case "PDT_LONG":
		return KnxPropertyDataType_PDT_LONG
	}
	return 0
}

func CastKnxPropertyDataType(structType interface{}) KnxPropertyDataType {
	castFunc := func(typ interface{}) KnxPropertyDataType {
		if sKnxPropertyDataType, ok := typ.(KnxPropertyDataType); ok {
			return sKnxPropertyDataType
		}
		return 0
	}
	return castFunc(structType)
}

func (m KnxPropertyDataType) LengthInBits() uint16 {
	return 8
}

func (m KnxPropertyDataType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxPropertyDataTypeParse(io *utils.ReadBuffer) (KnxPropertyDataType, error) {
	val, err := io.ReadUint8(8)
	if err != nil {
		return 0, nil
	}
	return KnxPropertyDataTypeByValue(val), nil
}

func (e KnxPropertyDataType) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint8(8, uint8(e))
	return err
}

func (m *KnxPropertyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = KnxPropertyDataTypeByName(string(tok))
		}
	}
}

func (e KnxPropertyDataType) String() string {
	switch e {
	case KnxPropertyDataType_PDT_UNKNOWN:
		return "PDT_UNKNOWN"
	case KnxPropertyDataType_PDT_CONTROL:
		return "PDT_CONTROL"
	case KnxPropertyDataType_PDT_UNSIGNED_LONG:
		return "PDT_UNSIGNED_LONG"
	case KnxPropertyDataType_PDT_FLOAT:
		return "PDT_FLOAT"
	case KnxPropertyDataType_PDT_DOUBLE:
		return "PDT_DOUBLE"
	case KnxPropertyDataType_PDT_CHAR_BLOCK:
		return "PDT_CHAR_BLOCK"
	case KnxPropertyDataType_PDT_POLL_GROUP_SETTINGS:
		return "PDT_POLL_GROUP_SETTINGS"
	case KnxPropertyDataType_PDT_SHORT_CHAR_BLOCK:
		return "PDT_SHORT_CHAR_BLOCK"
	case KnxPropertyDataType_PDT_DATE_TIME:
		return "PDT_DATE_TIME"
	case KnxPropertyDataType_PDT_VARIABLE_LENGTH:
		return "PDT_VARIABLE_LENGTH"
	case KnxPropertyDataType_PDT_GENERIC_01:
		return "PDT_GENERIC_01"
	case KnxPropertyDataType_PDT_GENERIC_02:
		return "PDT_GENERIC_02"
	case KnxPropertyDataType_PDT_CHAR:
		return "PDT_CHAR"
	case KnxPropertyDataType_PDT_GENERIC_03:
		return "PDT_GENERIC_03"
	case KnxPropertyDataType_PDT_GENERIC_04:
		return "PDT_GENERIC_04"
	case KnxPropertyDataType_PDT_GENERIC_05:
		return "PDT_GENERIC_05"
	case KnxPropertyDataType_PDT_GENERIC_06:
		return "PDT_GENERIC_06"
	case KnxPropertyDataType_PDT_GENERIC_07:
		return "PDT_GENERIC_07"
	case KnxPropertyDataType_PDT_GENERIC_08:
		return "PDT_GENERIC_08"
	case KnxPropertyDataType_PDT_GENERIC_09:
		return "PDT_GENERIC_09"
	case KnxPropertyDataType_PDT_GENERIC_10:
		return "PDT_GENERIC_10"
	case KnxPropertyDataType_PDT_GENERIC_11:
		return "PDT_GENERIC_11"
	case KnxPropertyDataType_PDT_GENERIC_12:
		return "PDT_GENERIC_12"
	case KnxPropertyDataType_PDT_UNSIGNED_CHAR:
		return "PDT_UNSIGNED_CHAR"
	case KnxPropertyDataType_PDT_GENERIC_13:
		return "PDT_GENERIC_13"
	case KnxPropertyDataType_PDT_GENERIC_14:
		return "PDT_GENERIC_14"
	case KnxPropertyDataType_PDT_GENERIC_15:
		return "PDT_GENERIC_15"
	case KnxPropertyDataType_PDT_GENERIC_16:
		return "PDT_GENERIC_16"
	case KnxPropertyDataType_PDT_GENERIC_17:
		return "PDT_GENERIC_17"
	case KnxPropertyDataType_PDT_GENERIC_18:
		return "PDT_GENERIC_18"
	case KnxPropertyDataType_PDT_GENERIC_19:
		return "PDT_GENERIC_19"
	case KnxPropertyDataType_PDT_GENERIC_20:
		return "PDT_GENERIC_20"
	case KnxPropertyDataType_PDT_UTF_8:
		return "PDT_UTF_8"
	case KnxPropertyDataType_PDT_VERSION:
		return "PDT_VERSION"
	case KnxPropertyDataType_PDT_INT:
		return "PDT_INT"
	case KnxPropertyDataType_PDT_ALARM_INFO:
		return "PDT_ALARM_INFO"
	case KnxPropertyDataType_PDT_BINARY_INFORMATION:
		return "PDT_BINARY_INFORMATION"
	case KnxPropertyDataType_PDT_BITSET8:
		return "PDT_BITSET8"
	case KnxPropertyDataType_PDT_BITSET16:
		return "PDT_BITSET16"
	case KnxPropertyDataType_PDT_ENUM8:
		return "PDT_ENUM8"
	case KnxPropertyDataType_PDT_SCALING:
		return "PDT_SCALING"
	case KnxPropertyDataType_PDT_NE_VL:
		return "PDT_NE_VL"
	case KnxPropertyDataType_PDT_NE_FL:
		return "PDT_NE_FL"
	case KnxPropertyDataType_PDT_FUNCTION:
		return "PDT_FUNCTION"
	case KnxPropertyDataType_PDT_ESCAPE:
		return "PDT_ESCAPE"
	case KnxPropertyDataType_PDT_UNSIGNED_INT:
		return "PDT_UNSIGNED_INT"
	case KnxPropertyDataType_PDT_KNX_FLOAT:
		return "PDT_KNX_FLOAT"
	case KnxPropertyDataType_PDT_DATE:
		return "PDT_DATE"
	case KnxPropertyDataType_PDT_TIME:
		return "PDT_TIME"
	case KnxPropertyDataType_PDT_LONG:
		return "PDT_LONG"
	}
	return ""
}
