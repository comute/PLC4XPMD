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

// BACnetEventParameterChangeOfStateListOfValues is the corresponding interface of BACnetEventParameterChangeOfStateListOfValues
type BACnetEventParameterChangeOfStateListOfValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfStateListOfValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfStateListOfValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfStateListOfValuesExactly interface {
	BACnetEventParameterChangeOfStateListOfValues
	isBACnetEventParameterChangeOfStateListOfValues() bool
}

// _BACnetEventParameterChangeOfStateListOfValues is the data-structure of this message
type _BACnetEventParameterChangeOfStateListOfValues struct {
	OpeningTag   BACnetOpeningTag
	ListOfValues []BACnetPropertyStates
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetListOfValues() []BACnetPropertyStates {
	return m.ListOfValues
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfStateListOfValues factory function for _BACnetEventParameterChangeOfStateListOfValues
func NewBACnetEventParameterChangeOfStateListOfValues(openingTag BACnetOpeningTag, listOfValues []BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfStateListOfValues {
	return &_BACnetEventParameterChangeOfStateListOfValues{OpeningTag: openingTag, ListOfValues: listOfValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfStateListOfValues(structType any) BACnetEventParameterChangeOfStateListOfValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfStateListOfValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfStateListOfValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfStateListOfValues"
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfStateListOfValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfStateListOfValues, error) {
	return BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfStateListOfValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfStateListOfValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfStateListOfValues, error) {
		return BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfStateListOfValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfStateListOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfStateListOfValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfValues, err := ReadTerminatedArrayField[BACnetPropertyStates](ctx, "listOfValues", ReadComplex[BACnetPropertyStates](BACnetPropertyStatesParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfStateListOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfStateListOfValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfStateListOfValues{
		TagNumber:    tagNumber,
		OpeningTag:   openingTag,
		ListOfValues: listOfValues,
		ClosingTag:   closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfStateListOfValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfStateListOfValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfValues", m.GetListOfValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfValues' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfStateListOfValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfStateListOfValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfStateListOfValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfStateListOfValues) isBACnetEventParameterChangeOfStateListOfValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfStateListOfValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
