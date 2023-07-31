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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// HistoryReadResult is the corresponding interface of HistoryReadResult
type HistoryReadResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// GetHistoryData returns HistoryData (property field)
	GetHistoryData() ExtensionObject
}

// HistoryReadResultExactly can be used when we want exactly this type and not a type which fulfills HistoryReadResult.
// This is useful for switch cases.
type HistoryReadResultExactly interface {
	HistoryReadResult
	isHistoryReadResult() bool
}

// _HistoryReadResult is the data-structure of this message
type _HistoryReadResult struct {
	*_ExtensionObjectDefinition
	StatusCode        StatusCode
	ContinuationPoint PascalByteString
	HistoryData       ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadResult) GetIdentifier() string {
	return "640"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryReadResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_HistoryReadResult) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

func (m *_HistoryReadResult) GetHistoryData() ExtensionObject {
	return m.HistoryData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryReadResult factory function for _HistoryReadResult
func NewHistoryReadResult(statusCode StatusCode, continuationPoint PascalByteString, historyData ExtensionObject) *_HistoryReadResult {
	_result := &_HistoryReadResult{
		StatusCode:                 statusCode,
		ContinuationPoint:          continuationPoint,
		HistoryData:                historyData,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryReadResult(structType any) HistoryReadResult {
	if casted, ok := structType.(HistoryReadResult); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadResult); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadResult) GetTypeName() string {
	return "HistoryReadResult"
}

func (m *_HistoryReadResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	// Simple field (historyData)
	lengthInBits += m.HistoryData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_HistoryReadResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryReadResultParse(ctx context.Context, theBytes []byte, identifier string) (HistoryReadResult, error) {
	return HistoryReadResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryReadResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryReadResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryReadResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of HistoryReadResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (continuationPoint)
	if pullErr := readBuffer.PullContext("continuationPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for continuationPoint")
	}
	_continuationPoint, _continuationPointErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _continuationPointErr != nil {
		return nil, errors.Wrap(_continuationPointErr, "Error parsing 'continuationPoint' field of HistoryReadResult")
	}
	continuationPoint := _continuationPoint.(PascalByteString)
	if closeErr := readBuffer.CloseContext("continuationPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for continuationPoint")
	}

	// Simple Field (historyData)
	if pullErr := readBuffer.PullContext("historyData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for historyData")
	}
	_historyData, _historyDataErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _historyDataErr != nil {
		return nil, errors.Wrap(_historyDataErr, "Error parsing 'historyData' field of HistoryReadResult")
	}
	historyData := _historyData.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("historyData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for historyData")
	}

	if closeErr := readBuffer.CloseContext("HistoryReadResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadResult")
	}

	// Create a partially initialized instance
	_child := &_HistoryReadResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		ContinuationPoint:          continuationPoint,
		HistoryData:                historyData,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryReadResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (continuationPoint)
		if pushErr := writeBuffer.PushContext("continuationPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for continuationPoint")
		}
		_continuationPointErr := writeBuffer.WriteSerializable(ctx, m.GetContinuationPoint())
		if popErr := writeBuffer.PopContext("continuationPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for continuationPoint")
		}
		if _continuationPointErr != nil {
			return errors.Wrap(_continuationPointErr, "Error serializing 'continuationPoint' field")
		}

		// Simple Field (historyData)
		if pushErr := writeBuffer.PushContext("historyData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for historyData")
		}
		_historyDataErr := writeBuffer.WriteSerializable(ctx, m.GetHistoryData())
		if popErr := writeBuffer.PopContext("historyData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for historyData")
		}
		if _historyDataErr != nil {
			return errors.Wrap(_historyDataErr, "Error serializing 'historyData' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadResult) isHistoryReadResult() bool {
	return true
}

func (m *_HistoryReadResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}