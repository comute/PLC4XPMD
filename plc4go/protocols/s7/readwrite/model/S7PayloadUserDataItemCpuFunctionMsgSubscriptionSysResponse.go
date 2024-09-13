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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
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
	// IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse()
}

// _S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse struct {
	S7PayloadUserDataItemContract
	Result     uint8
	Reserved01 uint8
}

var _ S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
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
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(result uint8, reserved01 uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Result:                        result,
		Reserved01:                    reserved01,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(structType any) S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse {
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
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadSimpleField(ctx, "result", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}
	m.Result = result

	reserved01, err := ReadSimpleField(ctx, "reserved01", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reserved01' field"))
	}
	m.Reserved01 = reserved01

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
	}

	return m, nil
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
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
		}

		if err := WriteSimpleField[uint8](ctx, "result", m.GetResult(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'result' field")
		}

		if err := WriteSimpleField[uint8](ctx, "reserved01", m.GetReserved01(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved01' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) IsS7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse() {
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
