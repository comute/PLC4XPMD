/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataFDSubscriptionLifetime is the data-structure of this message
type BACnetConstructedDataFDSubscriptionLifetime struct {
	*BACnetConstructedData
	FdSubscriptionLifetime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataFDSubscriptionLifetime is the corresponding interface of BACnetConstructedDataFDSubscriptionLifetime
type IBACnetConstructedDataFDSubscriptionLifetime interface {
	IBACnetConstructedData
	// GetFdSubscriptionLifetime returns FdSubscriptionLifetime (property field)
	GetFdSubscriptionLifetime() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FD_SUBSCRIPTION_LIFETIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFDSubscriptionLifetime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetFdSubscriptionLifetime() *BACnetApplicationTagUnsignedInteger {
	return m.FdSubscriptionLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFDSubscriptionLifetime factory function for BACnetConstructedDataFDSubscriptionLifetime
func NewBACnetConstructedDataFDSubscriptionLifetime(fdSubscriptionLifetime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataFDSubscriptionLifetime {
	_result := &BACnetConstructedDataFDSubscriptionLifetime{
		FdSubscriptionLifetime: fdSubscriptionLifetime,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFDSubscriptionLifetime(structType interface{}) *BACnetConstructedDataFDSubscriptionLifetime {
	if casted, ok := structType.(BACnetConstructedDataFDSubscriptionLifetime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFDSubscriptionLifetime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFDSubscriptionLifetime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFDSubscriptionLifetime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetTypeName() string {
	return "BACnetConstructedDataFDSubscriptionLifetime"
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fdSubscriptionLifetime)
	lengthInBits += m.FdSubscriptionLifetime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFDSubscriptionLifetimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataFDSubscriptionLifetime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFDSubscriptionLifetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFDSubscriptionLifetime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fdSubscriptionLifetime)
	if pullErr := readBuffer.PullContext("fdSubscriptionLifetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fdSubscriptionLifetime")
	}
	_fdSubscriptionLifetime, _fdSubscriptionLifetimeErr := BACnetApplicationTagParse(readBuffer)
	if _fdSubscriptionLifetimeErr != nil {
		return nil, errors.Wrap(_fdSubscriptionLifetimeErr, "Error parsing 'fdSubscriptionLifetime' field")
	}
	fdSubscriptionLifetime := CastBACnetApplicationTagUnsignedInteger(_fdSubscriptionLifetime)
	if closeErr := readBuffer.CloseContext("fdSubscriptionLifetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fdSubscriptionLifetime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFDSubscriptionLifetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFDSubscriptionLifetime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFDSubscriptionLifetime{
		FdSubscriptionLifetime: CastBACnetApplicationTagUnsignedInteger(fdSubscriptionLifetime),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFDSubscriptionLifetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFDSubscriptionLifetime")
		}

		// Simple Field (fdSubscriptionLifetime)
		if pushErr := writeBuffer.PushContext("fdSubscriptionLifetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fdSubscriptionLifetime")
		}
		_fdSubscriptionLifetimeErr := writeBuffer.WriteSerializable(m.FdSubscriptionLifetime)
		if popErr := writeBuffer.PopContext("fdSubscriptionLifetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fdSubscriptionLifetime")
		}
		if _fdSubscriptionLifetimeErr != nil {
			return errors.Wrap(_fdSubscriptionLifetimeErr, "Error serializing 'fdSubscriptionLifetime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFDSubscriptionLifetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFDSubscriptionLifetime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFDSubscriptionLifetime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
