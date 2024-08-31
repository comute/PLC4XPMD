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

// Services is the corresponding interface of Services
type Services interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOffsets returns Offsets (property field)
	GetOffsets() []uint16
	// GetServices returns Services (property field)
	GetServices() []CipService
}

// ServicesExactly can be used when we want exactly this type and not a type which fulfills Services.
// This is useful for switch cases.
type ServicesExactly interface {
	Services
	isServices() bool
}

// _Services is the data-structure of this message
type _Services struct {
	Offsets  []uint16
	Services []CipService

	// Arguments.
	ServicesLen uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Services) GetOffsets() []uint16 {
	return m.Offsets
}

func (m *_Services) GetServices() []CipService {
	return m.Services
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServices factory function for _Services
func NewServices(offsets []uint16, services []CipService, servicesLen uint16) *_Services {
	return &_Services{Offsets: offsets, Services: services, ServicesLen: servicesLen}
}

// Deprecated: use the interface for direct cast
func CastServices(structType any) Services {
	if casted, ok := structType.(Services); ok {
		return casted
	}
	if casted, ok := structType.(*Services); ok {
		return *casted
	}
	return nil
}

func (m *_Services) GetTypeName() string {
	return "Services"
}

func (m *_Services) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (serviceNb)
	lengthInBits += 16

	// Array field
	if len(m.Offsets) > 0 {
		lengthInBits += 16 * uint16(len(m.Offsets))
	}

	// Array field
	if len(m.Services) > 0 {
		for _curItem, element := range m.Services {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Services), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_Services) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServicesParse(ctx context.Context, theBytes []byte, servicesLen uint16) (Services, error) {
	return ServicesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), servicesLen)
}

func ServicesParseWithBufferProducer(servicesLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (Services, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Services, error) {
		return ServicesParseWithBuffer(ctx, readBuffer, servicesLen)
	}
}

func ServicesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, servicesLen uint16) (Services, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Services"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Services")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceNb, err := ReadImplicitField[uint16](ctx, "serviceNb", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNb' field"))
	}
	_ = serviceNb

	offsets, err := ReadCountArrayField[uint16](ctx, "offsets", ReadUnsignedShort(readBuffer, uint8(16)), uint64(serviceNb))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offsets' field"))
	}

	services, err := ReadCountArrayField[CipService](ctx, "services", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(false)), (uint16)(uint16(servicesLen)/uint16(serviceNb))), readBuffer), uint64(serviceNb))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'services' field"))
	}

	if closeErr := readBuffer.CloseContext("Services"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Services")
	}

	// Create the instance
	return &_Services{
		ServicesLen: servicesLen,
		Offsets:     offsets,
		Services:    services,
	}, nil
}

func (m *_Services) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Services) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Services"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Services")
	}

	// Implicit Field (serviceNb) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	serviceNb := uint16(uint16(len(m.GetOffsets())))
	_serviceNbErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("serviceNb", 16, uint16((serviceNb)))
	if _serviceNbErr != nil {
		return errors.Wrap(_serviceNbErr, "Error serializing 'serviceNb' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "offsets", m.GetOffsets(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'offsets' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "services", m.GetServices(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'services' field")
	}

	if popErr := writeBuffer.PopContext("Services"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Services")
	}
	return nil
}

////
// Arguments Getter

func (m *_Services) GetServicesLen() uint16 {
	return m.ServicesLen
}

//
////

func (m *_Services) isServices() bool {
	return true
}

func (m *_Services) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
