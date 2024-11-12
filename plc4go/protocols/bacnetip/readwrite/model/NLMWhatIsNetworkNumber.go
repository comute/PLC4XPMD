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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NLMWhatIsNetworkNumber is the corresponding interface of NLMWhatIsNetworkNumber
type NLMWhatIsNetworkNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// IsNLMWhatIsNetworkNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMWhatIsNetworkNumber()
	// CreateBuilder creates a NLMWhatIsNetworkNumberBuilder
	CreateNLMWhatIsNetworkNumberBuilder() NLMWhatIsNetworkNumberBuilder
}

// _NLMWhatIsNetworkNumber is the data-structure of this message
type _NLMWhatIsNetworkNumber struct {
	NLMContract
}

var _ NLMWhatIsNetworkNumber = (*_NLMWhatIsNetworkNumber)(nil)
var _ NLMRequirements = (*_NLMWhatIsNetworkNumber)(nil)

// NewNLMWhatIsNetworkNumber factory function for _NLMWhatIsNetworkNumber
func NewNLMWhatIsNetworkNumber(apduLength uint16) *_NLMWhatIsNetworkNumber {
	_result := &_NLMWhatIsNetworkNumber{
		NLMContract: NewNLM(apduLength),
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMWhatIsNetworkNumberBuilder is a builder for NLMWhatIsNetworkNumber
type NLMWhatIsNetworkNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NLMWhatIsNetworkNumberBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NLMBuilder
	// Build builds the NLMWhatIsNetworkNumber or returns an error if something is wrong
	Build() (NLMWhatIsNetworkNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMWhatIsNetworkNumber
}

// NewNLMWhatIsNetworkNumberBuilder() creates a NLMWhatIsNetworkNumberBuilder
func NewNLMWhatIsNetworkNumberBuilder() NLMWhatIsNetworkNumberBuilder {
	return &_NLMWhatIsNetworkNumberBuilder{_NLMWhatIsNetworkNumber: new(_NLMWhatIsNetworkNumber)}
}

type _NLMWhatIsNetworkNumberBuilder struct {
	*_NLMWhatIsNetworkNumber

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMWhatIsNetworkNumberBuilder) = (*_NLMWhatIsNetworkNumberBuilder)(nil)

func (b *_NLMWhatIsNetworkNumberBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
	contract.(*_NLM)._SubType = b._NLMWhatIsNetworkNumber
}

func (b *_NLMWhatIsNetworkNumberBuilder) WithMandatoryFields() NLMWhatIsNetworkNumberBuilder {
	return b
}

func (b *_NLMWhatIsNetworkNumberBuilder) Build() (NLMWhatIsNetworkNumber, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMWhatIsNetworkNumber.deepCopy(), nil
}

func (b *_NLMWhatIsNetworkNumberBuilder) MustBuild() NLMWhatIsNetworkNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMWhatIsNetworkNumberBuilder) Done() NLMBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNLMBuilder().(*_NLMBuilder)
	}
	return b.parentBuilder
}

func (b *_NLMWhatIsNetworkNumberBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMWhatIsNetworkNumberBuilder) DeepCopy() any {
	_copy := b.CreateNLMWhatIsNetworkNumberBuilder().(*_NLMWhatIsNetworkNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMWhatIsNetworkNumberBuilder creates a NLMWhatIsNetworkNumberBuilder
func (b *_NLMWhatIsNetworkNumber) CreateNLMWhatIsNetworkNumberBuilder() NLMWhatIsNetworkNumberBuilder {
	if b == nil {
		return NewNLMWhatIsNetworkNumberBuilder()
	}
	return &_NLMWhatIsNetworkNumberBuilder{_NLMWhatIsNetworkNumber: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMWhatIsNetworkNumber) GetMessageType() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMWhatIsNetworkNumber) GetParent() NLMContract {
	return m.NLMContract
}

// Deprecated: use the interface for direct cast
func CastNLMWhatIsNetworkNumber(structType any) NLMWhatIsNetworkNumber {
	if casted, ok := structType.(NLMWhatIsNetworkNumber); ok {
		return casted
	}
	if casted, ok := structType.(*NLMWhatIsNetworkNumber); ok {
		return *casted
	}
	return nil
}

func (m *_NLMWhatIsNetworkNumber) GetTypeName() string {
	return "NLMWhatIsNetworkNumber"
}

func (m *_NLMWhatIsNetworkNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NLMWhatIsNetworkNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMWhatIsNetworkNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMWhatIsNetworkNumber NLMWhatIsNetworkNumber, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMWhatIsNetworkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMWhatIsNetworkNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NLMWhatIsNetworkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMWhatIsNetworkNumber")
	}

	return m, nil
}

func (m *_NLMWhatIsNetworkNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMWhatIsNetworkNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMWhatIsNetworkNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMWhatIsNetworkNumber")
		}

		if popErr := writeBuffer.PopContext("NLMWhatIsNetworkNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMWhatIsNetworkNumber")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMWhatIsNetworkNumber) IsNLMWhatIsNetworkNumber() {}

func (m *_NLMWhatIsNetworkNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMWhatIsNetworkNumber) deepCopy() *_NLMWhatIsNetworkNumber {
	if m == nil {
		return nil
	}
	_NLMWhatIsNetworkNumberCopy := &_NLMWhatIsNetworkNumber{
		m.NLMContract.(*_NLM).deepCopy(),
	}
	m.NLMContract.(*_NLM)._SubType = m
	return _NLMWhatIsNetworkNumberCopy
}

func (m *_NLMWhatIsNetworkNumber) String() string {
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
