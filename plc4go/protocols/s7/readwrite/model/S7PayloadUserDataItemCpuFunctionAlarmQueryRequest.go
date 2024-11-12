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

// Constant values.
const S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_FUNCTIONID uint8 = 0x00
const S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_NUMBERMESSAGEOBJ uint8 = 0x01
const S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_VARIABLESPEC uint8 = 0x12
const S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_LENGTH uint8 = 0x08

// S7PayloadUserDataItemCpuFunctionAlarmQueryRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
type S7PayloadUserDataItemCpuFunctionAlarmQueryRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetQueryType returns QueryType (property field)
	GetQueryType() QueryType
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// IsS7PayloadUserDataItemCpuFunctionAlarmQueryRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionAlarmQueryRequest()
	// CreateBuilder creates a S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
	CreateS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder() S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
}

// _S7PayloadUserDataItemCpuFunctionAlarmQueryRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmQueryRequest struct {
	S7PayloadUserDataItemContract
	SyntaxId  SyntaxIdType
	QueryType QueryType
	AlarmType AlarmType
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ S7PayloadUserDataItemCpuFunctionAlarmQueryRequest = (*_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest)(nil)

// NewS7PayloadUserDataItemCpuFunctionAlarmQueryRequest factory function for _S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
func NewS7PayloadUserDataItemCpuFunctionAlarmQueryRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		SyntaxId:                      syntaxId,
		QueryType:                     queryType,
		AlarmType:                     alarmType,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder is a builder for S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
type S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
	// WithSyntaxId adds SyntaxId (property field)
	WithSyntaxId(SyntaxIdType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
	// WithQueryType adds QueryType (property field)
	WithQueryType(QueryType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
	// WithAlarmType adds AlarmType (property field)
	WithAlarmType(AlarmType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadUserDataItemBuilder
	// Build builds the S7PayloadUserDataItemCpuFunctionAlarmQueryRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemCpuFunctionAlarmQueryRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
}

// NewS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder() creates a S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
func NewS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder() S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	return &_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder{_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest: new(_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest)}
}

type _S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder struct {
	*_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) = (*_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder)(nil)

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
	contract.(*_S7PayloadUserDataItem)._SubType = b._S7PayloadUserDataItemCpuFunctionAlarmQueryRequest
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) WithMandatoryFields(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	return b.WithSyntaxId(syntaxId).WithQueryType(queryType).WithAlarmType(alarmType)
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) WithSyntaxId(syntaxId SyntaxIdType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	b.SyntaxId = syntaxId
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) WithQueryType(queryType QueryType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	b.QueryType = queryType
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) WithAlarmType(alarmType AlarmType) S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	b.AlarmType = alarmType
	return b
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) Build() (S7PayloadUserDataItemCpuFunctionAlarmQueryRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemCpuFunctionAlarmQueryRequest.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) MustBuild() S7PayloadUserDataItemCpuFunctionAlarmQueryRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) Done() S7PayloadUserDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadUserDataItemBuilder().(*_S7PayloadUserDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder().(*_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder creates a S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder
func (b *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) CreateS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder() S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder()
	}
	return &_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestBuilder{_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetCpuSubfunction() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetQueryType() QueryType {
	return m.QueryType
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetAlarmType() AlarmType {
	return m.AlarmType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetFunctionId() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_FUNCTIONID
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetNumberMessageObj() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_NUMBERMESSAGEOBJ
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetVariableSpec() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_VARIABLESPEC
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetLength() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmQueryRequest(structType any) S7PayloadUserDataItemCpuFunctionAlarmQueryRequest {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmQueryRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmQueryRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmQueryRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	// Const Field (variableSpec)
	lengthInBits += 8

	// Const Field (length)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (queryType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionAlarmQueryRequest S7PayloadUserDataItemCpuFunctionAlarmQueryRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmQueryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmQueryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadConstField[uint8](ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)), S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_FUNCTIONID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	_ = functionId

	numberMessageObj, err := ReadConstField[uint8](ctx, "numberMessageObj", ReadUnsignedByte(readBuffer, uint8(8)), S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_NUMBERMESSAGEOBJ)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberMessageObj' field"))
	}
	_ = numberMessageObj

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	length, err := ReadConstField[uint8](ctx, "length", ReadUnsignedByte(readBuffer, uint8(8)), S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_LENGTH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	syntaxId, err := ReadEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", ReadEnum(SyntaxIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}
	m.SyntaxId = syntaxId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	queryType, err := ReadEnumField[QueryType](ctx, "queryType", "QueryType", ReadEnum(QueryTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queryType' field"))
	}
	m.QueryType = queryType

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x34))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	alarmType, err := ReadEnumField[AlarmType](ctx, "alarmType", "AlarmType", ReadEnum(AlarmTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmType' field"))
	}
	m.AlarmType = alarmType

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmQueryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmQueryRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmQueryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmQueryRequest")
		}

		if err := WriteConstField(ctx, "functionId", S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_FUNCTIONID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'functionId' field")
		}

		if err := WriteConstField(ctx, "numberMessageObj", S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_NUMBERMESSAGEOBJ, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberMessageObj' field")
		}

		if err := WriteConstField(ctx, "variableSpec", S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'variableSpec' field")
		}

		if err := WriteConstField(ctx, "length", S7PayloadUserDataItemCpuFunctionAlarmQueryRequest_LENGTH, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", m.GetSyntaxId(), WriteEnum[SyntaxIdType, uint8](SyntaxIdType.GetValue, SyntaxIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'syntaxId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleEnumField[QueryType](ctx, "queryType", "QueryType", m.GetQueryType(), WriteEnum[QueryType, uint8](QueryType.GetValue, QueryType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'queryType' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x34), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleEnumField[AlarmType](ctx, "alarmType", "AlarmType", m.GetAlarmType(), WriteEnum[AlarmType, uint8](AlarmType.GetValue, AlarmType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmType' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmQueryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmQueryRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) IsS7PayloadUserDataItemCpuFunctionAlarmQueryRequest() {
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) deepCopy() *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemCpuFunctionAlarmQueryRequestCopy := &_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.SyntaxId,
		m.QueryType,
		m.AlarmType,
		m.reservedField0,
		m.reservedField1,
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemCpuFunctionAlarmQueryRequestCopy
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryRequest) String() string {
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
