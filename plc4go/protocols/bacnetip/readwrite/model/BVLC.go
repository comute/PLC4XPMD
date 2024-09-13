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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BVLC_BACNETTYPE uint8 = 0x81

// BVLC is the corresponding interface of BVLC
type BVLC interface {
	BVLCContract
	BVLCRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBVLC is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLC()
}

// BVLCContract provides a set of functions which can be overwritten by a sub struct
type BVLCContract interface {
	// GetBvlcPayloadLength returns BvlcPayloadLength (virtual field)
	GetBvlcPayloadLength() uint16
	// IsBVLC is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLC()
}

// BVLCRequirements provides a set of functions which need to be implemented by a sub struct
type BVLCRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetBvlcFunction returns BvlcFunction (discriminator field)
	GetBvlcFunction() uint8
}

// _BVLC is the data-structure of this message
type _BVLC struct {
	_SubType BVLC
}

var _ BVLCContract = (*_BVLC)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BVLC) GetBvlcPayloadLength() uint16 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint16(uint16(uint16(m.GetLengthInBytes(ctx))) - uint16(uint16(4)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_BVLC) GetBacnetType() uint8 {
	return BVLC_BACNETTYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLC factory function for _BVLC
func NewBVLC() *_BVLC {
	return &_BVLC{}
}

// Deprecated: use the interface for direct cast
func CastBVLC(structType any) BVLC {
	if casted, ok := structType.(BVLC); ok {
		return casted
	}
	if casted, ok := structType.(*BVLC); ok {
		return *casted
	}
	return nil
}

func (m *_BVLC) GetTypeName() string {
	return "BVLC"
}

func (m *_BVLC) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetType)
	lengthInBits += 8
	// Discriminator Field (bvlcFunction)
	lengthInBits += 8

	// Implicit Field (bvlcLength)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BVLC) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BVLCParse[T BVLC](ctx context.Context, theBytes []byte) (T, error) {
	return BVLCParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func BVLCParseWithBufferProducer[T BVLC]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BVLCParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BVLCParseWithBuffer[T BVLC](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BVLC{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BVLC) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bVLC BVLC, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLC"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLC")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetType, err := ReadConstField[uint8](ctx, "bacnetType", ReadUnsignedByte(readBuffer, uint8(8)), BVLC_BACNETTYPE, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetType' field"))
	}
	_ = bacnetType

	bvlcFunction, err := ReadDiscriminatorField[uint8](ctx, "bvlcFunction", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bvlcFunction' field"))
	}

	bvlcLength, err := ReadImplicitField[uint16](ctx, "bvlcLength", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bvlcLength' field"))
	}
	_ = bvlcLength

	bvlcPayloadLength, err := ReadVirtualField[uint16](ctx, "bvlcPayloadLength", (*uint16)(nil), uint16(bvlcLength)-uint16(uint16(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bvlcPayloadLength' field"))
	}
	_ = bvlcPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BVLC
	switch {
	case bvlcFunction == 0x00: // BVLCResult
		if _child, err = (&_BVLCResult{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCResult for type-switch of BVLC")
		}
	case bvlcFunction == 0x01: // BVLCWriteBroadcastDistributionTable
		if _child, err = (&_BVLCWriteBroadcastDistributionTable{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCWriteBroadcastDistributionTable for type-switch of BVLC")
		}
	case bvlcFunction == 0x02: // BVLCReadBroadcastDistributionTable
		if _child, err = (&_BVLCReadBroadcastDistributionTable{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCReadBroadcastDistributionTable for type-switch of BVLC")
		}
	case bvlcFunction == 0x03: // BVLCReadBroadcastDistributionTableAck
		if _child, err = (&_BVLCReadBroadcastDistributionTableAck{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCReadBroadcastDistributionTableAck for type-switch of BVLC")
		}
	case bvlcFunction == 0x04: // BVLCForwardedNPDU
		if _child, err = (&_BVLCForwardedNPDU{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCForwardedNPDU for type-switch of BVLC")
		}
	case bvlcFunction == 0x05: // BVLCRegisterForeignDevice
		if _child, err = (&_BVLCRegisterForeignDevice{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCRegisterForeignDevice for type-switch of BVLC")
		}
	case bvlcFunction == 0x06: // BVLCReadForeignDeviceTable
		if _child, err = (&_BVLCReadForeignDeviceTable{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCReadForeignDeviceTable for type-switch of BVLC")
		}
	case bvlcFunction == 0x07: // BVLCReadForeignDeviceTableAck
		if _child, err = (&_BVLCReadForeignDeviceTableAck{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCReadForeignDeviceTableAck for type-switch of BVLC")
		}
	case bvlcFunction == 0x08: // BVLCDeleteForeignDeviceTableEntry
		if _child, err = (&_BVLCDeleteForeignDeviceTableEntry{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCDeleteForeignDeviceTableEntry for type-switch of BVLC")
		}
	case bvlcFunction == 0x09: // BVLCDistributeBroadcastToNetwork
		if _child, err = (&_BVLCDistributeBroadcastToNetwork{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCDistributeBroadcastToNetwork for type-switch of BVLC")
		}
	case bvlcFunction == 0x0A: // BVLCOriginalUnicastNPDU
		if _child, err = (&_BVLCOriginalUnicastNPDU{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCOriginalUnicastNPDU for type-switch of BVLC")
		}
	case bvlcFunction == 0x0B: // BVLCOriginalBroadcastNPDU
		if _child, err = (&_BVLCOriginalBroadcastNPDU{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCOriginalBroadcastNPDU for type-switch of BVLC")
		}
	case bvlcFunction == 0x0C: // BVLCSecureBVLL
		if _child, err = (&_BVLCSecureBVLL{}).parse(ctx, readBuffer, m, bvlcPayloadLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BVLCSecureBVLL for type-switch of BVLC")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [bvlcFunction=%v]", bvlcFunction)
	}

	if closeErr := readBuffer.CloseContext("BVLC"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLC")
	}

	return _child, nil
}

func (pm *_BVLC) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BVLC, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BVLC"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLC")
	}

	if err := WriteConstField(ctx, "bacnetType", BVLC_BACNETTYPE, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'bacnetType' field")
	}

	if err := WriteDiscriminatorField(ctx, "bvlcFunction", m.GetBvlcFunction(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'bvlcFunction' field")
	}
	bvlcLength := uint16(uint16(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "bvlcLength", bvlcLength, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'bvlcLength' field")
	}
	// Virtual field
	bvlcPayloadLength := m.GetBvlcPayloadLength()
	_ = bvlcPayloadLength
	if _bvlcPayloadLengthErr := writeBuffer.WriteVirtual(ctx, "bvlcPayloadLength", m.GetBvlcPayloadLength()); _bvlcPayloadLengthErr != nil {
		return errors.Wrap(_bvlcPayloadLengthErr, "Error serializing 'bvlcPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BVLC"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLC")
	}
	return nil
}

func (m *_BVLC) IsBVLC() {}
