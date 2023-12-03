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

// BACnetProgramError is an enum
type BACnetProgramError uint16

type IBACnetProgramError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetProgramError_NORMAL                   BACnetProgramError = 0
	BACnetProgramError_LOAD_FAILED              BACnetProgramError = 1
	BACnetProgramError_INTERNAL                 BACnetProgramError = 2
	BACnetProgramError_PROGRAM                  BACnetProgramError = 3
	BACnetProgramError_OTHER                    BACnetProgramError = 4
	BACnetProgramError_VENDOR_PROPRIETARY_VALUE BACnetProgramError = 0xFFFF
)

var BACnetProgramErrorValues []BACnetProgramError

func init() {
	_ = errors.New
	BACnetProgramErrorValues = []BACnetProgramError{
		BACnetProgramError_NORMAL,
		BACnetProgramError_LOAD_FAILED,
		BACnetProgramError_INTERNAL,
		BACnetProgramError_PROGRAM,
		BACnetProgramError_OTHER,
		BACnetProgramError_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetProgramErrorByValue(value uint16) (enum BACnetProgramError, ok bool) {
	switch value {
	case 0:
		return BACnetProgramError_NORMAL, true
	case 0xFFFF:
		return BACnetProgramError_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetProgramError_LOAD_FAILED, true
	case 2:
		return BACnetProgramError_INTERNAL, true
	case 3:
		return BACnetProgramError_PROGRAM, true
	case 4:
		return BACnetProgramError_OTHER, true
	}
	return 0, false
}

func BACnetProgramErrorByName(value string) (enum BACnetProgramError, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetProgramError_NORMAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetProgramError_VENDOR_PROPRIETARY_VALUE, true
	case "LOAD_FAILED":
		return BACnetProgramError_LOAD_FAILED, true
	case "INTERNAL":
		return BACnetProgramError_INTERNAL, true
	case "PROGRAM":
		return BACnetProgramError_PROGRAM, true
	case "OTHER":
		return BACnetProgramError_OTHER, true
	}
	return 0, false
}

func BACnetProgramErrorKnows(value uint16) bool {
	for _, typeValue := range BACnetProgramErrorValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetProgramError(structType any) BACnetProgramError {
	castFunc := func(typ any) BACnetProgramError {
		if sBACnetProgramError, ok := typ.(BACnetProgramError); ok {
			return sBACnetProgramError
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetProgramError) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetProgramError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProgramErrorParse(ctx context.Context, theBytes []byte) (BACnetProgramError, error) {
	return BACnetProgramErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetProgramErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramError, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetProgramError", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetProgramError")
	}
	if enum, ok := BACnetProgramErrorByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetProgramError")
		return BACnetProgramError(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetProgramError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetProgramError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetProgramError", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetProgramError) PLC4XEnumName() string {
	switch e {
	case BACnetProgramError_NORMAL:
		return "NORMAL"
	case BACnetProgramError_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetProgramError_LOAD_FAILED:
		return "LOAD_FAILED"
	case BACnetProgramError_INTERNAL:
		return "INTERNAL"
	case BACnetProgramError_PROGRAM:
		return "PROGRAM"
	case BACnetProgramError_OTHER:
		return "OTHER"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetProgramError) String() string {
	return e.PLC4XEnumName()
}
