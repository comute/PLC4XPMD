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

// SALDataPoolsSpasPondsFountainsControl is the corresponding interface of SALDataPoolsSpasPondsFountainsControl
type SALDataPoolsSpasPondsFountainsControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetPoolsSpaPondsFountainsData returns PoolsSpaPondsFountainsData (property field)
	GetPoolsSpaPondsFountainsData() LightingData
}

// SALDataPoolsSpasPondsFountainsControlExactly can be used when we want exactly this type and not a type which fulfills SALDataPoolsSpasPondsFountainsControl.
// This is useful for switch cases.
type SALDataPoolsSpasPondsFountainsControlExactly interface {
	SALDataPoolsSpasPondsFountainsControl
	isSALDataPoolsSpasPondsFountainsControl() bool
}

// _SALDataPoolsSpasPondsFountainsControl is the data-structure of this message
type _SALDataPoolsSpasPondsFountainsControl struct {
	*_SALData
	PoolsSpaPondsFountainsData LightingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) GetApplicationId() ApplicationId {
	return ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataPoolsSpasPondsFountainsControl) GetPoolsSpaPondsFountainsData() LightingData {
	return m.PoolsSpaPondsFountainsData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataPoolsSpasPondsFountainsControl factory function for _SALDataPoolsSpasPondsFountainsControl
func NewSALDataPoolsSpasPondsFountainsControl(poolsSpaPondsFountainsData LightingData, salData SALData) *_SALDataPoolsSpasPondsFountainsControl {
	_result := &_SALDataPoolsSpasPondsFountainsControl{
		PoolsSpaPondsFountainsData: poolsSpaPondsFountainsData,
		_SALData:                   NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataPoolsSpasPondsFountainsControl(structType any) SALDataPoolsSpasPondsFountainsControl {
	if casted, ok := structType.(SALDataPoolsSpasPondsFountainsControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataPoolsSpasPondsFountainsControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetTypeName() string {
	return "SALDataPoolsSpasPondsFountainsControl"
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (poolsSpaPondsFountainsData)
	lengthInBits += m.PoolsSpaPondsFountainsData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataPoolsSpasPondsFountainsControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataPoolsSpasPondsFountainsControlParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataPoolsSpasPondsFountainsControl, error) {
	return SALDataPoolsSpasPondsFountainsControlParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataPoolsSpasPondsFountainsControlParseWithBufferProducer(applicationId ApplicationId) func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataPoolsSpasPondsFountainsControl, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SALDataPoolsSpasPondsFountainsControl, error) {
		return SALDataPoolsSpasPondsFountainsControlParseWithBuffer(ctx, readBuffer, applicationId)
	}
}

func SALDataPoolsSpasPondsFountainsControlParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataPoolsSpasPondsFountainsControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataPoolsSpasPondsFountainsControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataPoolsSpasPondsFountainsControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	poolsSpaPondsFountainsData, err := ReadSimpleField[LightingData](ctx, "poolsSpaPondsFountainsData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'poolsSpaPondsFountainsData' field"))
	}

	if closeErr := readBuffer.CloseContext("SALDataPoolsSpasPondsFountainsControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataPoolsSpasPondsFountainsControl")
	}

	// Create a partially initialized instance
	_child := &_SALDataPoolsSpasPondsFountainsControl{
		_SALData:                   &_SALData{},
		PoolsSpaPondsFountainsData: poolsSpaPondsFountainsData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataPoolsSpasPondsFountainsControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataPoolsSpasPondsFountainsControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataPoolsSpasPondsFountainsControl")
		}

		// Simple Field (poolsSpaPondsFountainsData)
		if pushErr := writeBuffer.PushContext("poolsSpaPondsFountainsData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for poolsSpaPondsFountainsData")
		}
		_poolsSpaPondsFountainsDataErr := writeBuffer.WriteSerializable(ctx, m.GetPoolsSpaPondsFountainsData())
		if popErr := writeBuffer.PopContext("poolsSpaPondsFountainsData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for poolsSpaPondsFountainsData")
		}
		if _poolsSpaPondsFountainsDataErr != nil {
			return errors.Wrap(_poolsSpaPondsFountainsDataErr, "Error serializing 'poolsSpaPondsFountainsData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataPoolsSpasPondsFountainsControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataPoolsSpasPondsFountainsControl")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataPoolsSpasPondsFountainsControl) isSALDataPoolsSpasPondsFountainsControl() bool {
	return true
}

func (m *_SALDataPoolsSpasPondsFountainsControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
