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

// NodeIdGuid is the corresponding interface of NodeIdGuid
type NodeIdGuid interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() []byte
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdGuidExactly can be used when we want exactly this type and not a type which fulfills NodeIdGuid.
// This is useful for switch cases.
type NodeIdGuidExactly interface {
	NodeIdGuid
	isNodeIdGuid() bool
}

// _NodeIdGuid is the data-structure of this message
type _NodeIdGuid struct {
	*_NodeIdTypeDefinition
	NamespaceIndex uint16
	Id             []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdGuid) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeGuid
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdGuid) InitializeParent(parent NodeIdTypeDefinition) {}

func (m *_NodeIdGuid) GetParent() NodeIdTypeDefinition {
	return m._NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdGuid) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdGuid) GetId() []byte {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdGuid) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdGuid factory function for _NodeIdGuid
func NewNodeIdGuid(namespaceIndex uint16, id []byte) *_NodeIdGuid {
	_result := &_NodeIdGuid{
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
		_NodeIdTypeDefinition: NewNodeIdTypeDefinition(),
	}
	_result._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdGuid(structType any) NodeIdGuid {
	if casted, ok := structType.(NodeIdGuid); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdGuid); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdGuid) GetTypeName() string {
	return "NodeIdGuid"
}

func (m *_NodeIdGuid) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Array field
	if len(m.Id) > 0 {
		lengthInBits += 8 * uint16(len(m.Id))
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdGuid) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdGuidParse(ctx context.Context, theBytes []byte) (NodeIdGuid, error) {
	return NodeIdGuidParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdGuidParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdGuid, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdGuid, error) {
		return NodeIdGuidParseWithBuffer(ctx, readBuffer)
	}
}

func NodeIdGuidParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdGuid, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdGuid"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdGuid")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}

	id, err := readBuffer.ReadByteArray("id", int(int32(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdGuid"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdGuid")
	}

	// Create a partially initialized instance
	_child := &_NodeIdGuid{
		_NodeIdTypeDefinition: &_NodeIdTypeDefinition{},
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
	}
	_child._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeIdGuid) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdGuid) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdGuid"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdGuid")
		}

		// Simple Field (namespaceIndex)
		namespaceIndex := uint16(m.GetNamespaceIndex())
		_namespaceIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("namespaceIndex", 16, uint16((namespaceIndex)))
		if _namespaceIndexErr != nil {
			return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteByteArrayField(ctx, "id", m.GetId(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdGuid"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdGuid")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdGuid) isNodeIdGuid() bool {
	return true
}

func (m *_NodeIdGuid) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
