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

// Union is the corresponding interface of Union
type Union interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// UnionExactly can be used when we want exactly this type and not a type which fulfills Union.
// This is useful for switch cases.
type UnionExactly interface {
	Union
	isUnion() bool
}

// _Union is the data-structure of this message
type _Union struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Union) GetIdentifier() string {
	return "12758"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Union) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_Union) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewUnion factory function for _Union
func NewUnion() *_Union {
	_result := &_Union{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnion(structType any) Union {
	if casted, ok := structType.(Union); ok {
		return casted
	}
	if casted, ok := structType.(*Union); ok {
		return *casted
	}
	return nil
}

func (m *_Union) GetTypeName() string {
	return "Union"
}

func (m *_Union) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_Union) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnionParse(ctx context.Context, theBytes []byte, identifier string) (Union, error) {
	return UnionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func UnionParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (Union, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Union, error) {
		return UnionParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func UnionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (Union, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Union"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Union")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Union"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Union")
	}

	// Create a partially initialized instance
	_child := &_Union{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_Union) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Union) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Union"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Union")
		}

		if popErr := writeBuffer.PopContext("Union"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Union")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Union) isUnion() bool {
	return true
}

func (m *_Union) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
