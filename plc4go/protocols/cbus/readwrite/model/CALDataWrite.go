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

// CALDataWrite is the corresponding interface of CALDataWrite
type CALDataWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCode returns Code (property field)
	GetCode() byte
	// GetParameterValue returns ParameterValue (property field)
	GetParameterValue() ParameterValue
}

// CALDataWriteExactly can be used when we want exactly this type and not a type which fulfills CALDataWrite.
// This is useful for switch cases.
type CALDataWriteExactly interface {
	CALDataWrite
	isCALDataWrite() bool
}

// _CALDataWrite is the data-structure of this message
type _CALDataWrite struct {
	*_CALData
	ParamNo        Parameter
	Code           byte
	ParameterValue ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataWrite) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataWrite) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataWrite) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataWrite) GetCode() byte {
	return m.Code
}

func (m *_CALDataWrite) GetParameterValue() ParameterValue {
	return m.ParameterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataWrite factory function for _CALDataWrite
func NewCALDataWrite(paramNo Parameter, code byte, parameterValue ParameterValue, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataWrite {
	_result := &_CALDataWrite{
		ParamNo:        paramNo,
		Code:           code,
		ParameterValue: parameterValue,
		_CALData:       NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataWrite(structType any) CALDataWrite {
	if casted, ok := structType.(CALDataWrite); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataWrite); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataWrite) GetTypeName() string {
	return "CALDataWrite"
}

func (m *_CALDataWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	// Simple field (parameterValue)
	lengthInBits += m.ParameterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALDataWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALDataWriteParse(ctx context.Context, theBytes []byte, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (CALDataWrite, error) {
	return CALDataWriteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), commandTypeContainer, requestContext)
}

func CALDataWriteParseWithBufferProducer(commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (CALDataWrite, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CALDataWrite, error) {
		return CALDataWriteParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	}
}

func CALDataWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (CALDataWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	paramNo, err := ReadEnumField[Parameter](ctx, "paramNo", "Parameter", ReadEnum(ParameterByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paramNo' field"))
	}

	code, err := ReadSimpleField(ctx, "code", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'code' field"))
	}

	parameterValue, err := ReadSimpleField[ParameterValue](ctx, "parameterValue", ReadComplex[ParameterValue](ParameterValueParseWithBufferProducer[ParameterValue]((ParameterType)(paramNo.ParameterType()), (uint8)(uint8(commandTypeContainer.NumBytes())-uint8(uint8(2)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterValue' field"))
	}

	if closeErr := readBuffer.CloseContext("CALDataWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataWrite")
	}

	// Create a partially initialized instance
	_child := &_CALDataWrite{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		ParamNo:        paramNo,
		Code:           code,
		ParameterValue: parameterValue,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataWrite")
		}

		// Simple Field (paramNo)
		if pushErr := writeBuffer.PushContext("paramNo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for paramNo")
		}
		_paramNoErr := writeBuffer.WriteSerializable(ctx, m.GetParamNo())
		if popErr := writeBuffer.PopContext("paramNo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for paramNo")
		}
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (code)
		code := byte(m.GetCode())
		_codeErr := /*TODO: migrate me*/ writeBuffer.WriteByte("code", (code))
		if _codeErr != nil {
			return errors.Wrap(_codeErr, "Error serializing 'code' field")
		}

		// Simple Field (parameterValue)
		if pushErr := writeBuffer.PushContext("parameterValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameterValue")
		}
		_parameterValueErr := writeBuffer.WriteSerializable(ctx, m.GetParameterValue())
		if popErr := writeBuffer.PopContext("parameterValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameterValue")
		}
		if _parameterValueErr != nil {
			return errors.Wrap(_parameterValueErr, "Error serializing 'parameterValue' field")
		}

		if popErr := writeBuffer.PopContext("CALDataWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataWrite")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataWrite) isCALDataWrite() bool {
	return true
}

func (m *_CALDataWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
