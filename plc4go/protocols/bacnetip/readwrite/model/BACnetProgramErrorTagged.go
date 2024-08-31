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

// BACnetProgramErrorTagged is the corresponding interface of BACnetProgramErrorTagged
type BACnetProgramErrorTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramError
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetProgramErrorTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetProgramErrorTagged.
// This is useful for switch cases.
type BACnetProgramErrorTaggedExactly interface {
	BACnetProgramErrorTagged
	isBACnetProgramErrorTagged() bool
}

// _BACnetProgramErrorTagged is the data-structure of this message
type _BACnetProgramErrorTagged struct {
	Header           BACnetTagHeader
	Value            BACnetProgramError
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProgramErrorTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProgramErrorTagged) GetValue() BACnetProgramError {
	return m.Value
}

func (m *_BACnetProgramErrorTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetProgramErrorTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProgramErrorTagged factory function for _BACnetProgramErrorTagged
func NewBACnetProgramErrorTagged(header BACnetTagHeader, value BACnetProgramError, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetProgramErrorTagged {
	return &_BACnetProgramErrorTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetProgramErrorTagged(structType any) BACnetProgramErrorTagged {
	if casted, ok := structType.(BACnetProgramErrorTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProgramErrorTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProgramErrorTagged) GetTypeName() string {
	return "BACnetProgramErrorTagged"
}

func (m *_BACnetProgramErrorTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetProgramErrorTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProgramErrorTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProgramErrorTagged, error) {
	return BACnetProgramErrorTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProgramErrorTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramErrorTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramErrorTagged, error) {
		return BACnetProgramErrorTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetProgramErrorTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProgramErrorTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramErrorTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProgramErrorTagged")
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

	value, err := ReadManualField[BACnetProgramError](ctx, "value", readBuffer, EnsureType[BACnetProgramError](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetProgramError_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetProgramErrorTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProgramErrorTagged")
	}

	// Create the instance
	return &_BACnetProgramErrorTagged{
		TagNumber:        tagNumber,
		TagClass:         tagClass,
		Header:           header,
		Value:            value,
		ProprietaryValue: proprietaryValue,
	}, nil
}

func (m *_BACnetProgramErrorTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProgramErrorTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProgramErrorTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProgramErrorTagged")
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
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProgramErrorTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProgramErrorTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProgramErrorTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProgramErrorTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProgramErrorTagged) isBACnetProgramErrorTagged() bool {
	return true
}

func (m *_BACnetProgramErrorTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
