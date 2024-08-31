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

// SecurityArmCode is the corresponding interface of SecurityArmCode
type SecurityArmCode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCode returns Code (property field)
	GetCode() uint8
	// GetIsDisarmed returns IsDisarmed (virtual field)
	GetIsDisarmed() bool
	// GetIsFullyArmed returns IsFullyArmed (virtual field)
	GetIsFullyArmed() bool
	// GetIsPartiallyArmed returns IsPartiallyArmed (virtual field)
	GetIsPartiallyArmed() bool
	// GetIsArmSubtype returns IsArmSubtype (virtual field)
	GetIsArmSubtype() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// SecurityArmCodeExactly can be used when we want exactly this type and not a type which fulfills SecurityArmCode.
// This is useful for switch cases.
type SecurityArmCodeExactly interface {
	SecurityArmCode
	isSecurityArmCode() bool
}

// _SecurityArmCode is the data-structure of this message
type _SecurityArmCode struct {
	Code uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityArmCode) GetCode() uint8 {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityArmCode) GetIsDisarmed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x00)))
}

func (m *_SecurityArmCode) GetIsFullyArmed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x01)))
}

func (m *_SecurityArmCode) GetIsPartiallyArmed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) == (0x02)))
}

func (m *_SecurityArmCode) GetIsArmSubtype() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetCode()) >= (0x03))) && bool(bool((m.GetCode()) <= (0x7F))))
}

func (m *_SecurityArmCode) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetCode()) > (0x7F)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityArmCode factory function for _SecurityArmCode
func NewSecurityArmCode(code uint8) *_SecurityArmCode {
	return &_SecurityArmCode{Code: code}
}

// Deprecated: use the interface for direct cast
func CastSecurityArmCode(structType any) SecurityArmCode {
	if casted, ok := structType.(SecurityArmCode); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityArmCode); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityArmCode) GetTypeName() string {
	return "SecurityArmCode"
}

func (m *_SecurityArmCode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (code)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityArmCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityArmCodeParse(ctx context.Context, theBytes []byte) (SecurityArmCode, error) {
	return SecurityArmCodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityArmCodeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityArmCode, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityArmCode, error) {
		return SecurityArmCodeParseWithBuffer(ctx, readBuffer)
	}
}

func SecurityArmCodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityArmCode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityArmCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityArmCode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	code, err := ReadSimpleField(ctx, "code", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}

	isDisarmed, err := ReadVirtualField[bool](ctx, "isDisarmed", (*bool)(nil), bool((code) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDisarmed' field"))
	}
	_ = isDisarmed

	isFullyArmed, err := ReadVirtualField[bool](ctx, "isFullyArmed", (*bool)(nil), bool((code) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isFullyArmed' field"))
	}
	_ = isFullyArmed

	isPartiallyArmed, err := ReadVirtualField[bool](ctx, "isPartiallyArmed", (*bool)(nil), bool((code) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPartiallyArmed' field"))
	}
	_ = isPartiallyArmed

	isArmSubtype, err := ReadVirtualField[bool](ctx, "isArmSubtype", (*bool)(nil), bool(bool((code) >= (0x03))) && bool(bool((code) <= (0x7F))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmSubtype' field"))
	}
	_ = isArmSubtype

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool((code) > (0x7F)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	if closeErr := readBuffer.CloseContext("SecurityArmCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityArmCode")
	}

	// Create the instance
	return &_SecurityArmCode{
		Code: code,
	}, nil
}

func (m *_SecurityArmCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityArmCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SecurityArmCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SecurityArmCode")
	}

	// Simple Field (code)
	code := uint8(m.GetCode())
	_codeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("code", 8, uint8((code)))
	if _codeErr != nil {
		return errors.Wrap(_codeErr, "Error serializing 'code' field")
	}
	// Virtual field
	isDisarmed := m.GetIsDisarmed()
	_ = isDisarmed
	if _isDisarmedErr := writeBuffer.WriteVirtual(ctx, "isDisarmed", m.GetIsDisarmed()); _isDisarmedErr != nil {
		return errors.Wrap(_isDisarmedErr, "Error serializing 'isDisarmed' field")
	}
	// Virtual field
	isFullyArmed := m.GetIsFullyArmed()
	_ = isFullyArmed
	if _isFullyArmedErr := writeBuffer.WriteVirtual(ctx, "isFullyArmed", m.GetIsFullyArmed()); _isFullyArmedErr != nil {
		return errors.Wrap(_isFullyArmedErr, "Error serializing 'isFullyArmed' field")
	}
	// Virtual field
	isPartiallyArmed := m.GetIsPartiallyArmed()
	_ = isPartiallyArmed
	if _isPartiallyArmedErr := writeBuffer.WriteVirtual(ctx, "isPartiallyArmed", m.GetIsPartiallyArmed()); _isPartiallyArmedErr != nil {
		return errors.Wrap(_isPartiallyArmedErr, "Error serializing 'isPartiallyArmed' field")
	}
	// Virtual field
	isArmSubtype := m.GetIsArmSubtype()
	_ = isArmSubtype
	if _isArmSubtypeErr := writeBuffer.WriteVirtual(ctx, "isArmSubtype", m.GetIsArmSubtype()); _isArmSubtypeErr != nil {
		return errors.Wrap(_isArmSubtypeErr, "Error serializing 'isArmSubtype' field")
	}
	// Virtual field
	isReserved := m.GetIsReserved()
	_ = isReserved
	if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}

	if popErr := writeBuffer.PopContext("SecurityArmCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SecurityArmCode")
	}
	return nil
}

func (m *_SecurityArmCode) isSecurityArmCode() bool {
	return true
}

func (m *_SecurityArmCode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
