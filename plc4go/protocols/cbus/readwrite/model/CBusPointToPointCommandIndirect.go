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

// CBusPointToPointCommandIndirect is the corresponding interface of CBusPointToPointCommandIndirect
type CBusPointToPointCommandIndirect interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusPointToPointCommand
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// IsCBusPointToPointCommandIndirect is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusPointToPointCommandIndirect()
	// CreateBuilder creates a CBusPointToPointCommandIndirectBuilder
	CreateCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder
}

// _CBusPointToPointCommandIndirect is the data-structure of this message
type _CBusPointToPointCommandIndirect struct {
	CBusPointToPointCommandContract
	BridgeAddress BridgeAddress
	NetworkRoute  NetworkRoute
	UnitAddress   UnitAddress
}

var _ CBusPointToPointCommandIndirect = (*_CBusPointToPointCommandIndirect)(nil)
var _ CBusPointToPointCommandRequirements = (*_CBusPointToPointCommandIndirect)(nil)

// NewCBusPointToPointCommandIndirect factory function for _CBusPointToPointCommandIndirect
func NewCBusPointToPointCommandIndirect(bridgeAddressCountPeek uint16, calData CALData, bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress, cBusOptions CBusOptions) *_CBusPointToPointCommandIndirect {
	if bridgeAddress == nil {
		panic("bridgeAddress of type BridgeAddress for CBusPointToPointCommandIndirect must not be nil")
	}
	if networkRoute == nil {
		panic("networkRoute of type NetworkRoute for CBusPointToPointCommandIndirect must not be nil")
	}
	if unitAddress == nil {
		panic("unitAddress of type UnitAddress for CBusPointToPointCommandIndirect must not be nil")
	}
	_result := &_CBusPointToPointCommandIndirect{
		CBusPointToPointCommandContract: NewCBusPointToPointCommand(bridgeAddressCountPeek, calData, cBusOptions),
		BridgeAddress:                   bridgeAddress,
		NetworkRoute:                    networkRoute,
		UnitAddress:                     unitAddress,
	}
	_result.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusPointToPointCommandIndirectBuilder is a builder for CBusPointToPointCommandIndirect
type CBusPointToPointCommandIndirectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder
	// WithBridgeAddress adds BridgeAddress (property field)
	WithBridgeAddress(BridgeAddress) CBusPointToPointCommandIndirectBuilder
	// WithBridgeAddressBuilder adds BridgeAddress (property field) which is build by the builder
	WithBridgeAddressBuilder(func(BridgeAddressBuilder) BridgeAddressBuilder) CBusPointToPointCommandIndirectBuilder
	// WithNetworkRoute adds NetworkRoute (property field)
	WithNetworkRoute(NetworkRoute) CBusPointToPointCommandIndirectBuilder
	// WithNetworkRouteBuilder adds NetworkRoute (property field) which is build by the builder
	WithNetworkRouteBuilder(func(NetworkRouteBuilder) NetworkRouteBuilder) CBusPointToPointCommandIndirectBuilder
	// WithUnitAddress adds UnitAddress (property field)
	WithUnitAddress(UnitAddress) CBusPointToPointCommandIndirectBuilder
	// WithUnitAddressBuilder adds UnitAddress (property field) which is build by the builder
	WithUnitAddressBuilder(func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandIndirectBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CBusPointToPointCommandBuilder
	// Build builds the CBusPointToPointCommandIndirect or returns an error if something is wrong
	Build() (CBusPointToPointCommandIndirect, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusPointToPointCommandIndirect
}

// NewCBusPointToPointCommandIndirectBuilder() creates a CBusPointToPointCommandIndirectBuilder
func NewCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder {
	return &_CBusPointToPointCommandIndirectBuilder{_CBusPointToPointCommandIndirect: new(_CBusPointToPointCommandIndirect)}
}

type _CBusPointToPointCommandIndirectBuilder struct {
	*_CBusPointToPointCommandIndirect

	parentBuilder *_CBusPointToPointCommandBuilder

	err *utils.MultiError
}

var _ (CBusPointToPointCommandIndirectBuilder) = (*_CBusPointToPointCommandIndirectBuilder)(nil)

func (b *_CBusPointToPointCommandIndirectBuilder) setParent(contract CBusPointToPointCommandContract) {
	b.CBusPointToPointCommandContract = contract
	contract.(*_CBusPointToPointCommand)._SubType = b._CBusPointToPointCommandIndirect
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithMandatoryFields(bridgeAddress BridgeAddress, networkRoute NetworkRoute, unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder {
	return b.WithBridgeAddress(bridgeAddress).WithNetworkRoute(networkRoute).WithUnitAddress(unitAddress)
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithBridgeAddress(bridgeAddress BridgeAddress) CBusPointToPointCommandIndirectBuilder {
	b.BridgeAddress = bridgeAddress
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithBridgeAddressBuilder(builderSupplier func(BridgeAddressBuilder) BridgeAddressBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(b.BridgeAddress.CreateBridgeAddressBuilder())
	var err error
	b.BridgeAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BridgeAddressBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithNetworkRoute(networkRoute NetworkRoute) CBusPointToPointCommandIndirectBuilder {
	b.NetworkRoute = networkRoute
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithNetworkRouteBuilder(builderSupplier func(NetworkRouteBuilder) NetworkRouteBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(b.NetworkRoute.CreateNetworkRouteBuilder())
	var err error
	b.NetworkRoute, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NetworkRouteBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithUnitAddress(unitAddress UnitAddress) CBusPointToPointCommandIndirectBuilder {
	b.UnitAddress = unitAddress
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) WithUnitAddressBuilder(builderSupplier func(UnitAddressBuilder) UnitAddressBuilder) CBusPointToPointCommandIndirectBuilder {
	builder := builderSupplier(b.UnitAddress.CreateUnitAddressBuilder())
	var err error
	b.UnitAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "UnitAddressBuilder failed"))
	}
	return b
}

func (b *_CBusPointToPointCommandIndirectBuilder) Build() (CBusPointToPointCommandIndirect, error) {
	if b.BridgeAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bridgeAddress' not set"))
	}
	if b.NetworkRoute == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkRoute' not set"))
	}
	if b.UnitAddress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unitAddress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusPointToPointCommandIndirect.deepCopy(), nil
}

func (b *_CBusPointToPointCommandIndirectBuilder) MustBuild() CBusPointToPointCommandIndirect {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusPointToPointCommandIndirectBuilder) Done() CBusPointToPointCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCBusPointToPointCommandBuilder().(*_CBusPointToPointCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_CBusPointToPointCommandIndirectBuilder) buildForCBusPointToPointCommand() (CBusPointToPointCommand, error) {
	return b.Build()
}

func (b *_CBusPointToPointCommandIndirectBuilder) DeepCopy() any {
	_copy := b.CreateCBusPointToPointCommandIndirectBuilder().(*_CBusPointToPointCommandIndirectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusPointToPointCommandIndirectBuilder creates a CBusPointToPointCommandIndirectBuilder
func (b *_CBusPointToPointCommandIndirect) CreateCBusPointToPointCommandIndirectBuilder() CBusPointToPointCommandIndirectBuilder {
	if b == nil {
		return NewCBusPointToPointCommandIndirectBuilder()
	}
	return &_CBusPointToPointCommandIndirectBuilder{_CBusPointToPointCommandIndirect: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToPointCommandIndirect) GetParent() CBusPointToPointCommandContract {
	return m.CBusPointToPointCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointCommandIndirect) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointCommandIndirect) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointCommandIndirect) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusPointToPointCommandIndirect(structType any) CBusPointToPointCommandIndirect {
	if casted, ok := structType.(CBusPointToPointCommandIndirect); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointCommandIndirect); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointCommandIndirect) GetTypeName() string {
	return "CBusPointToPointCommandIndirect"
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).getLengthInBits(ctx))

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	// Simple field (unitAddress)
	lengthInBits += m.UnitAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointCommandIndirect) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusPointToPointCommandIndirect) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusPointToPointCommand, cBusOptions CBusOptions) (__cBusPointToPointCommandIndirect CBusPointToPointCommandIndirect, err error) {
	m.CBusPointToPointCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointCommandIndirect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointCommandIndirect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bridgeAddress, err := ReadSimpleField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	m.BridgeAddress = bridgeAddress

	networkRoute, err := ReadSimpleField[NetworkRoute](ctx, "networkRoute", ReadComplex[NetworkRoute](NetworkRouteParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkRoute' field"))
	}
	m.NetworkRoute = networkRoute

	unitAddress, err := ReadSimpleField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	m.UnitAddress = unitAddress

	if closeErr := readBuffer.CloseContext("CBusPointToPointCommandIndirect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointCommandIndirect")
	}

	return m, nil
}

func (m *_CBusPointToPointCommandIndirect) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointCommandIndirect) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointCommandIndirect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointCommandIndirect")
		}

		if err := WriteSimpleField[BridgeAddress](ctx, "bridgeAddress", m.GetBridgeAddress(), WriteComplex[BridgeAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
		}

		if err := WriteSimpleField[NetworkRoute](ctx, "networkRoute", m.GetNetworkRoute(), WriteComplex[NetworkRoute](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkRoute' field")
		}

		if err := WriteSimpleField[UnitAddress](ctx, "unitAddress", m.GetUnitAddress(), WriteComplex[UnitAddress](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToPointCommandIndirect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointCommandIndirect")
		}
		return nil
	}
	return m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToPointCommandIndirect) IsCBusPointToPointCommandIndirect() {}

func (m *_CBusPointToPointCommandIndirect) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusPointToPointCommandIndirect) deepCopy() *_CBusPointToPointCommandIndirect {
	if m == nil {
		return nil
	}
	_CBusPointToPointCommandIndirectCopy := &_CBusPointToPointCommandIndirect{
		m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand).deepCopy(),
		utils.DeepCopy[BridgeAddress](m.BridgeAddress),
		utils.DeepCopy[NetworkRoute](m.NetworkRoute),
		utils.DeepCopy[UnitAddress](m.UnitAddress),
	}
	m.CBusPointToPointCommandContract.(*_CBusPointToPointCommand)._SubType = m
	return _CBusPointToPointCommandIndirectCopy
}

func (m *_CBusPointToPointCommandIndirect) String() string {
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
