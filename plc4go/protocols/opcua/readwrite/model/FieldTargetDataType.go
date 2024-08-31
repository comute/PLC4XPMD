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

// FieldTargetDataType is the corresponding interface of FieldTargetDataType
type FieldTargetDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetDataSetFieldId returns DataSetFieldId (property field)
	GetDataSetFieldId() GuidValue
	// GetReceiverIndexRange returns ReceiverIndexRange (property field)
	GetReceiverIndexRange() PascalString
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetWriteIndexRange returns WriteIndexRange (property field)
	GetWriteIndexRange() PascalString
	// GetOverrideValueHandling returns OverrideValueHandling (property field)
	GetOverrideValueHandling() OverrideValueHandling
	// GetOverrideValue returns OverrideValue (property field)
	GetOverrideValue() Variant
}

// FieldTargetDataTypeExactly can be used when we want exactly this type and not a type which fulfills FieldTargetDataType.
// This is useful for switch cases.
type FieldTargetDataTypeExactly interface {
	FieldTargetDataType
	isFieldTargetDataType() bool
}

// _FieldTargetDataType is the data-structure of this message
type _FieldTargetDataType struct {
	*_ExtensionObjectDefinition
	DataSetFieldId        GuidValue
	ReceiverIndexRange    PascalString
	TargetNodeId          NodeId
	AttributeId           uint32
	WriteIndexRange       PascalString
	OverrideValueHandling OverrideValueHandling
	OverrideValue         Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FieldTargetDataType) GetIdentifier() string {
	return "14746"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FieldTargetDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_FieldTargetDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FieldTargetDataType) GetDataSetFieldId() GuidValue {
	return m.DataSetFieldId
}

func (m *_FieldTargetDataType) GetReceiverIndexRange() PascalString {
	return m.ReceiverIndexRange
}

func (m *_FieldTargetDataType) GetTargetNodeId() NodeId {
	return m.TargetNodeId
}

func (m *_FieldTargetDataType) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_FieldTargetDataType) GetWriteIndexRange() PascalString {
	return m.WriteIndexRange
}

func (m *_FieldTargetDataType) GetOverrideValueHandling() OverrideValueHandling {
	return m.OverrideValueHandling
}

func (m *_FieldTargetDataType) GetOverrideValue() Variant {
	return m.OverrideValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFieldTargetDataType factory function for _FieldTargetDataType
func NewFieldTargetDataType(dataSetFieldId GuidValue, receiverIndexRange PascalString, targetNodeId NodeId, attributeId uint32, writeIndexRange PascalString, overrideValueHandling OverrideValueHandling, overrideValue Variant) *_FieldTargetDataType {
	_result := &_FieldTargetDataType{
		DataSetFieldId:             dataSetFieldId,
		ReceiverIndexRange:         receiverIndexRange,
		TargetNodeId:               targetNodeId,
		AttributeId:                attributeId,
		WriteIndexRange:            writeIndexRange,
		OverrideValueHandling:      overrideValueHandling,
		OverrideValue:              overrideValue,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFieldTargetDataType(structType any) FieldTargetDataType {
	if casted, ok := structType.(FieldTargetDataType); ok {
		return casted
	}
	if casted, ok := structType.(*FieldTargetDataType); ok {
		return *casted
	}
	return nil
}

func (m *_FieldTargetDataType) GetTypeName() string {
	return "FieldTargetDataType"
}

func (m *_FieldTargetDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dataSetFieldId)
	lengthInBits += m.DataSetFieldId.GetLengthInBits(ctx)

	// Simple field (receiverIndexRange)
	lengthInBits += m.ReceiverIndexRange.GetLengthInBits(ctx)

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (writeIndexRange)
	lengthInBits += m.WriteIndexRange.GetLengthInBits(ctx)

	// Simple field (overrideValueHandling)
	lengthInBits += 32

	// Simple field (overrideValue)
	lengthInBits += m.OverrideValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_FieldTargetDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FieldTargetDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (FieldTargetDataType, error) {
	return FieldTargetDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func FieldTargetDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (FieldTargetDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (FieldTargetDataType, error) {
		return FieldTargetDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func FieldTargetDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (FieldTargetDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FieldTargetDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FieldTargetDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataSetFieldId, err := ReadSimpleField[GuidValue](ctx, "dataSetFieldId", ReadComplex[GuidValue](GuidValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetFieldId' field"))
	}

	receiverIndexRange, err := ReadSimpleField[PascalString](ctx, "receiverIndexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiverIndexRange' field"))
	}

	targetNodeId, err := ReadSimpleField[NodeId](ctx, "targetNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeId' field"))
	}

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}

	writeIndexRange, err := ReadSimpleField[PascalString](ctx, "writeIndexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeIndexRange' field"))
	}

	overrideValueHandling, err := ReadEnumField[OverrideValueHandling](ctx, "overrideValueHandling", "OverrideValueHandling", ReadEnum(OverrideValueHandlingByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'overrideValueHandling' field"))
	}

	overrideValue, err := ReadSimpleField[Variant](ctx, "overrideValue", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'overrideValue' field"))
	}

	if closeErr := readBuffer.CloseContext("FieldTargetDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FieldTargetDataType")
	}

	// Create a partially initialized instance
	_child := &_FieldTargetDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		DataSetFieldId:             dataSetFieldId,
		ReceiverIndexRange:         receiverIndexRange,
		TargetNodeId:               targetNodeId,
		AttributeId:                attributeId,
		WriteIndexRange:            writeIndexRange,
		OverrideValueHandling:      overrideValueHandling,
		OverrideValue:              overrideValue,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_FieldTargetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FieldTargetDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FieldTargetDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FieldTargetDataType")
		}

		// Simple Field (dataSetFieldId)
		if pushErr := writeBuffer.PushContext("dataSetFieldId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataSetFieldId")
		}
		_dataSetFieldIdErr := writeBuffer.WriteSerializable(ctx, m.GetDataSetFieldId())
		if popErr := writeBuffer.PopContext("dataSetFieldId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataSetFieldId")
		}
		if _dataSetFieldIdErr != nil {
			return errors.Wrap(_dataSetFieldIdErr, "Error serializing 'dataSetFieldId' field")
		}

		// Simple Field (receiverIndexRange)
		if pushErr := writeBuffer.PushContext("receiverIndexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for receiverIndexRange")
		}
		_receiverIndexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetReceiverIndexRange())
		if popErr := writeBuffer.PopContext("receiverIndexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for receiverIndexRange")
		}
		if _receiverIndexRangeErr != nil {
			return errors.Wrap(_receiverIndexRangeErr, "Error serializing 'receiverIndexRange' field")
		}

		// Simple Field (targetNodeId)
		if pushErr := writeBuffer.PushContext("targetNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetNodeId")
		}
		_targetNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetTargetNodeId())
		if popErr := writeBuffer.PopContext("targetNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetNodeId")
		}
		if _targetNodeIdErr != nil {
			return errors.Wrap(_targetNodeIdErr, "Error serializing 'targetNodeId' field")
		}

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (writeIndexRange)
		if pushErr := writeBuffer.PushContext("writeIndexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writeIndexRange")
		}
		_writeIndexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetWriteIndexRange())
		if popErr := writeBuffer.PopContext("writeIndexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writeIndexRange")
		}
		if _writeIndexRangeErr != nil {
			return errors.Wrap(_writeIndexRangeErr, "Error serializing 'writeIndexRange' field")
		}

		// Simple Field (overrideValueHandling)
		if pushErr := writeBuffer.PushContext("overrideValueHandling"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for overrideValueHandling")
		}
		_overrideValueHandlingErr := writeBuffer.WriteSerializable(ctx, m.GetOverrideValueHandling())
		if popErr := writeBuffer.PopContext("overrideValueHandling"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for overrideValueHandling")
		}
		if _overrideValueHandlingErr != nil {
			return errors.Wrap(_overrideValueHandlingErr, "Error serializing 'overrideValueHandling' field")
		}

		// Simple Field (overrideValue)
		if pushErr := writeBuffer.PushContext("overrideValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for overrideValue")
		}
		_overrideValueErr := writeBuffer.WriteSerializable(ctx, m.GetOverrideValue())
		if popErr := writeBuffer.PopContext("overrideValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for overrideValue")
		}
		if _overrideValueErr != nil {
			return errors.Wrap(_overrideValueErr, "Error serializing 'overrideValue' field")
		}

		if popErr := writeBuffer.PopContext("FieldTargetDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FieldTargetDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FieldTargetDataType) isFieldTargetDataType() bool {
	return true
}

func (m *_FieldTargetDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
