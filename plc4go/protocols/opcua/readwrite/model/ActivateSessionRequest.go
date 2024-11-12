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

// ActivateSessionRequest is the corresponding interface of ActivateSessionRequest
type ActivateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetClientSignature returns ClientSignature (property field)
	GetClientSignature() SignatureData
	// GetClientSoftwareCertificates returns ClientSoftwareCertificates (property field)
	GetClientSoftwareCertificates() []SignedSoftwareCertificate
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetUserIdentityToken returns UserIdentityToken (property field)
	GetUserIdentityToken() ExtensionObject
	// GetUserTokenSignature returns UserTokenSignature (property field)
	GetUserTokenSignature() SignatureData
	// IsActivateSessionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsActivateSessionRequest()
	// CreateBuilder creates a ActivateSessionRequestBuilder
	CreateActivateSessionRequestBuilder() ActivateSessionRequestBuilder
}

// _ActivateSessionRequest is the data-structure of this message
type _ActivateSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader              RequestHeader
	ClientSignature            SignatureData
	ClientSoftwareCertificates []SignedSoftwareCertificate
	LocaleIds                  []PascalString
	UserIdentityToken          ExtensionObject
	UserTokenSignature         SignatureData
}

var _ ActivateSessionRequest = (*_ActivateSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ActivateSessionRequest)(nil)

// NewActivateSessionRequest factory function for _ActivateSessionRequest
func NewActivateSessionRequest(requestHeader RequestHeader, clientSignature SignatureData, clientSoftwareCertificates []SignedSoftwareCertificate, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature SignatureData) *_ActivateSessionRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for ActivateSessionRequest must not be nil")
	}
	if clientSignature == nil {
		panic("clientSignature of type SignatureData for ActivateSessionRequest must not be nil")
	}
	if userIdentityToken == nil {
		panic("userIdentityToken of type ExtensionObject for ActivateSessionRequest must not be nil")
	}
	if userTokenSignature == nil {
		panic("userTokenSignature of type SignatureData for ActivateSessionRequest must not be nil")
	}
	_result := &_ActivateSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ClientSignature:                   clientSignature,
		ClientSoftwareCertificates:        clientSoftwareCertificates,
		LocaleIds:                         localeIds,
		UserIdentityToken:                 userIdentityToken,
		UserTokenSignature:                userTokenSignature,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ActivateSessionRequestBuilder is a builder for ActivateSessionRequest
type ActivateSessionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, clientSignature SignatureData, clientSoftwareCertificates []SignedSoftwareCertificate, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature SignatureData) ActivateSessionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) ActivateSessionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) ActivateSessionRequestBuilder
	// WithClientSignature adds ClientSignature (property field)
	WithClientSignature(SignatureData) ActivateSessionRequestBuilder
	// WithClientSignatureBuilder adds ClientSignature (property field) which is build by the builder
	WithClientSignatureBuilder(func(SignatureDataBuilder) SignatureDataBuilder) ActivateSessionRequestBuilder
	// WithClientSoftwareCertificates adds ClientSoftwareCertificates (property field)
	WithClientSoftwareCertificates(...SignedSoftwareCertificate) ActivateSessionRequestBuilder
	// WithLocaleIds adds LocaleIds (property field)
	WithLocaleIds(...PascalString) ActivateSessionRequestBuilder
	// WithUserIdentityToken adds UserIdentityToken (property field)
	WithUserIdentityToken(ExtensionObject) ActivateSessionRequestBuilder
	// WithUserIdentityTokenBuilder adds UserIdentityToken (property field) which is build by the builder
	WithUserIdentityTokenBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) ActivateSessionRequestBuilder
	// WithUserTokenSignature adds UserTokenSignature (property field)
	WithUserTokenSignature(SignatureData) ActivateSessionRequestBuilder
	// WithUserTokenSignatureBuilder adds UserTokenSignature (property field) which is build by the builder
	WithUserTokenSignatureBuilder(func(SignatureDataBuilder) SignatureDataBuilder) ActivateSessionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ActivateSessionRequest or returns an error if something is wrong
	Build() (ActivateSessionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ActivateSessionRequest
}

// NewActivateSessionRequestBuilder() creates a ActivateSessionRequestBuilder
func NewActivateSessionRequestBuilder() ActivateSessionRequestBuilder {
	return &_ActivateSessionRequestBuilder{_ActivateSessionRequest: new(_ActivateSessionRequest)}
}

type _ActivateSessionRequestBuilder struct {
	*_ActivateSessionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ActivateSessionRequestBuilder) = (*_ActivateSessionRequestBuilder)(nil)

func (b *_ActivateSessionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ActivateSessionRequest
}

func (b *_ActivateSessionRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, clientSignature SignatureData, clientSoftwareCertificates []SignedSoftwareCertificate, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature SignatureData) ActivateSessionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithClientSignature(clientSignature).WithClientSoftwareCertificates(clientSoftwareCertificates...).WithLocaleIds(localeIds...).WithUserIdentityToken(userIdentityToken).WithUserTokenSignature(userTokenSignature)
}

func (b *_ActivateSessionRequestBuilder) WithRequestHeader(requestHeader RequestHeader) ActivateSessionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_ActivateSessionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSignature(clientSignature SignatureData) ActivateSessionRequestBuilder {
	b.ClientSignature = clientSignature
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSignatureBuilder(builderSupplier func(SignatureDataBuilder) SignatureDataBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.ClientSignature.CreateSignatureDataBuilder())
	var err error
	b.ClientSignature, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SignatureDataBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSoftwareCertificates(clientSoftwareCertificates ...SignedSoftwareCertificate) ActivateSessionRequestBuilder {
	b.ClientSoftwareCertificates = clientSoftwareCertificates
	return b
}

func (b *_ActivateSessionRequestBuilder) WithLocaleIds(localeIds ...PascalString) ActivateSessionRequestBuilder {
	b.LocaleIds = localeIds
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserIdentityToken(userIdentityToken ExtensionObject) ActivateSessionRequestBuilder {
	b.UserIdentityToken = userIdentityToken
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserIdentityTokenBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.UserIdentityToken.CreateExtensionObjectBuilder())
	var err error
	b.UserIdentityToken, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserTokenSignature(userTokenSignature SignatureData) ActivateSessionRequestBuilder {
	b.UserTokenSignature = userTokenSignature
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserTokenSignatureBuilder(builderSupplier func(SignatureDataBuilder) SignatureDataBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.UserTokenSignature.CreateSignatureDataBuilder())
	var err error
	b.UserTokenSignature, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SignatureDataBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) Build() (ActivateSessionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.ClientSignature == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientSignature' not set"))
	}
	if b.UserIdentityToken == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userIdentityToken' not set"))
	}
	if b.UserTokenSignature == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userTokenSignature' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ActivateSessionRequest.deepCopy(), nil
}

func (b *_ActivateSessionRequestBuilder) MustBuild() ActivateSessionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ActivateSessionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ActivateSessionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ActivateSessionRequestBuilder) DeepCopy() any {
	_copy := b.CreateActivateSessionRequestBuilder().(*_ActivateSessionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateActivateSessionRequestBuilder creates a ActivateSessionRequestBuilder
func (b *_ActivateSessionRequest) CreateActivateSessionRequestBuilder() ActivateSessionRequestBuilder {
	if b == nil {
		return NewActivateSessionRequestBuilder()
	}
	return &_ActivateSessionRequestBuilder{_ActivateSessionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ActivateSessionRequest) GetExtensionId() int32 {
	return int32(467)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ActivateSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ActivateSessionRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_ActivateSessionRequest) GetClientSignature() SignatureData {
	return m.ClientSignature
}

func (m *_ActivateSessionRequest) GetClientSoftwareCertificates() []SignedSoftwareCertificate {
	return m.ClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_ActivateSessionRequest) GetUserIdentityToken() ExtensionObject {
	return m.UserIdentityToken
}

func (m *_ActivateSessionRequest) GetUserTokenSignature() SignatureData {
	return m.UserTokenSignature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastActivateSessionRequest(structType any) ActivateSessionRequest {
	if casted, ok := structType.(ActivateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ActivateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ActivateSessionRequest) GetTypeName() string {
	return "ActivateSessionRequest"
}

func (m *_ActivateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientSignature)
	lengthInBits += m.ClientSignature.GetLengthInBits(ctx)

	// Implicit Field (noOfClientSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ClientSoftwareCertificates) > 0 {
		for _curItem, element := range m.ClientSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ClientSoftwareCertificates), _curItem)
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

	// Simple field (userIdentityToken)
	lengthInBits += m.UserIdentityToken.GetLengthInBits(ctx)

	// Simple field (userTokenSignature)
	lengthInBits += m.UserTokenSignature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ActivateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ActivateSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__activateSessionRequest ActivateSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ActivateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ActivateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	clientSignature, err := ReadSimpleField[SignatureData](ctx, "clientSignature", ReadComplex[SignatureData](ExtensionObjectDefinitionParseWithBufferProducer[SignatureData]((int32)(int32(458))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSignature' field"))
	}
	m.ClientSignature = clientSignature

	noOfClientSoftwareCertificates, err := ReadImplicitField[int32](ctx, "noOfClientSoftwareCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfClientSoftwareCertificates' field"))
	}
	_ = noOfClientSoftwareCertificates

	clientSoftwareCertificates, err := ReadCountArrayField[SignedSoftwareCertificate](ctx, "clientSoftwareCertificates", ReadComplex[SignedSoftwareCertificate](ExtensionObjectDefinitionParseWithBufferProducer[SignedSoftwareCertificate]((int32)(int32(346))), readBuffer), uint64(noOfClientSoftwareCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSoftwareCertificates' field"))
	}
	m.ClientSoftwareCertificates = clientSoftwareCertificates

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

	userIdentityToken, err := ReadSimpleField[ExtensionObject](ctx, "userIdentityToken", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userIdentityToken' field"))
	}
	m.UserIdentityToken = userIdentityToken

	userTokenSignature, err := ReadSimpleField[SignatureData](ctx, "userTokenSignature", ReadComplex[SignatureData](ExtensionObjectDefinitionParseWithBufferProducer[SignatureData]((int32)(int32(458))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userTokenSignature' field"))
	}
	m.UserTokenSignature = userTokenSignature

	if closeErr := readBuffer.CloseContext("ActivateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ActivateSessionRequest")
	}

	return m, nil
}

func (m *_ActivateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ActivateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ActivateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ActivateSessionRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[SignatureData](ctx, "clientSignature", m.GetClientSignature(), WriteComplex[SignatureData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSignature' field")
		}
		noOfClientSoftwareCertificates := int32(utils.InlineIf(bool((m.GetClientSoftwareCertificates()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetClientSoftwareCertificates()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfClientSoftwareCertificates", noOfClientSoftwareCertificates, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfClientSoftwareCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "clientSoftwareCertificates", m.GetClientSoftwareCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSoftwareCertificates' field")
		}
		noOfLocaleIds := int32(utils.InlineIf(bool((m.GetLocaleIds()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLocaleIds()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLocaleIds", noOfLocaleIds, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "userIdentityToken", m.GetUserIdentityToken(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userIdentityToken' field")
		}

		if err := WriteSimpleField[SignatureData](ctx, "userTokenSignature", m.GetUserTokenSignature(), WriteComplex[SignatureData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userTokenSignature' field")
		}

		if popErr := writeBuffer.PopContext("ActivateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ActivateSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ActivateSessionRequest) IsActivateSessionRequest() {}

func (m *_ActivateSessionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ActivateSessionRequest) deepCopy() *_ActivateSessionRequest {
	if m == nil {
		return nil
	}
	_ActivateSessionRequestCopy := &_ActivateSessionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopy[SignatureData](m.ClientSignature),
		utils.DeepCopySlice[SignedSoftwareCertificate, SignedSoftwareCertificate](m.ClientSoftwareCertificates),
		utils.DeepCopySlice[PascalString, PascalString](m.LocaleIds),
		utils.DeepCopy[ExtensionObject](m.UserIdentityToken),
		utils.DeepCopy[SignatureData](m.UserTokenSignature),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ActivateSessionRequestCopy
}

func (m *_ActivateSessionRequest) String() string {
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
