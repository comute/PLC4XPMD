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

// BACnetConstructedDataAlertEnrollmentAll is the corresponding interface of BACnetConstructedDataAlertEnrollmentAll
type BACnetConstructedDataAlertEnrollmentAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataAlertEnrollmentAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAlertEnrollmentAll.
// This is useful for switch cases.
type BACnetConstructedDataAlertEnrollmentAllExactly interface {
	BACnetConstructedDataAlertEnrollmentAll
	isBACnetConstructedDataAlertEnrollmentAll() bool
}

// _BACnetConstructedDataAlertEnrollmentAll is the data-structure of this message
type _BACnetConstructedDataAlertEnrollmentAll struct {
	*_BACnetConstructedData

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ALERT_ENROLLMENT
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlertEnrollmentAll) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

// NewBACnetConstructedDataAlertEnrollmentAll factory function for _BACnetConstructedDataAlertEnrollmentAll
func NewBACnetConstructedDataAlertEnrollmentAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlertEnrollmentAll {
	_result := &_BACnetConstructedDataAlertEnrollmentAll{
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlertEnrollmentAll(structType interface{}) BACnetConstructedDataAlertEnrollmentAll {
	if casted, ok := structType.(BACnetConstructedDataAlertEnrollmentAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlertEnrollmentAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetTypeName() string {
	return "BACnetConstructedDataAlertEnrollmentAll"
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAlertEnrollmentAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAlertEnrollmentAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlertEnrollmentAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlertEnrollmentAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlertEnrollmentAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlertEnrollmentAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAlertEnrollmentAll{
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlertEnrollmentAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlertEnrollmentAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlertEnrollmentAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlertEnrollmentAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) isBACnetConstructedDataAlertEnrollmentAll() bool {
	return true
}

func (m *_BACnetConstructedDataAlertEnrollmentAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
