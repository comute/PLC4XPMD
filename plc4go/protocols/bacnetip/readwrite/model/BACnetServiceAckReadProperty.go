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

// BACnetServiceAckReadProperty is the data-structure of this message
type BACnetServiceAckReadProperty struct {
	*BACnetServiceAck
	ObjectIdentifier   *BACnetContextTagObjectIdentifier
	PropertyIdentifier *BACnetPropertyIdentifierTagged
	ArrayIndex         *BACnetContextTagUnsignedInteger
	Values             *BACnetConstructedData

	// Arguments.
	ServiceAckLength uint16
}

// IBACnetServiceAckReadProperty is the corresponding interface of BACnetServiceAckReadProperty
type IBACnetServiceAckReadProperty interface {
	IBACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() *BACnetContextTagUnsignedInteger
	// GetValues returns Values (property field)
	GetValues() *BACnetConstructedData
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

func (m *BACnetServiceAckReadProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckReadProperty) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckReadProperty) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckReadProperty) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetServiceAckReadProperty) GetPropertyIdentifier() *BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *BACnetServiceAckReadProperty) GetArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *BACnetServiceAckReadProperty) GetValues() *BACnetConstructedData {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadProperty factory function for BACnetServiceAckReadProperty
func NewBACnetServiceAckReadProperty(objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetPropertyIdentifierTagged, arrayIndex *BACnetContextTagUnsignedInteger, values *BACnetConstructedData, serviceAckLength uint16) *BACnetServiceAckReadProperty {
	_result := &BACnetServiceAckReadProperty{
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		Values:             values,
		BACnetServiceAck:   NewBACnetServiceAck(serviceAckLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckReadProperty(structType interface{}) *BACnetServiceAckReadProperty {
	if casted, ok := structType.(BACnetServiceAckReadProperty); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadProperty); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckReadProperty(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckReadProperty(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckReadProperty) GetTypeName() string {
	return "BACnetServiceAckReadProperty"
}

func (m *BACnetServiceAckReadProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).GetLengthInBits()
	}

	// Optional Field (values)
	if m.Values != nil {
		lengthInBits += (*m.Values).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetServiceAckReadProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckReadPropertyParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (*BACnetServiceAckReadProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckReadProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetPropertyIdentifierTagged(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (values) (Can be skipped, if a given expression evaluates to false)
	var values *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("values"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for values")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(3), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue(), CastBACnetTagPayloadUnsignedInteger(CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((*arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'values' field")
		default:
			values = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("values"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for values")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckReadProperty")
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadProperty{
		ObjectIdentifier:   CastBACnetContextTagObjectIdentifier(objectIdentifier),
		PropertyIdentifier: CastBACnetPropertyIdentifierTagged(propertyIdentifier),
		ArrayIndex:         CastBACnetContextTagUnsignedInteger(arrayIndex),
		Values:             CastBACnetConstructedData(values),
		BACnetServiceAck:   &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckReadProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckReadProperty")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
		}
		_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyIdentifier")
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (arrayIndex) (Can be skipped, if the value is null)
		var arrayIndex *BACnetContextTagUnsignedInteger = nil
		if m.ArrayIndex != nil {
			if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for arrayIndex")
			}
			arrayIndex = m.ArrayIndex
			_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for arrayIndex")
			}
			if _arrayIndexErr != nil {
				return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
			}
		}

		// Optional Field (values) (Can be skipped, if the value is null)
		var values *BACnetConstructedData = nil
		if m.Values != nil {
			if pushErr := writeBuffer.PushContext("values"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for values")
			}
			values = m.Values
			_valuesErr := values.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("values"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for values")
			}
			if _valuesErr != nil {
				return errors.Wrap(_valuesErr, "Error serializing 'values' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckReadProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
