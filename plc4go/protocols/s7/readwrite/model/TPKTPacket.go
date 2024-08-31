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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const TPKTPacket_PROTOCOLID uint8 = 0x03

// TPKTPacket is the corresponding interface of TPKTPacket
type TPKTPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPayload returns Payload (property field)
	GetPayload() COTPPacket
}

// TPKTPacketExactly can be used when we want exactly this type and not a type which fulfills TPKTPacket.
// This is useful for switch cases.
type TPKTPacketExactly interface {
	TPKTPacket
	isTPKTPacket() bool
}

// _TPKTPacket is the data-structure of this message
type _TPKTPacket struct {
	Payload COTPPacket
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TPKTPacket) GetPayload() COTPPacket {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_TPKTPacket) GetProtocolId() uint8 {
	return TPKTPacket_PROTOCOLID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTPKTPacket factory function for _TPKTPacket
func NewTPKTPacket(payload COTPPacket) *_TPKTPacket {
	return &_TPKTPacket{Payload: payload}
}

// Deprecated: use the interface for direct cast
func CastTPKTPacket(structType any) TPKTPacket {
	if casted, ok := structType.(TPKTPacket); ok {
		return casted
	}
	if casted, ok := structType.(*TPKTPacket); ok {
		return *casted
	}
	return nil
}

func (m *_TPKTPacket) GetTypeName() string {
	return "TPKTPacket"
}

func (m *_TPKTPacket) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Implicit Field (len)
	lengthInBits += 16

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_TPKTPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TPKTPacketParse(ctx context.Context, theBytes []byte) (TPKTPacket, error) {
	return TPKTPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func TPKTPacketParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
		return TPKTPacketParseWithBuffer(ctx, readBuffer)
	}
}

func TPKTPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TPKTPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TPKTPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TPKTPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolId, err := ReadConstField[uint8](ctx, "protocolId", ReadUnsignedByte(readBuffer, uint8(8)), TPKTPacket_PROTOCOLID, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolId' field"))
	}
	_ = protocolId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	len, err := ReadImplicitField[uint16](ctx, "len", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'len' field"))
	}
	_ = len

	payload, err := ReadSimpleField[COTPPacket](ctx, "payload", ReadComplex[COTPPacket](COTPPacketParseWithBufferProducer[COTPPacket]((uint16)(uint16(len)-uint16(uint16(4)))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}

	if closeErr := readBuffer.CloseContext("TPKTPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TPKTPacket")
	}

	// Create the instance
	return &_TPKTPacket{
		Payload:        payload,
		reservedField0: reservedField0,
	}, nil
}

func (m *_TPKTPacket) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TPKTPacket) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TPKTPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TPKTPacket")
	}

	// Const Field (protocolId)
	_protocolIdErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("protocolId", 8, uint8(0x03))
	if _protocolIdErr != nil {
		return errors.Wrap(_protocolIdErr, "Error serializing 'protocolId' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len := uint16(uint16(m.GetPayload().GetLengthInBytes(ctx)) + uint16(uint16(4)))
	_lenErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("len", 16, uint16((len)))
	if _lenErr != nil {
		return errors.Wrap(_lenErr, "Error serializing 'len' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("TPKTPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TPKTPacket")
	}
	return nil
}

func (m *_TPKTPacket) isTPKTPacket() bool {
	return true
}

func (m *_TPKTPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
