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

// BACnetConstructedDataLastAccessEvent is the corresponding interface of BACnetConstructedDataLastAccessEvent
type BACnetConstructedDataLastAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastAccessEvent returns LastAccessEvent (property field)
	GetLastAccessEvent() BACnetAccessEventTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessEventTagged
}

// BACnetConstructedDataLastAccessEventExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastAccessEvent.
// This is useful for switch cases.
type BACnetConstructedDataLastAccessEventExactly interface {
	BACnetConstructedDataLastAccessEvent
	isBACnetConstructedDataLastAccessEvent() bool
}

// _BACnetConstructedDataLastAccessEvent is the data-structure of this message
type _BACnetConstructedDataLastAccessEvent struct {
	*_BACnetConstructedData
	LastAccessEvent BACnetAccessEventTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastAccessEvent) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_ACCESS_EVENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastAccessEvent) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetLastAccessEvent() BACnetAccessEventTagged {
	return m.LastAccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastAccessEvent) GetActualValue() BACnetAccessEventTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessEventTagged(m.GetLastAccessEvent())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastAccessEvent factory function for _BACnetConstructedDataLastAccessEvent
func NewBACnetConstructedDataLastAccessEvent(lastAccessEvent BACnetAccessEventTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastAccessEvent {
	_result := &_BACnetConstructedDataLastAccessEvent{
		LastAccessEvent:        lastAccessEvent,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastAccessEvent(structType any) BACnetConstructedDataLastAccessEvent {
	if casted, ok := structType.(BACnetConstructedDataLastAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastAccessEvent) GetTypeName() string {
	return "BACnetConstructedDataLastAccessEvent"
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastAccessEvent)
	lengthInBits += m.LastAccessEvent.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastAccessEventParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastAccessEvent, error) {
	return BACnetConstructedDataLastAccessEventParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastAccessEventParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLastAccessEvent, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataLastAccessEvent, error) {
		return BACnetConstructedDataLastAccessEventParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataLastAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastAccessEvent, err := ReadSimpleField[BACnetAccessEventTagged](ctx, "lastAccessEvent", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastAccessEvent' field"))
	}

	actualValue, err := ReadVirtualField[BACnetAccessEventTagged](ctx, "actualValue", (*BACnetAccessEventTagged)(nil), lastAccessEvent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastAccessEvent{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastAccessEvent: lastAccessEvent,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastAccessEvent")
		}

		// Simple Field (lastAccessEvent)
		if pushErr := writeBuffer.PushContext("lastAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastAccessEvent")
		}
		_lastAccessEventErr := writeBuffer.WriteSerializable(ctx, m.GetLastAccessEvent())
		if popErr := writeBuffer.PopContext("lastAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastAccessEvent")
		}
		if _lastAccessEventErr != nil {
			return errors.Wrap(_lastAccessEventErr, "Error serializing 'lastAccessEvent' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastAccessEvent) isBACnetConstructedDataLastAccessEvent() bool {
	return true
}

func (m *_BACnetConstructedDataLastAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
