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

// BACnetConstructedDataTransactionNotificationClass is the data-structure of this message
type BACnetConstructedDataTransactionNotificationClass struct {
	*BACnetConstructedData
	TransactionNotificationClass *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTransactionNotificationClass is the corresponding interface of BACnetConstructedDataTransactionNotificationClass
type IBACnetConstructedDataTransactionNotificationClass interface {
	IBACnetConstructedData
	// GetTransactionNotificationClass returns TransactionNotificationClass (property field)
	GetTransactionNotificationClass() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataTransactionNotificationClass) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRANSACTION_NOTIFICATION_CLASS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTransactionNotificationClass) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTransactionNotificationClass) GetTransactionNotificationClass() *BACnetApplicationTagUnsignedInteger {
	return m.TransactionNotificationClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTransactionNotificationClass factory function for BACnetConstructedDataTransactionNotificationClass
func NewBACnetConstructedDataTransactionNotificationClass(transactionNotificationClass *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTransactionNotificationClass {
	_result := &BACnetConstructedDataTransactionNotificationClass{
		TransactionNotificationClass: transactionNotificationClass,
		BACnetConstructedData:        NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTransactionNotificationClass(structType interface{}) *BACnetConstructedDataTransactionNotificationClass {
	if casted, ok := structType.(BACnetConstructedDataTransactionNotificationClass); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTransactionNotificationClass); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTransactionNotificationClass(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTransactionNotificationClass(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetTypeName() string {
	return "BACnetConstructedDataTransactionNotificationClass"
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (transactionNotificationClass)
	lengthInBits += m.TransactionNotificationClass.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTransactionNotificationClass) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTransactionNotificationClassParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTransactionNotificationClass, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTransactionNotificationClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTransactionNotificationClass")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (transactionNotificationClass)
	if pullErr := readBuffer.PullContext("transactionNotificationClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transactionNotificationClass")
	}
	_transactionNotificationClass, _transactionNotificationClassErr := BACnetApplicationTagParse(readBuffer)
	if _transactionNotificationClassErr != nil {
		return nil, errors.Wrap(_transactionNotificationClassErr, "Error parsing 'transactionNotificationClass' field")
	}
	transactionNotificationClass := CastBACnetApplicationTagUnsignedInteger(_transactionNotificationClass)
	if closeErr := readBuffer.CloseContext("transactionNotificationClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transactionNotificationClass")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTransactionNotificationClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTransactionNotificationClass")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTransactionNotificationClass{
		TransactionNotificationClass: CastBACnetApplicationTagUnsignedInteger(transactionNotificationClass),
		BACnetConstructedData:        &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTransactionNotificationClass) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTransactionNotificationClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTransactionNotificationClass")
		}

		// Simple Field (transactionNotificationClass)
		if pushErr := writeBuffer.PushContext("transactionNotificationClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transactionNotificationClass")
		}
		_transactionNotificationClassErr := writeBuffer.WriteSerializable(m.TransactionNotificationClass)
		if popErr := writeBuffer.PopContext("transactionNotificationClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transactionNotificationClass")
		}
		if _transactionNotificationClassErr != nil {
			return errors.Wrap(_transactionNotificationClassErr, "Error serializing 'transactionNotificationClass' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTransactionNotificationClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTransactionNotificationClass")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTransactionNotificationClass) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
