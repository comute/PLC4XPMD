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

// BACnetHostAddressIpAddress is the corresponding interface of BACnetHostAddressIpAddress
type BACnetHostAddressIpAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetHostAddress
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() BACnetContextTagOctetString
	// IsBACnetHostAddressIpAddress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostAddressIpAddress()
}

// _BACnetHostAddressIpAddress is the data-structure of this message
type _BACnetHostAddressIpAddress struct {
	BACnetHostAddressContract
	IpAddress BACnetContextTagOctetString
}

var _ BACnetHostAddressIpAddress = (*_BACnetHostAddressIpAddress)(nil)
var _ BACnetHostAddressRequirements = (*_BACnetHostAddressIpAddress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetHostAddressIpAddress) GetParent() BACnetHostAddressContract {
	return m.BACnetHostAddressContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressIpAddress) GetIpAddress() BACnetContextTagOctetString {
	return m.IpAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressIpAddress factory function for _BACnetHostAddressIpAddress
func NewBACnetHostAddressIpAddress(ipAddress BACnetContextTagOctetString, peekedTagHeader BACnetTagHeader) *_BACnetHostAddressIpAddress {
	if ipAddress == nil {
		panic("ipAddress of type BACnetContextTagOctetString for BACnetHostAddressIpAddress must not be nil")
	}
	_result := &_BACnetHostAddressIpAddress{
		BACnetHostAddressContract: NewBACnetHostAddress(peekedTagHeader),
		IpAddress:                 ipAddress,
	}
	_result.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressIpAddress(structType any) BACnetHostAddressIpAddress {
	if casted, ok := structType.(BACnetHostAddressIpAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressIpAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressIpAddress) GetTypeName() string {
	return "BACnetHostAddressIpAddress"
}

func (m *_BACnetHostAddressIpAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetHostAddressContract.(*_BACnetHostAddress).getLengthInBits(ctx))

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressIpAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetHostAddressIpAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetHostAddress) (__bACnetHostAddressIpAddress BACnetHostAddressIpAddress, err error) {
	m.BACnetHostAddressContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressIpAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressIpAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipAddress, err := ReadSimpleField[BACnetContextTagOctetString](ctx, "ipAddress", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipAddress' field"))
	}
	m.IpAddress = ipAddress

	if closeErr := readBuffer.CloseContext("BACnetHostAddressIpAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressIpAddress")
	}

	return m, nil
}

func (m *_BACnetHostAddressIpAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressIpAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressIpAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressIpAddress")
		}

		if err := WriteSimpleField[BACnetContextTagOctetString](ctx, "ipAddress", m.GetIpAddress(), WriteComplex[BACnetContextTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipAddress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressIpAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressIpAddress")
		}
		return nil
	}
	return m.BACnetHostAddressContract.(*_BACnetHostAddress).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetHostAddressIpAddress) IsBACnetHostAddressIpAddress() {}

func (m *_BACnetHostAddressIpAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
