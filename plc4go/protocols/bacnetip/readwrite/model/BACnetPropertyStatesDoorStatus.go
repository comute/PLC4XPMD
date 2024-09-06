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

// BACnetPropertyStatesDoorStatus is the corresponding interface of BACnetPropertyStatesDoorStatus
type BACnetPropertyStatesDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
	// IsBACnetPropertyStatesDoorStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesDoorStatus()
}

// _BACnetPropertyStatesDoorStatus is the data-structure of this message
type _BACnetPropertyStatesDoorStatus struct {
	BACnetPropertyStatesContract
	DoorStatus BACnetDoorStatusTagged
}

var _ BACnetPropertyStatesDoorStatus = (*_BACnetPropertyStatesDoorStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesDoorStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorStatus) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorStatus factory function for _BACnetPropertyStatesDoorStatus
func NewBACnetPropertyStatesDoorStatus(doorStatus BACnetDoorStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesDoorStatus {
	if doorStatus == nil {
		panic("doorStatus of type BACnetDoorStatusTagged for BACnetPropertyStatesDoorStatus must not be nil")
	}
	_result := &_BACnetPropertyStatesDoorStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		DoorStatus:                   doorStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorStatus(structType any) BACnetPropertyStatesDoorStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorStatus"
}

func (m *_BACnetPropertyStatesDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesDoorStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesDoorStatus BACnetPropertyStatesDoorStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorStatus, err := ReadSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", ReadComplex[BACnetDoorStatusTagged](BACnetDoorStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorStatus' field"))
	}
	m.DoorStatus = doorStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorStatus")
		}

		if err := WriteSimpleField[BACnetDoorStatusTagged](ctx, "doorStatus", m.GetDoorStatus(), WriteComplex[BACnetDoorStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorStatus) IsBACnetPropertyStatesDoorStatus() {}

func (m *_BACnetPropertyStatesDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
