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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUUnknown is the corresponding interface of APDUUnknown
type APDUUnknown interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetUnknownTypeRest returns UnknownTypeRest (property field)
	GetUnknownTypeRest() uint8
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
}

// APDUUnknownExactly can be used when we want exactly this type and not a type which fulfills APDUUnknown.
// This is useful for switch cases.
type APDUUnknownExactly interface {
	APDUUnknown
	isAPDUUnknown() bool
}

// _APDUUnknown is the data-structure of this message
type _APDUUnknown struct {
	*_APDU
	UnknownTypeRest uint8
	UnknownBytes    []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUUnknown) GetApduType() ApduType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUUnknown) InitializeParent(parent APDU) {}

func (m *_APDUUnknown) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUUnknown) GetUnknownTypeRest() uint8 {
	return m.UnknownTypeRest
}

func (m *_APDUUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUUnknown factory function for _APDUUnknown
func NewAPDUUnknown(unknownTypeRest uint8, unknownBytes []byte, apduLength uint16) *_APDUUnknown {
	_result := &_APDUUnknown{
		UnknownTypeRest: unknownTypeRest,
		UnknownBytes:    unknownBytes,
		_APDU:           NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUUnknown(structType interface{}) APDUUnknown {
	if casted, ok := structType.(APDUUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*APDUUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_APDUUnknown) GetTypeName() string {
	return "APDUUnknown"
}

func (m *_APDUUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unknownTypeRest)
	lengthInBits += 4

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_APDUUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUUnknownParse(theBytes []byte, apduLength uint16) (APDUUnknown, error) {
	return APDUUnknownParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUUnknownParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unknownTypeRest)
	_unknownTypeRest, _unknownTypeRestErr := readBuffer.ReadUint8("unknownTypeRest", 4)
	if _unknownTypeRestErr != nil {
		return nil, errors.Wrap(_unknownTypeRestErr, "Error parsing 'unknownTypeRest' field of APDUUnknown")
	}
	unknownTypeRest := _unknownTypeRest
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(utils.InlineIf((bool((apduLength) > (0))), func() interface{} { return uint16(apduLength) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field of APDUUnknown")
	}

	if closeErr := readBuffer.CloseContext("APDUUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUUnknown")
	}

	// Create a partially initialized instance
	_child := &_APDUUnknown{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		UnknownTypeRest: unknownTypeRest,
		UnknownBytes:    unknownBytes,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUUnknown")
		}

		// Simple Field (unknownTypeRest)
		unknownTypeRest := uint8(m.GetUnknownTypeRest())
		_unknownTypeRestErr := writeBuffer.WriteUint8("unknownTypeRest", 4, (unknownTypeRest))
		if _unknownTypeRestErr != nil {
			return errors.Wrap(_unknownTypeRestErr, "Error serializing 'unknownTypeRest' field")
		}

		// Array Field (unknownBytes)
		// Byte Array field (unknownBytes)
		if err := writeBuffer.WriteByteArray("unknownBytes", m.GetUnknownBytes()); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUUnknown")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUUnknown) isAPDUUnknown() bool {
	return true
}

func (m *_APDUUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
