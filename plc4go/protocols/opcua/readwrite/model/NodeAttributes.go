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

// NodeAttributes is the corresponding interface of NodeAttributes
type NodeAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSpecifiedAttributes returns SpecifiedAttributes (property field)
	GetSpecifiedAttributes() uint32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
}

// NodeAttributesExactly can be used when we want exactly this type and not a type which fulfills NodeAttributes.
// This is useful for switch cases.
type NodeAttributesExactly interface {
	NodeAttributes
	isNodeAttributes() bool
}

// _NodeAttributes is the data-structure of this message
type _NodeAttributes struct {
	*_ExtensionObjectDefinition
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeAttributes) GetIdentifier() string {
	return "351"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeAttributes) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_NodeAttributes) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_NodeAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_NodeAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_NodeAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_NodeAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeAttributes factory function for _NodeAttributes
func NewNodeAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32) *_NodeAttributes {
	_result := &_NodeAttributes{
		SpecifiedAttributes:        specifiedAttributes,
		DisplayName:                displayName,
		Description:                description,
		WriteMask:                  writeMask,
		UserWriteMask:              userWriteMask,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeAttributes(structType any) NodeAttributes {
	if casted, ok := structType.(NodeAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*NodeAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_NodeAttributes) GetTypeName() string {
	return "NodeAttributes"
}

func (m *_NodeAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (specifiedAttributes)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_NodeAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeAttributesParse(ctx context.Context, theBytes []byte, identifier string) (NodeAttributes, error) {
	return NodeAttributesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NodeAttributesParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (NodeAttributes, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NodeAttributes, error) {
		return NodeAttributesParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func NodeAttributesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NodeAttributes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}

	if closeErr := readBuffer.CloseContext("NodeAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeAttributes")
	}

	// Create a partially initialized instance
	_child := &_NodeAttributes{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SpecifiedAttributes:        specifiedAttributes,
		DisplayName:                displayName,
		Description:                description,
		WriteMask:                  writeMask,
		UserWriteMask:              userWriteMask,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeAttributes")
		}

		// Simple Field (specifiedAttributes)
		specifiedAttributes := uint32(m.GetSpecifiedAttributes())
		_specifiedAttributesErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("specifiedAttributes", 32, uint32((specifiedAttributes)))
		if _specifiedAttributesErr != nil {
			return errors.Wrap(_specifiedAttributesErr, "Error serializing 'specifiedAttributes' field")
		}

		// Simple Field (displayName)
		if pushErr := writeBuffer.PushContext("displayName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for displayName")
		}
		_displayNameErr := writeBuffer.WriteSerializable(ctx, m.GetDisplayName())
		if popErr := writeBuffer.PopContext("displayName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for displayName")
		}
		if _displayNameErr != nil {
			return errors.Wrap(_displayNameErr, "Error serializing 'displayName' field")
		}

		// Simple Field (description)
		if pushErr := writeBuffer.PushContext("description"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for description")
		}
		_descriptionErr := writeBuffer.WriteSerializable(ctx, m.GetDescription())
		if popErr := writeBuffer.PopContext("description"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for description")
		}
		if _descriptionErr != nil {
			return errors.Wrap(_descriptionErr, "Error serializing 'description' field")
		}

		// Simple Field (writeMask)
		writeMask := uint32(m.GetWriteMask())
		_writeMaskErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("writeMask", 32, uint32((writeMask)))
		if _writeMaskErr != nil {
			return errors.Wrap(_writeMaskErr, "Error serializing 'writeMask' field")
		}

		// Simple Field (userWriteMask)
		userWriteMask := uint32(m.GetUserWriteMask())
		_userWriteMaskErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("userWriteMask", 32, uint32((userWriteMask)))
		if _userWriteMaskErr != nil {
			return errors.Wrap(_userWriteMaskErr, "Error serializing 'userWriteMask' field")
		}

		if popErr := writeBuffer.PopContext("NodeAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeAttributes")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeAttributes) isNodeAttributes() bool {
	return true
}

func (m *_NodeAttributes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
