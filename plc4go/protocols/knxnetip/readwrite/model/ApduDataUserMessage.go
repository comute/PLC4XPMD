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

// ApduDataUserMessage is the corresponding interface of ApduDataUserMessage
type ApduDataUserMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataUserMessageExactly can be used when we want exactly this type and not a type which fulfills ApduDataUserMessage.
// This is useful for switch cases.
type ApduDataUserMessageExactly interface {
	ApduDataUserMessage
	isApduDataUserMessage() bool
}

// _ApduDataUserMessage is the data-structure of this message
type _ApduDataUserMessage struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataUserMessage) GetApciType() uint8 {
	return 0xB
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataUserMessage) InitializeParent(parent ApduData) {}

func (m *_ApduDataUserMessage) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataUserMessage factory function for _ApduDataUserMessage
func NewApduDataUserMessage(dataLength uint8) *_ApduDataUserMessage {
	_result := &_ApduDataUserMessage{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataUserMessage(structType interface{}) ApduDataUserMessage {
	if casted, ok := structType.(ApduDataUserMessage); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataUserMessage); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataUserMessage) GetTypeName() string {
	return "ApduDataUserMessage"
}

func (m *_ApduDataUserMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataUserMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataUserMessageParse(theBytes []byte, dataLength uint8) (ApduDataUserMessage, error) {
	return ApduDataUserMessageParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataUserMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataUserMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataUserMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataUserMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataUserMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataUserMessage")
	}

	// Create a partially initialized instance
	_child := &_ApduDataUserMessage{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataUserMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataUserMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataUserMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataUserMessage")
		}

		if popErr := writeBuffer.PopContext("ApduDataUserMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataUserMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataUserMessage) isApduDataUserMessage() bool {
	return true
}

func (m *_ApduDataUserMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
