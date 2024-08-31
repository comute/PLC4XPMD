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

// Constant values.
const S7MessageObjectRequest_VARIABLESPEC uint8 = 0x12
const S7MessageObjectRequest_LENGTH uint8 = 0x08

// S7MessageObjectRequest is the corresponding interface of S7MessageObjectRequest
type S7MessageObjectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7DataAlarmMessage
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetQueryType returns QueryType (property field)
	GetQueryType() QueryType
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
}

// S7MessageObjectRequestExactly can be used when we want exactly this type and not a type which fulfills S7MessageObjectRequest.
// This is useful for switch cases.
type S7MessageObjectRequestExactly interface {
	S7MessageObjectRequest
	isS7MessageObjectRequest() bool
}

// _S7MessageObjectRequest is the data-structure of this message
type _S7MessageObjectRequest struct {
	*_S7DataAlarmMessage
	SyntaxId  SyntaxIdType
	QueryType QueryType
	AlarmType AlarmType
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageObjectRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageObjectRequest) InitializeParent(parent S7DataAlarmMessage) {}

func (m *_S7MessageObjectRequest) GetParent() S7DataAlarmMessage {
	return m._S7DataAlarmMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageObjectRequest) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *_S7MessageObjectRequest) GetQueryType() QueryType {
	return m.QueryType
}

func (m *_S7MessageObjectRequest) GetAlarmType() AlarmType {
	return m.AlarmType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7MessageObjectRequest) GetVariableSpec() uint8 {
	return S7MessageObjectRequest_VARIABLESPEC
}

func (m *_S7MessageObjectRequest) GetLength() uint8 {
	return S7MessageObjectRequest_LENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7MessageObjectRequest factory function for _S7MessageObjectRequest
func NewS7MessageObjectRequest(syntaxId SyntaxIdType, queryType QueryType, alarmType AlarmType) *_S7MessageObjectRequest {
	_result := &_S7MessageObjectRequest{
		SyntaxId:            syntaxId,
		QueryType:           queryType,
		AlarmType:           alarmType,
		_S7DataAlarmMessage: NewS7DataAlarmMessage(),
	}
	_result._S7DataAlarmMessage._S7DataAlarmMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageObjectRequest(structType any) S7MessageObjectRequest {
	if casted, ok := structType.(S7MessageObjectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageObjectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageObjectRequest) GetTypeName() string {
	return "S7MessageObjectRequest"
}

func (m *_S7MessageObjectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (variableSpec)
	lengthInBits += 8

	// Const Field (length)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (queryType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageObjectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageObjectRequestParse(ctx context.Context, theBytes []byte, cpuFunctionType uint8) (S7MessageObjectRequest, error) {
	return S7MessageObjectRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionType)
}

func S7MessageObjectRequestParseWithBufferProducer(cpuFunctionType uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageObjectRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7MessageObjectRequest, error) {
		return S7MessageObjectRequestParseWithBuffer(ctx, readBuffer, cpuFunctionType)
	}
}

func S7MessageObjectRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8) (S7MessageObjectRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageObjectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageObjectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), S7MessageObjectRequest_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	length, err := ReadConstField[uint8](ctx, "length", ReadUnsignedByte(readBuffer, uint8(8)), S7MessageObjectRequest_LENGTH)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	syntaxId, err := ReadEnumField[SyntaxIdType](ctx, "syntaxId", "SyntaxIdType", ReadEnum(SyntaxIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	queryType, err := ReadEnumField[QueryType](ctx, "queryType", "QueryType", ReadEnum(QueryTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queryType' field"))
	}

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x34))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	alarmType, err := ReadEnumField[AlarmType](ctx, "alarmType", "AlarmType", ReadEnum(AlarmTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmType' field"))
	}

	if closeErr := readBuffer.CloseContext("S7MessageObjectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageObjectRequest")
	}

	// Create a partially initialized instance
	_child := &_S7MessageObjectRequest{
		_S7DataAlarmMessage: &_S7DataAlarmMessage{},
		SyntaxId:            syntaxId,
		QueryType:           queryType,
		AlarmType:           alarmType,
		reservedField0:      reservedField0,
		reservedField1:      reservedField1,
	}
	_child._S7DataAlarmMessage._S7DataAlarmMessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageObjectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageObjectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageObjectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageObjectRequest")
		}

		// Const Field (variableSpec)
		_variableSpecErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("variableSpec", 8, uint8(0x12))
		if _variableSpecErr != nil {
			return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
		}

		// Const Field (length)
		_lengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint8("length", 8, uint8(0x08))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (syntaxId)
		if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for syntaxId")
		}
		_syntaxIdErr := writeBuffer.WriteSerializable(ctx, m.GetSyntaxId())
		if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for syntaxId")
		}
		if _syntaxIdErr != nil {
			return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
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

		// Simple Field (queryType)
		if pushErr := writeBuffer.PushContext("queryType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for queryType")
		}
		_queryTypeErr := writeBuffer.WriteSerializable(ctx, m.GetQueryType())
		if popErr := writeBuffer.PopContext("queryType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for queryType")
		}
		if _queryTypeErr != nil {
			return errors.Wrap(_queryTypeErr, "Error serializing 'queryType' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x34)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x34),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (alarmType)
		if pushErr := writeBuffer.PushContext("alarmType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmType")
		}
		_alarmTypeErr := writeBuffer.WriteSerializable(ctx, m.GetAlarmType())
		if popErr := writeBuffer.PopContext("alarmType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmType")
		}
		if _alarmTypeErr != nil {
			return errors.Wrap(_alarmTypeErr, "Error serializing 'alarmType' field")
		}

		if popErr := writeBuffer.PopContext("S7MessageObjectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageObjectRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageObjectRequest) isS7MessageObjectRequest() bool {
	return true
}

func (m *_S7MessageObjectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
