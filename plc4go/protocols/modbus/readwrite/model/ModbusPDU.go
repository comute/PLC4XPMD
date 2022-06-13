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

// ModbusPDU is the data-structure of this message
type ModbusPDU struct {
	Child IModbusPDUChild
}

// IModbusPDU is the corresponding interface of ModbusPDU
type IModbusPDU interface {
	// GetErrorFlag returns ErrorFlag (discriminator field)
	GetErrorFlag() bool
	// GetFunctionFlag returns FunctionFlag (discriminator field)
	GetFunctionFlag() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IModbusPDUParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IModbusPDU, serializeChildFunction func() error) error
	GetTypeName() string
}

type IModbusPDUChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ModbusPDU)
	GetParent() *ModbusPDU

	GetTypeName() string
	IModbusPDU
}

// NewModbusPDU factory function for ModbusPDU
func NewModbusPDU() *ModbusPDU {
	return &ModbusPDU{}
}

func CastModbusPDU(structType interface{}) *ModbusPDU {
	if casted, ok := structType.(ModbusPDU); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return casted
	}
	if casted, ok := structType.(IModbusPDUChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *ModbusPDU) GetTypeName() string {
	return "ModbusPDU"
}

func (m *ModbusPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *ModbusPDU) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (errorFlag)
	lengthInBits += 1
	// Discriminator Field (functionFlag)
	lengthInBits += 7

	return lengthInBits
}

func (m *ModbusPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (errorFlag) (Used as input to a switch field)
	errorFlag, _errorFlagErr := readBuffer.ReadBit("errorFlag")
	if _errorFlagErr != nil {
		return nil, errors.Wrap(_errorFlagErr, "Error parsing 'errorFlag' field")
	}

	// Discriminator Field (functionFlag) (Used as input to a switch field)
	functionFlag, _functionFlagErr := readBuffer.ReadUint8("functionFlag", 7)
	if _functionFlagErr != nil {
		return nil, errors.Wrap(_functionFlagErr, "Error parsing 'functionFlag' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ModbusPDUChild interface {
		InitializeParent(*ModbusPDU)
		GetParent() *ModbusPDU
	}
	var _child ModbusPDUChild
	var typeSwitchError error
	switch {
	case errorFlag == bool(true): // ModbusPDUError
		_child, typeSwitchError = ModbusPDUErrorParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(false): // ModbusPDUReadDiscreteInputsRequest
		_child, typeSwitchError = ModbusPDUReadDiscreteInputsRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(true): // ModbusPDUReadDiscreteInputsResponse
		_child, typeSwitchError = ModbusPDUReadDiscreteInputsResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(false): // ModbusPDUReadCoilsRequest
		_child, typeSwitchError = ModbusPDUReadCoilsRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(true): // ModbusPDUReadCoilsResponse
		_child, typeSwitchError = ModbusPDUReadCoilsResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(false): // ModbusPDUWriteSingleCoilRequest
		_child, typeSwitchError = ModbusPDUWriteSingleCoilRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(true): // ModbusPDUWriteSingleCoilResponse
		_child, typeSwitchError = ModbusPDUWriteSingleCoilResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(false): // ModbusPDUWriteMultipleCoilsRequest
		_child, typeSwitchError = ModbusPDUWriteMultipleCoilsRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(true): // ModbusPDUWriteMultipleCoilsResponse
		_child, typeSwitchError = ModbusPDUWriteMultipleCoilsResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(false): // ModbusPDUReadInputRegistersRequest
		_child, typeSwitchError = ModbusPDUReadInputRegistersRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(true): // ModbusPDUReadInputRegistersResponse
		_child, typeSwitchError = ModbusPDUReadInputRegistersResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(false): // ModbusPDUReadHoldingRegistersRequest
		_child, typeSwitchError = ModbusPDUReadHoldingRegistersRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(true): // ModbusPDUReadHoldingRegistersResponse
		_child, typeSwitchError = ModbusPDUReadHoldingRegistersResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(false): // ModbusPDUWriteSingleRegisterRequest
		_child, typeSwitchError = ModbusPDUWriteSingleRegisterRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(true): // ModbusPDUWriteSingleRegisterResponse
		_child, typeSwitchError = ModbusPDUWriteSingleRegisterResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(false): // ModbusPDUWriteMultipleHoldingRegistersRequest
		_child, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(true): // ModbusPDUWriteMultipleHoldingRegistersResponse
		_child, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(false): // ModbusPDUReadWriteMultipleHoldingRegistersRequest
		_child, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(true): // ModbusPDUReadWriteMultipleHoldingRegistersResponse
		_child, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(false): // ModbusPDUMaskWriteHoldingRegisterRequest
		_child, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(true): // ModbusPDUMaskWriteHoldingRegisterResponse
		_child, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(false): // ModbusPDUReadFifoQueueRequest
		_child, typeSwitchError = ModbusPDUReadFifoQueueRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(true): // ModbusPDUReadFifoQueueResponse
		_child, typeSwitchError = ModbusPDUReadFifoQueueResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(false): // ModbusPDUReadFileRecordRequest
		_child, typeSwitchError = ModbusPDUReadFileRecordRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(true): // ModbusPDUReadFileRecordResponse
		_child, typeSwitchError = ModbusPDUReadFileRecordResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(false): // ModbusPDUWriteFileRecordRequest
		_child, typeSwitchError = ModbusPDUWriteFileRecordRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(true): // ModbusPDUWriteFileRecordResponse
		_child, typeSwitchError = ModbusPDUWriteFileRecordResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(false): // ModbusPDUReadExceptionStatusRequest
		_child, typeSwitchError = ModbusPDUReadExceptionStatusRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(true): // ModbusPDUReadExceptionStatusResponse
		_child, typeSwitchError = ModbusPDUReadExceptionStatusResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(false): // ModbusPDUDiagnosticRequest
		_child, typeSwitchError = ModbusPDUDiagnosticRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(true): // ModbusPDUDiagnosticResponse
		_child, typeSwitchError = ModbusPDUDiagnosticResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(false): // ModbusPDUGetComEventCounterRequest
		_child, typeSwitchError = ModbusPDUGetComEventCounterRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(true): // ModbusPDUGetComEventCounterResponse
		_child, typeSwitchError = ModbusPDUGetComEventCounterResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(false): // ModbusPDUGetComEventLogRequest
		_child, typeSwitchError = ModbusPDUGetComEventLogRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(true): // ModbusPDUGetComEventLogResponse
		_child, typeSwitchError = ModbusPDUGetComEventLogResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(false): // ModbusPDUReportServerIdRequest
		_child, typeSwitchError = ModbusPDUReportServerIdRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(true): // ModbusPDUReportServerIdResponse
		_child, typeSwitchError = ModbusPDUReportServerIdResponseParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(false): // ModbusPDUReadDeviceIdentificationRequest
		_child, typeSwitchError = ModbusPDUReadDeviceIdentificationRequestParse(readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(true): // ModbusPDUReadDeviceIdentificationResponse
		_child, typeSwitchError = ModbusPDUReadDeviceIdentificationResponseParse(readBuffer, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDU")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *ModbusPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ModbusPDU) SerializeParent(writeBuffer utils.WriteBuffer, child IModbusPDU, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusPDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDU")
	}

	// Discriminator Field (errorFlag) (Used as input to a switch field)
	errorFlag := bool(child.GetErrorFlag())
	_errorFlagErr := writeBuffer.WriteBit("errorFlag", (errorFlag))

	if _errorFlagErr != nil {
		return errors.Wrap(_errorFlagErr, "Error serializing 'errorFlag' field")
	}

	// Discriminator Field (functionFlag) (Used as input to a switch field)
	functionFlag := uint8(child.GetFunctionFlag())
	_functionFlagErr := writeBuffer.WriteUint8("functionFlag", 7, (functionFlag))

	if _functionFlagErr != nil {
		return errors.Wrap(_functionFlagErr, "Error serializing 'functionFlag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDU")
	}
	return nil
}

func (m *ModbusPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
