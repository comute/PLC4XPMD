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

// LineOffHookReason is an enum
type LineOffHookReason uint8

type ILineOffHookReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	LineOffHookReason_INCOMING_VOICE_CALL LineOffHookReason = 0x01
	LineOffHookReason_INCOMING_DATA_CALL  LineOffHookReason = 0x02
	LineOffHookReason_INCOMING_CALL       LineOffHookReason = 0x03
	LineOffHookReason_OUTGOING_VOICE_CALL LineOffHookReason = 0x10
	LineOffHookReason_OUTGOING_DATA_CALL  LineOffHookReason = 0x20
	LineOffHookReason_OUTGOING_CALL       LineOffHookReason = 0x30
	LineOffHookReason_CBTI_IS_SETTING     LineOffHookReason = 0x40
	LineOffHookReason_CBTI_IS_CLEARING    LineOffHookReason = 0x50
)

var LineOffHookReasonValues []LineOffHookReason

func init() {
	_ = errors.New
	LineOffHookReasonValues = []LineOffHookReason{
		LineOffHookReason_INCOMING_VOICE_CALL,
		LineOffHookReason_INCOMING_DATA_CALL,
		LineOffHookReason_INCOMING_CALL,
		LineOffHookReason_OUTGOING_VOICE_CALL,
		LineOffHookReason_OUTGOING_DATA_CALL,
		LineOffHookReason_OUTGOING_CALL,
		LineOffHookReason_CBTI_IS_SETTING,
		LineOffHookReason_CBTI_IS_CLEARING,
	}
}

func LineOffHookReasonByValue(value uint8) (enum LineOffHookReason, ok bool) {
	switch value {
	case 0x01:
		return LineOffHookReason_INCOMING_VOICE_CALL, true
	case 0x02:
		return LineOffHookReason_INCOMING_DATA_CALL, true
	case 0x03:
		return LineOffHookReason_INCOMING_CALL, true
	case 0x10:
		return LineOffHookReason_OUTGOING_VOICE_CALL, true
	case 0x20:
		return LineOffHookReason_OUTGOING_DATA_CALL, true
	case 0x30:
		return LineOffHookReason_OUTGOING_CALL, true
	case 0x40:
		return LineOffHookReason_CBTI_IS_SETTING, true
	case 0x50:
		return LineOffHookReason_CBTI_IS_CLEARING, true
	}
	return 0, false
}

func LineOffHookReasonByName(value string) (enum LineOffHookReason, ok bool) {
	switch value {
	case "INCOMING_VOICE_CALL":
		return LineOffHookReason_INCOMING_VOICE_CALL, true
	case "INCOMING_DATA_CALL":
		return LineOffHookReason_INCOMING_DATA_CALL, true
	case "INCOMING_CALL":
		return LineOffHookReason_INCOMING_CALL, true
	case "OUTGOING_VOICE_CALL":
		return LineOffHookReason_OUTGOING_VOICE_CALL, true
	case "OUTGOING_DATA_CALL":
		return LineOffHookReason_OUTGOING_DATA_CALL, true
	case "OUTGOING_CALL":
		return LineOffHookReason_OUTGOING_CALL, true
	case "CBTI_IS_SETTING":
		return LineOffHookReason_CBTI_IS_SETTING, true
	case "CBTI_IS_CLEARING":
		return LineOffHookReason_CBTI_IS_CLEARING, true
	}
	return 0, false
}

func LineOffHookReasonKnows(value uint8) bool {
	for _, typeValue := range LineOffHookReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLineOffHookReason(structType any) LineOffHookReason {
	castFunc := func(typ any) LineOffHookReason {
		if sLineOffHookReason, ok := typ.(LineOffHookReason); ok {
			return sLineOffHookReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m LineOffHookReason) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m LineOffHookReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LineOffHookReasonParse(ctx context.Context, theBytes []byte) (LineOffHookReason, error) {
	return LineOffHookReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LineOffHookReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LineOffHookReason, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("LineOffHookReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LineOffHookReason")
	}
	if enum, ok := LineOffHookReasonByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for LineOffHookReason")
		return LineOffHookReason(val), nil
	} else {
		return enum, nil
	}
}

func (e LineOffHookReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LineOffHookReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("LineOffHookReason", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LineOffHookReason) PLC4XEnumName() string {
	switch e {
	case LineOffHookReason_INCOMING_VOICE_CALL:
		return "INCOMING_VOICE_CALL"
	case LineOffHookReason_INCOMING_DATA_CALL:
		return "INCOMING_DATA_CALL"
	case LineOffHookReason_INCOMING_CALL:
		return "INCOMING_CALL"
	case LineOffHookReason_OUTGOING_VOICE_CALL:
		return "OUTGOING_VOICE_CALL"
	case LineOffHookReason_OUTGOING_DATA_CALL:
		return "OUTGOING_DATA_CALL"
	case LineOffHookReason_OUTGOING_CALL:
		return "OUTGOING_CALL"
	case LineOffHookReason_CBTI_IS_SETTING:
		return "CBTI_IS_SETTING"
	case LineOffHookReason_CBTI_IS_CLEARING:
		return "CBTI_IS_CLEARING"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e LineOffHookReason) String() string {
	return e.PLC4XEnumName()
}
