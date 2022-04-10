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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification struct {
	*BACnetUnconfirmedServiceRequest
	SubscriberProcessIdentifier *BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  *BACnetContextTagObjectIdentifier
	MonitoredObjectIdentifier   *BACnetContextTagObjectIdentifier
	LifetimeInSeconds           *BACnetContextTagUnsignedInteger
	ListOfValues                *BACnetPropertyValues

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification interface {
	IBACnetUnconfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() *BACnetPropertyValues
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
func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetServiceChoice() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetListOfValues() *BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification factory function for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(subscriberProcessIdentifier *BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier *BACnetContextTagObjectIdentifier, monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, lifetimeInSeconds *BACnetContextTagUnsignedInteger, listOfValues *BACnetPropertyValues, len uint16) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification {
	_result := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification{
		SubscriberProcessIdentifier:     subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:      initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:       monitoredObjectIdentifier,
		LifetimeInSeconds:               lifetimeInSeconds,
		ListOfValues:                    listOfValues,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotification(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits()

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := CastBACnetContextTagUnsignedInteger(_subscriberProcessIdentifier)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field")
	}
	initiatingDeviceIdentifier := CastBACnetContextTagObjectIdentifier(_initiatingDeviceIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lifetimeInSeconds)
	if pullErr := readBuffer.PullContext("lifetimeInSeconds"); pullErr != nil {
		return nil, pullErr
	}
	_lifetimeInSeconds, _lifetimeInSecondsErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lifetimeInSecondsErr != nil {
		return nil, errors.Wrap(_lifetimeInSecondsErr, "Error parsing 'lifetimeInSeconds' field")
	}
	lifetimeInSeconds := CastBACnetContextTagUnsignedInteger(_lifetimeInSeconds)
	if closeErr := readBuffer.CloseContext("lifetimeInSeconds"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParse(readBuffer, uint8(uint8(4)), BACnetObjectType(monitoredObjectIdentifier.GetObjectType()))
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field")
	}
	listOfValues := CastBACnetPropertyValues(_listOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification{
		SubscriberProcessIdentifier:     CastBACnetContextTagUnsignedInteger(subscriberProcessIdentifier),
		InitiatingDeviceIdentifier:      CastBACnetContextTagObjectIdentifier(initiatingDeviceIdentifier),
		MonitoredObjectIdentifier:       CastBACnetContextTagObjectIdentifier(monitoredObjectIdentifier),
		LifetimeInSeconds:               CastBACnetContextTagUnsignedInteger(lifetimeInSeconds),
		ListOfValues:                    CastBACnetPropertyValues(listOfValues),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_subscriberProcessIdentifierErr := m.SubscriberProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_initiatingDeviceIdentifierErr := m.InitiatingDeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return popErr
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Simple Field (lifetimeInSeconds)
		if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
			return pushErr
		}
		_lifetimeInSecondsErr := m.LifetimeInSeconds.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
			return popErr
		}
		if _lifetimeInSecondsErr != nil {
			return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
		}

		// Simple Field (listOfValues)
		if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
			return pushErr
		}
		_listOfValuesErr := m.ListOfValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
			return popErr
		}
		if _listOfValuesErr != nil {
			return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
