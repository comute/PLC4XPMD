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

// FilterOperand is the corresponding interface of FilterOperand
type FilterOperand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// IsFilterOperand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFilterOperand()
}

// _FilterOperand is the data-structure of this message
type _FilterOperand struct {
	ExtensionObjectDefinitionContract
}

var _ FilterOperand = (*_FilterOperand)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FilterOperand)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FilterOperand) GetIdentifier() string {
	return "591"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FilterOperand) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// NewFilterOperand factory function for _FilterOperand
func NewFilterOperand() *_FilterOperand {
	_result := &_FilterOperand{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFilterOperand(structType any) FilterOperand {
	if casted, ok := structType.(FilterOperand); ok {
		return casted
	}
	if casted, ok := structType.(*FilterOperand); ok {
		return *casted
	}
	return nil
}

func (m *_FilterOperand) GetTypeName() string {
	return "FilterOperand"
}

func (m *_FilterOperand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_FilterOperand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FilterOperand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__filterOperand FilterOperand, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FilterOperand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FilterOperand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("FilterOperand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FilterOperand")
	}

	return m, nil
}

func (m *_FilterOperand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FilterOperand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FilterOperand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FilterOperand")
		}

		if popErr := writeBuffer.PopContext("FilterOperand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FilterOperand")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FilterOperand) IsFilterOperand() {}

func (m *_FilterOperand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
