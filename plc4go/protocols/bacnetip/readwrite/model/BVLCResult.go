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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCResult is the corresponding interface of BVLCResult
type BVLCResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetCode returns Code (property field)
	GetCode() BVLCResultCode
}

// BVLCResultExactly can be used when we want exactly this type and not a type which fulfills BVLCResult.
// This is useful for switch cases.
type BVLCResultExactly interface {
	BVLCResult
	isBVLCResult() bool
}

// _BVLCResult is the data-structure of this message
type _BVLCResult struct {
	*_BVLC
	Code BVLCResultCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCResult) GetBvlcFunction() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCResult) InitializeParent(parent BVLC) {}

func (m *_BVLCResult) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCResult) GetCode() BVLCResultCode {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCResult factory function for _BVLCResult
func NewBVLCResult(code BVLCResultCode) *_BVLCResult {
	_result := &_BVLCResult{
		Code:  code,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCResult(structType any) BVLCResult {
	if casted, ok := structType.(BVLCResult); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCResult); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCResult) GetTypeName() string {
	return "BVLCResult"
}

func (m *_BVLCResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (code)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCResultParse(ctx context.Context, theBytes []byte) (BVLCResult, error) {
	return BVLCResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func BVLCResultParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCResult, error) {
		return BVLCResultParseWithBuffer(ctx, readBuffer)
	}
}

func BVLCResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	code, err := ReadEnumField[BVLCResultCode](ctx, "code", "BVLCResultCode", ReadEnum(BVLCResultCodeByValue, ReadUnsignedShort(readBuffer, uint8(16))), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}

	if closeErr := readBuffer.CloseContext("BVLCResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCResult")
	}

	// Create a partially initialized instance
	_child := &_BVLCResult{
		_BVLC: &_BVLC{},
		Code:  code,
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCResult")
		}

		// Simple Field (code)
		if pushErr := writeBuffer.PushContext("code"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for code")
		}
		_codeErr := writeBuffer.WriteSerializable(ctx, m.GetCode())
		if popErr := writeBuffer.PopContext("code"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for code")
		}
		if _codeErr != nil {
			return errors.Wrap(_codeErr, "Error serializing 'code' field")
		}

		if popErr := writeBuffer.PopContext("BVLCResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCResult) isBVLCResult() bool {
	return true
}

func (m *_BVLCResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
