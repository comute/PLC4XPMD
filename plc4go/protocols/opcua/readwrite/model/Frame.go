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

// Frame is the corresponding interface of Frame
type Frame interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsFrame is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFrame()
	// CreateBuilder creates a FrameBuilder
	CreateFrameBuilder() FrameBuilder
}

// _Frame is the data-structure of this message
type _Frame struct {
	ExtensionObjectDefinitionContract
}

var _ Frame = (*_Frame)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Frame)(nil)

// NewFrame factory function for _Frame
func NewFrame() *_Frame {
	_result := &_Frame{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FrameBuilder is a builder for Frame
type FrameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() FrameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the Frame or returns an error if something is wrong
	Build() (Frame, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Frame
}

// NewFrameBuilder() creates a FrameBuilder
func NewFrameBuilder() FrameBuilder {
	return &_FrameBuilder{_Frame: new(_Frame)}
}

type _FrameBuilder struct {
	*_Frame

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FrameBuilder) = (*_FrameBuilder)(nil)

func (b *_FrameBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._Frame
}

func (b *_FrameBuilder) WithMandatoryFields() FrameBuilder {
	return b
}

func (b *_FrameBuilder) Build() (Frame, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Frame.deepCopy(), nil
}

func (b *_FrameBuilder) MustBuild() Frame {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FrameBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_FrameBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FrameBuilder) DeepCopy() any {
	_copy := b.CreateFrameBuilder().(*_FrameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFrameBuilder creates a FrameBuilder
func (b *_Frame) CreateFrameBuilder() FrameBuilder {
	if b == nil {
		return NewFrameBuilder()
	}
	return &_FrameBuilder{_Frame: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Frame) GetExtensionId() int32 {
	return int32(18815)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Frame) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastFrame(structType any) Frame {
	if casted, ok := structType.(Frame); ok {
		return casted
	}
	if casted, ok := structType.(*Frame); ok {
		return *casted
	}
	return nil
}

func (m *_Frame) GetTypeName() string {
	return "Frame"
}

func (m *_Frame) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_Frame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Frame) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__frame Frame, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Frame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Frame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Frame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Frame")
	}

	return m, nil
}

func (m *_Frame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Frame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Frame"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Frame")
		}

		if popErr := writeBuffer.PopContext("Frame"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Frame")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Frame) IsFrame() {}

func (m *_Frame) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Frame) deepCopy() *_Frame {
	if m == nil {
		return nil
	}
	_FrameCopy := &_Frame{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FrameCopy
}

func (m *_Frame) String() string {
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
