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

// MessageSecurityMode is an enum
type MessageSecurityMode uint32

type IMessageSecurityMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	MessageSecurityMode_messageSecurityModeInvalid        MessageSecurityMode = 0
	MessageSecurityMode_messageSecurityModeNone           MessageSecurityMode = 1
	MessageSecurityMode_messageSecurityModeSign           MessageSecurityMode = 2
	MessageSecurityMode_messageSecurityModeSignAndEncrypt MessageSecurityMode = 3
)

var MessageSecurityModeValues []MessageSecurityMode

func init() {
	_ = errors.New
	MessageSecurityModeValues = []MessageSecurityMode{
		MessageSecurityMode_messageSecurityModeInvalid,
		MessageSecurityMode_messageSecurityModeNone,
		MessageSecurityMode_messageSecurityModeSign,
		MessageSecurityMode_messageSecurityModeSignAndEncrypt,
	}
}

func MessageSecurityModeByValue(value uint32) (enum MessageSecurityMode, ok bool) {
	switch value {
	case 0:
		return MessageSecurityMode_messageSecurityModeInvalid, true
	case 1:
		return MessageSecurityMode_messageSecurityModeNone, true
	case 2:
		return MessageSecurityMode_messageSecurityModeSign, true
	case 3:
		return MessageSecurityMode_messageSecurityModeSignAndEncrypt, true
	}
	return 0, false
}

func MessageSecurityModeByName(value string) (enum MessageSecurityMode, ok bool) {
	switch value {
	case "messageSecurityModeInvalid":
		return MessageSecurityMode_messageSecurityModeInvalid, true
	case "messageSecurityModeNone":
		return MessageSecurityMode_messageSecurityModeNone, true
	case "messageSecurityModeSign":
		return MessageSecurityMode_messageSecurityModeSign, true
	case "messageSecurityModeSignAndEncrypt":
		return MessageSecurityMode_messageSecurityModeSignAndEncrypt, true
	}
	return 0, false
}

func MessageSecurityModeKnows(value uint32) bool {
	for _, typeValue := range MessageSecurityModeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMessageSecurityMode(structType any) MessageSecurityMode {
	castFunc := func(typ any) MessageSecurityMode {
		if sMessageSecurityMode, ok := typ.(MessageSecurityMode); ok {
			return sMessageSecurityMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m MessageSecurityMode) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m MessageSecurityMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MessageSecurityModeParse(ctx context.Context, theBytes []byte) (MessageSecurityMode, error) {
	return MessageSecurityModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MessageSecurityModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MessageSecurityMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("MessageSecurityMode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MessageSecurityMode")
	}
	if enum, ok := MessageSecurityModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for MessageSecurityMode")
		return MessageSecurityMode(val), nil
	} else {
		return enum, nil
	}
}

func (e MessageSecurityMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MessageSecurityMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("MessageSecurityMode", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MessageSecurityMode) PLC4XEnumName() string {
	switch e {
	case MessageSecurityMode_messageSecurityModeInvalid:
		return "messageSecurityModeInvalid"
	case MessageSecurityMode_messageSecurityModeNone:
		return "messageSecurityModeNone"
	case MessageSecurityMode_messageSecurityModeSign:
		return "messageSecurityModeSign"
	case MessageSecurityMode_messageSecurityModeSignAndEncrypt:
		return "messageSecurityModeSignAndEncrypt"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e MessageSecurityMode) String() string {
	return e.PLC4XEnumName()
}
