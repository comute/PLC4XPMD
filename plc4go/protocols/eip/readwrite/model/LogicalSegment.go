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

// LogicalSegment is the corresponding interface of LogicalSegment
type LogicalSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() LogicalSegmentType
	// IsLogicalSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLogicalSegment()
	// CreateBuilder creates a LogicalSegmentBuilder
	CreateLogicalSegmentBuilder() LogicalSegmentBuilder
}

// _LogicalSegment is the data-structure of this message
type _LogicalSegment struct {
	PathSegmentContract
	SegmentType LogicalSegmentType
}

var _ LogicalSegment = (*_LogicalSegment)(nil)
var _ PathSegmentRequirements = (*_LogicalSegment)(nil)

// NewLogicalSegment factory function for _LogicalSegment
func NewLogicalSegment(segmentType LogicalSegmentType) *_LogicalSegment {
	if segmentType == nil {
		panic("segmentType of type LogicalSegmentType for LogicalSegment must not be nil")
	}
	_result := &_LogicalSegment{
		PathSegmentContract: NewPathSegment(),
		SegmentType:         segmentType,
	}
	_result.PathSegmentContract.(*_PathSegment)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LogicalSegmentBuilder is a builder for LogicalSegment
type LogicalSegmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentType LogicalSegmentType) LogicalSegmentBuilder
	// WithSegmentType adds SegmentType (property field)
	WithSegmentType(LogicalSegmentType) LogicalSegmentBuilder
	// WithSegmentTypeBuilder adds SegmentType (property field) which is build by the builder
	WithSegmentTypeBuilder(func(LogicalSegmentTypeBuilder) LogicalSegmentTypeBuilder) LogicalSegmentBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() PathSegmentBuilder
	// Build builds the LogicalSegment or returns an error if something is wrong
	Build() (LogicalSegment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LogicalSegment
}

// NewLogicalSegmentBuilder() creates a LogicalSegmentBuilder
func NewLogicalSegmentBuilder() LogicalSegmentBuilder {
	return &_LogicalSegmentBuilder{_LogicalSegment: new(_LogicalSegment)}
}

type _LogicalSegmentBuilder struct {
	*_LogicalSegment

	parentBuilder *_PathSegmentBuilder

	err *utils.MultiError
}

var _ (LogicalSegmentBuilder) = (*_LogicalSegmentBuilder)(nil)

func (b *_LogicalSegmentBuilder) setParent(contract PathSegmentContract) {
	b.PathSegmentContract = contract
	contract.(*_PathSegment)._SubType = b._LogicalSegment
}

func (b *_LogicalSegmentBuilder) WithMandatoryFields(segmentType LogicalSegmentType) LogicalSegmentBuilder {
	return b.WithSegmentType(segmentType)
}

func (b *_LogicalSegmentBuilder) WithSegmentType(segmentType LogicalSegmentType) LogicalSegmentBuilder {
	b.SegmentType = segmentType
	return b
}

func (b *_LogicalSegmentBuilder) WithSegmentTypeBuilder(builderSupplier func(LogicalSegmentTypeBuilder) LogicalSegmentTypeBuilder) LogicalSegmentBuilder {
	builder := builderSupplier(b.SegmentType.CreateLogicalSegmentTypeBuilder())
	var err error
	b.SegmentType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LogicalSegmentTypeBuilder failed"))
	}
	return b
}

func (b *_LogicalSegmentBuilder) Build() (LogicalSegment, error) {
	if b.SegmentType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'segmentType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LogicalSegment.deepCopy(), nil
}

func (b *_LogicalSegmentBuilder) MustBuild() LogicalSegment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LogicalSegmentBuilder) Done() PathSegmentBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewPathSegmentBuilder().(*_PathSegmentBuilder)
	}
	return b.parentBuilder
}

func (b *_LogicalSegmentBuilder) buildForPathSegment() (PathSegment, error) {
	return b.Build()
}

func (b *_LogicalSegmentBuilder) DeepCopy() any {
	_copy := b.CreateLogicalSegmentBuilder().(*_LogicalSegmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLogicalSegmentBuilder creates a LogicalSegmentBuilder
func (b *_LogicalSegment) CreateLogicalSegmentBuilder() LogicalSegmentBuilder {
	if b == nil {
		return NewLogicalSegmentBuilder()
	}
	return &_LogicalSegmentBuilder{_LogicalSegment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LogicalSegment) GetPathSegment() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LogicalSegment) GetParent() PathSegmentContract {
	return m.PathSegmentContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LogicalSegment) GetSegmentType() LogicalSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLogicalSegment(structType any) LogicalSegment {
	if casted, ok := structType.(LogicalSegment); ok {
		return casted
	}
	if casted, ok := structType.(*LogicalSegment); ok {
		return *casted
	}
	return nil
}

func (m *_LogicalSegment) GetTypeName() string {
	return "LogicalSegment"
}

func (m *_LogicalSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PathSegmentContract.(*_PathSegment).getLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LogicalSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LogicalSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PathSegment) (__logicalSegment LogicalSegment, err error) {
	m.PathSegmentContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LogicalSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LogicalSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentType, err := ReadSimpleField[LogicalSegmentType](ctx, "segmentType", ReadComplex[LogicalSegmentType](LogicalSegmentTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentType' field"))
	}
	m.SegmentType = segmentType

	if closeErr := readBuffer.CloseContext("LogicalSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LogicalSegment")
	}

	return m, nil
}

func (m *_LogicalSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LogicalSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LogicalSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LogicalSegment")
		}

		if err := WriteSimpleField[LogicalSegmentType](ctx, "segmentType", m.GetSegmentType(), WriteComplex[LogicalSegmentType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("LogicalSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LogicalSegment")
		}
		return nil
	}
	return m.PathSegmentContract.(*_PathSegment).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LogicalSegment) IsLogicalSegment() {}

func (m *_LogicalSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LogicalSegment) deepCopy() *_LogicalSegment {
	if m == nil {
		return nil
	}
	_LogicalSegmentCopy := &_LogicalSegment{
		m.PathSegmentContract.(*_PathSegment).deepCopy(),
		utils.DeepCopy[LogicalSegmentType](m.SegmentType),
	}
	m.PathSegmentContract.(*_PathSegment)._SubType = m
	return _LogicalSegmentCopy
}

func (m *_LogicalSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
