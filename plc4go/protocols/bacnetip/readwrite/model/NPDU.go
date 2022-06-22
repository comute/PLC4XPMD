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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// NPDU is the corresponding interface of NPDU
type NPDU interface {
	utils.LengthAware
	utils.Serializable
	// GetProtocolVersionNumber returns ProtocolVersionNumber (property field)
	GetProtocolVersionNumber() uint8
	// GetControl returns Control (property field)
	GetControl() NPDUControl
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() *uint16
	// GetDestinationLength returns DestinationLength (property field)
	GetDestinationLength() *uint8
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() []uint8
	// GetSourceNetworkAddress returns SourceNetworkAddress (property field)
	GetSourceNetworkAddress() *uint16
	// GetSourceLength returns SourceLength (property field)
	GetSourceLength() *uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() []uint8
	// GetHopCount returns HopCount (property field)
	GetHopCount() *uint8
	// GetNlm returns Nlm (property field)
	GetNlm() NLM
	// GetApdu returns Apdu (property field)
	GetApdu() APDU
	// GetDestinationLengthAddon returns DestinationLengthAddon (virtual field)
	GetDestinationLengthAddon() uint16
	// GetSourceLengthAddon returns SourceLengthAddon (virtual field)
	GetSourceLengthAddon() uint16
	// GetPayloadSubtraction returns PayloadSubtraction (virtual field)
	GetPayloadSubtraction() uint16
}

// NPDUExactly can be used when we want exactly this type and not a type which fulfills NPDU.
// This is useful for switch cases.
type NPDUExactly interface {
	NPDU
	isNPDU() bool
}

// _NPDU is the data-structure of this message
type _NPDU struct {
	ProtocolVersionNumber     uint8
	Control                   NPDUControl
	DestinationNetworkAddress *uint16
	DestinationLength         *uint8
	DestinationAddress        []uint8
	SourceNetworkAddress      *uint16
	SourceLength              *uint8
	SourceAddress             []uint8
	HopCount                  *uint8
	Nlm                       NLM
	Apdu                      APDU

	// Arguments.
	NpduLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NPDU) GetProtocolVersionNumber() uint8 {
	return m.ProtocolVersionNumber
}

func (m *_NPDU) GetControl() NPDUControl {
	return m.Control
}

func (m *_NPDU) GetDestinationNetworkAddress() *uint16 {
	return m.DestinationNetworkAddress
}

func (m *_NPDU) GetDestinationLength() *uint8 {
	return m.DestinationLength
}

func (m *_NPDU) GetDestinationAddress() []uint8 {
	return m.DestinationAddress
}

func (m *_NPDU) GetSourceNetworkAddress() *uint16 {
	return m.SourceNetworkAddress
}

func (m *_NPDU) GetSourceLength() *uint8 {
	return m.SourceLength
}

func (m *_NPDU) GetSourceAddress() []uint8 {
	return m.SourceAddress
}

func (m *_NPDU) GetHopCount() *uint8 {
	return m.HopCount
}

func (m *_NPDU) GetNlm() NLM {
	return m.Nlm
}

func (m *_NPDU) GetApdu() APDU {
	return m.Apdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NPDU) GetDestinationLengthAddon() uint16 {
	destinationNetworkAddress := m.DestinationNetworkAddress
	_ = destinationNetworkAddress
	destinationLength := m.DestinationLength
	_ = destinationLength
	sourceNetworkAddress := m.SourceNetworkAddress
	_ = sourceNetworkAddress
	sourceLength := m.SourceLength
	_ = sourceLength
	hopCount := m.HopCount
	_ = hopCount
	nlm := m.Nlm
	_ = nlm
	apdu := m.Apdu
	_ = apdu
	return uint16(utils.InlineIf(m.GetControl().GetDestinationSpecified(), func() interface{} { return uint16(uint16(uint16(uint16(3)) + uint16((*m.GetDestinationLength())))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
}

func (m *_NPDU) GetSourceLengthAddon() uint16 {
	destinationNetworkAddress := m.DestinationNetworkAddress
	_ = destinationNetworkAddress
	destinationLength := m.DestinationLength
	_ = destinationLength
	sourceNetworkAddress := m.SourceNetworkAddress
	_ = sourceNetworkAddress
	sourceLength := m.SourceLength
	_ = sourceLength
	hopCount := m.HopCount
	_ = hopCount
	nlm := m.Nlm
	_ = nlm
	apdu := m.Apdu
	_ = apdu
	return uint16(utils.InlineIf(m.GetControl().GetSourceSpecified(), func() interface{} { return uint16(uint16(uint16(uint16(3)) + uint16((*m.GetSourceLength())))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
}

func (m *_NPDU) GetPayloadSubtraction() uint16 {
	destinationNetworkAddress := m.DestinationNetworkAddress
	_ = destinationNetworkAddress
	destinationLength := m.DestinationLength
	_ = destinationLength
	sourceNetworkAddress := m.SourceNetworkAddress
	_ = sourceNetworkAddress
	sourceLength := m.SourceLength
	_ = sourceLength
	hopCount := m.HopCount
	_ = hopCount
	nlm := m.Nlm
	_ = nlm
	apdu := m.Apdu
	_ = apdu
	return uint16(uint16(uint16(2)) + uint16(uint16(uint16(uint16(m.GetSourceLengthAddon())+uint16(m.GetDestinationLengthAddon()))+uint16(uint16(utils.InlineIf(bool(m.GetControl().GetDestinationSpecified()), func() interface{} { return uint16(uint16(1)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNPDU factory function for _NPDU
func NewNPDU(protocolVersionNumber uint8, control NPDUControl, destinationNetworkAddress *uint16, destinationLength *uint8, destinationAddress []uint8, sourceNetworkAddress *uint16, sourceLength *uint8, sourceAddress []uint8, hopCount *uint8, nlm NLM, apdu APDU, npduLength uint16) *_NPDU {
	return &_NPDU{ProtocolVersionNumber: protocolVersionNumber, Control: control, DestinationNetworkAddress: destinationNetworkAddress, DestinationLength: destinationLength, DestinationAddress: destinationAddress, SourceNetworkAddress: sourceNetworkAddress, SourceLength: sourceLength, SourceAddress: sourceAddress, HopCount: hopCount, Nlm: nlm, Apdu: apdu, NpduLength: npduLength}
}

// Deprecated: use the interface for direct cast
func CastNPDU(structType interface{}) NPDU {
	if casted, ok := structType.(NPDU); ok {
		return casted
	}
	if casted, ok := structType.(*NPDU); ok {
		return *casted
	}
	return nil
}

func (m *_NPDU) GetTypeName() string {
	return "NPDU"
}

func (m *_NPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (protocolVersionNumber)
	lengthInBits += 8

	// Simple field (control)
	lengthInBits += m.Control.GetLengthInBits()

	// Optional Field (destinationNetworkAddress)
	if m.DestinationNetworkAddress != nil {
		lengthInBits += 16
	}

	// Optional Field (destinationLength)
	if m.DestinationLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (sourceNetworkAddress)
	if m.SourceNetworkAddress != nil {
		lengthInBits += 16
	}

	// Optional Field (sourceLength)
	if m.SourceLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.SourceAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.SourceAddress))
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (hopCount)
	if m.HopCount != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (nlm)
	if m.Nlm != nil {
		lengthInBits += m.Nlm.GetLengthInBits()
	}

	// Optional Field (apdu)
	if m.Apdu != nil {
		lengthInBits += m.Apdu.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_NPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NPDUParse(readBuffer utils.ReadBuffer, npduLength uint16) (NPDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolVersionNumber)
	_protocolVersionNumber, _protocolVersionNumberErr := readBuffer.ReadUint8("protocolVersionNumber", 8)
	if _protocolVersionNumberErr != nil {
		return nil, errors.Wrap(_protocolVersionNumberErr, "Error parsing 'protocolVersionNumber' field")
	}
	protocolVersionNumber := _protocolVersionNumber

	// Simple Field (control)
	if pullErr := readBuffer.PullContext("control"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for control")
	}
	_control, _controlErr := NPDUControlParse(readBuffer)
	if _controlErr != nil {
		return nil, errors.Wrap(_controlErr, "Error parsing 'control' field")
	}
	control := _control.(NPDUControl)
	if closeErr := readBuffer.CloseContext("control"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for control")
	}

	// Optional Field (destinationNetworkAddress) (Can be skipped, if a given expression evaluates to false)
	var destinationNetworkAddress *uint16 = nil
	if control.GetDestinationSpecified() {
		_val, _err := readBuffer.ReadUint16("destinationNetworkAddress", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field")
		}
		destinationNetworkAddress = &_val
	}

	// Optional Field (destinationLength) (Can be skipped, if a given expression evaluates to false)
	var destinationLength *uint8 = nil
	if control.GetDestinationSpecified() {
		_val, _err := readBuffer.ReadUint8("destinationLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'destinationLength' field")
		}
		destinationLength = &_val
	}

	// Array field (destinationAddress)
	if pullErr := readBuffer.PullContext("destinationAddress", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destinationAddress")
	}
	// Count array
	destinationAddress := make([]uint8, utils.InlineIf(control.GetDestinationSpecified(), func() interface{} { return uint16((*destinationLength)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	{
		for curItem := uint16(0); curItem < uint16(utils.InlineIf(control.GetDestinationSpecified(), func() interface{} { return uint16((*destinationLength)) }, func() interface{} { return uint16(uint16(0)) }).(uint16)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'destinationAddress' field")
			}
			destinationAddress[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("destinationAddress", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destinationAddress")
	}

	// Virtual field
	_destinationLengthAddon := utils.InlineIf(control.GetDestinationSpecified(), func() interface{} { return uint16(uint16(uint16(uint16(3)) + uint16((*destinationLength)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)
	destinationLengthAddon := uint16(_destinationLengthAddon)
	_ = destinationLengthAddon

	// Optional Field (sourceNetworkAddress) (Can be skipped, if a given expression evaluates to false)
	var sourceNetworkAddress *uint16 = nil
	if control.GetSourceSpecified() {
		_val, _err := readBuffer.ReadUint16("sourceNetworkAddress", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sourceNetworkAddress' field")
		}
		sourceNetworkAddress = &_val
	}

	// Optional Field (sourceLength) (Can be skipped, if a given expression evaluates to false)
	var sourceLength *uint8 = nil
	if control.GetSourceSpecified() {
		_val, _err := readBuffer.ReadUint8("sourceLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sourceLength' field")
		}
		sourceLength = &_val
	}

	// Array field (sourceAddress)
	if pullErr := readBuffer.PullContext("sourceAddress", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceAddress")
	}
	// Count array
	sourceAddress := make([]uint8, utils.InlineIf(control.GetSourceSpecified(), func() interface{} { return uint16((*sourceLength)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	{
		for curItem := uint16(0); curItem < uint16(utils.InlineIf(control.GetSourceSpecified(), func() interface{} { return uint16((*sourceLength)) }, func() interface{} { return uint16(uint16(0)) }).(uint16)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'sourceAddress' field")
			}
			sourceAddress[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("sourceAddress", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceAddress")
	}

	// Virtual field
	_sourceLengthAddon := utils.InlineIf(control.GetSourceSpecified(), func() interface{} { return uint16(uint16(uint16(uint16(3)) + uint16((*sourceLength)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)
	sourceLengthAddon := uint16(_sourceLengthAddon)
	_ = sourceLengthAddon

	// Optional Field (hopCount) (Can be skipped, if a given expression evaluates to false)
	var hopCount *uint8 = nil
	if control.GetDestinationSpecified() {
		_val, _err := readBuffer.ReadUint8("hopCount", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'hopCount' field")
		}
		hopCount = &_val
	}

	// Virtual field
	_payloadSubtraction := uint16(uint16(2)) + uint16(uint16(uint16(uint16(sourceLengthAddon)+uint16(destinationLengthAddon))+uint16(uint16(utils.InlineIf(bool(control.GetDestinationSpecified()), func() interface{} { return uint16(uint16(1)) }, func() interface{} { return uint16(uint16(0)) }).(uint16)))))
	payloadSubtraction := uint16(_payloadSubtraction)
	_ = payloadSubtraction

	// Optional Field (nlm) (Can be skipped, if a given expression evaluates to false)
	var nlm NLM = nil
	if control.GetMessageTypeFieldPresent() {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("nlm"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for nlm")
		}
		_val, _err := NLMParse(readBuffer, uint16(npduLength)-uint16(payloadSubtraction))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'nlm' field")
		default:
			nlm = _val.(NLM)
			if closeErr := readBuffer.CloseContext("nlm"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for nlm")
			}
		}
	}

	// Optional Field (apdu) (Can be skipped, if a given expression evaluates to false)
	var apdu APDU = nil
	if !(control.GetMessageTypeFieldPresent()) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("apdu"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for apdu")
		}
		_val, _err := APDUParse(readBuffer, uint16(npduLength)-uint16(payloadSubtraction))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'apdu' field")
		default:
			apdu = _val.(APDU)
			if closeErr := readBuffer.CloseContext("apdu"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for apdu")
			}
		}
	}

	// Validation
	if !(bool(bool((nlm) != (nil))) || bool(bool((apdu) != (nil)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"something is wrong here... apdu and nlm not set"})
	}

	if closeErr := readBuffer.CloseContext("NPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NPDU")
	}

	// Create the instance
	return NewNPDU(protocolVersionNumber, control, destinationNetworkAddress, destinationLength, destinationAddress, sourceNetworkAddress, sourceLength, sourceAddress, hopCount, nlm, apdu, npduLength), nil
}

func (m *_NPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NPDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NPDU")
	}

	// Simple Field (protocolVersionNumber)
	protocolVersionNumber := uint8(m.GetProtocolVersionNumber())
	_protocolVersionNumberErr := writeBuffer.WriteUint8("protocolVersionNumber", 8, (protocolVersionNumber))
	if _protocolVersionNumberErr != nil {
		return errors.Wrap(_protocolVersionNumberErr, "Error serializing 'protocolVersionNumber' field")
	}

	// Simple Field (control)
	if pushErr := writeBuffer.PushContext("control"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for control")
	}
	_controlErr := writeBuffer.WriteSerializable(m.GetControl())
	if popErr := writeBuffer.PopContext("control"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for control")
	}
	if _controlErr != nil {
		return errors.Wrap(_controlErr, "Error serializing 'control' field")
	}

	// Optional Field (destinationNetworkAddress) (Can be skipped, if the value is null)
	var destinationNetworkAddress *uint16 = nil
	if m.GetDestinationNetworkAddress() != nil {
		destinationNetworkAddress = m.GetDestinationNetworkAddress()
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, *(destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}
	}

	// Optional Field (destinationLength) (Can be skipped, if the value is null)
	var destinationLength *uint8 = nil
	if m.GetDestinationLength() != nil {
		destinationLength = m.GetDestinationLength()
		_destinationLengthErr := writeBuffer.WriteUint8("destinationLength", 8, *(destinationLength))
		if _destinationLengthErr != nil {
			return errors.Wrap(_destinationLengthErr, "Error serializing 'destinationLength' field")
		}
	}

	// Array Field (destinationAddress)
	if m.GetDestinationAddress() != nil {
		if pushErr := writeBuffer.PushContext("destinationAddress", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for destinationAddress")
		}
		for _, _element := range m.GetDestinationAddress() {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'destinationAddress' field")
			}
		}
		if popErr := writeBuffer.PopContext("destinationAddress", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for destinationAddress")
		}
	}
	// Virtual field
	if _destinationLengthAddonErr := writeBuffer.WriteVirtual("destinationLengthAddon", m.GetDestinationLengthAddon()); _destinationLengthAddonErr != nil {
		return errors.Wrap(_destinationLengthAddonErr, "Error serializing 'destinationLengthAddon' field")
	}

	// Optional Field (sourceNetworkAddress) (Can be skipped, if the value is null)
	var sourceNetworkAddress *uint16 = nil
	if m.GetSourceNetworkAddress() != nil {
		sourceNetworkAddress = m.GetSourceNetworkAddress()
		_sourceNetworkAddressErr := writeBuffer.WriteUint16("sourceNetworkAddress", 16, *(sourceNetworkAddress))
		if _sourceNetworkAddressErr != nil {
			return errors.Wrap(_sourceNetworkAddressErr, "Error serializing 'sourceNetworkAddress' field")
		}
	}

	// Optional Field (sourceLength) (Can be skipped, if the value is null)
	var sourceLength *uint8 = nil
	if m.GetSourceLength() != nil {
		sourceLength = m.GetSourceLength()
		_sourceLengthErr := writeBuffer.WriteUint8("sourceLength", 8, *(sourceLength))
		if _sourceLengthErr != nil {
			return errors.Wrap(_sourceLengthErr, "Error serializing 'sourceLength' field")
		}
	}

	// Array Field (sourceAddress)
	if m.GetSourceAddress() != nil {
		if pushErr := writeBuffer.PushContext("sourceAddress", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sourceAddress")
		}
		for _, _element := range m.GetSourceAddress() {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'sourceAddress' field")
			}
		}
		if popErr := writeBuffer.PopContext("sourceAddress", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sourceAddress")
		}
	}
	// Virtual field
	if _sourceLengthAddonErr := writeBuffer.WriteVirtual("sourceLengthAddon", m.GetSourceLengthAddon()); _sourceLengthAddonErr != nil {
		return errors.Wrap(_sourceLengthAddonErr, "Error serializing 'sourceLengthAddon' field")
	}

	// Optional Field (hopCount) (Can be skipped, if the value is null)
	var hopCount *uint8 = nil
	if m.GetHopCount() != nil {
		hopCount = m.GetHopCount()
		_hopCountErr := writeBuffer.WriteUint8("hopCount", 8, *(hopCount))
		if _hopCountErr != nil {
			return errors.Wrap(_hopCountErr, "Error serializing 'hopCount' field")
		}
	}
	// Virtual field
	if _payloadSubtractionErr := writeBuffer.WriteVirtual("payloadSubtraction", m.GetPayloadSubtraction()); _payloadSubtractionErr != nil {
		return errors.Wrap(_payloadSubtractionErr, "Error serializing 'payloadSubtraction' field")
	}

	// Optional Field (nlm) (Can be skipped, if the value is null)
	var nlm NLM = nil
	if m.GetNlm() != nil {
		if pushErr := writeBuffer.PushContext("nlm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nlm")
		}
		nlm = m.GetNlm()
		_nlmErr := writeBuffer.WriteSerializable(nlm)
		if popErr := writeBuffer.PopContext("nlm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nlm")
		}
		if _nlmErr != nil {
			return errors.Wrap(_nlmErr, "Error serializing 'nlm' field")
		}
	}

	// Optional Field (apdu) (Can be skipped, if the value is null)
	var apdu APDU = nil
	if m.GetApdu() != nil {
		if pushErr := writeBuffer.PushContext("apdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for apdu")
		}
		apdu = m.GetApdu()
		_apduErr := writeBuffer.WriteSerializable(apdu)
		if popErr := writeBuffer.PopContext("apdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for apdu")
		}
		if _apduErr != nil {
			return errors.Wrap(_apduErr, "Error serializing 'apdu' field")
		}
	}

	if popErr := writeBuffer.PopContext("NPDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NPDU")
	}
	return nil
}

func (m *_NPDU) isNPDU() bool {
	return true
}

func (m *_NPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
