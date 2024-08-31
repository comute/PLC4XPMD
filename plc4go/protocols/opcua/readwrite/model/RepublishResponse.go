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

// RepublishResponse is the corresponding interface of RepublishResponse
type RepublishResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNotificationMessage returns NotificationMessage (property field)
	GetNotificationMessage() ExtensionObjectDefinition
}

// RepublishResponseExactly can be used when we want exactly this type and not a type which fulfills RepublishResponse.
// This is useful for switch cases.
type RepublishResponseExactly interface {
	RepublishResponse
	isRepublishResponse() bool
}

// _RepublishResponse is the data-structure of this message
type _RepublishResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader      ExtensionObjectDefinition
	NotificationMessage ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RepublishResponse) GetIdentifier() string {
	return "835"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RepublishResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RepublishResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RepublishResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_RepublishResponse) GetNotificationMessage() ExtensionObjectDefinition {
	return m.NotificationMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRepublishResponse factory function for _RepublishResponse
func NewRepublishResponse(responseHeader ExtensionObjectDefinition, notificationMessage ExtensionObjectDefinition) *_RepublishResponse {
	_result := &_RepublishResponse{
		ResponseHeader:             responseHeader,
		NotificationMessage:        notificationMessage,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRepublishResponse(structType any) RepublishResponse {
	if casted, ok := structType.(RepublishResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RepublishResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RepublishResponse) GetTypeName() string {
	return "RepublishResponse"
}

func (m *_RepublishResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (notificationMessage)
	lengthInBits += m.NotificationMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RepublishResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RepublishResponseParse(ctx context.Context, theBytes []byte, identifier string) (RepublishResponse, error) {
	return RepublishResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RepublishResponseParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (RepublishResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RepublishResponse, error) {
		return RepublishResponseParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func RepublishResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RepublishResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RepublishResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RepublishResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}

	notificationMessage, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "notificationMessage", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("805")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationMessage' field"))
	}

	if closeErr := readBuffer.CloseContext("RepublishResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RepublishResponse")
	}

	// Create a partially initialized instance
	_child := &_RepublishResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NotificationMessage:        notificationMessage,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RepublishResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RepublishResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RepublishResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RepublishResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (notificationMessage)
		if pushErr := writeBuffer.PushContext("notificationMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notificationMessage")
		}
		_notificationMessageErr := writeBuffer.WriteSerializable(ctx, m.GetNotificationMessage())
		if popErr := writeBuffer.PopContext("notificationMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notificationMessage")
		}
		if _notificationMessageErr != nil {
			return errors.Wrap(_notificationMessageErr, "Error serializing 'notificationMessage' field")
		}

		if popErr := writeBuffer.PopContext("RepublishResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RepublishResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RepublishResponse) isRepublishResponse() bool {
	return true
}

func (m *_RepublishResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
