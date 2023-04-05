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
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUReadDiscreteInputsResponse is the corresponding interface of ModbusPDUReadDiscreteInputsResponse
type ModbusPDUReadDiscreteInputsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadDiscreteInputsResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadDiscreteInputsResponse.
// This is useful for switch cases.
type ModbusPDUReadDiscreteInputsResponseExactly interface {
	ModbusPDUReadDiscreteInputsResponse
	isModbusPDUReadDiscreteInputsResponse() bool
}

// _ModbusPDUReadDiscreteInputsResponse is the data-structure of this message
type _ModbusPDUReadDiscreteInputsResponse struct {
	*_ModbusPDU
	Value []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetFunctionFlag() uint8 {
	return 0x02
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDiscreteInputsResponse factory function for _ModbusPDUReadDiscreteInputsResponse
func NewModbusPDUReadDiscreteInputsResponse(value []byte) *_ModbusPDUReadDiscreteInputsResponse {
	_result := &_ModbusPDUReadDiscreteInputsResponse{
		Value:      value,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadDiscreteInputsResponse(structType interface{}) ModbusPDUReadDiscreteInputsResponse {
	if casted, ok := structType.(ModbusPDUReadDiscreteInputsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadDiscreteInputsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetTypeName() string {
	return "ModbusPDUReadDiscreteInputsResponse"
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadDiscreteInputsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadDiscreteInputsResponseParse(theBytes []byte, response bool) (ModbusPDUReadDiscreteInputsResponse, error) {
	return ModbusPDUReadDiscreteInputsResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadDiscreteInputsResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadDiscreteInputsResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDiscreteInputsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadDiscreteInputsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReadDiscreteInputsResponse")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUReadDiscreteInputsResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDiscreteInputsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadDiscreteInputsResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadDiscreteInputsResponse{
		_ModbusPDU: &_ModbusPDU{},
		Value:      value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadDiscreteInputsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDiscreteInputsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadDiscreteInputsResponse")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		// Byte Array field (value)
		if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDiscreteInputsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadDiscreteInputsResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadDiscreteInputsResponse) isModbusPDUReadDiscreteInputsResponse() bool {
	return true
}

func (m *_ModbusPDUReadDiscreteInputsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
