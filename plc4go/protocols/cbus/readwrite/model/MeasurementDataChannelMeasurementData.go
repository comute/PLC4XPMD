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

// MeasurementDataChannelMeasurementData is the corresponding interface of MeasurementDataChannelMeasurementData
type MeasurementDataChannelMeasurementData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeasurementData
	// GetDeviceId returns DeviceId (property field)
	GetDeviceId() uint8
	// GetChannel returns Channel (property field)
	GetChannel() uint8
	// GetUnits returns Units (property field)
	GetUnits() MeasurementUnits
	// GetMultiplier returns Multiplier (property field)
	GetMultiplier() int8
	// GetMsb returns Msb (property field)
	GetMsb() uint8
	// GetLsb returns Lsb (property field)
	GetLsb() uint8
	// GetRawValue returns RawValue (virtual field)
	GetRawValue() uint16
	// GetValue returns Value (virtual field)
	GetValue() float64
	// IsMeasurementDataChannelMeasurementData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMeasurementDataChannelMeasurementData()
}

// _MeasurementDataChannelMeasurementData is the data-structure of this message
type _MeasurementDataChannelMeasurementData struct {
	MeasurementDataContract
	DeviceId   uint8
	Channel    uint8
	Units      MeasurementUnits
	Multiplier int8
	Msb        uint8
	Lsb        uint8
}

var _ MeasurementDataChannelMeasurementData = (*_MeasurementDataChannelMeasurementData)(nil)
var _ MeasurementDataRequirements = (*_MeasurementDataChannelMeasurementData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeasurementDataChannelMeasurementData) GetParent() MeasurementDataContract {
	return m.MeasurementDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeasurementDataChannelMeasurementData) GetDeviceId() uint8 {
	return m.DeviceId
}

func (m *_MeasurementDataChannelMeasurementData) GetChannel() uint8 {
	return m.Channel
}

func (m *_MeasurementDataChannelMeasurementData) GetUnits() MeasurementUnits {
	return m.Units
}

func (m *_MeasurementDataChannelMeasurementData) GetMultiplier() int8 {
	return m.Multiplier
}

func (m *_MeasurementDataChannelMeasurementData) GetMsb() uint8 {
	return m.Msb
}

func (m *_MeasurementDataChannelMeasurementData) GetLsb() uint8 {
	return m.Lsb
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MeasurementDataChannelMeasurementData) GetRawValue() uint16 {
	ctx := context.Background()
	_ = ctx
	return uint16(m.GetMsb()<<uint16(8) | m.GetLsb())
}

func (m *_MeasurementDataChannelMeasurementData) GetValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(float64(float64(m.GetRawValue())*float64(m.GetMultiplier())) * float64(float64(10)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeasurementDataChannelMeasurementData factory function for _MeasurementDataChannelMeasurementData
func NewMeasurementDataChannelMeasurementData(deviceId uint8, channel uint8, units MeasurementUnits, multiplier int8, msb uint8, lsb uint8, commandTypeContainer MeasurementCommandTypeContainer) *_MeasurementDataChannelMeasurementData {
	_result := &_MeasurementDataChannelMeasurementData{
		MeasurementDataContract: NewMeasurementData(commandTypeContainer),
		DeviceId:                deviceId,
		Channel:                 channel,
		Units:                   units,
		Multiplier:              multiplier,
		Msb:                     msb,
		Lsb:                     lsb,
	}
	_result.MeasurementDataContract.(*_MeasurementData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeasurementDataChannelMeasurementData(structType any) MeasurementDataChannelMeasurementData {
	if casted, ok := structType.(MeasurementDataChannelMeasurementData); ok {
		return casted
	}
	if casted, ok := structType.(*MeasurementDataChannelMeasurementData); ok {
		return *casted
	}
	return nil
}

func (m *_MeasurementDataChannelMeasurementData) GetTypeName() string {
	return "MeasurementDataChannelMeasurementData"
}

func (m *_MeasurementDataChannelMeasurementData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MeasurementDataContract.(*_MeasurementData).getLengthInBits(ctx))

	// Simple field (deviceId)
	lengthInBits += 8

	// Simple field (channel)
	lengthInBits += 8

	// Simple field (units)
	lengthInBits += 8

	// Simple field (multiplier)
	lengthInBits += 8

	// Simple field (msb)
	lengthInBits += 8

	// Simple field (lsb)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MeasurementDataChannelMeasurementData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MeasurementDataChannelMeasurementData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MeasurementData) (__measurementDataChannelMeasurementData MeasurementDataChannelMeasurementData, err error) {
	m.MeasurementDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeasurementDataChannelMeasurementData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeasurementDataChannelMeasurementData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceId, err := ReadSimpleField(ctx, "deviceId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceId' field"))
	}
	m.DeviceId = deviceId

	channel, err := ReadSimpleField(ctx, "channel", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channel' field"))
	}
	m.Channel = channel

	units, err := ReadEnumField[MeasurementUnits](ctx, "units", "MeasurementUnits", ReadEnum(MeasurementUnitsByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	multiplier, err := ReadSimpleField(ctx, "multiplier", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'multiplier' field"))
	}
	m.Multiplier = multiplier

	msb, err := ReadSimpleField(ctx, "msb", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'msb' field"))
	}
	m.Msb = msb

	lsb, err := ReadSimpleField(ctx, "lsb", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lsb' field"))
	}
	m.Lsb = lsb

	rawValue, err := ReadVirtualField[uint16](ctx, "rawValue", (*uint16)(nil), msb<<uint16(8)|lsb)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawValue' field"))
	}
	_ = rawValue

	value, err := ReadVirtualField[float64](ctx, "value", (*float64)(nil), float64(float64(rawValue)*float64(multiplier))*float64(float64(10)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	_ = value

	if closeErr := readBuffer.CloseContext("MeasurementDataChannelMeasurementData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeasurementDataChannelMeasurementData")
	}

	return m, nil
}

func (m *_MeasurementDataChannelMeasurementData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeasurementDataChannelMeasurementData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeasurementDataChannelMeasurementData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeasurementDataChannelMeasurementData")
		}

		if err := WriteSimpleField[uint8](ctx, "deviceId", m.GetDeviceId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "channel", m.GetChannel(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'channel' field")
		}

		if err := WriteSimpleEnumField[MeasurementUnits](ctx, "units", "MeasurementUnits", m.GetUnits(), WriteEnum[MeasurementUnits, uint8](MeasurementUnits.GetValue, MeasurementUnits.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}

		if err := WriteSimpleField[int8](ctx, "multiplier", m.GetMultiplier(), WriteSignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'multiplier' field")
		}

		if err := WriteSimpleField[uint8](ctx, "msb", m.GetMsb(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'msb' field")
		}

		if err := WriteSimpleField[uint8](ctx, "lsb", m.GetLsb(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'lsb' field")
		}
		// Virtual field
		rawValue := m.GetRawValue()
		_ = rawValue
		if _rawValueErr := writeBuffer.WriteVirtual(ctx, "rawValue", m.GetRawValue()); _rawValueErr != nil {
			return errors.Wrap(_rawValueErr, "Error serializing 'rawValue' field")
		}
		// Virtual field
		value := m.GetValue()
		_ = value
		if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("MeasurementDataChannelMeasurementData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeasurementDataChannelMeasurementData")
		}
		return nil
	}
	return m.MeasurementDataContract.(*_MeasurementData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeasurementDataChannelMeasurementData) IsMeasurementDataChannelMeasurementData() {}

func (m *_MeasurementDataChannelMeasurementData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
