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

// BACnetNotificationParameters is the corresponding interface of BACnetNotificationParameters
type BACnetNotificationParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetNotificationParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParameters.
// This is useful for switch cases.
type BACnetNotificationParametersExactly interface {
	BACnetNotificationParameters
	isBACnetNotificationParameters() bool
}

// _BACnetNotificationParameters is the data-structure of this message
type _BACnetNotificationParameters struct {
	_BACnetNotificationParametersChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

type _BACnetNotificationParametersChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetNotificationParametersParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParameters, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetNotificationParametersChild interface {
	utils.Serializable
	InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetNotificationParameters

	GetTypeName() string
	BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParameters) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetNotificationParameters) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParameters factory function for _BACnetNotificationParameters
func NewBACnetNotificationParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParameters {
	return &_BACnetNotificationParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParameters(structType any) BACnetNotificationParameters {
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParameters) GetTypeName() string {
	return "BACnetNotificationParameters"
}

func (m *_BACnetNotificationParameters) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParameters, error) {
	return BACnetNotificationParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersParseWithBufferProducer[T BACnetNotificationParameters](tagNumber uint8, objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := BACnetNotificationParametersParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func BACnetNotificationParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetNotificationParametersChildSerializeRequirement interface {
		BACnetNotificationParameters
		InitializeParent(BACnetNotificationParameters, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetNotificationParameters
	}
	var _childTemp any
	var _child BACnetNotificationParametersChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfBitString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfBitStringParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfState
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfStateParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(2): // BACnetNotificationParametersChangeOfValue
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfValueParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(3): // BACnetNotificationParametersCommandFailure
		_childTemp, typeSwitchError = BACnetNotificationParametersCommandFailureParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(4): // BACnetNotificationParametersFloatingLimit
		_childTemp, typeSwitchError = BACnetNotificationParametersFloatingLimitParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(5): // BACnetNotificationParametersOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersOutOfRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(6): // BACnetNotificationParametersComplexEventType
		_childTemp, typeSwitchError = BACnetNotificationParametersComplexEventTypeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(8): // BACnetNotificationParametersChangeOfLifeSafety
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfLifeSafetyParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(9): // BACnetNotificationParametersExtended
		_childTemp, typeSwitchError = BACnetNotificationParametersExtendedParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(10): // BACnetNotificationParametersBufferReady
		_childTemp, typeSwitchError = BACnetNotificationParametersBufferReadyParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(11): // BACnetNotificationParametersUnsignedRange
		_childTemp, typeSwitchError = BACnetNotificationParametersUnsignedRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(13): // BACnetNotificationParametersAccessEvent
		_childTemp, typeSwitchError = BACnetNotificationParametersAccessEventParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(14): // BACnetNotificationParametersDoubleOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersDoubleOutOfRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(15): // BACnetNotificationParametersSignedOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersSignedOutOfRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(16): // BACnetNotificationParametersUnsignedOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersUnsignedOutOfRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(17): // BACnetNotificationParametersChangeOfCharacterString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfCharacterStringParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(18): // BACnetNotificationParametersChangeOfStatusFlags
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfStatusFlagsParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(19): // BACnetNotificationParametersChangeOfReliability
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfReliabilityParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(21): // BACnetNotificationParametersChangeOfDiscreteValue
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	case peekedTagNumber == uint8(22): // BACnetNotificationParametersChangeOfTimer
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfTimerParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetNotificationParameters")
	}
	_child = _childTemp.(BACnetNotificationParametersChildSerializeRequirement)

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParameters")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetNotificationParameters) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParameters, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotificationParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParameters")
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
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
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

	if popErr := writeBuffer.PopContext("BACnetNotificationParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParameters) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNotificationParameters) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetNotificationParameters) isBACnetNotificationParameters() bool {
	return true
}

func (m *_BACnetNotificationParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
