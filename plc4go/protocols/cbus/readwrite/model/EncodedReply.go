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

// EncodedReply is the corresponding interface of EncodedReply
type EncodedReply interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() byte
	// GetIsMonitoredSAL returns IsMonitoredSAL (virtual field)
	GetIsMonitoredSAL() bool
	// GetIsCalCommand returns IsCalCommand (virtual field)
	GetIsCalCommand() bool
	// GetIsStandardFormatStatus returns IsStandardFormatStatus (virtual field)
	GetIsStandardFormatStatus() bool
	// GetIsExtendedFormatStatus returns IsExtendedFormatStatus (virtual field)
	GetIsExtendedFormatStatus() bool
}

// EncodedReplyExactly can be used when we want exactly this type and not a type which fulfills EncodedReply.
// This is useful for switch cases.
type EncodedReplyExactly interface {
	EncodedReply
	isEncodedReply() bool
}

// _EncodedReply is the data-structure of this message
type _EncodedReply struct {
	_EncodedReplyChildRequirements
	PeekedByte byte

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

type _EncodedReplyChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type EncodedReplyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child EncodedReply, serializeChildFunction func() error) error
	GetTypeName() string
}

type EncodedReplyChild interface {
	utils.Serializable
	InitializeParent(parent EncodedReply, peekedByte byte)
	GetParent() *EncodedReply

	GetTypeName() string
	EncodedReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EncodedReply) GetPeekedByte() byte {
	return m.PeekedByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_EncodedReply) GetIsMonitoredSAL() bool {
	return bool(bool(bool(bool((m.GetPeekedByte()&0x3F) == (0x05))) || bool(bool((m.GetPeekedByte()) == (0x00)))) || bool(bool((m.GetPeekedByte()&0xF8) == (0x00))))
}

func (m *_EncodedReply) GetIsCalCommand() bool {
	return bool(bool(bool((m.GetPeekedByte()&0x3F) == (0x06))) || bool(m.RequestContext.GetSendCalCommandBefore()))
}

func (m *_EncodedReply) GetIsStandardFormatStatus() bool {
	return bool(bool(bool((m.GetPeekedByte()&0xC0) == (0xC0))) && bool(!(m.CBusOptions.GetExstat())))
}

func (m *_EncodedReply) GetIsExtendedFormatStatus() bool {
	return bool(bool(bool((m.GetPeekedByte()&0xE0) == (0xE0))) && bool((bool(m.CBusOptions.GetExstat()) || bool(m.RequestContext.GetSendStatusRequestLevelBefore()))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEncodedReply factory function for _EncodedReply
func NewEncodedReply(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_EncodedReply {
	return &_EncodedReply{PeekedByte: peekedByte, CBusOptions: cBusOptions, RequestContext: requestContext}
}

// Deprecated: use the interface for direct cast
func CastEncodedReply(structType interface{}) EncodedReply {
	if casted, ok := structType.(EncodedReply); ok {
		return casted
	}
	if casted, ok := structType.(*EncodedReply); ok {
		return *casted
	}
	return nil
}

func (m *_EncodedReply) GetTypeName() string {
	return "EncodedReply"
}

func (m *_EncodedReply) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_EncodedReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EncodedReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (EncodedReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EncodedReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EncodedReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedByte)
	currentPos = positionAware.GetPos()
	peekedByte, _err := readBuffer.ReadByte("peekedByte")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekedByte' field of EncodedReply")
	}

	readBuffer.Reset(currentPos)

	// Virtual field
	_isMonitoredSAL := bool(bool(bool((peekedByte&0x3F) == (0x05))) || bool(bool((peekedByte) == (0x00)))) || bool(bool((peekedByte&0xF8) == (0x00)))
	isMonitoredSAL := bool(_isMonitoredSAL)
	_ = isMonitoredSAL

	// Virtual field
	_isCalCommand := bool(bool((peekedByte&0x3F) == (0x06))) || bool(requestContext.GetSendCalCommandBefore())
	isCalCommand := bool(_isCalCommand)
	_ = isCalCommand

	// Virtual field
	_isStandardFormatStatus := bool(bool((peekedByte&0xC0) == (0xC0))) && bool(!(cBusOptions.GetExstat()))
	isStandardFormatStatus := bool(_isStandardFormatStatus)
	_ = isStandardFormatStatus

	// Virtual field
	_isExtendedFormatStatus := bool(bool((peekedByte&0xE0) == (0xE0))) && bool((bool(cBusOptions.GetExstat()) || bool(requestContext.GetSendStatusRequestLevelBefore())))
	isExtendedFormatStatus := bool(_isExtendedFormatStatus)
	_ = isExtendedFormatStatus

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type EncodedReplyChildSerializeRequirement interface {
		EncodedReply
		InitializeParent(EncodedReply, byte)
		GetParent() EncodedReply
	}
	var _childTemp interface{}
	var _child EncodedReplyChildSerializeRequirement
	var typeSwitchError error
	switch {
	case isMonitoredSAL == bool(true) && isCalCommand == bool(false) && isStandardFormatStatus == bool(false): // MonitoredSALReply
		_childTemp, typeSwitchError = MonitoredSALReplyParse(readBuffer, cBusOptions, requestContext)
	case 0 == 0 && 1 == 1 && isStandardFormatStatus == bool(true) && isExtendedFormatStatus == bool(false): // EncodedReplyStandardFormatStatusReply
		_childTemp, typeSwitchError = EncodedReplyStandardFormatStatusReplyParse(readBuffer, cBusOptions, requestContext)
	case 0 == 0 && 1 == 1 && 2 == 2 && isExtendedFormatStatus == bool(true): // EncodedReplyExtendedFormatStatusReply
		_childTemp, typeSwitchError = EncodedReplyExtendedFormatStatusReplyParse(readBuffer, cBusOptions, requestContext)
	case 0 == 0 && isCalCommand == bool(true) && 2 == 2 && 3 == 3: // EncodedReplyCALReply
		_childTemp, typeSwitchError = EncodedReplyCALReplyParse(readBuffer, cBusOptions, requestContext)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [isMonitoredSAL=%v, isCalCommand=%v, isStandardFormatStatus=%v, isExtendedFormatStatus=%v]", isMonitoredSAL, isCalCommand, isStandardFormatStatus, isExtendedFormatStatus)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of EncodedReply")
	}
	_child = _childTemp.(EncodedReplyChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("EncodedReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EncodedReply")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedByte)
	return _child, nil
}

func (pm *_EncodedReply) SerializeParent(writeBuffer utils.WriteBuffer, child EncodedReply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("EncodedReply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EncodedReply")
	}
	// Virtual field
	if _isMonitoredSALErr := writeBuffer.WriteVirtual("isMonitoredSAL", m.GetIsMonitoredSAL()); _isMonitoredSALErr != nil {
		return errors.Wrap(_isMonitoredSALErr, "Error serializing 'isMonitoredSAL' field")
	}
	// Virtual field
	if _isCalCommandErr := writeBuffer.WriteVirtual("isCalCommand", m.GetIsCalCommand()); _isCalCommandErr != nil {
		return errors.Wrap(_isCalCommandErr, "Error serializing 'isCalCommand' field")
	}
	// Virtual field
	if _isStandardFormatStatusErr := writeBuffer.WriteVirtual("isStandardFormatStatus", m.GetIsStandardFormatStatus()); _isStandardFormatStatusErr != nil {
		return errors.Wrap(_isStandardFormatStatusErr, "Error serializing 'isStandardFormatStatus' field")
	}
	// Virtual field
	if _isExtendedFormatStatusErr := writeBuffer.WriteVirtual("isExtendedFormatStatus", m.GetIsExtendedFormatStatus()); _isExtendedFormatStatusErr != nil {
		return errors.Wrap(_isExtendedFormatStatusErr, "Error serializing 'isExtendedFormatStatus' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("EncodedReply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EncodedReply")
	}
	return nil
}

////
// Arguments Getter

func (m *_EncodedReply) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_EncodedReply) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_EncodedReply) isEncodedReply() bool {
	return true
}

func (m *_EncodedReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
