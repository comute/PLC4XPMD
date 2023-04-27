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
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataMeasurement is the corresponding interface of SALDataMeasurement
type SALDataMeasurement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetMeasurementData returns MeasurementData (property field)
	GetMeasurementData() MeasurementData
}

// SALDataMeasurementExactly can be used when we want exactly this type and not a type which fulfills SALDataMeasurement.
// This is useful for switch cases.
type SALDataMeasurementExactly interface {
	SALDataMeasurement
	isSALDataMeasurement() bool
}

// _SALDataMeasurement is the data-structure of this message
type _SALDataMeasurement struct {
	*_SALData
	MeasurementData MeasurementData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataMeasurement) GetApplicationId() ApplicationId {
	return ApplicationId_MEASUREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataMeasurement) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataMeasurement) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataMeasurement) GetMeasurementData() MeasurementData {
	return m.MeasurementData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataMeasurement factory function for _SALDataMeasurement
func NewSALDataMeasurement(measurementData MeasurementData, salData SALData) *_SALDataMeasurement {
	_result := &_SALDataMeasurement{
		MeasurementData: measurementData,
		_SALData:        NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataMeasurement(structType any) SALDataMeasurement {
	if casted, ok := structType.(SALDataMeasurement); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataMeasurement); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataMeasurement) GetTypeName() string {
	return "SALDataMeasurement"
}

func (m *_SALDataMeasurement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (measurementData)
	lengthInBits += m.MeasurementData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataMeasurement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataMeasurementParse(theBytes []byte, applicationId ApplicationId) (SALDataMeasurement, error) {
	return SALDataMeasurementParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataMeasurementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataMeasurement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataMeasurement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataMeasurement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (measurementData)
	if pullErr := readBuffer.PullContext("measurementData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for measurementData")
	}
	_measurementData, _measurementDataErr := MeasurementDataParseWithBuffer(ctx, readBuffer)
	if _measurementDataErr != nil {
		return nil, errors.Wrap(_measurementDataErr, "Error parsing 'measurementData' field of SALDataMeasurement")
	}
	measurementData := _measurementData.(MeasurementData)
	if closeErr := readBuffer.CloseContext("measurementData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for measurementData")
	}

	if closeErr := readBuffer.CloseContext("SALDataMeasurement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataMeasurement")
	}

	// Create a partially initialized instance
	_child := &_SALDataMeasurement{
		_SALData:        &_SALData{},
		MeasurementData: measurementData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataMeasurement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataMeasurement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataMeasurement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataMeasurement")
		}

		// Simple Field (measurementData)
		if pushErr := writeBuffer.PushContext("measurementData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for measurementData")
		}
		_measurementDataErr := writeBuffer.WriteSerializable(ctx, m.GetMeasurementData())
		if popErr := writeBuffer.PopContext("measurementData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for measurementData")
		}
		if _measurementDataErr != nil {
			return errors.Wrap(_measurementDataErr, "Error serializing 'measurementData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataMeasurement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataMeasurement")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataMeasurement) isSALDataMeasurement() bool {
	return true
}

func (m *_SALDataMeasurement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
