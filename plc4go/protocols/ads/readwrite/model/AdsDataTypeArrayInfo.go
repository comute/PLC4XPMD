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

// AdsDataTypeArrayInfo is the corresponding interface of AdsDataTypeArrayInfo
type AdsDataTypeArrayInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLowerBound returns LowerBound (property field)
	GetLowerBound() uint32
	// GetNumElements returns NumElements (property field)
	GetNumElements() uint32
	// GetUpperBound returns UpperBound (virtual field)
	GetUpperBound() uint32
	// IsAdsDataTypeArrayInfo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDataTypeArrayInfo()
}

// _AdsDataTypeArrayInfo is the data-structure of this message
type _AdsDataTypeArrayInfo struct {
	LowerBound  uint32
	NumElements uint32
}

var _ AdsDataTypeArrayInfo = (*_AdsDataTypeArrayInfo)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDataTypeArrayInfo) GetLowerBound() uint32 {
	return m.LowerBound
}

func (m *_AdsDataTypeArrayInfo) GetNumElements() uint32 {
	return m.NumElements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_AdsDataTypeArrayInfo) GetUpperBound() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(uint32(m.GetLowerBound()) + uint32(m.GetNumElements()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDataTypeArrayInfo factory function for _AdsDataTypeArrayInfo
func NewAdsDataTypeArrayInfo(lowerBound uint32, numElements uint32) *_AdsDataTypeArrayInfo {
	return &_AdsDataTypeArrayInfo{LowerBound: lowerBound, NumElements: numElements}
}

// Deprecated: use the interface for direct cast
func CastAdsDataTypeArrayInfo(structType any) AdsDataTypeArrayInfo {
	if casted, ok := structType.(AdsDataTypeArrayInfo); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDataTypeArrayInfo); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDataTypeArrayInfo) GetTypeName() string {
	return "AdsDataTypeArrayInfo"
}

func (m *_AdsDataTypeArrayInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lowerBound)
	lengthInBits += 32

	// Simple field (numElements)
	lengthInBits += 32

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_AdsDataTypeArrayInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDataTypeArrayInfoParse(ctx context.Context, theBytes []byte) (AdsDataTypeArrayInfo, error) {
	return AdsDataTypeArrayInfoParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsDataTypeArrayInfoParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeArrayInfo, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeArrayInfo, error) {
		return AdsDataTypeArrayInfoParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDataTypeArrayInfoParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDataTypeArrayInfo, error) {
	v, err := (&_AdsDataTypeArrayInfo{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_AdsDataTypeArrayInfo) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__adsDataTypeArrayInfo AdsDataTypeArrayInfo, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDataTypeArrayInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDataTypeArrayInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowerBound, err := ReadSimpleField(ctx, "lowerBound", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowerBound' field"))
	}
	m.LowerBound = lowerBound

	numElements, err := ReadSimpleField(ctx, "numElements", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numElements' field"))
	}
	m.NumElements = numElements

	upperBound, err := ReadVirtualField[uint32](ctx, "upperBound", (*uint32)(nil), uint32(lowerBound)+uint32(numElements), codegen.WithByteOrder(binary.LittleEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'upperBound' field"))
	}
	_ = upperBound

	if closeErr := readBuffer.CloseContext("AdsDataTypeArrayInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDataTypeArrayInfo")
	}

	return m, nil
}

func (m *_AdsDataTypeArrayInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDataTypeArrayInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AdsDataTypeArrayInfo"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsDataTypeArrayInfo")
	}

	if err := WriteSimpleField[uint32](ctx, "lowerBound", m.GetLowerBound(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'lowerBound' field")
	}

	if err := WriteSimpleField[uint32](ctx, "numElements", m.GetNumElements(), WriteUnsignedInt(writeBuffer, 32), codegen.WithByteOrder(binary.LittleEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'numElements' field")
	}
	// Virtual field
	upperBound := m.GetUpperBound()
	_ = upperBound
	if _upperBoundErr := writeBuffer.WriteVirtual(ctx, "upperBound", m.GetUpperBound()); _upperBoundErr != nil {
		return errors.Wrap(_upperBoundErr, "Error serializing 'upperBound' field")
	}

	if popErr := writeBuffer.PopContext("AdsDataTypeArrayInfo"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsDataTypeArrayInfo")
	}
	return nil
}

func (m *_AdsDataTypeArrayInfo) IsAdsDataTypeArrayInfo() {}

func (m *_AdsDataTypeArrayInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
