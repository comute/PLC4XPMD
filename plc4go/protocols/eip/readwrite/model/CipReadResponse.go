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

// CipReadResponse is the corresponding interface of CipReadResponse
type CipReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetData returns Data (property field)
	GetData() CIPData
}

// CipReadResponseExactly can be used when we want exactly this type and not a type which fulfills CipReadResponse.
// This is useful for switch cases.
type CipReadResponseExactly interface {
	CipReadResponse
	isCipReadResponse() bool
}

// _CipReadResponse is the data-structure of this message
type _CipReadResponse struct {
	*_CipService
	Status    uint8
	ExtStatus uint8
	Data      CIPData
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipReadResponse) GetService() uint8 {
	return 0x4C
}

func (m *_CipReadResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipReadResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipReadResponse) InitializeParent(parent CipService) {}

func (m *_CipReadResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipReadResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipReadResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_CipReadResponse) GetData() CIPData {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipReadResponse factory function for _CipReadResponse
func NewCipReadResponse(status uint8, extStatus uint8, data CIPData, serviceLen uint16) *_CipReadResponse {
	_result := &_CipReadResponse{
		Status:      status,
		ExtStatus:   extStatus,
		Data:        data,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipReadResponse(structType any) CipReadResponse {
	if casted, ok := structType.(CipReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipReadResponse) GetTypeName() string {
	return "CipReadResponse"
}

func (m *_CipReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Optional Field (data)
	if m.Data != nil {
		lengthInBits += m.Data.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CipReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipReadResponseParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipReadResponse, error) {
	return CipReadResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipReadResponseParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CipReadResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CipReadResponse, error) {
		return CipReadResponseParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func CipReadResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}

	extStatus, err := ReadSimpleField(ctx, "extStatus", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extStatus' field"))
	}

	_data, err := ReadOptionalField[CIPData](ctx, "data", ReadComplex[CIPData](CIPDataParseWithBufferProducer((uint16)(uint16(serviceLen)-uint16(uint16(4)))), readBuffer), bool(((serviceLen)-(4)) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	var data CIPData
	if _data != nil {
		data = *_data
	}

	if closeErr := readBuffer.CloseContext("CipReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipReadResponse")
	}

	// Create a partially initialized instance
	_child := &_CipReadResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:         status,
		ExtStatus:      extStatus,
		Data:           data,
		reservedField0: reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipReadResponse")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.GetStatus())
		_statusErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("status", 8, uint8((status)))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.GetExtStatus())
		_extStatusErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("extStatus", 8, uint8((extStatus)))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		// Optional Field (data) (Can be skipped, if the value is null)
		var data CIPData = nil
		if m.GetData() != nil {
			if pushErr := writeBuffer.PushContext("data"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for data")
			}
			data = m.GetData()
			_dataErr := writeBuffer.WriteSerializable(ctx, data)
			if popErr := writeBuffer.PopContext("data"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for data")
			}
			if _dataErr != nil {
				return errors.Wrap(_dataErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("CipReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipReadResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipReadResponse) isCipReadResponse() bool {
	return true
}

func (m *_CipReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
