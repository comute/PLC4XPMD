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

// SubscribeCOVPropertyMultipleError is the corresponding interface of SubscribeCOVPropertyMultipleError
type SubscribeCOVPropertyMultipleError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedSubscription returns FirstFailedSubscription (property field)
	GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

// SubscribeCOVPropertyMultipleErrorExactly can be used when we want exactly this type and not a type which fulfills SubscribeCOVPropertyMultipleError.
// This is useful for switch cases.
type SubscribeCOVPropertyMultipleErrorExactly interface {
	SubscribeCOVPropertyMultipleError
	isSubscribeCOVPropertyMultipleError() bool
}

// _SubscribeCOVPropertyMultipleError is the data-structure of this message
type _SubscribeCOVPropertyMultipleError struct {
	*_BACnetError
	ErrorType               ErrorEnclosed
	FirstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribeCOVPropertyMultipleError) InitializeParent(parent BACnetError) {}

func (m *_SubscribeCOVPropertyMultipleError) GetParent() BACnetError {
	return m._BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleError) GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return m.FirstFailedSubscription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleError factory function for _SubscribeCOVPropertyMultipleError
func NewSubscribeCOVPropertyMultipleError(errorType ErrorEnclosed, firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) *_SubscribeCOVPropertyMultipleError {
	_result := &_SubscribeCOVPropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
		_BACnetError:            NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleError(structType any) SubscribeCOVPropertyMultipleError {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleError); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleError); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleError) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleError"
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedSubscription)
	lengthInBits += m.FirstFailedSubscription.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubscribeCOVPropertyMultipleErrorParse(ctx context.Context, theBytes []byte, errorChoice BACnetConfirmedServiceChoice) (SubscribeCOVPropertyMultipleError, error) {
	return SubscribeCOVPropertyMultipleErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), errorChoice)
}

func SubscribeCOVPropertyMultipleErrorParseWithBufferProducer(errorChoice BACnetConfirmedServiceChoice) func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleError, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleError, error) {
		return SubscribeCOVPropertyMultipleErrorParseWithBuffer(ctx, readBuffer, errorChoice)
	}
}

func SubscribeCOVPropertyMultipleErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (SubscribeCOVPropertyMultipleError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}

	firstFailedSubscription, err := ReadSimpleField[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](ctx, "firstFailedSubscription", ReadComplex[SubscribeCOVPropertyMultipleErrorFirstFailedSubscription](SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedSubscription' field"))
	}

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleError")
	}

	// Create a partially initialized instance
	_child := &_SubscribeCOVPropertyMultipleError{
		_BACnetError:            &_BACnetError{},
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_SubscribeCOVPropertyMultipleError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribeCOVPropertyMultipleError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleError")
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorType")
		}
		_errorTypeErr := writeBuffer.WriteSerializable(ctx, m.GetErrorType())
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorType")
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (firstFailedSubscription)
		if pushErr := writeBuffer.PushContext("firstFailedSubscription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for firstFailedSubscription")
		}
		_firstFailedSubscriptionErr := writeBuffer.WriteSerializable(ctx, m.GetFirstFailedSubscription())
		if popErr := writeBuffer.PopContext("firstFailedSubscription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for firstFailedSubscription")
		}
		if _firstFailedSubscriptionErr != nil {
			return errors.Wrap(_firstFailedSubscriptionErr, "Error serializing 'firstFailedSubscription' field")
		}

		if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleError")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscribeCOVPropertyMultipleError) isSubscribeCOVPropertyMultipleError() bool {
	return true
}

func (m *_SubscribeCOVPropertyMultipleError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
