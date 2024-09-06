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

// KeyValuePair is the corresponding interface of KeyValuePair
type KeyValuePair interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetKey returns Key (property field)
	GetKey() QualifiedName
	// GetValue returns Value (property field)
	GetValue() Variant
	// IsKeyValuePair is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKeyValuePair()
}

// _KeyValuePair is the data-structure of this message
type _KeyValuePair struct {
	ExtensionObjectDefinitionContract
	Key   QualifiedName
	Value Variant
}

var _ KeyValuePair = (*_KeyValuePair)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_KeyValuePair)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KeyValuePair) GetIdentifier() string {
	return "14535"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KeyValuePair) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KeyValuePair) GetKey() QualifiedName {
	return m.Key
}

func (m *_KeyValuePair) GetValue() Variant {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKeyValuePair factory function for _KeyValuePair
func NewKeyValuePair(key QualifiedName, value Variant) *_KeyValuePair {
	if key == nil {
		panic("key of type QualifiedName for KeyValuePair must not be nil")
	}
	if value == nil {
		panic("value of type Variant for KeyValuePair must not be nil")
	}
	_result := &_KeyValuePair{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Key:                               key,
		Value:                             value,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKeyValuePair(structType any) KeyValuePair {
	if casted, ok := structType.(KeyValuePair); ok {
		return casted
	}
	if casted, ok := structType.(*KeyValuePair); ok {
		return *casted
	}
	return nil
}

func (m *_KeyValuePair) GetTypeName() string {
	return "KeyValuePair"
}

func (m *_KeyValuePair) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (key)
	lengthInBits += m.Key.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_KeyValuePair) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KeyValuePair) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__keyValuePair KeyValuePair, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KeyValuePair"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KeyValuePair")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	key, err := ReadSimpleField[QualifiedName](ctx, "key", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'key' field"))
	}
	m.Key = key

	value, err := ReadSimpleField[Variant](ctx, "value", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("KeyValuePair"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KeyValuePair")
	}

	return m, nil
}

func (m *_KeyValuePair) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KeyValuePair) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KeyValuePair"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KeyValuePair")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "key", m.GetKey(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'key' field")
		}

		if err := WriteSimpleField[Variant](ctx, "value", m.GetValue(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("KeyValuePair"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KeyValuePair")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KeyValuePair) IsKeyValuePair() {}

func (m *_KeyValuePair) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
