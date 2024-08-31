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

// BACnetAccessPassbackModeTagged is the corresponding interface of BACnetAccessPassbackModeTagged
type BACnetAccessPassbackModeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessPassbackMode
}

// BACnetAccessPassbackModeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessPassbackModeTagged.
// This is useful for switch cases.
type BACnetAccessPassbackModeTaggedExactly interface {
	BACnetAccessPassbackModeTagged
	isBACnetAccessPassbackModeTagged() bool
}

// _BACnetAccessPassbackModeTagged is the data-structure of this message
type _BACnetAccessPassbackModeTagged struct {
	Header BACnetTagHeader
	Value  BACnetAccessPassbackMode

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessPassbackModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessPassbackModeTagged) GetValue() BACnetAccessPassbackMode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessPassbackModeTagged factory function for _BACnetAccessPassbackModeTagged
func NewBACnetAccessPassbackModeTagged(header BACnetTagHeader, value BACnetAccessPassbackMode, tagNumber uint8, tagClass TagClass) *_BACnetAccessPassbackModeTagged {
	return &_BACnetAccessPassbackModeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessPassbackModeTagged(structType any) BACnetAccessPassbackModeTagged {
	if casted, ok := structType.(BACnetAccessPassbackModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessPassbackModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessPassbackModeTagged) GetTypeName() string {
	return "BACnetAccessPassbackModeTagged"
}

func (m *_BACnetAccessPassbackModeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAccessPassbackModeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessPassbackModeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessPassbackModeTagged, error) {
	return BACnetAccessPassbackModeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessPassbackModeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessPassbackModeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessPassbackModeTagged, error) {
		return BACnetAccessPassbackModeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessPassbackModeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessPassbackModeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessPassbackModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessPassbackModeTagged")
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

	value, err := ReadManualField[BACnetAccessPassbackMode](ctx, "value", readBuffer, EnsureType[BACnetAccessPassbackMode](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAccessPassbackMode_PASSBACK_OFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessPassbackModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessPassbackModeTagged")
	}

	// Create the instance
	return &_BACnetAccessPassbackModeTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetAccessPassbackModeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessPassbackModeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessPassbackModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessPassbackModeTagged")
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

	if popErr := writeBuffer.PopContext("BACnetAccessPassbackModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessPassbackModeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessPassbackModeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessPassbackModeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessPassbackModeTagged) isBACnetAccessPassbackModeTagged() bool {
	return true
}

func (m *_BACnetAccessPassbackModeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
