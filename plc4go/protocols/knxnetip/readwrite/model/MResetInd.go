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

// MResetInd is the corresponding interface of MResetInd
type MResetInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMResetInd is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMResetInd()
	// CreateBuilder creates a MResetIndBuilder
	CreateMResetIndBuilder() MResetIndBuilder
}

// _MResetInd is the data-structure of this message
type _MResetInd struct {
	CEMIContract
}

var _ MResetInd = (*_MResetInd)(nil)
var _ CEMIRequirements = (*_MResetInd)(nil)

// NewMResetInd factory function for _MResetInd
func NewMResetInd(size uint16) *_MResetInd {
	_result := &_MResetInd{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MResetIndBuilder is a builder for MResetInd
type MResetIndBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MResetIndBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the MResetInd or returns an error if something is wrong
	Build() (MResetInd, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MResetInd
}

// NewMResetIndBuilder() creates a MResetIndBuilder
func NewMResetIndBuilder() MResetIndBuilder {
	return &_MResetIndBuilder{_MResetInd: new(_MResetInd)}
}

type _MResetIndBuilder struct {
	*_MResetInd

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MResetIndBuilder) = (*_MResetIndBuilder)(nil)

func (b *_MResetIndBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._MResetInd
}

func (b *_MResetIndBuilder) WithMandatoryFields() MResetIndBuilder {
	return b
}

func (b *_MResetIndBuilder) Build() (MResetInd, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MResetInd.deepCopy(), nil
}

func (b *_MResetIndBuilder) MustBuild() MResetInd {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MResetIndBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_MResetIndBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MResetIndBuilder) DeepCopy() any {
	_copy := b.CreateMResetIndBuilder().(*_MResetIndBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMResetIndBuilder creates a MResetIndBuilder
func (b *_MResetInd) CreateMResetIndBuilder() MResetIndBuilder {
	if b == nil {
		return NewMResetIndBuilder()
	}
	return &_MResetIndBuilder{_MResetInd: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MResetInd) GetMessageCode() uint8 {
	return 0xF0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MResetInd) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMResetInd(structType any) MResetInd {
	if casted, ok := structType.(MResetInd); ok {
		return casted
	}
	if casted, ok := structType.(*MResetInd); ok {
		return *casted
	}
	return nil
}

func (m *_MResetInd) GetTypeName() string {
	return "MResetInd"
}

func (m *_MResetInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MResetInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MResetInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mResetInd MResetInd, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MResetInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MResetInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MResetInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MResetInd")
	}

	return m, nil
}

func (m *_MResetInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MResetInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MResetInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MResetInd")
		}

		if popErr := writeBuffer.PopContext("MResetInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MResetInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MResetInd) IsMResetInd() {}

func (m *_MResetInd) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MResetInd) deepCopy() *_MResetInd {
	if m == nil {
		return nil
	}
	_MResetIndCopy := &_MResetInd{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MResetIndCopy
}

func (m *_MResetInd) String() string {
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
