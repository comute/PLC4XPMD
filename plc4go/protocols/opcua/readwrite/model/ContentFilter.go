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

// ContentFilter is the corresponding interface of ContentFilter
type ContentFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetElements returns Elements (property field)
	GetElements() []ContentFilterElement
	// IsContentFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsContentFilter()
	// CreateBuilder creates a ContentFilterBuilder
	CreateContentFilterBuilder() ContentFilterBuilder
}

// _ContentFilter is the data-structure of this message
type _ContentFilter struct {
	ExtensionObjectDefinitionContract
	Elements []ContentFilterElement
}

var _ ContentFilter = (*_ContentFilter)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ContentFilter)(nil)

// NewContentFilter factory function for _ContentFilter
func NewContentFilter(elements []ContentFilterElement) *_ContentFilter {
	_result := &_ContentFilter{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Elements:                          elements,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ContentFilterBuilder is a builder for ContentFilter
type ContentFilterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(elements []ContentFilterElement) ContentFilterBuilder
	// WithElements adds Elements (property field)
	WithElements(...ContentFilterElement) ContentFilterBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ContentFilter or returns an error if something is wrong
	Build() (ContentFilter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ContentFilter
}

// NewContentFilterBuilder() creates a ContentFilterBuilder
func NewContentFilterBuilder() ContentFilterBuilder {
	return &_ContentFilterBuilder{_ContentFilter: new(_ContentFilter)}
}

type _ContentFilterBuilder struct {
	*_ContentFilter

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ContentFilterBuilder) = (*_ContentFilterBuilder)(nil)

func (b *_ContentFilterBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ContentFilter
}

func (b *_ContentFilterBuilder) WithMandatoryFields(elements []ContentFilterElement) ContentFilterBuilder {
	return b.WithElements(elements...)
}

func (b *_ContentFilterBuilder) WithElements(elements ...ContentFilterElement) ContentFilterBuilder {
	b.Elements = elements
	return b
}

func (b *_ContentFilterBuilder) Build() (ContentFilter, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ContentFilter.deepCopy(), nil
}

func (b *_ContentFilterBuilder) MustBuild() ContentFilter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ContentFilterBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ContentFilterBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ContentFilterBuilder) DeepCopy() any {
	_copy := b.CreateContentFilterBuilder().(*_ContentFilterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateContentFilterBuilder creates a ContentFilterBuilder
func (b *_ContentFilter) CreateContentFilterBuilder() ContentFilterBuilder {
	if b == nil {
		return NewContentFilterBuilder()
	}
	return &_ContentFilterBuilder{_ContentFilter: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilter) GetExtensionId() int32 {
	return int32(588)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilter) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilter) GetElements() []ContentFilterElement {
	return m.Elements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastContentFilter(structType any) ContentFilter {
	if casted, ok := structType.(ContentFilter); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilter); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilter) GetTypeName() string {
	return "ContentFilter"
}

func (m *_ContentFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfElements)
	lengthInBits += 32

	// Array field
	if len(m.Elements) > 0 {
		for _curItem, element := range m.Elements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Elements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ContentFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__contentFilter ContentFilter, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContentFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElements, err := ReadImplicitField[int32](ctx, "noOfElements", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElements' field"))
	}
	_ = noOfElements

	elements, err := ReadCountArrayField[ContentFilterElement](ctx, "elements", ReadComplex[ContentFilterElement](ExtensionObjectDefinitionParseWithBufferProducer[ContentFilterElement]((int32)(int32(585))), readBuffer), uint64(noOfElements))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elements' field"))
	}
	m.Elements = elements

	if closeErr := readBuffer.CloseContext("ContentFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilter")
	}

	return m, nil
}

func (m *_ContentFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilter")
		}
		noOfElements := int32(utils.InlineIf(bool((m.GetElements()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetElements()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfElements", noOfElements, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "elements", m.GetElements(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'elements' field")
		}

		if popErr := writeBuffer.PopContext("ContentFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilter")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilter) IsContentFilter() {}

func (m *_ContentFilter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ContentFilter) deepCopy() *_ContentFilter {
	if m == nil {
		return nil
	}
	_ContentFilterCopy := &_ContentFilter{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[ContentFilterElement, ContentFilterElement](m.Elements),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ContentFilterCopy
}

func (m *_ContentFilter) String() string {
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
