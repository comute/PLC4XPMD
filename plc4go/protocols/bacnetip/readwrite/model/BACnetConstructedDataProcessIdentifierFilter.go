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

// BACnetConstructedDataProcessIdentifierFilter is the corresponding interface of BACnetConstructedDataProcessIdentifierFilter
type BACnetConstructedDataProcessIdentifierFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProcessIdentifierFilter returns ProcessIdentifierFilter (property field)
	GetProcessIdentifierFilter() BACnetProcessIdSelection
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProcessIdSelection
	// IsBACnetConstructedDataProcessIdentifierFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProcessIdentifierFilter()
}

// _BACnetConstructedDataProcessIdentifierFilter is the data-structure of this message
type _BACnetConstructedDataProcessIdentifierFilter struct {
	BACnetConstructedDataContract
	ProcessIdentifierFilter BACnetProcessIdSelection
}

var _ BACnetConstructedDataProcessIdentifierFilter = (*_BACnetConstructedDataProcessIdentifierFilter)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProcessIdentifierFilter)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROCESS_IDENTIFIER_FILTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetProcessIdentifierFilter() BACnetProcessIdSelection {
	return m.ProcessIdentifierFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetActualValue() BACnetProcessIdSelection {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProcessIdSelection(m.GetProcessIdentifierFilter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProcessIdentifierFilter factory function for _BACnetConstructedDataProcessIdentifierFilter
func NewBACnetConstructedDataProcessIdentifierFilter(processIdentifierFilter BACnetProcessIdSelection, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProcessIdentifierFilter {
	if processIdentifierFilter == nil {
		panic("processIdentifierFilter of type BACnetProcessIdSelection for BACnetConstructedDataProcessIdentifierFilter must not be nil")
	}
	_result := &_BACnetConstructedDataProcessIdentifierFilter{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProcessIdentifierFilter:       processIdentifierFilter,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProcessIdentifierFilter(structType any) BACnetConstructedDataProcessIdentifierFilter {
	if casted, ok := structType.(BACnetConstructedDataProcessIdentifierFilter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProcessIdentifierFilter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetTypeName() string {
	return "BACnetConstructedDataProcessIdentifierFilter"
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (processIdentifierFilter)
	lengthInBits += m.ProcessIdentifierFilter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProcessIdentifierFilter BACnetConstructedDataProcessIdentifierFilter, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProcessIdentifierFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProcessIdentifierFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifierFilter, err := ReadSimpleField[BACnetProcessIdSelection](ctx, "processIdentifierFilter", ReadComplex[BACnetProcessIdSelection](BACnetProcessIdSelectionParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifierFilter' field"))
	}
	m.ProcessIdentifierFilter = processIdentifierFilter

	actualValue, err := ReadVirtualField[BACnetProcessIdSelection](ctx, "actualValue", (*BACnetProcessIdSelection)(nil), processIdentifierFilter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProcessIdentifierFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProcessIdentifierFilter")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProcessIdentifierFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProcessIdentifierFilter")
		}

		if err := WriteSimpleField[BACnetProcessIdSelection](ctx, "processIdentifierFilter", m.GetProcessIdentifierFilter(), WriteComplex[BACnetProcessIdSelection](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifierFilter' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProcessIdentifierFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProcessIdentifierFilter")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) IsBACnetConstructedDataProcessIdentifierFilter() {
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
