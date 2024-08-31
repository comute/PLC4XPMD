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

// Constant values.
const ModbusPDUReadDeviceIdentificationRequest_MEITYPE uint8 = 0x0E

// ModbusPDUReadDeviceIdentificationRequest is the corresponding interface of ModbusPDUReadDeviceIdentificationRequest
type ModbusPDUReadDeviceIdentificationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetLevel returns Level (property field)
	GetLevel() ModbusDeviceInformationLevel
	// GetObjectId returns ObjectId (property field)
	GetObjectId() uint8
}

// ModbusPDUReadDeviceIdentificationRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadDeviceIdentificationRequest.
// This is useful for switch cases.
type ModbusPDUReadDeviceIdentificationRequestExactly interface {
	ModbusPDUReadDeviceIdentificationRequest
	isModbusPDUReadDeviceIdentificationRequest() bool
}

// _ModbusPDUReadDeviceIdentificationRequest is the data-structure of this message
type _ModbusPDUReadDeviceIdentificationRequest struct {
	*_ModbusPDU
	Level    ModbusDeviceInformationLevel
	ObjectId uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetFunctionFlag() uint8 {
	return 0x2B
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadDeviceIdentificationRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetLevel() ModbusDeviceInformationLevel {
	return m.Level
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetObjectId() uint8 {
	return m.ObjectId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetMeiType() uint8 {
	return ModbusPDUReadDeviceIdentificationRequest_MEITYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDeviceIdentificationRequest factory function for _ModbusPDUReadDeviceIdentificationRequest
func NewModbusPDUReadDeviceIdentificationRequest(level ModbusDeviceInformationLevel, objectId uint8) *_ModbusPDUReadDeviceIdentificationRequest {
	_result := &_ModbusPDUReadDeviceIdentificationRequest{
		Level:      level,
		ObjectId:   objectId,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadDeviceIdentificationRequest(structType any) ModbusPDUReadDeviceIdentificationRequest {
	if casted, ok := structType.(ModbusPDUReadDeviceIdentificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadDeviceIdentificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationRequest"
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (meiType)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	// Simple field (objectId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadDeviceIdentificationRequestParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadDeviceIdentificationRequest, error) {
	return ModbusPDUReadDeviceIdentificationRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadDeviceIdentificationRequestParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadDeviceIdentificationRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadDeviceIdentificationRequest, error) {
		return ModbusPDUReadDeviceIdentificationRequestParseWithBuffer(ctx, readBuffer, response)
	}
}

func ModbusPDUReadDeviceIdentificationRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadDeviceIdentificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDeviceIdentificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadDeviceIdentificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	meiType, err := ReadConstField[uint8](ctx, "meiType", ReadUnsignedByte(readBuffer, uint8(8)), ModbusPDUReadDeviceIdentificationRequest_MEITYPE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'meiType' field"))
	}
	_ = meiType

	level, err := ReadEnumField[ModbusDeviceInformationLevel](ctx, "level", "ModbusDeviceInformationLevel", ReadEnum(ModbusDeviceInformationLevelByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}

	objectId, err := ReadSimpleField(ctx, "objectId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectId' field"))
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDeviceIdentificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadDeviceIdentificationRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadDeviceIdentificationRequest{
		_ModbusPDU: &_ModbusPDU{},
		Level:      level,
		ObjectId:   objectId,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDeviceIdentificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadDeviceIdentificationRequest")
		}

		// Const Field (meiType)
		_meiTypeErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("meiType", 8, uint8(0x0E))
		if _meiTypeErr != nil {
			return errors.Wrap(_meiTypeErr, "Error serializing 'meiType' field")
		}

		// Simple Field (level)
		if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for level")
		}
		_levelErr := writeBuffer.WriteSerializable(ctx, m.GetLevel())
		if popErr := writeBuffer.PopContext("level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for level")
		}
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		// Simple Field (objectId)
		objectId := uint8(m.GetObjectId())
		_objectIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("objectId", 8, uint8((objectId)))
		if _objectIdErr != nil {
			return errors.Wrap(_objectIdErr, "Error serializing 'objectId' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDeviceIdentificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadDeviceIdentificationRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) isModbusPDUReadDeviceIdentificationRequest() bool {
	return true
}

func (m *_ModbusPDUReadDeviceIdentificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
