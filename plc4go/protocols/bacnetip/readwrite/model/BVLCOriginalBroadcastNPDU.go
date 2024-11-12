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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCOriginalBroadcastNPDU is the corresponding interface of BVLCOriginalBroadcastNPDU
type BVLCOriginalBroadcastNPDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
	// IsBVLCOriginalBroadcastNPDU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCOriginalBroadcastNPDU()
	// CreateBuilder creates a BVLCOriginalBroadcastNPDUBuilder
	CreateBVLCOriginalBroadcastNPDUBuilder() BVLCOriginalBroadcastNPDUBuilder
}

// _BVLCOriginalBroadcastNPDU is the data-structure of this message
type _BVLCOriginalBroadcastNPDU struct {
	BVLCContract
	Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCOriginalBroadcastNPDU = (*_BVLCOriginalBroadcastNPDU)(nil)
var _ BVLCRequirements = (*_BVLCOriginalBroadcastNPDU)(nil)

// NewBVLCOriginalBroadcastNPDU factory function for _BVLCOriginalBroadcastNPDU
func NewBVLCOriginalBroadcastNPDU(npdu NPDU, bvlcPayloadLength uint16) *_BVLCOriginalBroadcastNPDU {
	if npdu == nil {
		panic("npdu of type NPDU for BVLCOriginalBroadcastNPDU must not be nil")
	}
	_result := &_BVLCOriginalBroadcastNPDU{
		BVLCContract: NewBVLC(),
		Npdu:         npdu,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCOriginalBroadcastNPDUBuilder is a builder for BVLCOriginalBroadcastNPDU
type BVLCOriginalBroadcastNPDUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(npdu NPDU) BVLCOriginalBroadcastNPDUBuilder
	// WithNpdu adds Npdu (property field)
	WithNpdu(NPDU) BVLCOriginalBroadcastNPDUBuilder
	// WithNpduBuilder adds Npdu (property field) which is build by the builder
	WithNpduBuilder(func(NPDUBuilder) NPDUBuilder) BVLCOriginalBroadcastNPDUBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BVLCBuilder
	// Build builds the BVLCOriginalBroadcastNPDU or returns an error if something is wrong
	Build() (BVLCOriginalBroadcastNPDU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCOriginalBroadcastNPDU
}

// NewBVLCOriginalBroadcastNPDUBuilder() creates a BVLCOriginalBroadcastNPDUBuilder
func NewBVLCOriginalBroadcastNPDUBuilder() BVLCOriginalBroadcastNPDUBuilder {
	return &_BVLCOriginalBroadcastNPDUBuilder{_BVLCOriginalBroadcastNPDU: new(_BVLCOriginalBroadcastNPDU)}
}

type _BVLCOriginalBroadcastNPDUBuilder struct {
	*_BVLCOriginalBroadcastNPDU

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCOriginalBroadcastNPDUBuilder) = (*_BVLCOriginalBroadcastNPDUBuilder)(nil)

func (b *_BVLCOriginalBroadcastNPDUBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
	contract.(*_BVLC)._SubType = b._BVLCOriginalBroadcastNPDU
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) WithMandatoryFields(npdu NPDU) BVLCOriginalBroadcastNPDUBuilder {
	return b.WithNpdu(npdu)
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) WithNpdu(npdu NPDU) BVLCOriginalBroadcastNPDUBuilder {
	b.Npdu = npdu
	return b
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) WithNpduBuilder(builderSupplier func(NPDUBuilder) NPDUBuilder) BVLCOriginalBroadcastNPDUBuilder {
	builder := builderSupplier(b.Npdu.CreateNPDUBuilder())
	var err error
	b.Npdu, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NPDUBuilder failed"))
	}
	return b
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) Build() (BVLCOriginalBroadcastNPDU, error) {
	if b.Npdu == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'npdu' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCOriginalBroadcastNPDU.deepCopy(), nil
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) MustBuild() BVLCOriginalBroadcastNPDU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) Done() BVLCBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBVLCBuilder().(*_BVLCBuilder)
	}
	return b.parentBuilder
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCOriginalBroadcastNPDUBuilder) DeepCopy() any {
	_copy := b.CreateBVLCOriginalBroadcastNPDUBuilder().(*_BVLCOriginalBroadcastNPDUBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCOriginalBroadcastNPDUBuilder creates a BVLCOriginalBroadcastNPDUBuilder
func (b *_BVLCOriginalBroadcastNPDU) CreateBVLCOriginalBroadcastNPDUBuilder() BVLCOriginalBroadcastNPDUBuilder {
	if b == nil {
		return NewBVLCOriginalBroadcastNPDUBuilder()
	}
	return &_BVLCOriginalBroadcastNPDUBuilder{_BVLCOriginalBroadcastNPDU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCOriginalBroadcastNPDU) GetBvlcFunction() uint8 {
	return 0x0B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCOriginalBroadcastNPDU) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCOriginalBroadcastNPDU) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCOriginalBroadcastNPDU(structType any) BVLCOriginalBroadcastNPDU {
	if casted, ok := structType.(BVLCOriginalBroadcastNPDU); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCOriginalBroadcastNPDU); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCOriginalBroadcastNPDU) GetTypeName() string {
	return "BVLCOriginalBroadcastNPDU"
}

func (m *_BVLCOriginalBroadcastNPDU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BVLCOriginalBroadcastNPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCOriginalBroadcastNPDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCOriginalBroadcastNPDU BVLCOriginalBroadcastNPDU, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCOriginalBroadcastNPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCOriginalBroadcastNPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	npdu, err := ReadSimpleField[NPDU](ctx, "npdu", ReadComplex[NPDU](NPDUParseWithBufferProducer((uint16)(bvlcPayloadLength)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'npdu' field"))
	}
	m.Npdu = npdu

	if closeErr := readBuffer.CloseContext("BVLCOriginalBroadcastNPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCOriginalBroadcastNPDU")
	}

	return m, nil
}

func (m *_BVLCOriginalBroadcastNPDU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCOriginalBroadcastNPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCOriginalBroadcastNPDU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCOriginalBroadcastNPDU")
		}

		if err := WriteSimpleField[NPDU](ctx, "npdu", m.GetNpdu(), WriteComplex[NPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCOriginalBroadcastNPDU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCOriginalBroadcastNPDU")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCOriginalBroadcastNPDU) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCOriginalBroadcastNPDU) IsBVLCOriginalBroadcastNPDU() {}

func (m *_BVLCOriginalBroadcastNPDU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCOriginalBroadcastNPDU) deepCopy() *_BVLCOriginalBroadcastNPDU {
	if m == nil {
		return nil
	}
	_BVLCOriginalBroadcastNPDUCopy := &_BVLCOriginalBroadcastNPDU{
		m.BVLCContract.(*_BVLC).deepCopy(),
		utils.DeepCopy[NPDU](m.Npdu),
		m.BvlcPayloadLength,
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCOriginalBroadcastNPDUCopy
}

func (m *_BVLCOriginalBroadcastNPDU) String() string {
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
