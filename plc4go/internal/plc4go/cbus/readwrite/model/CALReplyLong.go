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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CALReplyLong struct {
	*CALReply
	TerminatingByte        uint32
	UnitAddress            *UnitAddress
	BridgeAddress          *BridgeAddress
	SerialInterfaceAddress *SerialInterfaceAddress
	ReservedByte           *byte
	ReplyNetwork           *ReplyNetwork
}

// The corresponding interface
type ICALReplyLong interface {
	ICALReply
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() *UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() *BridgeAddress
	// GetSerialInterfaceAddress returns SerialInterfaceAddress (property field)
	GetSerialInterfaceAddress() *SerialInterfaceAddress
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() *ReplyNetwork
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
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

func (m *CALReplyLong) InitializeParent(parent *CALReply, calType byte, calData *CALData) {
	m.CALReply.CalType = calType
	m.CALReply.CalData = calData
}

func (m *CALReplyLong) GetParent() *CALReply {
	return m.CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *CALReplyLong) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *CALReplyLong) GetUnitAddress() *UnitAddress {
	return m.UnitAddress
}

func (m *CALReplyLong) GetBridgeAddress() *BridgeAddress {
	return m.BridgeAddress
}

func (m *CALReplyLong) GetSerialInterfaceAddress() *SerialInterfaceAddress {
	return m.SerialInterfaceAddress
}

func (m *CALReplyLong) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *CALReplyLong) GetReplyNetwork() *ReplyNetwork {
	return m.ReplyNetwork
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *CALReplyLong) GetIsUnitAddress() bool {
	unitAddress := m.UnitAddress
	_ = unitAddress
	bridgeAddress := m.BridgeAddress
	_ = bridgeAddress
	reservedByte := m.ReservedByte
	_ = reservedByte
	replyNetwork := m.ReplyNetwork
	_ = replyNetwork
	return bool(bool(((m.GetTerminatingByte()) & (0xff)) == (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReplyLong factory function for CALReplyLong
func NewCALReplyLong(terminatingByte uint32, unitAddress *UnitAddress, bridgeAddress *BridgeAddress, serialInterfaceAddress *SerialInterfaceAddress, reservedByte *byte, replyNetwork *ReplyNetwork, calType byte, calData *CALData) *CALReplyLong {
	_result := &CALReplyLong{
		TerminatingByte:        terminatingByte,
		UnitAddress:            unitAddress,
		BridgeAddress:          bridgeAddress,
		SerialInterfaceAddress: serialInterfaceAddress,
		ReservedByte:           reservedByte,
		ReplyNetwork:           replyNetwork,
		CALReply:               NewCALReply(calType, calData),
	}
	_result.Child = _result
	return _result
}

func CastCALReplyLong(structType interface{}) *CALReplyLong {
	if casted, ok := structType.(CALReplyLong); ok {
		return &casted
	}
	if casted, ok := structType.(*CALReplyLong); ok {
		return casted
	}
	if casted, ok := structType.(CALReply); ok {
		return CastCALReplyLong(casted.Child)
	}
	if casted, ok := structType.(*CALReply); ok {
		return CastCALReplyLong(casted.Child)
	}
	return nil
}

func (m *CALReplyLong) GetTypeName() string {
	return "CALReplyLong"
}

func (m *CALReplyLong) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALReplyLong) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += (*m.UnitAddress).GetLengthInBits()
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += (*m.BridgeAddress).GetLengthInBits()
	}

	// Simple field (serialInterfaceAddress)
	lengthInBits += m.SerialInterfaceAddress.GetLengthInBits()

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += (*m.ReplyNetwork).GetLengthInBits()
	}

	return lengthInBits
}

func (m *CALReplyLong) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyLongParse(readBuffer utils.ReadBuffer) (*CALReplyLong, error) {
	if pullErr := readBuffer.PullContext("CALReplyLong"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x86) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x86),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Peek Field (terminatingByte)
	currentPos = readBuffer.GetPos()
	terminatingByte, _err := readBuffer.ReadUint32("terminatingByte", 24)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'terminatingByte' field")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isUnitAddress := bool(((terminatingByte) & (0xff)) == (0x00))
	isUnitAddress := bool(_isUnitAddress)
	_ = isUnitAddress

	// Optional Field (unitAddress) (Can be skipped, if a given expression evaluates to false)
	var unitAddress *UnitAddress = nil
	if isUnitAddress {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("unitAddress"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := UnitAddressParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'unitAddress' field")
		default:
			unitAddress = CastUnitAddress(_val)
			if closeErr := readBuffer.CloseContext("unitAddress"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (bridgeAddress) (Can be skipped, if a given expression evaluates to false)
	var bridgeAddress *BridgeAddress = nil
	if !(isUnitAddress) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BridgeAddressParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'bridgeAddress' field")
		default:
			bridgeAddress = CastBridgeAddress(_val)
			if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (serialInterfaceAddress)
	if pullErr := readBuffer.PullContext("serialInterfaceAddress"); pullErr != nil {
		return nil, pullErr
	}
	_serialInterfaceAddress, _serialInterfaceAddressErr := SerialInterfaceAddressParse(readBuffer)
	if _serialInterfaceAddressErr != nil {
		return nil, errors.Wrap(_serialInterfaceAddressErr, "Error parsing 'serialInterfaceAddress' field")
	}
	serialInterfaceAddress := CastSerialInterfaceAddress(_serialInterfaceAddress)
	if closeErr := readBuffer.CloseContext("serialInterfaceAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (reservedByte) (Can be skipped, if a given expression evaluates to false)
	var reservedByte *byte = nil
	if isUnitAddress {
		_val, _err := readBuffer.ReadByte("reservedByte")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reservedByte' field")
		}
		reservedByte = &_val
	}

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, utils.ParseAssertError{"wrong reservedByte"}
	}

	// Optional Field (replyNetwork) (Can be skipped, if a given expression evaluates to false)
	var replyNetwork *ReplyNetwork = nil
	if !(isUnitAddress) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("replyNetwork"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := ReplyNetworkParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'replyNetwork' field")
		default:
			replyNetwork = CastReplyNetwork(_val)
			if closeErr := readBuffer.CloseContext("replyNetwork"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("CALReplyLong"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALReplyLong{
		TerminatingByte:        terminatingByte,
		UnitAddress:            CastUnitAddress(unitAddress),
		BridgeAddress:          CastBridgeAddress(bridgeAddress),
		SerialInterfaceAddress: CastSerialInterfaceAddress(serialInterfaceAddress),
		ReservedByte:           reservedByte,
		ReplyNetwork:           CastReplyNetwork(replyNetwork),
		CALReply:               &CALReply{},
	}
	_child.CALReply.Child = _child
	return _child, nil
}

func (m *CALReplyLong) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyLong"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x86))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}
		// Virtual field
		if _isUnitAddressErr := writeBuffer.WriteVirtual("isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		// Optional Field (unitAddress) (Can be skipped, if the value is null)
		var unitAddress *UnitAddress = nil
		if m.UnitAddress != nil {
			if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
				return pushErr
			}
			unitAddress = m.UnitAddress
			_unitAddressErr := unitAddress.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
				return popErr
			}
			if _unitAddressErr != nil {
				return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
			}
		}

		// Optional Field (bridgeAddress) (Can be skipped, if the value is null)
		var bridgeAddress *BridgeAddress = nil
		if m.BridgeAddress != nil {
			if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
				return pushErr
			}
			bridgeAddress = m.BridgeAddress
			_bridgeAddressErr := bridgeAddress.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
				return popErr
			}
			if _bridgeAddressErr != nil {
				return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
			}
		}

		// Simple Field (serialInterfaceAddress)
		if pushErr := writeBuffer.PushContext("serialInterfaceAddress"); pushErr != nil {
			return pushErr
		}
		_serialInterfaceAddressErr := m.SerialInterfaceAddress.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serialInterfaceAddress"); popErr != nil {
			return popErr
		}
		if _serialInterfaceAddressErr != nil {
			return errors.Wrap(_serialInterfaceAddressErr, "Error serializing 'serialInterfaceAddress' field")
		}

		// Optional Field (reservedByte) (Can be skipped, if the value is null)
		var reservedByte *byte = nil
		if m.ReservedByte != nil {
			reservedByte = m.ReservedByte
			_reservedByteErr := writeBuffer.WriteByte("reservedByte", *(reservedByte))
			if _reservedByteErr != nil {
				return errors.Wrap(_reservedByteErr, "Error serializing 'reservedByte' field")
			}
		}

		// Optional Field (replyNetwork) (Can be skipped, if the value is null)
		var replyNetwork *ReplyNetwork = nil
		if m.ReplyNetwork != nil {
			if pushErr := writeBuffer.PushContext("replyNetwork"); pushErr != nil {
				return pushErr
			}
			replyNetwork = m.ReplyNetwork
			_replyNetworkErr := replyNetwork.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("replyNetwork"); popErr != nil {
				return popErr
			}
			if _replyNetworkErr != nil {
				return errors.Wrap(_replyNetworkErr, "Error serializing 'replyNetwork' field")
			}
		}

		if popErr := writeBuffer.PopContext("CALReplyLong"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALReplyLong) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
