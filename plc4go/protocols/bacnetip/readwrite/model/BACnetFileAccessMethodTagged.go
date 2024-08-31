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

// BACnetFileAccessMethodTagged is the corresponding interface of BACnetFileAccessMethodTagged
type BACnetFileAccessMethodTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetFileAccessMethod
}

// BACnetFileAccessMethodTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetFileAccessMethodTagged.
// This is useful for switch cases.
type BACnetFileAccessMethodTaggedExactly interface {
	BACnetFileAccessMethodTagged
	isBACnetFileAccessMethodTagged() bool
}

// _BACnetFileAccessMethodTagged is the data-structure of this message
type _BACnetFileAccessMethodTagged struct {
	Header BACnetTagHeader
	Value  BACnetFileAccessMethod

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFileAccessMethodTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetFileAccessMethodTagged) GetValue() BACnetFileAccessMethod {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFileAccessMethodTagged factory function for _BACnetFileAccessMethodTagged
func NewBACnetFileAccessMethodTagged(header BACnetTagHeader, value BACnetFileAccessMethod, tagNumber uint8, tagClass TagClass) *_BACnetFileAccessMethodTagged {
	return &_BACnetFileAccessMethodTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetFileAccessMethodTagged(structType any) BACnetFileAccessMethodTagged {
	if casted, ok := structType.(BACnetFileAccessMethodTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFileAccessMethodTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFileAccessMethodTagged) GetTypeName() string {
	return "BACnetFileAccessMethodTagged"
}

func (m *_BACnetFileAccessMethodTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetFileAccessMethodTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFileAccessMethodTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetFileAccessMethodTagged, error) {
	return BACnetFileAccessMethodTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetFileAccessMethodTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFileAccessMethodTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFileAccessMethodTagged, error) {
		return BACnetFileAccessMethodTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetFileAccessMethodTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetFileAccessMethodTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFileAccessMethodTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFileAccessMethodTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetFileAccessMethod](ctx, "value", readBuffer, EnsureType[BACnetFileAccessMethod](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetFileAccessMethod_RECORD_ACCESS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetFileAccessMethodTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFileAccessMethodTagged")
	}

	// Create the instance
	return &_BACnetFileAccessMethodTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetFileAccessMethodTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFileAccessMethodTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFileAccessMethodTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFileAccessMethodTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFileAccessMethodTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFileAccessMethodTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFileAccessMethodTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetFileAccessMethodTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetFileAccessMethodTagged) isBACnetFileAccessMethodTagged() bool {
	return true
}

func (m *_BACnetFileAccessMethodTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
