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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Services is the corresponding interface of Services
type Services interface {
	utils.LengthAware
	utils.Serializable
	// GetServiceNb returns ServiceNb (property field)
	GetServiceNb() uint16
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
	ServiceNb uint16
	Offsets   []uint16
	Services  []CipService

	// Arguments.
	ServicesLen uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Services) GetServiceNb() uint16 {
	return m.ServiceNb
}

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
func NewServices(serviceNb uint16, offsets []uint16, services []CipService, servicesLen uint16) *_Services {
	return &_Services{ServiceNb: serviceNb, Offsets: offsets, Services: services, ServicesLen: servicesLen}
}

// Deprecated: use the interface for direct cast
func CastServices(structType interface{}) Services {
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

func (m *_Services) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_Services) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (serviceNb)
	lengthInBits += 16

	// Array field
	if len(m.Offsets) > 0 {
		lengthInBits += 16 * uint16(len(m.Offsets))
	}

	// Array field
	if len(m.Services) > 0 {
		for i, element := range m.Services {
			last := i == len(m.Services)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_Services) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ServicesParse(readBuffer utils.ReadBuffer, servicesLen uint16) (Services, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Services"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Services")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (serviceNb)
	_serviceNb, _serviceNbErr := readBuffer.ReadUint16("serviceNb", 16)
	if _serviceNbErr != nil {
		return nil, errors.Wrap(_serviceNbErr, "Error parsing 'serviceNb' field of Services")
	}
	serviceNb := _serviceNb

	// Array field (offsets)
	if pullErr := readBuffer.PullContext("offsets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for offsets")
	}
	// Count array
	offsets := make([]uint16, serviceNb)
	// This happens when the size is set conditional to 0
	if len(offsets) == 0 {
		offsets = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(serviceNb); curItem++ {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'offsets' field of Services")
			}
			offsets[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("offsets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for offsets")
	}

	// Array field (services)
	if pullErr := readBuffer.PullContext("services", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for services")
	}
	// Count array
	services := make([]CipService, serviceNb)
	// This happens when the size is set conditional to 0
	if len(services) == 0 {
		services = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(serviceNb); curItem++ {
			_item, _err := CipServiceParse(readBuffer, uint16(servicesLen)/uint16(serviceNb))
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'services' field of Services")
			}
			services[curItem] = _item.(CipService)
		}
	}
	if closeErr := readBuffer.CloseContext("services", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for services")
	}

	if closeErr := readBuffer.CloseContext("Services"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Services")
	}

	// Create the instance
	return NewServices(serviceNb, offsets, services, servicesLen), nil
}

func (m *_Services) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Services"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Services")
	}

	// Simple Field (serviceNb)
	serviceNb := uint16(m.GetServiceNb())
	_serviceNbErr := writeBuffer.WriteUint16("serviceNb", 16, (serviceNb))
	if _serviceNbErr != nil {
		return errors.Wrap(_serviceNbErr, "Error serializing 'serviceNb' field")
	}

	// Array Field (offsets)
	if pushErr := writeBuffer.PushContext("offsets", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for offsets")
	}
	for _, _element := range m.GetOffsets() {
		_elementErr := writeBuffer.WriteUint16("", 16, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'offsets' field")
		}
	}
	if popErr := writeBuffer.PopContext("offsets", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for offsets")
	}

	// Array Field (services)
	if pushErr := writeBuffer.PushContext("services", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for services")
	}
	for _, _element := range m.GetServices() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'services' field")
		}
	}
	if popErr := writeBuffer.PopContext("services", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for services")
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
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
