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

// BACnetChannelValueOctetString is the corresponding interface of BACnetChannelValueOctetString
type BACnetChannelValueOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
}

// BACnetChannelValueOctetStringExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueOctetString.
// This is useful for switch cases.
type BACnetChannelValueOctetStringExactly interface {
	BACnetChannelValueOctetString
	isBACnetChannelValueOctetString() bool
}

// _BACnetChannelValueOctetString is the data-structure of this message
type _BACnetChannelValueOctetString struct {
	*_BACnetChannelValue
	OctetStringValue BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueOctetString) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueOctetString) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueOctetString factory function for _BACnetChannelValueOctetString
func NewBACnetChannelValueOctetString(octetStringValue BACnetApplicationTagOctetString, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueOctetString {
	_result := &_BACnetChannelValueOctetString{
		OctetStringValue:    octetStringValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueOctetString(structType any) BACnetChannelValueOctetString {
	if casted, ok := structType.(BACnetChannelValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueOctetString) GetTypeName() string {
	return "BACnetChannelValueOctetString"
}

func (m *_BACnetChannelValueOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueOctetStringParse(ctx context.Context, theBytes []byte) (BACnetChannelValueOctetString, error) {
	return BACnetChannelValueOctetStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueOctetStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueOctetString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueOctetString, error) {
		return BACnetChannelValueOctetStringParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetChannelValueOctetStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octetStringValue, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueOctetString")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueOctetString{
		_BACnetChannelValue: &_BACnetChannelValue{},
		OctetStringValue:    octetStringValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueOctetString")
		}

		// Simple Field (octetStringValue)
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for octetStringValue")
		}
		_octetStringValueErr := writeBuffer.WriteSerializable(ctx, m.GetOctetStringValue())
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for octetStringValue")
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueOctetString")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueOctetString) isBACnetChannelValueOctetString() bool {
	return true
}

func (m *_BACnetChannelValueOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
