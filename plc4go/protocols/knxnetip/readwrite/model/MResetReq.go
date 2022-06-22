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

// MResetReq is the corresponding interface of MResetReq
type MResetReq interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MResetReqExactly can be used when we want exactly this type and not a type which fulfills MResetReq.
// This is useful for switch cases.
type MResetReqExactly interface {
	MResetReq
	isMResetReq() bool
}

// _MResetReq is the data-structure of this message
type _MResetReq struct {
	*_CEMI

	// Arguments.
	Size uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MResetReq) GetMessageCode() uint8 {
	return 0xF1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MResetReq) InitializeParent(parent CEMI) {}

func (m *_MResetReq) GetParent() CEMI {
	return m._CEMI
}

// NewMResetReq factory function for _MResetReq
func NewMResetReq(size uint16) *_MResetReq {
	_result := &_MResetReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMResetReq(structType interface{}) MResetReq {
	if casted, ok := structType.(MResetReq); ok {
		return casted
	}
	if casted, ok := structType.(*MResetReq); ok {
		return *casted
	}
	return nil
}

func (m *_MResetReq) GetTypeName() string {
	return "MResetReq"
}

func (m *_MResetReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MResetReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_MResetReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MResetReqParse(readBuffer utils.ReadBuffer, size uint16) (MResetReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MResetReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MResetReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MResetReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MResetReq")
	}

	// Create a partially initialized instance
	_child := &_MResetReq{
		_CEMI: &_CEMI{},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MResetReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MResetReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MResetReq")
		}

		if popErr := writeBuffer.PopContext("MResetReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MResetReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MResetReq) isMResetReq() bool {
	return true
}

func (m *_MResetReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
