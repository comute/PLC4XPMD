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

// PathSegment is the corresponding interface of PathSegment
type PathSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPathSegment returns PathSegment (discriminator field)
	GetPathSegment() uint8
}

// PathSegmentExactly can be used when we want exactly this type and not a type which fulfills PathSegment.
// This is useful for switch cases.
type PathSegmentExactly interface {
	PathSegment
	isPathSegment() bool
}

// _PathSegment is the data-structure of this message
type _PathSegment struct {
	_PathSegmentChildRequirements
}

type _PathSegmentChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPathSegment() uint8
}

type PathSegmentParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child PathSegment, serializeChildFunction func() error) error
	GetTypeName() string
}

type PathSegmentChild interface {
	utils.Serializable
	InitializeParent(parent PathSegment)
	GetParent() *PathSegment

	GetTypeName() string
	PathSegment
}

// NewPathSegment factory function for _PathSegment
func NewPathSegment() *_PathSegment {
	return &_PathSegment{}
}

// Deprecated: use the interface for direct cast
func CastPathSegment(structType any) PathSegment {
	if casted, ok := structType.(PathSegment); ok {
		return casted
	}
	if casted, ok := structType.(*PathSegment); ok {
		return *casted
	}
	return nil
}

func (m *_PathSegment) GetTypeName() string {
	return "PathSegment"
}

func (m *_PathSegment) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (pathSegment)
	lengthInBits += 3

	return lengthInBits
}

func (m *_PathSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PathSegmentParse(ctx context.Context, theBytes []byte) (PathSegment, error) {
	return PathSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PathSegmentParseWithBufferProducer[T PathSegment]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := PathSegmentParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func PathSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PathSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PathSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PathSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pathSegment, err := ReadDiscriminatorField[uint8](ctx, "pathSegment", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pathSegment' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type PathSegmentChildSerializeRequirement interface {
		PathSegment
		InitializeParent(PathSegment)
		GetParent() PathSegment
	}
	var _childTemp any
	var _child PathSegmentChildSerializeRequirement
	var typeSwitchError error
	switch {
	case pathSegment == 0x00: // PortSegment
		_childTemp, typeSwitchError = PortSegmentParseWithBuffer(ctx, readBuffer)
	case pathSegment == 0x01: // LogicalSegment
		_childTemp, typeSwitchError = LogicalSegmentParseWithBuffer(ctx, readBuffer)
	case pathSegment == 0x04: // DataSegment
		_childTemp, typeSwitchError = DataSegmentParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [pathSegment=%v]", pathSegment)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of PathSegment")
	}
	_child = _childTemp.(PathSegmentChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("PathSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PathSegment")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_PathSegment) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child PathSegment, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PathSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PathSegment")
	}

	// Discriminator Field (pathSegment) (Used as input to a switch field)
	pathSegment := uint8(child.GetPathSegment())
	_pathSegmentErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("pathSegment", 3, uint8((pathSegment)))

	if _pathSegmentErr != nil {
		return errors.Wrap(_pathSegmentErr, "Error serializing 'pathSegment' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("PathSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PathSegment")
	}
	return nil
}

func (m *_PathSegment) isPathSegment() bool {
	return true
}

func (m *_PathSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
