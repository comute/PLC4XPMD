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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataOutOfService is the corresponding interface of BACnetConstructedDataOutOfService
type BACnetConstructedDataOutOfService interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOutOfService returns OutOfService (property field)
	GetOutOfService() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataOutOfServiceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOutOfService.
// This is useful for switch cases.
type BACnetConstructedDataOutOfServiceExactly interface {
	BACnetConstructedDataOutOfService
	isBACnetConstructedDataOutOfService() bool
}

// _BACnetConstructedDataOutOfService is the data-structure of this message
type _BACnetConstructedDataOutOfService struct {
	*_BACnetConstructedData
	OutOfService BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOutOfService) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OUT_OF_SERVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOutOfService) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOutOfService) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetOutOfService() BACnetApplicationTagBoolean {
	return m.OutOfService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOutOfService) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetOutOfService())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOutOfService factory function for _BACnetConstructedDataOutOfService
func NewBACnetConstructedDataOutOfService(outOfService BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOutOfService {
	_result := &_BACnetConstructedDataOutOfService{
		OutOfService:           outOfService,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOutOfService(structType interface{}) BACnetConstructedDataOutOfService {
	if casted, ok := structType.(BACnetConstructedDataOutOfService); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOutOfService); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOutOfService) GetTypeName() string {
	return "BACnetConstructedDataOutOfService"
}

func (m *_BACnetConstructedDataOutOfService) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOutOfService) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (outOfService)
	lengthInBits += m.OutOfService.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOutOfService) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOutOfServiceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOutOfService, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOutOfService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOutOfService")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (outOfService)
	if pullErr := readBuffer.PullContext("outOfService"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for outOfService")
	}
	_outOfService, _outOfServiceErr := BACnetApplicationTagParse(readBuffer)
	if _outOfServiceErr != nil {
		return nil, errors.Wrap(_outOfServiceErr, "Error parsing 'outOfService' field")
	}
	outOfService := _outOfService.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("outOfService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for outOfService")
	}

	// Virtual field
	_actualValue := outOfService
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOutOfService"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOutOfService")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOutOfService{
		OutOfService: outOfService,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOutOfService) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOutOfService"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOutOfService")
		}

		// Simple Field (outOfService)
		if pushErr := writeBuffer.PushContext("outOfService"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for outOfService")
		}
		_outOfServiceErr := writeBuffer.WriteSerializable(m.GetOutOfService())
		if popErr := writeBuffer.PopContext("outOfService"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for outOfService")
		}
		if _outOfServiceErr != nil {
			return errors.Wrap(_outOfServiceErr, "Error serializing 'outOfService' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOutOfService"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOutOfService")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOutOfService) isBACnetConstructedDataOutOfService() bool {
	return true
}

func (m *_BACnetConstructedDataOutOfService) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
