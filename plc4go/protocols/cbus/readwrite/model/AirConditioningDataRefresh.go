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

// AirConditioningDataRefresh is the corresponding interface of AirConditioningDataRefresh
type AirConditioningDataRefresh interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
}

// AirConditioningDataRefreshExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataRefresh.
// This is useful for switch cases.
type AirConditioningDataRefreshExactly interface {
	AirConditioningDataRefresh
	isAirConditioningDataRefresh() bool
}

// _AirConditioningDataRefresh is the data-structure of this message
type _AirConditioningDataRefresh struct {
	*_AirConditioningData
	ZoneGroup byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataRefresh) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataRefresh) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataRefresh) GetZoneGroup() byte {
	return m.ZoneGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataRefresh factory function for _AirConditioningDataRefresh
func NewAirConditioningDataRefresh(zoneGroup byte, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataRefresh {
	_result := &_AirConditioningDataRefresh{
		ZoneGroup:            zoneGroup,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataRefresh(structType any) AirConditioningDataRefresh {
	if casted, ok := structType.(AirConditioningDataRefresh); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataRefresh); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataRefresh) GetTypeName() string {
	return "AirConditioningDataRefresh"
}

func (m *_AirConditioningDataRefresh) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataRefresh) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataRefreshParse(ctx context.Context, theBytes []byte) (AirConditioningDataRefresh, error) {
	return AirConditioningDataRefreshParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataRefreshParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataRefresh, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataRefresh, error) {
		return AirConditioningDataRefreshParseWithBuffer(ctx, readBuffer)
	}
}

func AirConditioningDataRefreshParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataRefresh, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataRefresh"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataRefresh")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataRefresh"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataRefresh")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataRefresh{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataRefresh) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataRefresh) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataRefresh"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataRefresh")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := /*TODO: migrate me*/ writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataRefresh"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataRefresh")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataRefresh) isAirConditioningDataRefresh() bool {
	return true
}

func (m *_AirConditioningDataRefresh) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
