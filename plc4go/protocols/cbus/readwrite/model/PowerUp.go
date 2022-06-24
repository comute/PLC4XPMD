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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const PowerUp_PLUS byte = 0x02B
const PowerUp_CR byte = 0x0D
const PowerUp_LF byte = 0x0A

// PowerUp is the corresponding interface of PowerUp
type PowerUp interface {
	utils.LengthAware
	utils.Serializable
}

// PowerUpExactly can be used when we want exactly this type and not a type which fulfills PowerUp.
// This is useful for switch cases.
type PowerUpExactly interface {
	PowerUp
	isPowerUp() bool
}

// _PowerUp is the data-structure of this message
type _PowerUp struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_PowerUp) GetPlus() byte {
	return PowerUp_PLUS
}

func (m *_PowerUp) GetCr() byte {
	return PowerUp_CR
}

func (m *_PowerUp) GetLf() byte {
	return PowerUp_LF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPowerUp factory function for _PowerUp
func NewPowerUp() *_PowerUp {
	return &_PowerUp{}
}

// Deprecated: use the interface for direct cast
func CastPowerUp(structType interface{}) PowerUp {
	if casted, ok := structType.(PowerUp); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUp); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUp) GetTypeName() string {
	return "PowerUp"
}

func (m *_PowerUp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_PowerUp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (plus)
	lengthInBits += 8

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *_PowerUp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func PowerUpParse(readBuffer utils.ReadBuffer) (PowerUp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PowerUp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (plus)
	plus, _plusErr := readBuffer.ReadByte("plus")
	if _plusErr != nil {
		return nil, errors.Wrap(_plusErr, "Error parsing 'plus' field")
	}
	if plus != PowerUp_PLUS {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", PowerUp_PLUS) + " but got " + fmt.Sprintf("%d", plus))
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != PowerUp_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", PowerUp_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field")
	}
	if lf != PowerUp_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", PowerUp_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("PowerUp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUp")
	}

	// Create the instance
	return NewPowerUp(), nil
}

func (m *_PowerUp) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("PowerUp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PowerUp")
	}

	// Const Field (plus)
	_plusErr := writeBuffer.WriteByte("plus", 0x02B)
	if _plusErr != nil {
		return errors.Wrap(_plusErr, "Error serializing 'plus' field")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("PowerUp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PowerUp")
	}
	return nil
}

func (m *_PowerUp) isPowerUp() bool {
	return true
}

func (m *_PowerUp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
