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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataVTClassesSupported is the corresponding interface of BACnetConstructedDataVTClassesSupported
type BACnetConstructedDataVTClassesSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVtClassesSupported returns VtClassesSupported (property field)
	GetVtClassesSupported() []BACnetVTClassTagged
}

// BACnetConstructedDataVTClassesSupportedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVTClassesSupported.
// This is useful for switch cases.
type BACnetConstructedDataVTClassesSupportedExactly interface {
	BACnetConstructedDataVTClassesSupported
	isBACnetConstructedDataVTClassesSupported() bool
}

// _BACnetConstructedDataVTClassesSupported is the data-structure of this message
type _BACnetConstructedDataVTClassesSupported struct {
	*_BACnetConstructedData
	VtClassesSupported []BACnetVTClassTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVTClassesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VT_CLASSES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVTClassesSupported) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetVtClassesSupported() []BACnetVTClassTagged {
	return m.VtClassesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVTClassesSupported factory function for _BACnetConstructedDataVTClassesSupported
func NewBACnetConstructedDataVTClassesSupported(vtClassesSupported []BACnetVTClassTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVTClassesSupported {
	_result := &_BACnetConstructedDataVTClassesSupported{
		VtClassesSupported:     vtClassesSupported,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVTClassesSupported(structType interface{}) BACnetConstructedDataVTClassesSupported {
	if casted, ok := structType.(BACnetConstructedDataVTClassesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVTClassesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVTClassesSupported) GetTypeName() string {
	return "BACnetConstructedDataVTClassesSupported"
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.VtClassesSupported) > 0 {
		for _, element := range m.VtClassesSupported {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataVTClassesSupportedParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVTClassesSupported, error) {
	return BACnetConstructedDataVTClassesSupportedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataVTClassesSupportedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVTClassesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVTClassesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVTClassesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (vtClassesSupported)
	if pullErr := readBuffer.PullContext("vtClassesSupported", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtClassesSupported")
	}
	// Terminated array
	var vtClassesSupported []BACnetVTClassTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetVTClassTaggedParseWithBuffer(ctx, readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'vtClassesSupported' field of BACnetConstructedDataVTClassesSupported")
			}
			vtClassesSupported = append(vtClassesSupported, _item.(BACnetVTClassTagged))
		}
	}
	if closeErr := readBuffer.CloseContext("vtClassesSupported", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtClassesSupported")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVTClassesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVTClassesSupported")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVTClassesSupported{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VtClassesSupported: vtClassesSupported,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVTClassesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVTClassesSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVTClassesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVTClassesSupported")
		}

		// Array Field (vtClassesSupported)
		if pushErr := writeBuffer.PushContext("vtClassesSupported", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtClassesSupported")
		}
		for _curItem, _element := range m.GetVtClassesSupported() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetVtClassesSupported()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'vtClassesSupported' field")
			}
		}
		if popErr := writeBuffer.PopContext("vtClassesSupported", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtClassesSupported")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVTClassesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVTClassesSupported")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVTClassesSupported) isBACnetConstructedDataVTClassesSupported() bool {
	return true
}

func (m *_BACnetConstructedDataVTClassesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
