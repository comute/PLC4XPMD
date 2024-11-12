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

// PubSubConfiguration2DataType is the corresponding interface of PubSubConfiguration2DataType
type PubSubConfiguration2DataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPublishedDataSets returns PublishedDataSets (property field)
	GetPublishedDataSets() []PublishedDataSetDataType
	// GetConnections returns Connections (property field)
	GetConnections() []PubSubConnectionDataType
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetSubscribedDataSets returns SubscribedDataSets (property field)
	GetSubscribedDataSets() []StandaloneSubscribedDataSetDataType
	// GetDataSetClasses returns DataSetClasses (property field)
	GetDataSetClasses() []DataSetMetaDataType
	// GetDefaultSecurityKeyServices returns DefaultSecurityKeyServices (property field)
	GetDefaultSecurityKeyServices() []EndpointDescription
	// GetSecurityGroups returns SecurityGroups (property field)
	GetSecurityGroups() []SecurityGroupDataType
	// GetPubSubKeyPushTargets returns PubSubKeyPushTargets (property field)
	GetPubSubKeyPushTargets() []PubSubKeyPushTargetDataType
	// GetConfigurationVersion returns ConfigurationVersion (property field)
	GetConfigurationVersion() uint32
	// GetConfigurationProperties returns ConfigurationProperties (property field)
	GetConfigurationProperties() []KeyValuePair
	// IsPubSubConfiguration2DataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubConfiguration2DataType()
	// CreateBuilder creates a PubSubConfiguration2DataTypeBuilder
	CreatePubSubConfiguration2DataTypeBuilder() PubSubConfiguration2DataTypeBuilder
}

// _PubSubConfiguration2DataType is the data-structure of this message
type _PubSubConfiguration2DataType struct {
	ExtensionObjectDefinitionContract
	PublishedDataSets          []PublishedDataSetDataType
	Connections                []PubSubConnectionDataType
	Enabled                    bool
	SubscribedDataSets         []StandaloneSubscribedDataSetDataType
	DataSetClasses             []DataSetMetaDataType
	DefaultSecurityKeyServices []EndpointDescription
	SecurityGroups             []SecurityGroupDataType
	PubSubKeyPushTargets       []PubSubKeyPushTargetDataType
	ConfigurationVersion       uint32
	ConfigurationProperties    []KeyValuePair
	// Reserved Fields
	reservedField0 *uint8
}

var _ PubSubConfiguration2DataType = (*_PubSubConfiguration2DataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubConfiguration2DataType)(nil)

// NewPubSubConfiguration2DataType factory function for _PubSubConfiguration2DataType
func NewPubSubConfiguration2DataType(publishedDataSets []PublishedDataSetDataType, connections []PubSubConnectionDataType, enabled bool, subscribedDataSets []StandaloneSubscribedDataSetDataType, dataSetClasses []DataSetMetaDataType, defaultSecurityKeyServices []EndpointDescription, securityGroups []SecurityGroupDataType, pubSubKeyPushTargets []PubSubKeyPushTargetDataType, configurationVersion uint32, configurationProperties []KeyValuePair) *_PubSubConfiguration2DataType {
	_result := &_PubSubConfiguration2DataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PublishedDataSets:                 publishedDataSets,
		Connections:                       connections,
		Enabled:                           enabled,
		SubscribedDataSets:                subscribedDataSets,
		DataSetClasses:                    dataSetClasses,
		DefaultSecurityKeyServices:        defaultSecurityKeyServices,
		SecurityGroups:                    securityGroups,
		PubSubKeyPushTargets:              pubSubKeyPushTargets,
		ConfigurationVersion:              configurationVersion,
		ConfigurationProperties:           configurationProperties,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PubSubConfiguration2DataTypeBuilder is a builder for PubSubConfiguration2DataType
type PubSubConfiguration2DataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(publishedDataSets []PublishedDataSetDataType, connections []PubSubConnectionDataType, enabled bool, subscribedDataSets []StandaloneSubscribedDataSetDataType, dataSetClasses []DataSetMetaDataType, defaultSecurityKeyServices []EndpointDescription, securityGroups []SecurityGroupDataType, pubSubKeyPushTargets []PubSubKeyPushTargetDataType, configurationVersion uint32, configurationProperties []KeyValuePair) PubSubConfiguration2DataTypeBuilder
	// WithPublishedDataSets adds PublishedDataSets (property field)
	WithPublishedDataSets(...PublishedDataSetDataType) PubSubConfiguration2DataTypeBuilder
	// WithConnections adds Connections (property field)
	WithConnections(...PubSubConnectionDataType) PubSubConfiguration2DataTypeBuilder
	// WithEnabled adds Enabled (property field)
	WithEnabled(bool) PubSubConfiguration2DataTypeBuilder
	// WithSubscribedDataSets adds SubscribedDataSets (property field)
	WithSubscribedDataSets(...StandaloneSubscribedDataSetDataType) PubSubConfiguration2DataTypeBuilder
	// WithDataSetClasses adds DataSetClasses (property field)
	WithDataSetClasses(...DataSetMetaDataType) PubSubConfiguration2DataTypeBuilder
	// WithDefaultSecurityKeyServices adds DefaultSecurityKeyServices (property field)
	WithDefaultSecurityKeyServices(...EndpointDescription) PubSubConfiguration2DataTypeBuilder
	// WithSecurityGroups adds SecurityGroups (property field)
	WithSecurityGroups(...SecurityGroupDataType) PubSubConfiguration2DataTypeBuilder
	// WithPubSubKeyPushTargets adds PubSubKeyPushTargets (property field)
	WithPubSubKeyPushTargets(...PubSubKeyPushTargetDataType) PubSubConfiguration2DataTypeBuilder
	// WithConfigurationVersion adds ConfigurationVersion (property field)
	WithConfigurationVersion(uint32) PubSubConfiguration2DataTypeBuilder
	// WithConfigurationProperties adds ConfigurationProperties (property field)
	WithConfigurationProperties(...KeyValuePair) PubSubConfiguration2DataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the PubSubConfiguration2DataType or returns an error if something is wrong
	Build() (PubSubConfiguration2DataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PubSubConfiguration2DataType
}

// NewPubSubConfiguration2DataTypeBuilder() creates a PubSubConfiguration2DataTypeBuilder
func NewPubSubConfiguration2DataTypeBuilder() PubSubConfiguration2DataTypeBuilder {
	return &_PubSubConfiguration2DataTypeBuilder{_PubSubConfiguration2DataType: new(_PubSubConfiguration2DataType)}
}

type _PubSubConfiguration2DataTypeBuilder struct {
	*_PubSubConfiguration2DataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PubSubConfiguration2DataTypeBuilder) = (*_PubSubConfiguration2DataTypeBuilder)(nil)

func (b *_PubSubConfiguration2DataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._PubSubConfiguration2DataType
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithMandatoryFields(publishedDataSets []PublishedDataSetDataType, connections []PubSubConnectionDataType, enabled bool, subscribedDataSets []StandaloneSubscribedDataSetDataType, dataSetClasses []DataSetMetaDataType, defaultSecurityKeyServices []EndpointDescription, securityGroups []SecurityGroupDataType, pubSubKeyPushTargets []PubSubKeyPushTargetDataType, configurationVersion uint32, configurationProperties []KeyValuePair) PubSubConfiguration2DataTypeBuilder {
	return b.WithPublishedDataSets(publishedDataSets...).WithConnections(connections...).WithEnabled(enabled).WithSubscribedDataSets(subscribedDataSets...).WithDataSetClasses(dataSetClasses...).WithDefaultSecurityKeyServices(defaultSecurityKeyServices...).WithSecurityGroups(securityGroups...).WithPubSubKeyPushTargets(pubSubKeyPushTargets...).WithConfigurationVersion(configurationVersion).WithConfigurationProperties(configurationProperties...)
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithPublishedDataSets(publishedDataSets ...PublishedDataSetDataType) PubSubConfiguration2DataTypeBuilder {
	b.PublishedDataSets = publishedDataSets
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithConnections(connections ...PubSubConnectionDataType) PubSubConfiguration2DataTypeBuilder {
	b.Connections = connections
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithEnabled(enabled bool) PubSubConfiguration2DataTypeBuilder {
	b.Enabled = enabled
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithSubscribedDataSets(subscribedDataSets ...StandaloneSubscribedDataSetDataType) PubSubConfiguration2DataTypeBuilder {
	b.SubscribedDataSets = subscribedDataSets
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithDataSetClasses(dataSetClasses ...DataSetMetaDataType) PubSubConfiguration2DataTypeBuilder {
	b.DataSetClasses = dataSetClasses
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithDefaultSecurityKeyServices(defaultSecurityKeyServices ...EndpointDescription) PubSubConfiguration2DataTypeBuilder {
	b.DefaultSecurityKeyServices = defaultSecurityKeyServices
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithSecurityGroups(securityGroups ...SecurityGroupDataType) PubSubConfiguration2DataTypeBuilder {
	b.SecurityGroups = securityGroups
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithPubSubKeyPushTargets(pubSubKeyPushTargets ...PubSubKeyPushTargetDataType) PubSubConfiguration2DataTypeBuilder {
	b.PubSubKeyPushTargets = pubSubKeyPushTargets
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithConfigurationVersion(configurationVersion uint32) PubSubConfiguration2DataTypeBuilder {
	b.ConfigurationVersion = configurationVersion
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) WithConfigurationProperties(configurationProperties ...KeyValuePair) PubSubConfiguration2DataTypeBuilder {
	b.ConfigurationProperties = configurationProperties
	return b
}

func (b *_PubSubConfiguration2DataTypeBuilder) Build() (PubSubConfiguration2DataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PubSubConfiguration2DataType.deepCopy(), nil
}

func (b *_PubSubConfiguration2DataTypeBuilder) MustBuild() PubSubConfiguration2DataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PubSubConfiguration2DataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_PubSubConfiguration2DataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PubSubConfiguration2DataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePubSubConfiguration2DataTypeBuilder().(*_PubSubConfiguration2DataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePubSubConfiguration2DataTypeBuilder creates a PubSubConfiguration2DataTypeBuilder
func (b *_PubSubConfiguration2DataType) CreatePubSubConfiguration2DataTypeBuilder() PubSubConfiguration2DataTypeBuilder {
	if b == nil {
		return NewPubSubConfiguration2DataTypeBuilder()
	}
	return &_PubSubConfiguration2DataTypeBuilder{_PubSubConfiguration2DataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfiguration2DataType) GetExtensionId() int32 {
	return int32(23604)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfiguration2DataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfiguration2DataType) GetPublishedDataSets() []PublishedDataSetDataType {
	return m.PublishedDataSets
}

func (m *_PubSubConfiguration2DataType) GetConnections() []PubSubConnectionDataType {
	return m.Connections
}

func (m *_PubSubConfiguration2DataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_PubSubConfiguration2DataType) GetSubscribedDataSets() []StandaloneSubscribedDataSetDataType {
	return m.SubscribedDataSets
}

func (m *_PubSubConfiguration2DataType) GetDataSetClasses() []DataSetMetaDataType {
	return m.DataSetClasses
}

func (m *_PubSubConfiguration2DataType) GetDefaultSecurityKeyServices() []EndpointDescription {
	return m.DefaultSecurityKeyServices
}

func (m *_PubSubConfiguration2DataType) GetSecurityGroups() []SecurityGroupDataType {
	return m.SecurityGroups
}

func (m *_PubSubConfiguration2DataType) GetPubSubKeyPushTargets() []PubSubKeyPushTargetDataType {
	return m.PubSubKeyPushTargets
}

func (m *_PubSubConfiguration2DataType) GetConfigurationVersion() uint32 {
	return m.ConfigurationVersion
}

func (m *_PubSubConfiguration2DataType) GetConfigurationProperties() []KeyValuePair {
	return m.ConfigurationProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPubSubConfiguration2DataType(structType any) PubSubConfiguration2DataType {
	if casted, ok := structType.(PubSubConfiguration2DataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfiguration2DataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfiguration2DataType) GetTypeName() string {
	return "PubSubConfiguration2DataType"
}

func (m *_PubSubConfiguration2DataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfPublishedDataSets)
	lengthInBits += 32

	// Array field
	if len(m.PublishedDataSets) > 0 {
		for _curItem, element := range m.PublishedDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PublishedDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfConnections)
	lengthInBits += 32

	// Array field
	if len(m.Connections) > 0 {
		for _curItem, element := range m.Connections {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Connections), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Implicit Field (noOfSubscribedDataSets)
	lengthInBits += 32

	// Array field
	if len(m.SubscribedDataSets) > 0 {
		for _curItem, element := range m.SubscribedDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SubscribedDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDataSetClasses)
	lengthInBits += 32

	// Array field
	if len(m.DataSetClasses) > 0 {
		for _curItem, element := range m.DataSetClasses {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataSetClasses), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDefaultSecurityKeyServices)
	lengthInBits += 32

	// Array field
	if len(m.DefaultSecurityKeyServices) > 0 {
		for _curItem, element := range m.DefaultSecurityKeyServices {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DefaultSecurityKeyServices), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfSecurityGroups)
	lengthInBits += 32

	// Array field
	if len(m.SecurityGroups) > 0 {
		for _curItem, element := range m.SecurityGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfPubSubKeyPushTargets)
	lengthInBits += 32

	// Array field
	if len(m.PubSubKeyPushTargets) > 0 {
		for _curItem, element := range m.PubSubKeyPushTargets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PubSubKeyPushTargets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (configurationVersion)
	lengthInBits += 32

	// Implicit Field (noOfConfigurationProperties)
	lengthInBits += 32

	// Array field
	if len(m.ConfigurationProperties) > 0 {
		for _curItem, element := range m.ConfigurationProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ConfigurationProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PubSubConfiguration2DataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubConfiguration2DataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__pubSubConfiguration2DataType PubSubConfiguration2DataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubConfiguration2DataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfiguration2DataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfPublishedDataSets, err := ReadImplicitField[int32](ctx, "noOfPublishedDataSets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfPublishedDataSets' field"))
	}
	_ = noOfPublishedDataSets

	publishedDataSets, err := ReadCountArrayField[PublishedDataSetDataType](ctx, "publishedDataSets", ReadComplex[PublishedDataSetDataType](ExtensionObjectDefinitionParseWithBufferProducer[PublishedDataSetDataType]((int32)(int32(15580))), readBuffer), uint64(noOfPublishedDataSets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishedDataSets' field"))
	}
	m.PublishedDataSets = publishedDataSets

	noOfConnections, err := ReadImplicitField[int32](ctx, "noOfConnections", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConnections' field"))
	}
	_ = noOfConnections

	connections, err := ReadCountArrayField[PubSubConnectionDataType](ctx, "connections", ReadComplex[PubSubConnectionDataType](ExtensionObjectDefinitionParseWithBufferProducer[PubSubConnectionDataType]((int32)(int32(15619))), readBuffer), uint64(noOfConnections))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connections' field"))
	}
	m.Connections = connections

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	noOfSubscribedDataSets, err := ReadImplicitField[int32](ctx, "noOfSubscribedDataSets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSubscribedDataSets' field"))
	}
	_ = noOfSubscribedDataSets

	subscribedDataSets, err := ReadCountArrayField[StandaloneSubscribedDataSetDataType](ctx, "subscribedDataSets", ReadComplex[StandaloneSubscribedDataSetDataType](ExtensionObjectDefinitionParseWithBufferProducer[StandaloneSubscribedDataSetDataType]((int32)(int32(23602))), readBuffer), uint64(noOfSubscribedDataSets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscribedDataSets' field"))
	}
	m.SubscribedDataSets = subscribedDataSets

	noOfDataSetClasses, err := ReadImplicitField[int32](ctx, "noOfDataSetClasses", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataSetClasses' field"))
	}
	_ = noOfDataSetClasses

	dataSetClasses, err := ReadCountArrayField[DataSetMetaDataType](ctx, "dataSetClasses", ReadComplex[DataSetMetaDataType](ExtensionObjectDefinitionParseWithBufferProducer[DataSetMetaDataType]((int32)(int32(14525))), readBuffer), uint64(noOfDataSetClasses))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetClasses' field"))
	}
	m.DataSetClasses = dataSetClasses

	noOfDefaultSecurityKeyServices, err := ReadImplicitField[int32](ctx, "noOfDefaultSecurityKeyServices", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDefaultSecurityKeyServices' field"))
	}
	_ = noOfDefaultSecurityKeyServices

	defaultSecurityKeyServices, err := ReadCountArrayField[EndpointDescription](ctx, "defaultSecurityKeyServices", ReadComplex[EndpointDescription](ExtensionObjectDefinitionParseWithBufferProducer[EndpointDescription]((int32)(int32(314))), readBuffer), uint64(noOfDefaultSecurityKeyServices))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'defaultSecurityKeyServices' field"))
	}
	m.DefaultSecurityKeyServices = defaultSecurityKeyServices

	noOfSecurityGroups, err := ReadImplicitField[int32](ctx, "noOfSecurityGroups", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSecurityGroups' field"))
	}
	_ = noOfSecurityGroups

	securityGroups, err := ReadCountArrayField[SecurityGroupDataType](ctx, "securityGroups", ReadComplex[SecurityGroupDataType](ExtensionObjectDefinitionParseWithBufferProducer[SecurityGroupDataType]((int32)(int32(23603))), readBuffer), uint64(noOfSecurityGroups))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroups' field"))
	}
	m.SecurityGroups = securityGroups

	noOfPubSubKeyPushTargets, err := ReadImplicitField[int32](ctx, "noOfPubSubKeyPushTargets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfPubSubKeyPushTargets' field"))
	}
	_ = noOfPubSubKeyPushTargets

	pubSubKeyPushTargets, err := ReadCountArrayField[PubSubKeyPushTargetDataType](ctx, "pubSubKeyPushTargets", ReadComplex[PubSubKeyPushTargetDataType](ExtensionObjectDefinitionParseWithBufferProducer[PubSubKeyPushTargetDataType]((int32)(int32(25272))), readBuffer), uint64(noOfPubSubKeyPushTargets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pubSubKeyPushTargets' field"))
	}
	m.PubSubKeyPushTargets = pubSubKeyPushTargets

	configurationVersion, err := ReadSimpleField(ctx, "configurationVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationVersion' field"))
	}
	m.ConfigurationVersion = configurationVersion

	noOfConfigurationProperties, err := ReadImplicitField[int32](ctx, "noOfConfigurationProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConfigurationProperties' field"))
	}
	_ = noOfConfigurationProperties

	configurationProperties, err := ReadCountArrayField[KeyValuePair](ctx, "configurationProperties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfConfigurationProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationProperties' field"))
	}
	m.ConfigurationProperties = configurationProperties

	if closeErr := readBuffer.CloseContext("PubSubConfiguration2DataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfiguration2DataType")
	}

	return m, nil
}

func (m *_PubSubConfiguration2DataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfiguration2DataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfiguration2DataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfiguration2DataType")
		}
		noOfPublishedDataSets := int32(utils.InlineIf(bool((m.GetPublishedDataSets()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetPublishedDataSets()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfPublishedDataSets", noOfPublishedDataSets, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfPublishedDataSets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "publishedDataSets", m.GetPublishedDataSets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'publishedDataSets' field")
		}
		noOfConnections := int32(utils.InlineIf(bool((m.GetConnections()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetConnections()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfConnections", noOfConnections, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConnections' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "connections", m.GetConnections(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'connections' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}
		noOfSubscribedDataSets := int32(utils.InlineIf(bool((m.GetSubscribedDataSets()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSubscribedDataSets()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSubscribedDataSets", noOfSubscribedDataSets, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSubscribedDataSets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "subscribedDataSets", m.GetSubscribedDataSets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subscribedDataSets' field")
		}
		noOfDataSetClasses := int32(utils.InlineIf(bool((m.GetDataSetClasses()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDataSetClasses()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDataSetClasses", noOfDataSetClasses, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataSetClasses' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataSetClasses", m.GetDataSetClasses(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetClasses' field")
		}
		noOfDefaultSecurityKeyServices := int32(utils.InlineIf(bool((m.GetDefaultSecurityKeyServices()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDefaultSecurityKeyServices()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDefaultSecurityKeyServices", noOfDefaultSecurityKeyServices, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDefaultSecurityKeyServices' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "defaultSecurityKeyServices", m.GetDefaultSecurityKeyServices(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'defaultSecurityKeyServices' field")
		}
		noOfSecurityGroups := int32(utils.InlineIf(bool((m.GetSecurityGroups()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSecurityGroups()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSecurityGroups", noOfSecurityGroups, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSecurityGroups' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "securityGroups", m.GetSecurityGroups(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'securityGroups' field")
		}
		noOfPubSubKeyPushTargets := int32(utils.InlineIf(bool((m.GetPubSubKeyPushTargets()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetPubSubKeyPushTargets()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfPubSubKeyPushTargets", noOfPubSubKeyPushTargets, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfPubSubKeyPushTargets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "pubSubKeyPushTargets", m.GetPubSubKeyPushTargets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'pubSubKeyPushTargets' field")
		}

		if err := WriteSimpleField[uint32](ctx, "configurationVersion", m.GetConfigurationVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationVersion' field")
		}
		noOfConfigurationProperties := int32(utils.InlineIf(bool((m.GetConfigurationProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetConfigurationProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfConfigurationProperties", noOfConfigurationProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConfigurationProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "configurationProperties", m.GetConfigurationProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationProperties' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfiguration2DataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfiguration2DataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfiguration2DataType) IsPubSubConfiguration2DataType() {}

func (m *_PubSubConfiguration2DataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PubSubConfiguration2DataType) deepCopy() *_PubSubConfiguration2DataType {
	if m == nil {
		return nil
	}
	_PubSubConfiguration2DataTypeCopy := &_PubSubConfiguration2DataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[PublishedDataSetDataType, PublishedDataSetDataType](m.PublishedDataSets),
		utils.DeepCopySlice[PubSubConnectionDataType, PubSubConnectionDataType](m.Connections),
		m.Enabled,
		utils.DeepCopySlice[StandaloneSubscribedDataSetDataType, StandaloneSubscribedDataSetDataType](m.SubscribedDataSets),
		utils.DeepCopySlice[DataSetMetaDataType, DataSetMetaDataType](m.DataSetClasses),
		utils.DeepCopySlice[EndpointDescription, EndpointDescription](m.DefaultSecurityKeyServices),
		utils.DeepCopySlice[SecurityGroupDataType, SecurityGroupDataType](m.SecurityGroups),
		utils.DeepCopySlice[PubSubKeyPushTargetDataType, PubSubKeyPushTargetDataType](m.PubSubKeyPushTargets),
		m.ConfigurationVersion,
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.ConfigurationProperties),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubConfiguration2DataTypeCopy
}

func (m *_PubSubConfiguration2DataType) String() string {
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
