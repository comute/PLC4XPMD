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

// ModbusPDUReadInputRegistersResponse is the corresponding interface of ModbusPDUReadInputRegistersResponse
type ModbusPDUReadInputRegistersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadInputRegistersResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadInputRegistersResponse.
// This is useful for switch cases.
type ModbusPDUReadInputRegistersResponseExactly interface {
	ModbusPDUReadInputRegistersResponse
	isModbusPDUReadInputRegistersResponse() bool
}

// _ModbusPDUReadInputRegistersResponse is the data-structure of this message
type _ModbusPDUReadInputRegistersResponse struct {
	*_ModbusPDU
	Value []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadInputRegistersResponse) GetFunctionFlag() uint8 {
	return 0x04
}

func (m *_ModbusPDUReadInputRegistersResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadInputRegistersResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadInputRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadInputRegistersResponse factory function for _ModbusPDUReadInputRegistersResponse
func NewModbusPDUReadInputRegistersResponse(value []byte) *_ModbusPDUReadInputRegistersResponse {
	_result := &_ModbusPDUReadInputRegistersResponse{
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadInputRegistersResponse(structType any) ModbusPDUReadInputRegistersResponse {
	if casted, ok := structType.(ModbusPDUReadInputRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadInputRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadInputRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadInputRegistersResponse"
}

func (m *_ModbusPDUReadInputRegistersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadInputRegistersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadInputRegistersResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReadInputRegistersResponse, error) {
	return ModbusPDUReadInputRegistersResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadInputRegistersResponseParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadInputRegistersResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadInputRegistersResponse, error) {
		return ModbusPDUReadInputRegistersResponseParseWithBuffer(ctx, readBuffer, response)
	}
}

func ModbusPDUReadInputRegistersResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadInputRegistersResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadInputRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadInputRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadInputRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadInputRegistersResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadInputRegistersResponse{
		_ModbusPDU: &_ModbusPDU{},
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadInputRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadInputRegistersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadInputRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadInputRegistersResponse")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("byteCount", 8, uint8((byteCount)))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadInputRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadInputRegistersResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadInputRegistersResponse) isModbusPDUReadInputRegistersResponse() bool {
	return true
}

func (m *_ModbusPDUReadInputRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
