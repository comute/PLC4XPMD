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

// BACnetConstructedDataAuthenticationPolicyNames is the corresponding interface of BACnetConstructedDataAuthenticationPolicyNames
type BACnetConstructedDataAuthenticationPolicyNames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationPolicyNames returns AuthenticationPolicyNames (property field)
	GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataAuthenticationPolicyNamesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthenticationPolicyNames.
// This is useful for switch cases.
type BACnetConstructedDataAuthenticationPolicyNamesExactly interface {
	BACnetConstructedDataAuthenticationPolicyNames
	isBACnetConstructedDataAuthenticationPolicyNames() bool
}

// _BACnetConstructedDataAuthenticationPolicyNames is the data-structure of this message
type _BACnetConstructedDataAuthenticationPolicyNames struct {
	*_BACnetConstructedData
	NumberOfDataElements      BACnetApplicationTagUnsignedInteger
	AuthenticationPolicyNames []BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_NAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString {
	return m.AuthenticationPolicyNames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthenticationPolicyNames factory function for _BACnetConstructedDataAuthenticationPolicyNames
func NewBACnetConstructedDataAuthenticationPolicyNames(numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationPolicyNames []BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationPolicyNames {
	_result := &_BACnetConstructedDataAuthenticationPolicyNames{
		NumberOfDataElements:      numberOfDataElements,
		AuthenticationPolicyNames: authenticationPolicyNames,
		_BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthenticationPolicyNames(structType any) BACnetConstructedDataAuthenticationPolicyNames {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyNames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyNames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyNames"
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AuthenticationPolicyNames) > 0 {
		for _, element := range m.AuthenticationPolicyNames {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAuthenticationPolicyNamesParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationPolicyNames, error) {
	return BACnetConstructedDataAuthenticationPolicyNamesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAuthenticationPolicyNamesParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAuthenticationPolicyNames, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataAuthenticationPolicyNames, error) {
		return BACnetConstructedDataAuthenticationPolicyNamesParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataAuthenticationPolicyNamesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationPolicyNames, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyNames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationPolicyNames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
	}

	authenticationPolicyNames, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "authenticationPolicyNames", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationPolicyNames' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyNames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationPolicyNames")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthenticationPolicyNames{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements:      numberOfDataElements,
		AuthenticationPolicyNames: authenticationPolicyNames,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyNames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationPolicyNames")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		if err := WriteComplexTypeArrayField(ctx, "authenticationPolicyNames", m.GetAuthenticationPolicyNames(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationPolicyNames' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyNames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationPolicyNames")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) isBACnetConstructedDataAuthenticationPolicyNames() bool {
	return true
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
