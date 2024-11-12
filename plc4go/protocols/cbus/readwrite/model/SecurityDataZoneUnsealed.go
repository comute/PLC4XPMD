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

// SecurityDataZoneUnsealed is the corresponding interface of SecurityDataZoneUnsealed
type SecurityDataZoneUnsealed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// IsSecurityDataZoneUnsealed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataZoneUnsealed()
	// CreateBuilder creates a SecurityDataZoneUnsealedBuilder
	CreateSecurityDataZoneUnsealedBuilder() SecurityDataZoneUnsealedBuilder
}

// _SecurityDataZoneUnsealed is the data-structure of this message
type _SecurityDataZoneUnsealed struct {
	SecurityDataContract
	ZoneNumber uint8
}

var _ SecurityDataZoneUnsealed = (*_SecurityDataZoneUnsealed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneUnsealed)(nil)

// NewSecurityDataZoneUnsealed factory function for _SecurityDataZoneUnsealed
func NewSecurityDataZoneUnsealed(commandTypeContainer SecurityCommandTypeContainer, argument byte, zoneNumber uint8) *_SecurityDataZoneUnsealed {
	_result := &_SecurityDataZoneUnsealed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataZoneUnsealedBuilder is a builder for SecurityDataZoneUnsealed
type SecurityDataZoneUnsealedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneNumber uint8) SecurityDataZoneUnsealedBuilder
	// WithZoneNumber adds ZoneNumber (property field)
	WithZoneNumber(uint8) SecurityDataZoneUnsealedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataZoneUnsealed or returns an error if something is wrong
	Build() (SecurityDataZoneUnsealed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataZoneUnsealed
}

// NewSecurityDataZoneUnsealedBuilder() creates a SecurityDataZoneUnsealedBuilder
func NewSecurityDataZoneUnsealedBuilder() SecurityDataZoneUnsealedBuilder {
	return &_SecurityDataZoneUnsealedBuilder{_SecurityDataZoneUnsealed: new(_SecurityDataZoneUnsealed)}
}

type _SecurityDataZoneUnsealedBuilder struct {
	*_SecurityDataZoneUnsealed

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataZoneUnsealedBuilder) = (*_SecurityDataZoneUnsealedBuilder)(nil)

func (b *_SecurityDataZoneUnsealedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataZoneUnsealed
}

func (b *_SecurityDataZoneUnsealedBuilder) WithMandatoryFields(zoneNumber uint8) SecurityDataZoneUnsealedBuilder {
	return b.WithZoneNumber(zoneNumber)
}

func (b *_SecurityDataZoneUnsealedBuilder) WithZoneNumber(zoneNumber uint8) SecurityDataZoneUnsealedBuilder {
	b.ZoneNumber = zoneNumber
	return b
}

func (b *_SecurityDataZoneUnsealedBuilder) Build() (SecurityDataZoneUnsealed, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataZoneUnsealed.deepCopy(), nil
}

func (b *_SecurityDataZoneUnsealedBuilder) MustBuild() SecurityDataZoneUnsealed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataZoneUnsealedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataZoneUnsealedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataZoneUnsealedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataZoneUnsealedBuilder().(*_SecurityDataZoneUnsealedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataZoneUnsealedBuilder creates a SecurityDataZoneUnsealedBuilder
func (b *_SecurityDataZoneUnsealed) CreateSecurityDataZoneUnsealedBuilder() SecurityDataZoneUnsealedBuilder {
	if b == nil {
		return NewSecurityDataZoneUnsealedBuilder()
	}
	return &_SecurityDataZoneUnsealedBuilder{_SecurityDataZoneUnsealed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneUnsealed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneUnsealed) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneUnsealed(structType any) SecurityDataZoneUnsealed {
	if casted, ok := structType.(SecurityDataZoneUnsealed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneUnsealed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneUnsealed) GetTypeName() string {
	return "SecurityDataZoneUnsealed"
}

func (m *_SecurityDataZoneUnsealed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneUnsealed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneUnsealed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneUnsealed SecurityDataZoneUnsealed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneUnsealed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneUnsealed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneUnsealed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneUnsealed")
	}

	return m, nil
}

func (m *_SecurityDataZoneUnsealed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneUnsealed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneUnsealed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneUnsealed")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneUnsealed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneUnsealed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneUnsealed) IsSecurityDataZoneUnsealed() {}

func (m *_SecurityDataZoneUnsealed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataZoneUnsealed) deepCopy() *_SecurityDataZoneUnsealed {
	if m == nil {
		return nil
	}
	_SecurityDataZoneUnsealedCopy := &_SecurityDataZoneUnsealed{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.ZoneNumber,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataZoneUnsealedCopy
}

func (m *_SecurityDataZoneUnsealed) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
