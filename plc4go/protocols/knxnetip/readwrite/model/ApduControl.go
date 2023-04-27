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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduControl is the corresponding interface of ApduControl
type ApduControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetControlType returns ControlType (discriminator field)
	GetControlType() uint8
}

// ApduControlExactly can be used when we want exactly this type and not a type which fulfills ApduControl.
// This is useful for switch cases.
type ApduControlExactly interface {
	ApduControl
	isApduControl() bool
}

// _ApduControl is the data-structure of this message
type _ApduControl struct {
	_ApduControlChildRequirements
}

type _ApduControlChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetControlType() uint8
}

type ApduControlParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduControl, serializeChildFunction func() error) error
	GetTypeName() string
}

type ApduControlChild interface {
	utils.Serializable
	InitializeParent(parent ApduControl)
	GetParent() *ApduControl

	GetTypeName() string
	ApduControl
}

// NewApduControl factory function for _ApduControl
func NewApduControl() *_ApduControl {
	return &_ApduControl{}
}

// Deprecated: use the interface for direct cast
func CastApduControl(structType any) ApduControl {
	if casted, ok := structType.(ApduControl); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControl); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControl) GetTypeName() string {
	return "ApduControl"
}

func (m *_ApduControl) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (controlType)
	lengthInBits += 2

	return lengthInBits
}

func (m *_ApduControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduControlParse(theBytes []byte) (ApduControl, error) {
	return ApduControlParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ApduControlParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApduControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType, _controlTypeErr := readBuffer.ReadUint8("controlType", 2)
	if _controlTypeErr != nil {
		return nil, errors.Wrap(_controlTypeErr, "Error parsing 'controlType' field of ApduControl")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ApduControlChildSerializeRequirement interface {
		ApduControl
		InitializeParent(ApduControl)
		GetParent() ApduControl
	}
	var _childTemp any
	var _child ApduControlChildSerializeRequirement
	var typeSwitchError error
	switch {
	case controlType == 0x0: // ApduControlConnect
		_childTemp, typeSwitchError = ApduControlConnectParseWithBuffer(ctx, readBuffer)
	case controlType == 0x1: // ApduControlDisconnect
		_childTemp, typeSwitchError = ApduControlDisconnectParseWithBuffer(ctx, readBuffer)
	case controlType == 0x2: // ApduControlAck
		_childTemp, typeSwitchError = ApduControlAckParseWithBuffer(ctx, readBuffer)
	case controlType == 0x3: // ApduControlNack
		_childTemp, typeSwitchError = ApduControlNackParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [controlType=%v]", controlType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ApduControl")
	}
	_child = _childTemp.(ApduControlChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ApduControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControl")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ApduControl) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduControl, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ApduControl"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduControl")
	}

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType := uint8(child.GetControlType())
	_controlTypeErr := writeBuffer.WriteUint8("controlType", 2, (controlType))

	if _controlTypeErr != nil {
		return errors.Wrap(_controlTypeErr, "Error serializing 'controlType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduControl"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduControl")
	}
	return nil
}

func (m *_ApduControl) isApduControl() bool {
	return true
}

func (m *_ApduControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
