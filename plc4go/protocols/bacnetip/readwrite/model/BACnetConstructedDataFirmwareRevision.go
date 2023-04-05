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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataFirmwareRevision is the corresponding interface of BACnetConstructedDataFirmwareRevision
type BACnetConstructedDataFirmwareRevision interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFirmwareRevision returns FirmwareRevision (property field)
	GetFirmwareRevision() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataFirmwareRevisionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFirmwareRevision.
// This is useful for switch cases.
type BACnetConstructedDataFirmwareRevisionExactly interface {
	BACnetConstructedDataFirmwareRevision
	isBACnetConstructedDataFirmwareRevision() bool
}

// _BACnetConstructedDataFirmwareRevision is the data-structure of this message
type _BACnetConstructedDataFirmwareRevision struct {
	*_BACnetConstructedData
	FirmwareRevision BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFirmwareRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFirmwareRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FIRMWARE_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFirmwareRevision) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFirmwareRevision) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFirmwareRevision) GetFirmwareRevision() BACnetApplicationTagCharacterString {
	return m.FirmwareRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFirmwareRevision) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetFirmwareRevision())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFirmwareRevision factory function for _BACnetConstructedDataFirmwareRevision
func NewBACnetConstructedDataFirmwareRevision(firmwareRevision BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFirmwareRevision {
	_result := &_BACnetConstructedDataFirmwareRevision{
		FirmwareRevision:       firmwareRevision,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFirmwareRevision(structType interface{}) BACnetConstructedDataFirmwareRevision {
	if casted, ok := structType.(BACnetConstructedDataFirmwareRevision); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFirmwareRevision); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFirmwareRevision) GetTypeName() string {
	return "BACnetConstructedDataFirmwareRevision"
}

func (m *_BACnetConstructedDataFirmwareRevision) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (firmwareRevision)
	lengthInBits += m.FirmwareRevision.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFirmwareRevision) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFirmwareRevisionParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFirmwareRevision, error) {
	return BACnetConstructedDataFirmwareRevisionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFirmwareRevisionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFirmwareRevision, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFirmwareRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFirmwareRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (firmwareRevision)
	if pullErr := readBuffer.PullContext("firmwareRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for firmwareRevision")
	}
	_firmwareRevision, _firmwareRevisionErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _firmwareRevisionErr != nil {
		return nil, errors.Wrap(_firmwareRevisionErr, "Error parsing 'firmwareRevision' field of BACnetConstructedDataFirmwareRevision")
	}
	firmwareRevision := _firmwareRevision.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("firmwareRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for firmwareRevision")
	}

	// Virtual field
	_actualValue := firmwareRevision
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFirmwareRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFirmwareRevision")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFirmwareRevision{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FirmwareRevision: firmwareRevision,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFirmwareRevision) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFirmwareRevision) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFirmwareRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFirmwareRevision")
		}

		// Simple Field (firmwareRevision)
		if pushErr := writeBuffer.PushContext("firmwareRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for firmwareRevision")
		}
		_firmwareRevisionErr := writeBuffer.WriteSerializable(ctx, m.GetFirmwareRevision())
		if popErr := writeBuffer.PopContext("firmwareRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for firmwareRevision")
		}
		if _firmwareRevisionErr != nil {
			return errors.Wrap(_firmwareRevisionErr, "Error serializing 'firmwareRevision' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFirmwareRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFirmwareRevision")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFirmwareRevision) isBACnetConstructedDataFirmwareRevision() bool {
	return true
}

func (m *_BACnetConstructedDataFirmwareRevision) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
