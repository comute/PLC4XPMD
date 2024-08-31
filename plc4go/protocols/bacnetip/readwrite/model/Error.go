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

// Error is the corresponding interface of Error
type Error interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetErrorClass returns ErrorClass (property field)
	GetErrorClass() ErrorClassTagged
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() ErrorCodeTagged
}

// ErrorExactly can be used when we want exactly this type and not a type which fulfills Error.
// This is useful for switch cases.
type ErrorExactly interface {
	Error
	isError() bool
}

// _Error is the data-structure of this message
type _Error struct {
	ErrorClass ErrorClassTagged
	ErrorCode  ErrorCodeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Error) GetErrorClass() ErrorClassTagged {
	return m.ErrorClass
}

func (m *_Error) GetErrorCode() ErrorCodeTagged {
	return m.ErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewError factory function for _Error
func NewError(errorClass ErrorClassTagged, errorCode ErrorCodeTagged) *_Error {
	return &_Error{ErrorClass: errorClass, ErrorCode: errorCode}
}

// Deprecated: use the interface for direct cast
func CastError(structType any) Error {
	if casted, ok := structType.(Error); ok {
		return casted
	}
	if casted, ok := structType.(*Error); ok {
		return *casted
	}
	return nil
}

func (m *_Error) GetTypeName() string {
	return "Error"
}

func (m *_Error) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (errorClass)
	lengthInBits += m.ErrorClass.GetLengthInBits(ctx)

	// Simple field (errorCode)
	lengthInBits += m.ErrorCode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Error) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorParse(ctx context.Context, theBytes []byte) (Error, error) {
	return ErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
		return ErrorParseWithBuffer(ctx, readBuffer)
	}
}

func ErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Error, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Error")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorClass, err := ReadSimpleField[ErrorClassTagged](ctx, "errorClass", ReadComplex[ErrorClassTagged](ErrorClassTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorClass' field"))
	}

	errorCode, err := ReadSimpleField[ErrorCodeTagged](ctx, "errorCode", ReadComplex[ErrorCodeTagged](ErrorCodeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorCode' field"))
	}

	if closeErr := readBuffer.CloseContext("Error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Error")
	}

	// Create the instance
	return &_Error{
		ErrorClass: errorClass,
		ErrorCode:  errorCode,
	}, nil
}

func (m *_Error) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Error) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Error"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Error")
	}

	// Simple Field (errorClass)
	if pushErr := writeBuffer.PushContext("errorClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for errorClass")
	}
	_errorClassErr := writeBuffer.WriteSerializable(ctx, m.GetErrorClass())
	if popErr := writeBuffer.PopContext("errorClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for errorClass")
	}
	if _errorClassErr != nil {
		return errors.Wrap(_errorClassErr, "Error serializing 'errorClass' field")
	}

	// Simple Field (errorCode)
	if pushErr := writeBuffer.PushContext("errorCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for errorCode")
	}
	_errorCodeErr := writeBuffer.WriteSerializable(ctx, m.GetErrorCode())
	if popErr := writeBuffer.PopContext("errorCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for errorCode")
	}
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	if popErr := writeBuffer.PopContext("Error"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Error")
	}
	return nil
}

func (m *_Error) isError() bool {
	return true
}

func (m *_Error) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
