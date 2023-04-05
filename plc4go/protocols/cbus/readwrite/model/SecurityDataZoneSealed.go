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

// SecurityDataZoneSealed is the corresponding interface of SecurityDataZoneSealed
type SecurityDataZoneSealed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataZoneSealedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneSealed.
// This is useful for switch cases.
type SecurityDataZoneSealedExactly interface {
	SecurityDataZoneSealed
	isSecurityDataZoneSealed() bool
}

// _SecurityDataZoneSealed is the data-structure of this message
type _SecurityDataZoneSealed struct {
	*_SecurityData
	ZoneNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneSealed) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataZoneSealed) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneSealed) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataZoneSealed factory function for _SecurityDataZoneSealed
func NewSecurityDataZoneSealed(zoneNumber uint8, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataZoneSealed {
	_result := &_SecurityDataZoneSealed{
		ZoneNumber:    zoneNumber,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneSealed(structType interface{}) SecurityDataZoneSealed {
	if casted, ok := structType.(SecurityDataZoneSealed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneSealed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneSealed) GetTypeName() string {
	return "SecurityDataZoneSealed"
}

func (m *_SecurityDataZoneSealed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneSealed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataZoneSealedParse(theBytes []byte) (SecurityDataZoneSealed, error) {
	return SecurityDataZoneSealedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataZoneSealedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataZoneSealed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneSealed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneSealed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
	_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataZoneSealed")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneSealed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneSealed")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataZoneSealed{
		_SecurityData: &_SecurityData{},
		ZoneNumber:    zoneNumber,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataZoneSealed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneSealed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneSealed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneSealed")
		}

		// Simple Field (zoneNumber)
		zoneNumber := uint8(m.GetZoneNumber())
		_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
		if _zoneNumberErr != nil {
			return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneSealed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneSealed")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneSealed) isSecurityDataZoneSealed() bool {
	return true
}

func (m *_SecurityDataZoneSealed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
