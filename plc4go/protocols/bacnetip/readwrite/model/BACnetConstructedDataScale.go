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

// BACnetConstructedDataScale is the corresponding interface of BACnetConstructedDataScale
type BACnetConstructedDataScale interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetScale returns Scale (property field)
	GetScale() BACnetScale
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetScale
}

// BACnetConstructedDataScaleExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataScale.
// This is useful for switch cases.
type BACnetConstructedDataScaleExactly interface {
	BACnetConstructedDataScale
	isBACnetConstructedDataScale() bool
}

// _BACnetConstructedDataScale is the data-structure of this message
type _BACnetConstructedDataScale struct {
	*_BACnetConstructedData
	Scale BACnetScale
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataScale) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataScale) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SCALE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataScale) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataScale) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataScale) GetScale() BACnetScale {
	return m.Scale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataScale) GetActualValue() BACnetScale {
	return CastBACnetScale(m.GetScale())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataScale factory function for _BACnetConstructedDataScale
func NewBACnetConstructedDataScale(scale BACnetScale, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataScale {
	_result := &_BACnetConstructedDataScale{
		Scale:                  scale,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataScale(structType interface{}) BACnetConstructedDataScale {
	if casted, ok := structType.(BACnetConstructedDataScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataScale) GetTypeName() string {
	return "BACnetConstructedDataScale"
}

func (m *_BACnetConstructedDataScale) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataScale) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (scale)
	lengthInBits += m.Scale.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataScale) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataScaleParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataScale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (scale)
	if pullErr := readBuffer.PullContext("scale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for scale")
	}
	_scale, _scaleErr := BACnetScaleParse(readBuffer)
	if _scaleErr != nil {
		return nil, errors.Wrap(_scaleErr, "Error parsing 'scale' field")
	}
	scale := _scale.(BACnetScale)
	if closeErr := readBuffer.CloseContext("scale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for scale")
	}

	// Virtual field
	_actualValue := scale
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScale")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataScale{
		Scale: scale,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataScale) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScale")
		}

		// Simple Field (scale)
		if pushErr := writeBuffer.PushContext("scale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for scale")
		}
		_scaleErr := writeBuffer.WriteSerializable(m.GetScale())
		if popErr := writeBuffer.PopContext("scale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for scale")
		}
		if _scaleErr != nil {
			return errors.Wrap(_scaleErr, "Error serializing 'scale' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScale")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataScale) isBACnetConstructedDataScale() bool {
	return true
}

func (m *_BACnetConstructedDataScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
