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

// BACnetConstructedDataCredentialDisable is the corresponding interface of BACnetConstructedDataCredentialDisable
type BACnetConstructedDataCredentialDisable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCredentialDisable returns CredentialDisable (property field)
	GetCredentialDisable() BACnetAccessCredentialDisableTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessCredentialDisableTagged
}

// BACnetConstructedDataCredentialDisableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialDisable.
// This is useful for switch cases.
type BACnetConstructedDataCredentialDisableExactly interface {
	BACnetConstructedDataCredentialDisable
	isBACnetConstructedDataCredentialDisable() bool
}

// _BACnetConstructedDataCredentialDisable is the data-structure of this message
type _BACnetConstructedDataCredentialDisable struct {
	*_BACnetConstructedData
	CredentialDisable BACnetAccessCredentialDisableTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialDisable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIAL_DISABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialDisable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialDisable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetCredentialDisable() BACnetAccessCredentialDisableTagged {
	return m.CredentialDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetActualValue() BACnetAccessCredentialDisableTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessCredentialDisableTagged(m.GetCredentialDisable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialDisable factory function for _BACnetConstructedDataCredentialDisable
func NewBACnetConstructedDataCredentialDisable(credentialDisable BACnetAccessCredentialDisableTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialDisable {
	_result := &_BACnetConstructedDataCredentialDisable{
		CredentialDisable:      credentialDisable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialDisable(structType any) BACnetConstructedDataCredentialDisable {
	if casted, ok := structType.(BACnetConstructedDataCredentialDisable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDisable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialDisable) GetTypeName() string {
	return "BACnetConstructedDataCredentialDisable"
}

func (m *_BACnetConstructedDataCredentialDisable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (credentialDisable)
	lengthInBits += m.CredentialDisable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialDisable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCredentialDisableParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialDisable, error) {
	return BACnetConstructedDataCredentialDisableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCredentialDisableParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataCredentialDisable, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataCredentialDisable, error) {
		return BACnetConstructedDataCredentialDisableParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataCredentialDisableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialDisable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	credentialDisable, err := ReadSimpleField[BACnetAccessCredentialDisableTagged](ctx, "credentialDisable", ReadComplex[BACnetAccessCredentialDisableTagged](BACnetAccessCredentialDisableTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'credentialDisable' field"))
	}

	actualValue, err := ReadVirtualField[BACnetAccessCredentialDisableTagged](ctx, "actualValue", (*BACnetAccessCredentialDisableTagged)(nil), credentialDisable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialDisable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialDisable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CredentialDisable: credentialDisable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialDisable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialDisable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialDisable")
		}

		// Simple Field (credentialDisable)
		if pushErr := writeBuffer.PushContext("credentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for credentialDisable")
		}
		_credentialDisableErr := writeBuffer.WriteSerializable(ctx, m.GetCredentialDisable())
		if popErr := writeBuffer.PopContext("credentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for credentialDisable")
		}
		if _credentialDisableErr != nil {
			return errors.Wrap(_credentialDisableErr, "Error serializing 'credentialDisable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialDisable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialDisable) isBACnetConstructedDataCredentialDisable() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
