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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataAcknowledge is the corresponding interface of CALDataAcknowledge
type CALDataAcknowledge interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetCode returns Code (property field)
	GetCode() uint8
}

// CALDataAcknowledgeExactly can be used when we want exactly this type and not a type which fulfills CALDataAcknowledge.
// This is useful for switch cases.
type CALDataAcknowledgeExactly interface {
	CALDataAcknowledge
	isCALDataAcknowledge() bool
}

// _CALDataAcknowledge is the data-structure of this message
type _CALDataAcknowledge struct {
	*_CALData
	ParamNo Parameter
	Code    uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataAcknowledge) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataAcknowledge) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataAcknowledge) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataAcknowledge) GetCode() uint8 {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataAcknowledge factory function for _CALDataAcknowledge
func NewCALDataAcknowledge(paramNo Parameter, code uint8, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataAcknowledge {
	_result := &_CALDataAcknowledge{
		ParamNo:  paramNo,
		Code:     code,
		_CALData: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataAcknowledge(structType any) CALDataAcknowledge {
	if casted, ok := structType.(CALDataAcknowledge); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataAcknowledge); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataAcknowledge) GetTypeName() string {
	return "CALDataAcknowledge"
}

func (m *_CALDataAcknowledge) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (code)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataAcknowledge) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALDataAcknowledgeParse(ctx context.Context, theBytes []byte, requestContext RequestContext) (CALDataAcknowledge, error) {
	return CALDataAcknowledgeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), requestContext)
}

func CALDataAcknowledgeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, requestContext RequestContext) (CALDataAcknowledge, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CALDataAcknowledge"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataAcknowledge")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	if pullErr := readBuffer.PullContext("paramNo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for paramNo")
	}
	_paramNo, _paramNoErr := ParameterParseWithBuffer(ctx, readBuffer)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataAcknowledge")
	}
	paramNo := _paramNo
	if closeErr := readBuffer.CloseContext("paramNo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for paramNo")
	}

	// Simple Field (code)
	_code, _codeErr := readBuffer.ReadUint8("code", 8)
	if _codeErr != nil {
		return nil, errors.Wrap(_codeErr, "Error parsing 'code' field of CALDataAcknowledge")
	}
	code := _code

	if closeErr := readBuffer.CloseContext("CALDataAcknowledge"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataAcknowledge")
	}

	// Create a partially initialized instance
	_child := &_CALDataAcknowledge{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		ParamNo: paramNo,
		Code:    code,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataAcknowledge) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataAcknowledge) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataAcknowledge"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataAcknowledge")
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
		code := uint8(m.GetCode())
		_codeErr := writeBuffer.WriteUint8("code", 8, uint8((code)))
		if _codeErr != nil {
			return errors.Wrap(_codeErr, "Error serializing 'code' field")
		}

		if popErr := writeBuffer.PopContext("CALDataAcknowledge"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataAcknowledge")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataAcknowledge) isCALDataAcknowledge() bool {
	return true
}

func (m *_CALDataAcknowledge) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
