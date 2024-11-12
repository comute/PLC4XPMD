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

// IdentifyReplyCommandNetworkTerminalLevels is the corresponding interface of IdentifyReplyCommandNetworkTerminalLevels
type IdentifyReplyCommandNetworkTerminalLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	IdentifyReplyCommand
	// GetNetworkTerminalLevels returns NetworkTerminalLevels (property field)
	GetNetworkTerminalLevels() []byte
	// IsIdentifyReplyCommandNetworkTerminalLevels is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommandNetworkTerminalLevels()
	// CreateBuilder creates a IdentifyReplyCommandNetworkTerminalLevelsBuilder
	CreateIdentifyReplyCommandNetworkTerminalLevelsBuilder() IdentifyReplyCommandNetworkTerminalLevelsBuilder
}

// _IdentifyReplyCommandNetworkTerminalLevels is the data-structure of this message
type _IdentifyReplyCommandNetworkTerminalLevels struct {
	IdentifyReplyCommandContract
	NetworkTerminalLevels []byte
}

var _ IdentifyReplyCommandNetworkTerminalLevels = (*_IdentifyReplyCommandNetworkTerminalLevels)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandNetworkTerminalLevels)(nil)

// NewIdentifyReplyCommandNetworkTerminalLevels factory function for _IdentifyReplyCommandNetworkTerminalLevels
func NewIdentifyReplyCommandNetworkTerminalLevels(networkTerminalLevels []byte, numBytes uint8) *_IdentifyReplyCommandNetworkTerminalLevels {
	_result := &_IdentifyReplyCommandNetworkTerminalLevels{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		NetworkTerminalLevels:        networkTerminalLevels,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IdentifyReplyCommandNetworkTerminalLevelsBuilder is a builder for IdentifyReplyCommandNetworkTerminalLevels
type IdentifyReplyCommandNetworkTerminalLevelsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkTerminalLevels []byte) IdentifyReplyCommandNetworkTerminalLevelsBuilder
	// WithNetworkTerminalLevels adds NetworkTerminalLevels (property field)
	WithNetworkTerminalLevels(...byte) IdentifyReplyCommandNetworkTerminalLevelsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() IdentifyReplyCommandBuilder
	// Build builds the IdentifyReplyCommandNetworkTerminalLevels or returns an error if something is wrong
	Build() (IdentifyReplyCommandNetworkTerminalLevels, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IdentifyReplyCommandNetworkTerminalLevels
}

// NewIdentifyReplyCommandNetworkTerminalLevelsBuilder() creates a IdentifyReplyCommandNetworkTerminalLevelsBuilder
func NewIdentifyReplyCommandNetworkTerminalLevelsBuilder() IdentifyReplyCommandNetworkTerminalLevelsBuilder {
	return &_IdentifyReplyCommandNetworkTerminalLevelsBuilder{_IdentifyReplyCommandNetworkTerminalLevels: new(_IdentifyReplyCommandNetworkTerminalLevels)}
}

type _IdentifyReplyCommandNetworkTerminalLevelsBuilder struct {
	*_IdentifyReplyCommandNetworkTerminalLevels

	parentBuilder *_IdentifyReplyCommandBuilder

	err *utils.MultiError
}

var _ (IdentifyReplyCommandNetworkTerminalLevelsBuilder) = (*_IdentifyReplyCommandNetworkTerminalLevelsBuilder)(nil)

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) setParent(contract IdentifyReplyCommandContract) {
	b.IdentifyReplyCommandContract = contract
	contract.(*_IdentifyReplyCommand)._SubType = b._IdentifyReplyCommandNetworkTerminalLevels
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) WithMandatoryFields(networkTerminalLevels []byte) IdentifyReplyCommandNetworkTerminalLevelsBuilder {
	return b.WithNetworkTerminalLevels(networkTerminalLevels...)
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) WithNetworkTerminalLevels(networkTerminalLevels ...byte) IdentifyReplyCommandNetworkTerminalLevelsBuilder {
	b.NetworkTerminalLevels = networkTerminalLevels
	return b
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) Build() (IdentifyReplyCommandNetworkTerminalLevels, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._IdentifyReplyCommandNetworkTerminalLevels.deepCopy(), nil
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) MustBuild() IdentifyReplyCommandNetworkTerminalLevels {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) Done() IdentifyReplyCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewIdentifyReplyCommandBuilder().(*_IdentifyReplyCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) buildForIdentifyReplyCommand() (IdentifyReplyCommand, error) {
	return b.Build()
}

func (b *_IdentifyReplyCommandNetworkTerminalLevelsBuilder) DeepCopy() any {
	_copy := b.CreateIdentifyReplyCommandNetworkTerminalLevelsBuilder().(*_IdentifyReplyCommandNetworkTerminalLevelsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIdentifyReplyCommandNetworkTerminalLevelsBuilder creates a IdentifyReplyCommandNetworkTerminalLevelsBuilder
func (b *_IdentifyReplyCommandNetworkTerminalLevels) CreateIdentifyReplyCommandNetworkTerminalLevelsBuilder() IdentifyReplyCommandNetworkTerminalLevelsBuilder {
	if b == nil {
		return NewIdentifyReplyCommandNetworkTerminalLevelsBuilder()
	}
	return &_IdentifyReplyCommandNetworkTerminalLevelsBuilder{_IdentifyReplyCommandNetworkTerminalLevels: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetAttribute() Attribute {
	return Attribute_NetworkTerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetNetworkTerminalLevels() []byte {
	return m.NetworkTerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandNetworkTerminalLevels(structType any) IdentifyReplyCommandNetworkTerminalLevels {
	if casted, ok := structType.(IdentifyReplyCommandNetworkTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandNetworkTerminalLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandNetworkTerminalLevels"
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.NetworkTerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.NetworkTerminalLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandNetworkTerminalLevels IdentifyReplyCommandNetworkTerminalLevels, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandNetworkTerminalLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandNetworkTerminalLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkTerminalLevels, err := readBuffer.ReadByteArray("networkTerminalLevels", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkTerminalLevels' field"))
	}
	m.NetworkTerminalLevels = networkTerminalLevels

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandNetworkTerminalLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandNetworkTerminalLevels")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandNetworkTerminalLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandNetworkTerminalLevels")
		}

		if err := WriteByteArrayField(ctx, "networkTerminalLevels", m.GetNetworkTerminalLevels(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkTerminalLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandNetworkTerminalLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandNetworkTerminalLevels")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) IsIdentifyReplyCommandNetworkTerminalLevels() {}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) deepCopy() *_IdentifyReplyCommandNetworkTerminalLevels {
	if m == nil {
		return nil
	}
	_IdentifyReplyCommandNetworkTerminalLevelsCopy := &_IdentifyReplyCommandNetworkTerminalLevels{
		m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.NetworkTerminalLevels),
	}
	m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = m
	return _IdentifyReplyCommandNetworkTerminalLevelsCopy
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) String() string {
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
