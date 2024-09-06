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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ComObjectTable is the corresponding interface of ComObjectTable
type ComObjectTable interface {
	ComObjectTableContract
	ComObjectTableRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsComObjectTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsComObjectTable()
}

// ComObjectTableContract provides a set of functions which can be overwritten by a sub struct
type ComObjectTableContract interface {
	// IsComObjectTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsComObjectTable()
}

// ComObjectTableRequirements provides a set of functions which need to be implemented by a sub struct
type ComObjectTableRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetFirmwareType returns FirmwareType (discriminator field)
	GetFirmwareType() FirmwareType
}

// _ComObjectTable is the data-structure of this message
type _ComObjectTable struct {
	_SubType ComObjectTable
}

var _ ComObjectTableContract = (*_ComObjectTable)(nil)

// NewComObjectTable factory function for _ComObjectTable
func NewComObjectTable() *_ComObjectTable {
	return &_ComObjectTable{}
}

// Deprecated: use the interface for direct cast
func CastComObjectTable(structType any) ComObjectTable {
	if casted, ok := structType.(ComObjectTable); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTable); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTable) GetTypeName() string {
	return "ComObjectTable"
}

func (m *_ComObjectTable) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ComObjectTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ComObjectTableParse[T ComObjectTable](ctx context.Context, theBytes []byte, firmwareType FirmwareType) (T, error) {
	return ComObjectTableParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), firmwareType)
}

func ComObjectTableParseWithBufferProducer[T ComObjectTable](firmwareType FirmwareType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ComObjectTableParseWithBuffer[T](ctx, readBuffer, firmwareType)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ComObjectTableParseWithBuffer[T ComObjectTable](ctx context.Context, readBuffer utils.ReadBuffer, firmwareType FirmwareType) (T, error) {
	v, err := (&_ComObjectTable{}).parse(ctx, readBuffer, firmwareType)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ComObjectTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, firmwareType FirmwareType) (__comObjectTable ComObjectTable, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ComObjectTable
	switch {
	case firmwareType == FirmwareType_SYSTEM_1: // ComObjectTableRealisationType1
		if _child, err = (&_ComObjectTableRealisationType1{}).parse(ctx, readBuffer, m, firmwareType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ComObjectTableRealisationType1 for type-switch of ComObjectTable")
		}
	case firmwareType == FirmwareType_SYSTEM_2: // ComObjectTableRealisationType2
		if _child, err = (&_ComObjectTableRealisationType2{}).parse(ctx, readBuffer, m, firmwareType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ComObjectTableRealisationType2 for type-switch of ComObjectTable")
		}
	case firmwareType == FirmwareType_SYSTEM_300: // ComObjectTableRealisationType6
		if _child, err = (&_ComObjectTableRealisationType6{}).parse(ctx, readBuffer, m, firmwareType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ComObjectTableRealisationType6 for type-switch of ComObjectTable")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [firmwareType=%v]", firmwareType)
	}

	if closeErr := readBuffer.CloseContext("ComObjectTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTable")
	}

	return _child, nil
}

func (pm *_ComObjectTable) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ComObjectTable, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ComObjectTable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ComObjectTable")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ComObjectTable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ComObjectTable")
	}
	return nil
}

func (m *_ComObjectTable) IsComObjectTable() {}
