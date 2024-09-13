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

// BACnetShedStateTagged is the corresponding interface of BACnetShedStateTagged
type BACnetShedStateTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetShedState
	// IsBACnetShedStateTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedStateTagged()
}

// _BACnetShedStateTagged is the data-structure of this message
type _BACnetShedStateTagged struct {
	Header BACnetTagHeader
	Value  BACnetShedState

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetShedStateTagged = (*_BACnetShedStateTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedStateTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetShedStateTagged) GetValue() BACnetShedState {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetShedStateTagged factory function for _BACnetShedStateTagged
func NewBACnetShedStateTagged(header BACnetTagHeader, value BACnetShedState, tagNumber uint8, tagClass TagClass) *_BACnetShedStateTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetShedStateTagged must not be nil")
	}
	return &_BACnetShedStateTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetShedStateTagged(structType any) BACnetShedStateTagged {
	if casted, ok := structType.(BACnetShedStateTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedStateTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedStateTagged) GetTypeName() string {
	return "BACnetShedStateTagged"
}

func (m *_BACnetShedStateTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetShedStateTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetShedStateTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetShedStateTagged, error) {
	return BACnetShedStateTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetShedStateTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetShedStateTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetShedStateTagged, error) {
		return BACnetShedStateTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetShedStateTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetShedStateTagged, error) {
	v, err := (&_BACnetShedStateTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetShedStateTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetShedStateTagged BACnetShedStateTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetShedState](ctx, "value", readBuffer, EnsureType[BACnetShedState](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetShedState_SHED_INACTIVE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetShedStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedStateTagged")
	}

	return m, nil
}

func (m *_BACnetShedStateTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedStateTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetShedStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetShedStateTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetShedState](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetShedStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetShedStateTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetShedStateTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetShedStateTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetShedStateTagged) IsBACnetShedStateTagged() {}

func (m *_BACnetShedStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
