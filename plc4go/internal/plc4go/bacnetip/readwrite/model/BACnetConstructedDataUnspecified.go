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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetConstructedDataUnspecified struct {
	*BACnetConstructedData
	Data               []*BACnetConstructedDataElement
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	Content            *BACnetApplicationTag

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
}

// The corresponding interface
type IBACnetConstructedDataUnspecified interface {
	IBACnetConstructedData
	// GetData returns Data (property field)
	GetData() []*BACnetConstructedDataElement
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier
	// GetContent returns Content (property field)
	GetContent() *BACnetApplicationTag
	// GetHasData returns HasData (virtual field)
	GetHasData() bool
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
func (m *BACnetConstructedDataUnspecified) GetObjectType() BACnetObjectType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataUnspecified) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataUnspecified) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetConstructedDataUnspecified) GetData() []*BACnetConstructedDataElement {
	return m.Data
}

func (m *BACnetConstructedDataUnspecified) GetPropertyIdentifier() *BACnetContextTagPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetConstructedDataUnspecified) GetContent() *BACnetApplicationTag {
	return m.Content
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *BACnetConstructedDataUnspecified) GetHasData() bool {
	propertyIdentifier := m.PropertyIdentifier
	_ = propertyIdentifier
	content := m.Content
	_ = content
	return bool(bool((len(m.GetData())) == (0)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUnspecified factory function for BACnetConstructedDataUnspecified
func NewBACnetConstructedDataUnspecified(data []*BACnetConstructedDataElement, propertyIdentifier *BACnetContextTagPropertyIdentifier, content *BACnetApplicationTag, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedDataUnspecified {
	_result := &BACnetConstructedDataUnspecified{
		Data:                  data,
		PropertyIdentifier:    propertyIdentifier,
		Content:               content,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber, propertyIdentifierArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataUnspecified(structType interface{}) *BACnetConstructedDataUnspecified {
	if casted, ok := structType.(BACnetConstructedDataUnspecified); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUnspecified); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataUnspecified(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataUnspecified(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataUnspecified) GetTypeName() string {
	return "BACnetConstructedDataUnspecified"
}

func (m *BACnetConstructedDataUnspecified) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataUnspecified) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (propertyIdentifier)
	if m.PropertyIdentifier != nil {
		lengthInBits += (*m.PropertyIdentifier).GetLengthInBits()
	}

	// Optional Field (content)
	if m.Content != nil {
		lengthInBits += (*m.Content).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConstructedDataUnspecified) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUnspecifiedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedDataUnspecified, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUnspecified"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	data := make([]*BACnetConstructedDataElement, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetConstructedDataElementParse(readBuffer, objectType, propertyIdentifierArgument)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, CastBACnetConstructedDataElement(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_hasData := bool((len(data)) == (0))
	hasData := bool(_hasData)
	_ = hasData

	// Optional Field (propertyIdentifier) (Can be skipped, if a given expression evaluates to false)
	var propertyIdentifier *BACnetContextTagPropertyIdentifier = nil
	if hasData {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_PROPERTY_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyIdentifier' field")
		default:
			propertyIdentifier = CastBACnetContextTagPropertyIdentifier(_val)
			if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (content) (Can be skipped, if a given expression evaluates to false)
	var content *BACnetApplicationTag = nil
	if hasData {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("content"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'content' field")
		default:
			content = CastBACnetApplicationTag(_val)
			if closeErr := readBuffer.CloseContext("content"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUnspecified"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataUnspecified{
		Data:                  data,
		PropertyIdentifier:    CastBACnetContextTagPropertyIdentifier(propertyIdentifier),
		Content:               CastBACnetApplicationTag(content),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataUnspecified) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUnspecified"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}
		// Virtual field
		if _hasDataErr := writeBuffer.WriteVirtual("hasData", m.GetHasData()); _hasDataErr != nil {
			return errors.Wrap(_hasDataErr, "Error serializing 'hasData' field")
		}

		// Optional Field (propertyIdentifier) (Can be skipped, if the value is null)
		var propertyIdentifier *BACnetContextTagPropertyIdentifier = nil
		if m.PropertyIdentifier != nil {
			if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
				return pushErr
			}
			propertyIdentifier = m.PropertyIdentifier
			_propertyIdentifierErr := propertyIdentifier.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
				return popErr
			}
			if _propertyIdentifierErr != nil {
				return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
			}
		}

		// Optional Field (content) (Can be skipped, if the value is null)
		var content *BACnetApplicationTag = nil
		if m.Content != nil {
			if pushErr := writeBuffer.PushContext("content"); pushErr != nil {
				return pushErr
			}
			content = m.Content
			_contentErr := content.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("content"); popErr != nil {
				return popErr
			}
			if _contentErr != nil {
				return errors.Wrap(_contentErr, "Error serializing 'content' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUnspecified"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataUnspecified) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
