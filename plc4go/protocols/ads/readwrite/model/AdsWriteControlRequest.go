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

// AdsWriteControlRequest is the corresponding interface of AdsWriteControlRequest
type AdsWriteControlRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// GetAdsState returns AdsState (property field)
	GetAdsState() uint16
	// GetDeviceState returns DeviceState (property field)
	GetDeviceState() uint16
	// GetData returns Data (property field)
	GetData() []byte
	// IsAdsWriteControlRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsWriteControlRequest()
	// CreateBuilder creates a AdsWriteControlRequestBuilder
	CreateAdsWriteControlRequestBuilder() AdsWriteControlRequestBuilder
}

// _AdsWriteControlRequest is the data-structure of this message
type _AdsWriteControlRequest struct {
	AmsPacketContract
	AdsState    uint16
	DeviceState uint16
	Data        []byte
}

var _ AdsWriteControlRequest = (*_AdsWriteControlRequest)(nil)
var _ AmsPacketRequirements = (*_AdsWriteControlRequest)(nil)

// NewAdsWriteControlRequest factory function for _AdsWriteControlRequest
func NewAdsWriteControlRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32, adsState uint16, deviceState uint16, data []byte) *_AdsWriteControlRequest {
	_result := &_AdsWriteControlRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
		AdsState:          adsState,
		DeviceState:       deviceState,
		Data:              data,
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsWriteControlRequestBuilder is a builder for AdsWriteControlRequest
type AdsWriteControlRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(adsState uint16, deviceState uint16, data []byte) AdsWriteControlRequestBuilder
	// WithAdsState adds AdsState (property field)
	WithAdsState(uint16) AdsWriteControlRequestBuilder
	// WithDeviceState adds DeviceState (property field)
	WithDeviceState(uint16) AdsWriteControlRequestBuilder
	// WithData adds Data (property field)
	WithData(...byte) AdsWriteControlRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsWriteControlRequest or returns an error if something is wrong
	Build() (AdsWriteControlRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsWriteControlRequest
}

// NewAdsWriteControlRequestBuilder() creates a AdsWriteControlRequestBuilder
func NewAdsWriteControlRequestBuilder() AdsWriteControlRequestBuilder {
	return &_AdsWriteControlRequestBuilder{_AdsWriteControlRequest: new(_AdsWriteControlRequest)}
}

type _AdsWriteControlRequestBuilder struct {
	*_AdsWriteControlRequest

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsWriteControlRequestBuilder) = (*_AdsWriteControlRequestBuilder)(nil)

func (b *_AdsWriteControlRequestBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
	contract.(*_AmsPacket)._SubType = b._AdsWriteControlRequest
}

func (b *_AdsWriteControlRequestBuilder) WithMandatoryFields(adsState uint16, deviceState uint16, data []byte) AdsWriteControlRequestBuilder {
	return b.WithAdsState(adsState).WithDeviceState(deviceState).WithData(data...)
}

func (b *_AdsWriteControlRequestBuilder) WithAdsState(adsState uint16) AdsWriteControlRequestBuilder {
	b.AdsState = adsState
	return b
}

func (b *_AdsWriteControlRequestBuilder) WithDeviceState(deviceState uint16) AdsWriteControlRequestBuilder {
	b.DeviceState = deviceState
	return b
}

func (b *_AdsWriteControlRequestBuilder) WithData(data ...byte) AdsWriteControlRequestBuilder {
	b.Data = data
	return b
}

func (b *_AdsWriteControlRequestBuilder) Build() (AdsWriteControlRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsWriteControlRequest.deepCopy(), nil
}

func (b *_AdsWriteControlRequestBuilder) MustBuild() AdsWriteControlRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsWriteControlRequestBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsWriteControlRequestBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsWriteControlRequestBuilder) DeepCopy() any {
	_copy := b.CreateAdsWriteControlRequestBuilder().(*_AdsWriteControlRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsWriteControlRequestBuilder creates a AdsWriteControlRequestBuilder
func (b *_AdsWriteControlRequest) CreateAdsWriteControlRequestBuilder() AdsWriteControlRequestBuilder {
	if b == nil {
		return NewAdsWriteControlRequestBuilder()
	}
	return &_AdsWriteControlRequestBuilder{_AdsWriteControlRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsWriteControlRequest) GetCommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *_AdsWriteControlRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsWriteControlRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsWriteControlRequest) GetAdsState() uint16 {
	return m.AdsState
}

func (m *_AdsWriteControlRequest) GetDeviceState() uint16 {
	return m.DeviceState
}

func (m *_AdsWriteControlRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsWriteControlRequest(structType any) AdsWriteControlRequest {
	if casted, ok := structType.(AdsWriteControlRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsWriteControlRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsWriteControlRequest) GetTypeName() string {
	return "AdsWriteControlRequest"
}

func (m *_AdsWriteControlRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsWriteControlRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsWriteControlRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsWriteControlRequest AdsWriteControlRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsWriteControlRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsWriteControlRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adsState, err := ReadSimpleField(ctx, "adsState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsState' field"))
	}
	m.AdsState = adsState

	deviceState, err := ReadSimpleField(ctx, "deviceState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceState' field"))
	}
	m.DeviceState = deviceState

	length, err := ReadImplicitField[uint32](ctx, "length", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	data, err := readBuffer.ReadByteArray("data", int(length))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("AdsWriteControlRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsWriteControlRequest")
	}

	return m, nil
}

func (m *_AdsWriteControlRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsWriteControlRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteControlRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsWriteControlRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "adsState", m.GetAdsState(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'adsState' field")
		}

		if err := WriteSimpleField[uint16](ctx, "deviceState", m.GetDeviceState(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceState' field")
		}
		length := uint32(uint32(len(m.GetData())))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteControlRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsWriteControlRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsWriteControlRequest) IsAdsWriteControlRequest() {}

func (m *_AdsWriteControlRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsWriteControlRequest) deepCopy() *_AdsWriteControlRequest {
	if m == nil {
		return nil
	}
	_AdsWriteControlRequestCopy := &_AdsWriteControlRequest{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
		m.AdsState,
		m.DeviceState,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsWriteControlRequestCopy
}

func (m *_AdsWriteControlRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
