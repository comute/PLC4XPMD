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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// TimeZoneDataType is the corresponding interface of TimeZoneDataType
type TimeZoneDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetOffset returns Offset (property field)
	GetOffset() int16
	// GetDaylightSavingInOffset returns DaylightSavingInOffset (property field)
	GetDaylightSavingInOffset() bool
}

// TimeZoneDataTypeExactly can be used when we want exactly this type and not a type which fulfills TimeZoneDataType.
// This is useful for switch cases.
type TimeZoneDataTypeExactly interface {
	TimeZoneDataType
	isTimeZoneDataType() bool
}

// _TimeZoneDataType is the data-structure of this message
type _TimeZoneDataType struct {
	*_ExtensionObjectDefinition
	Offset                 int16
	DaylightSavingInOffset bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TimeZoneDataType) GetIdentifier() string {
	return "8914"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TimeZoneDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TimeZoneDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TimeZoneDataType) GetOffset() int16 {
	return m.Offset
}

func (m *_TimeZoneDataType) GetDaylightSavingInOffset() bool {
	return m.DaylightSavingInOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTimeZoneDataType factory function for _TimeZoneDataType
func NewTimeZoneDataType(offset int16, daylightSavingInOffset bool) *_TimeZoneDataType {
	_result := &_TimeZoneDataType{
		Offset:                     offset,
		DaylightSavingInOffset:     daylightSavingInOffset,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTimeZoneDataType(structType any) TimeZoneDataType {
	if casted, ok := structType.(TimeZoneDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TimeZoneDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TimeZoneDataType) GetTypeName() string {
	return "TimeZoneDataType"
}

func (m *_TimeZoneDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (offset)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (daylightSavingInOffset)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TimeZoneDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TimeZoneDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (TimeZoneDataType, error) {
	return TimeZoneDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TimeZoneDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TimeZoneDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TimeZoneDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeZoneDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (offset)
	_offset, _offsetErr := readBuffer.ReadInt16("offset", 16)
	if _offsetErr != nil {
		return nil, errors.Wrap(_offsetErr, "Error parsing 'offset' field of TimeZoneDataType")
	}
	offset := _offset

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TimeZoneDataType")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (daylightSavingInOffset)
	_daylightSavingInOffset, _daylightSavingInOffsetErr := readBuffer.ReadBit("daylightSavingInOffset")
	if _daylightSavingInOffsetErr != nil {
		return nil, errors.Wrap(_daylightSavingInOffsetErr, "Error parsing 'daylightSavingInOffset' field of TimeZoneDataType")
	}
	daylightSavingInOffset := _daylightSavingInOffset

	if closeErr := readBuffer.CloseContext("TimeZoneDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeZoneDataType")
	}

	// Create a partially initialized instance
	_child := &_TimeZoneDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Offset:                     offset,
		DaylightSavingInOffset:     daylightSavingInOffset,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TimeZoneDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TimeZoneDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TimeZoneDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TimeZoneDataType")
		}

		// Simple Field (offset)
		offset := int16(m.GetOffset())
		_offsetErr := writeBuffer.WriteInt16("offset", 16, int16((offset)))
		if _offsetErr != nil {
			return errors.Wrap(_offsetErr, "Error serializing 'offset' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (daylightSavingInOffset)
		daylightSavingInOffset := bool(m.GetDaylightSavingInOffset())
		_daylightSavingInOffsetErr := writeBuffer.WriteBit("daylightSavingInOffset", (daylightSavingInOffset))
		if _daylightSavingInOffsetErr != nil {
			return errors.Wrap(_daylightSavingInOffsetErr, "Error serializing 'daylightSavingInOffset' field")
		}

		if popErr := writeBuffer.PopContext("TimeZoneDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TimeZoneDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TimeZoneDataType) isTimeZoneDataType() bool {
	return true
}

func (m *_TimeZoneDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
