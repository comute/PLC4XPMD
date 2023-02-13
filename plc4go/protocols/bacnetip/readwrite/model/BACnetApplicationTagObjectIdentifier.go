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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetApplicationTagObjectIdentifier is the corresponding interface of BACnetApplicationTagObjectIdentifier
type BACnetApplicationTagObjectIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadObjectIdentifier
	// GetObjectType returns ObjectType (virtual field)
	GetObjectType() BACnetObjectType
	// GetInstanceNumber returns InstanceNumber (virtual field)
	GetInstanceNumber() uint32
}

// BACnetApplicationTagObjectIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagObjectIdentifier.
// This is useful for switch cases.
type BACnetApplicationTagObjectIdentifierExactly interface {
	BACnetApplicationTagObjectIdentifier
	isBACnetApplicationTagObjectIdentifier() bool
}

// _BACnetApplicationTagObjectIdentifier is the data-structure of this message
type _BACnetApplicationTagObjectIdentifier struct {
	*_BACnetApplicationTag
	Payload BACnetTagPayloadObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagObjectIdentifier) InitializeParent(parent BACnetApplicationTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetApplicationTagObjectIdentifier) GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagObjectIdentifier) GetPayload() BACnetTagPayloadObjectIdentifier {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagObjectIdentifier) GetObjectType() BACnetObjectType {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectType(m.GetPayload().GetObjectType())
}

func (m *_BACnetApplicationTagObjectIdentifier) GetInstanceNumber() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetPayload().GetInstanceNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagObjectIdentifier factory function for _BACnetApplicationTagObjectIdentifier
func NewBACnetApplicationTagObjectIdentifier(payload BACnetTagPayloadObjectIdentifier, header BACnetTagHeader) *_BACnetApplicationTagObjectIdentifier {
	_result := &_BACnetApplicationTagObjectIdentifier{
		Payload:               payload,
		_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagObjectIdentifier(structType interface{}) BACnetApplicationTagObjectIdentifier {
	if casted, ok := structType.(BACnetApplicationTagObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagObjectIdentifier) GetTypeName() string {
	return "BACnetApplicationTagObjectIdentifier"
}

func (m *_BACnetApplicationTagObjectIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagObjectIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetApplicationTagObjectIdentifierParse(theBytes []byte) (BACnetApplicationTagObjectIdentifier, error) {
	return BACnetApplicationTagObjectIdentifierParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetApplicationTagObjectIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTagObjectIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadObjectIdentifierParseWithBuffer(ctx, readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetApplicationTagObjectIdentifier")
	}
	payload := _payload.(BACnetTagPayloadObjectIdentifier)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_objectType := payload.GetObjectType()
	objectType := BACnetObjectType(_objectType)
	_ = objectType

	// Virtual field
	_instanceNumber := payload.GetInstanceNumber()
	instanceNumber := uint32(_instanceNumber)
	_ = instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagObjectIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagObjectIdentifier{
		_BACnetApplicationTag: &_BACnetApplicationTag{},
		Payload:               payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagObjectIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagObjectIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagObjectIdentifier")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _objectTypeErr := writeBuffer.WriteVirtual(ctx, "objectType", m.GetObjectType()); _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}
		// Virtual field
		if _instanceNumberErr := writeBuffer.WriteVirtual(ctx, "instanceNumber", m.GetInstanceNumber()); _instanceNumberErr != nil {
			return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagObjectIdentifier")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagObjectIdentifier) isBACnetApplicationTagObjectIdentifier() bool {
	return true
}

func (m *_BACnetApplicationTagObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
