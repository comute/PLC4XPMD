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

// S7PayloadUserData is the corresponding interface of S7PayloadUserData
type S7PayloadUserData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7PayloadUserDataItem
	// IsS7PayloadUserData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserData()
	// CreateBuilder creates a S7PayloadUserDataBuilder
	CreateS7PayloadUserDataBuilder() S7PayloadUserDataBuilder
}

// _S7PayloadUserData is the data-structure of this message
type _S7PayloadUserData struct {
	S7PayloadContract
	Items []S7PayloadUserDataItem
}

var _ S7PayloadUserData = (*_S7PayloadUserData)(nil)
var _ S7PayloadRequirements = (*_S7PayloadUserData)(nil)

// NewS7PayloadUserData factory function for _S7PayloadUserData
func NewS7PayloadUserData(items []S7PayloadUserDataItem, parameter S7Parameter) *_S7PayloadUserData {
	_result := &_S7PayloadUserData{
		S7PayloadContract: NewS7Payload(parameter),
		Items:             items,
	}
	_result.S7PayloadContract.(*_S7Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataBuilder is a builder for S7PayloadUserData
type S7PayloadUserDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []S7PayloadUserDataItem) S7PayloadUserDataBuilder
	// WithItems adds Items (property field)
	WithItems(...S7PayloadUserDataItem) S7PayloadUserDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadBuilder
	// Build builds the S7PayloadUserData or returns an error if something is wrong
	Build() (S7PayloadUserData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserData
}

// NewS7PayloadUserDataBuilder() creates a S7PayloadUserDataBuilder
func NewS7PayloadUserDataBuilder() S7PayloadUserDataBuilder {
	return &_S7PayloadUserDataBuilder{_S7PayloadUserData: new(_S7PayloadUserData)}
}

type _S7PayloadUserDataBuilder struct {
	*_S7PayloadUserData

	parentBuilder *_S7PayloadBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataBuilder) = (*_S7PayloadUserDataBuilder)(nil)

func (b *_S7PayloadUserDataBuilder) setParent(contract S7PayloadContract) {
	b.S7PayloadContract = contract
	contract.(*_S7Payload)._SubType = b._S7PayloadUserData
}

func (b *_S7PayloadUserDataBuilder) WithMandatoryFields(items []S7PayloadUserDataItem) S7PayloadUserDataBuilder {
	return b.WithItems(items...)
}

func (b *_S7PayloadUserDataBuilder) WithItems(items ...S7PayloadUserDataItem) S7PayloadUserDataBuilder {
	b.Items = items
	return b
}

func (b *_S7PayloadUserDataBuilder) Build() (S7PayloadUserData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserData.deepCopy(), nil
}

func (b *_S7PayloadUserDataBuilder) MustBuild() S7PayloadUserData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadUserDataBuilder) Done() S7PayloadBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadBuilder().(*_S7PayloadBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadUserDataBuilder) buildForS7Payload() (S7Payload, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataBuilder().(*_S7PayloadUserDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataBuilder creates a S7PayloadUserDataBuilder
func (b *_S7PayloadUserData) CreateS7PayloadUserDataBuilder() S7PayloadUserDataBuilder {
	if b == nil {
		return NewS7PayloadUserDataBuilder()
	}
	return &_S7PayloadUserDataBuilder{_S7PayloadUserData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserData) GetParameterParameterType() uint8 {
	return 0x00
}

func (m *_S7PayloadUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserData) GetParent() S7PayloadContract {
	return m.S7PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserData) GetItems() []S7PayloadUserDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadUserData(structType any) S7PayloadUserData {
	if casted, ok := structType.(S7PayloadUserData); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserData); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserData) GetTypeName() string {
	return "S7PayloadUserData"
}

func (m *_S7PayloadUserData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadContract.(*_S7Payload).getLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Payload, messageType uint8, parameter S7Parameter) (__s7PayloadUserData S7PayloadUserData, err error) {
	m.S7PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := ReadCountArrayField[S7PayloadUserDataItem](ctx, "items", ReadComplex[S7PayloadUserDataItem](S7PayloadUserDataItemParseWithBufferProducer[S7PayloadUserDataItem]((uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuFunctionGroup()), (uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuFunctionType()), (uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuSubfunction())), readBuffer), uint64(int32(len(CastS7ParameterUserData(parameter).GetItems()))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserData")
	}

	return m, nil
}

func (m *_S7PayloadUserData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserData")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserData")
		}
		return nil
	}
	return m.S7PayloadContract.(*_S7Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserData) IsS7PayloadUserData() {}

func (m *_S7PayloadUserData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserData) deepCopy() *_S7PayloadUserData {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataCopy := &_S7PayloadUserData{
		m.S7PayloadContract.(*_S7Payload).deepCopy(),
		utils.DeepCopySlice[S7PayloadUserDataItem, S7PayloadUserDataItem](m.Items),
	}
	m.S7PayloadContract.(*_S7Payload)._SubType = m
	return _S7PayloadUserDataCopy
}

func (m *_S7PayloadUserData) String() string {
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
