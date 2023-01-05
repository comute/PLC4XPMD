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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersChangeOfReliability is the corresponding interface of BACnetNotificationParametersChangeOfReliability
type BACnetNotificationParametersChangeOfReliability interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetPropertyValues returns PropertyValues (property field)
	GetPropertyValues() BACnetPropertyValues
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersChangeOfReliabilityExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfReliability.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfReliabilityExactly interface {
	BACnetNotificationParametersChangeOfReliability
	isBACnetNotificationParametersChangeOfReliability() bool
}

// _BACnetNotificationParametersChangeOfReliability is the data-structure of this message
type _BACnetNotificationParametersChangeOfReliability struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	Reliability     BACnetReliabilityTagged
	StatusFlags     BACnetStatusFlagsTagged
	PropertyValues  BACnetPropertyValues
	InnerClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfReliability) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfReliability) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetPropertyValues() BACnetPropertyValues {
	return m.PropertyValues
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfReliability factory function for _BACnetNotificationParametersChangeOfReliability
func NewBACnetNotificationParametersChangeOfReliability(innerOpeningTag BACnetOpeningTag, reliability BACnetReliabilityTagged, statusFlags BACnetStatusFlagsTagged, propertyValues BACnetPropertyValues, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfReliability {
	_result := &_BACnetNotificationParametersChangeOfReliability{
		InnerOpeningTag:               innerOpeningTag,
		Reliability:                   reliability,
		StatusFlags:                   statusFlags,
		PropertyValues:                propertyValues,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfReliability(structType interface{}) BACnetNotificationParametersChangeOfReliability {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfReliability"
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (propertyValues)
	lengthInBits += m.PropertyValues.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfReliabilityParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfReliability, error) {
	return BACnetNotificationParametersChangeOfReliabilityParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, peekedTagNumber)
}

func BACnetNotificationParametersChangeOfReliabilityParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfReliability, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersChangeOfReliability")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (reliability)
	if pullErr := readBuffer.PullContext("reliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reliability")
	}
	_reliability, _reliabilityErr := BACnetReliabilityTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reliabilityErr != nil {
		return nil, errors.Wrap(_reliabilityErr, "Error parsing 'reliability' field of BACnetNotificationParametersChangeOfReliability")
	}
	reliability := _reliability.(BACnetReliabilityTagged)
	if closeErr := readBuffer.CloseContext("reliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reliability")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParseWithBuffer(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field of BACnetNotificationParametersChangeOfReliability")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (propertyValues)
	if pullErr := readBuffer.PullContext("propertyValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyValues")
	}
	_propertyValues, _propertyValuesErr := BACnetPropertyValuesParseWithBuffer(readBuffer, uint8(uint8(2)), BACnetObjectType(objectTypeArgument))
	if _propertyValuesErr != nil {
		return nil, errors.Wrap(_propertyValuesErr, "Error parsing 'propertyValues' field of BACnetNotificationParametersChangeOfReliability")
	}
	propertyValues := _propertyValues.(BACnetPropertyValues)
	if closeErr := readBuffer.CloseContext("propertyValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyValues")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersChangeOfReliability")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfReliability")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfReliability{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag: innerOpeningTag,
		Reliability:     reliability,
		StatusFlags:     statusFlags,
		PropertyValues:  propertyValues,
		InnerClosingTag: innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfReliability")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (reliability)
		if pushErr := writeBuffer.PushContext("reliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reliability")
		}
		_reliabilityErr := writeBuffer.WriteSerializable(m.GetReliability())
		if popErr := writeBuffer.PopContext("reliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reliability")
		}
		if _reliabilityErr != nil {
			return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (propertyValues)
		if pushErr := writeBuffer.PushContext("propertyValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValues")
		}
		_propertyValuesErr := writeBuffer.WriteSerializable(m.GetPropertyValues())
		if popErr := writeBuffer.PopContext("propertyValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValues")
		}
		if _propertyValuesErr != nil {
			return errors.Wrap(_propertyValuesErr, "Error serializing 'propertyValues' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfReliability")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfReliability) isBACnetNotificationParametersChangeOfReliability() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
