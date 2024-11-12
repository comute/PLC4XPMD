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

// BACnetLogDataLogDataEntryFailure is the corresponding interface of BACnetLogDataLogDataEntryFailure
type BACnetLogDataLogDataEntryFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetFailure returns Failure (property field)
	GetFailure() ErrorEnclosed
	// IsBACnetLogDataLogDataEntryFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryFailure()
	// CreateBuilder creates a BACnetLogDataLogDataEntryFailureBuilder
	CreateBACnetLogDataLogDataEntryFailureBuilder() BACnetLogDataLogDataEntryFailureBuilder
}

// _BACnetLogDataLogDataEntryFailure is the data-structure of this message
type _BACnetLogDataLogDataEntryFailure struct {
	BACnetLogDataLogDataEntryContract
	Failure ErrorEnclosed
}

var _ BACnetLogDataLogDataEntryFailure = (*_BACnetLogDataLogDataEntryFailure)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryFailure)(nil)

// NewBACnetLogDataLogDataEntryFailure factory function for _BACnetLogDataLogDataEntryFailure
func NewBACnetLogDataLogDataEntryFailure(peekedTagHeader BACnetTagHeader, failure ErrorEnclosed) *_BACnetLogDataLogDataEntryFailure {
	if failure == nil {
		panic("failure of type ErrorEnclosed for BACnetLogDataLogDataEntryFailure must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryFailure{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		Failure:                           failure,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogDataEntryFailureBuilder is a builder for BACnetLogDataLogDataEntryFailure
type BACnetLogDataLogDataEntryFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(failure ErrorEnclosed) BACnetLogDataLogDataEntryFailureBuilder
	// WithFailure adds Failure (property field)
	WithFailure(ErrorEnclosed) BACnetLogDataLogDataEntryFailureBuilder
	// WithFailureBuilder adds Failure (property field) which is build by the builder
	WithFailureBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) BACnetLogDataLogDataEntryFailureBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetLogDataLogDataEntryBuilder
	// Build builds the BACnetLogDataLogDataEntryFailure or returns an error if something is wrong
	Build() (BACnetLogDataLogDataEntryFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogDataEntryFailure
}

// NewBACnetLogDataLogDataEntryFailureBuilder() creates a BACnetLogDataLogDataEntryFailureBuilder
func NewBACnetLogDataLogDataEntryFailureBuilder() BACnetLogDataLogDataEntryFailureBuilder {
	return &_BACnetLogDataLogDataEntryFailureBuilder{_BACnetLogDataLogDataEntryFailure: new(_BACnetLogDataLogDataEntryFailure)}
}

type _BACnetLogDataLogDataEntryFailureBuilder struct {
	*_BACnetLogDataLogDataEntryFailure

	parentBuilder *_BACnetLogDataLogDataEntryBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogDataEntryFailureBuilder) = (*_BACnetLogDataLogDataEntryFailureBuilder)(nil)

func (b *_BACnetLogDataLogDataEntryFailureBuilder) setParent(contract BACnetLogDataLogDataEntryContract) {
	b.BACnetLogDataLogDataEntryContract = contract
	contract.(*_BACnetLogDataLogDataEntry)._SubType = b._BACnetLogDataLogDataEntryFailure
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) WithMandatoryFields(failure ErrorEnclosed) BACnetLogDataLogDataEntryFailureBuilder {
	return b.WithFailure(failure)
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) WithFailure(failure ErrorEnclosed) BACnetLogDataLogDataEntryFailureBuilder {
	b.Failure = failure
	return b
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) WithFailureBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) BACnetLogDataLogDataEntryFailureBuilder {
	builder := builderSupplier(b.Failure.CreateErrorEnclosedBuilder())
	var err error
	b.Failure, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) Build() (BACnetLogDataLogDataEntryFailure, error) {
	if b.Failure == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'failure' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogDataEntryFailure.deepCopy(), nil
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) MustBuild() BACnetLogDataLogDataEntryFailure {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) Done() BACnetLogDataLogDataEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetLogDataLogDataEntryBuilder().(*_BACnetLogDataLogDataEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) buildForBACnetLogDataLogDataEntry() (BACnetLogDataLogDataEntry, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogDataEntryFailureBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogDataEntryFailureBuilder().(*_BACnetLogDataLogDataEntryFailureBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogDataEntryFailureBuilder creates a BACnetLogDataLogDataEntryFailureBuilder
func (b *_BACnetLogDataLogDataEntryFailure) CreateBACnetLogDataLogDataEntryFailureBuilder() BACnetLogDataLogDataEntryFailureBuilder {
	if b == nil {
		return NewBACnetLogDataLogDataEntryFailureBuilder()
	}
	return &_BACnetLogDataLogDataEntryFailureBuilder{_BACnetLogDataLogDataEntryFailure: b.deepCopy()}
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

func (m *_BACnetLogDataLogDataEntryFailure) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryFailure) GetFailure() ErrorEnclosed {
	return m.Failure
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryFailure(structType any) BACnetLogDataLogDataEntryFailure {
	if casted, ok := structType.(BACnetLogDataLogDataEntryFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryFailure) GetTypeName() string {
	return "BACnetLogDataLogDataEntryFailure"
}

func (m *_BACnetLogDataLogDataEntryFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (failure)
	lengthInBits += m.Failure.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryFailure BACnetLogDataLogDataEntryFailure, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	failure, err := ReadSimpleField[ErrorEnclosed](ctx, "failure", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(7))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'failure' field"))
	}
	m.Failure = failure

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryFailure")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryFailure")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "failure", m.GetFailure(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'failure' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryFailure")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryFailure) IsBACnetLogDataLogDataEntryFailure() {}

func (m *_BACnetLogDataLogDataEntryFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryFailure) deepCopy() *_BACnetLogDataLogDataEntryFailure {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryFailureCopy := &_BACnetLogDataLogDataEntryFailure{
		m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).deepCopy(),
		utils.DeepCopy[ErrorEnclosed](m.Failure),
	}
	m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryFailureCopy
}

func (m *_BACnetLogDataLogDataEntryFailure) String() string {
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
