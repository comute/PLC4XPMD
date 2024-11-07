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

// EnumDescription is the corresponding interface of EnumDescription
type EnumDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDataTypeId returns DataTypeId (property field)
	GetDataTypeId() NodeId
	// GetName returns Name (property field)
	GetName() QualifiedName
	// GetEnumDefinition returns EnumDefinition (property field)
	GetEnumDefinition() EnumDefinition
	// GetBuiltInType returns BuiltInType (property field)
	GetBuiltInType() uint8
	// IsEnumDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEnumDescription()
	// CreateBuilder creates a EnumDescriptionBuilder
	CreateEnumDescriptionBuilder() EnumDescriptionBuilder
}

// _EnumDescription is the data-structure of this message
type _EnumDescription struct {
	ExtensionObjectDefinitionContract
	DataTypeId     NodeId
	Name           QualifiedName
	EnumDefinition EnumDefinition
	BuiltInType    uint8
}

var _ EnumDescription = (*_EnumDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EnumDescription)(nil)

// NewEnumDescription factory function for _EnumDescription
func NewEnumDescription(dataTypeId NodeId, name QualifiedName, enumDefinition EnumDefinition, builtInType uint8) *_EnumDescription {
	if dataTypeId == nil {
		panic("dataTypeId of type NodeId for EnumDescription must not be nil")
	}
	if name == nil {
		panic("name of type QualifiedName for EnumDescription must not be nil")
	}
	if enumDefinition == nil {
		panic("enumDefinition of type EnumDefinition for EnumDescription must not be nil")
	}
	_result := &_EnumDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataTypeId:                        dataTypeId,
		Name:                              name,
		EnumDefinition:                    enumDefinition,
		BuiltInType:                       builtInType,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EnumDescriptionBuilder is a builder for EnumDescription
type EnumDescriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataTypeId NodeId, name QualifiedName, enumDefinition EnumDefinition, builtInType uint8) EnumDescriptionBuilder
	// WithDataTypeId adds DataTypeId (property field)
	WithDataTypeId(NodeId) EnumDescriptionBuilder
	// WithDataTypeIdBuilder adds DataTypeId (property field) which is build by the builder
	WithDataTypeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) EnumDescriptionBuilder
	// WithName adds Name (property field)
	WithName(QualifiedName) EnumDescriptionBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) EnumDescriptionBuilder
	// WithEnumDefinition adds EnumDefinition (property field)
	WithEnumDefinition(EnumDefinition) EnumDescriptionBuilder
	// WithEnumDefinitionBuilder adds EnumDefinition (property field) which is build by the builder
	WithEnumDefinitionBuilder(func(EnumDefinitionBuilder) EnumDefinitionBuilder) EnumDescriptionBuilder
	// WithBuiltInType adds BuiltInType (property field)
	WithBuiltInType(uint8) EnumDescriptionBuilder
	// Build builds the EnumDescription or returns an error if something is wrong
	Build() (EnumDescription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EnumDescription
}

// NewEnumDescriptionBuilder() creates a EnumDescriptionBuilder
func NewEnumDescriptionBuilder() EnumDescriptionBuilder {
	return &_EnumDescriptionBuilder{_EnumDescription: new(_EnumDescription)}
}

type _EnumDescriptionBuilder struct {
	*_EnumDescription

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EnumDescriptionBuilder) = (*_EnumDescriptionBuilder)(nil)

func (b *_EnumDescriptionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_EnumDescriptionBuilder) WithMandatoryFields(dataTypeId NodeId, name QualifiedName, enumDefinition EnumDefinition, builtInType uint8) EnumDescriptionBuilder {
	return b.WithDataTypeId(dataTypeId).WithName(name).WithEnumDefinition(enumDefinition).WithBuiltInType(builtInType)
}

func (b *_EnumDescriptionBuilder) WithDataTypeId(dataTypeId NodeId) EnumDescriptionBuilder {
	b.DataTypeId = dataTypeId
	return b
}

func (b *_EnumDescriptionBuilder) WithDataTypeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) EnumDescriptionBuilder {
	builder := builderSupplier(b.DataTypeId.CreateNodeIdBuilder())
	var err error
	b.DataTypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_EnumDescriptionBuilder) WithName(name QualifiedName) EnumDescriptionBuilder {
	b.Name = name
	return b
}

func (b *_EnumDescriptionBuilder) WithNameBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) EnumDescriptionBuilder {
	builder := builderSupplier(b.Name.CreateQualifiedNameBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return b
}

func (b *_EnumDescriptionBuilder) WithEnumDefinition(enumDefinition EnumDefinition) EnumDescriptionBuilder {
	b.EnumDefinition = enumDefinition
	return b
}

func (b *_EnumDescriptionBuilder) WithEnumDefinitionBuilder(builderSupplier func(EnumDefinitionBuilder) EnumDefinitionBuilder) EnumDescriptionBuilder {
	builder := builderSupplier(b.EnumDefinition.CreateEnumDefinitionBuilder())
	var err error
	b.EnumDefinition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "EnumDefinitionBuilder failed"))
	}
	return b
}

func (b *_EnumDescriptionBuilder) WithBuiltInType(builtInType uint8) EnumDescriptionBuilder {
	b.BuiltInType = builtInType
	return b
}

func (b *_EnumDescriptionBuilder) Build() (EnumDescription, error) {
	if b.DataTypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataTypeId' not set"))
	}
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.EnumDefinition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'enumDefinition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EnumDescription.deepCopy(), nil
}

func (b *_EnumDescriptionBuilder) MustBuild() EnumDescription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_EnumDescriptionBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_EnumDescriptionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EnumDescriptionBuilder) DeepCopy() any {
	_copy := b.CreateEnumDescriptionBuilder().(*_EnumDescriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEnumDescriptionBuilder creates a EnumDescriptionBuilder
func (b *_EnumDescription) CreateEnumDescriptionBuilder() EnumDescriptionBuilder {
	if b == nil {
		return NewEnumDescriptionBuilder()
	}
	return &_EnumDescriptionBuilder{_EnumDescription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EnumDescription) GetExtensionId() int32 {
	return int32(15490)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EnumDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EnumDescription) GetDataTypeId() NodeId {
	return m.DataTypeId
}

func (m *_EnumDescription) GetName() QualifiedName {
	return m.Name
}

func (m *_EnumDescription) GetEnumDefinition() EnumDefinition {
	return m.EnumDefinition
}

func (m *_EnumDescription) GetBuiltInType() uint8 {
	return m.BuiltInType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEnumDescription(structType any) EnumDescription {
	if casted, ok := structType.(EnumDescription); ok {
		return casted
	}
	if casted, ok := structType.(*EnumDescription); ok {
		return *casted
	}
	return nil
}

func (m *_EnumDescription) GetTypeName() string {
	return "EnumDescription"
}

func (m *_EnumDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (dataTypeId)
	lengthInBits += m.DataTypeId.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (enumDefinition)
	lengthInBits += m.EnumDefinition.GetLengthInBits(ctx)

	// Simple field (builtInType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_EnumDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EnumDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__enumDescription EnumDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EnumDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EnumDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataTypeId, err := ReadSimpleField[NodeId](ctx, "dataTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeId' field"))
	}
	m.DataTypeId = dataTypeId

	name, err := ReadSimpleField[QualifiedName](ctx, "name", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	enumDefinition, err := ReadSimpleField[EnumDefinition](ctx, "enumDefinition", ReadComplex[EnumDefinition](ExtensionObjectDefinitionParseWithBufferProducer[EnumDefinition]((int32)(int32(102))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumDefinition' field"))
	}
	m.EnumDefinition = enumDefinition

	builtInType, err := ReadSimpleField(ctx, "builtInType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'builtInType' field"))
	}
	m.BuiltInType = builtInType

	if closeErr := readBuffer.CloseContext("EnumDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EnumDescription")
	}

	return m, nil
}

func (m *_EnumDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EnumDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EnumDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EnumDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataTypeId", m.GetDataTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataTypeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "name", m.GetName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[EnumDefinition](ctx, "enumDefinition", m.GetEnumDefinition(), WriteComplex[EnumDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enumDefinition' field")
		}

		if err := WriteSimpleField[uint8](ctx, "builtInType", m.GetBuiltInType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'builtInType' field")
		}

		if popErr := writeBuffer.PopContext("EnumDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EnumDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EnumDescription) IsEnumDescription() {}

func (m *_EnumDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EnumDescription) deepCopy() *_EnumDescription {
	if m == nil {
		return nil
	}
	_EnumDescriptionCopy := &_EnumDescription{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.DataTypeId.DeepCopy().(NodeId),
		m.Name.DeepCopy().(QualifiedName),
		m.EnumDefinition.DeepCopy().(EnumDefinition),
		m.BuiltInType,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EnumDescriptionCopy
}

func (m *_EnumDescription) String() string {
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