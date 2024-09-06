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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() ListOfCovNotificationsList
	// IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple()
}

// _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	TimeRemaining               BACnetContextTagUnsignedInteger
	Timestamp                   BACnetTimeStampEnclosed
	ListOfCovNotifications      ListOfCovNotificationsList
}

var _ BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetListOfCovNotifications() ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple factory function for _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, timestamp BACnetTimeStampEnclosed, listOfCovNotifications ListOfCovNotificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if listOfCovNotifications == nil {
		panic("listOfCovNotifications of type ListOfCovNotificationsList for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:            initiatingDeviceIdentifier,
		TimeRemaining:                         timeRemaining,
		Timestamp:                             timestamp,
		ListOfCovNotifications:                listOfCovNotifications,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(structType any) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += m.Timestamp.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	var timestamp BACnetTimeStampEnclosed
	_timestamp, err := ReadOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	if _timestamp != nil {
		timestamp = *_timestamp
		m.Timestamp = timestamp
	}

	listOfCovNotifications, err := ReadSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", ReadComplex[ListOfCovNotificationsList](ListOfCovNotificationsListParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovNotifications' field"))
	}
	m.ListOfCovNotifications = listOfCovNotifications

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeRemaining' field")
		}

		if err := WriteOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", GetRef(m.GetTimestamp()), WriteComplex[BACnetTimeStampEnclosed](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", m.GetListOfCovNotifications(), WriteComplex[ListOfCovNotificationsList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfCovNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
