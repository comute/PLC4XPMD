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

// IdentifyReplyCommandLogicalAssignment is the corresponding interface of IdentifyReplyCommandLogicalAssignment
type IdentifyReplyCommandLogicalAssignment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetLogicAssigment returns LogicAssigment (property field)
	GetLogicAssigment() []LogicAssignment
}

// IdentifyReplyCommandLogicalAssignmentExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandLogicalAssignment.
// This is useful for switch cases.
type IdentifyReplyCommandLogicalAssignmentExactly interface {
	IdentifyReplyCommandLogicalAssignment
	isIdentifyReplyCommandLogicalAssignment() bool
}

// _IdentifyReplyCommandLogicalAssignment is the data-structure of this message
type _IdentifyReplyCommandLogicalAssignment struct {
	*_IdentifyReplyCommand
	LogicAssigment []LogicAssignment
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetAttribute() Attribute {
	return Attribute_LogicalAssignment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandLogicalAssignment) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetLogicAssigment() []LogicAssignment {
	return m.LogicAssigment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandLogicalAssignment factory function for _IdentifyReplyCommandLogicalAssignment
func NewIdentifyReplyCommandLogicalAssignment(logicAssigment []LogicAssignment, numBytes uint8) *_IdentifyReplyCommandLogicalAssignment {
	_result := &_IdentifyReplyCommandLogicalAssignment{
		LogicAssigment:        logicAssigment,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandLogicalAssignment(structType any) IdentifyReplyCommandLogicalAssignment {
	if casted, ok := structType.(IdentifyReplyCommandLogicalAssignment); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandLogicalAssignment); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetTypeName() string {
	return "IdentifyReplyCommandLogicalAssignment"
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.LogicAssigment) > 0 {
		for _curItem, element := range m.LogicAssigment {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LogicAssigment), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandLogicalAssignmentParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandLogicalAssignment, error) {
	return IdentifyReplyCommandLogicalAssignmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandLogicalAssignmentParseWithBufferProducer(attribute Attribute, numBytes uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandLogicalAssignment, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandLogicalAssignment, error) {
		return IdentifyReplyCommandLogicalAssignmentParseWithBuffer(ctx, readBuffer, attribute, numBytes)
	}
}

func IdentifyReplyCommandLogicalAssignmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandLogicalAssignment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandLogicalAssignment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandLogicalAssignment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logicAssigment, err := ReadCountArrayField[LogicAssignment](ctx, "logicAssigment", ReadComplex[LogicAssignment](LogicAssignmentParseWithBuffer, readBuffer), uint64(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicAssigment' field"))
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandLogicalAssignment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandLogicalAssignment")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandLogicalAssignment{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		LogicAssigment: logicAssigment,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandLogicalAssignment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandLogicalAssignment")
		}

		if err := WriteComplexTypeArrayField(ctx, "logicAssigment", m.GetLogicAssigment(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'logicAssigment' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandLogicalAssignment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandLogicalAssignment")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandLogicalAssignment) isIdentifyReplyCommandLogicalAssignment() bool {
	return true
}

func (m *_IdentifyReplyCommandLogicalAssignment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
