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

// ErrorReportingSystemCategoryTypeBuildingManagementSystems is the corresponding interface of ErrorReportingSystemCategoryTypeBuildingManagementSystems
type ErrorReportingSystemCategoryTypeBuildingManagementSystems interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForBuildingManagementSystems
	// IsErrorReportingSystemCategoryTypeBuildingManagementSystems is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryTypeBuildingManagementSystems()
}

// _ErrorReportingSystemCategoryTypeBuildingManagementSystems is the data-structure of this message
type _ErrorReportingSystemCategoryTypeBuildingManagementSystems struct {
	ErrorReportingSystemCategoryTypeContract
	CategoryForType ErrorReportingSystemCategoryTypeForBuildingManagementSystems
}

var _ ErrorReportingSystemCategoryTypeBuildingManagementSystems = (*_ErrorReportingSystemCategoryTypeBuildingManagementSystems)(nil)
var _ ErrorReportingSystemCategoryTypeRequirements = (*_ErrorReportingSystemCategoryTypeBuildingManagementSystems)(nil)

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

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
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
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		CategoryForType:                          categoryForType,
	}
	_result.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = _result
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
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeBuildingManagementSystems ErrorReportingSystemCategoryTypeBuildingManagementSystems, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryForType, err := ReadEnumField[ErrorReportingSystemCategoryTypeForBuildingManagementSystems](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForBuildingManagementSystems", ReadEnum(ErrorReportingSystemCategoryTypeForBuildingManagementSystemsByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryForType' field"))
	}
	m.CategoryForType = categoryForType

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
	}

	return m, nil
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
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
		}

		if err := WriteSimpleEnumField[ErrorReportingSystemCategoryTypeForBuildingManagementSystems](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForBuildingManagementSystems", m.GetCategoryForType(), WriteEnum[ErrorReportingSystemCategoryTypeForBuildingManagementSystems, uint8](ErrorReportingSystemCategoryTypeForBuildingManagementSystems.GetValue, ErrorReportingSystemCategoryTypeForBuildingManagementSystems.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeBuildingManagementSystems"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeBuildingManagementSystems")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeBuildingManagementSystems) IsErrorReportingSystemCategoryTypeBuildingManagementSystems() {
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
