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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersExtended is the corresponding interface of BACnetNotificationParametersExtended
type BACnetNotificationParametersExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedEventType returns ExtendedEventType (property field)
	GetExtendedEventType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetNotificationParametersExtendedParameters
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersExtendedExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersExtended.
// This is useful for switch cases.
type BACnetNotificationParametersExtendedExactly interface {
	BACnetNotificationParametersExtended
	isBACnetNotificationParametersExtended() bool
}

// _BACnetNotificationParametersExtended is the data-structure of this message
type _BACnetNotificationParametersExtended struct {
	*_BACnetNotificationParameters
	InnerOpeningTag   BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedEventType BACnetContextTagUnsignedInteger
	Parameters        BACnetNotificationParametersExtendedParameters
	InnerClosingTag   BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersExtended) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersExtended) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersExtended) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetNotificationParametersExtended) GetExtendedEventType() BACnetContextTagUnsignedInteger {
	return m.ExtendedEventType
}

func (m *_BACnetNotificationParametersExtended) GetParameters() BACnetNotificationParametersExtendedParameters {
	return m.Parameters
}

func (m *_BACnetNotificationParametersExtended) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersExtended factory function for _BACnetNotificationParametersExtended
func NewBACnetNotificationParametersExtended(innerOpeningTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetNotificationParametersExtendedParameters, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersExtended {
	_result := &_BACnetNotificationParametersExtended{
		InnerOpeningTag:               innerOpeningTag,
		VendorId:                      vendorId,
		ExtendedEventType:             extendedEventType,
		Parameters:                    parameters,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersExtended(structType interface{}) BACnetNotificationParametersExtended {
	if casted, ok := structType.(BACnetNotificationParametersExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersExtended) GetTypeName() string {
	return "BACnetNotificationParametersExtended"
}

func (m *_BACnetNotificationParametersExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersExtendedParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersExtended, error) {
	return BACnetNotificationParametersExtendedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersExtendedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersExtended")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of BACnetNotificationParametersExtended")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (extendedEventType)
	if pullErr := readBuffer.PullContext("extendedEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedEventType")
	}
	_extendedEventType, _extendedEventTypeErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _extendedEventTypeErr != nil {
		return nil, errors.Wrap(_extendedEventTypeErr, "Error parsing 'extendedEventType' field of BACnetNotificationParametersExtended")
	}
	extendedEventType := _extendedEventType.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("extendedEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedEventType")
	}

	// Simple Field (parameters)
	if pullErr := readBuffer.PullContext("parameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	_parameters, _parametersErr := BACnetNotificationParametersExtendedParametersParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _parametersErr != nil {
		return nil, errors.Wrap(_parametersErr, "Error parsing 'parameters' field of BACnetNotificationParametersExtended")
	}
	parameters := _parameters.(BACnetNotificationParametersExtendedParameters)
	if closeErr := readBuffer.CloseContext("parameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersExtended")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersExtended")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersExtended{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag:   innerOpeningTag,
		VendorId:          vendorId,
		ExtendedEventType: extendedEventType,
		Parameters:        parameters,
		InnerClosingTag:   innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersExtended")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (extendedEventType)
		if pushErr := writeBuffer.PushContext("extendedEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedEventType")
		}
		_extendedEventTypeErr := writeBuffer.WriteSerializable(ctx, m.GetExtendedEventType())
		if popErr := writeBuffer.PopContext("extendedEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedEventType")
		}
		if _extendedEventTypeErr != nil {
			return errors.Wrap(_extendedEventTypeErr, "Error serializing 'extendedEventType' field")
		}

		// Simple Field (parameters)
		if pushErr := writeBuffer.PushContext("parameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameters")
		}
		_parametersErr := writeBuffer.WriteSerializable(ctx, m.GetParameters())
		if popErr := writeBuffer.PopContext("parameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameters")
		}
		if _parametersErr != nil {
			return errors.Wrap(_parametersErr, "Error serializing 'parameters' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersExtended")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersExtended) isBACnetNotificationParametersExtended() bool {
	return true
}

func (m *_BACnetNotificationParametersExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
