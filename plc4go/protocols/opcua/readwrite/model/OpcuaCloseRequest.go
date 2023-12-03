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

// OpcuaCloseRequest is the corresponding interface of OpcuaCloseRequest
type OpcuaCloseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetChunk returns Chunk (property field)
	GetChunk() string
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetSecureTokenId returns SecureTokenId (property field)
	GetSecureTokenId() int32
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() int32
	// GetRequestId returns RequestId (property field)
	GetRequestId() int32
	// GetMessage returns Message (property field)
	GetMessage() ExtensionObject
}

// OpcuaCloseRequestExactly can be used when we want exactly this type and not a type which fulfills OpcuaCloseRequest.
// This is useful for switch cases.
type OpcuaCloseRequestExactly interface {
	OpcuaCloseRequest
	isOpcuaCloseRequest() bool
}

// _OpcuaCloseRequest is the data-structure of this message
type _OpcuaCloseRequest struct {
	*_MessagePDU
	Chunk           string
	SecureChannelId int32
	SecureTokenId   int32
	SequenceNumber  int32
	RequestId       int32
	Message         ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaCloseRequest) GetMessageType() string {
	return "CLO"
}

func (m *_OpcuaCloseRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaCloseRequest) InitializeParent(parent MessagePDU) {}

func (m *_OpcuaCloseRequest) GetParent() MessagePDU {
	return m._MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaCloseRequest) GetChunk() string {
	return m.Chunk
}

func (m *_OpcuaCloseRequest) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpcuaCloseRequest) GetSecureTokenId() int32 {
	return m.SecureTokenId
}

func (m *_OpcuaCloseRequest) GetSequenceNumber() int32 {
	return m.SequenceNumber
}

func (m *_OpcuaCloseRequest) GetRequestId() int32 {
	return m.RequestId
}

func (m *_OpcuaCloseRequest) GetMessage() ExtensionObject {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaCloseRequest factory function for _OpcuaCloseRequest
func NewOpcuaCloseRequest(chunk string, secureChannelId int32, secureTokenId int32, sequenceNumber int32, requestId int32, message ExtensionObject) *_OpcuaCloseRequest {
	_result := &_OpcuaCloseRequest{
		Chunk:           chunk,
		SecureChannelId: secureChannelId,
		SecureTokenId:   secureTokenId,
		SequenceNumber:  sequenceNumber,
		RequestId:       requestId,
		Message:         message,
		_MessagePDU:     NewMessagePDU(),
	}
	_result._MessagePDU._MessagePDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaCloseRequest(structType any) OpcuaCloseRequest {
	if casted, ok := structType.(OpcuaCloseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaCloseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaCloseRequest) GetTypeName() string {
	return "OpcuaCloseRequest"
}

func (m *_OpcuaCloseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (messageSize)
	lengthInBits += 32

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (secureTokenId)
	lengthInBits += 32

	// Simple field (sequenceNumber)
	lengthInBits += 32

	// Simple field (requestId)
	lengthInBits += 32

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaCloseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaCloseRequestParse(ctx context.Context, theBytes []byte, response bool) (OpcuaCloseRequest, error) {
	return OpcuaCloseRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func OpcuaCloseRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (OpcuaCloseRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaCloseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaCloseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (chunk)
	_chunk, _chunkErr := readBuffer.ReadString("chunk", uint32(8), "UTF-8")
	if _chunkErr != nil {
		return nil, errors.Wrap(_chunkErr, "Error parsing 'chunk' field of OpcuaCloseRequest")
	}
	chunk := _chunk

	// Implicit Field (messageSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadInt32("messageSize", 32)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field of OpcuaCloseRequest")
	}

	// Simple Field (secureChannelId)
	_secureChannelId, _secureChannelIdErr := readBuffer.ReadInt32("secureChannelId", 32)
	if _secureChannelIdErr != nil {
		return nil, errors.Wrap(_secureChannelIdErr, "Error parsing 'secureChannelId' field of OpcuaCloseRequest")
	}
	secureChannelId := _secureChannelId

	// Simple Field (secureTokenId)
	_secureTokenId, _secureTokenIdErr := readBuffer.ReadInt32("secureTokenId", 32)
	if _secureTokenIdErr != nil {
		return nil, errors.Wrap(_secureTokenIdErr, "Error parsing 'secureTokenId' field of OpcuaCloseRequest")
	}
	secureTokenId := _secureTokenId

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadInt32("sequenceNumber", 32)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of OpcuaCloseRequest")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (requestId)
	_requestId, _requestIdErr := readBuffer.ReadInt32("requestId", 32)
	if _requestIdErr != nil {
		return nil, errors.Wrap(_requestIdErr, "Error parsing 'requestId' field of OpcuaCloseRequest")
	}
	requestId := _requestId

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
	_message, _messageErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(false)))
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of OpcuaCloseRequest")
	}
	message := _message.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	if closeErr := readBuffer.CloseContext("OpcuaCloseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaCloseRequest")
	}

	// Create a partially initialized instance
	_child := &_OpcuaCloseRequest{
		_MessagePDU:     &_MessagePDU{},
		Chunk:           chunk,
		SecureChannelId: secureChannelId,
		SecureTokenId:   secureTokenId,
		SequenceNumber:  sequenceNumber,
		RequestId:       requestId,
		Message:         message,
	}
	_child._MessagePDU._MessagePDUChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaCloseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaCloseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaCloseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaCloseRequest")
		}

		// Simple Field (chunk)
		chunk := string(m.GetChunk())
		_chunkErr := writeBuffer.WriteString("chunk", uint32(8), "UTF-8", (chunk))
		if _chunkErr != nil {
			return errors.Wrap(_chunkErr, "Error serializing 'chunk' field")
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := int32(int32(m.GetLengthInBytes(ctx)))
		_messageSizeErr := writeBuffer.WriteInt32("messageSize", 32, int32((messageSize)))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (secureChannelId)
		secureChannelId := int32(m.GetSecureChannelId())
		_secureChannelIdErr := writeBuffer.WriteInt32("secureChannelId", 32, int32((secureChannelId)))
		if _secureChannelIdErr != nil {
			return errors.Wrap(_secureChannelIdErr, "Error serializing 'secureChannelId' field")
		}

		// Simple Field (secureTokenId)
		secureTokenId := int32(m.GetSecureTokenId())
		_secureTokenIdErr := writeBuffer.WriteInt32("secureTokenId", 32, int32((secureTokenId)))
		if _secureTokenIdErr != nil {
			return errors.Wrap(_secureTokenIdErr, "Error serializing 'secureTokenId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := int32(m.GetSequenceNumber())
		_sequenceNumberErr := writeBuffer.WriteInt32("sequenceNumber", 32, int32((sequenceNumber)))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (requestId)
		requestId := int32(m.GetRequestId())
		_requestIdErr := writeBuffer.WriteInt32("requestId", 32, int32((requestId)))
		if _requestIdErr != nil {
			return errors.Wrap(_requestIdErr, "Error serializing 'requestId' field")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(ctx, m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaCloseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaCloseRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaCloseRequest) isOpcuaCloseRequest() bool {
	return true
}

func (m *_OpcuaCloseRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
