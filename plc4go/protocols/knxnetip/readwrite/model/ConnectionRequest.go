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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ConnectionRequest is the corresponding interface of ConnectionRequest
type ConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetHpaiDiscoveryEndpoint returns HpaiDiscoveryEndpoint (property field)
	GetHpaiDiscoveryEndpoint() HPAIDiscoveryEndpoint
	// GetHpaiDataEndpoint returns HpaiDataEndpoint (property field)
	GetHpaiDataEndpoint() HPAIDataEndpoint
	// GetConnectionRequestInformation returns ConnectionRequestInformation (property field)
	GetConnectionRequestInformation() ConnectionRequestInformation
	// IsConnectionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConnectionRequest()
	// CreateBuilder creates a ConnectionRequestBuilder
	CreateConnectionRequestBuilder() ConnectionRequestBuilder
}

// _ConnectionRequest is the data-structure of this message
type _ConnectionRequest struct {
	KnxNetIpMessageContract
	HpaiDiscoveryEndpoint        HPAIDiscoveryEndpoint
	HpaiDataEndpoint             HPAIDataEndpoint
	ConnectionRequestInformation ConnectionRequestInformation
}

var _ ConnectionRequest = (*_ConnectionRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_ConnectionRequest)(nil)

// NewConnectionRequest factory function for _ConnectionRequest
func NewConnectionRequest(hpaiDiscoveryEndpoint HPAIDiscoveryEndpoint, hpaiDataEndpoint HPAIDataEndpoint, connectionRequestInformation ConnectionRequestInformation) *_ConnectionRequest {
	if hpaiDiscoveryEndpoint == nil {
		panic("hpaiDiscoveryEndpoint of type HPAIDiscoveryEndpoint for ConnectionRequest must not be nil")
	}
	if hpaiDataEndpoint == nil {
		panic("hpaiDataEndpoint of type HPAIDataEndpoint for ConnectionRequest must not be nil")
	}
	if connectionRequestInformation == nil {
		panic("connectionRequestInformation of type ConnectionRequestInformation for ConnectionRequest must not be nil")
	}
	_result := &_ConnectionRequest{
		KnxNetIpMessageContract:      NewKnxNetIpMessage(),
		HpaiDiscoveryEndpoint:        hpaiDiscoveryEndpoint,
		HpaiDataEndpoint:             hpaiDataEndpoint,
		ConnectionRequestInformation: connectionRequestInformation,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConnectionRequestBuilder is a builder for ConnectionRequest
type ConnectionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hpaiDiscoveryEndpoint HPAIDiscoveryEndpoint, hpaiDataEndpoint HPAIDataEndpoint, connectionRequestInformation ConnectionRequestInformation) ConnectionRequestBuilder
	// WithHpaiDiscoveryEndpoint adds HpaiDiscoveryEndpoint (property field)
	WithHpaiDiscoveryEndpoint(HPAIDiscoveryEndpoint) ConnectionRequestBuilder
	// WithHpaiDiscoveryEndpointBuilder adds HpaiDiscoveryEndpoint (property field) which is build by the builder
	WithHpaiDiscoveryEndpointBuilder(func(HPAIDiscoveryEndpointBuilder) HPAIDiscoveryEndpointBuilder) ConnectionRequestBuilder
	// WithHpaiDataEndpoint adds HpaiDataEndpoint (property field)
	WithHpaiDataEndpoint(HPAIDataEndpoint) ConnectionRequestBuilder
	// WithHpaiDataEndpointBuilder adds HpaiDataEndpoint (property field) which is build by the builder
	WithHpaiDataEndpointBuilder(func(HPAIDataEndpointBuilder) HPAIDataEndpointBuilder) ConnectionRequestBuilder
	// WithConnectionRequestInformation adds ConnectionRequestInformation (property field)
	WithConnectionRequestInformation(ConnectionRequestInformation) ConnectionRequestBuilder
	// WithConnectionRequestInformationBuilder adds ConnectionRequestInformation (property field) which is build by the builder
	WithConnectionRequestInformationBuilder(func(ConnectionRequestInformationBuilder) ConnectionRequestInformationBuilder) ConnectionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the ConnectionRequest or returns an error if something is wrong
	Build() (ConnectionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ConnectionRequest
}

// NewConnectionRequestBuilder() creates a ConnectionRequestBuilder
func NewConnectionRequestBuilder() ConnectionRequestBuilder {
	return &_ConnectionRequestBuilder{_ConnectionRequest: new(_ConnectionRequest)}
}

type _ConnectionRequestBuilder struct {
	*_ConnectionRequest

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (ConnectionRequestBuilder) = (*_ConnectionRequestBuilder)(nil)

func (b *_ConnectionRequestBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._ConnectionRequest
}

func (b *_ConnectionRequestBuilder) WithMandatoryFields(hpaiDiscoveryEndpoint HPAIDiscoveryEndpoint, hpaiDataEndpoint HPAIDataEndpoint, connectionRequestInformation ConnectionRequestInformation) ConnectionRequestBuilder {
	return b.WithHpaiDiscoveryEndpoint(hpaiDiscoveryEndpoint).WithHpaiDataEndpoint(hpaiDataEndpoint).WithConnectionRequestInformation(connectionRequestInformation)
}

func (b *_ConnectionRequestBuilder) WithHpaiDiscoveryEndpoint(hpaiDiscoveryEndpoint HPAIDiscoveryEndpoint) ConnectionRequestBuilder {
	b.HpaiDiscoveryEndpoint = hpaiDiscoveryEndpoint
	return b
}

func (b *_ConnectionRequestBuilder) WithHpaiDiscoveryEndpointBuilder(builderSupplier func(HPAIDiscoveryEndpointBuilder) HPAIDiscoveryEndpointBuilder) ConnectionRequestBuilder {
	builder := builderSupplier(b.HpaiDiscoveryEndpoint.CreateHPAIDiscoveryEndpointBuilder())
	var err error
	b.HpaiDiscoveryEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIDiscoveryEndpointBuilder failed"))
	}
	return b
}

func (b *_ConnectionRequestBuilder) WithHpaiDataEndpoint(hpaiDataEndpoint HPAIDataEndpoint) ConnectionRequestBuilder {
	b.HpaiDataEndpoint = hpaiDataEndpoint
	return b
}

func (b *_ConnectionRequestBuilder) WithHpaiDataEndpointBuilder(builderSupplier func(HPAIDataEndpointBuilder) HPAIDataEndpointBuilder) ConnectionRequestBuilder {
	builder := builderSupplier(b.HpaiDataEndpoint.CreateHPAIDataEndpointBuilder())
	var err error
	b.HpaiDataEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIDataEndpointBuilder failed"))
	}
	return b
}

func (b *_ConnectionRequestBuilder) WithConnectionRequestInformation(connectionRequestInformation ConnectionRequestInformation) ConnectionRequestBuilder {
	b.ConnectionRequestInformation = connectionRequestInformation
	return b
}

func (b *_ConnectionRequestBuilder) WithConnectionRequestInformationBuilder(builderSupplier func(ConnectionRequestInformationBuilder) ConnectionRequestInformationBuilder) ConnectionRequestBuilder {
	builder := builderSupplier(b.ConnectionRequestInformation.CreateConnectionRequestInformationBuilder())
	var err error
	b.ConnectionRequestInformation, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ConnectionRequestInformationBuilder failed"))
	}
	return b
}

func (b *_ConnectionRequestBuilder) Build() (ConnectionRequest, error) {
	if b.HpaiDiscoveryEndpoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hpaiDiscoveryEndpoint' not set"))
	}
	if b.HpaiDataEndpoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hpaiDataEndpoint' not set"))
	}
	if b.ConnectionRequestInformation == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'connectionRequestInformation' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ConnectionRequest.deepCopy(), nil
}

func (b *_ConnectionRequestBuilder) MustBuild() ConnectionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConnectionRequestBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_ConnectionRequestBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_ConnectionRequestBuilder) DeepCopy() any {
	_copy := b.CreateConnectionRequestBuilder().(*_ConnectionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConnectionRequestBuilder creates a ConnectionRequestBuilder
func (b *_ConnectionRequest) CreateConnectionRequestBuilder() ConnectionRequestBuilder {
	if b == nil {
		return NewConnectionRequestBuilder()
	}
	return &_ConnectionRequestBuilder{_ConnectionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequest) GetMsgType() uint16 {
	return 0x0205
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectionRequest) GetHpaiDiscoveryEndpoint() HPAIDiscoveryEndpoint {
	return m.HpaiDiscoveryEndpoint
}

func (m *_ConnectionRequest) GetHpaiDataEndpoint() HPAIDataEndpoint {
	return m.HpaiDataEndpoint
}

func (m *_ConnectionRequest) GetConnectionRequestInformation() ConnectionRequestInformation {
	return m.ConnectionRequestInformation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConnectionRequest(structType any) ConnectionRequest {
	if casted, ok := structType.(ConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequest) GetTypeName() string {
	return "ConnectionRequest"
}

func (m *_ConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (hpaiDiscoveryEndpoint)
	lengthInBits += m.HpaiDiscoveryEndpoint.GetLengthInBits(ctx)

	// Simple field (hpaiDataEndpoint)
	lengthInBits += m.HpaiDataEndpoint.GetLengthInBits(ctx)

	// Simple field (connectionRequestInformation)
	lengthInBits += m.ConnectionRequestInformation.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ConnectionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__connectionRequest ConnectionRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hpaiDiscoveryEndpoint, err := ReadSimpleField[HPAIDiscoveryEndpoint](ctx, "hpaiDiscoveryEndpoint", ReadComplex[HPAIDiscoveryEndpoint](HPAIDiscoveryEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiDiscoveryEndpoint' field"))
	}
	m.HpaiDiscoveryEndpoint = hpaiDiscoveryEndpoint

	hpaiDataEndpoint, err := ReadSimpleField[HPAIDataEndpoint](ctx, "hpaiDataEndpoint", ReadComplex[HPAIDataEndpoint](HPAIDataEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiDataEndpoint' field"))
	}
	m.HpaiDataEndpoint = hpaiDataEndpoint

	connectionRequestInformation, err := ReadSimpleField[ConnectionRequestInformation](ctx, "connectionRequestInformation", ReadComplex[ConnectionRequestInformation](ConnectionRequestInformationParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionRequestInformation' field"))
	}
	m.ConnectionRequestInformation = connectionRequestInformation

	if closeErr := readBuffer.CloseContext("ConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequest")
	}

	return m, nil
}

func (m *_ConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequest")
		}

		if err := WriteSimpleField[HPAIDiscoveryEndpoint](ctx, "hpaiDiscoveryEndpoint", m.GetHpaiDiscoveryEndpoint(), WriteComplex[HPAIDiscoveryEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiDiscoveryEndpoint' field")
		}

		if err := WriteSimpleField[HPAIDataEndpoint](ctx, "hpaiDataEndpoint", m.GetHpaiDataEndpoint(), WriteComplex[HPAIDataEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiDataEndpoint' field")
		}

		if err := WriteSimpleField[ConnectionRequestInformation](ctx, "connectionRequestInformation", m.GetConnectionRequestInformation(), WriteComplex[ConnectionRequestInformation](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionRequestInformation' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequest) IsConnectionRequest() {}

func (m *_ConnectionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ConnectionRequest) deepCopy() *_ConnectionRequest {
	if m == nil {
		return nil
	}
	_ConnectionRequestCopy := &_ConnectionRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopy[HPAIDiscoveryEndpoint](m.HpaiDiscoveryEndpoint),
		utils.DeepCopy[HPAIDataEndpoint](m.HpaiDataEndpoint),
		utils.DeepCopy[ConnectionRequestInformation](m.ConnectionRequestInformation),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _ConnectionRequestCopy
}

func (m *_ConnectionRequest) String() string {
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
