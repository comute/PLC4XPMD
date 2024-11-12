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

// DataSetMetaDataType is the corresponding interface of DataSetMetaDataType
type DataSetMetaDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNamespaces returns Namespaces (property field)
	GetNamespaces() []PascalString
	// GetStructureDataTypes returns StructureDataTypes (property field)
	GetStructureDataTypes() []StructureDescription
	// GetEnumDataTypes returns EnumDataTypes (property field)
	GetEnumDataTypes() []EnumDescription
	// GetSimpleDataTypes returns SimpleDataTypes (property field)
	GetSimpleDataTypes() []SimpleTypeDescription
	// GetName returns Name (property field)
	GetName() PascalString
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetFields returns Fields (property field)
	GetFields() []FieldMetaData
	// GetDataSetClassId returns DataSetClassId (property field)
	GetDataSetClassId() GuidValue
	// GetConfigurationVersion returns ConfigurationVersion (property field)
	GetConfigurationVersion() ConfigurationVersionDataType
	// IsDataSetMetaDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetMetaDataType()
	// CreateBuilder creates a DataSetMetaDataTypeBuilder
	CreateDataSetMetaDataTypeBuilder() DataSetMetaDataTypeBuilder
}

// _DataSetMetaDataType is the data-structure of this message
type _DataSetMetaDataType struct {
	ExtensionObjectDefinitionContract
	Namespaces           []PascalString
	StructureDataTypes   []StructureDescription
	EnumDataTypes        []EnumDescription
	SimpleDataTypes      []SimpleTypeDescription
	Name                 PascalString
	Description          LocalizedText
	Fields               []FieldMetaData
	DataSetClassId       GuidValue
	ConfigurationVersion ConfigurationVersionDataType
}

var _ DataSetMetaDataType = (*_DataSetMetaDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetMetaDataType)(nil)

// NewDataSetMetaDataType factory function for _DataSetMetaDataType
func NewDataSetMetaDataType(namespaces []PascalString, structureDataTypes []StructureDescription, enumDataTypes []EnumDescription, simpleDataTypes []SimpleTypeDescription, name PascalString, description LocalizedText, fields []FieldMetaData, dataSetClassId GuidValue, configurationVersion ConfigurationVersionDataType) *_DataSetMetaDataType {
	if name == nil {
		panic("name of type PascalString for DataSetMetaDataType must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for DataSetMetaDataType must not be nil")
	}
	if dataSetClassId == nil {
		panic("dataSetClassId of type GuidValue for DataSetMetaDataType must not be nil")
	}
	if configurationVersion == nil {
		panic("configurationVersion of type ConfigurationVersionDataType for DataSetMetaDataType must not be nil")
	}
	_result := &_DataSetMetaDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Namespaces:                        namespaces,
		StructureDataTypes:                structureDataTypes,
		EnumDataTypes:                     enumDataTypes,
		SimpleDataTypes:                   simpleDataTypes,
		Name:                              name,
		Description:                       description,
		Fields:                            fields,
		DataSetClassId:                    dataSetClassId,
		ConfigurationVersion:              configurationVersion,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSetMetaDataTypeBuilder is a builder for DataSetMetaDataType
type DataSetMetaDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaces []PascalString, structureDataTypes []StructureDescription, enumDataTypes []EnumDescription, simpleDataTypes []SimpleTypeDescription, name PascalString, description LocalizedText, fields []FieldMetaData, dataSetClassId GuidValue, configurationVersion ConfigurationVersionDataType) DataSetMetaDataTypeBuilder
	// WithNamespaces adds Namespaces (property field)
	WithNamespaces(...PascalString) DataSetMetaDataTypeBuilder
	// WithStructureDataTypes adds StructureDataTypes (property field)
	WithStructureDataTypes(...StructureDescription) DataSetMetaDataTypeBuilder
	// WithEnumDataTypes adds EnumDataTypes (property field)
	WithEnumDataTypes(...EnumDescription) DataSetMetaDataTypeBuilder
	// WithSimpleDataTypes adds SimpleDataTypes (property field)
	WithSimpleDataTypes(...SimpleTypeDescription) DataSetMetaDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) DataSetMetaDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) DataSetMetaDataTypeBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) DataSetMetaDataTypeBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) DataSetMetaDataTypeBuilder
	// WithFields adds Fields (property field)
	WithFields(...FieldMetaData) DataSetMetaDataTypeBuilder
	// WithDataSetClassId adds DataSetClassId (property field)
	WithDataSetClassId(GuidValue) DataSetMetaDataTypeBuilder
	// WithDataSetClassIdBuilder adds DataSetClassId (property field) which is build by the builder
	WithDataSetClassIdBuilder(func(GuidValueBuilder) GuidValueBuilder) DataSetMetaDataTypeBuilder
	// WithConfigurationVersion adds ConfigurationVersion (property field)
	WithConfigurationVersion(ConfigurationVersionDataType) DataSetMetaDataTypeBuilder
	// WithConfigurationVersionBuilder adds ConfigurationVersion (property field) which is build by the builder
	WithConfigurationVersionBuilder(func(ConfigurationVersionDataTypeBuilder) ConfigurationVersionDataTypeBuilder) DataSetMetaDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DataSetMetaDataType or returns an error if something is wrong
	Build() (DataSetMetaDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSetMetaDataType
}

// NewDataSetMetaDataTypeBuilder() creates a DataSetMetaDataTypeBuilder
func NewDataSetMetaDataTypeBuilder() DataSetMetaDataTypeBuilder {
	return &_DataSetMetaDataTypeBuilder{_DataSetMetaDataType: new(_DataSetMetaDataType)}
}

type _DataSetMetaDataTypeBuilder struct {
	*_DataSetMetaDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DataSetMetaDataTypeBuilder) = (*_DataSetMetaDataTypeBuilder)(nil)

func (b *_DataSetMetaDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._DataSetMetaDataType
}

func (b *_DataSetMetaDataTypeBuilder) WithMandatoryFields(namespaces []PascalString, structureDataTypes []StructureDescription, enumDataTypes []EnumDescription, simpleDataTypes []SimpleTypeDescription, name PascalString, description LocalizedText, fields []FieldMetaData, dataSetClassId GuidValue, configurationVersion ConfigurationVersionDataType) DataSetMetaDataTypeBuilder {
	return b.WithNamespaces(namespaces...).WithStructureDataTypes(structureDataTypes...).WithEnumDataTypes(enumDataTypes...).WithSimpleDataTypes(simpleDataTypes...).WithName(name).WithDescription(description).WithFields(fields...).WithDataSetClassId(dataSetClassId).WithConfigurationVersion(configurationVersion)
}

func (b *_DataSetMetaDataTypeBuilder) WithNamespaces(namespaces ...PascalString) DataSetMetaDataTypeBuilder {
	b.Namespaces = namespaces
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithStructureDataTypes(structureDataTypes ...StructureDescription) DataSetMetaDataTypeBuilder {
	b.StructureDataTypes = structureDataTypes
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithEnumDataTypes(enumDataTypes ...EnumDescription) DataSetMetaDataTypeBuilder {
	b.EnumDataTypes = enumDataTypes
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithSimpleDataTypes(simpleDataTypes ...SimpleTypeDescription) DataSetMetaDataTypeBuilder {
	b.SimpleDataTypes = simpleDataTypes
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithName(name PascalString) DataSetMetaDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) DataSetMetaDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithDescription(description LocalizedText) DataSetMetaDataTypeBuilder {
	b.Description = description
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) DataSetMetaDataTypeBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithFields(fields ...FieldMetaData) DataSetMetaDataTypeBuilder {
	b.Fields = fields
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithDataSetClassId(dataSetClassId GuidValue) DataSetMetaDataTypeBuilder {
	b.DataSetClassId = dataSetClassId
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithDataSetClassIdBuilder(builderSupplier func(GuidValueBuilder) GuidValueBuilder) DataSetMetaDataTypeBuilder {
	builder := builderSupplier(b.DataSetClassId.CreateGuidValueBuilder())
	var err error
	b.DataSetClassId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "GuidValueBuilder failed"))
	}
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithConfigurationVersion(configurationVersion ConfigurationVersionDataType) DataSetMetaDataTypeBuilder {
	b.ConfigurationVersion = configurationVersion
	return b
}

func (b *_DataSetMetaDataTypeBuilder) WithConfigurationVersionBuilder(builderSupplier func(ConfigurationVersionDataTypeBuilder) ConfigurationVersionDataTypeBuilder) DataSetMetaDataTypeBuilder {
	builder := builderSupplier(b.ConfigurationVersion.CreateConfigurationVersionDataTypeBuilder())
	var err error
	b.ConfigurationVersion, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ConfigurationVersionDataTypeBuilder failed"))
	}
	return b
}

func (b *_DataSetMetaDataTypeBuilder) Build() (DataSetMetaDataType, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.DataSetClassId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataSetClassId' not set"))
	}
	if b.ConfigurationVersion == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'configurationVersion' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataSetMetaDataType.deepCopy(), nil
}

func (b *_DataSetMetaDataTypeBuilder) MustBuild() DataSetMetaDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataSetMetaDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DataSetMetaDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DataSetMetaDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateDataSetMetaDataTypeBuilder().(*_DataSetMetaDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataSetMetaDataTypeBuilder creates a DataSetMetaDataTypeBuilder
func (b *_DataSetMetaDataType) CreateDataSetMetaDataTypeBuilder() DataSetMetaDataTypeBuilder {
	if b == nil {
		return NewDataSetMetaDataTypeBuilder()
	}
	return &_DataSetMetaDataTypeBuilder{_DataSetMetaDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetMetaDataType) GetExtensionId() int32 {
	return int32(14525)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetMetaDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSetMetaDataType) GetNamespaces() []PascalString {
	return m.Namespaces
}

func (m *_DataSetMetaDataType) GetStructureDataTypes() []StructureDescription {
	return m.StructureDataTypes
}

func (m *_DataSetMetaDataType) GetEnumDataTypes() []EnumDescription {
	return m.EnumDataTypes
}

func (m *_DataSetMetaDataType) GetSimpleDataTypes() []SimpleTypeDescription {
	return m.SimpleDataTypes
}

func (m *_DataSetMetaDataType) GetName() PascalString {
	return m.Name
}

func (m *_DataSetMetaDataType) GetDescription() LocalizedText {
	return m.Description
}

func (m *_DataSetMetaDataType) GetFields() []FieldMetaData {
	return m.Fields
}

func (m *_DataSetMetaDataType) GetDataSetClassId() GuidValue {
	return m.DataSetClassId
}

func (m *_DataSetMetaDataType) GetConfigurationVersion() ConfigurationVersionDataType {
	return m.ConfigurationVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataSetMetaDataType(structType any) DataSetMetaDataType {
	if casted, ok := structType.(DataSetMetaDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetMetaDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetMetaDataType) GetTypeName() string {
	return "DataSetMetaDataType"
}

func (m *_DataSetMetaDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfNamespaces)
	lengthInBits += 32

	// Array field
	if len(m.Namespaces) > 0 {
		for _curItem, element := range m.Namespaces {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Namespaces), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfStructureDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.StructureDataTypes) > 0 {
		for _curItem, element := range m.StructureDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StructureDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfEnumDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.EnumDataTypes) > 0 {
		for _curItem, element := range m.EnumDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EnumDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfSimpleDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.SimpleDataTypes) > 0 {
		for _curItem, element := range m.SimpleDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SimpleDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Implicit Field (noOfFields)
	lengthInBits += 32

	// Array field
	if len(m.Fields) > 0 {
		for _curItem, element := range m.Fields {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Fields), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (dataSetClassId)
	lengthInBits += m.DataSetClassId.GetLengthInBits(ctx)

	// Simple field (configurationVersion)
	lengthInBits += m.ConfigurationVersion.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSetMetaDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetMetaDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__dataSetMetaDataType DataSetMetaDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetMetaDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetMetaDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfNamespaces, err := ReadImplicitField[int32](ctx, "noOfNamespaces", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNamespaces' field"))
	}
	_ = noOfNamespaces

	namespaces, err := ReadCountArrayField[PascalString](ctx, "namespaces", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfNamespaces))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaces' field"))
	}
	m.Namespaces = namespaces

	noOfStructureDataTypes, err := ReadImplicitField[int32](ctx, "noOfStructureDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfStructureDataTypes' field"))
	}
	_ = noOfStructureDataTypes

	structureDataTypes, err := ReadCountArrayField[StructureDescription](ctx, "structureDataTypes", ReadComplex[StructureDescription](ExtensionObjectDefinitionParseWithBufferProducer[StructureDescription]((int32)(int32(15489))), readBuffer), uint64(noOfStructureDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureDataTypes' field"))
	}
	m.StructureDataTypes = structureDataTypes

	noOfEnumDataTypes, err := ReadImplicitField[int32](ctx, "noOfEnumDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEnumDataTypes' field"))
	}
	_ = noOfEnumDataTypes

	enumDataTypes, err := ReadCountArrayField[EnumDescription](ctx, "enumDataTypes", ReadComplex[EnumDescription](ExtensionObjectDefinitionParseWithBufferProducer[EnumDescription]((int32)(int32(15490))), readBuffer), uint64(noOfEnumDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumDataTypes' field"))
	}
	m.EnumDataTypes = enumDataTypes

	noOfSimpleDataTypes, err := ReadImplicitField[int32](ctx, "noOfSimpleDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSimpleDataTypes' field"))
	}
	_ = noOfSimpleDataTypes

	simpleDataTypes, err := ReadCountArrayField[SimpleTypeDescription](ctx, "simpleDataTypes", ReadComplex[SimpleTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[SimpleTypeDescription]((int32)(int32(15007))), readBuffer), uint64(noOfSimpleDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleDataTypes' field"))
	}
	m.SimpleDataTypes = simpleDataTypes

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	noOfFields, err := ReadImplicitField[int32](ctx, "noOfFields", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfFields' field"))
	}
	_ = noOfFields

	fields, err := ReadCountArrayField[FieldMetaData](ctx, "fields", ReadComplex[FieldMetaData](ExtensionObjectDefinitionParseWithBufferProducer[FieldMetaData]((int32)(int32(14526))), readBuffer), uint64(noOfFields))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fields' field"))
	}
	m.Fields = fields

	dataSetClassId, err := ReadSimpleField[GuidValue](ctx, "dataSetClassId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetClassId' field"))
	}
	m.DataSetClassId = dataSetClassId

	configurationVersion, err := ReadSimpleField[ConfigurationVersionDataType](ctx, "configurationVersion", ReadComplex[ConfigurationVersionDataType](ExtensionObjectDefinitionParseWithBufferProducer[ConfigurationVersionDataType]((int32)(int32(14595))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationVersion' field"))
	}
	m.ConfigurationVersion = configurationVersion

	if closeErr := readBuffer.CloseContext("DataSetMetaDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetMetaDataType")
	}

	return m, nil
}

func (m *_DataSetMetaDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetMetaDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetMetaDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetMetaDataType")
		}
		noOfNamespaces := int32(utils.InlineIf(bool((m.GetNamespaces()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNamespaces()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNamespaces", noOfNamespaces, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNamespaces' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "namespaces", m.GetNamespaces(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaces' field")
		}
		noOfStructureDataTypes := int32(utils.InlineIf(bool((m.GetStructureDataTypes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetStructureDataTypes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfStructureDataTypes", noOfStructureDataTypes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfStructureDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "structureDataTypes", m.GetStructureDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'structureDataTypes' field")
		}
		noOfEnumDataTypes := int32(utils.InlineIf(bool((m.GetEnumDataTypes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetEnumDataTypes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfEnumDataTypes", noOfEnumDataTypes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEnumDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "enumDataTypes", m.GetEnumDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'enumDataTypes' field")
		}
		noOfSimpleDataTypes := int32(utils.InlineIf(bool((m.GetSimpleDataTypes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSimpleDataTypes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSimpleDataTypes", noOfSimpleDataTypes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSimpleDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "simpleDataTypes", m.GetSimpleDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'simpleDataTypes' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}
		noOfFields := int32(utils.InlineIf(bool((m.GetFields()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetFields()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfFields", noOfFields, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfFields' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "fields", m.GetFields(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'fields' field")
		}

		if err := WriteSimpleField[GuidValue](ctx, "dataSetClassId", m.GetDataSetClassId(), WriteComplex[GuidValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetClassId' field")
		}

		if err := WriteSimpleField[ConfigurationVersionDataType](ctx, "configurationVersion", m.GetConfigurationVersion(), WriteComplex[ConfigurationVersionDataType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationVersion' field")
		}

		if popErr := writeBuffer.PopContext("DataSetMetaDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetMetaDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetMetaDataType) IsDataSetMetaDataType() {}

func (m *_DataSetMetaDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSetMetaDataType) deepCopy() *_DataSetMetaDataType {
	if m == nil {
		return nil
	}
	_DataSetMetaDataTypeCopy := &_DataSetMetaDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[PascalString, PascalString](m.Namespaces),
		utils.DeepCopySlice[StructureDescription, StructureDescription](m.StructureDataTypes),
		utils.DeepCopySlice[EnumDescription, EnumDescription](m.EnumDataTypes),
		utils.DeepCopySlice[SimpleTypeDescription, SimpleTypeDescription](m.SimpleDataTypes),
		utils.DeepCopy[PascalString](m.Name),
		utils.DeepCopy[LocalizedText](m.Description),
		utils.DeepCopySlice[FieldMetaData, FieldMetaData](m.Fields),
		utils.DeepCopy[GuidValue](m.DataSetClassId),
		utils.DeepCopy[ConfigurationVersionDataType](m.ConfigurationVersion),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataSetMetaDataTypeCopy
}

func (m *_DataSetMetaDataType) String() string {
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
