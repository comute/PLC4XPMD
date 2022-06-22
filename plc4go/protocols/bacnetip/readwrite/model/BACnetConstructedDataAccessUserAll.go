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

// BACnetConstructedDataAccessUserAll is the corresponding interface of BACnetConstructedDataAccessUserAll
type BACnetConstructedDataAccessUserAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataAccessUserAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessUserAll.
// This is useful for switch cases.
type BACnetConstructedDataAccessUserAllExactly interface {
	BACnetConstructedDataAccessUserAll
	isBACnetConstructedDataAccessUserAll() bool
}

// _BACnetConstructedDataAccessUserAll is the data-structure of this message
type _BACnetConstructedDataAccessUserAll struct {
	*_BACnetConstructedData

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessUserAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_USER
}

func (m *_BACnetConstructedDataAccessUserAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessUserAll) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessUserAll) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

// NewBACnetConstructedDataAccessUserAll factory function for _BACnetConstructedDataAccessUserAll
func NewBACnetConstructedDataAccessUserAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessUserAll {
	_result := &_BACnetConstructedDataAccessUserAll{
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessUserAll(structType interface{}) BACnetConstructedDataAccessUserAll {
	if casted, ok := structType.(BACnetConstructedDataAccessUserAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessUserAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessUserAll) GetTypeName() string {
	return "BACnetConstructedDataAccessUserAll"
}

func (m *_BACnetConstructedDataAccessUserAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccessUserAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessUserAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessUserAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessUserAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessUserAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessUserAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessUserAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessUserAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessUserAll{
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessUserAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessUserAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessUserAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessUserAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessUserAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessUserAll) isBACnetConstructedDataAccessUserAll() bool {
	return true
}

func (m *_BACnetConstructedDataAccessUserAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
