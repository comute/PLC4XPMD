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

// BACnetConstructedDataTotalRecordCount is the data-structure of this message
type BACnetConstructedDataTotalRecordCount struct {
	*BACnetConstructedData
	TotalRecordCount *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTotalRecordCount is the corresponding interface of BACnetConstructedDataTotalRecordCount
type IBACnetConstructedDataTotalRecordCount interface {
	IBACnetConstructedData
	// GetTotalRecordCount returns TotalRecordCount (property field)
	GetTotalRecordCount() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataTotalRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTotalRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TOTAL_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTotalRecordCount) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTotalRecordCount) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTotalRecordCount) GetTotalRecordCount() *BACnetApplicationTagUnsignedInteger {
	return m.TotalRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTotalRecordCount factory function for BACnetConstructedDataTotalRecordCount
func NewBACnetConstructedDataTotalRecordCount(totalRecordCount *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTotalRecordCount {
	_result := &BACnetConstructedDataTotalRecordCount{
		TotalRecordCount:      totalRecordCount,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTotalRecordCount(structType interface{}) *BACnetConstructedDataTotalRecordCount {
	if casted, ok := structType.(BACnetConstructedDataTotalRecordCount); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTotalRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTotalRecordCount(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTotalRecordCount(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTotalRecordCount) GetTypeName() string {
	return "BACnetConstructedDataTotalRecordCount"
}

func (m *BACnetConstructedDataTotalRecordCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTotalRecordCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (totalRecordCount)
	lengthInBits += m.TotalRecordCount.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTotalRecordCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTotalRecordCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTotalRecordCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTotalRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTotalRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (totalRecordCount)
	if pullErr := readBuffer.PullContext("totalRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for totalRecordCount")
	}
	_totalRecordCount, _totalRecordCountErr := BACnetApplicationTagParse(readBuffer)
	if _totalRecordCountErr != nil {
		return nil, errors.Wrap(_totalRecordCountErr, "Error parsing 'totalRecordCount' field")
	}
	totalRecordCount := CastBACnetApplicationTagUnsignedInteger(_totalRecordCount)
	if closeErr := readBuffer.CloseContext("totalRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for totalRecordCount")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTotalRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTotalRecordCount")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTotalRecordCount{
		TotalRecordCount:      CastBACnetApplicationTagUnsignedInteger(totalRecordCount),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTotalRecordCount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTotalRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTotalRecordCount")
		}

		// Simple Field (totalRecordCount)
		if pushErr := writeBuffer.PushContext("totalRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for totalRecordCount")
		}
		_totalRecordCountErr := m.TotalRecordCount.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("totalRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for totalRecordCount")
		}
		if _totalRecordCountErr != nil {
			return errors.Wrap(_totalRecordCountErr, "Error serializing 'totalRecordCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTotalRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTotalRecordCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTotalRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
