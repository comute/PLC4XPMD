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

// BACnetConstructedDataEntryPoints is the data-structure of this message
type BACnetConstructedDataEntryPoints struct {
	*BACnetConstructedData
	EntryPoints []*BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEntryPoints is the corresponding interface of BACnetConstructedDataEntryPoints
type IBACnetConstructedDataEntryPoints interface {
	IBACnetConstructedData
	// GetEntryPoints returns EntryPoints (property field)
	GetEntryPoints() []*BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataEntryPoints) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEntryPoints) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENTRY_POINTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEntryPoints) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEntryPoints) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEntryPoints) GetEntryPoints() []*BACnetDeviceObjectReference {
	return m.EntryPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEntryPoints factory function for BACnetConstructedDataEntryPoints
func NewBACnetConstructedDataEntryPoints(entryPoints []*BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEntryPoints {
	_result := &BACnetConstructedDataEntryPoints{
		EntryPoints:           entryPoints,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEntryPoints(structType interface{}) *BACnetConstructedDataEntryPoints {
	if casted, ok := structType.(BACnetConstructedDataEntryPoints); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEntryPoints); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEntryPoints(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEntryPoints(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEntryPoints) GetTypeName() string {
	return "BACnetConstructedDataEntryPoints"
}

func (m *BACnetConstructedDataEntryPoints) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEntryPoints) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.EntryPoints) > 0 {
		for _, element := range m.EntryPoints {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataEntryPoints) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEntryPointsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEntryPoints, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEntryPoints"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEntryPoints")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (entryPoints)
	if pullErr := readBuffer.PullContext("entryPoints", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for entryPoints")
	}
	// Terminated array
	entryPoints := make([]*BACnetDeviceObjectReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'entryPoints' field")
			}
			entryPoints = append(entryPoints, CastBACnetDeviceObjectReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("entryPoints", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for entryPoints")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEntryPoints"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEntryPoints")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEntryPoints{
		EntryPoints:           entryPoints,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEntryPoints) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEntryPoints"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEntryPoints")
		}

		// Array Field (entryPoints)
		if m.EntryPoints != nil {
			if pushErr := writeBuffer.PushContext("entryPoints", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for entryPoints")
			}
			for _, _element := range m.EntryPoints {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'entryPoints' field")
				}
			}
			if popErr := writeBuffer.PopContext("entryPoints", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for entryPoints")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEntryPoints"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEntryPoints")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEntryPoints) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
