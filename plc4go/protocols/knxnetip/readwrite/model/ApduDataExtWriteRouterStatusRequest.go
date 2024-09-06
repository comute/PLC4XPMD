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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtWriteRouterStatusRequest is the corresponding interface of ApduDataExtWriteRouterStatusRequest
type ApduDataExtWriteRouterStatusRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// IsApduDataExtWriteRouterStatusRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtWriteRouterStatusRequest()
}

// _ApduDataExtWriteRouterStatusRequest is the data-structure of this message
type _ApduDataExtWriteRouterStatusRequest struct {
	ApduDataExtContract
}

var _ ApduDataExtWriteRouterStatusRequest = (*_ApduDataExtWriteRouterStatusRequest)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtWriteRouterStatusRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtWriteRouterStatusRequest) GetExtApciType() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtWriteRouterStatusRequest) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtWriteRouterStatusRequest factory function for _ApduDataExtWriteRouterStatusRequest
func NewApduDataExtWriteRouterStatusRequest(length uint8) *_ApduDataExtWriteRouterStatusRequest {
	_result := &_ApduDataExtWriteRouterStatusRequest{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtWriteRouterStatusRequest(structType any) ApduDataExtWriteRouterStatusRequest {
	if casted, ok := structType.(ApduDataExtWriteRouterStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtWriteRouterStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetTypeName() string {
	return "ApduDataExtWriteRouterStatusRequest"
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtWriteRouterStatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtWriteRouterStatusRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtWriteRouterStatusRequest ApduDataExtWriteRouterStatusRequest, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRouterStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtWriteRouterStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRouterStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtWriteRouterStatusRequest")
	}

	return m, nil
}

func (m *_ApduDataExtWriteRouterStatusRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtWriteRouterStatusRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRouterStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtWriteRouterStatusRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRouterStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtWriteRouterStatusRequest")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtWriteRouterStatusRequest) IsApduDataExtWriteRouterStatusRequest() {}

func (m *_ApduDataExtWriteRouterStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
