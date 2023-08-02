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

// BACnetAction is an enum
type BACnetAction uint8

type IBACnetAction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAction_DIRECT  BACnetAction = 0
	BACnetAction_REVERSE BACnetAction = 1
)

var BACnetActionValues []BACnetAction

func init() {
	_ = errors.New
	BACnetActionValues = []BACnetAction{
		BACnetAction_DIRECT,
		BACnetAction_REVERSE,
	}
}

func BACnetActionByValue(value uint8) (enum BACnetAction, ok bool) {
	switch value {
	case 0:
		return BACnetAction_DIRECT, true
	case 1:
		return BACnetAction_REVERSE, true
	}
	return 0, false
}

func BACnetActionByName(value string) (enum BACnetAction, ok bool) {
	switch value {
	case "DIRECT":
		return BACnetAction_DIRECT, true
	case "REVERSE":
		return BACnetAction_REVERSE, true
	}
	return 0, false
}

func BACnetActionKnows(value uint8) bool {
	for _, typeValue := range BACnetActionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAction(structType any) BACnetAction {
	castFunc := func(typ any) BACnetAction {
		if sBACnetAction, ok := typ.(BACnetAction); ok {
			return sBACnetAction
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAction) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetActionParse(ctx context.Context, theBytes []byte) (BACnetAction, error) {
	return BACnetActionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetActionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAction, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAction", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAction")
	}
	if enum, ok := BACnetActionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAction")
		return BACnetAction(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAction", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAction) PLC4XEnumName() string {
	switch e {
	case BACnetAction_DIRECT:
		return "DIRECT"
	case BACnetAction_REVERSE:
		return "REVERSE"
	}
	return ""
}

func (e BACnetAction) String() string {
	return e.PLC4XEnumName()
}
