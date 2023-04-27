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

// ErrorReportingSystemCategoryTypeBuildingManagementSystems is the corresponding interface of ErrorReportingSystemCategoryTypeBuildingManagementSystems
type ErrorReportingSystemCategoryTypeBuildingManagementSystems interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForBuildingManagementSystems
}

// ErrorReportingSystemCategoryTypeBuildingManagementSystemsExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryTypeBuildingManagementSystems.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeBuildingManagementSystemsExactly interface {
	ErrorReportingSystemCategoryTypeBuildingManagementSystems
	isErrorReportingSystemCategoryTypeBuildingManagementSystems() bool
}

// _ErrorReportingSystemCategoryTypeBuildingManagementSystems is the data-structure of this message
type _ErrorReportingSystemCategoryTypeBuildingManagementSystems struct {
	*_ErrorReportingSystemCategoryType
	CategoryForType ErrorReportingSystemCategoryTypeForBuildingManagementSystems
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) InitializeParent(parent ErrorReportingSystemCategoryType) {
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetParent() ErrorReportingSystemCategoryType {
	return m._ErrorReportingSystemCategoryType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetCategoryForType() ErrorReportingSystemCategoryTypeForBuildingManagementSystems {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategoryTypeBuildingManagementSystems factory function for _ErrorReportingSystemCategoryTypeBuildingManagementSystems
func NewErrorReportingSystemCategoryTypeBuildingManagementSystems(categoryForType ErrorReportingSystemCategoryTypeForBuildingManagementSystems) *_ErrorReportingSystemCategoryTypeBuildingManagementSystems {
	_result := &_ErrorReportingSystemCategoryTypeBuildingManagementSystems{
		CategoryForType:                   categoryForType,
		_ErrorReportingSystemCategoryType: NewErrorReportingSystemCategoryType(),
	}
	_result._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeBuildingManagementSystems(structType any) ErrorReportingSystemCategoryTypeBuildingManagementSystems {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeBuildingManagementSystems); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeBuildingManagementSystems); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeBuildingManagementSystems"
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeBuildingManagementSystemsParse(theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeBuildingManagementSystems, error) {
	return ErrorReportingSystemCategoryTypeBuildingManagementSystemsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeBuildingManagementSystemsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeBuildingManagementSystems, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (categoryForType)
	if pullErr := readBuffer.PullContext("categoryForType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for categoryForType")
	}
	_categoryForType, _categoryForTypeErr := ErrorReportingSystemCategoryTypeForBuildingManagementSystemsParseWithBuffer(ctx, readBuffer)
	if _categoryForTypeErr != nil {
		return nil, errors.Wrap(_categoryForTypeErr, "Error parsing 'categoryForType' field of ErrorReportingSystemCategoryTypeBuildingManagementSystems")
	}
	categoryForType := _categoryForType
	if closeErr := readBuffer.CloseContext("categoryForType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for categoryForType")
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
	}

	// Create a partially initialized instance
	_child := &_ErrorReportingSystemCategoryTypeBuildingManagementSystems{
		_ErrorReportingSystemCategoryType: &_ErrorReportingSystemCategoryType{},
		CategoryForType:                   categoryForType,
	}
	_child._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _child
	return _child, nil
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
		}

		// Simple Field (categoryForType)
		if pushErr := writeBuffer.PushContext("categoryForType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for categoryForType")
		}
		_categoryForTypeErr := writeBuffer.WriteSerializable(ctx, m.GetCategoryForType())
		if popErr := writeBuffer.PopContext("categoryForType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for categoryForType")
		}
		if _categoryForTypeErr != nil {
			return errors.Wrap(_categoryForTypeErr, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) isErrorReportingSystemCategoryTypeBuildingManagementSystems() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
