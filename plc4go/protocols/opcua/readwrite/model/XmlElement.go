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

// XmlElement is the corresponding interface of XmlElement
type XmlElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLength returns Length (property field)
	GetLength() int32
	// GetValue returns Value (property field)
	GetValue() []string
}

// XmlElementExactly can be used when we want exactly this type and not a type which fulfills XmlElement.
// This is useful for switch cases.
type XmlElementExactly interface {
	XmlElement
	isXmlElement() bool
}

// _XmlElement is the data-structure of this message
type _XmlElement struct {
	Length int32
	Value  []string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_XmlElement) GetLength() int32 {
	return m.Length
}

func (m *_XmlElement) GetValue() []string {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewXmlElement factory function for _XmlElement
func NewXmlElement(length int32, value []string) *_XmlElement {
	return &_XmlElement{Length: length, Value: value}
}

// Deprecated: use the interface for direct cast
func CastXmlElement(structType any) XmlElement {
	if casted, ok := structType.(XmlElement); ok {
		return casted
	}
	if casted, ok := structType.(*XmlElement); ok {
		return *casted
	}
	return nil
}

func (m *_XmlElement) GetTypeName() string {
	return "XmlElement"
}

func (m *_XmlElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (length)
	lengthInBits += 32

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_XmlElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func XmlElementParse(ctx context.Context, theBytes []byte) (XmlElement, error) {
	return XmlElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func XmlElementParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (XmlElement, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (XmlElement, error) {
		return XmlElementParseWithBuffer(ctx, readBuffer)
	}
}

func XmlElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (XmlElement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("XmlElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for XmlElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	length, err := ReadSimpleField(ctx, "length", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}

	value, err := ReadCountArrayField[string](ctx, "value", ReadString(readBuffer, uint32(8)), uint64(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("XmlElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for XmlElement")
	}

	// Create the instance
	return &_XmlElement{
		Length: length,
		Value:  value,
	}, nil
}

func (m *_XmlElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_XmlElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("XmlElement"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for XmlElement")
	}

	// Simple Field (length)
	length := int32(m.GetLength())
	_lengthErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("length", 32, int32((length)))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteString(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("XmlElement"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for XmlElement")
	}
	return nil
}

func (m *_XmlElement) isXmlElement() bool {
	return true
}

func (m *_XmlElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
