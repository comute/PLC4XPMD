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

// BACnetStatusFlags is an enum
type BACnetStatusFlags uint8

type IBACnetStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetStatusFlags_IN_ALARM       BACnetStatusFlags = 0
	BACnetStatusFlags_FAULT          BACnetStatusFlags = 1
	BACnetStatusFlags_OVERRIDDEN     BACnetStatusFlags = 2
	BACnetStatusFlags_OUT_OF_SERVICE BACnetStatusFlags = 3
)

var BACnetStatusFlagsValues []BACnetStatusFlags

func init() {
	_ = errors.New
	BACnetStatusFlagsValues = []BACnetStatusFlags{
		BACnetStatusFlags_IN_ALARM,
		BACnetStatusFlags_FAULT,
		BACnetStatusFlags_OVERRIDDEN,
		BACnetStatusFlags_OUT_OF_SERVICE,
	}
}

func BACnetStatusFlagsByValue(value uint8) (enum BACnetStatusFlags, ok bool) {
	switch value {
	case 0:
		return BACnetStatusFlags_IN_ALARM, true
	case 1:
		return BACnetStatusFlags_FAULT, true
	case 2:
		return BACnetStatusFlags_OVERRIDDEN, true
	case 3:
		return BACnetStatusFlags_OUT_OF_SERVICE, true
	}
	return 0, false
}

func BACnetStatusFlagsByName(value string) (enum BACnetStatusFlags, ok bool) {
	switch value {
	case "IN_ALARM":
		return BACnetStatusFlags_IN_ALARM, true
	case "FAULT":
		return BACnetStatusFlags_FAULT, true
	case "OVERRIDDEN":
		return BACnetStatusFlags_OVERRIDDEN, true
	case "OUT_OF_SERVICE":
		return BACnetStatusFlags_OUT_OF_SERVICE, true
	}
	return 0, false
}

func BACnetStatusFlagsKnows(value uint8) bool {
	for _, typeValue := range BACnetStatusFlagsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetStatusFlags(structType any) BACnetStatusFlags {
	castFunc := func(typ any) BACnetStatusFlags {
		if sBACnetStatusFlags, ok := typ.(BACnetStatusFlags); ok {
			return sBACnetStatusFlags
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetStatusFlagsParse(ctx context.Context, theBytes []byte) (BACnetStatusFlags, error) {
	return BACnetStatusFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetStatusFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetStatusFlags, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetStatusFlags", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetStatusFlags")
	}
	if enum, ok := BACnetStatusFlagsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetStatusFlags")
		return BACnetStatusFlags(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetStatusFlags", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetStatusFlags) PLC4XEnumName() string {
	switch e {
	case BACnetStatusFlags_IN_ALARM:
		return "IN_ALARM"
	case BACnetStatusFlags_FAULT:
		return "FAULT"
	case BACnetStatusFlags_OVERRIDDEN:
		return "OVERRIDDEN"
	case BACnetStatusFlags_OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	}
	return ""
}

func (e BACnetStatusFlags) String() string {
	return e.PLC4XEnumName()
}
