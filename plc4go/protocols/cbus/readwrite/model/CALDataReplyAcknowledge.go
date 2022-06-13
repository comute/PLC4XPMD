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

// CALDataReplyAcknowledge is the data-structure of this message
type CALDataReplyAcknowledge struct {
	*CALData
	ParamNo uint8
	Code    uint8
}

// ICALDataReplyAcknowledge is the corresponding interface of CALDataReplyAcknowledge
type ICALDataReplyAcknowledge interface {
	ICALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() uint8
	// GetCode returns Code (property field)
	GetCode() uint8
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

func (m *CALDataReplyAcknowledge) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

func (m *CALDataReplyAcknowledge) GetParent() *CALData {
	return m.CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALDataReplyAcknowledge) GetParamNo() uint8 {
	return m.ParamNo
}

func (m *CALDataReplyAcknowledge) GetCode() uint8 {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataReplyAcknowledge factory function for CALDataReplyAcknowledge
func NewCALDataReplyAcknowledge(paramNo uint8, code uint8, commandTypeContainer CALCommandTypeContainer) *CALDataReplyAcknowledge {
	_result := &CALDataReplyAcknowledge{
		ParamNo: paramNo,
		Code:    code,
		CALData: NewCALData(commandTypeContainer),
	}
	_result.Child = _result
	return _result
}

func CastCALDataReplyAcknowledge(structType interface{}) *CALDataReplyAcknowledge {
	if casted, ok := structType.(CALDataReplyAcknowledge); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataReplyAcknowledge); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataReplyAcknowledge(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataReplyAcknowledge(casted.Child)
	}
	return nil
}

func (m *CALDataReplyAcknowledge) GetTypeName() string {
	return "CALDataReplyAcknowledge"
}

func (m *CALDataReplyAcknowledge) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataReplyAcknowledge) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALDataReplyAcknowledge) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyAcknowledgeParse(readBuffer utils.ReadBuffer) (*CALDataReplyAcknowledge, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReplyAcknowledge"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReplyAcknowledge")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	_paramNo, _paramNoErr := readBuffer.ReadUint8("paramNo", 8)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field")
	}
	paramNo := _paramNo

	// Simple Field (code)
	_code, _codeErr := readBuffer.ReadUint8("code", 8)
	if _codeErr != nil {
		return nil, errors.Wrap(_codeErr, "Error parsing 'code' field")
	}
	code := _code

	if closeErr := readBuffer.CloseContext("CALDataReplyAcknowledge"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReplyAcknowledge")
	}

	// Create a partially initialized instance
	_child := &CALDataReplyAcknowledge{
		ParamNo: paramNo,
		Code:    code,
		CALData: &CALData{},
	}
	_child.CALData.Child = _child
	return _child, nil
}

func (m *CALDataReplyAcknowledge) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReplyAcknowledge"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReplyAcknowledge")
		}

		// Simple Field (paramNo)
		paramNo := uint8(m.ParamNo)
		_paramNoErr := writeBuffer.WriteUint8("paramNo", 8, (paramNo))
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (code)
		code := uint8(m.Code)
		_codeErr := writeBuffer.WriteUint8("code", 8, (code))
		if _codeErr != nil {
			return errors.Wrap(_codeErr, "Error serializing 'code' field")
		}

		if popErr := writeBuffer.PopContext("CALDataReplyAcknowledge"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReplyAcknowledge")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataReplyAcknowledge) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
