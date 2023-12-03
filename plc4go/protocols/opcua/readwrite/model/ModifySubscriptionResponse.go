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

// ModifySubscriptionResponse is the corresponding interface of ModifySubscriptionResponse
type ModifySubscriptionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetRevisedPublishingInterval returns RevisedPublishingInterval (property field)
	GetRevisedPublishingInterval() float64
	// GetRevisedLifetimeCount returns RevisedLifetimeCount (property field)
	GetRevisedLifetimeCount() uint32
	// GetRevisedMaxKeepAliveCount returns RevisedMaxKeepAliveCount (property field)
	GetRevisedMaxKeepAliveCount() uint32
}

// ModifySubscriptionResponseExactly can be used when we want exactly this type and not a type which fulfills ModifySubscriptionResponse.
// This is useful for switch cases.
type ModifySubscriptionResponseExactly interface {
	ModifySubscriptionResponse
	isModifySubscriptionResponse() bool
}

// _ModifySubscriptionResponse is the data-structure of this message
type _ModifySubscriptionResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader            ExtensionObjectDefinition
	RevisedPublishingInterval float64
	RevisedLifetimeCount      uint32
	RevisedMaxKeepAliveCount  uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifySubscriptionResponse) GetIdentifier() string {
	return "796"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifySubscriptionResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ModifySubscriptionResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifySubscriptionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_ModifySubscriptionResponse) GetRevisedPublishingInterval() float64 {
	return m.RevisedPublishingInterval
}

func (m *_ModifySubscriptionResponse) GetRevisedLifetimeCount() uint32 {
	return m.RevisedLifetimeCount
}

func (m *_ModifySubscriptionResponse) GetRevisedMaxKeepAliveCount() uint32 {
	return m.RevisedMaxKeepAliveCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModifySubscriptionResponse factory function for _ModifySubscriptionResponse
func NewModifySubscriptionResponse(responseHeader ExtensionObjectDefinition, revisedPublishingInterval float64, revisedLifetimeCount uint32, revisedMaxKeepAliveCount uint32) *_ModifySubscriptionResponse {
	_result := &_ModifySubscriptionResponse{
		ResponseHeader:             responseHeader,
		RevisedPublishingInterval:  revisedPublishingInterval,
		RevisedLifetimeCount:       revisedLifetimeCount,
		RevisedMaxKeepAliveCount:   revisedMaxKeepAliveCount,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModifySubscriptionResponse(structType any) ModifySubscriptionResponse {
	if casted, ok := structType.(ModifySubscriptionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModifySubscriptionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModifySubscriptionResponse) GetTypeName() string {
	return "ModifySubscriptionResponse"
}

func (m *_ModifySubscriptionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (revisedPublishingInterval)
	lengthInBits += 64

	// Simple field (revisedLifetimeCount)
	lengthInBits += 32

	// Simple field (revisedMaxKeepAliveCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ModifySubscriptionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModifySubscriptionResponseParse(ctx context.Context, theBytes []byte, identifier string) (ModifySubscriptionResponse, error) {
	return ModifySubscriptionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ModifySubscriptionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ModifySubscriptionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModifySubscriptionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifySubscriptionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of ModifySubscriptionResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (revisedPublishingInterval)
	_revisedPublishingInterval, _revisedPublishingIntervalErr := readBuffer.ReadFloat64("revisedPublishingInterval", 64)
	if _revisedPublishingIntervalErr != nil {
		return nil, errors.Wrap(_revisedPublishingIntervalErr, "Error parsing 'revisedPublishingInterval' field of ModifySubscriptionResponse")
	}
	revisedPublishingInterval := _revisedPublishingInterval

	// Simple Field (revisedLifetimeCount)
	_revisedLifetimeCount, _revisedLifetimeCountErr := readBuffer.ReadUint32("revisedLifetimeCount", 32)
	if _revisedLifetimeCountErr != nil {
		return nil, errors.Wrap(_revisedLifetimeCountErr, "Error parsing 'revisedLifetimeCount' field of ModifySubscriptionResponse")
	}
	revisedLifetimeCount := _revisedLifetimeCount

	// Simple Field (revisedMaxKeepAliveCount)
	_revisedMaxKeepAliveCount, _revisedMaxKeepAliveCountErr := readBuffer.ReadUint32("revisedMaxKeepAliveCount", 32)
	if _revisedMaxKeepAliveCountErr != nil {
		return nil, errors.Wrap(_revisedMaxKeepAliveCountErr, "Error parsing 'revisedMaxKeepAliveCount' field of ModifySubscriptionResponse")
	}
	revisedMaxKeepAliveCount := _revisedMaxKeepAliveCount

	if closeErr := readBuffer.CloseContext("ModifySubscriptionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifySubscriptionResponse")
	}

	// Create a partially initialized instance
	_child := &_ModifySubscriptionResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		RevisedPublishingInterval:  revisedPublishingInterval,
		RevisedLifetimeCount:       revisedLifetimeCount,
		RevisedMaxKeepAliveCount:   revisedMaxKeepAliveCount,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ModifySubscriptionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifySubscriptionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifySubscriptionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifySubscriptionResponse")
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

		// Simple Field (revisedPublishingInterval)
		revisedPublishingInterval := float64(m.GetRevisedPublishingInterval())
		_revisedPublishingIntervalErr := writeBuffer.WriteFloat64("revisedPublishingInterval", 64, (revisedPublishingInterval))
		if _revisedPublishingIntervalErr != nil {
			return errors.Wrap(_revisedPublishingIntervalErr, "Error serializing 'revisedPublishingInterval' field")
		}

		// Simple Field (revisedLifetimeCount)
		revisedLifetimeCount := uint32(m.GetRevisedLifetimeCount())
		_revisedLifetimeCountErr := writeBuffer.WriteUint32("revisedLifetimeCount", 32, uint32((revisedLifetimeCount)))
		if _revisedLifetimeCountErr != nil {
			return errors.Wrap(_revisedLifetimeCountErr, "Error serializing 'revisedLifetimeCount' field")
		}

		// Simple Field (revisedMaxKeepAliveCount)
		revisedMaxKeepAliveCount := uint32(m.GetRevisedMaxKeepAliveCount())
		_revisedMaxKeepAliveCountErr := writeBuffer.WriteUint32("revisedMaxKeepAliveCount", 32, uint32((revisedMaxKeepAliveCount)))
		if _revisedMaxKeepAliveCountErr != nil {
			return errors.Wrap(_revisedMaxKeepAliveCountErr, "Error serializing 'revisedMaxKeepAliveCount' field")
		}

		if popErr := writeBuffer.PopContext("ModifySubscriptionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifySubscriptionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModifySubscriptionResponse) isModifySubscriptionResponse() bool {
	return true
}

func (m *_ModifySubscriptionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
