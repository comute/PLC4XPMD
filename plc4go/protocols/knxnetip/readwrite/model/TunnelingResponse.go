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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TunnelingResponse is the corresponding interface of TunnelingResponse
type TunnelingResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetTunnelingResponseDataBlock returns TunnelingResponseDataBlock (property field)
	GetTunnelingResponseDataBlock() TunnelingResponseDataBlock
}

// TunnelingResponseExactly can be used when we want exactly this type and not a type which fulfills TunnelingResponse.
// This is useful for switch cases.
type TunnelingResponseExactly interface {
	TunnelingResponse
	isTunnelingResponse() bool
}

// _TunnelingResponse is the data-structure of this message
type _TunnelingResponse struct {
	*_KnxNetIpMessage
	TunnelingResponseDataBlock TunnelingResponseDataBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TunnelingResponse) GetMsgType() uint16 {
	return 0x0421
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TunnelingResponse) InitializeParent(parent KnxNetIpMessage) {}

func (m *_TunnelingResponse) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingResponse) GetTunnelingResponseDataBlock() TunnelingResponseDataBlock {
	return m.TunnelingResponseDataBlock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTunnelingResponse factory function for _TunnelingResponse
func NewTunnelingResponse(tunnelingResponseDataBlock TunnelingResponseDataBlock) *_TunnelingResponse {
	_result := &_TunnelingResponse{
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
		_KnxNetIpMessage:           NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTunnelingResponse(structType interface{}) TunnelingResponse {
	if casted, ok := structType.(TunnelingResponse); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingResponse); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingResponse) GetTypeName() string {
	return "TunnelingResponse"
}

func (m *_TunnelingResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (tunnelingResponseDataBlock)
	lengthInBits += m.TunnelingResponseDataBlock.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TunnelingResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TunnelingResponseParse(theBytes []byte) (TunnelingResponse, error) {
	return TunnelingResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func TunnelingResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TunnelingResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tunnelingResponseDataBlock)
	if pullErr := readBuffer.PullContext("tunnelingResponseDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tunnelingResponseDataBlock")
	}
	_tunnelingResponseDataBlock, _tunnelingResponseDataBlockErr := TunnelingResponseDataBlockParseWithBuffer(ctx, readBuffer)
	if _tunnelingResponseDataBlockErr != nil {
		return nil, errors.Wrap(_tunnelingResponseDataBlockErr, "Error parsing 'tunnelingResponseDataBlock' field of TunnelingResponse")
	}
	tunnelingResponseDataBlock := _tunnelingResponseDataBlock.(TunnelingResponseDataBlock)
	if closeErr := readBuffer.CloseContext("tunnelingResponseDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tunnelingResponseDataBlock")
	}

	if closeErr := readBuffer.CloseContext("TunnelingResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingResponse")
	}

	// Create a partially initialized instance
	_child := &_TunnelingResponse{
		_KnxNetIpMessage:           &_KnxNetIpMessage{},
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_TunnelingResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TunnelingResponse")
		}

		// Simple Field (tunnelingResponseDataBlock)
		if pushErr := writeBuffer.PushContext("tunnelingResponseDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for tunnelingResponseDataBlock")
		}
		_tunnelingResponseDataBlockErr := writeBuffer.WriteSerializable(ctx, m.GetTunnelingResponseDataBlock())
		if popErr := writeBuffer.PopContext("tunnelingResponseDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for tunnelingResponseDataBlock")
		}
		if _tunnelingResponseDataBlockErr != nil {
			return errors.Wrap(_tunnelingResponseDataBlockErr, "Error serializing 'tunnelingResponseDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TunnelingResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TunnelingResponse) isTunnelingResponse() bool {
	return true
}

func (m *_TunnelingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
