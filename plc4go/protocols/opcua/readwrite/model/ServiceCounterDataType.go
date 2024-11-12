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

// ServiceCounterDataType is the corresponding interface of ServiceCounterDataType
type ServiceCounterDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetTotalCount returns TotalCount (property field)
	GetTotalCount() uint32
	// GetErrorCount returns ErrorCount (property field)
	GetErrorCount() uint32
	// IsServiceCounterDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServiceCounterDataType()
	// CreateBuilder creates a ServiceCounterDataTypeBuilder
	CreateServiceCounterDataTypeBuilder() ServiceCounterDataTypeBuilder
}

// _ServiceCounterDataType is the data-structure of this message
type _ServiceCounterDataType struct {
	ExtensionObjectDefinitionContract
	TotalCount uint32
	ErrorCount uint32
}

var _ ServiceCounterDataType = (*_ServiceCounterDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServiceCounterDataType)(nil)

// NewServiceCounterDataType factory function for _ServiceCounterDataType
func NewServiceCounterDataType(totalCount uint32, errorCount uint32) *_ServiceCounterDataType {
	_result := &_ServiceCounterDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TotalCount:                        totalCount,
		ErrorCount:                        errorCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ServiceCounterDataTypeBuilder is a builder for ServiceCounterDataType
type ServiceCounterDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(totalCount uint32, errorCount uint32) ServiceCounterDataTypeBuilder
	// WithTotalCount adds TotalCount (property field)
	WithTotalCount(uint32) ServiceCounterDataTypeBuilder
	// WithErrorCount adds ErrorCount (property field)
	WithErrorCount(uint32) ServiceCounterDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ServiceCounterDataType or returns an error if something is wrong
	Build() (ServiceCounterDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ServiceCounterDataType
}

// NewServiceCounterDataTypeBuilder() creates a ServiceCounterDataTypeBuilder
func NewServiceCounterDataTypeBuilder() ServiceCounterDataTypeBuilder {
	return &_ServiceCounterDataTypeBuilder{_ServiceCounterDataType: new(_ServiceCounterDataType)}
}

type _ServiceCounterDataTypeBuilder struct {
	*_ServiceCounterDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ServiceCounterDataTypeBuilder) = (*_ServiceCounterDataTypeBuilder)(nil)

func (b *_ServiceCounterDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ServiceCounterDataType
}

func (b *_ServiceCounterDataTypeBuilder) WithMandatoryFields(totalCount uint32, errorCount uint32) ServiceCounterDataTypeBuilder {
	return b.WithTotalCount(totalCount).WithErrorCount(errorCount)
}

func (b *_ServiceCounterDataTypeBuilder) WithTotalCount(totalCount uint32) ServiceCounterDataTypeBuilder {
	b.TotalCount = totalCount
	return b
}

func (b *_ServiceCounterDataTypeBuilder) WithErrorCount(errorCount uint32) ServiceCounterDataTypeBuilder {
	b.ErrorCount = errorCount
	return b
}

func (b *_ServiceCounterDataTypeBuilder) Build() (ServiceCounterDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ServiceCounterDataType.deepCopy(), nil
}

func (b *_ServiceCounterDataTypeBuilder) MustBuild() ServiceCounterDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ServiceCounterDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ServiceCounterDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ServiceCounterDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateServiceCounterDataTypeBuilder().(*_ServiceCounterDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateServiceCounterDataTypeBuilder creates a ServiceCounterDataTypeBuilder
func (b *_ServiceCounterDataType) CreateServiceCounterDataTypeBuilder() ServiceCounterDataTypeBuilder {
	if b == nil {
		return NewServiceCounterDataTypeBuilder()
	}
	return &_ServiceCounterDataTypeBuilder{_ServiceCounterDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServiceCounterDataType) GetExtensionId() int32 {
	return int32(873)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServiceCounterDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServiceCounterDataType) GetTotalCount() uint32 {
	return m.TotalCount
}

func (m *_ServiceCounterDataType) GetErrorCount() uint32 {
	return m.ErrorCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServiceCounterDataType(structType any) ServiceCounterDataType {
	if casted, ok := structType.(ServiceCounterDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServiceCounterDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServiceCounterDataType) GetTypeName() string {
	return "ServiceCounterDataType"
}

func (m *_ServiceCounterDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (totalCount)
	lengthInBits += 32

	// Simple field (errorCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ServiceCounterDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServiceCounterDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__serviceCounterDataType ServiceCounterDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServiceCounterDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServiceCounterDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	totalCount, err := ReadSimpleField(ctx, "totalCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'totalCount' field"))
	}
	m.TotalCount = totalCount

	errorCount, err := ReadSimpleField(ctx, "errorCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCount' field"))
	}
	m.ErrorCount = errorCount

	if closeErr := readBuffer.CloseContext("ServiceCounterDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServiceCounterDataType")
	}

	return m, nil
}

func (m *_ServiceCounterDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServiceCounterDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServiceCounterDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServiceCounterDataType")
		}

		if err := WriteSimpleField[uint32](ctx, "totalCount", m.GetTotalCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'totalCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "errorCount", m.GetErrorCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorCount' field")
		}

		if popErr := writeBuffer.PopContext("ServiceCounterDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServiceCounterDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServiceCounterDataType) IsServiceCounterDataType() {}

func (m *_ServiceCounterDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ServiceCounterDataType) deepCopy() *_ServiceCounterDataType {
	if m == nil {
		return nil
	}
	_ServiceCounterDataTypeCopy := &_ServiceCounterDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.TotalCount,
		m.ErrorCount,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ServiceCounterDataTypeCopy
}

func (m *_ServiceCounterDataType) String() string {
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
