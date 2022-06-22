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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestConfirmedPrivateTransfer is the corresponding interface of BACnetConfirmedServiceRequestConfirmedPrivateTransfer
type BACnetConfirmedServiceRequestConfirmedPrivateTransfer interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetServiceParameters returns ServiceParameters (property field)
	GetServiceParameters() BACnetConstructedData
}

// BACnetConfirmedServiceRequestConfirmedPrivateTransferExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedPrivateTransfer.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedPrivateTransferExactly interface {
	BACnetConfirmedServiceRequestConfirmedPrivateTransfer
	isBACnetConfirmedServiceRequestConfirmedPrivateTransfer() bool
}

// _BACnetConfirmedServiceRequestConfirmedPrivateTransfer is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedPrivateTransfer struct {
	*_BACnetConfirmedServiceRequest
	VendorId          BACnetVendorIdTagged
	ServiceNumber     BACnetContextTagUnsignedInteger
	ServiceParameters BACnetConstructedData

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceParameters() BACnetConstructedData {
	return m.ServiceParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedPrivateTransfer factory function for _BACnetConfirmedServiceRequestConfirmedPrivateTransfer
func NewBACnetConfirmedServiceRequestConfirmedPrivateTransfer(vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, serviceParameters BACnetConstructedData, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer {
	_result := &_BACnetConfirmedServiceRequestConfirmedPrivateTransfer{
		VendorId:                       vendorId,
		ServiceNumber:                  serviceNumber,
		ServiceParameters:              serviceParameters,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedPrivateTransfer(structType interface{}) BACnetConfirmedServiceRequestConfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedPrivateTransfer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedPrivateTransfer"
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits()

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits()

	// Optional Field (serviceParameters)
	if m.ServiceParameters != nil {
		lengthInBits += m.ServiceParameters.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedPrivateTransferParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestConfirmedPrivateTransfer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceNumber")
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}
	serviceNumber := _serviceNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceNumber")
	}

	// Optional Field (serviceParameters) (Can be skipped, if a given expression evaluates to false)
	var serviceParameters BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("serviceParameters"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for serviceParameters")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(2), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'serviceParameters' field")
		default:
			serviceParameters = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("serviceParameters"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for serviceParameters")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedPrivateTransfer{
		VendorId:                       vendorId,
		ServiceNumber:                  serviceNumber,
		ServiceParameters:              serviceParameters,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceNumber")
		}
		_serviceNumberErr := writeBuffer.WriteSerializable(m.GetServiceNumber())
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceNumber")
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (serviceParameters) (Can be skipped, if the value is null)
		var serviceParameters BACnetConstructedData = nil
		if m.GetServiceParameters() != nil {
			if pushErr := writeBuffer.PushContext("serviceParameters"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceParameters")
			}
			serviceParameters = m.GetServiceParameters()
			_serviceParametersErr := writeBuffer.WriteSerializable(serviceParameters)
			if popErr := writeBuffer.PopContext("serviceParameters"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceParameters")
			}
			if _serviceParametersErr != nil {
				return errors.Wrap(_serviceParametersErr, "Error serializing 'serviceParameters' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) isBACnetConfirmedServiceRequestConfirmedPrivateTransfer() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
