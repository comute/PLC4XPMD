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

// BACnetNameValue is the corresponding interface of BACnetNameValue
type BACnetNameValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetName returns Name (property field)
	GetName() BACnetContextTagCharacterString
	// GetValue returns Value (property field)
	GetValue() BACnetConstructedData
}

// BACnetNameValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNameValue.
// This is useful for switch cases.
type BACnetNameValueExactly interface {
	BACnetNameValue
	isBACnetNameValue() bool
}

// _BACnetNameValue is the data-structure of this message
type _BACnetNameValue struct {
	Name  BACnetContextTagCharacterString
	Value BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNameValue) GetName() BACnetContextTagCharacterString {
	return m.Name
}

func (m *_BACnetNameValue) GetValue() BACnetConstructedData {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNameValue factory function for _BACnetNameValue
func NewBACnetNameValue(name BACnetContextTagCharacterString, value BACnetConstructedData) *_BACnetNameValue {
	return &_BACnetNameValue{Name: name, Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetNameValue(structType any) BACnetNameValue {
	if casted, ok := structType.(BACnetNameValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNameValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNameValue) GetTypeName() string {
	return "BACnetNameValue"
}

func (m *_BACnetNameValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Optional Field (value)
	if m.Value != nil {
		lengthInBits += m.Value.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetNameValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNameValueParse(ctx context.Context, theBytes []byte) (BACnetNameValue, error) {
	return BACnetNameValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetNameValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNameValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNameValue, error) {
		return BACnetNameValueParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetNameValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNameValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNameValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNameValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "name", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}

	_value, err := ReadOptionalField[BACnetConstructedData](ctx, "value", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(1)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	var value BACnetConstructedData
	if _value != nil {
		value = *_value
	}

	if closeErr := readBuffer.CloseContext("BACnetNameValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNameValue")
	}

	// Create the instance
	return &_BACnetNameValue{
		Name:  name,
		Value: value,
	}, nil
}

func (m *_BACnetNameValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNameValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNameValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNameValue")
	}

	// Simple Field (name)
	if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for name")
	}
	_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
	if popErr := writeBuffer.PopContext("name"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for name")
	}
	if _nameErr != nil {
		return errors.Wrap(_nameErr, "Error serializing 'name' field")
	}

	// Optional Field (value) (Can be skipped, if the value is null)
	var value BACnetConstructedData = nil
	if m.GetValue() != nil {
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		value = m.GetValue()
		_valueErr := writeBuffer.WriteSerializable(ctx, value)
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetNameValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNameValue")
	}
	return nil
}

func (m *_BACnetNameValue) isBACnetNameValue() bool {
	return true
}

func (m *_BACnetNameValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
