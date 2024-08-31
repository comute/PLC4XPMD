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

// MonitoredSALLongFormSmartMode is the corresponding interface of MonitoredSALLongFormSmartMode
type MonitoredSALLongFormSmartMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MonitoredSAL
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() ReplyNetwork
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
}

// MonitoredSALLongFormSmartModeExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALLongFormSmartMode.
// This is useful for switch cases.
type MonitoredSALLongFormSmartModeExactly interface {
	MonitoredSALLongFormSmartMode
	isMonitoredSALLongFormSmartMode() bool
}

// _MonitoredSALLongFormSmartMode is the data-structure of this message
type _MonitoredSALLongFormSmartMode struct {
	*_MonitoredSAL
	TerminatingByte uint32
	UnitAddress     UnitAddress
	BridgeAddress   BridgeAddress
	Application     ApplicationIdContainer
	ReservedByte    *byte
	ReplyNetwork    ReplyNetwork
	SalData         SALData
	// Reserved Fields
	reservedField0 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALLongFormSmartMode) InitializeParent(parent MonitoredSAL, salType byte) {
	m.SalType = salType
}

func (m *_MonitoredSALLongFormSmartMode) GetParent() MonitoredSAL {
	return m._MonitoredSAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *_MonitoredSALLongFormSmartMode) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALLongFormSmartMode) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *_MonitoredSALLongFormSmartMode) GetReplyNetwork() ReplyNetwork {
	return m.ReplyNetwork
}

func (m *_MonitoredSALLongFormSmartMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetIsUnitAddress() bool {
	ctx := context.Background()
	_ = ctx
	unitAddress := m.UnitAddress
	_ = unitAddress
	bridgeAddress := m.BridgeAddress
	_ = bridgeAddress
	reservedByte := m.ReservedByte
	_ = reservedByte
	replyNetwork := m.ReplyNetwork
	_ = replyNetwork
	salData := m.SalData
	_ = salData
	return bool(bool((m.GetTerminatingByte() & 0xff) == (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSALLongFormSmartMode factory function for _MonitoredSALLongFormSmartMode
func NewMonitoredSALLongFormSmartMode(terminatingByte uint32, unitAddress UnitAddress, bridgeAddress BridgeAddress, application ApplicationIdContainer, reservedByte *byte, replyNetwork ReplyNetwork, salData SALData, salType byte, cBusOptions CBusOptions) *_MonitoredSALLongFormSmartMode {
	_result := &_MonitoredSALLongFormSmartMode{
		TerminatingByte: terminatingByte,
		UnitAddress:     unitAddress,
		BridgeAddress:   bridgeAddress,
		Application:     application,
		ReservedByte:    reservedByte,
		ReplyNetwork:    replyNetwork,
		SalData:         salData,
		_MonitoredSAL:   NewMonitoredSAL(salType, cBusOptions),
	}
	_result._MonitoredSAL._MonitoredSALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALLongFormSmartMode(structType any) MonitoredSALLongFormSmartMode {
	if casted, ok := structType.(MonitoredSALLongFormSmartMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALLongFormSmartMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALLongFormSmartMode) GetTypeName() string {
	return "MonitoredSALLongFormSmartMode"
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += m.UnitAddress.GetLengthInBits(ctx)
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += m.ReplyNetwork.GetLengthInBits(ctx)
	}

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALLongFormSmartModeParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (MonitoredSALLongFormSmartMode, error) {
	return MonitoredSALLongFormSmartModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func MonitoredSALLongFormSmartModeParseWithBufferProducer(cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALLongFormSmartMode, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALLongFormSmartMode, error) {
		return MonitoredSALLongFormSmartModeParseWithBuffer(ctx, readBuffer, cBusOptions)
	}
}

func MonitoredSALLongFormSmartModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (MonitoredSALLongFormSmartMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALLongFormSmartMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALLongFormSmartMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x05))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	terminatingByte, err := ReadPeekField[uint32](ctx, "terminatingByte", ReadUnsignedInt(readBuffer, uint8(24)), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminatingByte' field"))
	}

	isUnitAddress, err := ReadVirtualField[bool](ctx, "isUnitAddress", (*bool)(nil), bool((terminatingByte&0xff) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnitAddress' field"))
	}
	_ = isUnitAddress

	_unitAddress, err := ReadOptionalField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	var unitAddress UnitAddress
	if _unitAddress != nil {
		unitAddress = *_unitAddress
	}

	_bridgeAddress, err := ReadOptionalField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	var bridgeAddress BridgeAddress
	if _bridgeAddress != nil {
		bridgeAddress = *_bridgeAddress
	}

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}

	reservedByte, err := ReadOptionalField[byte](ctx, "reservedByte", ReadByte(readBuffer, 8), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reservedByte' field"))
	}

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "invalid unit address"})
	}

	_replyNetwork, err := ReadOptionalField[ReplyNetwork](ctx, "replyNetwork", ReadComplex[ReplyNetwork](ReplyNetworkParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'replyNetwork' field"))
	}
	var replyNetwork ReplyNetwork
	if _replyNetwork != nil {
		replyNetwork = *_replyNetwork
	}

	_salData, err := ReadOptionalField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(application.ApplicationId())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	var salData SALData
	if _salData != nil {
		salData = *_salData
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALLongFormSmartMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALLongFormSmartMode")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALLongFormSmartMode{
		_MonitoredSAL: &_MonitoredSAL{
			CBusOptions: cBusOptions,
		},
		TerminatingByte: terminatingByte,
		UnitAddress:     unitAddress,
		BridgeAddress:   bridgeAddress,
		Application:     application,
		ReservedByte:    reservedByte,
		ReplyNetwork:    replyNetwork,
		SalData:         salData,
		reservedField0:  reservedField0,
	}
	_child._MonitoredSAL._MonitoredSALChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALLongFormSmartMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALLongFormSmartMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALLongFormSmartMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALLongFormSmartMode")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x05)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0x05),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}
		// Virtual field
		isUnitAddress := m.GetIsUnitAddress()
		_ = isUnitAddress
		if _isUnitAddressErr := writeBuffer.WriteVirtual(ctx, "isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		// Optional Field (unitAddress) (Can be skipped, if the value is null)
		var unitAddress UnitAddress = nil
		if m.GetUnitAddress() != nil {
			if pushErr := writeBuffer.PushContext("unitAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for unitAddress")
			}
			unitAddress = m.GetUnitAddress()
			_unitAddressErr := writeBuffer.WriteSerializable(ctx, unitAddress)
			if popErr := writeBuffer.PopContext("unitAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for unitAddress")
			}
			if _unitAddressErr != nil {
				return errors.Wrap(_unitAddressErr, "Error serializing 'unitAddress' field")
			}
		}

		// Optional Field (bridgeAddress) (Can be skipped, if the value is null)
		var bridgeAddress BridgeAddress = nil
		if m.GetBridgeAddress() != nil {
			if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bridgeAddress")
			}
			bridgeAddress = m.GetBridgeAddress()
			_bridgeAddressErr := writeBuffer.WriteSerializable(ctx, bridgeAddress)
			if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bridgeAddress")
			}
			if _bridgeAddressErr != nil {
				return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Optional Field (reservedByte) (Can be skipped, if the value is null)
		var reservedByte *byte = nil
		if m.GetReservedByte() != nil {
			reservedByte = m.GetReservedByte()
			_reservedByteErr := /*TODO: migrate me*/ writeBuffer.WriteByte("reservedByte", *(reservedByte))
			if _reservedByteErr != nil {
				return errors.Wrap(_reservedByteErr, "Error serializing 'reservedByte' field")
			}
		}

		// Optional Field (replyNetwork) (Can be skipped, if the value is null)
		var replyNetwork ReplyNetwork = nil
		if m.GetReplyNetwork() != nil {
			if pushErr := writeBuffer.PushContext("replyNetwork"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for replyNetwork")
			}
			replyNetwork = m.GetReplyNetwork()
			_replyNetworkErr := writeBuffer.WriteSerializable(ctx, replyNetwork)
			if popErr := writeBuffer.PopContext("replyNetwork"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for replyNetwork")
			}
			if _replyNetworkErr != nil {
				return errors.Wrap(_replyNetworkErr, "Error serializing 'replyNetwork' field")
			}
		}

		// Optional Field (salData) (Can be skipped, if the value is null)
		var salData SALData = nil
		if m.GetSalData() != nil {
			if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for salData")
			}
			salData = m.GetSalData()
			_salDataErr := writeBuffer.WriteSerializable(ctx, salData)
			if popErr := writeBuffer.PopContext("salData"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for salData")
			}
			if _salDataErr != nil {
				return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
			}
		}

		if popErr := writeBuffer.PopContext("MonitoredSALLongFormSmartMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALLongFormSmartMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALLongFormSmartMode) isMonitoredSALLongFormSmartMode() bool {
	return true
}

func (m *_MonitoredSALLongFormSmartMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
