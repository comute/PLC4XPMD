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

// BACnetServiceAckAtomicWriteFile is the corresponding interface of BACnetServiceAckAtomicWriteFile
type BACnetServiceAckAtomicWriteFile interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetContextTagSignedInteger
	// IsBACnetServiceAckAtomicWriteFile is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicWriteFile()
	// CreateBuilder creates a BACnetServiceAckAtomicWriteFileBuilder
	CreateBACnetServiceAckAtomicWriteFileBuilder() BACnetServiceAckAtomicWriteFileBuilder
}

// _BACnetServiceAckAtomicWriteFile is the data-structure of this message
type _BACnetServiceAckAtomicWriteFile struct {
	BACnetServiceAckContract
	FileStartPosition BACnetContextTagSignedInteger
}

var _ BACnetServiceAckAtomicWriteFile = (*_BACnetServiceAckAtomicWriteFile)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckAtomicWriteFile)(nil)

// NewBACnetServiceAckAtomicWriteFile factory function for _BACnetServiceAckAtomicWriteFile
func NewBACnetServiceAckAtomicWriteFile(fileStartPosition BACnetContextTagSignedInteger, serviceAckLength uint32) *_BACnetServiceAckAtomicWriteFile {
	if fileStartPosition == nil {
		panic("fileStartPosition of type BACnetContextTagSignedInteger for BACnetServiceAckAtomicWriteFile must not be nil")
	}
	_result := &_BACnetServiceAckAtomicWriteFile{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		FileStartPosition:        fileStartPosition,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckAtomicWriteFileBuilder is a builder for BACnetServiceAckAtomicWriteFile
type BACnetServiceAckAtomicWriteFileBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fileStartPosition BACnetContextTagSignedInteger) BACnetServiceAckAtomicWriteFileBuilder
	// WithFileStartPosition adds FileStartPosition (property field)
	WithFileStartPosition(BACnetContextTagSignedInteger) BACnetServiceAckAtomicWriteFileBuilder
	// WithFileStartPositionBuilder adds FileStartPosition (property field) which is build by the builder
	WithFileStartPositionBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetServiceAckAtomicWriteFileBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetServiceAckBuilder
	// Build builds the BACnetServiceAckAtomicWriteFile or returns an error if something is wrong
	Build() (BACnetServiceAckAtomicWriteFile, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckAtomicWriteFile
}

// NewBACnetServiceAckAtomicWriteFileBuilder() creates a BACnetServiceAckAtomicWriteFileBuilder
func NewBACnetServiceAckAtomicWriteFileBuilder() BACnetServiceAckAtomicWriteFileBuilder {
	return &_BACnetServiceAckAtomicWriteFileBuilder{_BACnetServiceAckAtomicWriteFile: new(_BACnetServiceAckAtomicWriteFile)}
}

type _BACnetServiceAckAtomicWriteFileBuilder struct {
	*_BACnetServiceAckAtomicWriteFile

	parentBuilder *_BACnetServiceAckBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckAtomicWriteFileBuilder) = (*_BACnetServiceAckAtomicWriteFileBuilder)(nil)

func (b *_BACnetServiceAckAtomicWriteFileBuilder) setParent(contract BACnetServiceAckContract) {
	b.BACnetServiceAckContract = contract
	contract.(*_BACnetServiceAck)._SubType = b._BACnetServiceAckAtomicWriteFile
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) WithMandatoryFields(fileStartPosition BACnetContextTagSignedInteger) BACnetServiceAckAtomicWriteFileBuilder {
	return b.WithFileStartPosition(fileStartPosition)
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) WithFileStartPosition(fileStartPosition BACnetContextTagSignedInteger) BACnetServiceAckAtomicWriteFileBuilder {
	b.FileStartPosition = fileStartPosition
	return b
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) WithFileStartPositionBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetServiceAckAtomicWriteFileBuilder {
	builder := builderSupplier(b.FileStartPosition.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.FileStartPosition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) Build() (BACnetServiceAckAtomicWriteFile, error) {
	if b.FileStartPosition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fileStartPosition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckAtomicWriteFile.deepCopy(), nil
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) MustBuild() BACnetServiceAckAtomicWriteFile {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) Done() BACnetServiceAckBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetServiceAckBuilder().(*_BACnetServiceAckBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) buildForBACnetServiceAck() (BACnetServiceAck, error) {
	return b.Build()
}

func (b *_BACnetServiceAckAtomicWriteFileBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckAtomicWriteFileBuilder().(*_BACnetServiceAckAtomicWriteFileBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckAtomicWriteFileBuilder creates a BACnetServiceAckAtomicWriteFileBuilder
func (b *_BACnetServiceAckAtomicWriteFile) CreateBACnetServiceAckAtomicWriteFileBuilder() BACnetServiceAckAtomicWriteFileBuilder {
	if b == nil {
		return NewBACnetServiceAckAtomicWriteFileBuilder()
	}
	return &_BACnetServiceAckAtomicWriteFileBuilder{_BACnetServiceAckAtomicWriteFile: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicWriteFile) GetFileStartPosition() BACnetContextTagSignedInteger {
	return m.FileStartPosition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicWriteFile(structType any) BACnetServiceAckAtomicWriteFile {
	if casted, ok := structType.(BACnetServiceAckAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicWriteFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicWriteFile) GetTypeName() string {
	return "BACnetServiceAckAtomicWriteFile"
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicWriteFile) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckAtomicWriteFile) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckAtomicWriteFile BACnetServiceAckAtomicWriteFile, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicWriteFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicWriteFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicWriteFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicWriteFile")
	}

	return m, nil
}

func (m *_BACnetServiceAckAtomicWriteFile) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicWriteFile) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicWriteFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicWriteFile")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicWriteFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicWriteFile")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicWriteFile) IsBACnetServiceAckAtomicWriteFile() {}

func (m *_BACnetServiceAckAtomicWriteFile) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckAtomicWriteFile) deepCopy() *_BACnetServiceAckAtomicWriteFile {
	if m == nil {
		return nil
	}
	_BACnetServiceAckAtomicWriteFileCopy := &_BACnetServiceAckAtomicWriteFile{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.FileStartPosition),
	}
	m.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckAtomicWriteFileCopy
}

func (m *_BACnetServiceAckAtomicWriteFile) String() string {
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
