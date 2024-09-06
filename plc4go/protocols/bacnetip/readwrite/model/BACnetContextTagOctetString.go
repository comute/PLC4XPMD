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

// BACnetContextTagOctetString is the corresponding interface of BACnetContextTagOctetString
type BACnetContextTagOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadOctetString
	// IsBACnetContextTagOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagOctetString()
}

// _BACnetContextTagOctetString is the data-structure of this message
type _BACnetContextTagOctetString struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadOctetString
}

var _ BACnetContextTagOctetString = (*_BACnetContextTagOctetString)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagOctetString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagOctetString) GetDataType() BACnetDataType {
	return BACnetDataType_OCTET_STRING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagOctetString) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagOctetString) GetPayload() BACnetTagPayloadOctetString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagOctetString factory function for _BACnetContextTagOctetString
func NewBACnetContextTagOctetString(payload BACnetTagPayloadOctetString, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagOctetString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadOctetString for BACnetContextTagOctetString must not be nil")
	}
	_result := &_BACnetContextTagOctetString{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagOctetString(structType any) BACnetContextTagOctetString {
	if casted, ok := structType.(BACnetContextTagOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagOctetString) GetTypeName() string {
	return "BACnetContextTagOctetString"
}

func (m *_BACnetContextTagOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagOctetString BACnetContextTagOctetString, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadOctetString](ctx, "payload", ReadComplex[BACnetTagPayloadOctetString](BACnetTagPayloadOctetStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetContextTagOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagOctetString")
	}

	return m, nil
}

func (m *_BACnetContextTagOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagOctetString")
		}

		if err := WriteSimpleField[BACnetTagPayloadOctetString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagOctetString")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagOctetString) IsBACnetContextTagOctetString() {}

func (m *_BACnetContextTagOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
