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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestWriteGroup is the corresponding interface of BACnetUnconfirmedServiceRequestWriteGroup
type BACnetUnconfirmedServiceRequestWriteGroup interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetGroupNumber returns GroupNumber (property field)
	GetGroupNumber() BACnetContextTagUnsignedInteger
	// GetWritePriority returns WritePriority (property field)
	GetWritePriority() BACnetContextTagUnsignedInteger
	// GetChangeList returns ChangeList (property field)
	GetChangeList() BACnetGroupChannelValueList
	// GetInhibitDelay returns InhibitDelay (property field)
	GetInhibitDelay() BACnetContextTagUnsignedInteger
}

// BACnetUnconfirmedServiceRequestWriteGroupExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWriteGroup.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWriteGroupExactly interface {
	BACnetUnconfirmedServiceRequestWriteGroup
	isBACnetUnconfirmedServiceRequestWriteGroup() bool
}

// _BACnetUnconfirmedServiceRequestWriteGroup is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWriteGroup struct {
	*_BACnetUnconfirmedServiceRequest
	GroupNumber   BACnetContextTagUnsignedInteger
	WritePriority BACnetContextTagUnsignedInteger
	ChangeList    BACnetGroupChannelValueList
	InhibitDelay  BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_WRITE_GROUP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetGroupNumber() BACnetContextTagUnsignedInteger {
	return m.GroupNumber
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetWritePriority() BACnetContextTagUnsignedInteger {
	return m.WritePriority
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetChangeList() BACnetGroupChannelValueList {
	return m.ChangeList
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetInhibitDelay() BACnetContextTagUnsignedInteger {
	return m.InhibitDelay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWriteGroup factory function for _BACnetUnconfirmedServiceRequestWriteGroup
func NewBACnetUnconfirmedServiceRequestWriteGroup(groupNumber BACnetContextTagUnsignedInteger, writePriority BACnetContextTagUnsignedInteger, changeList BACnetGroupChannelValueList, inhibitDelay BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestWriteGroup {
	_result := &_BACnetUnconfirmedServiceRequestWriteGroup{
		GroupNumber:                      groupNumber,
		WritePriority:                    writePriority,
		ChangeList:                       changeList,
		InhibitDelay:                     inhibitDelay,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWriteGroup(structType interface{}) BACnetUnconfirmedServiceRequestWriteGroup {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWriteGroup); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWriteGroup); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWriteGroup"
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (groupNumber)
	lengthInBits += m.GroupNumber.GetLengthInBits(ctx)

	// Simple field (writePriority)
	lengthInBits += m.WritePriority.GetLengthInBits(ctx)

	// Simple field (changeList)
	lengthInBits += m.ChangeList.GetLengthInBits(ctx)

	// Optional Field (inhibitDelay)
	if m.InhibitDelay != nil {
		lengthInBits += m.InhibitDelay.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestWriteGroupParse(theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWriteGroup, error) {
	return BACnetUnconfirmedServiceRequestWriteGroupParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestWriteGroupParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWriteGroup, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWriteGroup"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWriteGroup")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupNumber)
	if pullErr := readBuffer.PullContext("groupNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupNumber")
	}
	_groupNumber, _groupNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _groupNumberErr != nil {
		return nil, errors.Wrap(_groupNumberErr, "Error parsing 'groupNumber' field of BACnetUnconfirmedServiceRequestWriteGroup")
	}
	groupNumber := _groupNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("groupNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupNumber")
	}

	// Simple Field (writePriority)
	if pullErr := readBuffer.PullContext("writePriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writePriority")
	}
	_writePriority, _writePriorityErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _writePriorityErr != nil {
		return nil, errors.Wrap(_writePriorityErr, "Error parsing 'writePriority' field of BACnetUnconfirmedServiceRequestWriteGroup")
	}
	writePriority := _writePriority.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("writePriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writePriority")
	}

	// Simple Field (changeList)
	if pullErr := readBuffer.PullContext("changeList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changeList")
	}
	_changeList, _changeListErr := BACnetGroupChannelValueListParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _changeListErr != nil {
		return nil, errors.Wrap(_changeListErr, "Error parsing 'changeList' field of BACnetUnconfirmedServiceRequestWriteGroup")
	}
	changeList := _changeList.(BACnetGroupChannelValueList)
	if closeErr := readBuffer.CloseContext("changeList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changeList")
	}

	// Optional Field (inhibitDelay) (Can be skipped, if a given expression evaluates to false)
	var inhibitDelay BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("inhibitDelay"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for inhibitDelay")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'inhibitDelay' field of BACnetUnconfirmedServiceRequestWriteGroup")
		default:
			inhibitDelay = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("inhibitDelay"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for inhibitDelay")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWriteGroup"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWriteGroup")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWriteGroup{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		GroupNumber:   groupNumber,
		WritePriority: writePriority,
		ChangeList:    changeList,
		InhibitDelay:  inhibitDelay,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWriteGroup"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWriteGroup")
		}

		// Simple Field (groupNumber)
		if pushErr := writeBuffer.PushContext("groupNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for groupNumber")
		}
		_groupNumberErr := writeBuffer.WriteSerializable(ctx, m.GetGroupNumber())
		if popErr := writeBuffer.PopContext("groupNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for groupNumber")
		}
		if _groupNumberErr != nil {
			return errors.Wrap(_groupNumberErr, "Error serializing 'groupNumber' field")
		}

		// Simple Field (writePriority)
		if pushErr := writeBuffer.PushContext("writePriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writePriority")
		}
		_writePriorityErr := writeBuffer.WriteSerializable(ctx, m.GetWritePriority())
		if popErr := writeBuffer.PopContext("writePriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writePriority")
		}
		if _writePriorityErr != nil {
			return errors.Wrap(_writePriorityErr, "Error serializing 'writePriority' field")
		}

		// Simple Field (changeList)
		if pushErr := writeBuffer.PushContext("changeList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changeList")
		}
		_changeListErr := writeBuffer.WriteSerializable(ctx, m.GetChangeList())
		if popErr := writeBuffer.PopContext("changeList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changeList")
		}
		if _changeListErr != nil {
			return errors.Wrap(_changeListErr, "Error serializing 'changeList' field")
		}

		// Optional Field (inhibitDelay) (Can be skipped, if the value is null)
		var inhibitDelay BACnetContextTagUnsignedInteger = nil
		if m.GetInhibitDelay() != nil {
			if pushErr := writeBuffer.PushContext("inhibitDelay"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for inhibitDelay")
			}
			inhibitDelay = m.GetInhibitDelay()
			_inhibitDelayErr := writeBuffer.WriteSerializable(ctx, inhibitDelay)
			if popErr := writeBuffer.PopContext("inhibitDelay"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for inhibitDelay")
			}
			if _inhibitDelayErr != nil {
				return errors.Wrap(_inhibitDelayErr, "Error serializing 'inhibitDelay' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWriteGroup"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWriteGroup")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) isBACnetUnconfirmedServiceRequestWriteGroup() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWriteGroup) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
