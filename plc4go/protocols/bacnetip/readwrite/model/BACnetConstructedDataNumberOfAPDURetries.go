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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataNumberOfAPDURetries is the corresponding interface of BACnetConstructedDataNumberOfAPDURetries
type BACnetConstructedDataNumberOfAPDURetries interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfApduRetries returns NumberOfApduRetries (property field)
	GetNumberOfApduRetries() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataNumberOfAPDURetriesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNumberOfAPDURetries.
// This is useful for switch cases.
type BACnetConstructedDataNumberOfAPDURetriesExactly interface {
	BACnetConstructedDataNumberOfAPDURetries
	isBACnetConstructedDataNumberOfAPDURetries() bool
}

// _BACnetConstructedDataNumberOfAPDURetries is the data-structure of this message
type _BACnetConstructedDataNumberOfAPDURetries struct {
	*_BACnetConstructedData
	NumberOfApduRetries BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_APDU_RETRIES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetNumberOfApduRetries() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfApduRetries
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetNumberOfApduRetries())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNumberOfAPDURetries factory function for _BACnetConstructedDataNumberOfAPDURetries
func NewBACnetConstructedDataNumberOfAPDURetries(numberOfApduRetries BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNumberOfAPDURetries {
	_result := &_BACnetConstructedDataNumberOfAPDURetries{
		NumberOfApduRetries:    numberOfApduRetries,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNumberOfAPDURetries(structType interface{}) BACnetConstructedDataNumberOfAPDURetries {
	if casted, ok := structType.(BACnetConstructedDataNumberOfAPDURetries); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfAPDURetries); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetTypeName() string {
	return "BACnetConstructedDataNumberOfAPDURetries"
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numberOfApduRetries)
	lengthInBits += m.NumberOfApduRetries.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNumberOfAPDURetriesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNumberOfAPDURetries, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfAPDURetries"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNumberOfAPDURetries")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfApduRetries)
	if pullErr := readBuffer.PullContext("numberOfApduRetries"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for numberOfApduRetries")
	}
	_numberOfApduRetries, _numberOfApduRetriesErr := BACnetApplicationTagParse(readBuffer)
	if _numberOfApduRetriesErr != nil {
		return nil, errors.Wrap(_numberOfApduRetriesErr, "Error parsing 'numberOfApduRetries' field")
	}
	numberOfApduRetries := _numberOfApduRetries.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("numberOfApduRetries"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for numberOfApduRetries")
	}

	// Virtual field
	_actualValue := numberOfApduRetries
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfAPDURetries"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNumberOfAPDURetries")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNumberOfAPDURetries{
		NumberOfApduRetries:    numberOfApduRetries,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfAPDURetries"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNumberOfAPDURetries")
		}

		// Simple Field (numberOfApduRetries)
		if pushErr := writeBuffer.PushContext("numberOfApduRetries"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfApduRetries")
		}
		_numberOfApduRetriesErr := writeBuffer.WriteSerializable(m.GetNumberOfApduRetries())
		if popErr := writeBuffer.PopContext("numberOfApduRetries"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfApduRetries")
		}
		if _numberOfApduRetriesErr != nil {
			return errors.Wrap(_numberOfApduRetriesErr, "Error serializing 'numberOfApduRetries' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfAPDURetries"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNumberOfAPDURetries")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) isBACnetConstructedDataNumberOfAPDURetries() bool {
	return true
}

func (m *_BACnetConstructedDataNumberOfAPDURetries) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
