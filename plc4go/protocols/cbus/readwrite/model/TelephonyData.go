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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TelephonyData is the corresponding interface of TelephonyData
type TelephonyData interface {
	TelephonyDataContract
	TelephonyDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsTelephonyData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyData()
}

// TelephonyDataContract provides a set of functions which can be overwritten by a sub struct
type TelephonyDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() TelephonyCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() TelephonyCommandType
	// IsTelephonyData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyData()
}

// TelephonyDataRequirements provides a set of functions which need to be implemented by a sub struct
type TelephonyDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetArgument returns Argument (discriminator field)
	GetArgument() byte
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() TelephonyCommandType
}

// _TelephonyData is the data-structure of this message
type _TelephonyData struct {
	_SubType             TelephonyData
	CommandTypeContainer TelephonyCommandTypeContainer
	Argument             byte
}

var _ TelephonyDataContract = (*_TelephonyData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyData) GetCommandTypeContainer() TelephonyCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_TelephonyData) GetArgument() byte {
	return m.Argument
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_TelephonyData) GetCommandType() TelephonyCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastTelephonyCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyData factory function for _TelephonyData
func NewTelephonyData(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyData {
	return &_TelephonyData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastTelephonyData(structType any) TelephonyData {
	if casted, ok := structType.(TelephonyData); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyData); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyData) GetTypeName() string {
	return "TelephonyData"
}

func (m *_TelephonyData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func TelephonyDataParse[T TelephonyData](ctx context.Context, theBytes []byte) (T, error) {
	return TelephonyDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataParseWithBufferProducer[T TelephonyData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := TelephonyDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func TelephonyDataParseWithBuffer[T TelephonyData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_TelephonyData{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_TelephonyData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__telephonyData TelephonyData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsTelephonyCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[TelephonyCommandTypeContainer](ctx, "commandTypeContainer", "TelephonyCommandTypeContainer", ReadEnum(TelephonyCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[TelephonyCommandType](ctx, "commandType", (*TelephonyCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	argument, err := ReadSimpleField(ctx, "argument", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'argument' field"))
	}
	m.Argument = argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child TelephonyData
	switch {
	case commandType == TelephonyCommandType_EVENT && argument == 0x01: // TelephonyDataLineOnHook
		if _child, err = (&_TelephonyDataLineOnHook{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataLineOnHook for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x02: // TelephonyDataLineOffHook
		if _child, err = (&_TelephonyDataLineOffHook{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataLineOffHook for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x03: // TelephonyDataDialOutFailure
		if _child, err = (&_TelephonyDataDialOutFailure{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataDialOutFailure for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x04: // TelephonyDataDialInFailure
		if _child, err = (&_TelephonyDataDialInFailure{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataDialInFailure for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x05: // TelephonyDataRinging
		if _child, err = (&_TelephonyDataRinging{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataRinging for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x06: // TelephonyDataRecallLastNumber
		if _child, err = (&_TelephonyDataRecallLastNumber{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataRecallLastNumber for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x07: // TelephonyDataInternetConnectionRequestMade
		if _child, err = (&_TelephonyDataInternetConnectionRequestMade{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataInternetConnectionRequestMade for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x80: // TelephonyDataIsolateSecondaryOutlet
		if _child, err = (&_TelephonyDataIsolateSecondaryOutlet{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataIsolateSecondaryOutlet for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x81: // TelephonyDataRecallLastNumberRequest
		if _child, err = (&_TelephonyDataRecallLastNumberRequest{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataRecallLastNumberRequest for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x82: // TelephonyDataRejectIncomingCall
		if _child, err = (&_TelephonyDataRejectIncomingCall{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataRejectIncomingCall for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x83: // TelephonyDataDivert
		if _child, err = (&_TelephonyDataDivert{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataDivert for type-switch of TelephonyData")
		}
	case commandType == TelephonyCommandType_EVENT && argument == 0x84: // TelephonyDataClearDiversion
		if _child, err = (&_TelephonyDataClearDiversion{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type TelephonyDataClearDiversion for type-switch of TelephonyData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}

	if closeErr := readBuffer.CloseContext("TelephonyData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyData")
	}

	return _child, nil
}

func (pm *_TelephonyData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TelephonyData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TelephonyData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TelephonyData")
	}

	if err := WriteSimpleEnumField[TelephonyCommandTypeContainer](ctx, "commandTypeContainer", "TelephonyCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[TelephonyCommandTypeContainer, uint8](TelephonyCommandTypeContainer.GetValue, TelephonyCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "argument", m.GetArgument(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("TelephonyData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TelephonyData")
	}
	return nil
}

func (m *_TelephonyData) IsTelephonyData() {}
