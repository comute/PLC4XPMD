/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataSecurityPDUTimeout is the data-structure of this message
type BACnetConstructedDataSecurityPDUTimeout struct {
	*BACnetConstructedData
	SecurityPduTimeout *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSecurityPDUTimeout is the corresponding interface of BACnetConstructedDataSecurityPDUTimeout
type IBACnetConstructedDataSecurityPDUTimeout interface {
	IBACnetConstructedData
	// GetSecurityPduTimeout returns SecurityPduTimeout (property field)
	GetSecurityPduTimeout() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataSecurityPDUTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURITY_PDU_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSecurityPDUTimeout) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSecurityPDUTimeout) GetSecurityPduTimeout() *BACnetApplicationTagUnsignedInteger {
	return m.SecurityPduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSecurityPDUTimeout factory function for BACnetConstructedDataSecurityPDUTimeout
func NewBACnetConstructedDataSecurityPDUTimeout(securityPduTimeout *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSecurityPDUTimeout {
	_result := &BACnetConstructedDataSecurityPDUTimeout{
		SecurityPduTimeout:    securityPduTimeout,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSecurityPDUTimeout(structType interface{}) *BACnetConstructedDataSecurityPDUTimeout {
	if casted, ok := structType.(BACnetConstructedDataSecurityPDUTimeout); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecurityPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSecurityPDUTimeout(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSecurityPDUTimeout(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataSecurityPDUTimeout"
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (securityPduTimeout)
	lengthInBits += m.SecurityPduTimeout.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataSecurityPDUTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSecurityPDUTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSecurityPDUTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecurityPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecurityPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityPduTimeout)
	if pullErr := readBuffer.PullContext("securityPduTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPduTimeout")
	}
	_securityPduTimeout, _securityPduTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _securityPduTimeoutErr != nil {
		return nil, errors.Wrap(_securityPduTimeoutErr, "Error parsing 'securityPduTimeout' field")
	}
	securityPduTimeout := CastBACnetApplicationTagUnsignedInteger(_securityPduTimeout)
	if closeErr := readBuffer.CloseContext("securityPduTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPduTimeout")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecurityPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecurityPDUTimeout")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSecurityPDUTimeout{
		SecurityPduTimeout:    CastBACnetApplicationTagUnsignedInteger(securityPduTimeout),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSecurityPDUTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecurityPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecurityPDUTimeout")
		}

		// Simple Field (securityPduTimeout)
		if pushErr := writeBuffer.PushContext("securityPduTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPduTimeout")
		}
		_securityPduTimeoutErr := writeBuffer.WriteSerializable(m.SecurityPduTimeout)
		if popErr := writeBuffer.PopContext("securityPduTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPduTimeout")
		}
		if _securityPduTimeoutErr != nil {
			return errors.Wrap(_securityPduTimeoutErr, "Error serializing 'securityPduTimeout' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecurityPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecurityPDUTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSecurityPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
