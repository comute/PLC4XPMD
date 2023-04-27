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

// SetAttributeSingleResponse is the corresponding interface of SetAttributeSingleResponse
type SetAttributeSingleResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// SetAttributeSingleResponseExactly can be used when we want exactly this type and not a type which fulfills SetAttributeSingleResponse.
// This is useful for switch cases.
type SetAttributeSingleResponseExactly interface {
	SetAttributeSingleResponse
	isSetAttributeSingleResponse() bool
}

// _SetAttributeSingleResponse is the data-structure of this message
type _SetAttributeSingleResponse struct {
	*_CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeSingleResponse) GetService() uint8 {
	return 0x10
}

func (m *_SetAttributeSingleResponse) GetResponse() bool {
	return bool(true)
}

func (m *_SetAttributeSingleResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeSingleResponse) InitializeParent(parent CipService) {}

func (m *_SetAttributeSingleResponse) GetParent() CipService {
	return m._CipService
}

// NewSetAttributeSingleResponse factory function for _SetAttributeSingleResponse
func NewSetAttributeSingleResponse(serviceLen uint16) *_SetAttributeSingleResponse {
	_result := &_SetAttributeSingleResponse{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetAttributeSingleResponse(structType any) SetAttributeSingleResponse {
	if casted, ok := structType.(SetAttributeSingleResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeSingleResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeSingleResponse) GetTypeName() string {
	return "SetAttributeSingleResponse"
}

func (m *_SetAttributeSingleResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeSingleResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetAttributeSingleResponseParse(theBytes []byte, connected bool, serviceLen uint16) (SetAttributeSingleResponse, error) {
	return SetAttributeSingleResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func SetAttributeSingleResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (SetAttributeSingleResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeSingleResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeSingleResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeSingleResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeSingleResponse")
	}

	// Create a partially initialized instance
	_child := &_SetAttributeSingleResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_SetAttributeSingleResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeSingleResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeSingleResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeSingleResponse")
		}

		if popErr := writeBuffer.PopContext("SetAttributeSingleResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeSingleResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeSingleResponse) isSetAttributeSingleResponse() bool {
	return true
}

func (m *_SetAttributeSingleResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
