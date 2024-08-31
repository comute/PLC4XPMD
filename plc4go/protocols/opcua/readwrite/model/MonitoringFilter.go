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

// MonitoringFilter is the corresponding interface of MonitoringFilter
type MonitoringFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// MonitoringFilterExactly can be used when we want exactly this type and not a type which fulfills MonitoringFilter.
// This is useful for switch cases.
type MonitoringFilterExactly interface {
	MonitoringFilter
	isMonitoringFilter() bool
}

// _MonitoringFilter is the data-structure of this message
type _MonitoringFilter struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoringFilter) GetIdentifier() string {
	return "721"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoringFilter) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoringFilter) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewMonitoringFilter factory function for _MonitoringFilter
func NewMonitoringFilter() *_MonitoringFilter {
	_result := &_MonitoringFilter{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoringFilter(structType any) MonitoringFilter {
	if casted, ok := structType.(MonitoringFilter); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoringFilter); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoringFilter) GetTypeName() string {
	return "MonitoringFilter"
}

func (m *_MonitoringFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MonitoringFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoringFilterParse(ctx context.Context, theBytes []byte, identifier string) (MonitoringFilter, error) {
	return MonitoringFilterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoringFilterParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoringFilter, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoringFilter, error) {
		return MonitoringFilterParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func MonitoringFilterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoringFilter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoringFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoringFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MonitoringFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoringFilter")
	}

	// Create a partially initialized instance
	_child := &_MonitoringFilter{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoringFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoringFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoringFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoringFilter")
		}

		if popErr := writeBuffer.PopContext("MonitoringFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoringFilter")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoringFilter) isMonitoringFilter() bool {
	return true
}

func (m *_MonitoringFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
