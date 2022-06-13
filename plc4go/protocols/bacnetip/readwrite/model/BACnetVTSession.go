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

// BACnetVTSession is the data-structure of this message
type BACnetVTSession struct {
	LocalVtSessionId  *BACnetApplicationTagUnsignedInteger
	RemoveVtSessionId *BACnetApplicationTagUnsignedInteger
	RemoteVtAddress   *BACnetAddress
}

// IBACnetVTSession is the corresponding interface of BACnetVTSession
type IBACnetVTSession interface {
	// GetLocalVtSessionId returns LocalVtSessionId (property field)
	GetLocalVtSessionId() *BACnetApplicationTagUnsignedInteger
	// GetRemoveVtSessionId returns RemoveVtSessionId (property field)
	GetRemoveVtSessionId() *BACnetApplicationTagUnsignedInteger
	// GetRemoteVtAddress returns RemoteVtAddress (property field)
	GetRemoteVtAddress() *BACnetAddress
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

func (m *BACnetVTSession) GetLocalVtSessionId() *BACnetApplicationTagUnsignedInteger {
	return m.LocalVtSessionId
}

func (m *BACnetVTSession) GetRemoveVtSessionId() *BACnetApplicationTagUnsignedInteger {
	return m.RemoveVtSessionId
}

func (m *BACnetVTSession) GetRemoteVtAddress() *BACnetAddress {
	return m.RemoteVtAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVTSession factory function for BACnetVTSession
func NewBACnetVTSession(localVtSessionId *BACnetApplicationTagUnsignedInteger, removeVtSessionId *BACnetApplicationTagUnsignedInteger, remoteVtAddress *BACnetAddress) *BACnetVTSession {
	return &BACnetVTSession{LocalVtSessionId: localVtSessionId, RemoveVtSessionId: removeVtSessionId, RemoteVtAddress: remoteVtAddress}
}

func CastBACnetVTSession(structType interface{}) *BACnetVTSession {
	if casted, ok := structType.(BACnetVTSession); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetVTSession); ok {
		return casted
	}
	return nil
}

func (m *BACnetVTSession) GetTypeName() string {
	return "BACnetVTSession"
}

func (m *BACnetVTSession) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetVTSession) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (localVtSessionId)
	lengthInBits += m.LocalVtSessionId.GetLengthInBits()

	// Simple field (removeVtSessionId)
	lengthInBits += m.RemoveVtSessionId.GetLengthInBits()

	// Simple field (remoteVtAddress)
	lengthInBits += m.RemoteVtAddress.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetVTSession) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetVTSessionParse(readBuffer utils.ReadBuffer) (*BACnetVTSession, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVTSession"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVTSession")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localVtSessionId)
	if pullErr := readBuffer.PullContext("localVtSessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localVtSessionId")
	}
	_localVtSessionId, _localVtSessionIdErr := BACnetApplicationTagParse(readBuffer)
	if _localVtSessionIdErr != nil {
		return nil, errors.Wrap(_localVtSessionIdErr, "Error parsing 'localVtSessionId' field")
	}
	localVtSessionId := CastBACnetApplicationTagUnsignedInteger(_localVtSessionId)
	if closeErr := readBuffer.CloseContext("localVtSessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localVtSessionId")
	}

	// Simple Field (removeVtSessionId)
	if pullErr := readBuffer.PullContext("removeVtSessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for removeVtSessionId")
	}
	_removeVtSessionId, _removeVtSessionIdErr := BACnetApplicationTagParse(readBuffer)
	if _removeVtSessionIdErr != nil {
		return nil, errors.Wrap(_removeVtSessionIdErr, "Error parsing 'removeVtSessionId' field")
	}
	removeVtSessionId := CastBACnetApplicationTagUnsignedInteger(_removeVtSessionId)
	if closeErr := readBuffer.CloseContext("removeVtSessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for removeVtSessionId")
	}

	// Simple Field (remoteVtAddress)
	if pullErr := readBuffer.PullContext("remoteVtAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for remoteVtAddress")
	}
	_remoteVtAddress, _remoteVtAddressErr := BACnetAddressParse(readBuffer)
	if _remoteVtAddressErr != nil {
		return nil, errors.Wrap(_remoteVtAddressErr, "Error parsing 'remoteVtAddress' field")
	}
	remoteVtAddress := CastBACnetAddress(_remoteVtAddress)
	if closeErr := readBuffer.CloseContext("remoteVtAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for remoteVtAddress")
	}

	if closeErr := readBuffer.CloseContext("BACnetVTSession"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVTSession")
	}

	// Create the instance
	return NewBACnetVTSession(localVtSessionId, removeVtSessionId, remoteVtAddress), nil
}

func (m *BACnetVTSession) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetVTSession"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVTSession")
	}

	// Simple Field (localVtSessionId)
	if pushErr := writeBuffer.PushContext("localVtSessionId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for localVtSessionId")
	}
	_localVtSessionIdErr := writeBuffer.WriteSerializable(m.LocalVtSessionId)
	if popErr := writeBuffer.PopContext("localVtSessionId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for localVtSessionId")
	}
	if _localVtSessionIdErr != nil {
		return errors.Wrap(_localVtSessionIdErr, "Error serializing 'localVtSessionId' field")
	}

	// Simple Field (removeVtSessionId)
	if pushErr := writeBuffer.PushContext("removeVtSessionId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for removeVtSessionId")
	}
	_removeVtSessionIdErr := writeBuffer.WriteSerializable(m.RemoveVtSessionId)
	if popErr := writeBuffer.PopContext("removeVtSessionId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for removeVtSessionId")
	}
	if _removeVtSessionIdErr != nil {
		return errors.Wrap(_removeVtSessionIdErr, "Error serializing 'removeVtSessionId' field")
	}

	// Simple Field (remoteVtAddress)
	if pushErr := writeBuffer.PushContext("remoteVtAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for remoteVtAddress")
	}
	_remoteVtAddressErr := writeBuffer.WriteSerializable(m.RemoteVtAddress)
	if popErr := writeBuffer.PopContext("remoteVtAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for remoteVtAddress")
	}
	if _remoteVtAddressErr != nil {
		return errors.Wrap(_remoteVtAddressErr, "Error serializing 'remoteVtAddress' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVTSession"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVTSession")
	}
	return nil
}

func (m *BACnetVTSession) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
