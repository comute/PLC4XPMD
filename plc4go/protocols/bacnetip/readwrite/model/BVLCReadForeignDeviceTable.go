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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadForeignDeviceTable is the corresponding interface of BVLCReadForeignDeviceTable
type BVLCReadForeignDeviceTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// IsBVLCReadForeignDeviceTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCReadForeignDeviceTable()
	// CreateBuilder creates a BVLCReadForeignDeviceTableBuilder
	CreateBVLCReadForeignDeviceTableBuilder() BVLCReadForeignDeviceTableBuilder
}

// _BVLCReadForeignDeviceTable is the data-structure of this message
type _BVLCReadForeignDeviceTable struct {
	BVLCContract
}

var _ BVLCReadForeignDeviceTable = (*_BVLCReadForeignDeviceTable)(nil)
var _ BVLCRequirements = (*_BVLCReadForeignDeviceTable)(nil)

// NewBVLCReadForeignDeviceTable factory function for _BVLCReadForeignDeviceTable
func NewBVLCReadForeignDeviceTable() *_BVLCReadForeignDeviceTable {
	_result := &_BVLCReadForeignDeviceTable{
		BVLCContract: NewBVLC(),
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCReadForeignDeviceTableBuilder is a builder for BVLCReadForeignDeviceTable
type BVLCReadForeignDeviceTableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BVLCReadForeignDeviceTableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BVLCBuilder
	// Build builds the BVLCReadForeignDeviceTable or returns an error if something is wrong
	Build() (BVLCReadForeignDeviceTable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCReadForeignDeviceTable
}

// NewBVLCReadForeignDeviceTableBuilder() creates a BVLCReadForeignDeviceTableBuilder
func NewBVLCReadForeignDeviceTableBuilder() BVLCReadForeignDeviceTableBuilder {
	return &_BVLCReadForeignDeviceTableBuilder{_BVLCReadForeignDeviceTable: new(_BVLCReadForeignDeviceTable)}
}

type _BVLCReadForeignDeviceTableBuilder struct {
	*_BVLCReadForeignDeviceTable

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCReadForeignDeviceTableBuilder) = (*_BVLCReadForeignDeviceTableBuilder)(nil)

func (b *_BVLCReadForeignDeviceTableBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
	contract.(*_BVLC)._SubType = b._BVLCReadForeignDeviceTable
}

func (b *_BVLCReadForeignDeviceTableBuilder) WithMandatoryFields() BVLCReadForeignDeviceTableBuilder {
	return b
}

func (b *_BVLCReadForeignDeviceTableBuilder) Build() (BVLCReadForeignDeviceTable, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCReadForeignDeviceTable.deepCopy(), nil
}

func (b *_BVLCReadForeignDeviceTableBuilder) MustBuild() BVLCReadForeignDeviceTable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCReadForeignDeviceTableBuilder) Done() BVLCBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBVLCBuilder().(*_BVLCBuilder)
	}
	return b.parentBuilder
}

func (b *_BVLCReadForeignDeviceTableBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCReadForeignDeviceTableBuilder) DeepCopy() any {
	_copy := b.CreateBVLCReadForeignDeviceTableBuilder().(*_BVLCReadForeignDeviceTableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCReadForeignDeviceTableBuilder creates a BVLCReadForeignDeviceTableBuilder
func (b *_BVLCReadForeignDeviceTable) CreateBVLCReadForeignDeviceTableBuilder() BVLCReadForeignDeviceTableBuilder {
	if b == nil {
		return NewBVLCReadForeignDeviceTableBuilder()
	}
	return &_BVLCReadForeignDeviceTableBuilder{_BVLCReadForeignDeviceTable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadForeignDeviceTable) GetBvlcFunction() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadForeignDeviceTable) GetParent() BVLCContract {
	return m.BVLCContract
}

// Deprecated: use the interface for direct cast
func CastBVLCReadForeignDeviceTable(structType any) BVLCReadForeignDeviceTable {
	if casted, ok := structType.(BVLCReadForeignDeviceTable); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadForeignDeviceTable); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadForeignDeviceTable) GetTypeName() string {
	return "BVLCReadForeignDeviceTable"
}

func (m *_BVLCReadForeignDeviceTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BVLCReadForeignDeviceTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCReadForeignDeviceTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC) (__bVLCReadForeignDeviceTable BVLCReadForeignDeviceTable, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadForeignDeviceTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadForeignDeviceTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCReadForeignDeviceTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadForeignDeviceTable")
	}

	return m, nil
}

func (m *_BVLCReadForeignDeviceTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadForeignDeviceTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadForeignDeviceTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadForeignDeviceTable")
		}

		if popErr := writeBuffer.PopContext("BVLCReadForeignDeviceTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadForeignDeviceTable")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCReadForeignDeviceTable) IsBVLCReadForeignDeviceTable() {}

func (m *_BVLCReadForeignDeviceTable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCReadForeignDeviceTable) deepCopy() *_BVLCReadForeignDeviceTable {
	if m == nil {
		return nil
	}
	_BVLCReadForeignDeviceTableCopy := &_BVLCReadForeignDeviceTable{
		m.BVLCContract.(*_BVLC).deepCopy(),
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCReadForeignDeviceTableCopy
}

func (m *_BVLCReadForeignDeviceTable) String() string {
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
