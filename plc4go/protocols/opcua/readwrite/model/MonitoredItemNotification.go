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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoredItemNotification is the corresponding interface of MonitoredItemNotification
type MonitoredItemNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetValue returns Value (property field)
	GetValue() DataValue
}

// MonitoredItemNotificationExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemNotification.
// This is useful for switch cases.
type MonitoredItemNotificationExactly interface {
	MonitoredItemNotification
	isMonitoredItemNotification() bool
}

// _MonitoredItemNotification is the data-structure of this message
type _MonitoredItemNotification struct {
	*_ExtensionObjectDefinition
	ClientHandle uint32
	Value        DataValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemNotification) GetIdentifier() string {
	return "808"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemNotification) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoredItemNotification) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemNotification) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_MonitoredItemNotification) GetValue() DataValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemNotification factory function for _MonitoredItemNotification
func NewMonitoredItemNotification(clientHandle uint32, value DataValue) *_MonitoredItemNotification {
	_result := &_MonitoredItemNotification{
		ClientHandle:               clientHandle,
		Value:                      value,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemNotification(structType any) MonitoredItemNotification {
	if casted, ok := structType.(MonitoredItemNotification); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemNotification); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemNotification) GetTypeName() string {
	return "MonitoredItemNotification"
}

func (m *_MonitoredItemNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredItemNotificationParse(ctx context.Context, theBytes []byte, identifier string) (MonitoredItemNotification, error) {
	return MonitoredItemNotificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoredItemNotificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoredItemNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredItemNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (clientHandle)
	_clientHandle, _clientHandleErr := readBuffer.ReadUint32("clientHandle", 32)
	if _clientHandleErr != nil {
		return nil, errors.Wrap(_clientHandleErr, "Error parsing 'clientHandle' field of MonitoredItemNotification")
	}
	clientHandle := _clientHandle

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := DataValueParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of MonitoredItemNotification")
	}
	value := _value.(DataValue)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("MonitoredItemNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemNotification")
	}

	// Create a partially initialized instance
	_child := &_MonitoredItemNotification{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ClientHandle:               clientHandle,
		Value:                      value,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredItemNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemNotification")
		}

		// Simple Field (clientHandle)
		clientHandle := uint32(m.GetClientHandle())
		_clientHandleErr := writeBuffer.WriteUint32("clientHandle", 32, uint32((clientHandle)))
		if _clientHandleErr != nil {
			return errors.Wrap(_clientHandleErr, "Error serializing 'clientHandle' field")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemNotification")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemNotification) isMonitoredItemNotification() bool {
	return true
}

func (m *_MonitoredItemNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
