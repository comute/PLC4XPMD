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

// BACnetActionCommand is the corresponding interface of BACnetActionCommand
type BACnetActionCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// GetPostDelay returns PostDelay (property field)
	GetPostDelay() BACnetContextTagBoolean
	// GetQuitOnFailure returns QuitOnFailure (property field)
	GetQuitOnFailure() BACnetContextTagBoolean
	// GetWriteSuccessful returns WriteSuccessful (property field)
	GetWriteSuccessful() BACnetContextTagBoolean
}

// BACnetActionCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetActionCommand.
// This is useful for switch cases.
type BACnetActionCommandExactly interface {
	BACnetActionCommand
	isBACnetActionCommand() bool
}

// _BACnetActionCommand is the data-structure of this message
type _BACnetActionCommand struct {
	DeviceIdentifier   BACnetContextTagObjectIdentifier
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger
	PostDelay          BACnetContextTagBoolean
	QuitOnFailure      BACnetContextTagBoolean
	WriteSuccessful    BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetActionCommand) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetActionCommand) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetActionCommand) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetActionCommand) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetActionCommand) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetActionCommand) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetActionCommand) GetPostDelay() BACnetContextTagBoolean {
	return m.PostDelay
}

func (m *_BACnetActionCommand) GetQuitOnFailure() BACnetContextTagBoolean {
	return m.QuitOnFailure
}

func (m *_BACnetActionCommand) GetWriteSuccessful() BACnetContextTagBoolean {
	return m.WriteSuccessful
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetActionCommand factory function for _BACnetActionCommand
func NewBACnetActionCommand(deviceIdentifier BACnetContextTagObjectIdentifier, objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, postDelay BACnetContextTagBoolean, quitOnFailure BACnetContextTagBoolean, writeSuccessful BACnetContextTagBoolean) *_BACnetActionCommand {
	return &_BACnetActionCommand{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, PostDelay: postDelay, QuitOnFailure: quitOnFailure, WriteSuccessful: writeSuccessful}
}

// Deprecated: use the interface for direct cast
func CastBACnetActionCommand(structType any) BACnetActionCommand {
	if casted, ok := structType.(BACnetActionCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetActionCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetActionCommand) GetTypeName() string {
	return "BACnetActionCommand"
}

func (m *_BACnetActionCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	// Optional Field (postDelay)
	if m.PostDelay != nil {
		lengthInBits += m.PostDelay.GetLengthInBits(ctx)
	}

	// Simple field (quitOnFailure)
	lengthInBits += m.QuitOnFailure.GetLengthInBits(ctx)

	// Simple field (writeSuccessful)
	lengthInBits += m.WriteSuccessful.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetActionCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetActionCommandParse(theBytes []byte) (BACnetActionCommand, error) {
	return BACnetActionCommandParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetActionCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field of BACnetActionCommand")
		default:
			deviceIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
			}
		}
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetActionCommand")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetActionCommand")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetActionCommand")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(4), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field of BACnetActionCommand")
		default:
			propertyValue = _val.(BACnetConstructedData)
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
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(5), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field of BACnetActionCommand")
		default:
			priority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	// Optional Field (postDelay) (Can be skipped, if a given expression evaluates to false)
	var postDelay BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("postDelay"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for postDelay")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(6), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'postDelay' field of BACnetActionCommand")
		default:
			postDelay = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("postDelay"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for postDelay")
			}
		}
	}

	// Simple Field (quitOnFailure)
	if pullErr := readBuffer.PullContext("quitOnFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for quitOnFailure")
	}
	_quitOnFailure, _quitOnFailureErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(7)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _quitOnFailureErr != nil {
		return nil, errors.Wrap(_quitOnFailureErr, "Error parsing 'quitOnFailure' field of BACnetActionCommand")
	}
	quitOnFailure := _quitOnFailure.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("quitOnFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for quitOnFailure")
	}

	// Simple Field (writeSuccessful)
	if pullErr := readBuffer.PullContext("writeSuccessful"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writeSuccessful")
	}
	_writeSuccessful, _writeSuccessfulErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(8)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _writeSuccessfulErr != nil {
		return nil, errors.Wrap(_writeSuccessfulErr, "Error parsing 'writeSuccessful' field of BACnetActionCommand")
	}
	writeSuccessful := _writeSuccessful.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("writeSuccessful"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writeSuccessful")
	}

	if closeErr := readBuffer.CloseContext("BACnetActionCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionCommand")
	}

	// Create the instance
	return &_BACnetActionCommand{
		DeviceIdentifier:   deviceIdentifier,
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
		PostDelay:          postDelay,
		QuitOnFailure:      quitOnFailure,
		WriteSuccessful:    writeSuccessful,
	}, nil
}

func (m *_BACnetActionCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetActionCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetActionCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionCommand")
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	if m.GetDeviceIdentifier() != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		deviceIdentifier = m.GetDeviceIdentifier()
		_deviceIdentifierErr := writeBuffer.WriteSerializable(ctx, deviceIdentifier)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
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
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(ctx, arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue BACnetConstructedData = nil
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

	// Optional Field (postDelay) (Can be skipped, if the value is null)
	var postDelay BACnetContextTagBoolean = nil
	if m.GetPostDelay() != nil {
		if pushErr := writeBuffer.PushContext("postDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for postDelay")
		}
		postDelay = m.GetPostDelay()
		_postDelayErr := writeBuffer.WriteSerializable(ctx, postDelay)
		if popErr := writeBuffer.PopContext("postDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for postDelay")
		}
		if _postDelayErr != nil {
			return errors.Wrap(_postDelayErr, "Error serializing 'postDelay' field")
		}
	}

	// Simple Field (quitOnFailure)
	if pushErr := writeBuffer.PushContext("quitOnFailure"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for quitOnFailure")
	}
	_quitOnFailureErr := writeBuffer.WriteSerializable(ctx, m.GetQuitOnFailure())
	if popErr := writeBuffer.PopContext("quitOnFailure"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for quitOnFailure")
	}
	if _quitOnFailureErr != nil {
		return errors.Wrap(_quitOnFailureErr, "Error serializing 'quitOnFailure' field")
	}

	// Simple Field (writeSuccessful)
	if pushErr := writeBuffer.PushContext("writeSuccessful"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for writeSuccessful")
	}
	_writeSuccessfulErr := writeBuffer.WriteSerializable(ctx, m.GetWriteSuccessful())
	if popErr := writeBuffer.PopContext("writeSuccessful"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for writeSuccessful")
	}
	if _writeSuccessfulErr != nil {
		return errors.Wrap(_writeSuccessfulErr, "Error serializing 'writeSuccessful' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionCommand")
	}
	return nil
}

func (m *_BACnetActionCommand) isBACnetActionCommand() bool {
	return true
}

func (m *_BACnetActionCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
