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
	ExtensionObjectDefinition
}

// FrameExactly can be used when we want exactly this type and not a type which fulfills Frame.
// This is useful for switch cases.
type FrameExactly interface {
	Frame
	isFrame() bool
}

// _Frame is the data-structure of this message
type _Frame struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Frame) GetIdentifier() string {
	return "18815"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Frame) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_Frame) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewFrame factory function for _Frame
func NewFrame() *_Frame {
	_result := &_Frame{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
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
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_Frame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FrameParse(ctx context.Context, theBytes []byte, identifier string) (Frame, error) {
	return FrameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func FrameParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (Frame, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Frame, error) {
		return FrameParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func FrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (Frame, error) {
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

	// Create a partially initialized instance
	_child := &_Frame{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
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
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Frame) isFrame() bool {
	return true
}

func (m *_Frame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
