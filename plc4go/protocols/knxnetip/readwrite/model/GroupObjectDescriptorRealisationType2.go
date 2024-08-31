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

// GroupObjectDescriptorRealisationType2 is the corresponding interface of GroupObjectDescriptorRealisationType2
type GroupObjectDescriptorRealisationType2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDataPointer returns DataPointer (property field)
	GetDataPointer() uint8
	// GetUpdateEnable returns UpdateEnable (property field)
	GetUpdateEnable() bool
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
}

// GroupObjectDescriptorRealisationType2Exactly can be used when we want exactly this type and not a type which fulfills GroupObjectDescriptorRealisationType2.
// This is useful for switch cases.
type GroupObjectDescriptorRealisationType2Exactly interface {
	GroupObjectDescriptorRealisationType2
	isGroupObjectDescriptorRealisationType2() bool
}

// _GroupObjectDescriptorRealisationType2 is the data-structure of this message
type _GroupObjectDescriptorRealisationType2 struct {
	DataPointer           uint8
	UpdateEnable          bool
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationType2) GetDataPointer() uint8 {
	return m.DataPointer
}

func (m *_GroupObjectDescriptorRealisationType2) GetUpdateEnable() bool {
	return m.UpdateEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationType2) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationType2) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGroupObjectDescriptorRealisationType2 factory function for _GroupObjectDescriptorRealisationType2
func NewGroupObjectDescriptorRealisationType2(dataPointer uint8, updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationType2 {
	return &_GroupObjectDescriptorRealisationType2{DataPointer: dataPointer, UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType2(structType any) GroupObjectDescriptorRealisationType2 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType2); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType2); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType2) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType2"
}

func (m *_GroupObjectDescriptorRealisationType2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataPointer)
	lengthInBits += 8

	// Simple field (updateEnable)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupObjectDescriptorRealisationType2Parse(ctx context.Context, theBytes []byte) (GroupObjectDescriptorRealisationType2, error) {
	return GroupObjectDescriptorRealisationType2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupObjectDescriptorRealisationType2ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType2, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType2, error) {
		return GroupObjectDescriptorRealisationType2ParseWithBuffer(ctx, readBuffer)
	}
}

func GroupObjectDescriptorRealisationType2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataPointer, err := ReadSimpleField(ctx, "dataPointer", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataPointer' field"))
	}

	updateEnable, err := ReadSimpleField(ctx, "updateEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateEnable' field"))
	}

	transmitEnable, err := ReadSimpleField(ctx, "transmitEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmitEnable' field"))
	}

	segmentSelectorEnable, err := ReadSimpleField(ctx, "segmentSelectorEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentSelectorEnable' field"))
	}

	writeEnable, err := ReadSimpleField(ctx, "writeEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeEnable' field"))
	}

	readEnable, err := ReadSimpleField(ctx, "readEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readEnable' field"))
	}

	communicationEnable, err := ReadSimpleField(ctx, "communicationEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationEnable' field"))
	}

	priority, err := ReadEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", ReadEnum(CEMIPriorityByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	valueType, err := ReadEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", ReadEnum(ComObjectValueTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueType' field"))
	}

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType2")
	}

	// Create the instance
	return &_GroupObjectDescriptorRealisationType2{
		DataPointer:           dataPointer,
		UpdateEnable:          updateEnable,
		TransmitEnable:        transmitEnable,
		SegmentSelectorEnable: segmentSelectorEnable,
		WriteEnable:           writeEnable,
		ReadEnable:            readEnable,
		CommunicationEnable:   communicationEnable,
		Priority:              priority,
		ValueType:             valueType,
	}, nil
}

func (m *_GroupObjectDescriptorRealisationType2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GroupObjectDescriptorRealisationType2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType2")
	}

	// Simple Field (dataPointer)
	dataPointer := uint8(m.GetDataPointer())
	_dataPointerErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("dataPointer", 8, uint8((dataPointer)))
	if _dataPointerErr != nil {
		return errors.Wrap(_dataPointerErr, "Error serializing 'dataPointer' field")
	}

	// Simple Field (updateEnable)
	updateEnable := bool(m.GetUpdateEnable())
	_updateEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("updateEnable", (updateEnable))
	if _updateEnableErr != nil {
		return errors.Wrap(_updateEnableErr, "Error serializing 'updateEnable' field")
	}

	// Simple Field (transmitEnable)
	transmitEnable := bool(m.GetTransmitEnable())
	_transmitEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("transmitEnable", (transmitEnable))
	if _transmitEnableErr != nil {
		return errors.Wrap(_transmitEnableErr, "Error serializing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable := bool(m.GetSegmentSelectorEnable())
	_segmentSelectorEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("segmentSelectorEnable", (segmentSelectorEnable))
	if _segmentSelectorEnableErr != nil {
		return errors.Wrap(_segmentSelectorEnableErr, "Error serializing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable := bool(m.GetWriteEnable())
	_writeEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("writeEnable", (writeEnable))
	if _writeEnableErr != nil {
		return errors.Wrap(_writeEnableErr, "Error serializing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable := bool(m.GetReadEnable())
	_readEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("readEnable", (readEnable))
	if _readEnableErr != nil {
		return errors.Wrap(_readEnableErr, "Error serializing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable := bool(m.GetCommunicationEnable())
	_communicationEnableErr := /*TODO: migrate me*/ writeBuffer.WriteBit("communicationEnable", (communicationEnable))
	if _communicationEnableErr != nil {
		return errors.Wrap(_communicationEnableErr, "Error serializing 'communicationEnable' field")
	}

	// Simple Field (priority)
	if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for priority")
	}
	_priorityErr := writeBuffer.WriteSerializable(ctx, m.GetPriority())
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priority")
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (valueType)
	if pushErr := writeBuffer.PushContext("valueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueType")
	}
	_valueTypeErr := writeBuffer.WriteSerializable(ctx, m.GetValueType())
	if popErr := writeBuffer.PopContext("valueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueType")
	}
	if _valueTypeErr != nil {
		return errors.Wrap(_valueTypeErr, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType2")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType2) isGroupObjectDescriptorRealisationType2() bool {
	return true
}

func (m *_GroupObjectDescriptorRealisationType2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
