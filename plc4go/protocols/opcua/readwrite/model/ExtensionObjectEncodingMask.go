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

// ExtensionObjectEncodingMask is the corresponding interface of ExtensionObjectEncodingMask
type ExtensionObjectEncodingMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTypeIdSpecified returns TypeIdSpecified (property field)
	GetTypeIdSpecified() bool
	// GetXmlbody returns Xmlbody (property field)
	GetXmlbody() bool
	// GetBinaryBody returns BinaryBody (property field)
	GetBinaryBody() bool
	// IsExtensionObjectEncodingMask is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsExtensionObjectEncodingMask()
}

// _ExtensionObjectEncodingMask is the data-structure of this message
type _ExtensionObjectEncodingMask struct {
	TypeIdSpecified bool
	Xmlbody         bool
	BinaryBody      bool
	// Reserved Fields
	reservedField0 *int8
}

var _ ExtensionObjectEncodingMask = (*_ExtensionObjectEncodingMask)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtensionObjectEncodingMask) GetTypeIdSpecified() bool {
	return m.TypeIdSpecified
}

func (m *_ExtensionObjectEncodingMask) GetXmlbody() bool {
	return m.Xmlbody
}

func (m *_ExtensionObjectEncodingMask) GetBinaryBody() bool {
	return m.BinaryBody
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExtensionObjectEncodingMask factory function for _ExtensionObjectEncodingMask
func NewExtensionObjectEncodingMask(typeIdSpecified bool, xmlbody bool, binaryBody bool) *_ExtensionObjectEncodingMask {
	return &_ExtensionObjectEncodingMask{TypeIdSpecified: typeIdSpecified, Xmlbody: xmlbody, BinaryBody: binaryBody}
}

// Deprecated: use the interface for direct cast
func CastExtensionObjectEncodingMask(structType any) ExtensionObjectEncodingMask {
	if casted, ok := structType.(ExtensionObjectEncodingMask); ok {
		return casted
	}
	if casted, ok := structType.(*ExtensionObjectEncodingMask); ok {
		return *casted
	}
	return nil
}

func (m *_ExtensionObjectEncodingMask) GetTypeName() string {
	return "ExtensionObjectEncodingMask"
}

func (m *_ExtensionObjectEncodingMask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 5

	// Simple field (typeIdSpecified)
	lengthInBits += 1

	// Simple field (xmlbody)
	lengthInBits += 1

	// Simple field (binaryBody)
	lengthInBits += 1

	return lengthInBits
}

func (m *_ExtensionObjectEncodingMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExtensionObjectEncodingMaskParse(ctx context.Context, theBytes []byte) (ExtensionObjectEncodingMask, error) {
	return ExtensionObjectEncodingMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ExtensionObjectEncodingMaskParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObjectEncodingMask, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObjectEncodingMask, error) {
		return ExtensionObjectEncodingMaskParseWithBuffer(ctx, readBuffer)
	}
}

func ExtensionObjectEncodingMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ExtensionObjectEncodingMask, error) {
	v, err := (&_ExtensionObjectEncodingMask{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ExtensionObjectEncodingMask) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__extensionObjectEncodingMask ExtensionObjectEncodingMask, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtensionObjectEncodingMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtensionObjectEncodingMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadSignedByte(readBuffer, uint8(5)), int8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	typeIdSpecified, err := ReadSimpleField(ctx, "typeIdSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeIdSpecified' field"))
	}
	m.TypeIdSpecified = typeIdSpecified

	xmlbody, err := ReadSimpleField(ctx, "xmlbody", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'xmlbody' field"))
	}
	m.Xmlbody = xmlbody

	binaryBody, err := ReadSimpleField(ctx, "binaryBody", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryBody' field"))
	}
	m.BinaryBody = binaryBody

	if closeErr := readBuffer.CloseContext("ExtensionObjectEncodingMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtensionObjectEncodingMask")
	}

	return m, nil
}

func (m *_ExtensionObjectEncodingMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ExtensionObjectEncodingMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ExtensionObjectEncodingMask"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtensionObjectEncodingMask")
	}

	if err := WriteReservedField[int8](ctx, "reserved", int8(0x00), WriteSignedByte(writeBuffer, 5)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "typeIdSpecified", m.GetTypeIdSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'typeIdSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "xmlbody", m.GetXmlbody(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'xmlbody' field")
	}

	if err := WriteSimpleField[bool](ctx, "binaryBody", m.GetBinaryBody(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'binaryBody' field")
	}

	if popErr := writeBuffer.PopContext("ExtensionObjectEncodingMask"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtensionObjectEncodingMask")
	}
	return nil
}

func (m *_ExtensionObjectEncodingMask) IsExtensionObjectEncodingMask() {}

func (m *_ExtensionObjectEncodingMask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
