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

// ObjectTypeAttributes is the corresponding interface of ObjectTypeAttributes
type ObjectTypeAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
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
	// GetIsAbstract returns IsAbstract (property field)
	GetIsAbstract() bool
	// IsObjectTypeAttributes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsObjectTypeAttributes()
	// CreateBuilder creates a ObjectTypeAttributesBuilder
	CreateObjectTypeAttributesBuilder() ObjectTypeAttributesBuilder
}

// _ObjectTypeAttributes is the data-structure of this message
type _ObjectTypeAttributes struct {
	ExtensionObjectDefinitionContract
	SpecifiedAttributes uint32
	DisplayName         LocalizedText
	Description         LocalizedText
	WriteMask           uint32
	UserWriteMask       uint32
	IsAbstract          bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ ObjectTypeAttributes = (*_ObjectTypeAttributes)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ObjectTypeAttributes)(nil)

// NewObjectTypeAttributes factory function for _ObjectTypeAttributes
func NewObjectTypeAttributes(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, isAbstract bool) *_ObjectTypeAttributes {
	if displayName == nil {
		panic("displayName of type LocalizedText for ObjectTypeAttributes must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for ObjectTypeAttributes must not be nil")
	}
	_result := &_ObjectTypeAttributes{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SpecifiedAttributes:               specifiedAttributes,
		DisplayName:                       displayName,
		Description:                       description,
		WriteMask:                         writeMask,
		UserWriteMask:                     userWriteMask,
		IsAbstract:                        isAbstract,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ObjectTypeAttributesBuilder is a builder for ObjectTypeAttributes
type ObjectTypeAttributesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, isAbstract bool) ObjectTypeAttributesBuilder
	// WithSpecifiedAttributes adds SpecifiedAttributes (property field)
	WithSpecifiedAttributes(uint32) ObjectTypeAttributesBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) ObjectTypeAttributesBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectTypeAttributesBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) ObjectTypeAttributesBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectTypeAttributesBuilder
	// WithWriteMask adds WriteMask (property field)
	WithWriteMask(uint32) ObjectTypeAttributesBuilder
	// WithUserWriteMask adds UserWriteMask (property field)
	WithUserWriteMask(uint32) ObjectTypeAttributesBuilder
	// WithIsAbstract adds IsAbstract (property field)
	WithIsAbstract(bool) ObjectTypeAttributesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ObjectTypeAttributes or returns an error if something is wrong
	Build() (ObjectTypeAttributes, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ObjectTypeAttributes
}

// NewObjectTypeAttributesBuilder() creates a ObjectTypeAttributesBuilder
func NewObjectTypeAttributesBuilder() ObjectTypeAttributesBuilder {
	return &_ObjectTypeAttributesBuilder{_ObjectTypeAttributes: new(_ObjectTypeAttributes)}
}

type _ObjectTypeAttributesBuilder struct {
	*_ObjectTypeAttributes

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ObjectTypeAttributesBuilder) = (*_ObjectTypeAttributesBuilder)(nil)

func (b *_ObjectTypeAttributesBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ObjectTypeAttributes
}

func (b *_ObjectTypeAttributesBuilder) WithMandatoryFields(specifiedAttributes uint32, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, isAbstract bool) ObjectTypeAttributesBuilder {
	return b.WithSpecifiedAttributes(specifiedAttributes).WithDisplayName(displayName).WithDescription(description).WithWriteMask(writeMask).WithUserWriteMask(userWriteMask).WithIsAbstract(isAbstract)
}

func (b *_ObjectTypeAttributesBuilder) WithSpecifiedAttributes(specifiedAttributes uint32) ObjectTypeAttributesBuilder {
	b.SpecifiedAttributes = specifiedAttributes
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithDisplayName(displayName LocalizedText) ObjectTypeAttributesBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectTypeAttributesBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithDescription(description LocalizedText) ObjectTypeAttributesBuilder {
	b.Description = description
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ObjectTypeAttributesBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithWriteMask(writeMask uint32) ObjectTypeAttributesBuilder {
	b.WriteMask = writeMask
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithUserWriteMask(userWriteMask uint32) ObjectTypeAttributesBuilder {
	b.UserWriteMask = userWriteMask
	return b
}

func (b *_ObjectTypeAttributesBuilder) WithIsAbstract(isAbstract bool) ObjectTypeAttributesBuilder {
	b.IsAbstract = isAbstract
	return b
}

func (b *_ObjectTypeAttributesBuilder) Build() (ObjectTypeAttributes, error) {
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ObjectTypeAttributes.deepCopy(), nil
}

func (b *_ObjectTypeAttributesBuilder) MustBuild() ObjectTypeAttributes {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ObjectTypeAttributesBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ObjectTypeAttributesBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ObjectTypeAttributesBuilder) DeepCopy() any {
	_copy := b.CreateObjectTypeAttributesBuilder().(*_ObjectTypeAttributesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateObjectTypeAttributesBuilder creates a ObjectTypeAttributesBuilder
func (b *_ObjectTypeAttributes) CreateObjectTypeAttributesBuilder() ObjectTypeAttributesBuilder {
	if b == nil {
		return NewObjectTypeAttributesBuilder()
	}
	return &_ObjectTypeAttributesBuilder{_ObjectTypeAttributes: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ObjectTypeAttributes) GetExtensionId() int32 {
	return int32(363)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ObjectTypeAttributes) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ObjectTypeAttributes) GetSpecifiedAttributes() uint32 {
	return m.SpecifiedAttributes
}

func (m *_ObjectTypeAttributes) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_ObjectTypeAttributes) GetDescription() LocalizedText {
	return m.Description
}

func (m *_ObjectTypeAttributes) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_ObjectTypeAttributes) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

func (m *_ObjectTypeAttributes) GetIsAbstract() bool {
	return m.IsAbstract
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastObjectTypeAttributes(structType any) ObjectTypeAttributes {
	if casted, ok := structType.(ObjectTypeAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*ObjectTypeAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_ObjectTypeAttributes) GetTypeName() string {
	return "ObjectTypeAttributes"
}

func (m *_ObjectTypeAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

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

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isAbstract)
	lengthInBits += 1

	return lengthInBits
}

func (m *_ObjectTypeAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ObjectTypeAttributes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__objectTypeAttributes ObjectTypeAttributes, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ObjectTypeAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ObjectTypeAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specifiedAttributes, err := ReadSimpleField(ctx, "specifiedAttributes", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specifiedAttributes' field"))
	}
	m.SpecifiedAttributes = specifiedAttributes

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	writeMask, err := ReadSimpleField(ctx, "writeMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeMask' field"))
	}
	m.WriteMask = writeMask

	userWriteMask, err := ReadSimpleField(ctx, "userWriteMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userWriteMask' field"))
	}
	m.UserWriteMask = userWriteMask

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isAbstract, err := ReadSimpleField(ctx, "isAbstract", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAbstract' field"))
	}
	m.IsAbstract = isAbstract

	if closeErr := readBuffer.CloseContext("ObjectTypeAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ObjectTypeAttributes")
	}

	return m, nil
}

func (m *_ObjectTypeAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ObjectTypeAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ObjectTypeAttributes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ObjectTypeAttributes")
		}

		if err := WriteSimpleField[uint32](ctx, "specifiedAttributes", m.GetSpecifiedAttributes(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'specifiedAttributes' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if err := WriteSimpleField[uint32](ctx, "writeMask", m.GetWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "userWriteMask", m.GetUserWriteMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'userWriteMask' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isAbstract", m.GetIsAbstract(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isAbstract' field")
		}

		if popErr := writeBuffer.PopContext("ObjectTypeAttributes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ObjectTypeAttributes")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ObjectTypeAttributes) IsObjectTypeAttributes() {}

func (m *_ObjectTypeAttributes) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ObjectTypeAttributes) deepCopy() *_ObjectTypeAttributes {
	if m == nil {
		return nil
	}
	_ObjectTypeAttributesCopy := &_ObjectTypeAttributes{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SpecifiedAttributes,
		utils.DeepCopy[LocalizedText](m.DisplayName),
		utils.DeepCopy[LocalizedText](m.Description),
		m.WriteMask,
		m.UserWriteMask,
		m.IsAbstract,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ObjectTypeAttributesCopy
}

func (m *_ObjectTypeAttributes) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
