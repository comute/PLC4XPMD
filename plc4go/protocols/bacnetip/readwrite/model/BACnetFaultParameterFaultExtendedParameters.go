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

// BACnetFaultParameterFaultExtendedParameters is the corresponding interface of BACnetFaultParameterFaultExtendedParameters
type BACnetFaultParameterFaultExtendedParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetParameters returns Parameters (property field)
	GetParameters() []BACnetFaultParameterFaultExtendedParametersEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultExtendedParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParameters.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersExactly interface {
	BACnetFaultParameterFaultExtendedParameters
	isBACnetFaultParameterFaultExtendedParameters() bool
}

// _BACnetFaultParameterFaultExtendedParameters is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParameters struct {
	OpeningTag BACnetOpeningTag
	Parameters []BACnetFaultParameterFaultExtendedParametersEntry
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetParameters() []BACnetFaultParameterFaultExtendedParametersEntry {
	return m.Parameters
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParameters factory function for _BACnetFaultParameterFaultExtendedParameters
func NewBACnetFaultParameterFaultExtendedParameters(openingTag BACnetOpeningTag, parameters []BACnetFaultParameterFaultExtendedParametersEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultExtendedParameters {
	return &_BACnetFaultParameterFaultExtendedParameters{OpeningTag: openingTag, Parameters: parameters, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParameters(structType any) BACnetFaultParameterFaultExtendedParameters {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParameters"
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultExtendedParameters, error) {
	return BACnetFaultParameterFaultExtendedParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultExtendedParametersParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParameters, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParameters, error) {
		return BACnetFaultParameterFaultExtendedParametersParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetFaultParameterFaultExtendedParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultExtendedParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	parameters, err := ReadTerminatedArrayField[BACnetFaultParameterFaultExtendedParametersEntry](ctx, "parameters", ReadComplex[BACnetFaultParameterFaultExtendedParametersEntry](BACnetFaultParameterFaultExtendedParametersEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParameters")
	}

	// Create the instance
	return &_BACnetFaultParameterFaultExtendedParameters{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Parameters: parameters,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetFaultParameterFaultExtendedParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParameters")
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

	if err := WriteComplexTypeArrayField(ctx, "parameters", m.GetParameters(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'parameters' field")
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

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultExtendedParameters) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultExtendedParameters) isBACnetFaultParameterFaultExtendedParameters() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
