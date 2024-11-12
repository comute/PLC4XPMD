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

// TelephonyDataDialInFailure is the corresponding interface of TelephonyDataDialInFailure
type TelephonyDataDialInFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() DialInFailureReason
	// IsTelephonyDataDialInFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataDialInFailure()
	// CreateBuilder creates a TelephonyDataDialInFailureBuilder
	CreateTelephonyDataDialInFailureBuilder() TelephonyDataDialInFailureBuilder
}

// _TelephonyDataDialInFailure is the data-structure of this message
type _TelephonyDataDialInFailure struct {
	TelephonyDataContract
	Reason DialInFailureReason
}

var _ TelephonyDataDialInFailure = (*_TelephonyDataDialInFailure)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataDialInFailure)(nil)

// NewTelephonyDataDialInFailure factory function for _TelephonyDataDialInFailure
func NewTelephonyDataDialInFailure(commandTypeContainer TelephonyCommandTypeContainer, argument byte, reason DialInFailureReason) *_TelephonyDataDialInFailure {
	_result := &_TelephonyDataDialInFailure{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Reason:                reason,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataDialInFailureBuilder is a builder for TelephonyDataDialInFailure
type TelephonyDataDialInFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reason DialInFailureReason) TelephonyDataDialInFailureBuilder
	// WithReason adds Reason (property field)
	WithReason(DialInFailureReason) TelephonyDataDialInFailureBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TelephonyDataBuilder
	// Build builds the TelephonyDataDialInFailure or returns an error if something is wrong
	Build() (TelephonyDataDialInFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataDialInFailure
}

// NewTelephonyDataDialInFailureBuilder() creates a TelephonyDataDialInFailureBuilder
func NewTelephonyDataDialInFailureBuilder() TelephonyDataDialInFailureBuilder {
	return &_TelephonyDataDialInFailureBuilder{_TelephonyDataDialInFailure: new(_TelephonyDataDialInFailure)}
}

type _TelephonyDataDialInFailureBuilder struct {
	*_TelephonyDataDialInFailure

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataDialInFailureBuilder) = (*_TelephonyDataDialInFailureBuilder)(nil)

func (b *_TelephonyDataDialInFailureBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
	contract.(*_TelephonyData)._SubType = b._TelephonyDataDialInFailure
}

func (b *_TelephonyDataDialInFailureBuilder) WithMandatoryFields(reason DialInFailureReason) TelephonyDataDialInFailureBuilder {
	return b.WithReason(reason)
}

func (b *_TelephonyDataDialInFailureBuilder) WithReason(reason DialInFailureReason) TelephonyDataDialInFailureBuilder {
	b.Reason = reason
	return b
}

func (b *_TelephonyDataDialInFailureBuilder) Build() (TelephonyDataDialInFailure, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataDialInFailure.deepCopy(), nil
}

func (b *_TelephonyDataDialInFailureBuilder) MustBuild() TelephonyDataDialInFailure {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TelephonyDataDialInFailureBuilder) Done() TelephonyDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTelephonyDataBuilder().(*_TelephonyDataBuilder)
	}
	return b.parentBuilder
}

func (b *_TelephonyDataDialInFailureBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataDialInFailureBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataDialInFailureBuilder().(*_TelephonyDataDialInFailureBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataDialInFailureBuilder creates a TelephonyDataDialInFailureBuilder
func (b *_TelephonyDataDialInFailure) CreateTelephonyDataDialInFailureBuilder() TelephonyDataDialInFailureBuilder {
	if b == nil {
		return NewTelephonyDataDialInFailureBuilder()
	}
	return &_TelephonyDataDialInFailureBuilder{_TelephonyDataDialInFailure: b.deepCopy()}
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

func (m *_TelephonyDataDialInFailure) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDialInFailure) GetReason() DialInFailureReason {
	return m.Reason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTelephonyDataDialInFailure(structType any) TelephonyDataDialInFailure {
	if casted, ok := structType.(TelephonyDataDialInFailure); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDialInFailure); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDialInFailure) GetTypeName() string {
	return "TelephonyDataDialInFailure"
}

func (m *_TelephonyDataDialInFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (reason)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyDataDialInFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataDialInFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataDialInFailure TelephonyDataDialInFailure, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDialInFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDialInFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reason, err := ReadEnumField[DialInFailureReason](ctx, "reason", "DialInFailureReason", ReadEnum(DialInFailureReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reason' field"))
	}
	m.Reason = reason

	if closeErr := readBuffer.CloseContext("TelephonyDataDialInFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDialInFailure")
	}

	return m, nil
}

func (m *_TelephonyDataDialInFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataDialInFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDialInFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDialInFailure")
		}

		if err := WriteSimpleEnumField[DialInFailureReason](ctx, "reason", "DialInFailureReason", m.GetReason(), WriteEnum[DialInFailureReason, uint8](DialInFailureReason.GetValue, DialInFailureReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDialInFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDialInFailure")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataDialInFailure) IsTelephonyDataDialInFailure() {}

func (m *_TelephonyDataDialInFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataDialInFailure) deepCopy() *_TelephonyDataDialInFailure {
	if m == nil {
		return nil
	}
	_TelephonyDataDialInFailureCopy := &_TelephonyDataDialInFailure{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
		m.Reason,
	}
	m.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataDialInFailureCopy
}

func (m *_TelephonyDataDialInFailure) String() string {
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
