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
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCpuFunctionMsgSubscription is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscription
type S7PayloadUserDataItemCpuFunctionMsgSubscription interface {
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetSubscription returns Subscription (property field)
	GetSubscription() uint8
	// GetMagicKey returns MagicKey (property field)
	GetMagicKey() string
	// GetAlarmtype returns Alarmtype (property field)
	GetAlarmtype() *AlarmStateType
	// GetReserve returns Reserve (property field)
	GetReserve() *uint8
}

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionMsgSubscription.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionExactly interface {
	S7PayloadUserDataItemCpuFunctionMsgSubscription
	isS7PayloadUserDataItemCpuFunctionMsgSubscription() bool
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscription is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscription struct {
	*_S7PayloadUserDataItem
	Subscription uint8
	MagicKey     string
	Alarmtype    *AlarmStateType
	Reserve      *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetCpuSubfunction() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetSubscription() uint8 {
	return m.Subscription
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetMagicKey() string {
	return m.MagicKey
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetAlarmtype() *AlarmStateType {
	return m.Alarmtype
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetReserve() *uint8 {
	return m.Reserve
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscription factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscription
func NewS7PayloadUserDataItemCpuFunctionMsgSubscription(Subscription uint8, magicKey string, Alarmtype *AlarmStateType, Reserve *uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadUserDataItemCpuFunctionMsgSubscription {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscription{
		Subscription:           Subscription,
		MagicKey:               magicKey,
		Alarmtype:              Alarmtype,
		Reserve:                Reserve,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscription(structType interface{}) S7PayloadUserDataItemCpuFunctionMsgSubscription {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscription"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (Subscription)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (magicKey)
	lengthInBits += 64

	// Optional Field (Alarmtype)
	if m.Alarmtype != nil {
		lengthInBits += 8
	}

	// Optional Field (Reserve)
	if m.Reserve != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (Subscription)
	_Subscription, _SubscriptionErr := readBuffer.ReadUint8("Subscription", 8)
	if _SubscriptionErr != nil {
		return nil, errors.Wrap(_SubscriptionErr, "Error parsing 'Subscription' field")
	}
	Subscription := _Subscription

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (magicKey)
	_magicKey, _magicKeyErr := readBuffer.ReadString("magicKey", uint32(64))
	if _magicKeyErr != nil {
		return nil, errors.Wrap(_magicKeyErr, "Error parsing 'magicKey' field")
	}
	magicKey := _magicKey

	// Optional Field (Alarmtype) (Can be skipped, if a given expression evaluates to false)
	var Alarmtype *AlarmStateType = nil
	if bool((Subscription) >= (128)) {
		if pullErr := readBuffer.PullContext("Alarmtype"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for Alarmtype")
		}
		_val, _err := AlarmStateTypeParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'Alarmtype' field")
		}
		Alarmtype = &_val
		if closeErr := readBuffer.CloseContext("Alarmtype"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for Alarmtype")
		}
	}

	// Optional Field (Reserve) (Can be skipped, if a given expression evaluates to false)
	var Reserve *uint8 = nil
	if bool((Subscription) >= (128)) {
		_val, _err := readBuffer.ReadUint8("Reserve", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'Reserve' field")
		}
		Reserve = &_val
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscription")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionMsgSubscription{
		Subscription:           Subscription,
		MagicKey:               magicKey,
		Alarmtype:              Alarmtype,
		Reserve:                Reserve,
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscription")
		}

		// Simple Field (Subscription)
		Subscription := uint8(m.GetSubscription())
		_SubscriptionErr := writeBuffer.WriteUint8("Subscription", 8, (Subscription))
		if _SubscriptionErr != nil {
			return errors.Wrap(_SubscriptionErr, "Error serializing 'Subscription' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (magicKey)
		magicKey := string(m.GetMagicKey())
		_magicKeyErr := writeBuffer.WriteString("magicKey", uint32(64), "UTF-8", (magicKey))
		if _magicKeyErr != nil {
			return errors.Wrap(_magicKeyErr, "Error serializing 'magicKey' field")
		}

		// Optional Field (Alarmtype) (Can be skipped, if the value is null)
		var Alarmtype *AlarmStateType = nil
		if m.GetAlarmtype() != nil {
			if pushErr := writeBuffer.PushContext("Alarmtype"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for Alarmtype")
			}
			Alarmtype = m.GetAlarmtype()
			_AlarmtypeErr := writeBuffer.WriteSerializable(Alarmtype)
			if popErr := writeBuffer.PopContext("Alarmtype"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for Alarmtype")
			}
			if _AlarmtypeErr != nil {
				return errors.Wrap(_AlarmtypeErr, "Error serializing 'Alarmtype' field")
			}
		}

		// Optional Field (Reserve) (Can be skipped, if the value is null)
		var Reserve *uint8 = nil
		if m.GetReserve() != nil {
			Reserve = m.GetReserve()
			_ReserveErr := writeBuffer.WriteUint8("Reserve", 8, *(Reserve))
			if _ReserveErr != nil {
				return errors.Wrap(_ReserveErr, "Error serializing 'Reserve' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscription")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) isS7PayloadUserDataItemCpuFunctionMsgSubscription() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
