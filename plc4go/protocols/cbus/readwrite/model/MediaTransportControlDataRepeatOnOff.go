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

// MediaTransportControlDataRepeatOnOff is the corresponding interface of MediaTransportControlDataRepeatOnOff
type MediaTransportControlDataRepeatOnOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetRepeatType returns RepeatType (property field)
	GetRepeatType() byte
	// GetIsOff returns IsOff (virtual field)
	GetIsOff() bool
	// GetIsRepeatCurrent returns IsRepeatCurrent (virtual field)
	GetIsRepeatCurrent() bool
	// GetIsRepeatTracks returns IsRepeatTracks (virtual field)
	GetIsRepeatTracks() bool
}

// MediaTransportControlDataRepeatOnOffExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataRepeatOnOff.
// This is useful for switch cases.
type MediaTransportControlDataRepeatOnOffExactly interface {
	MediaTransportControlDataRepeatOnOff
	isMediaTransportControlDataRepeatOnOff() bool
}

// _MediaTransportControlDataRepeatOnOff is the data-structure of this message
type _MediaTransportControlDataRepeatOnOff struct {
	*_MediaTransportControlData
	RepeatType byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataRepeatOnOff) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataRepeatOnOff) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataRepeatOnOff) GetRepeatType() byte {
	return m.RepeatType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataRepeatOnOff) GetIsOff() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRepeatType()) == (0x00)))
}

func (m *_MediaTransportControlDataRepeatOnOff) GetIsRepeatCurrent() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetRepeatType()) > (0x00))) && bool(bool((m.GetRepeatType()) <= (0xFE))))
}

func (m *_MediaTransportControlDataRepeatOnOff) GetIsRepeatTracks() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRepeatType()) >= (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataRepeatOnOff factory function for _MediaTransportControlDataRepeatOnOff
func NewMediaTransportControlDataRepeatOnOff(repeatType byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataRepeatOnOff {
	_result := &_MediaTransportControlDataRepeatOnOff{
		RepeatType:                 repeatType,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataRepeatOnOff(structType any) MediaTransportControlDataRepeatOnOff {
	if casted, ok := structType.(MediaTransportControlDataRepeatOnOff); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataRepeatOnOff); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataRepeatOnOff) GetTypeName() string {
	return "MediaTransportControlDataRepeatOnOff"
}

func (m *_MediaTransportControlDataRepeatOnOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (repeatType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataRepeatOnOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataRepeatOnOffParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataRepeatOnOff, error) {
	return MediaTransportControlDataRepeatOnOffParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataRepeatOnOffParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRepeatOnOff, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRepeatOnOff, error) {
		return MediaTransportControlDataRepeatOnOffParseWithBuffer(ctx, readBuffer)
	}
}

func MediaTransportControlDataRepeatOnOffParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataRepeatOnOff, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataRepeatOnOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataRepeatOnOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	repeatType, err := ReadSimpleField(ctx, "repeatType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'repeatType' field"))
	}

	isOff, err := ReadVirtualField[bool](ctx, "isOff", (*bool)(nil), bool((repeatType) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isOff' field"))
	}
	_ = isOff

	isRepeatCurrent, err := ReadVirtualField[bool](ctx, "isRepeatCurrent", (*bool)(nil), bool(bool((repeatType) > (0x00))) && bool(bool((repeatType) <= (0xFE))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isRepeatCurrent' field"))
	}
	_ = isRepeatCurrent

	isRepeatTracks, err := ReadVirtualField[bool](ctx, "isRepeatTracks", (*bool)(nil), bool((repeatType) >= (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isRepeatTracks' field"))
	}
	_ = isRepeatTracks

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataRepeatOnOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataRepeatOnOff")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataRepeatOnOff{
		_MediaTransportControlData: &_MediaTransportControlData{},
		RepeatType:                 repeatType,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataRepeatOnOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataRepeatOnOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataRepeatOnOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataRepeatOnOff")
		}

		// Simple Field (repeatType)
		repeatType := byte(m.GetRepeatType())
		_repeatTypeErr := /*TODO: migrate me*/ writeBuffer.WriteByte("repeatType", (repeatType))
		if _repeatTypeErr != nil {
			return errors.Wrap(_repeatTypeErr, "Error serializing 'repeatType' field")
		}
		// Virtual field
		isOff := m.GetIsOff()
		_ = isOff
		if _isOffErr := writeBuffer.WriteVirtual(ctx, "isOff", m.GetIsOff()); _isOffErr != nil {
			return errors.Wrap(_isOffErr, "Error serializing 'isOff' field")
		}
		// Virtual field
		isRepeatCurrent := m.GetIsRepeatCurrent()
		_ = isRepeatCurrent
		if _isRepeatCurrentErr := writeBuffer.WriteVirtual(ctx, "isRepeatCurrent", m.GetIsRepeatCurrent()); _isRepeatCurrentErr != nil {
			return errors.Wrap(_isRepeatCurrentErr, "Error serializing 'isRepeatCurrent' field")
		}
		// Virtual field
		isRepeatTracks := m.GetIsRepeatTracks()
		_ = isRepeatTracks
		if _isRepeatTracksErr := writeBuffer.WriteVirtual(ctx, "isRepeatTracks", m.GetIsRepeatTracks()); _isRepeatTracksErr != nil {
			return errors.Wrap(_isRepeatTracksErr, "Error serializing 'isRepeatTracks' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataRepeatOnOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataRepeatOnOff")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataRepeatOnOff) isMediaTransportControlDataRepeatOnOff() bool {
	return true
}

func (m *_MediaTransportControlDataRepeatOnOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
