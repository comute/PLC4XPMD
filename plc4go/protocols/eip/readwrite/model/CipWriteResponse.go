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

// CipWriteResponse is the corresponding interface of CipWriteResponse
type CipWriteResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// IsCipWriteResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipWriteResponse()
	// CreateBuilder creates a CipWriteResponseBuilder
	CreateCipWriteResponseBuilder() CipWriteResponseBuilder
}

// _CipWriteResponse is the data-structure of this message
type _CipWriteResponse struct {
	CipServiceContract
	Status    uint8
	ExtStatus uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ CipWriteResponse = (*_CipWriteResponse)(nil)
var _ CipServiceRequirements = (*_CipWriteResponse)(nil)

// NewCipWriteResponse factory function for _CipWriteResponse
func NewCipWriteResponse(status uint8, extStatus uint8, serviceLen uint16) *_CipWriteResponse {
	_result := &_CipWriteResponse{
		CipServiceContract: NewCipService(serviceLen),
		Status:             status,
		ExtStatus:          extStatus,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipWriteResponseBuilder is a builder for CipWriteResponse
type CipWriteResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status uint8, extStatus uint8) CipWriteResponseBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint8) CipWriteResponseBuilder
	// WithExtStatus adds ExtStatus (property field)
	WithExtStatus(uint8) CipWriteResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the CipWriteResponse or returns an error if something is wrong
	Build() (CipWriteResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipWriteResponse
}

// NewCipWriteResponseBuilder() creates a CipWriteResponseBuilder
func NewCipWriteResponseBuilder() CipWriteResponseBuilder {
	return &_CipWriteResponseBuilder{_CipWriteResponse: new(_CipWriteResponse)}
}

type _CipWriteResponseBuilder struct {
	*_CipWriteResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipWriteResponseBuilder) = (*_CipWriteResponseBuilder)(nil)

func (b *_CipWriteResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._CipWriteResponse
}

func (b *_CipWriteResponseBuilder) WithMandatoryFields(status uint8, extStatus uint8) CipWriteResponseBuilder {
	return b.WithStatus(status).WithExtStatus(extStatus)
}

func (b *_CipWriteResponseBuilder) WithStatus(status uint8) CipWriteResponseBuilder {
	b.Status = status
	return b
}

func (b *_CipWriteResponseBuilder) WithExtStatus(extStatus uint8) CipWriteResponseBuilder {
	b.ExtStatus = extStatus
	return b
}

func (b *_CipWriteResponseBuilder) Build() (CipWriteResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipWriteResponse.deepCopy(), nil
}

func (b *_CipWriteResponseBuilder) MustBuild() CipWriteResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipWriteResponseBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_CipWriteResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipWriteResponseBuilder) DeepCopy() any {
	_copy := b.CreateCipWriteResponseBuilder().(*_CipWriteResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipWriteResponseBuilder creates a CipWriteResponseBuilder
func (b *_CipWriteResponse) CreateCipWriteResponseBuilder() CipWriteResponseBuilder {
	if b == nil {
		return NewCipWriteResponseBuilder()
	}
	return &_CipWriteResponseBuilder{_CipWriteResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipWriteResponse) GetService() uint8 {
	return 0x4D
}

func (m *_CipWriteResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipWriteResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipWriteResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipWriteResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipWriteResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipWriteResponse(structType any) CipWriteResponse {
	if casted, ok := structType.(CipWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipWriteResponse) GetTypeName() string {
	return "CipWriteResponse"
}

func (m *_CipWriteResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipWriteResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipWriteResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipWriteResponse CipWriteResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	extStatus, err := ReadSimpleField(ctx, "extStatus", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extStatus' field"))
	}
	m.ExtStatus = extStatus

	if closeErr := readBuffer.CloseContext("CipWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipWriteResponse")
	}

	return m, nil
}

func (m *_CipWriteResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipWriteResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipWriteResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "extStatus", m.GetExtStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'extStatus' field")
		}

		if popErr := writeBuffer.PopContext("CipWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipWriteResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipWriteResponse) IsCipWriteResponse() {}

func (m *_CipWriteResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipWriteResponse) deepCopy() *_CipWriteResponse {
	if m == nil {
		return nil
	}
	_CipWriteResponseCopy := &_CipWriteResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.Status,
		m.ExtStatus,
		m.reservedField0,
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _CipWriteResponseCopy
}

func (m *_CipWriteResponse) String() string {
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
