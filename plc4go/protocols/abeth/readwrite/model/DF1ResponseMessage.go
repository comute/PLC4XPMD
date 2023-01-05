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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1ResponseMessage is the corresponding interface of DF1ResponseMessage
type DF1ResponseMessage interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
}

// DF1ResponseMessageExactly can be used when we want exactly this type and not a type which fulfills DF1ResponseMessage.
// This is useful for switch cases.
type DF1ResponseMessageExactly interface {
	DF1ResponseMessage
	isDF1ResponseMessage() bool
}

// _DF1ResponseMessage is the data-structure of this message
type _DF1ResponseMessage struct {
	_DF1ResponseMessageChildRequirements
	DestinationAddress uint8
	SourceAddress      uint8
	Status             uint8
	TransactionCounter uint16

	// Arguments.
	PayloadLength uint16
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

type _DF1ResponseMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetCommandCode() uint8
}

type DF1ResponseMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child DF1ResponseMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type DF1ResponseMessageChild interface {
	utils.Serializable
	InitializeParent(parent DF1ResponseMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16)
	GetParent() *DF1ResponseMessage

	GetTypeName() string
	DF1ResponseMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1ResponseMessage) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1ResponseMessage) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1ResponseMessage) GetStatus() uint8 {
	return m.Status
}

func (m *_DF1ResponseMessage) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1ResponseMessage factory function for _DF1ResponseMessage
func NewDF1ResponseMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, payloadLength uint16) *_DF1ResponseMessage {
	return &_DF1ResponseMessage{DestinationAddress: destinationAddress, SourceAddress: sourceAddress, Status: status, TransactionCounter: transactionCounter, PayloadLength: payloadLength}
}

// Deprecated: use the interface for direct cast
func CastDF1ResponseMessage(structType interface{}) DF1ResponseMessage {
	if casted, ok := structType.(DF1ResponseMessage); ok {
		return casted
	}
	if casted, ok := structType.(*DF1ResponseMessage); ok {
		return *casted
	}
	return nil
}

func (m *_DF1ResponseMessage) GetTypeName() string {
	return "DF1ResponseMessage"
}

func (m *_DF1ResponseMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1ResponseMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1ResponseMessageParse(theBytes []byte, payloadLength uint16) (DF1ResponseMessage, error) {
	return DF1ResponseMessageParseWithBuffer(utils.NewReadBufferByteBased(theBytes), payloadLength)
}

func DF1ResponseMessageParseWithBuffer(readBuffer utils.ReadBuffer, payloadLength uint16) (DF1ResponseMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1ResponseMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1ResponseMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DF1ResponseMessage")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (destinationAddress)
	_destinationAddress, _destinationAddressErr := readBuffer.ReadUint8("destinationAddress", 8)
	if _destinationAddressErr != nil {
		return nil, errors.Wrap(_destinationAddressErr, "Error parsing 'destinationAddress' field of DF1ResponseMessage")
	}
	destinationAddress := _destinationAddress

	// Simple Field (sourceAddress)
	_sourceAddress, _sourceAddressErr := readBuffer.ReadUint8("sourceAddress", 8)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field of DF1ResponseMessage")
	}
	sourceAddress := _sourceAddress

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DF1ResponseMessage")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := readBuffer.ReadUint8("commandCode", 8)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field of DF1ResponseMessage")
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of DF1ResponseMessage")
	}
	status := _status

	// Simple Field (transactionCounter)
	_transactionCounter, _transactionCounterErr := readBuffer.ReadUint16("transactionCounter", 16)
	if _transactionCounterErr != nil {
		return nil, errors.Wrap(_transactionCounterErr, "Error parsing 'transactionCounter' field of DF1ResponseMessage")
	}
	transactionCounter := _transactionCounter

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type DF1ResponseMessageChildSerializeRequirement interface {
		DF1ResponseMessage
		InitializeParent(DF1ResponseMessage, uint8, uint8, uint8, uint16)
		GetParent() DF1ResponseMessage
	}
	var _childTemp interface{}
	var _child DF1ResponseMessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandCode == 0x4F: // DF1CommandResponseMessageProtectedTypedLogicalRead
		_childTemp, typeSwitchError = DF1CommandResponseMessageProtectedTypedLogicalReadParseWithBuffer(readBuffer, payloadLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of DF1ResponseMessage")
	}
	_child = _childTemp.(DF1ResponseMessageChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("DF1ResponseMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1ResponseMessage")
	}

	// Finish initializing
	_child.InitializeParent(_child, destinationAddress, sourceAddress, status, transactionCounter)
	_child.GetParent().(*_DF1ResponseMessage).reservedField0 = reservedField0
	_child.GetParent().(*_DF1ResponseMessage).reservedField1 = reservedField1
	return _child, nil
}

func (pm *_DF1ResponseMessage) SerializeParent(writeBuffer utils.WriteBuffer, child DF1ResponseMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DF1ResponseMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1ResponseMessage")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if pm.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *pm.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 8, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (destinationAddress)
	destinationAddress := uint8(m.GetDestinationAddress())
	_destinationAddressErr := writeBuffer.WriteUint8("destinationAddress", 8, (destinationAddress))
	if _destinationAddressErr != nil {
		return errors.Wrap(_destinationAddressErr, "Error serializing 'destinationAddress' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress := uint8(m.GetSourceAddress())
	_sourceAddressErr := writeBuffer.WriteUint8("sourceAddress", 8, (sourceAddress))
	if _sourceAddressErr != nil {
		return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if pm.reservedField1 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *pm.reservedField1
		}
		_err := writeBuffer.WriteUint8("reserved", 8, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.GetCommandCode())
	_commandCodeErr := writeBuffer.WriteUint8("commandCode", 8, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Simple Field (status)
	status := uint8(m.GetStatus())
	_statusErr := writeBuffer.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter := uint16(m.GetTransactionCounter())
	_transactionCounterErr := writeBuffer.WriteUint16("transactionCounter", 16, (transactionCounter))
	if _transactionCounterErr != nil {
		return errors.Wrap(_transactionCounterErr, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1ResponseMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1ResponseMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_DF1ResponseMessage) GetPayloadLength() uint16 {
	return m.PayloadLength
}

//
////

func (m *_DF1ResponseMessage) isDF1ResponseMessage() bool {
	return true
}

func (m *_DF1ResponseMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
