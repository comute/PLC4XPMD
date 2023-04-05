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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyValue is the corresponding interface of BACnetPropertyValue
type BACnetPropertyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedDataElement
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetPropertyValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyValue.
// This is useful for switch cases.
type BACnetPropertyValueExactly interface {
	BACnetPropertyValue
	isBACnetPropertyValue() bool
}

// _BACnetPropertyValue is the data-structure of this message
type _BACnetPropertyValue struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	PropertyArrayIndex BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedDataElement
	Priority           BACnetContextTagUnsignedInteger

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyValue) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyValue) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetPropertyValue) GetPropertyValue() BACnetConstructedDataElement {
	return m.PropertyValue
}

func (m *_BACnetPropertyValue) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyValue factory function for _BACnetPropertyValue
func NewBACnetPropertyValue(propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedDataElement, priority BACnetContextTagUnsignedInteger, objectTypeArgument BACnetObjectType) *_BACnetPropertyValue {
	return &_BACnetPropertyValue{PropertyIdentifier: propertyIdentifier, PropertyArrayIndex: propertyArrayIndex, PropertyValue: propertyValue, Priority: priority, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyValue(structType interface{}) BACnetPropertyValue {
	if casted, ok := structType.(BACnetPropertyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyValue) GetTypeName() string {
	return "BACnetPropertyValue"
}

func (m *_BACnetPropertyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetPropertyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyValueParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPropertyValue, error) {
	return BACnetPropertyValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPropertyValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPropertyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetPropertyValue")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if a given expression evaluates to false)
	var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyArrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyArrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field of BACnetPropertyValue")
		default:
			propertyArrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyArrayIndex")
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue BACnetConstructedDataElement = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
		_val, _err := BACnetConstructedDataElementParseWithBuffer(ctx, readBuffer, objectTypeArgument, propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((propertyArrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field of BACnetPropertyValue")
		default:
			propertyValue = _val.(BACnetConstructedDataElement)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field of BACnetPropertyValue")
		default:
			priority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyValue")
	}

	// Create the instance
	return &_BACnetPropertyValue{
		ObjectTypeArgument: objectTypeArgument,
		PropertyIdentifier: propertyIdentifier,
		PropertyArrayIndex: propertyArrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
	}, nil
}

func (m *_BACnetPropertyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyValue")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
	var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetPropertyArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("propertyArrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyArrayIndex")
		}
		propertyArrayIndex = m.GetPropertyArrayIndex()
		_propertyArrayIndexErr := writeBuffer.WriteSerializable(ctx, propertyArrayIndex)
		if popErr := writeBuffer.PopContext("propertyArrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyArrayIndex")
		}
		if _propertyArrayIndexErr != nil {
			return errors.Wrap(_propertyArrayIndexErr, "Error serializing 'propertyArrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue BACnetConstructedDataElement = nil
	if m.GetPropertyValue() != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		propertyValue = m.GetPropertyValue()
		_propertyValueErr := writeBuffer.WriteSerializable(ctx, propertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority BACnetContextTagUnsignedInteger = nil
	if m.GetPriority() != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.GetPriority()
		_priorityErr := writeBuffer.WriteSerializable(ctx, priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPropertyValue) isBACnetPropertyValue() bool {
	return true
}

func (m *_BACnetPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
