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

// ErrorReportingSystemCategoryType is the corresponding interface of ErrorReportingSystemCategoryType
type ErrorReportingSystemCategoryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetErrorReportingSystemCategoryClass returns ErrorReportingSystemCategoryClass (discriminator field)
	GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass
}

// ErrorReportingSystemCategoryTypeExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryType.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeExactly interface {
	ErrorReportingSystemCategoryType
	isErrorReportingSystemCategoryType() bool
}

// _ErrorReportingSystemCategoryType is the data-structure of this message
type _ErrorReportingSystemCategoryType struct {
	_ErrorReportingSystemCategoryTypeChildRequirements
}

type _ErrorReportingSystemCategoryTypeChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass
}

type ErrorReportingSystemCategoryTypeParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingSystemCategoryType, serializeChildFunction func() error) error
	GetTypeName() string
}

type ErrorReportingSystemCategoryTypeChild interface {
	utils.Serializable
	InitializeParent(parent ErrorReportingSystemCategoryType)
	GetParent() *ErrorReportingSystemCategoryType

	GetTypeName() string
	ErrorReportingSystemCategoryType
}

// NewErrorReportingSystemCategoryType factory function for _ErrorReportingSystemCategoryType
func NewErrorReportingSystemCategoryType() *_ErrorReportingSystemCategoryType {
	return &_ErrorReportingSystemCategoryType{}
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryType(structType any) ErrorReportingSystemCategoryType {
	if casted, ok := structType.(ErrorReportingSystemCategoryType); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryType); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryType) GetTypeName() string {
	return "ErrorReportingSystemCategoryType"
}

func (m *_ErrorReportingSystemCategoryType) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeParse(ctx context.Context, theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryType, error) {
	return ErrorReportingSystemCategoryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeParseWithBufferProducer[T ErrorReportingSystemCategoryType](errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := ErrorReportingSystemCategoryTypeParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func ErrorReportingSystemCategoryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ErrorReportingSystemCategoryTypeChildSerializeRequirement interface {
		ErrorReportingSystemCategoryType
		InitializeParent(ErrorReportingSystemCategoryType)
		GetParent() ErrorReportingSystemCategoryType
	}
	var _childTemp any
	var _child ErrorReportingSystemCategoryTypeChildSerializeRequirement
	var typeSwitchError error
	switch {
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_INPUT_UNITS: // ErrorReportingSystemCategoryTypeInputUnits
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeInputUnitsParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_SUPPORT_UNITS: // ErrorReportingSystemCategoryTypeSupportUnits
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeSupportUnitsParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS: // ErrorReportingSystemCategoryTypeBuildingManagementSystems
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeBuildingManagementSystemsParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_OUTPUT_UNITS: // ErrorReportingSystemCategoryTypeOutputUnits
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeOutputUnitsParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS: // ErrorReportingSystemCategoryTypeClimateControllers
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeClimateControllersParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	case 0 == 0: // ErrorReportingSystemCategoryTypeReserved
		_childTemp, typeSwitchError = ErrorReportingSystemCategoryTypeReservedParseWithBuffer(ctx, readBuffer, errorReportingSystemCategoryClass)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [errorReportingSystemCategoryClass=%v]", errorReportingSystemCategoryClass)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ErrorReportingSystemCategoryType")
	}
	_child = _childTemp.(ErrorReportingSystemCategoryTypeChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryType")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ErrorReportingSystemCategoryType) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingSystemCategoryType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryType")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryType")
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryType) isErrorReportingSystemCategoryType() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
