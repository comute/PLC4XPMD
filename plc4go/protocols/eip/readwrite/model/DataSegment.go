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

// DataSegment is the corresponding interface of DataSegment
type DataSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() DataSegmentType
}

// DataSegmentExactly can be used when we want exactly this type and not a type which fulfills DataSegment.
// This is useful for switch cases.
type DataSegmentExactly interface {
	DataSegment
	isDataSegment() bool
}

// _DataSegment is the data-structure of this message
type _DataSegment struct {
	*_PathSegment
	SegmentType DataSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSegment) GetPathSegment() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSegment) InitializeParent(parent PathSegment) {}

func (m *_DataSegment) GetParent() PathSegment {
	return m._PathSegment
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSegment) GetSegmentType() DataSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataSegment factory function for _DataSegment
func NewDataSegment(segmentType DataSegmentType) *_DataSegment {
	_result := &_DataSegment{
		SegmentType:  segmentType,
		_PathSegment: NewPathSegment(),
	}
	_result._PathSegment._PathSegmentChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataSegment(structType any) DataSegment {
	if casted, ok := structType.(DataSegment); ok {
		return casted
	}
	if casted, ok := structType.(*DataSegment); ok {
		return *casted
	}
	return nil
}

func (m *_DataSegment) GetTypeName() string {
	return "DataSegment"
}

func (m *_DataSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSegmentParse(ctx context.Context, theBytes []byte) (DataSegment, error) {
	return DataSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSegmentParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (DataSegment, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DataSegment, error) {
		return DataSegmentParseWithBuffer(ctx, readBuffer)
	}
}

func DataSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentType, err := ReadSimpleField[DataSegmentType](ctx, "segmentType", ReadComplex[DataSegmentType](DataSegmentTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentType' field"))
	}

	if closeErr := readBuffer.CloseContext("DataSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSegment")
	}

	// Create a partially initialized instance
	_child := &_DataSegment{
		_PathSegment: &_PathSegment{},
		SegmentType:  segmentType,
	}
	_child._PathSegment._PathSegmentChildRequirements = _child
	return _child, nil
}

func (m *_DataSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSegment")
		}

		// Simple Field (segmentType)
		if pushErr := writeBuffer.PushContext("segmentType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for segmentType")
		}
		_segmentTypeErr := writeBuffer.WriteSerializable(ctx, m.GetSegmentType())
		if popErr := writeBuffer.PopContext("segmentType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for segmentType")
		}
		if _segmentTypeErr != nil {
			return errors.Wrap(_segmentTypeErr, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("DataSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSegment")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSegment) isDataSegment() bool {
	return true
}

func (m *_DataSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
