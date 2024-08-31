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

// SALDataErrorReporting is the corresponding interface of SALDataErrorReporting
type SALDataErrorReporting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetErrorReportingData returns ErrorReportingData (property field)
	GetErrorReportingData() ErrorReportingData
}

// SALDataErrorReportingExactly can be used when we want exactly this type and not a type which fulfills SALDataErrorReporting.
// This is useful for switch cases.
type SALDataErrorReportingExactly interface {
	SALDataErrorReporting
	isSALDataErrorReporting() bool
}

// _SALDataErrorReporting is the data-structure of this message
type _SALDataErrorReporting struct {
	*_SALData
	ErrorReportingData ErrorReportingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataErrorReporting) GetApplicationId() ApplicationId {
	return ApplicationId_ERROR_REPORTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataErrorReporting) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataErrorReporting) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataErrorReporting) GetErrorReportingData() ErrorReportingData {
	return m.ErrorReportingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataErrorReporting factory function for _SALDataErrorReporting
func NewSALDataErrorReporting(errorReportingData ErrorReportingData, salData SALData) *_SALDataErrorReporting {
	_result := &_SALDataErrorReporting{
		ErrorReportingData: errorReportingData,
		_SALData:           NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataErrorReporting(structType any) SALDataErrorReporting {
	if casted, ok := structType.(SALDataErrorReporting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataErrorReporting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataErrorReporting) GetTypeName() string {
	return "SALDataErrorReporting"
}

func (m *_SALDataErrorReporting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorReportingData)
	lengthInBits += m.ErrorReportingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataErrorReporting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataErrorReportingParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataErrorReporting, error) {
	return SALDataErrorReportingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataErrorReportingParseWithBufferProducer(applicationId ApplicationId) func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataErrorReporting, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataErrorReporting, error) {
		return SALDataErrorReportingParseWithBuffer(ctx, readBuffer, applicationId)
	}
}

func SALDataErrorReportingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataErrorReporting, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataErrorReporting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataErrorReporting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorReportingData, err := ReadSimpleField[ErrorReportingData](ctx, "errorReportingData", ReadComplex[ErrorReportingData](ErrorReportingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorReportingData' field"))
	}

	if closeErr := readBuffer.CloseContext("SALDataErrorReporting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataErrorReporting")
	}

	// Create a partially initialized instance
	_child := &_SALDataErrorReporting{
		_SALData:           &_SALData{},
		ErrorReportingData: errorReportingData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataErrorReporting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataErrorReporting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataErrorReporting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataErrorReporting")
		}

		// Simple Field (errorReportingData)
		if pushErr := writeBuffer.PushContext("errorReportingData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorReportingData")
		}
		_errorReportingDataErr := writeBuffer.WriteSerializable(ctx, m.GetErrorReportingData())
		if popErr := writeBuffer.PopContext("errorReportingData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorReportingData")
		}
		if _errorReportingDataErr != nil {
			return errors.Wrap(_errorReportingDataErr, "Error serializing 'errorReportingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataErrorReporting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataErrorReporting")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataErrorReporting) isSALDataErrorReporting() bool {
	return true
}

func (m *_SALDataErrorReporting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
