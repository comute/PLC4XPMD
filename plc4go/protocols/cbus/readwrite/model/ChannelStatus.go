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

// ChannelStatus is an enum
type ChannelStatus uint8

type IChannelStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ChannelStatus_OK                     ChannelStatus = 0
	ChannelStatus_LAMP_FAULT             ChannelStatus = 2
	ChannelStatus_CURRENT_LIMIT_OR_SHORT ChannelStatus = 3
)

var ChannelStatusValues []ChannelStatus

func init() {
	_ = errors.New
	ChannelStatusValues = []ChannelStatus{
		ChannelStatus_OK,
		ChannelStatus_LAMP_FAULT,
		ChannelStatus_CURRENT_LIMIT_OR_SHORT,
	}
}

func ChannelStatusByValue(value uint8) (enum ChannelStatus, ok bool) {
	switch value {
	case 0:
		return ChannelStatus_OK, true
	case 2:
		return ChannelStatus_LAMP_FAULT, true
	case 3:
		return ChannelStatus_CURRENT_LIMIT_OR_SHORT, true
	}
	return 0, false
}

func ChannelStatusByName(value string) (enum ChannelStatus, ok bool) {
	switch value {
	case "OK":
		return ChannelStatus_OK, true
	case "LAMP_FAULT":
		return ChannelStatus_LAMP_FAULT, true
	case "CURRENT_LIMIT_OR_SHORT":
		return ChannelStatus_CURRENT_LIMIT_OR_SHORT, true
	}
	return 0, false
}

func ChannelStatusKnows(value uint8) bool {
	for _, typeValue := range ChannelStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastChannelStatus(structType any) ChannelStatus {
	castFunc := func(typ any) ChannelStatus {
		if sChannelStatus, ok := typ.(ChannelStatus); ok {
			return sChannelStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m ChannelStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ChannelStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChannelStatusParse(ctx context.Context, theBytes []byte) (ChannelStatus, error) {
	return ChannelStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ChannelStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ChannelStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ChannelStatus", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ChannelStatus")
	}
	if enum, ok := ChannelStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ChannelStatus")
		return ChannelStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e ChannelStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ChannelStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ChannelStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ChannelStatus) PLC4XEnumName() string {
	switch e {
	case ChannelStatus_OK:
		return "OK"
	case ChannelStatus_LAMP_FAULT:
		return "LAMP_FAULT"
	case ChannelStatus_CURRENT_LIMIT_OR_SHORT:
		return "CURRENT_LIMIT_OR_SHORT"
	}
	return ""
}

func (e ChannelStatus) String() string {
	return e.PLC4XEnumName()
}
