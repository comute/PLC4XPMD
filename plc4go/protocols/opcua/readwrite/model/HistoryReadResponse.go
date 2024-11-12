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

// HistoryReadResponse is the corresponding interface of HistoryReadResponse
type HistoryReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetResults returns Results (property field)
	GetResults() []HistoryReadResult
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsHistoryReadResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryReadResponse()
	// CreateBuilder creates a HistoryReadResponseBuilder
	CreateHistoryReadResponseBuilder() HistoryReadResponseBuilder
}

// _HistoryReadResponse is the data-structure of this message
type _HistoryReadResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader  ResponseHeader
	Results         []HistoryReadResult
	DiagnosticInfos []DiagnosticInfo
}

var _ HistoryReadResponse = (*_HistoryReadResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadResponse)(nil)

// NewHistoryReadResponse factory function for _HistoryReadResponse
func NewHistoryReadResponse(responseHeader ResponseHeader, results []HistoryReadResult, diagnosticInfos []DiagnosticInfo) *_HistoryReadResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for HistoryReadResponse must not be nil")
	}
	_result := &_HistoryReadResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		Results:                           results,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryReadResponseBuilder is a builder for HistoryReadResponse
type HistoryReadResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, results []HistoryReadResult, diagnosticInfos []DiagnosticInfo) HistoryReadResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) HistoryReadResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) HistoryReadResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...HistoryReadResult) HistoryReadResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) HistoryReadResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the HistoryReadResponse or returns an error if something is wrong
	Build() (HistoryReadResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryReadResponse
}

// NewHistoryReadResponseBuilder() creates a HistoryReadResponseBuilder
func NewHistoryReadResponseBuilder() HistoryReadResponseBuilder {
	return &_HistoryReadResponseBuilder{_HistoryReadResponse: new(_HistoryReadResponse)}
}

type _HistoryReadResponseBuilder struct {
	*_HistoryReadResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryReadResponseBuilder) = (*_HistoryReadResponseBuilder)(nil)

func (b *_HistoryReadResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._HistoryReadResponse
}

func (b *_HistoryReadResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, results []HistoryReadResult, diagnosticInfos []DiagnosticInfo) HistoryReadResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithResults(results...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_HistoryReadResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) HistoryReadResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_HistoryReadResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) HistoryReadResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_HistoryReadResponseBuilder) WithResults(results ...HistoryReadResult) HistoryReadResponseBuilder {
	b.Results = results
	return b
}

func (b *_HistoryReadResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) HistoryReadResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_HistoryReadResponseBuilder) Build() (HistoryReadResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryReadResponse.deepCopy(), nil
}

func (b *_HistoryReadResponseBuilder) MustBuild() HistoryReadResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HistoryReadResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_HistoryReadResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryReadResponseBuilder) DeepCopy() any {
	_copy := b.CreateHistoryReadResponseBuilder().(*_HistoryReadResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryReadResponseBuilder creates a HistoryReadResponseBuilder
func (b *_HistoryReadResponse) CreateHistoryReadResponseBuilder() HistoryReadResponseBuilder {
	if b == nil {
		return NewHistoryReadResponseBuilder()
	}
	return &_HistoryReadResponseBuilder{_HistoryReadResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadResponse) GetExtensionId() int32 {
	return int32(667)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_HistoryReadResponse) GetResults() []HistoryReadResult {
	return m.Results
}

func (m *_HistoryReadResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryReadResponse(structType any) HistoryReadResponse {
	if casted, ok := structType.(HistoryReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadResponse) GetTypeName() string {
	return "HistoryReadResponse"
}

func (m *_HistoryReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyReadResponse HistoryReadResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadImplicitField[int32](ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	_ = noOfResults

	results, err := ReadCountArrayField[HistoryReadResult](ctx, "results", ReadComplex[HistoryReadResult](ExtensionObjectDefinitionParseWithBufferProducer[HistoryReadResult]((int32)(int32(640))), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	_ = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("HistoryReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadResponse")
	}

	return m, nil
}

func (m *_HistoryReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfResults := int32(utils.InlineIf(bool((m.GetResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfResults", noOfResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadResponse) IsHistoryReadResponse() {}

func (m *_HistoryReadResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryReadResponse) deepCopy() *_HistoryReadResponse {
	if m == nil {
		return nil
	}
	_HistoryReadResponseCopy := &_HistoryReadResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		utils.DeepCopySlice[HistoryReadResult, HistoryReadResult](m.Results),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryReadResponseCopy
}

func (m *_HistoryReadResponse) String() string {
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
