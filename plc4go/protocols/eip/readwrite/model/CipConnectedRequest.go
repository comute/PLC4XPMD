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

// CipConnectedRequest is the corresponding interface of CipConnectedRequest
type CipConnectedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetPathSegments returns PathSegments (property field)
	GetPathSegments() []byte
	// IsCipConnectedRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipConnectedRequest()
	// CreateBuilder creates a CipConnectedRequestBuilder
	CreateCipConnectedRequestBuilder() CipConnectedRequestBuilder
}

// _CipConnectedRequest is the data-structure of this message
type _CipConnectedRequest struct {
	CipServiceContract
	PathSegments []byte
	// Reserved Fields
	reservedField0 *uint16
	reservedField1 *uint32
}

var _ CipConnectedRequest = (*_CipConnectedRequest)(nil)
var _ CipServiceRequirements = (*_CipConnectedRequest)(nil)

// NewCipConnectedRequest factory function for _CipConnectedRequest
func NewCipConnectedRequest(pathSegments []byte, serviceLen uint16) *_CipConnectedRequest {
	_result := &_CipConnectedRequest{
		CipServiceContract: NewCipService(serviceLen),
		PathSegments:       pathSegments,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipConnectedRequestBuilder is a builder for CipConnectedRequest
type CipConnectedRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pathSegments []byte) CipConnectedRequestBuilder
	// WithPathSegments adds PathSegments (property field)
	WithPathSegments(...byte) CipConnectedRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the CipConnectedRequest or returns an error if something is wrong
	Build() (CipConnectedRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipConnectedRequest
}

// NewCipConnectedRequestBuilder() creates a CipConnectedRequestBuilder
func NewCipConnectedRequestBuilder() CipConnectedRequestBuilder {
	return &_CipConnectedRequestBuilder{_CipConnectedRequest: new(_CipConnectedRequest)}
}

type _CipConnectedRequestBuilder struct {
	*_CipConnectedRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipConnectedRequestBuilder) = (*_CipConnectedRequestBuilder)(nil)

func (b *_CipConnectedRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._CipConnectedRequest
}

func (b *_CipConnectedRequestBuilder) WithMandatoryFields(pathSegments []byte) CipConnectedRequestBuilder {
	return b.WithPathSegments(pathSegments...)
}

func (b *_CipConnectedRequestBuilder) WithPathSegments(pathSegments ...byte) CipConnectedRequestBuilder {
	b.PathSegments = pathSegments
	return b
}

func (b *_CipConnectedRequestBuilder) Build() (CipConnectedRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipConnectedRequest.deepCopy(), nil
}

func (b *_CipConnectedRequestBuilder) MustBuild() CipConnectedRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipConnectedRequestBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_CipConnectedRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipConnectedRequestBuilder) DeepCopy() any {
	_copy := b.CreateCipConnectedRequestBuilder().(*_CipConnectedRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipConnectedRequestBuilder creates a CipConnectedRequestBuilder
func (b *_CipConnectedRequest) CreateCipConnectedRequestBuilder() CipConnectedRequestBuilder {
	if b == nil {
		return NewCipConnectedRequestBuilder()
	}
	return &_CipConnectedRequestBuilder{_CipConnectedRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectedRequest) GetService() uint8 {
	return 0x52
}

func (m *_CipConnectedRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectedRequest) GetConnected() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectedRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectedRequest) GetPathSegments() []byte {
	return m.PathSegments
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipConnectedRequest(structType any) CipConnectedRequest {
	if casted, ok := structType.(CipConnectedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectedRequest) GetTypeName() string {
	return "CipConnectedRequest"
}

func (m *_CipConnectedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.PathSegments) > 0 {
		lengthInBits += 8 * uint16(len(m.PathSegments))
	}

	// Reserved Field (reserved)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CipConnectedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectedRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectedRequest CipConnectedRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	pathSegments, err := readBuffer.ReadByteArray("pathSegments", int(int32(requestPathSize)*int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pathSegments' field"))
	}
	m.PathSegments = pathSegments

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0001))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(32)), uint32(0x00000000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	if closeErr := readBuffer.CloseContext("CipConnectedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectedRequest")
	}

	return m, nil
}

func (m *_CipConnectedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectedRequest")
		}
		requestPathSize := uint8(uint8(uint8(len(m.GetPathSegments()))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteByteArrayField(ctx, "pathSegments", m.GetPathSegments(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'pathSegments' field")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0001), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteReservedField[uint32](ctx, "reserved", uint32(0x00000000), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("CipConnectedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectedRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectedRequest) IsCipConnectedRequest() {}

func (m *_CipConnectedRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipConnectedRequest) deepCopy() *_CipConnectedRequest {
	if m == nil {
		return nil
	}
	_CipConnectedRequestCopy := &_CipConnectedRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.PathSegments),
		m.reservedField0,
		m.reservedField1,
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _CipConnectedRequestCopy
}

func (m *_CipConnectedRequest) String() string {
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
