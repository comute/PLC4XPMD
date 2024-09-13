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

// ConnectionRequestInformationTunnelConnection is the corresponding interface of ConnectionRequestInformationTunnelConnection
type ConnectionRequestInformationTunnelConnection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ConnectionRequestInformation
	// GetKnxLayer returns KnxLayer (property field)
	GetKnxLayer() KnxLayer
	// IsConnectionRequestInformationTunnelConnection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequestInformationTunnelConnection()
}

// _ConnectionRequestInformationTunnelConnection is the data-structure of this message
type _ConnectionRequestInformationTunnelConnection struct {
	ConnectionRequestInformationContract
	KnxLayer KnxLayer
	// Reserved Fields
	reservedField0 *uint8
}

var _ ConnectionRequestInformationTunnelConnection = (*_ConnectionRequestInformationTunnelConnection)(nil)
var _ ConnectionRequestInformationRequirements = (*_ConnectionRequestInformationTunnelConnection)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetParent() ConnectionRequestInformationContract {
	return m.ConnectionRequestInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionRequestInformationTunnelConnection) GetKnxLayer() KnxLayer {
	return m.KnxLayer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectionRequestInformationTunnelConnection factory function for _ConnectionRequestInformationTunnelConnection
func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer) *_ConnectionRequestInformationTunnelConnection {
	_result := &_ConnectionRequestInformationTunnelConnection{
		ConnectionRequestInformationContract: NewConnectionRequestInformation(),
		KnxLayer:                             knxLayer,
	}
	_result.ConnectionRequestInformationContract.(*_ConnectionRequestInformation)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformationTunnelConnection(structType any) ConnectionRequestInformationTunnelConnection {
	if casted, ok := structType.(ConnectionRequestInformationTunnelConnection); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationTunnelConnection); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformationTunnelConnection) GetTypeName() string {
	return "ConnectionRequestInformationTunnelConnection"
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).getLengthInBits(ctx))

	// Simple field (knxLayer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ConnectionRequestInformationTunnelConnection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionRequestInformationTunnelConnection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ConnectionRequestInformation) (__connectionRequestInformationTunnelConnection ConnectionRequestInformationTunnelConnection, err error) {
	m.ConnectionRequestInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationTunnelConnection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationTunnelConnection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	knxLayer, err := ReadEnumField[KnxLayer](ctx, "knxLayer", "KnxLayer", ReadEnum(KnxLayerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'knxLayer' field"))
	}
	m.KnxLayer = knxLayer

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationTunnelConnection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationTunnelConnection")
	}

	return m, nil
}

func (m *_ConnectionRequestInformationTunnelConnection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequestInformationTunnelConnection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationTunnelConnection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationTunnelConnection")
		}

		if err := WriteSimpleEnumField[KnxLayer](ctx, "knxLayer", "KnxLayer", m.GetKnxLayer(), WriteEnum[KnxLayer, uint8](KnxLayer.GetValue, KnxLayer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'knxLayer' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationTunnelConnection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationTunnelConnection")
		}
		return nil
	}
	return m.ConnectionRequestInformationContract.(*_ConnectionRequestInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequestInformationTunnelConnection) IsConnectionRequestInformationTunnelConnection() {
}

func (m *_ConnectionRequestInformationTunnelConnection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
