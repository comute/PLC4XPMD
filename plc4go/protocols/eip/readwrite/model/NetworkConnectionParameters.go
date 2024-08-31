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

// NetworkConnectionParameters is the corresponding interface of NetworkConnectionParameters
type NetworkConnectionParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetConnectionSize returns ConnectionSize (property field)
	GetConnectionSize() uint16
	// GetOwner returns Owner (property field)
	GetOwner() bool
	// GetConnectionType returns ConnectionType (property field)
	GetConnectionType() uint8
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetConnectionSizeType returns ConnectionSizeType (property field)
	GetConnectionSizeType() bool
}

// NetworkConnectionParametersExactly can be used when we want exactly this type and not a type which fulfills NetworkConnectionParameters.
// This is useful for switch cases.
type NetworkConnectionParametersExactly interface {
	NetworkConnectionParameters
	isNetworkConnectionParameters() bool
}

// _NetworkConnectionParameters is the data-structure of this message
type _NetworkConnectionParameters struct {
	ConnectionSize     uint16
	Owner              bool
	ConnectionType     uint8
	Priority           uint8
	ConnectionSizeType bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *bool
	reservedField2 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkConnectionParameters) GetConnectionSize() uint16 {
	return m.ConnectionSize
}

func (m *_NetworkConnectionParameters) GetOwner() bool {
	return m.Owner
}

func (m *_NetworkConnectionParameters) GetConnectionType() uint8 {
	return m.ConnectionType
}

func (m *_NetworkConnectionParameters) GetPriority() uint8 {
	return m.Priority
}

func (m *_NetworkConnectionParameters) GetConnectionSizeType() bool {
	return m.ConnectionSizeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkConnectionParameters factory function for _NetworkConnectionParameters
func NewNetworkConnectionParameters(connectionSize uint16, owner bool, connectionType uint8, priority uint8, connectionSizeType bool) *_NetworkConnectionParameters {
	return &_NetworkConnectionParameters{ConnectionSize: connectionSize, Owner: owner, ConnectionType: connectionType, Priority: priority, ConnectionSizeType: connectionSizeType}
}

// Deprecated: use the interface for direct cast
func CastNetworkConnectionParameters(structType any) NetworkConnectionParameters {
	if casted, ok := structType.(NetworkConnectionParameters); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkConnectionParameters); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkConnectionParameters) GetTypeName() string {
	return "NetworkConnectionParameters"
}

func (m *_NetworkConnectionParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (connectionSize)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (owner)
	lengthInBits += 1

	// Simple field (connectionType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (connectionSizeType)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_NetworkConnectionParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NetworkConnectionParametersParse(ctx context.Context, theBytes []byte) (NetworkConnectionParameters, error) {
	return NetworkConnectionParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NetworkConnectionParametersParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
		return NetworkConnectionParametersParseWithBuffer(ctx, readBuffer)
	}
}

func NetworkConnectionParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkConnectionParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkConnectionParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	connectionSize, err := ReadSimpleField(ctx, "connectionSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSize' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	owner, err := ReadSimpleField(ctx, "owner", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'owner' field"))
	}

	connectionType, err := ReadSimpleField(ctx, "connectionType", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionType' field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	connectionSizeType, err := ReadSimpleField(ctx, "connectionSizeType", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSizeType' field"))
	}

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	if closeErr := readBuffer.CloseContext("NetworkConnectionParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkConnectionParameters")
	}

	// Create the instance
	return &_NetworkConnectionParameters{
		ConnectionSize:     connectionSize,
		Owner:              owner,
		ConnectionType:     connectionType,
		Priority:           priority,
		ConnectionSizeType: connectionSizeType,
		reservedField0:     reservedField0,
		reservedField1:     reservedField1,
		reservedField2:     reservedField2,
	}, nil
}

func (m *_NetworkConnectionParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkConnectionParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NetworkConnectionParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkConnectionParameters")
	}

	// Simple Field (connectionSize)
	connectionSize := uint16(m.GetConnectionSize())
	_connectionSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("connectionSize", 16, uint16((connectionSize)))
	if _connectionSizeErr != nil {
		return errors.Wrap(_connectionSizeErr, "Error serializing 'connectionSize' field")
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

	// Simple Field (owner)
	owner := bool(m.GetOwner())
	_ownerErr := /*TODO: migrate me*/ writeBuffer.WriteBit("owner", (owner))
	if _ownerErr != nil {
		return errors.Wrap(_ownerErr, "Error serializing 'owner' field")
	}

	// Simple Field (connectionType)
	connectionType := uint8(m.GetConnectionType())
	_connectionTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("connectionType", 2, uint8((connectionType)))
	if _connectionTypeErr != nil {
		return errors.Wrap(_connectionTypeErr, "Error serializing 'connectionType' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField1 != nil {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := /*TODO: migrate me*/ writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (priority)
	priority := uint8(m.GetPriority())
	_priorityErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("priority", 2, uint8((priority)))
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (connectionSizeType)
	connectionSizeType := bool(m.GetConnectionSizeType())
	_connectionSizeTypeErr := /*TODO: migrate me*/ writeBuffer.WriteBit("connectionSizeType", (connectionSizeType))
	if _connectionSizeTypeErr != nil {
		return errors.Wrap(_connectionSizeTypeErr, "Error serializing 'connectionSizeType' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField2 != nil {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField2
		}
		_err := /*TODO: migrate me*/ writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("NetworkConnectionParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkConnectionParameters")
	}
	return nil
}

func (m *_NetworkConnectionParameters) isNetworkConnectionParameters() bool {
	return true
}

func (m *_NetworkConnectionParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
