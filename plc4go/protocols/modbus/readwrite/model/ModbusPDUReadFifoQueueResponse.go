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

// ModbusPDUReadFifoQueueResponse is the corresponding interface of ModbusPDUReadFifoQueueResponse
type ModbusPDUReadFifoQueueResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetFifoValue returns FifoValue (property field)
	GetFifoValue() []uint16
}

// ModbusPDUReadFifoQueueResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFifoQueueResponse.
// This is useful for switch cases.
type ModbusPDUReadFifoQueueResponseExactly interface {
	ModbusPDUReadFifoQueueResponse
	isModbusPDUReadFifoQueueResponse() bool
}

// _ModbusPDUReadFifoQueueResponse is the data-structure of this message
type _ModbusPDUReadFifoQueueResponse struct {
	*_ModbusPDU
	FifoValue []uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFifoQueueResponse) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *_ModbusPDUReadFifoQueueResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadFifoQueueResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFifoQueueResponse) GetFifoValue() []uint16 {
	return m.FifoValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFifoQueueResponse factory function for _ModbusPDUReadFifoQueueResponse
func NewModbusPDUReadFifoQueueResponse(fifoValue []uint16) *_ModbusPDUReadFifoQueueResponse {
	_result := &_ModbusPDUReadFifoQueueResponse{
		FifoValue:  fifoValue,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFifoQueueResponse(structType any) ModbusPDUReadFifoQueueResponse {
	if casted, ok := structType.(ModbusPDUReadFifoQueueResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFifoQueueResponse) GetTypeName() string {
	return "ModbusPDUReadFifoQueueResponse"
}

func (m *_ModbusPDUReadFifoQueueResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 16

	// Implicit Field (fifoCount)
	lengthInBits += 16

	// Array field
	if len(m.FifoValue) > 0 {
		lengthInBits += 16 * uint16(len(m.FifoValue))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFifoQueueResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadFifoQueueResponseParse(theBytes []byte, response bool) (ModbusPDUReadFifoQueueResponse, error) {
	return ModbusPDUReadFifoQueueResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadFifoQueueResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadFifoQueueResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFifoQueueResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint16("byteCount", 16)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReadFifoQueueResponse")
	}

	// Implicit Field (fifoCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	fifoCount, _fifoCountErr := readBuffer.ReadUint16("fifoCount", 16)
	_ = fifoCount
	if _fifoCountErr != nil {
		return nil, errors.Wrap(_fifoCountErr, "Error parsing 'fifoCount' field of ModbusPDUReadFifoQueueResponse")
	}

	// Array field (fifoValue)
	if pullErr := readBuffer.PullContext("fifoValue", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fifoValue")
	}
	// Count array
	fifoValue := make([]uint16, fifoCount)
	// This happens when the size is set conditional to 0
	if len(fifoValue) == 0 {
		fifoValue = nil
	}
	{
		_numItems := uint16(fifoCount)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'fifoValue' field of ModbusPDUReadFifoQueueResponse")
			}
			fifoValue[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("fifoValue", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fifoValue")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFifoQueueResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadFifoQueueResponse{
		_ModbusPDU: &_ModbusPDU{},
		FifoValue:  fifoValue,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadFifoQueueResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFifoQueueResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFifoQueueResponse")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint16(uint16((uint16(uint16(len(m.GetFifoValue()))) * uint16(uint16(2)))) + uint16(uint16(2)))
		_byteCountErr := writeBuffer.WriteUint16("byteCount", 16, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Implicit Field (fifoCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		fifoCount := uint16(uint16((uint16(uint16(len(m.GetFifoValue()))) * uint16(uint16(2)))) / uint16(uint16(2)))
		_fifoCountErr := writeBuffer.WriteUint16("fifoCount", 16, (fifoCount))
		if _fifoCountErr != nil {
			return errors.Wrap(_fifoCountErr, "Error serializing 'fifoCount' field")
		}

		// Array Field (fifoValue)
		if pushErr := writeBuffer.PushContext("fifoValue", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fifoValue")
		}
		for _curItem, _element := range m.GetFifoValue() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint16("", 16, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'fifoValue' field")
			}
		}
		if popErr := writeBuffer.PopContext("fifoValue", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fifoValue")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFifoQueueResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFifoQueueResponse) isModbusPDUReadFifoQueueResponse() bool {
	return true
}

func (m *_ModbusPDUReadFifoQueueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
