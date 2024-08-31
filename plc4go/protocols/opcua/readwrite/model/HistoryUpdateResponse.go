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

// HistoryUpdateResponse is the corresponding interface of HistoryUpdateResponse
type HistoryUpdateResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// HistoryUpdateResponseExactly can be used when we want exactly this type and not a type which fulfills HistoryUpdateResponse.
// This is useful for switch cases.
type HistoryUpdateResponseExactly interface {
	HistoryUpdateResponse
	isHistoryUpdateResponse() bool
}

// _HistoryUpdateResponse is the data-structure of this message
type _HistoryUpdateResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateResponse) GetIdentifier() string {
	return "703"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryUpdateResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_HistoryUpdateResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_HistoryUpdateResponse) GetResults() []ExtensionObjectDefinition {
	return m.Results
}

func (m *_HistoryUpdateResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_HistoryUpdateResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryUpdateResponse factory function for _HistoryUpdateResponse
func NewHistoryUpdateResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_HistoryUpdateResponse {
	_result := &_HistoryUpdateResponse{
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryUpdateResponse(structType any) HistoryUpdateResponse {
	if casted, ok := structType.(HistoryUpdateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateResponse) GetTypeName() string {
	return "HistoryUpdateResponse"
}

func (m *_HistoryUpdateResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
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

	// Simple field (noOfDiagnosticInfos)
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

func (m *_HistoryUpdateResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryUpdateResponseParse(ctx context.Context, theBytes []byte, identifier string) (HistoryUpdateResponse, error) {
	return HistoryUpdateResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryUpdateResponseParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (HistoryUpdateResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HistoryUpdateResponse, error) {
		return HistoryUpdateResponseParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func HistoryUpdateResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryUpdateResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryUpdateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}

	noOfResults, err := ReadSimpleField(ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}

	results, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "results", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("697")), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}

	if closeErr := readBuffer.CloseContext("HistoryUpdateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateResponse")
	}

	// Create a partially initialized instance
	_child := &_HistoryUpdateResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryUpdateResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfResults)
		noOfResults := int32(m.GetNoOfResults())
		_noOfResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfResults", 32, int32((noOfResults)))
		if _noOfResultsErr != nil {
			return errors.Wrap(_noOfResultsErr, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, int32((noOfDiagnosticInfos)))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateResponse) isHistoryUpdateResponse() bool {
	return true
}

func (m *_HistoryUpdateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
