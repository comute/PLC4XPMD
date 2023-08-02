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

// HVACSensorStatus is an enum
type HVACSensorStatus uint8

type IHVACSensorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	HVACSensorStatus_NO_ERROR_OPERATING_NORMALLY               HVACSensorStatus = 0x00
	HVACSensorStatus_SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND HVACSensorStatus = 0x01
	HVACSensorStatus_SENSOR_OUT_OF_CALIBRATION                 HVACSensorStatus = 0x02
	HVACSensorStatus_SENSOR_TOTAL_FAILURE                      HVACSensorStatus = 0x03
)

var HVACSensorStatusValues []HVACSensorStatus

func init() {
	_ = errors.New
	HVACSensorStatusValues = []HVACSensorStatus{
		HVACSensorStatus_NO_ERROR_OPERATING_NORMALLY,
		HVACSensorStatus_SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND,
		HVACSensorStatus_SENSOR_OUT_OF_CALIBRATION,
		HVACSensorStatus_SENSOR_TOTAL_FAILURE,
	}
}

func HVACSensorStatusByValue(value uint8) (enum HVACSensorStatus, ok bool) {
	switch value {
	case 0x00:
		return HVACSensorStatus_NO_ERROR_OPERATING_NORMALLY, true
	case 0x01:
		return HVACSensorStatus_SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND, true
	case 0x02:
		return HVACSensorStatus_SENSOR_OUT_OF_CALIBRATION, true
	case 0x03:
		return HVACSensorStatus_SENSOR_TOTAL_FAILURE, true
	}
	return 0, false
}

func HVACSensorStatusByName(value string) (enum HVACSensorStatus, ok bool) {
	switch value {
	case "NO_ERROR_OPERATING_NORMALLY":
		return HVACSensorStatus_NO_ERROR_OPERATING_NORMALLY, true
	case "SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND":
		return HVACSensorStatus_SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND, true
	case "SENSOR_OUT_OF_CALIBRATION":
		return HVACSensorStatus_SENSOR_OUT_OF_CALIBRATION, true
	case "SENSOR_TOTAL_FAILURE":
		return HVACSensorStatus_SENSOR_TOTAL_FAILURE, true
	}
	return 0, false
}

func HVACSensorStatusKnows(value uint8) bool {
	for _, typeValue := range HVACSensorStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastHVACSensorStatus(structType any) HVACSensorStatus {
	castFunc := func(typ any) HVACSensorStatus {
		if sHVACSensorStatus, ok := typ.(HVACSensorStatus); ok {
			return sHVACSensorStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m HVACSensorStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m HVACSensorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACSensorStatusParse(ctx context.Context, theBytes []byte) (HVACSensorStatus, error) {
	return HVACSensorStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACSensorStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACSensorStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("HVACSensorStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HVACSensorStatus")
	}
	if enum, ok := HVACSensorStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for HVACSensorStatus")
		return HVACSensorStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e HVACSensorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HVACSensorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("HVACSensorStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HVACSensorStatus) PLC4XEnumName() string {
	switch e {
	case HVACSensorStatus_NO_ERROR_OPERATING_NORMALLY:
		return "NO_ERROR_OPERATING_NORMALLY"
	case HVACSensorStatus_SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND:
		return "SENSOR_OPERATING_IN_RELAXED_ACCURACY_BAND"
	case HVACSensorStatus_SENSOR_OUT_OF_CALIBRATION:
		return "SENSOR_OUT_OF_CALIBRATION"
	case HVACSensorStatus_SENSOR_TOTAL_FAILURE:
		return "SENSOR_TOTAL_FAILURE"
	}
	return ""
}

func (e HVACSensorStatus) String() string {
	return e.PLC4XEnumName()
}
