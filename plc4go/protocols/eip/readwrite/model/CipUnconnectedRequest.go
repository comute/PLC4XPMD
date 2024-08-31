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
const CipUnconnectedRequest_ROUTE uint16 = 0x0001

// CipUnconnectedRequest is the corresponding interface of CipUnconnectedRequest
type CipUnconnectedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetUnconnectedService returns UnconnectedService (property field)
	GetUnconnectedService() CipService
	// GetBackPlane returns BackPlane (property field)
	GetBackPlane() int8
	// GetSlot returns Slot (property field)
	GetSlot() int8
}

// CipUnconnectedRequestExactly can be used when we want exactly this type and not a type which fulfills CipUnconnectedRequest.
// This is useful for switch cases.
type CipUnconnectedRequestExactly interface {
	CipUnconnectedRequest
	isCipUnconnectedRequest() bool
}

// _CipUnconnectedRequest is the data-structure of this message
type _CipUnconnectedRequest struct {
	*_CipService
	ClassSegment       PathSegment
	InstanceSegment    PathSegment
	UnconnectedService CipService
	BackPlane          int8
	Slot               int8
	// Reserved Fields
	reservedField0 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipUnconnectedRequest) GetService() uint8 {
	return 0x52
}

func (m *_CipUnconnectedRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipUnconnectedRequest) GetConnected() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipUnconnectedRequest) InitializeParent(parent CipService) {}

func (m *_CipUnconnectedRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipUnconnectedRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipUnconnectedRequest) GetUnconnectedService() CipService {
	return m.UnconnectedService
}

func (m *_CipUnconnectedRequest) GetBackPlane() int8 {
	return m.BackPlane
}

func (m *_CipUnconnectedRequest) GetSlot() int8 {
	return m.Slot
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipUnconnectedRequest) GetRoute() uint16 {
	return CipUnconnectedRequest_ROUTE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipUnconnectedRequest factory function for _CipUnconnectedRequest
func NewCipUnconnectedRequest(classSegment PathSegment, instanceSegment PathSegment, unconnectedService CipService, backPlane int8, slot int8, serviceLen uint16) *_CipUnconnectedRequest {
	_result := &_CipUnconnectedRequest{
		ClassSegment:       classSegment,
		InstanceSegment:    instanceSegment,
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		_CipService:        NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipUnconnectedRequest(structType any) CipUnconnectedRequest {
	if casted, ok := structType.(CipUnconnectedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipUnconnectedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipUnconnectedRequest) GetTypeName() string {
	return "CipUnconnectedRequest"
}

func (m *_CipUnconnectedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (messageSize)
	lengthInBits += 16

	// Simple field (unconnectedService)
	lengthInBits += m.UnconnectedService.GetLengthInBits(ctx)

	// Const Field (route)
	lengthInBits += 16

	// Simple field (backPlane)
	lengthInBits += 8

	// Simple field (slot)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipUnconnectedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipUnconnectedRequestParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipUnconnectedRequest, error) {
	return CipUnconnectedRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipUnconnectedRequestParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CipUnconnectedRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CipUnconnectedRequest, error) {
		return CipUnconnectedRequestParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func CipUnconnectedRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipUnconnectedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipUnconnectedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipUnconnectedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x9D05))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	messageSize, err := ReadImplicitField[uint16](ctx, "messageSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageSize' field"))
	}
	_ = messageSize

	unconnectedService, err := ReadSimpleField[CipService](ctx, "unconnectedService", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(false)), (uint16)(messageSize)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unconnectedService' field"))
	}

	route, err := ReadConstField[uint16](ctx, "route", ReadUnsignedShort(readBuffer, uint8(16)), CipUnconnectedRequest_ROUTE)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'route' field"))
	}
	_ = route

	backPlane, err := ReadSimpleField(ctx, "backPlane", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backPlane' field"))
	}

	slot, err := ReadSimpleField(ctx, "slot", ReadSignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'slot' field"))
	}

	if closeErr := readBuffer.CloseContext("CipUnconnectedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipUnconnectedRequest")
	}

	// Create a partially initialized instance
	_child := &_CipUnconnectedRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		ClassSegment:       classSegment,
		InstanceSegment:    instanceSegment,
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		reservedField0:     reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipUnconnectedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipUnconnectedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipUnconnectedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipUnconnectedRequest")
		}

		// Implicit Field (requestPathSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		_requestPathSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("requestPathSize", 8, uint8((requestPathSize)))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Simple Field (classSegment)
		if pushErr := writeBuffer.PushContext("classSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for classSegment")
		}
		_classSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetClassSegment())
		if popErr := writeBuffer.PopContext("classSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for classSegment")
		}
		if _classSegmentErr != nil {
			return errors.Wrap(_classSegmentErr, "Error serializing 'classSegment' field")
		}

		// Simple Field (instanceSegment)
		if pushErr := writeBuffer.PushContext("instanceSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instanceSegment")
		}
		_instanceSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetInstanceSegment())
		if popErr := writeBuffer.PopContext("instanceSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instanceSegment")
		}
		if _instanceSegmentErr != nil {
			return errors.Wrap(_instanceSegmentErr, "Error serializing 'instanceSegment' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint16 = uint16(0x9D05)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint16(0x9D05),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint16("reserved", 16, uint16(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := uint16(uint16(uint16(uint16(m.GetLengthInBytes(ctx)))-uint16(uint16(10))) - uint16(uint16(4)))
		_messageSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("messageSize", 16, uint16((messageSize)))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (unconnectedService)
		if pushErr := writeBuffer.PushContext("unconnectedService"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unconnectedService")
		}
		_unconnectedServiceErr := writeBuffer.WriteSerializable(ctx, m.GetUnconnectedService())
		if popErr := writeBuffer.PopContext("unconnectedService"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unconnectedService")
		}
		if _unconnectedServiceErr != nil {
			return errors.Wrap(_unconnectedServiceErr, "Error serializing 'unconnectedService' field")
		}

		// Const Field (route)
		_routeErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint16("route", 16, uint16(0x0001))
		if _routeErr != nil {
			return errors.Wrap(_routeErr, "Error serializing 'route' field")
		}

		// Simple Field (backPlane)
		backPlane := int8(m.GetBackPlane())
		_backPlaneErr := /*TODO: migrate me*/ writeBuffer.WriteInt8("backPlane", 8, int8((backPlane)))
		if _backPlaneErr != nil {
			return errors.Wrap(_backPlaneErr, "Error serializing 'backPlane' field")
		}

		// Simple Field (slot)
		slot := int8(m.GetSlot())
		_slotErr := /*TODO: migrate me*/ writeBuffer.WriteInt8("slot", 8, int8((slot)))
		if _slotErr != nil {
			return errors.Wrap(_slotErr, "Error serializing 'slot' field")
		}

		if popErr := writeBuffer.PopContext("CipUnconnectedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipUnconnectedRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipUnconnectedRequest) isCipUnconnectedRequest() bool {
	return true
}

func (m *_CipUnconnectedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
