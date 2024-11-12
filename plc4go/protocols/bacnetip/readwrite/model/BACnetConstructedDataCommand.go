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

// BACnetConstructedDataCommand is the corresponding interface of BACnetConstructedDataCommand
type BACnetConstructedDataCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCommand returns Command (property field)
	GetCommand() BACnetNetworkPortCommandTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNetworkPortCommandTagged
	// IsBACnetConstructedDataCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCommand()
	// CreateBuilder creates a BACnetConstructedDataCommandBuilder
	CreateBACnetConstructedDataCommandBuilder() BACnetConstructedDataCommandBuilder
}

// _BACnetConstructedDataCommand is the data-structure of this message
type _BACnetConstructedDataCommand struct {
	BACnetConstructedDataContract
	Command BACnetNetworkPortCommandTagged
}

var _ BACnetConstructedDataCommand = (*_BACnetConstructedDataCommand)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCommand)(nil)

// NewBACnetConstructedDataCommand factory function for _BACnetConstructedDataCommand
func NewBACnetConstructedDataCommand(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, command BACnetNetworkPortCommandTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCommand {
	if command == nil {
		panic("command of type BACnetNetworkPortCommandTagged for BACnetConstructedDataCommand must not be nil")
	}
	_result := &_BACnetConstructedDataCommand{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Command:                       command,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCommandBuilder is a builder for BACnetConstructedDataCommand
type BACnetConstructedDataCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(command BACnetNetworkPortCommandTagged) BACnetConstructedDataCommandBuilder
	// WithCommand adds Command (property field)
	WithCommand(BACnetNetworkPortCommandTagged) BACnetConstructedDataCommandBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(BACnetNetworkPortCommandTaggedBuilder) BACnetNetworkPortCommandTaggedBuilder) BACnetConstructedDataCommandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCommand or returns an error if something is wrong
	Build() (BACnetConstructedDataCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCommand
}

// NewBACnetConstructedDataCommandBuilder() creates a BACnetConstructedDataCommandBuilder
func NewBACnetConstructedDataCommandBuilder() BACnetConstructedDataCommandBuilder {
	return &_BACnetConstructedDataCommandBuilder{_BACnetConstructedDataCommand: new(_BACnetConstructedDataCommand)}
}

type _BACnetConstructedDataCommandBuilder struct {
	*_BACnetConstructedDataCommand

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCommandBuilder) = (*_BACnetConstructedDataCommandBuilder)(nil)

func (b *_BACnetConstructedDataCommandBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCommand
}

func (b *_BACnetConstructedDataCommandBuilder) WithMandatoryFields(command BACnetNetworkPortCommandTagged) BACnetConstructedDataCommandBuilder {
	return b.WithCommand(command)
}

func (b *_BACnetConstructedDataCommandBuilder) WithCommand(command BACnetNetworkPortCommandTagged) BACnetConstructedDataCommandBuilder {
	b.Command = command
	return b
}

func (b *_BACnetConstructedDataCommandBuilder) WithCommandBuilder(builderSupplier func(BACnetNetworkPortCommandTaggedBuilder) BACnetNetworkPortCommandTaggedBuilder) BACnetConstructedDataCommandBuilder {
	builder := builderSupplier(b.Command.CreateBACnetNetworkPortCommandTaggedBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNetworkPortCommandTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCommandBuilder) Build() (BACnetConstructedDataCommand, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCommand.deepCopy(), nil
}

func (b *_BACnetConstructedDataCommandBuilder) MustBuild() BACnetConstructedDataCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCommandBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCommandBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCommandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCommandBuilder().(*_BACnetConstructedDataCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCommandBuilder creates a BACnetConstructedDataCommandBuilder
func (b *_BACnetConstructedDataCommand) CreateBACnetConstructedDataCommandBuilder() BACnetConstructedDataCommandBuilder {
	if b == nil {
		return NewBACnetConstructedDataCommandBuilder()
	}
	return &_BACnetConstructedDataCommandBuilder{_BACnetConstructedDataCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCommand) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCommand) GetCommand() BACnetNetworkPortCommandTagged {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCommand) GetActualValue() BACnetNetworkPortCommandTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetNetworkPortCommandTagged(m.GetCommand())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCommand(structType any) BACnetConstructedDataCommand {
	if casted, ok := structType.(BACnetConstructedDataCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCommand) GetTypeName() string {
	return "BACnetConstructedDataCommand"
}

func (m *_BACnetConstructedDataCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCommand BACnetConstructedDataCommand, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[BACnetNetworkPortCommandTagged](ctx, "command", ReadComplex[BACnetNetworkPortCommandTagged](BACnetNetworkPortCommandTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	actualValue, err := ReadVirtualField[BACnetNetworkPortCommandTagged](ctx, "actualValue", (*BACnetNetworkPortCommandTagged)(nil), command)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCommand")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCommand")
		}

		if err := WriteSimpleField[BACnetNetworkPortCommandTagged](ctx, "command", m.GetCommand(), WriteComplex[BACnetNetworkPortCommandTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCommand")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCommand) IsBACnetConstructedDataCommand() {}

func (m *_BACnetConstructedDataCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCommand) deepCopy() *_BACnetConstructedDataCommand {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCommandCopy := &_BACnetConstructedDataCommand{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetNetworkPortCommandTagged](m.Command),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCommandCopy
}

func (m *_BACnetConstructedDataCommand) String() string {
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
