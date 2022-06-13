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

// ApduDataExtDomainAddressResponse is the data-structure of this message
type ApduDataExtDomainAddressResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtDomainAddressResponse is the corresponding interface of ApduDataExtDomainAddressResponse
type IApduDataExtDomainAddressResponse interface {
	IApduDataExt
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

func (m *ApduDataExtDomainAddressResponse) GetExtApciType() uint8 {
	return 0x22
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtDomainAddressResponse) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtDomainAddressResponse) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtDomainAddressResponse factory function for ApduDataExtDomainAddressResponse
func NewApduDataExtDomainAddressResponse(length uint8) *ApduDataExtDomainAddressResponse {
	_result := &ApduDataExtDomainAddressResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtDomainAddressResponse(structType interface{}) *ApduDataExtDomainAddressResponse {
	if casted, ok := structType.(ApduDataExtDomainAddressResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtDomainAddressResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtDomainAddressResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtDomainAddressResponse) GetTypeName() string {
	return "ApduDataExtDomainAddressResponse"
}

func (m *ApduDataExtDomainAddressResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtDomainAddressResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtDomainAddressResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtDomainAddressResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtDomainAddressResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressResponse")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtDomainAddressResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtDomainAddressResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtDomainAddressResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
