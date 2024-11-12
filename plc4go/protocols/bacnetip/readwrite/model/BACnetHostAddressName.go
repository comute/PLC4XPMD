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

// BACnetHostAddressName is the corresponding interface of BACnetHostAddressName
type BACnetHostAddressName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetHostAddress
	// GetName returns Name (property field)
	GetName() BACnetContextTagCharacterString
	// IsBACnetHostAddressName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostAddressName()
	// CreateBuilder creates a BACnetHostAddressNameBuilder
	CreateBACnetHostAddressNameBuilder() BACnetHostAddressNameBuilder
}

// _BACnetHostAddressName is the data-structure of this message
type _BACnetHostAddressName struct {
	BACnetHostAddressContract
	Name BACnetContextTagCharacterString
}

var _ BACnetHostAddressName = (*_BACnetHostAddressName)(nil)
var _ BACnetHostAddressRequirements = (*_BACnetHostAddressName)(nil)

// NewBACnetHostAddressName factory function for _BACnetHostAddressName
func NewBACnetHostAddressName(peekedTagHeader BACnetTagHeader, name BACnetContextTagCharacterString) *_BACnetHostAddressName {
	if name == nil {
		panic("name of type BACnetContextTagCharacterString for BACnetHostAddressName must not be nil")
	}
	_result := &_BACnetHostAddressName{
		BACnetHostAddressContract: NewBACnetHostAddress(peekedTagHeader),
		Name:                      name,
	}
	_result.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetHostAddressNameBuilder is a builder for BACnetHostAddressName
type BACnetHostAddressNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name BACnetContextTagCharacterString) BACnetHostAddressNameBuilder
	// WithName adds Name (property field)
	WithName(BACnetContextTagCharacterString) BACnetHostAddressNameBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetHostAddressNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetHostAddressBuilder
	// Build builds the BACnetHostAddressName or returns an error if something is wrong
	Build() (BACnetHostAddressName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetHostAddressName
}

// NewBACnetHostAddressNameBuilder() creates a BACnetHostAddressNameBuilder
func NewBACnetHostAddressNameBuilder() BACnetHostAddressNameBuilder {
	return &_BACnetHostAddressNameBuilder{_BACnetHostAddressName: new(_BACnetHostAddressName)}
}

type _BACnetHostAddressNameBuilder struct {
	*_BACnetHostAddressName

	parentBuilder *_BACnetHostAddressBuilder

	err *utils.MultiError
}

var _ (BACnetHostAddressNameBuilder) = (*_BACnetHostAddressNameBuilder)(nil)

func (b *_BACnetHostAddressNameBuilder) setParent(contract BACnetHostAddressContract) {
	b.BACnetHostAddressContract = contract
	contract.(*_BACnetHostAddress)._SubType = b._BACnetHostAddressName
}

func (b *_BACnetHostAddressNameBuilder) WithMandatoryFields(name BACnetContextTagCharacterString) BACnetHostAddressNameBuilder {
	return b.WithName(name)
}

func (b *_BACnetHostAddressNameBuilder) WithName(name BACnetContextTagCharacterString) BACnetHostAddressNameBuilder {
	b.Name = name
	return b
}

func (b *_BACnetHostAddressNameBuilder) WithNameBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetHostAddressNameBuilder {
	builder := builderSupplier(b.Name.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetHostAddressNameBuilder) Build() (BACnetHostAddressName, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetHostAddressName.deepCopy(), nil
}

func (b *_BACnetHostAddressNameBuilder) MustBuild() BACnetHostAddressName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetHostAddressNameBuilder) Done() BACnetHostAddressBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetHostAddressBuilder().(*_BACnetHostAddressBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetHostAddressNameBuilder) buildForBACnetHostAddress() (BACnetHostAddress, error) {
	return b.Build()
}

func (b *_BACnetHostAddressNameBuilder) DeepCopy() any {
	_copy := b.CreateBACnetHostAddressNameBuilder().(*_BACnetHostAddressNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetHostAddressNameBuilder creates a BACnetHostAddressNameBuilder
func (b *_BACnetHostAddressName) CreateBACnetHostAddressNameBuilder() BACnetHostAddressNameBuilder {
	if b == nil {
		return NewBACnetHostAddressNameBuilder()
	}
	return &_BACnetHostAddressNameBuilder{_BACnetHostAddressName: b.deepCopy()}
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

func (m *_BACnetHostAddressName) GetParent() BACnetHostAddressContract {
	return m.BACnetHostAddressContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressName) GetName() BACnetContextTagCharacterString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressName(structType any) BACnetHostAddressName {
	if casted, ok := structType.(BACnetHostAddressName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressName) GetTypeName() string {
	return "BACnetHostAddressName"
}

func (m *_BACnetHostAddressName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetHostAddressContract.(*_BACnetHostAddress).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetHostAddressName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetHostAddress) (__bACnetHostAddressName BACnetHostAddressName, err error) {
	m.BACnetHostAddressContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "name", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	if closeErr := readBuffer.CloseContext("BACnetHostAddressName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressName")
	}

	return m, nil
}

func (m *_BACnetHostAddressName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressName")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "name", m.GetName(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressName")
		}
		return nil
	}
	return m.BACnetHostAddressContract.(*_BACnetHostAddress).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetHostAddressName) IsBACnetHostAddressName() {}

func (m *_BACnetHostAddressName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetHostAddressName) deepCopy() *_BACnetHostAddressName {
	if m == nil {
		return nil
	}
	_BACnetHostAddressNameCopy := &_BACnetHostAddressName{
		m.BACnetHostAddressContract.(*_BACnetHostAddress).deepCopy(),
		utils.DeepCopy[BACnetContextTagCharacterString](m.Name),
	}
	m.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = m
	return _BACnetHostAddressNameCopy
}

func (m *_BACnetHostAddressName) String() string {
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
