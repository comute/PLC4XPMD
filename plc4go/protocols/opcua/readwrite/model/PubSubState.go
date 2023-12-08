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

// PubSubState is an enum
type PubSubState uint32

type IPubSubState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PubSubState_pubSubStateDisabled       PubSubState = 0
	PubSubState_pubSubStatePaused         PubSubState = 1
	PubSubState_pubSubStateOperational    PubSubState = 2
	PubSubState_pubSubStateError          PubSubState = 3
	PubSubState_pubSubStatePreOperational PubSubState = 4
)

var PubSubStateValues []PubSubState

func init() {
	_ = errors.New
	PubSubStateValues = []PubSubState{
		PubSubState_pubSubStateDisabled,
		PubSubState_pubSubStatePaused,
		PubSubState_pubSubStateOperational,
		PubSubState_pubSubStateError,
		PubSubState_pubSubStatePreOperational,
	}
}

func PubSubStateByValue(value uint32) (enum PubSubState, ok bool) {
	switch value {
	case 0:
		return PubSubState_pubSubStateDisabled, true
	case 1:
		return PubSubState_pubSubStatePaused, true
	case 2:
		return PubSubState_pubSubStateOperational, true
	case 3:
		return PubSubState_pubSubStateError, true
	case 4:
		return PubSubState_pubSubStatePreOperational, true
	}
	return 0, false
}

func PubSubStateByName(value string) (enum PubSubState, ok bool) {
	switch value {
	case "pubSubStateDisabled":
		return PubSubState_pubSubStateDisabled, true
	case "pubSubStatePaused":
		return PubSubState_pubSubStatePaused, true
	case "pubSubStateOperational":
		return PubSubState_pubSubStateOperational, true
	case "pubSubStateError":
		return PubSubState_pubSubStateError, true
	case "pubSubStatePreOperational":
		return PubSubState_pubSubStatePreOperational, true
	}
	return 0, false
}

func PubSubStateKnows(value uint32) bool {
	for _, typeValue := range PubSubStateValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPubSubState(structType any) PubSubState {
	castFunc := func(typ any) PubSubState {
		if sPubSubState, ok := typ.(PubSubState); ok {
			return sPubSubState
		}
		return 0
	}
	return castFunc(structType)
}

func (m PubSubState) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m PubSubState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubStateParse(ctx context.Context, theBytes []byte) (PubSubState, error) {
	return PubSubStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PubSubStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PubSubState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("PubSubState", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PubSubState")
	}
	if enum, ok := PubSubStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PubSubState")
		return PubSubState(val), nil
	} else {
		return enum, nil
	}
}

func (e PubSubState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PubSubState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("PubSubState", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PubSubState) PLC4XEnumName() string {
	switch e {
	case PubSubState_pubSubStateDisabled:
		return "pubSubStateDisabled"
	case PubSubState_pubSubStatePaused:
		return "pubSubStatePaused"
	case PubSubState_pubSubStateOperational:
		return "pubSubStateOperational"
	case PubSubState_pubSubStateError:
		return "pubSubStateError"
	case PubSubState_pubSubStatePreOperational:
		return "pubSubStatePreOperational"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e PubSubState) String() string {
	return e.PLC4XEnumName()
}
