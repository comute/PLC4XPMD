/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type AlarmStateType uint8

type IAlarmStateType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	AlarmStateType_SCAN_ABORT       AlarmStateType = 0x00
	AlarmStateType_SCAN_INITIATE    AlarmStateType = 0x01
	AlarmStateType_ALARM_ABORT      AlarmStateType = 0x04
	AlarmStateType_ALARM_INITIATE   AlarmStateType = 0x05
	AlarmStateType_ALARM_S_ABORT    AlarmStateType = 0x08
	AlarmStateType_ALARM_S_INITIATE AlarmStateType = 0x09
)

var AlarmStateTypeValues []AlarmStateType

func init() {
	_ = errors.New
	AlarmStateTypeValues = []AlarmStateType{
		AlarmStateType_SCAN_ABORT,
		AlarmStateType_SCAN_INITIATE,
		AlarmStateType_ALARM_ABORT,
		AlarmStateType_ALARM_INITIATE,
		AlarmStateType_ALARM_S_ABORT,
		AlarmStateType_ALARM_S_INITIATE,
	}
}

func AlarmStateTypeByValue(value uint8) AlarmStateType {
	switch value {
	case 0x00:
		return AlarmStateType_SCAN_ABORT
	case 0x01:
		return AlarmStateType_SCAN_INITIATE
	case 0x04:
		return AlarmStateType_ALARM_ABORT
	case 0x05:
		return AlarmStateType_ALARM_INITIATE
	case 0x08:
		return AlarmStateType_ALARM_S_ABORT
	case 0x09:
		return AlarmStateType_ALARM_S_INITIATE
	}
	return 0
}

func AlarmStateTypeByName(value string) AlarmStateType {
	switch value {
	case "SCAN_ABORT":
		return AlarmStateType_SCAN_ABORT
	case "SCAN_INITIATE":
		return AlarmStateType_SCAN_INITIATE
	case "ALARM_ABORT":
		return AlarmStateType_ALARM_ABORT
	case "ALARM_INITIATE":
		return AlarmStateType_ALARM_INITIATE
	case "ALARM_S_ABORT":
		return AlarmStateType_ALARM_S_ABORT
	case "ALARM_S_INITIATE":
		return AlarmStateType_ALARM_S_INITIATE
	}
	return 0
}

func CastAlarmStateType(structType interface{}) AlarmStateType {
	castFunc := func(typ interface{}) AlarmStateType {
		if sAlarmStateType, ok := typ.(AlarmStateType); ok {
			return sAlarmStateType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AlarmStateType) LengthInBits() uint16 {
	return 8
}

func (m AlarmStateType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AlarmStateTypeParse(readBuffer utils.ReadBuffer) (AlarmStateType, error) {
	val, err := readBuffer.ReadUint8("AlarmStateType", 8)
	if err != nil {
		return 0, nil
	}
	return AlarmStateTypeByValue(val), nil
}

func (e AlarmStateType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AlarmStateType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e AlarmStateType) name() string {
	switch e {
	case AlarmStateType_SCAN_ABORT:
		return "SCAN_ABORT"
	case AlarmStateType_SCAN_INITIATE:
		return "SCAN_INITIATE"
	case AlarmStateType_ALARM_ABORT:
		return "ALARM_ABORT"
	case AlarmStateType_ALARM_INITIATE:
		return "ALARM_INITIATE"
	case AlarmStateType_ALARM_S_ABORT:
		return "ALARM_S_ABORT"
	case AlarmStateType_ALARM_S_INITIATE:
		return "ALARM_S_INITIATE"
	}
	return ""
}

func (e AlarmStateType) String() string {
	return e.name()
}
