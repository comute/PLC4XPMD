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

// BACnetConfirmedServiceRequestVTData is the corresponding interface of BACnetConfirmedServiceRequestVTData
type BACnetConfirmedServiceRequestVTData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetVtSessionIdentifier returns VtSessionIdentifier (property field)
	GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// GetVtNewData returns VtNewData (property field)
	GetVtNewData() BACnetApplicationTagOctetString
	// GetVtDataFlag returns VtDataFlag (property field)
	GetVtDataFlag() BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestVTDataExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestVTData.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestVTDataExactly interface {
	BACnetConfirmedServiceRequestVTData
	isBACnetConfirmedServiceRequestVTData() bool
}

// _BACnetConfirmedServiceRequestVTData is the data-structure of this message
type _BACnetConfirmedServiceRequestVTData struct {
	*_BACnetConfirmedServiceRequest
	VtSessionIdentifier BACnetApplicationTagUnsignedInteger
	VtNewData           BACnetApplicationTagOctetString
	VtDataFlag          BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestVTData) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTData) GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.VtSessionIdentifier
}

func (m *_BACnetConfirmedServiceRequestVTData) GetVtNewData() BACnetApplicationTagOctetString {
	return m.VtNewData
}

func (m *_BACnetConfirmedServiceRequestVTData) GetVtDataFlag() BACnetApplicationTagUnsignedInteger {
	return m.VtDataFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestVTData factory function for _BACnetConfirmedServiceRequestVTData
func NewBACnetConfirmedServiceRequestVTData(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestVTData {
	_result := &_BACnetConfirmedServiceRequestVTData{
		VtSessionIdentifier:            vtSessionIdentifier,
		VtNewData:                      vtNewData,
		VtDataFlag:                     vtDataFlag,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTData(structType any) BACnetConfirmedServiceRequestVTData {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTData) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTData"
}

func (m *_BACnetConfirmedServiceRequestVTData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (vtSessionIdentifier)
	lengthInBits += m.VtSessionIdentifier.GetLengthInBits(ctx)

	// Simple field (vtNewData)
	lengthInBits += m.VtNewData.GetLengthInBits(ctx)

	// Simple field (vtDataFlag)
	lengthInBits += m.VtDataFlag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestVTDataParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTData, error) {
	return BACnetConfirmedServiceRequestVTDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestVTDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vtSessionIdentifier)
	if pullErr := readBuffer.PullContext("vtSessionIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtSessionIdentifier")
	}
	_vtSessionIdentifier, _vtSessionIdentifierErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _vtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_vtSessionIdentifierErr, "Error parsing 'vtSessionIdentifier' field of BACnetConfirmedServiceRequestVTData")
	}
	vtSessionIdentifier := _vtSessionIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("vtSessionIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtSessionIdentifier")
	}

	// Simple Field (vtNewData)
	if pullErr := readBuffer.PullContext("vtNewData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtNewData")
	}
	_vtNewData, _vtNewDataErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _vtNewDataErr != nil {
		return nil, errors.Wrap(_vtNewDataErr, "Error parsing 'vtNewData' field of BACnetConfirmedServiceRequestVTData")
	}
	vtNewData := _vtNewData.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("vtNewData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtNewData")
	}

	// Simple Field (vtDataFlag)
	if pullErr := readBuffer.PullContext("vtDataFlag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtDataFlag")
	}
	_vtDataFlag, _vtDataFlagErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _vtDataFlagErr != nil {
		return nil, errors.Wrap(_vtDataFlagErr, "Error parsing 'vtDataFlag' field of BACnetConfirmedServiceRequestVTData")
	}
	vtDataFlag := _vtDataFlag.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("vtDataFlag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtDataFlag")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTData")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestVTData{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		VtSessionIdentifier: vtSessionIdentifier,
		VtNewData:           vtNewData,
		VtDataFlag:          vtDataFlag,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestVTData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTData")
		}

		// Simple Field (vtSessionIdentifier)
		if pushErr := writeBuffer.PushContext("vtSessionIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtSessionIdentifier")
		}
		_vtSessionIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetVtSessionIdentifier())
		if popErr := writeBuffer.PopContext("vtSessionIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtSessionIdentifier")
		}
		if _vtSessionIdentifierErr != nil {
			return errors.Wrap(_vtSessionIdentifierErr, "Error serializing 'vtSessionIdentifier' field")
		}

		// Simple Field (vtNewData)
		if pushErr := writeBuffer.PushContext("vtNewData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtNewData")
		}
		_vtNewDataErr := writeBuffer.WriteSerializable(ctx, m.GetVtNewData())
		if popErr := writeBuffer.PopContext("vtNewData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtNewData")
		}
		if _vtNewDataErr != nil {
			return errors.Wrap(_vtNewDataErr, "Error serializing 'vtNewData' field")
		}

		// Simple Field (vtDataFlag)
		if pushErr := writeBuffer.PushContext("vtDataFlag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtDataFlag")
		}
		_vtDataFlagErr := writeBuffer.WriteSerializable(ctx, m.GetVtDataFlag())
		if popErr := writeBuffer.PopContext("vtDataFlag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtDataFlag")
		}
		if _vtDataFlagErr != nil {
			return errors.Wrap(_vtDataFlagErr, "Error serializing 'vtDataFlag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestVTData) isBACnetConfirmedServiceRequestVTData() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestVTData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
