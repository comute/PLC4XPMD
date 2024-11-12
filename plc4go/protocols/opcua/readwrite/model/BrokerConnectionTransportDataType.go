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

// BrokerConnectionTransportDataType is the corresponding interface of BrokerConnectionTransportDataType
type BrokerConnectionTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResourceUri returns ResourceUri (property field)
	GetResourceUri() PascalString
	// GetAuthenticationProfileUri returns AuthenticationProfileUri (property field)
	GetAuthenticationProfileUri() PascalString
	// IsBrokerConnectionTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrokerConnectionTransportDataType()
	// CreateBuilder creates a BrokerConnectionTransportDataTypeBuilder
	CreateBrokerConnectionTransportDataTypeBuilder() BrokerConnectionTransportDataTypeBuilder
}

// _BrokerConnectionTransportDataType is the data-structure of this message
type _BrokerConnectionTransportDataType struct {
	ExtensionObjectDefinitionContract
	ResourceUri              PascalString
	AuthenticationProfileUri PascalString
}

var _ BrokerConnectionTransportDataType = (*_BrokerConnectionTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrokerConnectionTransportDataType)(nil)

// NewBrokerConnectionTransportDataType factory function for _BrokerConnectionTransportDataType
func NewBrokerConnectionTransportDataType(resourceUri PascalString, authenticationProfileUri PascalString) *_BrokerConnectionTransportDataType {
	if resourceUri == nil {
		panic("resourceUri of type PascalString for BrokerConnectionTransportDataType must not be nil")
	}
	if authenticationProfileUri == nil {
		panic("authenticationProfileUri of type PascalString for BrokerConnectionTransportDataType must not be nil")
	}
	_result := &_BrokerConnectionTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResourceUri:                       resourceUri,
		AuthenticationProfileUri:          authenticationProfileUri,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrokerConnectionTransportDataTypeBuilder is a builder for BrokerConnectionTransportDataType
type BrokerConnectionTransportDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(resourceUri PascalString, authenticationProfileUri PascalString) BrokerConnectionTransportDataTypeBuilder
	// WithResourceUri adds ResourceUri (property field)
	WithResourceUri(PascalString) BrokerConnectionTransportDataTypeBuilder
	// WithResourceUriBuilder adds ResourceUri (property field) which is build by the builder
	WithResourceUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerConnectionTransportDataTypeBuilder
	// WithAuthenticationProfileUri adds AuthenticationProfileUri (property field)
	WithAuthenticationProfileUri(PascalString) BrokerConnectionTransportDataTypeBuilder
	// WithAuthenticationProfileUriBuilder adds AuthenticationProfileUri (property field) which is build by the builder
	WithAuthenticationProfileUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerConnectionTransportDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the BrokerConnectionTransportDataType or returns an error if something is wrong
	Build() (BrokerConnectionTransportDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrokerConnectionTransportDataType
}

// NewBrokerConnectionTransportDataTypeBuilder() creates a BrokerConnectionTransportDataTypeBuilder
func NewBrokerConnectionTransportDataTypeBuilder() BrokerConnectionTransportDataTypeBuilder {
	return &_BrokerConnectionTransportDataTypeBuilder{_BrokerConnectionTransportDataType: new(_BrokerConnectionTransportDataType)}
}

type _BrokerConnectionTransportDataTypeBuilder struct {
	*_BrokerConnectionTransportDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrokerConnectionTransportDataTypeBuilder) = (*_BrokerConnectionTransportDataTypeBuilder)(nil)

func (b *_BrokerConnectionTransportDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._BrokerConnectionTransportDataType
}

func (b *_BrokerConnectionTransportDataTypeBuilder) WithMandatoryFields(resourceUri PascalString, authenticationProfileUri PascalString) BrokerConnectionTransportDataTypeBuilder {
	return b.WithResourceUri(resourceUri).WithAuthenticationProfileUri(authenticationProfileUri)
}

func (b *_BrokerConnectionTransportDataTypeBuilder) WithResourceUri(resourceUri PascalString) BrokerConnectionTransportDataTypeBuilder {
	b.ResourceUri = resourceUri
	return b
}

func (b *_BrokerConnectionTransportDataTypeBuilder) WithResourceUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerConnectionTransportDataTypeBuilder {
	builder := builderSupplier(b.ResourceUri.CreatePascalStringBuilder())
	var err error
	b.ResourceUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerConnectionTransportDataTypeBuilder) WithAuthenticationProfileUri(authenticationProfileUri PascalString) BrokerConnectionTransportDataTypeBuilder {
	b.AuthenticationProfileUri = authenticationProfileUri
	return b
}

func (b *_BrokerConnectionTransportDataTypeBuilder) WithAuthenticationProfileUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerConnectionTransportDataTypeBuilder {
	builder := builderSupplier(b.AuthenticationProfileUri.CreatePascalStringBuilder())
	var err error
	b.AuthenticationProfileUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerConnectionTransportDataTypeBuilder) Build() (BrokerConnectionTransportDataType, error) {
	if b.ResourceUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resourceUri' not set"))
	}
	if b.AuthenticationProfileUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'authenticationProfileUri' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrokerConnectionTransportDataType.deepCopy(), nil
}

func (b *_BrokerConnectionTransportDataTypeBuilder) MustBuild() BrokerConnectionTransportDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BrokerConnectionTransportDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_BrokerConnectionTransportDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrokerConnectionTransportDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateBrokerConnectionTransportDataTypeBuilder().(*_BrokerConnectionTransportDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrokerConnectionTransportDataTypeBuilder creates a BrokerConnectionTransportDataTypeBuilder
func (b *_BrokerConnectionTransportDataType) CreateBrokerConnectionTransportDataTypeBuilder() BrokerConnectionTransportDataTypeBuilder {
	if b == nil {
		return NewBrokerConnectionTransportDataTypeBuilder()
	}
	return &_BrokerConnectionTransportDataTypeBuilder{_BrokerConnectionTransportDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrokerConnectionTransportDataType) GetExtensionId() int32 {
	return int32(15009)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrokerConnectionTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrokerConnectionTransportDataType) GetResourceUri() PascalString {
	return m.ResourceUri
}

func (m *_BrokerConnectionTransportDataType) GetAuthenticationProfileUri() PascalString {
	return m.AuthenticationProfileUri
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrokerConnectionTransportDataType(structType any) BrokerConnectionTransportDataType {
	if casted, ok := structType.(BrokerConnectionTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*BrokerConnectionTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_BrokerConnectionTransportDataType) GetTypeName() string {
	return "BrokerConnectionTransportDataType"
}

func (m *_BrokerConnectionTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (resourceUri)
	lengthInBits += m.ResourceUri.GetLengthInBits(ctx)

	// Simple field (authenticationProfileUri)
	lengthInBits += m.AuthenticationProfileUri.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BrokerConnectionTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrokerConnectionTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__brokerConnectionTransportDataType BrokerConnectionTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrokerConnectionTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrokerConnectionTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resourceUri, err := ReadSimpleField[PascalString](ctx, "resourceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resourceUri' field"))
	}
	m.ResourceUri = resourceUri

	authenticationProfileUri, err := ReadSimpleField[PascalString](ctx, "authenticationProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationProfileUri' field"))
	}
	m.AuthenticationProfileUri = authenticationProfileUri

	if closeErr := readBuffer.CloseContext("BrokerConnectionTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrokerConnectionTransportDataType")
	}

	return m, nil
}

func (m *_BrokerConnectionTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrokerConnectionTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrokerConnectionTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrokerConnectionTransportDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "resourceUri", m.GetResourceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resourceUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "authenticationProfileUri", m.GetAuthenticationProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationProfileUri' field")
		}

		if popErr := writeBuffer.PopContext("BrokerConnectionTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrokerConnectionTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrokerConnectionTransportDataType) IsBrokerConnectionTransportDataType() {}

func (m *_BrokerConnectionTransportDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrokerConnectionTransportDataType) deepCopy() *_BrokerConnectionTransportDataType {
	if m == nil {
		return nil
	}
	_BrokerConnectionTransportDataTypeCopy := &_BrokerConnectionTransportDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.ResourceUri),
		utils.DeepCopy[PascalString](m.AuthenticationProfileUri),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrokerConnectionTransportDataTypeCopy
}

func (m *_BrokerConnectionTransportDataType) String() string {
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
