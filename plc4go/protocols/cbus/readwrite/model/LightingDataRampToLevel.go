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

// LightingDataRampToLevel is the corresponding interface of LightingDataRampToLevel
type LightingDataRampToLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
	// GetLevel returns Level (property field)
	GetLevel() byte
	// IsLightingDataRampToLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLightingDataRampToLevel()
	// CreateBuilder creates a LightingDataRampToLevelBuilder
	CreateLightingDataRampToLevelBuilder() LightingDataRampToLevelBuilder
}

// _LightingDataRampToLevel is the data-structure of this message
type _LightingDataRampToLevel struct {
	LightingDataContract
	Group byte
	Level byte
}

var _ LightingDataRampToLevel = (*_LightingDataRampToLevel)(nil)
var _ LightingDataRequirements = (*_LightingDataRampToLevel)(nil)

// NewLightingDataRampToLevel factory function for _LightingDataRampToLevel
func NewLightingDataRampToLevel(commandTypeContainer LightingCommandTypeContainer, group byte, level byte) *_LightingDataRampToLevel {
	_result := &_LightingDataRampToLevel{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
		Level:                level,
	}
	_result.LightingDataContract.(*_LightingData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LightingDataRampToLevelBuilder is a builder for LightingDataRampToLevel
type LightingDataRampToLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(group byte, level byte) LightingDataRampToLevelBuilder
	// WithGroup adds Group (property field)
	WithGroup(byte) LightingDataRampToLevelBuilder
	// WithLevel adds Level (property field)
	WithLevel(byte) LightingDataRampToLevelBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() LightingDataBuilder
	// Build builds the LightingDataRampToLevel or returns an error if something is wrong
	Build() (LightingDataRampToLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LightingDataRampToLevel
}

// NewLightingDataRampToLevelBuilder() creates a LightingDataRampToLevelBuilder
func NewLightingDataRampToLevelBuilder() LightingDataRampToLevelBuilder {
	return &_LightingDataRampToLevelBuilder{_LightingDataRampToLevel: new(_LightingDataRampToLevel)}
}

type _LightingDataRampToLevelBuilder struct {
	*_LightingDataRampToLevel

	parentBuilder *_LightingDataBuilder

	err *utils.MultiError
}

var _ (LightingDataRampToLevelBuilder) = (*_LightingDataRampToLevelBuilder)(nil)

func (b *_LightingDataRampToLevelBuilder) setParent(contract LightingDataContract) {
	b.LightingDataContract = contract
	contract.(*_LightingData)._SubType = b._LightingDataRampToLevel
}

func (b *_LightingDataRampToLevelBuilder) WithMandatoryFields(group byte, level byte) LightingDataRampToLevelBuilder {
	return b.WithGroup(group).WithLevel(level)
}

func (b *_LightingDataRampToLevelBuilder) WithGroup(group byte) LightingDataRampToLevelBuilder {
	b.Group = group
	return b
}

func (b *_LightingDataRampToLevelBuilder) WithLevel(level byte) LightingDataRampToLevelBuilder {
	b.Level = level
	return b
}

func (b *_LightingDataRampToLevelBuilder) Build() (LightingDataRampToLevel, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LightingDataRampToLevel.deepCopy(), nil
}

func (b *_LightingDataRampToLevelBuilder) MustBuild() LightingDataRampToLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LightingDataRampToLevelBuilder) Done() LightingDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewLightingDataBuilder().(*_LightingDataBuilder)
	}
	return b.parentBuilder
}

func (b *_LightingDataRampToLevelBuilder) buildForLightingData() (LightingData, error) {
	return b.Build()
}

func (b *_LightingDataRampToLevelBuilder) DeepCopy() any {
	_copy := b.CreateLightingDataRampToLevelBuilder().(*_LightingDataRampToLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLightingDataRampToLevelBuilder creates a LightingDataRampToLevelBuilder
func (b *_LightingDataRampToLevel) CreateLightingDataRampToLevelBuilder() LightingDataRampToLevelBuilder {
	if b == nil {
		return NewLightingDataRampToLevelBuilder()
	}
	return &_LightingDataRampToLevelBuilder{_LightingDataRampToLevel: b.deepCopy()}
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

func (m *_LightingDataRampToLevel) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataRampToLevel) GetGroup() byte {
	return m.Group
}

func (m *_LightingDataRampToLevel) GetLevel() byte {
	return m.Level
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLightingDataRampToLevel(structType any) LightingDataRampToLevel {
	if casted, ok := structType.(LightingDataRampToLevel); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataRampToLevel); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataRampToLevel) GetTypeName() string {
	return "LightingDataRampToLevel"
}

func (m *_LightingDataRampToLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataRampToLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataRampToLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData) (__lightingDataRampToLevel LightingDataRampToLevel, err error) {
	m.LightingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataRampToLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataRampToLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	level, err := ReadSimpleField(ctx, "level", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	if closeErr := readBuffer.CloseContext("LightingDataRampToLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataRampToLevel")
	}

	return m, nil
}

func (m *_LightingDataRampToLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataRampToLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataRampToLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataRampToLevel")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if err := WriteSimpleField[byte](ctx, "level", m.GetLevel(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataRampToLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataRampToLevel")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataRampToLevel) IsLightingDataRampToLevel() {}

func (m *_LightingDataRampToLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LightingDataRampToLevel) deepCopy() *_LightingDataRampToLevel {
	if m == nil {
		return nil
	}
	_LightingDataRampToLevelCopy := &_LightingDataRampToLevel{
		m.LightingDataContract.(*_LightingData).deepCopy(),
		m.Group,
		m.Level,
	}
	m.LightingDataContract.(*_LightingData)._SubType = m
	return _LightingDataRampToLevelCopy
}

func (m *_LightingDataRampToLevel) String() string {
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
