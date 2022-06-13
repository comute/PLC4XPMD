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

// BACnetConstructedDataEventAlgorithmInhibitRef is the data-structure of this message
type BACnetConstructedDataEventAlgorithmInhibitRef struct {
	*BACnetConstructedData
	EventAlgorithmInhibitRef *BACnetObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEventAlgorithmInhibitRef is the corresponding interface of BACnetConstructedDataEventAlgorithmInhibitRef
type IBACnetConstructedDataEventAlgorithmInhibitRef interface {
	IBACnetConstructedData
	// GetEventAlgorithmInhibitRef returns EventAlgorithmInhibitRef (property field)
	GetEventAlgorithmInhibitRef() *BACnetObjectPropertyReference
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

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ALGORITHM_INHIBIT_REF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetEventAlgorithmInhibitRef() *BACnetObjectPropertyReference {
	return m.EventAlgorithmInhibitRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventAlgorithmInhibitRef factory function for BACnetConstructedDataEventAlgorithmInhibitRef
func NewBACnetConstructedDataEventAlgorithmInhibitRef(eventAlgorithmInhibitRef *BACnetObjectPropertyReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEventAlgorithmInhibitRef {
	_result := &BACnetConstructedDataEventAlgorithmInhibitRef{
		EventAlgorithmInhibitRef: eventAlgorithmInhibitRef,
		BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEventAlgorithmInhibitRef(structType interface{}) *BACnetConstructedDataEventAlgorithmInhibitRef {
	if casted, ok := structType.(BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventAlgorithmInhibitRef(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEventAlgorithmInhibitRef(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetTypeName() string {
	return "BACnetConstructedDataEventAlgorithmInhibitRef"
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventAlgorithmInhibitRef)
	lengthInBits += m.EventAlgorithmInhibitRef.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventAlgorithmInhibitRefParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEventAlgorithmInhibitRef, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventAlgorithmInhibitRef")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventAlgorithmInhibitRef)
	if pullErr := readBuffer.PullContext("eventAlgorithmInhibitRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventAlgorithmInhibitRef")
	}
	_eventAlgorithmInhibitRef, _eventAlgorithmInhibitRefErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _eventAlgorithmInhibitRefErr != nil {
		return nil, errors.Wrap(_eventAlgorithmInhibitRefErr, "Error parsing 'eventAlgorithmInhibitRef' field")
	}
	eventAlgorithmInhibitRef := CastBACnetObjectPropertyReference(_eventAlgorithmInhibitRef)
	if closeErr := readBuffer.CloseContext("eventAlgorithmInhibitRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventAlgorithmInhibitRef")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventAlgorithmInhibitRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventAlgorithmInhibitRef")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEventAlgorithmInhibitRef{
		EventAlgorithmInhibitRef: CastBACnetObjectPropertyReference(eventAlgorithmInhibitRef),
		BACnetConstructedData:    &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventAlgorithmInhibitRef")
		}

		// Simple Field (eventAlgorithmInhibitRef)
		if pushErr := writeBuffer.PushContext("eventAlgorithmInhibitRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventAlgorithmInhibitRef")
		}
		_eventAlgorithmInhibitRefErr := writeBuffer.WriteSerializable(m.EventAlgorithmInhibitRef)
		if popErr := writeBuffer.PopContext("eventAlgorithmInhibitRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventAlgorithmInhibitRef")
		}
		if _eventAlgorithmInhibitRefErr != nil {
			return errors.Wrap(_eventAlgorithmInhibitRefErr, "Error serializing 'eventAlgorithmInhibitRef' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventAlgorithmInhibitRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventAlgorithmInhibitRef")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEventAlgorithmInhibitRef) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
