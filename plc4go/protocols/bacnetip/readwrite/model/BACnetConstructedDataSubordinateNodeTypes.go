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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSubordinateNodeTypes is the data-structure of this message
type BACnetConstructedDataSubordinateNodeTypes struct {
	*BACnetConstructedData
	NumberOfDataElements *BACnetApplicationTagUnsignedInteger
	SubordinateNodeTypes []*BACnetNodeTypeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSubordinateNodeTypes is the corresponding interface of BACnetConstructedDataSubordinateNodeTypes
type IBACnetConstructedDataSubordinateNodeTypes interface {
	IBACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger
	// GetSubordinateNodeTypes returns SubordinateNodeTypes (property field)
	GetSubordinateNodeTypes() []*BACnetNodeTypeTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
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

func (m *BACnetConstructedDataSubordinateNodeTypes) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_NODE_TYPES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSubordinateNodeTypes) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSubordinateNodeTypes) GetNumberOfDataElements() *BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetSubordinateNodeTypes() []*BACnetNodeTypeTagged {
	return m.SubordinateNodeTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataSubordinateNodeTypes) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateNodeTypes factory function for BACnetConstructedDataSubordinateNodeTypes
func NewBACnetConstructedDataSubordinateNodeTypes(numberOfDataElements *BACnetApplicationTagUnsignedInteger, subordinateNodeTypes []*BACnetNodeTypeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSubordinateNodeTypes {
	_result := &BACnetConstructedDataSubordinateNodeTypes{
		NumberOfDataElements:  numberOfDataElements,
		SubordinateNodeTypes:  subordinateNodeTypes,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSubordinateNodeTypes(structType interface{}) *BACnetConstructedDataSubordinateNodeTypes {
	if casted, ok := structType.(BACnetConstructedDataSubordinateNodeTypes); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateNodeTypes); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateNodeTypes(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateNodeTypes(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetTypeName() string {
	return "BACnetConstructedDataSubordinateNodeTypes"
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += (*m.NumberOfDataElements).GetLengthInBits()
	}

	// Array field
	if len(m.SubordinateNodeTypes) > 0 {
		for _, element := range m.SubordinateNodeTypes {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubordinateNodeTypesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSubordinateNodeTypes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateNodeTypes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateNodeTypes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = CastBACnetApplicationTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (subordinateNodeTypes)
	if pullErr := readBuffer.PullContext("subordinateNodeTypes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subordinateNodeTypes")
	}
	// Terminated array
	subordinateNodeTypes := make([]*BACnetNodeTypeTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetNodeTypeTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateNodeTypes' field")
			}
			subordinateNodeTypes = append(subordinateNodeTypes, CastBACnetNodeTypeTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("subordinateNodeTypes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subordinateNodeTypes")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateNodeTypes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateNodeTypes")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSubordinateNodeTypes{
		NumberOfDataElements:  CastBACnetApplicationTagUnsignedInteger(numberOfDataElements),
		SubordinateNodeTypes:  subordinateNodeTypes,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSubordinateNodeTypes) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateNodeTypes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateNodeTypes")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements *BACnetApplicationTagUnsignedInteger = nil
		if m.NumberOfDataElements != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.NumberOfDataElements
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (subordinateNodeTypes)
		if m.SubordinateNodeTypes != nil {
			if pushErr := writeBuffer.PushContext("subordinateNodeTypes", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for subordinateNodeTypes")
			}
			for _, _element := range m.SubordinateNodeTypes {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'subordinateNodeTypes' field")
				}
			}
			if popErr := writeBuffer.PopContext("subordinateNodeTypes", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for subordinateNodeTypes")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateNodeTypes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateNodeTypes")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSubordinateNodeTypes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
