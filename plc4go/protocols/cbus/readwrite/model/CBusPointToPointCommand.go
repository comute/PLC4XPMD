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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusPointToPointCommand_CR byte = 0xD

// CBusPointToPointCommand is the data-structure of this message
type CBusPointToPointCommand struct {
	BridgeAddressCountPeek uint16
	CalData                *CALData
	Crc                    *Checksum
	PeekAlpha              byte
	Alpha                  *Alpha

	// Arguments.
	Srchk bool
	Child ICBusPointToPointCommandChild
}

// ICBusPointToPointCommand is the corresponding interface of CBusPointToPointCommand
type ICBusPointToPointCommand interface {
	// GetBridgeAddressCountPeek returns BridgeAddressCountPeek (property field)
	GetBridgeAddressCountPeek() uint16
	// GetCalData returns CalData (property field)
	GetCalData() *CALData
	// GetCrc returns Crc (property field)
	GetCrc() *Checksum
	// GetPeekAlpha returns PeekAlpha (property field)
	GetPeekAlpha() byte
	// GetAlpha returns Alpha (property field)
	GetAlpha() *Alpha
	// GetIsDirect returns IsDirect (virtual field)
	GetIsDirect() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICBusPointToPointCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICBusPointToPointCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICBusPointToPointCommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CBusPointToPointCommand, bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha)
	GetParent() *CBusPointToPointCommand

	GetTypeName() string
	ICBusPointToPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CBusPointToPointCommand) GetBridgeAddressCountPeek() uint16 {
	return m.BridgeAddressCountPeek
}

func (m *CBusPointToPointCommand) GetCalData() *CALData {
	return m.CalData
}

func (m *CBusPointToPointCommand) GetCrc() *Checksum {
	return m.Crc
}

func (m *CBusPointToPointCommand) GetPeekAlpha() byte {
	return m.PeekAlpha
}

func (m *CBusPointToPointCommand) GetAlpha() *Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *CBusPointToPointCommand) GetIsDirect() bool {
	crc := m.Crc
	_ = crc
	alpha := m.Alpha
	_ = alpha
	return bool(bool(((m.GetBridgeAddressCountPeek()) & (0x00FF)) == (0x0000)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *CBusPointToPointCommand) GetCr() byte {
	return CBusPointToPointCommand_CR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointCommand factory function for CBusPointToPointCommand
func NewCBusPointToPointCommand(bridgeAddressCountPeek uint16, calData *CALData, crc *Checksum, peekAlpha byte, alpha *Alpha, srchk bool) *CBusPointToPointCommand {
	return &CBusPointToPointCommand{BridgeAddressCountPeek: bridgeAddressCountPeek, CalData: calData, Crc: crc, PeekAlpha: peekAlpha, Alpha: alpha, Srchk: srchk}
}

func CastCBusPointToPointCommand(structType interface{}) *CBusPointToPointCommand {
	if casted, ok := structType.(CBusPointToPointCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusPointToPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(ICBusPointToPointCommandChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *CBusPointToPointCommand) GetTypeName() string {
	return "CBusPointToPointCommand"
}

func (m *CBusPointToPointCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusPointToPointCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CBusPointToPointCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += (*m.Crc).GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += (*m.Alpha).GetLengthInBits()
	}

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *CBusPointToPointCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointCommandParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (bridgeAddressCountPeek)
	currentPos = positionAware.GetPos()
	bridgeAddressCountPeek, _err := readBuffer.ReadUint16("bridgeAddressCountPeek", 16)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'bridgeAddressCountPeek' field")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isDirect := bool(((bridgeAddressCountPeek) & (0x00FF)) == (0x0000))
	isDirect := bool(_isDirect)
	_ = isDirect

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusPointToPointCommandChild interface {
		InitializeParent(*CBusPointToPointCommand, uint16, *CALData, *Checksum, byte, *Alpha)
		GetParent() *CBusPointToPointCommand
	}
	var _child CBusPointToPointCommandChild
	var typeSwitchError error
	switch {
	case isDirect == bool(true): // CBusPointToPointCommandDirect
		_child, typeSwitchError = CBusPointToPointCommandDirectParse(readBuffer, srchk)
	case isDirect == bool(false): // CBusPointToPointCommandIndirect
		_child, typeSwitchError = CBusPointToPointCommandIndirectParse(readBuffer, srchk)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (calData)
	if pullErr := readBuffer.PullContext("calData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calData")
	}
	_calData, _calDataErr := CALDataParse(readBuffer)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field")
	}
	calData := CastCALData(_calData)
	if closeErr := readBuffer.CloseContext("calData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calData")
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *Checksum = nil
	if srchk {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for crc")
		}
		_val, _err := ChecksumParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		default:
			crc = CastChecksum(_val)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for crc")
			}
		}
	}

	// Peek Field (peekAlpha)
	currentPos = positionAware.GetPos()
	peekAlpha, _err := readBuffer.ReadByte("peekAlpha")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekAlpha' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha *Alpha = nil
	if bool(bool(bool((peekAlpha) >= (0x67)))) && bool(bool(bool((peekAlpha) <= (0x7A)))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alpha")
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field")
		default:
			alpha = CastAlpha(_val)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CBusPointToPointCommand_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusPointToPointCommand_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), bridgeAddressCountPeek, calData, crc, peekAlpha, alpha)
	return _child.GetParent(), nil
}

func (m *CBusPointToPointCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CBusPointToPointCommand) SerializeParent(writeBuffer utils.WriteBuffer, child ICBusPointToPointCommand, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusPointToPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommand")
	}
	// Virtual field
	if _isDirectErr := writeBuffer.WriteVirtual("isDirect", m.GetIsDirect()); _isDirectErr != nil {
		return errors.Wrap(_isDirectErr, "Error serializing 'isDirect' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (calData)
	if pushErr := writeBuffer.PushContext("calData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for calData")
	}
	_calDataErr := writeBuffer.WriteSerializable(m.CalData)
	if popErr := writeBuffer.PopContext("calData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for calData")
	}
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}

	// Optional Field (crc) (Can be skipped, if the value is null)
	var crc *Checksum = nil
	if m.Crc != nil {
		if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for crc")
		}
		crc = m.Crc
		_crcErr := writeBuffer.WriteSerializable(crc)
		if popErr := writeBuffer.PopContext("crc"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for crc")
		}
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}
	}

	// Optional Field (alpha) (Can be skipped, if the value is null)
	var alpha *Alpha = nil
	if m.Alpha != nil {
		if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alpha")
		}
		alpha = m.Alpha
		_alphaErr := writeBuffer.WriteSerializable(alpha)
		if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alpha")
		}
		if _alphaErr != nil {
			return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
		}
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0xD)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointCommand")
	}
	return nil
}

func (m *CBusPointToPointCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
