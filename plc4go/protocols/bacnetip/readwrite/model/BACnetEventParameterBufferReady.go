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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventParameterBufferReady is the corresponding interface of BACnetEventParameterBufferReady
type BACnetEventParameterBufferReady interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetNotificationThreshold returns NotificationThreshold (property field)
	GetNotificationThreshold() BACnetContextTagUnsignedInteger
	// GetPreviousNotificationCount returns PreviousNotificationCount (property field)
	GetPreviousNotificationCount() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterBufferReadyExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterBufferReady.
// This is useful for switch cases.
type BACnetEventParameterBufferReadyExactly interface {
	BACnetEventParameterBufferReady
	isBACnetEventParameterBufferReady() bool
}

// _BACnetEventParameterBufferReady is the data-structure of this message
type _BACnetEventParameterBufferReady struct {
	*_BACnetEventParameter
	OpeningTag                BACnetOpeningTag
	NotificationThreshold     BACnetContextTagUnsignedInteger
	PreviousNotificationCount BACnetContextTagUnsignedInteger
	ClosingTag                BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterBufferReady) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterBufferReady) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterBufferReady) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterBufferReady) GetNotificationThreshold() BACnetContextTagUnsignedInteger {
	return m.NotificationThreshold
}

func (m *_BACnetEventParameterBufferReady) GetPreviousNotificationCount() BACnetContextTagUnsignedInteger {
	return m.PreviousNotificationCount
}

func (m *_BACnetEventParameterBufferReady) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterBufferReady factory function for _BACnetEventParameterBufferReady
func NewBACnetEventParameterBufferReady(openingTag BACnetOpeningTag, notificationThreshold BACnetContextTagUnsignedInteger, previousNotificationCount BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterBufferReady {
	_result := &_BACnetEventParameterBufferReady{
		OpeningTag:                openingTag,
		NotificationThreshold:     notificationThreshold,
		PreviousNotificationCount: previousNotificationCount,
		ClosingTag:                closingTag,
		_BACnetEventParameter:     NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterBufferReady(structType interface{}) BACnetEventParameterBufferReady {
	if casted, ok := structType.(BACnetEventParameterBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterBufferReady); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterBufferReady) GetTypeName() string {
	return "BACnetEventParameterBufferReady"
}

func (m *_BACnetEventParameterBufferReady) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (notificationThreshold)
	lengthInBits += m.NotificationThreshold.GetLengthInBits(ctx)

	// Simple field (previousNotificationCount)
	lengthInBits += m.PreviousNotificationCount.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterBufferReady) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterBufferReadyParse(theBytes []byte) (BACnetEventParameterBufferReady, error) {
	return BACnetEventParameterBufferReadyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterBufferReadyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterBufferReady, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(10)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterBufferReady")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (notificationThreshold)
	if pullErr := readBuffer.PullContext("notificationThreshold"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notificationThreshold")
	}
	_notificationThreshold, _notificationThresholdErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _notificationThresholdErr != nil {
		return nil, errors.Wrap(_notificationThresholdErr, "Error parsing 'notificationThreshold' field of BACnetEventParameterBufferReady")
	}
	notificationThreshold := _notificationThreshold.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("notificationThreshold"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notificationThreshold")
	}

	// Simple Field (previousNotificationCount)
	if pullErr := readBuffer.PullContext("previousNotificationCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for previousNotificationCount")
	}
	_previousNotificationCount, _previousNotificationCountErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _previousNotificationCountErr != nil {
		return nil, errors.Wrap(_previousNotificationCountErr, "Error parsing 'previousNotificationCount' field of BACnetEventParameterBufferReady")
	}
	previousNotificationCount := _previousNotificationCount.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("previousNotificationCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for previousNotificationCount")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(10)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterBufferReady")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterBufferReady")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterBufferReady{
		_BACnetEventParameter:     &_BACnetEventParameter{},
		OpeningTag:                openingTag,
		NotificationThreshold:     notificationThreshold,
		PreviousNotificationCount: previousNotificationCount,
		ClosingTag:                closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterBufferReady) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterBufferReady) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterBufferReady")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (notificationThreshold)
		if pushErr := writeBuffer.PushContext("notificationThreshold"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notificationThreshold")
		}
		_notificationThresholdErr := writeBuffer.WriteSerializable(ctx, m.GetNotificationThreshold())
		if popErr := writeBuffer.PopContext("notificationThreshold"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notificationThreshold")
		}
		if _notificationThresholdErr != nil {
			return errors.Wrap(_notificationThresholdErr, "Error serializing 'notificationThreshold' field")
		}

		// Simple Field (previousNotificationCount)
		if pushErr := writeBuffer.PushContext("previousNotificationCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for previousNotificationCount")
		}
		_previousNotificationCountErr := writeBuffer.WriteSerializable(ctx, m.GetPreviousNotificationCount())
		if popErr := writeBuffer.PopContext("previousNotificationCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for previousNotificationCount")
		}
		if _previousNotificationCountErr != nil {
			return errors.Wrap(_previousNotificationCountErr, "Error serializing 'previousNotificationCount' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterBufferReady")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterBufferReady) isBACnetEventParameterBufferReady() bool {
	return true
}

func (m *_BACnetEventParameterBufferReady) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
