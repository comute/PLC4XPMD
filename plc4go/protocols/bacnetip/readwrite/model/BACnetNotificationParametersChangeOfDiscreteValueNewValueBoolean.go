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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean()
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(booleanValue BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if booleanValue == nil {
		panic("booleanValue of type BACnetApplicationTagBoolean for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		BooleanValue: booleanValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
