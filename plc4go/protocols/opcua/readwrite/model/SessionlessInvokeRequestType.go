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

// SessionlessInvokeRequestType is the corresponding interface of SessionlessInvokeRequestType
type SessionlessInvokeRequestType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetUrisVersion returns UrisVersion (property field)
	GetUrisVersion() uint32
	// GetNamespaceUris returns NamespaceUris (property field)
	GetNamespaceUris() []PascalString
	// GetServerUris returns ServerUris (property field)
	GetServerUris() []PascalString
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetServiceId returns ServiceId (property field)
	GetServiceId() uint32
	// IsSessionlessInvokeRequestType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSessionlessInvokeRequestType()
	// CreateBuilder creates a SessionlessInvokeRequestTypeBuilder
	CreateSessionlessInvokeRequestTypeBuilder() SessionlessInvokeRequestTypeBuilder
}

// _SessionlessInvokeRequestType is the data-structure of this message
type _SessionlessInvokeRequestType struct {
	ExtensionObjectDefinitionContract
	UrisVersion   uint32
	NamespaceUris []PascalString
	ServerUris    []PascalString
	LocaleIds     []PascalString
	ServiceId     uint32
}

var _ SessionlessInvokeRequestType = (*_SessionlessInvokeRequestType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SessionlessInvokeRequestType)(nil)

// NewSessionlessInvokeRequestType factory function for _SessionlessInvokeRequestType
func NewSessionlessInvokeRequestType(urisVersion uint32, namespaceUris []PascalString, serverUris []PascalString, localeIds []PascalString, serviceId uint32) *_SessionlessInvokeRequestType {
	_result := &_SessionlessInvokeRequestType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		UrisVersion:                       urisVersion,
		NamespaceUris:                     namespaceUris,
		ServerUris:                        serverUris,
		LocaleIds:                         localeIds,
		ServiceId:                         serviceId,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SessionlessInvokeRequestTypeBuilder is a builder for SessionlessInvokeRequestType
type SessionlessInvokeRequestTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(urisVersion uint32, namespaceUris []PascalString, serverUris []PascalString, localeIds []PascalString, serviceId uint32) SessionlessInvokeRequestTypeBuilder
	// WithUrisVersion adds UrisVersion (property field)
	WithUrisVersion(uint32) SessionlessInvokeRequestTypeBuilder
	// WithNamespaceUris adds NamespaceUris (property field)
	WithNamespaceUris(...PascalString) SessionlessInvokeRequestTypeBuilder
	// WithServerUris adds ServerUris (property field)
	WithServerUris(...PascalString) SessionlessInvokeRequestTypeBuilder
	// WithLocaleIds adds LocaleIds (property field)
	WithLocaleIds(...PascalString) SessionlessInvokeRequestTypeBuilder
	// WithServiceId adds ServiceId (property field)
	WithServiceId(uint32) SessionlessInvokeRequestTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the SessionlessInvokeRequestType or returns an error if something is wrong
	Build() (SessionlessInvokeRequestType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SessionlessInvokeRequestType
}

// NewSessionlessInvokeRequestTypeBuilder() creates a SessionlessInvokeRequestTypeBuilder
func NewSessionlessInvokeRequestTypeBuilder() SessionlessInvokeRequestTypeBuilder {
	return &_SessionlessInvokeRequestTypeBuilder{_SessionlessInvokeRequestType: new(_SessionlessInvokeRequestType)}
}

type _SessionlessInvokeRequestTypeBuilder struct {
	*_SessionlessInvokeRequestType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SessionlessInvokeRequestTypeBuilder) = (*_SessionlessInvokeRequestTypeBuilder)(nil)

func (b *_SessionlessInvokeRequestTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._SessionlessInvokeRequestType
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithMandatoryFields(urisVersion uint32, namespaceUris []PascalString, serverUris []PascalString, localeIds []PascalString, serviceId uint32) SessionlessInvokeRequestTypeBuilder {
	return b.WithUrisVersion(urisVersion).WithNamespaceUris(namespaceUris...).WithServerUris(serverUris...).WithLocaleIds(localeIds...).WithServiceId(serviceId)
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithUrisVersion(urisVersion uint32) SessionlessInvokeRequestTypeBuilder {
	b.UrisVersion = urisVersion
	return b
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithNamespaceUris(namespaceUris ...PascalString) SessionlessInvokeRequestTypeBuilder {
	b.NamespaceUris = namespaceUris
	return b
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithServerUris(serverUris ...PascalString) SessionlessInvokeRequestTypeBuilder {
	b.ServerUris = serverUris
	return b
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithLocaleIds(localeIds ...PascalString) SessionlessInvokeRequestTypeBuilder {
	b.LocaleIds = localeIds
	return b
}

func (b *_SessionlessInvokeRequestTypeBuilder) WithServiceId(serviceId uint32) SessionlessInvokeRequestTypeBuilder {
	b.ServiceId = serviceId
	return b
}

func (b *_SessionlessInvokeRequestTypeBuilder) Build() (SessionlessInvokeRequestType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SessionlessInvokeRequestType.deepCopy(), nil
}

func (b *_SessionlessInvokeRequestTypeBuilder) MustBuild() SessionlessInvokeRequestType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SessionlessInvokeRequestTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_SessionlessInvokeRequestTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SessionlessInvokeRequestTypeBuilder) DeepCopy() any {
	_copy := b.CreateSessionlessInvokeRequestTypeBuilder().(*_SessionlessInvokeRequestTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSessionlessInvokeRequestTypeBuilder creates a SessionlessInvokeRequestTypeBuilder
func (b *_SessionlessInvokeRequestType) CreateSessionlessInvokeRequestTypeBuilder() SessionlessInvokeRequestTypeBuilder {
	if b == nil {
		return NewSessionlessInvokeRequestTypeBuilder()
	}
	return &_SessionlessInvokeRequestTypeBuilder{_SessionlessInvokeRequestType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SessionlessInvokeRequestType) GetExtensionId() int32 {
	return int32(15903)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SessionlessInvokeRequestType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SessionlessInvokeRequestType) GetUrisVersion() uint32 {
	return m.UrisVersion
}

func (m *_SessionlessInvokeRequestType) GetNamespaceUris() []PascalString {
	return m.NamespaceUris
}

func (m *_SessionlessInvokeRequestType) GetServerUris() []PascalString {
	return m.ServerUris
}

func (m *_SessionlessInvokeRequestType) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_SessionlessInvokeRequestType) GetServiceId() uint32 {
	return m.ServiceId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSessionlessInvokeRequestType(structType any) SessionlessInvokeRequestType {
	if casted, ok := structType.(SessionlessInvokeRequestType); ok {
		return casted
	}
	if casted, ok := structType.(*SessionlessInvokeRequestType); ok {
		return *casted
	}
	return nil
}

func (m *_SessionlessInvokeRequestType) GetTypeName() string {
	return "SessionlessInvokeRequestType"
}

func (m *_SessionlessInvokeRequestType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (urisVersion)
	lengthInBits += 32

	// Implicit Field (noOfNamespaceUris)
	lengthInBits += 32

	// Array field
	if len(m.NamespaceUris) > 0 {
		for _curItem, element := range m.NamespaceUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NamespaceUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfServerUris)
	lengthInBits += 32

	// Array field
	if len(m.ServerUris) > 0 {
		for _curItem, element := range m.ServerUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (serviceId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SessionlessInvokeRequestType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SessionlessInvokeRequestType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__sessionlessInvokeRequestType SessionlessInvokeRequestType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SessionlessInvokeRequestType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SessionlessInvokeRequestType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	urisVersion, err := ReadSimpleField(ctx, "urisVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'urisVersion' field"))
	}
	m.UrisVersion = urisVersion

	noOfNamespaceUris, err := ReadImplicitField[int32](ctx, "noOfNamespaceUris", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNamespaceUris' field"))
	}
	_ = noOfNamespaceUris

	namespaceUris, err := ReadCountArrayField[PascalString](ctx, "namespaceUris", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfNamespaceUris))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceUris' field"))
	}
	m.NamespaceUris = namespaceUris

	noOfServerUris, err := ReadImplicitField[int32](ctx, "noOfServerUris", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerUris' field"))
	}
	_ = noOfServerUris

	serverUris, err := ReadCountArrayField[PascalString](ctx, "serverUris", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfServerUris))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverUris' field"))
	}
	m.ServerUris = serverUris

	noOfLocaleIds, err := ReadImplicitField[int32](ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	_ = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	serviceId, err := ReadSimpleField(ctx, "serviceId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceId' field"))
	}
	m.ServiceId = serviceId

	if closeErr := readBuffer.CloseContext("SessionlessInvokeRequestType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SessionlessInvokeRequestType")
	}

	return m, nil
}

func (m *_SessionlessInvokeRequestType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SessionlessInvokeRequestType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SessionlessInvokeRequestType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SessionlessInvokeRequestType")
		}

		if err := WriteSimpleField[uint32](ctx, "urisVersion", m.GetUrisVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'urisVersion' field")
		}
		noOfNamespaceUris := int32(utils.InlineIf(bool((m.GetNamespaceUris()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetNamespaceUris()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfNamespaceUris", noOfNamespaceUris, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNamespaceUris' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "namespaceUris", m.GetNamespaceUris(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceUris' field")
		}
		noOfServerUris := int32(utils.InlineIf(bool((m.GetServerUris()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetServerUris()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfServerUris", noOfServerUris, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerUris' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverUris", m.GetServerUris(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverUris' field")
		}
		noOfLocaleIds := int32(utils.InlineIf(bool((m.GetLocaleIds()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLocaleIds()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLocaleIds", noOfLocaleIds, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[uint32](ctx, "serviceId", m.GetServiceId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceId' field")
		}

		if popErr := writeBuffer.PopContext("SessionlessInvokeRequestType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SessionlessInvokeRequestType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SessionlessInvokeRequestType) IsSessionlessInvokeRequestType() {}

func (m *_SessionlessInvokeRequestType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SessionlessInvokeRequestType) deepCopy() *_SessionlessInvokeRequestType {
	if m == nil {
		return nil
	}
	_SessionlessInvokeRequestTypeCopy := &_SessionlessInvokeRequestType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.UrisVersion,
		utils.DeepCopySlice[PascalString, PascalString](m.NamespaceUris),
		utils.DeepCopySlice[PascalString, PascalString](m.ServerUris),
		utils.DeepCopySlice[PascalString, PascalString](m.LocaleIds),
		m.ServiceId,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SessionlessInvokeRequestTypeCopy
}

func (m *_SessionlessInvokeRequestType) String() string {
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
