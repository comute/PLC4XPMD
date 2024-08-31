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

// MFuncPropCommandReq is the corresponding interface of MFuncPropCommandReq
type MFuncPropCommandReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MFuncPropCommandReqExactly can be used when we want exactly this type and not a type which fulfills MFuncPropCommandReq.
// This is useful for switch cases.
type MFuncPropCommandReqExactly interface {
	MFuncPropCommandReq
	isMFuncPropCommandReq() bool
}

// _MFuncPropCommandReq is the data-structure of this message
type _MFuncPropCommandReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropCommandReq) GetMessageCode() uint8 {
	return 0xF8
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropCommandReq) InitializeParent(parent CEMI) {}

func (m *_MFuncPropCommandReq) GetParent() CEMI {
	return m._CEMI
}

// NewMFuncPropCommandReq factory function for _MFuncPropCommandReq
func NewMFuncPropCommandReq(size uint16) *_MFuncPropCommandReq {
	_result := &_MFuncPropCommandReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMFuncPropCommandReq(structType any) MFuncPropCommandReq {
	if casted, ok := structType.(MFuncPropCommandReq); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropCommandReq); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropCommandReq) GetTypeName() string {
	return "MFuncPropCommandReq"
}

func (m *_MFuncPropCommandReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropCommandReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MFuncPropCommandReqParse(ctx context.Context, theBytes []byte, size uint16) (MFuncPropCommandReq, error) {
	return MFuncPropCommandReqParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func MFuncPropCommandReqParseWithBufferProducer(size uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (MFuncPropCommandReq, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MFuncPropCommandReq, error) {
		return MFuncPropCommandReqParseWithBuffer(ctx, readBuffer, size)
	}
}

func MFuncPropCommandReqParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (MFuncPropCommandReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropCommandReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropCommandReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCommandReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropCommandReq")
	}

	// Create a partially initialized instance
	_child := &_MFuncPropCommandReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MFuncPropCommandReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropCommandReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCommandReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropCommandReq")
		}

		if popErr := writeBuffer.PopContext("MFuncPropCommandReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropCommandReq")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropCommandReq) isMFuncPropCommandReq() bool {
	return true
}

func (m *_MFuncPropCommandReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
