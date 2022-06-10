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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetRouterEntry is the data-structure of this message
type BACnetRouterEntry struct {
	NetworkNumber    *BACnetContextTagUnsignedInteger
	MacAddress       *BACnetContextTagOctetString
	Status           *BACnetRouterEntryStatusTagged
	PerformanceIndex *BACnetContextTagOctetString
}

// IBACnetRouterEntry is the corresponding interface of BACnetRouterEntry
type IBACnetRouterEntry interface {
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() *BACnetContextTagUnsignedInteger
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() *BACnetContextTagOctetString
	// GetStatus returns Status (property field)
	GetStatus() *BACnetRouterEntryStatusTagged
	// GetPerformanceIndex returns PerformanceIndex (property field)
	GetPerformanceIndex() *BACnetContextTagOctetString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetRouterEntry) GetNetworkNumber() *BACnetContextTagUnsignedInteger {
	return m.NetworkNumber
}

func (m *BACnetRouterEntry) GetMacAddress() *BACnetContextTagOctetString {
	return m.MacAddress
}

func (m *BACnetRouterEntry) GetStatus() *BACnetRouterEntryStatusTagged {
	return m.Status
}

func (m *BACnetRouterEntry) GetPerformanceIndex() *BACnetContextTagOctetString {
	return m.PerformanceIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRouterEntry factory function for BACnetRouterEntry
func NewBACnetRouterEntry(networkNumber *BACnetContextTagUnsignedInteger, macAddress *BACnetContextTagOctetString, status *BACnetRouterEntryStatusTagged, performanceIndex *BACnetContextTagOctetString) *BACnetRouterEntry {
	return &BACnetRouterEntry{NetworkNumber: networkNumber, MacAddress: macAddress, Status: status, PerformanceIndex: performanceIndex}
}

func CastBACnetRouterEntry(structType interface{}) *BACnetRouterEntry {
	if casted, ok := structType.(BACnetRouterEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetRouterEntry); ok {
		return casted
	}
	return nil
}

func (m *BACnetRouterEntry) GetTypeName() string {
	return "BACnetRouterEntry"
}

func (m *BACnetRouterEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetRouterEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (networkNumber)
	lengthInBits += m.NetworkNumber.GetLengthInBits()

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits()

	// Simple field (status)
	lengthInBits += m.Status.GetLengthInBits()

	// Optional Field (performanceIndex)
	if m.PerformanceIndex != nil {
		lengthInBits += (*m.PerformanceIndex).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetRouterEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRouterEntryParse(readBuffer utils.ReadBuffer) (*BACnetRouterEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRouterEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRouterEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumber)
	if pullErr := readBuffer.PullContext("networkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkNumber")
	}
	_networkNumber, _networkNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _networkNumberErr != nil {
		return nil, errors.Wrap(_networkNumberErr, "Error parsing 'networkNumber' field")
	}
	networkNumber := CastBACnetContextTagUnsignedInteger(_networkNumber)
	if closeErr := readBuffer.CloseContext("networkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkNumber")
	}

	// Simple Field (macAddress)
	if pullErr := readBuffer.PullContext("macAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for macAddress")
	}
	_macAddress, _macAddressErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_OCTET_STRING))
	if _macAddressErr != nil {
		return nil, errors.Wrap(_macAddressErr, "Error parsing 'macAddress' field")
	}
	macAddress := CastBACnetContextTagOctetString(_macAddress)
	if closeErr := readBuffer.CloseContext("macAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for macAddress")
	}

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
	_status, _statusErr := BACnetRouterEntryStatusTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := CastBACnetRouterEntryStatusTagged(_status)
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	// Optional Field (performanceIndex) (Can be skipped, if a given expression evaluates to false)
	var performanceIndex *BACnetContextTagOctetString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("performanceIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for performanceIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_OCTET_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'performanceIndex' field")
		default:
			performanceIndex = CastBACnetContextTagOctetString(_val)
			if closeErr := readBuffer.CloseContext("performanceIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for performanceIndex")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetRouterEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRouterEntry")
	}

	// Create the instance
	return NewBACnetRouterEntry(networkNumber, macAddress, status, performanceIndex), nil
}

func (m *BACnetRouterEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetRouterEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRouterEntry")
	}

	// Simple Field (networkNumber)
	if pushErr := writeBuffer.PushContext("networkNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for networkNumber")
	}
	_networkNumberErr := m.NetworkNumber.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("networkNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for networkNumber")
	}
	if _networkNumberErr != nil {
		return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
	}

	// Simple Field (macAddress)
	if pushErr := writeBuffer.PushContext("macAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for macAddress")
	}
	_macAddressErr := m.MacAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("macAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for macAddress")
	}
	if _macAddressErr != nil {
		return errors.Wrap(_macAddressErr, "Error serializing 'macAddress' field")
	}

	// Simple Field (status)
	if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for status")
	}
	_statusErr := m.Status.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for status")
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Optional Field (performanceIndex) (Can be skipped, if the value is null)
	var performanceIndex *BACnetContextTagOctetString = nil
	if m.PerformanceIndex != nil {
		if pushErr := writeBuffer.PushContext("performanceIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for performanceIndex")
		}
		performanceIndex = m.PerformanceIndex
		_performanceIndexErr := performanceIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("performanceIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for performanceIndex")
		}
		if _performanceIndexErr != nil {
			return errors.Wrap(_performanceIndexErr, "Error serializing 'performanceIndex' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetRouterEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRouterEntry")
	}
	return nil
}

func (m *BACnetRouterEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
