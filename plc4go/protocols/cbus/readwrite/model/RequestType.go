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

// RequestType is an enum
type RequestType uint8

type IRequestType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ControlChar() uint8
}

const (
	RequestType_UNKNOWN                RequestType = 0x00
	RequestType_SMART_CONNECT_SHORTCUT RequestType = 0x7C
	RequestType_RESET                  RequestType = 0x7E
	RequestType_DIRECT_COMMAND         RequestType = 0x40
	RequestType_REQUEST_COMMAND        RequestType = 0x5C
	RequestType_NULL                   RequestType = 0x6E
	RequestType_EMPTY                  RequestType = 0x0D
)

var RequestTypeValues []RequestType

func init() {
	_ = errors.New
	RequestTypeValues = []RequestType{
		RequestType_UNKNOWN,
		RequestType_SMART_CONNECT_SHORTCUT,
		RequestType_RESET,
		RequestType_DIRECT_COMMAND,
		RequestType_REQUEST_COMMAND,
		RequestType_NULL,
		RequestType_EMPTY,
	}
}

func (e RequestType) ControlChar() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0x00
		}
	case 0x0D:
		{ /* '0x0D' */
			return 0x00
		}
	case 0x40:
		{ /* '0x40' */
			return 0x40
		}
	case 0x5C:
		{ /* '0x5C' */
			return 0x5C
		}
	case 0x6E:
		{ /* '0x6E' */
			return 0x00
		}
	case 0x7C:
		{ /* '0x7C' */
			return 0x7C
		}
	case 0x7E:
		{ /* '0x7E' */
			return 0x7E
		}
	default:
		{
			return 0
		}
	}
}

func RequestTypeFirstEnumForFieldControlChar(value uint8) (enum RequestType, ok bool) {
	for _, sizeValue := range RequestTypeValues {
		if sizeValue.ControlChar() == value {
			return sizeValue, true
		}
	}
	return 0, false
}
func RequestTypeByValue(value uint8) (enum RequestType, ok bool) {
	switch value {
	case 0x00:
		return RequestType_UNKNOWN, true
	case 0x0D:
		return RequestType_EMPTY, true
	case 0x40:
		return RequestType_DIRECT_COMMAND, true
	case 0x5C:
		return RequestType_REQUEST_COMMAND, true
	case 0x6E:
		return RequestType_NULL, true
	case 0x7C:
		return RequestType_SMART_CONNECT_SHORTCUT, true
	case 0x7E:
		return RequestType_RESET, true
	}
	return 0, false
}

func RequestTypeByName(value string) (enum RequestType, ok bool) {
	switch value {
	case "UNKNOWN":
		return RequestType_UNKNOWN, true
	case "EMPTY":
		return RequestType_EMPTY, true
	case "DIRECT_COMMAND":
		return RequestType_DIRECT_COMMAND, true
	case "REQUEST_COMMAND":
		return RequestType_REQUEST_COMMAND, true
	case "NULL":
		return RequestType_NULL, true
	case "SMART_CONNECT_SHORTCUT":
		return RequestType_SMART_CONNECT_SHORTCUT, true
	case "RESET":
		return RequestType_RESET, true
	}
	return 0, false
}

func RequestTypeKnows(value uint8) bool {
	for _, typeValue := range RequestTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastRequestType(structType any) RequestType {
	castFunc := func(typ any) RequestType {
		if sRequestType, ok := typ.(RequestType); ok {
			return sRequestType
		}
		return 0
	}
	return castFunc(structType)
}

func (m RequestType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m RequestType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestTypeParse(ctx context.Context, theBytes []byte) (RequestType, error) {
	return RequestTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RequestTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RequestType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("RequestType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading RequestType")
	}
	if enum, ok := RequestTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for RequestType")
		return RequestType(val), nil
	} else {
		return enum, nil
	}
}

func (e RequestType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e RequestType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("RequestType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e RequestType) GetValue() uint8 {
	return uint8(e)
}

func (e RequestType) GetControlChar() uint8 {
	return e.ControlChar()
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e RequestType) PLC4XEnumName() string {
	switch e {
	case RequestType_UNKNOWN:
		return "UNKNOWN"
	case RequestType_EMPTY:
		return "EMPTY"
	case RequestType_DIRECT_COMMAND:
		return "DIRECT_COMMAND"
	case RequestType_REQUEST_COMMAND:
		return "REQUEST_COMMAND"
	case RequestType_NULL:
		return "NULL"
	case RequestType_SMART_CONNECT_SHORTCUT:
		return "SMART_CONNECT_SHORTCUT"
	case RequestType_RESET:
		return "RESET"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e RequestType) String() string {
	return e.PLC4XEnumName()
}
