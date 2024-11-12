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

// BACnetPropertyStatesSecurityLevel is the corresponding interface of BACnetPropertyStatesSecurityLevel
type BACnetPropertyStatesSecurityLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() BACnetSecurityLevelTagged
	// IsBACnetPropertyStatesSecurityLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesSecurityLevel()
	// CreateBuilder creates a BACnetPropertyStatesSecurityLevelBuilder
	CreateBACnetPropertyStatesSecurityLevelBuilder() BACnetPropertyStatesSecurityLevelBuilder
}

// _BACnetPropertyStatesSecurityLevel is the data-structure of this message
type _BACnetPropertyStatesSecurityLevel struct {
	BACnetPropertyStatesContract
	SecurityLevel BACnetSecurityLevelTagged
}

var _ BACnetPropertyStatesSecurityLevel = (*_BACnetPropertyStatesSecurityLevel)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesSecurityLevel)(nil)

// NewBACnetPropertyStatesSecurityLevel factory function for _BACnetPropertyStatesSecurityLevel
func NewBACnetPropertyStatesSecurityLevel(peekedTagHeader BACnetTagHeader, securityLevel BACnetSecurityLevelTagged) *_BACnetPropertyStatesSecurityLevel {
	if securityLevel == nil {
		panic("securityLevel of type BACnetSecurityLevelTagged for BACnetPropertyStatesSecurityLevel must not be nil")
	}
	_result := &_BACnetPropertyStatesSecurityLevel{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		SecurityLevel:                securityLevel,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesSecurityLevelBuilder is a builder for BACnetPropertyStatesSecurityLevel
type BACnetPropertyStatesSecurityLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(securityLevel BACnetSecurityLevelTagged) BACnetPropertyStatesSecurityLevelBuilder
	// WithSecurityLevel adds SecurityLevel (property field)
	WithSecurityLevel(BACnetSecurityLevelTagged) BACnetPropertyStatesSecurityLevelBuilder
	// WithSecurityLevelBuilder adds SecurityLevel (property field) which is build by the builder
	WithSecurityLevelBuilder(func(BACnetSecurityLevelTaggedBuilder) BACnetSecurityLevelTaggedBuilder) BACnetPropertyStatesSecurityLevelBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesSecurityLevel or returns an error if something is wrong
	Build() (BACnetPropertyStatesSecurityLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesSecurityLevel
}

// NewBACnetPropertyStatesSecurityLevelBuilder() creates a BACnetPropertyStatesSecurityLevelBuilder
func NewBACnetPropertyStatesSecurityLevelBuilder() BACnetPropertyStatesSecurityLevelBuilder {
	return &_BACnetPropertyStatesSecurityLevelBuilder{_BACnetPropertyStatesSecurityLevel: new(_BACnetPropertyStatesSecurityLevel)}
}

type _BACnetPropertyStatesSecurityLevelBuilder struct {
	*_BACnetPropertyStatesSecurityLevel

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesSecurityLevelBuilder) = (*_BACnetPropertyStatesSecurityLevelBuilder)(nil)

func (b *_BACnetPropertyStatesSecurityLevelBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesSecurityLevel
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) WithMandatoryFields(securityLevel BACnetSecurityLevelTagged) BACnetPropertyStatesSecurityLevelBuilder {
	return b.WithSecurityLevel(securityLevel)
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) WithSecurityLevel(securityLevel BACnetSecurityLevelTagged) BACnetPropertyStatesSecurityLevelBuilder {
	b.SecurityLevel = securityLevel
	return b
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) WithSecurityLevelBuilder(builderSupplier func(BACnetSecurityLevelTaggedBuilder) BACnetSecurityLevelTaggedBuilder) BACnetPropertyStatesSecurityLevelBuilder {
	builder := builderSupplier(b.SecurityLevel.CreateBACnetSecurityLevelTaggedBuilder())
	var err error
	b.SecurityLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetSecurityLevelTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) Build() (BACnetPropertyStatesSecurityLevel, error) {
	if b.SecurityLevel == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityLevel' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesSecurityLevel.deepCopy(), nil
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) MustBuild() BACnetPropertyStatesSecurityLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesSecurityLevelBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesSecurityLevelBuilder().(*_BACnetPropertyStatesSecurityLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesSecurityLevelBuilder creates a BACnetPropertyStatesSecurityLevelBuilder
func (b *_BACnetPropertyStatesSecurityLevel) CreateBACnetPropertyStatesSecurityLevelBuilder() BACnetPropertyStatesSecurityLevelBuilder {
	if b == nil {
		return NewBACnetPropertyStatesSecurityLevelBuilder()
	}
	return &_BACnetPropertyStatesSecurityLevelBuilder{_BACnetPropertyStatesSecurityLevel: b.deepCopy()}
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

func (m *_BACnetPropertyStatesSecurityLevel) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSecurityLevel) GetSecurityLevel() BACnetSecurityLevelTagged {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSecurityLevel(structType any) BACnetPropertyStatesSecurityLevel {
	if casted, ok := structType.(BACnetPropertyStatesSecurityLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSecurityLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSecurityLevel) GetTypeName() string {
	return "BACnetPropertyStatesSecurityLevel"
}

func (m *_BACnetPropertyStatesSecurityLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (securityLevel)
	lengthInBits += m.SecurityLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesSecurityLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesSecurityLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesSecurityLevel BACnetPropertyStatesSecurityLevel, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSecurityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSecurityLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityLevel, err := ReadSimpleField[BACnetSecurityLevelTagged](ctx, "securityLevel", ReadComplex[BACnetSecurityLevelTagged](BACnetSecurityLevelTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityLevel' field"))
	}
	m.SecurityLevel = securityLevel

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSecurityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSecurityLevel")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesSecurityLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesSecurityLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSecurityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSecurityLevel")
		}

		if err := WriteSimpleField[BACnetSecurityLevelTagged](ctx, "securityLevel", m.GetSecurityLevel(), WriteComplex[BACnetSecurityLevelTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSecurityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSecurityLevel")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesSecurityLevel) IsBACnetPropertyStatesSecurityLevel() {}

func (m *_BACnetPropertyStatesSecurityLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesSecurityLevel) deepCopy() *_BACnetPropertyStatesSecurityLevel {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesSecurityLevelCopy := &_BACnetPropertyStatesSecurityLevel{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetSecurityLevelTagged](m.SecurityLevel),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesSecurityLevelCopy
}

func (m *_BACnetPropertyStatesSecurityLevel) String() string {
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
