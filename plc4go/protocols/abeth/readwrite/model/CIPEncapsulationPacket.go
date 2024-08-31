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

// CIPEncapsulationPacket is the corresponding interface of CIPEncapsulationPacket
type CIPEncapsulationPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() uint16
	// GetSessionHandle returns SessionHandle (property field)
	GetSessionHandle() uint32
	// GetStatus returns Status (property field)
	GetStatus() uint32
	// GetSenderContext returns SenderContext (property field)
	GetSenderContext() []uint8
	// GetOptions returns Options (property field)
	GetOptions() uint32
}

// CIPEncapsulationPacketExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationPacket.
// This is useful for switch cases.
type CIPEncapsulationPacketExactly interface {
	CIPEncapsulationPacket
	isCIPEncapsulationPacket() bool
}

// _CIPEncapsulationPacket is the data-structure of this message
type _CIPEncapsulationPacket struct {
	_CIPEncapsulationPacketChildRequirements
	SessionHandle uint32
	Status        uint32
	SenderContext []uint8
	Options       uint32
	// Reserved Fields
	reservedField0 *uint32
}

type _CIPEncapsulationPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() uint16
}

type CIPEncapsulationPacketParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CIPEncapsulationPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type CIPEncapsulationPacketChild interface {
	utils.Serializable
	InitializeParent(parent CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32)
	GetParent() *CIPEncapsulationPacket

	GetTypeName() string
	CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationPacket) GetSessionHandle() uint32 {
	return m.SessionHandle
}

func (m *_CIPEncapsulationPacket) GetStatus() uint32 {
	return m.Status
}

func (m *_CIPEncapsulationPacket) GetSenderContext() []uint8 {
	return m.SenderContext
}

func (m *_CIPEncapsulationPacket) GetOptions() uint32 {
	return m.Options
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationPacket factory function for _CIPEncapsulationPacket
func NewCIPEncapsulationPacket(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationPacket {
	return &_CIPEncapsulationPacket{SessionHandle: sessionHandle, Status: status, SenderContext: senderContext, Options: options}
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationPacket(structType any) CIPEncapsulationPacket {
	if casted, ok := structType.(CIPEncapsulationPacket); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationPacket); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationPacket) GetTypeName() string {
	return "CIPEncapsulationPacket"
}

func (m *_CIPEncapsulationPacket) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandType)
	lengthInBits += 16

	// Implicit Field (packetLen)
	lengthInBits += 16

	// Simple field (sessionHandle)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += 32

	// Array field
	if len(m.SenderContext) > 0 {
		lengthInBits += 8 * uint16(len(m.SenderContext))
	}

	// Simple field (options)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CIPEncapsulationPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPEncapsulationPacketParse(ctx context.Context, theBytes []byte) (CIPEncapsulationPacket, error) {
	return CIPEncapsulationPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func CIPEncapsulationPacketParseWithBufferProducer[T CIPEncapsulationPacket]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := CIPEncapsulationPacketParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func CIPEncapsulationPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPEncapsulationPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	commandType, err := ReadDiscriminatorField[uint16](ctx, "commandType", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}

	packetLen, err := ReadImplicitField[uint16](ctx, "packetLen", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'packetLen' field"))
	}
	_ = packetLen

	sessionHandle, err := ReadSimpleField(ctx, "sessionHandle", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionHandle' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}

	senderContext, err := ReadCountArrayField[uint8](ctx, "senderContext", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'senderContext' field"))
	}

	options, err := ReadSimpleField(ctx, "options", ReadUnsignedInt(readBuffer, uint8(32)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'options' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(32)), uint32(0x00000000), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CIPEncapsulationPacketChildSerializeRequirement interface {
		CIPEncapsulationPacket
		InitializeParent(CIPEncapsulationPacket, uint32, uint32, []uint8, uint32)
		GetParent() CIPEncapsulationPacket
	}
	var _childTemp any
	var _child CIPEncapsulationPacketChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == 0x0101: // CIPEncapsulationConnectionRequest
		_childTemp, typeSwitchError = CIPEncapsulationConnectionRequestParseWithBuffer(ctx, readBuffer)
	case commandType == 0x0201: // CIPEncapsulationConnectionResponse
		_childTemp, typeSwitchError = CIPEncapsulationConnectionResponseParseWithBuffer(ctx, readBuffer)
	case commandType == 0x0107: // CIPEncapsulationReadRequest
		_childTemp, typeSwitchError = CIPEncapsulationReadRequestParseWithBuffer(ctx, readBuffer)
	case commandType == 0x0207: // CIPEncapsulationReadResponse
		_childTemp, typeSwitchError = CIPEncapsulationReadResponseParseWithBuffer(ctx, readBuffer, packetLen)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CIPEncapsulationPacket")
	}
	_child = _childTemp.(CIPEncapsulationPacketChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CIPEncapsulationPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationPacket")
	}

	// Finish initializing
	_child.InitializeParent(_child, sessionHandle, status, senderContext, options)
	_child.GetParent().(*_CIPEncapsulationPacket).reservedField0 = reservedField0
	return _child, nil
}

func (pm *_CIPEncapsulationPacket) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CIPEncapsulationPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPEncapsulationPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationPacket")
	}

	// Discriminator Field (commandType) (Used as input to a switch field)
	commandType := uint16(child.GetCommandType())
	_commandTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("commandType", 16, uint16((commandType)))

	if _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Implicit Field (packetLen) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	packetLen := uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(28)))
	_packetLenErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("packetLen", 16, uint16((packetLen)))
	if _packetLenErr != nil {
		return errors.Wrap(_packetLenErr, "Error serializing 'packetLen' field")
	}

	// Simple Field (sessionHandle)
	sessionHandle := uint32(m.GetSessionHandle())
	_sessionHandleErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("sessionHandle", 32, uint32((sessionHandle)))
	if _sessionHandleErr != nil {
		return errors.Wrap(_sessionHandleErr, "Error serializing 'sessionHandle' field")
	}

	// Simple Field (status)
	status := uint32(m.GetStatus())
	_statusErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("status", 32, uint32((status)))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "senderContext", m.GetSenderContext(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'senderContext' field")
	}

	// Simple Field (options)
	options := uint32(m.GetOptions())
	_optionsErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("options", 32, uint32((options)))
	if _optionsErr != nil {
		return errors.Wrap(_optionsErr, "Error serializing 'options' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint32 = uint32(0x00000000)
		if pm.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint32(0x00000000),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *pm.reservedField0
		}
		_err := /*TODO: migrate me*/ writeBuffer.WriteUint32("reserved", 32, uint32(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CIPEncapsulationPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPEncapsulationPacket")
	}
	return nil
}

func (m *_CIPEncapsulationPacket) isCIPEncapsulationPacket() bool {
	return true
}

func (m *_CIPEncapsulationPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
