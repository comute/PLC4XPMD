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

// Reply is the corresponding interface of Reply
type Reply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() byte
}

// ReplyExactly can be used when we want exactly this type and not a type which fulfills Reply.
// This is useful for switch cases.
type ReplyExactly interface {
	Reply
	isReply() bool
}

// _Reply is the data-structure of this message
type _Reply struct {
	_ReplyChildRequirements
	PeekedByte byte

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

type _ReplyChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedByte() byte
}

type ReplyParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Reply, serializeChildFunction func() error) error
	GetTypeName() string
}

type ReplyChild interface {
	utils.Serializable
	InitializeParent(parent Reply, peekedByte byte)
	GetParent() *Reply

	GetTypeName() string
	Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Reply) GetPeekedByte() byte {
	return m.PeekedByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReply factory function for _Reply
func NewReply(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_Reply {
	return &_Reply{PeekedByte: peekedByte, CBusOptions: cBusOptions, RequestContext: requestContext}
}

// Deprecated: use the interface for direct cast
func CastReply(structType any) Reply {
	if casted, ok := structType.(Reply); ok {
		return casted
	}
	if casted, ok := structType.(*Reply); ok {
		return *casted
	}
	return nil
}

func (m *_Reply) GetTypeName() string {
	return "Reply"
}

func (m *_Reply) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Reply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReplyParse(theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (Reply, error) {
	return ReplyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (Reply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Reply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Reply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedByte)
	currentPos = positionAware.GetPos()
	peekedByte, _err := readBuffer.ReadByte("peekedByte")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekedByte' field of Reply")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ReplyChildSerializeRequirement interface {
		Reply
		InitializeParent(Reply, byte)
		GetParent() Reply
	}
	var _childTemp any
	var _child ReplyChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedByte == 0x2B: // PowerUpReply
		_childTemp, typeSwitchError = PowerUpReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	case peekedByte == 0x3D: // ParameterChangeReply
		_childTemp, typeSwitchError = ParameterChangeReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	case 0 == 0: // ReplyEncodedReply
		_childTemp, typeSwitchError = ReplyEncodedReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedByte=%v]", peekedByte)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of Reply")
	}
	_child = _childTemp.(ReplyChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("Reply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Reply")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedByte)
	return _child, nil
}

func (pm *_Reply) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Reply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Reply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Reply")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Reply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Reply")
	}
	return nil
}

////
// Arguments Getter

func (m *_Reply) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_Reply) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_Reply) isReply() bool {
	return true
}

func (m *_Reply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
