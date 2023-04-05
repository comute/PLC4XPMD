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

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetResult returns Result (property field)
	GetResult() uint8
	// GetReserved01 returns Reserved01 (property field)
	GetReserved01() uint8
}

// S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
	isS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse struct {
	*_S7PayloadUserDataItem
	Result     uint8
	Reserved01 uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetDataLength() uint16 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetResult() uint8 {
	return m.Result
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetReserved01() uint8 {
	return m.Reserved01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse factory function for _S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(result uint8, reserved01 uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse{
		Result:                 result,
		Reserved01:             reserved01,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(structType interface{}) S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParse(theBytes []byte, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse, error) {
	return S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	_result, _resultErr := readBuffer.ReadUint8("result", 8)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}
	result := _result

	// Simple Field (reserved01)
	_reserved01, _reserved01Err := readBuffer.ReadUint8("reserved01", 8)
	if _reserved01Err != nil {
		return nil, errors.Wrap(_reserved01Err, "Error parsing 'reserved01' field of S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}
	reserved01 := _reserved01

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		Result:                 result,
		Reserved01:             reserved01,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
		}

		// Simple Field (result)
		result := uint8(m.GetResult())
		_resultErr := writeBuffer.WriteUint8("result", 8, (result))
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (reserved01)
		reserved01 := uint8(m.GetReserved01())
		_reserved01Err := writeBuffer.WriteUint8("reserved01", 8, (reserved01))
		if _reserved01Err != nil {
			return errors.Wrap(_reserved01Err, "Error serializing 'reserved01' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) isS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
