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

// AirConditioningDataHvacScheduleEntry is the corresponding interface of AirConditioningDataHvacScheduleEntry
type AirConditioningDataHvacScheduleEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetEntry returns Entry (property field)
	GetEntry() uint8
	// GetFormat returns Format (property field)
	GetFormat() byte
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// GetStartTime returns StartTime (property field)
	GetStartTime() HVACStartTime
	// GetLevel returns Level (property field)
	GetLevel() HVACTemperature
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
}

// AirConditioningDataHvacScheduleEntryExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataHvacScheduleEntry.
// This is useful for switch cases.
type AirConditioningDataHvacScheduleEntryExactly interface {
	AirConditioningDataHvacScheduleEntry
	isAirConditioningDataHvacScheduleEntry() bool
}

// _AirConditioningDataHvacScheduleEntry is the data-structure of this message
type _AirConditioningDataHvacScheduleEntry struct {
	*_AirConditioningData
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Entry            uint8
	Format           byte
	HvacModeAndFlags HVACModeAndFlags
	StartTime        HVACStartTime
	Level            HVACTemperature
	RawLevel         HVACRawLevels
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataHvacScheduleEntry) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataHvacScheduleEntry) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataHvacScheduleEntry) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataHvacScheduleEntry) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataHvacScheduleEntry) GetEntry() uint8 {
	return m.Entry
}

func (m *_AirConditioningDataHvacScheduleEntry) GetFormat() byte {
	return m.Format
}

func (m *_AirConditioningDataHvacScheduleEntry) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

func (m *_AirConditioningDataHvacScheduleEntry) GetStartTime() HVACStartTime {
	return m.StartTime
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLevel() HVACTemperature {
	return m.Level
}

func (m *_AirConditioningDataHvacScheduleEntry) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataHvacScheduleEntry factory function for _AirConditioningDataHvacScheduleEntry
func NewAirConditioningDataHvacScheduleEntry(zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, hvacModeAndFlags HVACModeAndFlags, startTime HVACStartTime, level HVACTemperature, rawLevel HVACRawLevels, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataHvacScheduleEntry {
	_result := &_AirConditioningDataHvacScheduleEntry{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Entry:                entry,
		Format:               format,
		HvacModeAndFlags:     hvacModeAndFlags,
		StartTime:            startTime,
		Level:                level,
		RawLevel:             rawLevel,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataHvacScheduleEntry(structType any) AirConditioningDataHvacScheduleEntry {
	if casted, ok := structType.(AirConditioningDataHvacScheduleEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataHvacScheduleEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataHvacScheduleEntry) GetTypeName() string {
	return "AirConditioningDataHvacScheduleEntry"
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (entry)
	lengthInBits += 8

	// Simple field (format)
	lengthInBits += 8

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataHvacScheduleEntryParse(ctx context.Context, theBytes []byte) (AirConditioningDataHvacScheduleEntry, error) {
	return AirConditioningDataHvacScheduleEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataHvacScheduleEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataHvacScheduleEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningDataHvacScheduleEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataHvacScheduleEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := /*TODO: migrate me*/ readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataHvacScheduleEntry")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataHvacScheduleEntry")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (entry)
	_entry, _entryErr := /*TODO: migrate me*/ readBuffer.ReadUint8("entry", 8)
	if _entryErr != nil {
		return nil, errors.Wrap(_entryErr, "Error parsing 'entry' field of AirConditioningDataHvacScheduleEntry")
	}
	entry := _entry

	// Simple Field (format)
	_format, _formatErr := /*TODO: migrate me*/ readBuffer.ReadByte("format")
	if _formatErr != nil {
		return nil, errors.Wrap(_formatErr, "Error parsing 'format' field of AirConditioningDataHvacScheduleEntry")
	}
	format := _format

	// Simple Field (hvacModeAndFlags)
	if pullErr := readBuffer.PullContext("hvacModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacModeAndFlags")
	}
	_hvacModeAndFlags, _hvacModeAndFlagsErr := HVACModeAndFlagsParseWithBuffer(ctx, readBuffer)
	if _hvacModeAndFlagsErr != nil {
		return nil, errors.Wrap(_hvacModeAndFlagsErr, "Error parsing 'hvacModeAndFlags' field of AirConditioningDataHvacScheduleEntry")
	}
	hvacModeAndFlags := _hvacModeAndFlags.(HVACModeAndFlags)
	if closeErr := readBuffer.CloseContext("hvacModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacModeAndFlags")
	}

	// Simple Field (startTime)
	if pullErr := readBuffer.PullContext("startTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for startTime")
	}
	_startTime, _startTimeErr := HVACStartTimeParseWithBuffer(ctx, readBuffer)
	if _startTimeErr != nil {
		return nil, errors.Wrap(_startTimeErr, "Error parsing 'startTime' field of AirConditioningDataHvacScheduleEntry")
	}
	startTime := _startTime.(HVACStartTime)
	if closeErr := readBuffer.CloseContext("startTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for startTime")
	}

	_level, err := ReadOptionalField[HVACTemperature](ctx, "level", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelTemperature())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	var level HVACTemperature
	if _level != nil {
		level = *_level
	}

	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	var rawLevel HVACRawLevels
	if _rawLevel != nil {
		rawLevel = *_rawLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataHvacScheduleEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataHvacScheduleEntry")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataHvacScheduleEntry{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Entry:                entry,
		Format:               format,
		HvacModeAndFlags:     hvacModeAndFlags,
		StartTime:            startTime,
		Level:                level,
		RawLevel:             rawLevel,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataHvacScheduleEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataHvacScheduleEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataHvacScheduleEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataHvacScheduleEntry")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := /*TODO: migrate me*/ writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (entry)
		entry := uint8(m.GetEntry())
		_entryErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("entry", 8, uint8((entry)))
		if _entryErr != nil {
			return errors.Wrap(_entryErr, "Error serializing 'entry' field")
		}

		// Simple Field (format)
		format := byte(m.GetFormat())
		_formatErr := /*TODO: migrate me*/ writeBuffer.WriteByte("format", (format))
		if _formatErr != nil {
			return errors.Wrap(_formatErr, "Error serializing 'format' field")
		}

		// Simple Field (hvacModeAndFlags)
		if pushErr := writeBuffer.PushContext("hvacModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hvacModeAndFlags")
		}
		_hvacModeAndFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetHvacModeAndFlags())
		if popErr := writeBuffer.PopContext("hvacModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hvacModeAndFlags")
		}
		if _hvacModeAndFlagsErr != nil {
			return errors.Wrap(_hvacModeAndFlagsErr, "Error serializing 'hvacModeAndFlags' field")
		}

		// Simple Field (startTime)
		if pushErr := writeBuffer.PushContext("startTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for startTime")
		}
		_startTimeErr := writeBuffer.WriteSerializable(ctx, m.GetStartTime())
		if popErr := writeBuffer.PopContext("startTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for startTime")
		}
		if _startTimeErr != nil {
			return errors.Wrap(_startTimeErr, "Error serializing 'startTime' field")
		}

		// Optional Field (level) (Can be skipped, if the value is null)
		var level HVACTemperature = nil
		if m.GetLevel() != nil {
			if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for level")
			}
			level = m.GetLevel()
			_levelErr := writeBuffer.WriteSerializable(ctx, level)
			if popErr := writeBuffer.PopContext("level"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for level")
			}
			if _levelErr != nil {
				return errors.Wrap(_levelErr, "Error serializing 'level' field")
			}
		}

		// Optional Field (rawLevel) (Can be skipped, if the value is null)
		var rawLevel HVACRawLevels = nil
		if m.GetRawLevel() != nil {
			if pushErr := writeBuffer.PushContext("rawLevel"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for rawLevel")
			}
			rawLevel = m.GetRawLevel()
			_rawLevelErr := writeBuffer.WriteSerializable(ctx, rawLevel)
			if popErr := writeBuffer.PopContext("rawLevel"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for rawLevel")
			}
			if _rawLevelErr != nil {
				return errors.Wrap(_rawLevelErr, "Error serializing 'rawLevel' field")
			}
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataHvacScheduleEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataHvacScheduleEntry")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataHvacScheduleEntry) isAirConditioningDataHvacScheduleEntry() bool {
	return true
}

func (m *_AirConditioningDataHvacScheduleEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
