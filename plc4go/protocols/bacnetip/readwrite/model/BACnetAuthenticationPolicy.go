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

// BACnetAuthenticationPolicy is the data-structure of this message
type BACnetAuthenticationPolicy struct {
	Policy        *BACnetAuthenticationPolicyList
	OrderEnforced *BACnetContextTagBoolean
	Timeout       *BACnetContextTagUnsignedInteger
}

// IBACnetAuthenticationPolicy is the corresponding interface of BACnetAuthenticationPolicy
type IBACnetAuthenticationPolicy interface {
	// GetPolicy returns Policy (property field)
	GetPolicy() *BACnetAuthenticationPolicyList
	// GetOrderEnforced returns OrderEnforced (property field)
	GetOrderEnforced() *BACnetContextTagBoolean
	// GetTimeout returns Timeout (property field)
	GetTimeout() *BACnetContextTagUnsignedInteger
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

func (m *BACnetAuthenticationPolicy) GetPolicy() *BACnetAuthenticationPolicyList {
	return m.Policy
}

func (m *BACnetAuthenticationPolicy) GetOrderEnforced() *BACnetContextTagBoolean {
	return m.OrderEnforced
}

func (m *BACnetAuthenticationPolicy) GetTimeout() *BACnetContextTagUnsignedInteger {
	return m.Timeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationPolicy factory function for BACnetAuthenticationPolicy
func NewBACnetAuthenticationPolicy(policy *BACnetAuthenticationPolicyList, orderEnforced *BACnetContextTagBoolean, timeout *BACnetContextTagUnsignedInteger) *BACnetAuthenticationPolicy {
	return &BACnetAuthenticationPolicy{Policy: policy, OrderEnforced: orderEnforced, Timeout: timeout}
}

func CastBACnetAuthenticationPolicy(structType interface{}) *BACnetAuthenticationPolicy {
	if casted, ok := structType.(BACnetAuthenticationPolicy); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicy); ok {
		return casted
	}
	return nil
}

func (m *BACnetAuthenticationPolicy) GetTypeName() string {
	return "BACnetAuthenticationPolicy"
}

func (m *BACnetAuthenticationPolicy) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAuthenticationPolicy) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (policy)
	lengthInBits += m.Policy.GetLengthInBits()

	// Simple field (orderEnforced)
	lengthInBits += m.OrderEnforced.GetLengthInBits()

	// Simple field (timeout)
	lengthInBits += m.Timeout.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAuthenticationPolicy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAuthenticationPolicyParse(readBuffer utils.ReadBuffer) (*BACnetAuthenticationPolicy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (policy)
	if pullErr := readBuffer.PullContext("policy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for policy")
	}
	_policy, _policyErr := BACnetAuthenticationPolicyListParse(readBuffer, uint8(uint8(0)))
	if _policyErr != nil {
		return nil, errors.Wrap(_policyErr, "Error parsing 'policy' field")
	}
	policy := CastBACnetAuthenticationPolicyList(_policy)
	if closeErr := readBuffer.CloseContext("policy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for policy")
	}

	// Simple Field (orderEnforced)
	if pullErr := readBuffer.PullContext("orderEnforced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for orderEnforced")
	}
	_orderEnforced, _orderEnforcedErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _orderEnforcedErr != nil {
		return nil, errors.Wrap(_orderEnforcedErr, "Error parsing 'orderEnforced' field")
	}
	orderEnforced := CastBACnetContextTagBoolean(_orderEnforced)
	if closeErr := readBuffer.CloseContext("orderEnforced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for orderEnforced")
	}

	// Simple Field (timeout)
	if pullErr := readBuffer.PullContext("timeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeout")
	}
	_timeout, _timeoutErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeoutErr != nil {
		return nil, errors.Wrap(_timeoutErr, "Error parsing 'timeout' field")
	}
	timeout := CastBACnetContextTagUnsignedInteger(_timeout)
	if closeErr := readBuffer.CloseContext("timeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeout")
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicy")
	}

	// Create the instance
	return NewBACnetAuthenticationPolicy(policy, orderEnforced, timeout), nil
}

func (m *BACnetAuthenticationPolicy) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicy")
	}

	// Simple Field (policy)
	if pushErr := writeBuffer.PushContext("policy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for policy")
	}
	_policyErr := writeBuffer.WriteSerializable(m.Policy)
	if popErr := writeBuffer.PopContext("policy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for policy")
	}
	if _policyErr != nil {
		return errors.Wrap(_policyErr, "Error serializing 'policy' field")
	}

	// Simple Field (orderEnforced)
	if pushErr := writeBuffer.PushContext("orderEnforced"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for orderEnforced")
	}
	_orderEnforcedErr := writeBuffer.WriteSerializable(m.OrderEnforced)
	if popErr := writeBuffer.PopContext("orderEnforced"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for orderEnforced")
	}
	if _orderEnforcedErr != nil {
		return errors.Wrap(_orderEnforcedErr, "Error serializing 'orderEnforced' field")
	}

	// Simple Field (timeout)
	if pushErr := writeBuffer.PushContext("timeout"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeout")
	}
	_timeoutErr := writeBuffer.WriteSerializable(m.Timeout)
	if popErr := writeBuffer.PopContext("timeout"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeout")
	}
	if _timeoutErr != nil {
		return errors.Wrap(_timeoutErr, "Error serializing 'timeout' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicy")
	}
	return nil
}

func (m *BACnetAuthenticationPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
