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

// TelephonyDataRecallLastNumberRequest is the corresponding interface of TelephonyDataRecallLastNumberRequest
type TelephonyDataRecallLastNumberRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// GetRecallLastNumberType returns RecallLastNumberType (property field)
	GetRecallLastNumberType() byte
	// GetIsNumberOfLastOutgoingCall returns IsNumberOfLastOutgoingCall (virtual field)
	GetIsNumberOfLastOutgoingCall() bool
	// GetIsNumberOfLastIncomingCall returns IsNumberOfLastIncomingCall (virtual field)
	GetIsNumberOfLastIncomingCall() bool
	// IsTelephonyDataRecallLastNumberRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataRecallLastNumberRequest()
	// CreateBuilder creates a TelephonyDataRecallLastNumberRequestBuilder
	CreateTelephonyDataRecallLastNumberRequestBuilder() TelephonyDataRecallLastNumberRequestBuilder
}

// _TelephonyDataRecallLastNumberRequest is the data-structure of this message
type _TelephonyDataRecallLastNumberRequest struct {
	TelephonyDataContract
	RecallLastNumberType byte
}

var _ TelephonyDataRecallLastNumberRequest = (*_TelephonyDataRecallLastNumberRequest)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataRecallLastNumberRequest)(nil)

// NewTelephonyDataRecallLastNumberRequest factory function for _TelephonyDataRecallLastNumberRequest
func NewTelephonyDataRecallLastNumberRequest(commandTypeContainer TelephonyCommandTypeContainer, argument byte, recallLastNumberType byte) *_TelephonyDataRecallLastNumberRequest {
	_result := &_TelephonyDataRecallLastNumberRequest{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		RecallLastNumberType:  recallLastNumberType,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataRecallLastNumberRequestBuilder is a builder for TelephonyDataRecallLastNumberRequest
type TelephonyDataRecallLastNumberRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recallLastNumberType byte) TelephonyDataRecallLastNumberRequestBuilder
	// WithRecallLastNumberType adds RecallLastNumberType (property field)
	WithRecallLastNumberType(byte) TelephonyDataRecallLastNumberRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TelephonyDataBuilder
	// Build builds the TelephonyDataRecallLastNumberRequest or returns an error if something is wrong
	Build() (TelephonyDataRecallLastNumberRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataRecallLastNumberRequest
}

// NewTelephonyDataRecallLastNumberRequestBuilder() creates a TelephonyDataRecallLastNumberRequestBuilder
func NewTelephonyDataRecallLastNumberRequestBuilder() TelephonyDataRecallLastNumberRequestBuilder {
	return &_TelephonyDataRecallLastNumberRequestBuilder{_TelephonyDataRecallLastNumberRequest: new(_TelephonyDataRecallLastNumberRequest)}
}

type _TelephonyDataRecallLastNumberRequestBuilder struct {
	*_TelephonyDataRecallLastNumberRequest

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataRecallLastNumberRequestBuilder) = (*_TelephonyDataRecallLastNumberRequestBuilder)(nil)

func (b *_TelephonyDataRecallLastNumberRequestBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
	contract.(*_TelephonyData)._SubType = b._TelephonyDataRecallLastNumberRequest
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) WithMandatoryFields(recallLastNumberType byte) TelephonyDataRecallLastNumberRequestBuilder {
	return b.WithRecallLastNumberType(recallLastNumberType)
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) WithRecallLastNumberType(recallLastNumberType byte) TelephonyDataRecallLastNumberRequestBuilder {
	b.RecallLastNumberType = recallLastNumberType
	return b
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) Build() (TelephonyDataRecallLastNumberRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataRecallLastNumberRequest.deepCopy(), nil
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) MustBuild() TelephonyDataRecallLastNumberRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) Done() TelephonyDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTelephonyDataBuilder().(*_TelephonyDataBuilder)
	}
	return b.parentBuilder
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataRecallLastNumberRequestBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataRecallLastNumberRequestBuilder().(*_TelephonyDataRecallLastNumberRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataRecallLastNumberRequestBuilder creates a TelephonyDataRecallLastNumberRequestBuilder
func (b *_TelephonyDataRecallLastNumberRequest) CreateTelephonyDataRecallLastNumberRequestBuilder() TelephonyDataRecallLastNumberRequestBuilder {
	if b == nil {
		return NewTelephonyDataRecallLastNumberRequestBuilder()
	}
	return &_TelephonyDataRecallLastNumberRequestBuilder{_TelephonyDataRecallLastNumberRequest: b.deepCopy()}
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

func (m *_TelephonyDataRecallLastNumberRequest) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetRecallLastNumberType() byte {
	return m.RecallLastNumberType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastOutgoingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x01)))
}

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastIncomingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x02)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTelephonyDataRecallLastNumberRequest(structType any) TelephonyDataRecallLastNumberRequest {
	if casted, ok := structType.(TelephonyDataRecallLastNumberRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRecallLastNumberRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRecallLastNumberRequest) GetTypeName() string {
	return "TelephonyDataRecallLastNumberRequest"
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (recallLastNumberType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataRecallLastNumberRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataRecallLastNumberRequest TelephonyDataRecallLastNumberRequest, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRecallLastNumberRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRecallLastNumberRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recallLastNumberType, err := ReadSimpleField(ctx, "recallLastNumberType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recallLastNumberType' field"))
	}
	m.RecallLastNumberType = recallLastNumberType

	isNumberOfLastOutgoingCall, err := ReadVirtualField[bool](ctx, "isNumberOfLastOutgoingCall", (*bool)(nil), bool((recallLastNumberType) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNumberOfLastOutgoingCall' field"))
	}
	_ = isNumberOfLastOutgoingCall

	isNumberOfLastIncomingCall, err := ReadVirtualField[bool](ctx, "isNumberOfLastIncomingCall", (*bool)(nil), bool((recallLastNumberType) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNumberOfLastIncomingCall' field"))
	}
	_ = isNumberOfLastIncomingCall

	if closeErr := readBuffer.CloseContext("TelephonyDataRecallLastNumberRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRecallLastNumberRequest")
	}

	return m, nil
}

func (m *_TelephonyDataRecallLastNumberRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRecallLastNumberRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRecallLastNumberRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRecallLastNumberRequest")
		}

		if err := WriteSimpleField[byte](ctx, "recallLastNumberType", m.GetRecallLastNumberType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'recallLastNumberType' field")
		}
		// Virtual field
		isNumberOfLastOutgoingCall := m.GetIsNumberOfLastOutgoingCall()
		_ = isNumberOfLastOutgoingCall
		if _isNumberOfLastOutgoingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastOutgoingCall", m.GetIsNumberOfLastOutgoingCall()); _isNumberOfLastOutgoingCallErr != nil {
			return errors.Wrap(_isNumberOfLastOutgoingCallErr, "Error serializing 'isNumberOfLastOutgoingCall' field")
		}
		// Virtual field
		isNumberOfLastIncomingCall := m.GetIsNumberOfLastIncomingCall()
		_ = isNumberOfLastIncomingCall
		if _isNumberOfLastIncomingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastIncomingCall", m.GetIsNumberOfLastIncomingCall()); _isNumberOfLastIncomingCallErr != nil {
			return errors.Wrap(_isNumberOfLastIncomingCallErr, "Error serializing 'isNumberOfLastIncomingCall' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRecallLastNumberRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRecallLastNumberRequest")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataRecallLastNumberRequest) IsTelephonyDataRecallLastNumberRequest() {}

func (m *_TelephonyDataRecallLastNumberRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataRecallLastNumberRequest) deepCopy() *_TelephonyDataRecallLastNumberRequest {
	if m == nil {
		return nil
	}
	_TelephonyDataRecallLastNumberRequestCopy := &_TelephonyDataRecallLastNumberRequest{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
		m.RecallLastNumberType,
	}
	m.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataRecallLastNumberRequestCopy
}

func (m *_TelephonyDataRecallLastNumberRequest) String() string {
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
