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

// BACnetChannelValueLightingCommand is the corresponding interface of BACnetChannelValueLightingCommand
type BACnetChannelValueLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
	// IsBACnetChannelValueLightingCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueLightingCommand()
	// CreateBuilder creates a BACnetChannelValueLightingCommandBuilder
	CreateBACnetChannelValueLightingCommandBuilder() BACnetChannelValueLightingCommandBuilder
}

// _BACnetChannelValueLightingCommand is the data-structure of this message
type _BACnetChannelValueLightingCommand struct {
	BACnetChannelValueContract
	LigthingCommandValue BACnetLightingCommandEnclosed
}

var _ BACnetChannelValueLightingCommand = (*_BACnetChannelValueLightingCommand)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueLightingCommand)(nil)

// NewBACnetChannelValueLightingCommand factory function for _BACnetChannelValueLightingCommand
func NewBACnetChannelValueLightingCommand(peekedTagHeader BACnetTagHeader, ligthingCommandValue BACnetLightingCommandEnclosed) *_BACnetChannelValueLightingCommand {
	if ligthingCommandValue == nil {
		panic("ligthingCommandValue of type BACnetLightingCommandEnclosed for BACnetChannelValueLightingCommand must not be nil")
	}
	_result := &_BACnetChannelValueLightingCommand{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		LigthingCommandValue:       ligthingCommandValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueLightingCommandBuilder is a builder for BACnetChannelValueLightingCommand
type BACnetChannelValueLightingCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetChannelValueLightingCommandBuilder
	// WithLigthingCommandValue adds LigthingCommandValue (property field)
	WithLigthingCommandValue(BACnetLightingCommandEnclosed) BACnetChannelValueLightingCommandBuilder
	// WithLigthingCommandValueBuilder adds LigthingCommandValue (property field) which is build by the builder
	WithLigthingCommandValueBuilder(func(BACnetLightingCommandEnclosedBuilder) BACnetLightingCommandEnclosedBuilder) BACnetChannelValueLightingCommandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueLightingCommand or returns an error if something is wrong
	Build() (BACnetChannelValueLightingCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueLightingCommand
}

// NewBACnetChannelValueLightingCommandBuilder() creates a BACnetChannelValueLightingCommandBuilder
func NewBACnetChannelValueLightingCommandBuilder() BACnetChannelValueLightingCommandBuilder {
	return &_BACnetChannelValueLightingCommandBuilder{_BACnetChannelValueLightingCommand: new(_BACnetChannelValueLightingCommand)}
}

type _BACnetChannelValueLightingCommandBuilder struct {
	*_BACnetChannelValueLightingCommand

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueLightingCommandBuilder) = (*_BACnetChannelValueLightingCommandBuilder)(nil)

func (b *_BACnetChannelValueLightingCommandBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueLightingCommand
}

func (b *_BACnetChannelValueLightingCommandBuilder) WithMandatoryFields(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetChannelValueLightingCommandBuilder {
	return b.WithLigthingCommandValue(ligthingCommandValue)
}

func (b *_BACnetChannelValueLightingCommandBuilder) WithLigthingCommandValue(ligthingCommandValue BACnetLightingCommandEnclosed) BACnetChannelValueLightingCommandBuilder {
	b.LigthingCommandValue = ligthingCommandValue
	return b
}

func (b *_BACnetChannelValueLightingCommandBuilder) WithLigthingCommandValueBuilder(builderSupplier func(BACnetLightingCommandEnclosedBuilder) BACnetLightingCommandEnclosedBuilder) BACnetChannelValueLightingCommandBuilder {
	builder := builderSupplier(b.LigthingCommandValue.CreateBACnetLightingCommandEnclosedBuilder())
	var err error
	b.LigthingCommandValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingCommandEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetChannelValueLightingCommandBuilder) Build() (BACnetChannelValueLightingCommand, error) {
	if b.LigthingCommandValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ligthingCommandValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueLightingCommand.deepCopy(), nil
}

func (b *_BACnetChannelValueLightingCommandBuilder) MustBuild() BACnetChannelValueLightingCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueLightingCommandBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueLightingCommandBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueLightingCommandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueLightingCommandBuilder().(*_BACnetChannelValueLightingCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueLightingCommandBuilder creates a BACnetChannelValueLightingCommandBuilder
func (b *_BACnetChannelValueLightingCommand) CreateBACnetChannelValueLightingCommandBuilder() BACnetChannelValueLightingCommandBuilder {
	if b == nil {
		return NewBACnetChannelValueLightingCommandBuilder()
	}
	return &_BACnetChannelValueLightingCommandBuilder{_BACnetChannelValueLightingCommand: b.deepCopy()}
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

func (m *_BACnetChannelValueLightingCommand) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueLightingCommand(structType any) BACnetChannelValueLightingCommand {
	if casted, ok := structType.(BACnetChannelValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueLightingCommand) GetTypeName() string {
	return "BACnetChannelValueLightingCommand"
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueLightingCommand BACnetChannelValueLightingCommand, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ligthingCommandValue, err := ReadSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", ReadComplex[BACnetLightingCommandEnclosed](BACnetLightingCommandEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ligthingCommandValue' field"))
	}
	m.LigthingCommandValue = ligthingCommandValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueLightingCommand")
	}

	return m, nil
}

func (m *_BACnetChannelValueLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", m.GetLigthingCommandValue(), WriteComplex[BACnetLightingCommandEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueLightingCommand")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueLightingCommand) IsBACnetChannelValueLightingCommand() {}

func (m *_BACnetChannelValueLightingCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueLightingCommand) deepCopy() *_BACnetChannelValueLightingCommand {
	if m == nil {
		return nil
	}
	_BACnetChannelValueLightingCommandCopy := &_BACnetChannelValueLightingCommand{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetLightingCommandEnclosed](m.LigthingCommandValue),
	}
	m.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueLightingCommandCopy
}

func (m *_BACnetChannelValueLightingCommand) String() string {
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
