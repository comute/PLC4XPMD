/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataRequestGetStatus is the data-structure of this message
type CALDataRequestGetStatus struct {
	*CALData
	ParamNo uint8
	Count   uint8
}

// ICALDataRequestGetStatus is the corresponding interface of CALDataRequestGetStatus
type ICALDataRequestGetStatus interface {
	ICALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CALDataRequestGetStatus) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

func (m *CALDataRequestGetStatus) GetParent() *CALData {
	return m.CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALDataRequestGetStatus) GetParamNo() uint8 {
	return m.ParamNo
}

func (m *CALDataRequestGetStatus) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataRequestGetStatus factory function for CALDataRequestGetStatus
func NewCALDataRequestGetStatus(paramNo uint8, count uint8, commandTypeContainer CALCommandTypeContainer) *CALDataRequestGetStatus {
	_result := &CALDataRequestGetStatus{
		ParamNo: paramNo,
		Count:   count,
		CALData: NewCALData(commandTypeContainer),
	}
	_result.Child = _result
	return _result
}

func CastCALDataRequestGetStatus(structType interface{}) *CALDataRequestGetStatus {
	if casted, ok := structType.(CALDataRequestGetStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataRequestGetStatus); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataRequestGetStatus(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataRequestGetStatus(casted.Child)
	}
	return nil
}

func (m *CALDataRequestGetStatus) GetTypeName() string {
	return "CALDataRequestGetStatus"
}

func (m *CALDataRequestGetStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataRequestGetStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALDataRequestGetStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRequestGetStatusParse(readBuffer utils.ReadBuffer) (*CALDataRequestGetStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRequestGetStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataRequestGetStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	_paramNo, _paramNoErr := readBuffer.ReadUint8("paramNo", 8)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field")
	}
	paramNo := _paramNo

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("CALDataRequestGetStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataRequestGetStatus")
	}

	// Create a partially initialized instance
	_child := &CALDataRequestGetStatus{
		ParamNo: paramNo,
		Count:   count,
		CALData: &CALData{},
	}
	_child.CALData.Child = _child
	return _child, nil
}

func (m *CALDataRequestGetStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRequestGetStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataRequestGetStatus")
		}

		// Simple Field (paramNo)
		paramNo := uint8(m.ParamNo)
		_paramNoErr := writeBuffer.WriteUint8("paramNo", 8, (paramNo))
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (count)
		count := uint8(m.Count)
		_countErr := writeBuffer.WriteUint8("count", 8, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRequestGetStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataRequestGetStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataRequestGetStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
