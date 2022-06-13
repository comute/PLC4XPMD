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

// BACnetConstructedDataLightingOutputTrackingValue is the data-structure of this message
type BACnetConstructedDataLightingOutputTrackingValue struct {
	*BACnetConstructedData
	TrackingValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataLightingOutputTrackingValue is the corresponding interface of BACnetConstructedDataLightingOutputTrackingValue
type IBACnetConstructedDataLightingOutputTrackingValue interface {
	IBACnetConstructedData
	// GetTrackingValue returns TrackingValue (property field)
	GetTrackingValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIGHTING_OUTPUT
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRACKING_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLightingOutputTrackingValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetTrackingValue() *BACnetApplicationTagReal {
	return m.TrackingValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingOutputTrackingValue factory function for BACnetConstructedDataLightingOutputTrackingValue
func NewBACnetConstructedDataLightingOutputTrackingValue(trackingValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataLightingOutputTrackingValue {
	_result := &BACnetConstructedDataLightingOutputTrackingValue{
		TrackingValue:         trackingValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLightingOutputTrackingValue(structType interface{}) *BACnetConstructedDataLightingOutputTrackingValue {
	if casted, ok := structType.(BACnetConstructedDataLightingOutputTrackingValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingOutputTrackingValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingOutputTrackingValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLightingOutputTrackingValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetTypeName() string {
	return "BACnetConstructedDataLightingOutputTrackingValue"
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (trackingValue)
	lengthInBits += m.TrackingValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLightingOutputTrackingValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataLightingOutputTrackingValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingOutputTrackingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingOutputTrackingValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (trackingValue)
	if pullErr := readBuffer.PullContext("trackingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for trackingValue")
	}
	_trackingValue, _trackingValueErr := BACnetApplicationTagParse(readBuffer)
	if _trackingValueErr != nil {
		return nil, errors.Wrap(_trackingValueErr, "Error parsing 'trackingValue' field")
	}
	trackingValue := CastBACnetApplicationTagReal(_trackingValue)
	if closeErr := readBuffer.CloseContext("trackingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for trackingValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingOutputTrackingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingOutputTrackingValue")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLightingOutputTrackingValue{
		TrackingValue:         CastBACnetApplicationTagReal(trackingValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingOutputTrackingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingOutputTrackingValue")
		}

		// Simple Field (trackingValue)
		if pushErr := writeBuffer.PushContext("trackingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for trackingValue")
		}
		_trackingValueErr := writeBuffer.WriteSerializable(m.TrackingValue)
		if popErr := writeBuffer.PopContext("trackingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for trackingValue")
		}
		if _trackingValueErr != nil {
			return errors.Wrap(_trackingValueErr, "Error serializing 'trackingValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingOutputTrackingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingOutputTrackingValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLightingOutputTrackingValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
