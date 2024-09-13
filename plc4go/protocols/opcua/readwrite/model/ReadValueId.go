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

// ReadValueId is the corresponding interface of ReadValueId
type ReadValueId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetDataEncoding returns DataEncoding (property field)
	GetDataEncoding() QualifiedName
	// IsReadValueId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReadValueId()
}

// _ReadValueId is the data-structure of this message
type _ReadValueId struct {
	ExtensionObjectDefinitionContract
	NodeId       NodeId
	AttributeId  uint32
	IndexRange   PascalString
	DataEncoding QualifiedName
}

var _ ReadValueId = (*_ReadValueId)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReadValueId)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadValueId) GetIdentifier() string {
	return "628"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadValueId) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadValueId) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_ReadValueId) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_ReadValueId) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_ReadValueId) GetDataEncoding() QualifiedName {
	return m.DataEncoding
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReadValueId factory function for _ReadValueId
func NewReadValueId(nodeId NodeId, attributeId uint32, indexRange PascalString, dataEncoding QualifiedName) *_ReadValueId {
	if nodeId == nil {
		panic("nodeId of type NodeId for ReadValueId must not be nil")
	}
	if indexRange == nil {
		panic("indexRange of type PascalString for ReadValueId must not be nil")
	}
	if dataEncoding == nil {
		panic("dataEncoding of type QualifiedName for ReadValueId must not be nil")
	}
	_result := &_ReadValueId{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		AttributeId:                       attributeId,
		IndexRange:                        indexRange,
		DataEncoding:                      dataEncoding,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReadValueId(structType any) ReadValueId {
	if casted, ok := structType.(ReadValueId); ok {
		return casted
	}
	if casted, ok := structType.(*ReadValueId); ok {
		return *casted
	}
	return nil
}

func (m *_ReadValueId) GetTypeName() string {
	return "ReadValueId"
}

func (m *_ReadValueId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (dataEncoding)
	lengthInBits += m.DataEncoding.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReadValueId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReadValueId) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__readValueId ReadValueId, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReadValueId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadValueId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	indexRange, err := ReadSimpleField[PascalString](ctx, "indexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexRange' field"))
	}
	m.IndexRange = indexRange

	dataEncoding, err := ReadSimpleField[QualifiedName](ctx, "dataEncoding", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataEncoding' field"))
	}
	m.DataEncoding = dataEncoding

	if closeErr := readBuffer.CloseContext("ReadValueId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadValueId")
	}

	return m, nil
}

func (m *_ReadValueId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadValueId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadValueId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadValueId")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "indexRange", m.GetIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexRange' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "dataEncoding", m.GetDataEncoding(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataEncoding' field")
		}

		if popErr := writeBuffer.PopContext("ReadValueId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadValueId")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadValueId) IsReadValueId() {}

func (m *_ReadValueId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
