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

// SearchResponse is the corresponding interface of SearchResponse
type SearchResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
	// GetDibDeviceInfo returns DibDeviceInfo (property field)
	GetDibDeviceInfo() DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies (property field)
	GetDibSuppSvcFamilies() DIBSuppSvcFamilies
	// IsSearchResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSearchResponse()
	// CreateBuilder creates a SearchResponseBuilder
	CreateSearchResponseBuilder() SearchResponseBuilder
}

// _SearchResponse is the data-structure of this message
type _SearchResponse struct {
	KnxNetIpMessageContract
	HpaiControlEndpoint HPAIControlEndpoint
	DibDeviceInfo       DIBDeviceInfo
	DibSuppSvcFamilies  DIBSuppSvcFamilies
}

var _ SearchResponse = (*_SearchResponse)(nil)
var _ KnxNetIpMessageRequirements = (*_SearchResponse)(nil)

// NewSearchResponse factory function for _SearchResponse
func NewSearchResponse(hpaiControlEndpoint HPAIControlEndpoint, dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) *_SearchResponse {
	if hpaiControlEndpoint == nil {
		panic("hpaiControlEndpoint of type HPAIControlEndpoint for SearchResponse must not be nil")
	}
	if dibDeviceInfo == nil {
		panic("dibDeviceInfo of type DIBDeviceInfo for SearchResponse must not be nil")
	}
	if dibSuppSvcFamilies == nil {
		panic("dibSuppSvcFamilies of type DIBSuppSvcFamilies for SearchResponse must not be nil")
	}
	_result := &_SearchResponse{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		HpaiControlEndpoint:     hpaiControlEndpoint,
		DibDeviceInfo:           dibDeviceInfo,
		DibSuppSvcFamilies:      dibSuppSvcFamilies,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SearchResponseBuilder is a builder for SearchResponse
type SearchResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hpaiControlEndpoint HPAIControlEndpoint, dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) SearchResponseBuilder
	// WithHpaiControlEndpoint adds HpaiControlEndpoint (property field)
	WithHpaiControlEndpoint(HPAIControlEndpoint) SearchResponseBuilder
	// WithHpaiControlEndpointBuilder adds HpaiControlEndpoint (property field) which is build by the builder
	WithHpaiControlEndpointBuilder(func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) SearchResponseBuilder
	// WithDibDeviceInfo adds DibDeviceInfo (property field)
	WithDibDeviceInfo(DIBDeviceInfo) SearchResponseBuilder
	// WithDibDeviceInfoBuilder adds DibDeviceInfo (property field) which is build by the builder
	WithDibDeviceInfoBuilder(func(DIBDeviceInfoBuilder) DIBDeviceInfoBuilder) SearchResponseBuilder
	// WithDibSuppSvcFamilies adds DibSuppSvcFamilies (property field)
	WithDibSuppSvcFamilies(DIBSuppSvcFamilies) SearchResponseBuilder
	// WithDibSuppSvcFamiliesBuilder adds DibSuppSvcFamilies (property field) which is build by the builder
	WithDibSuppSvcFamiliesBuilder(func(DIBSuppSvcFamiliesBuilder) DIBSuppSvcFamiliesBuilder) SearchResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the SearchResponse or returns an error if something is wrong
	Build() (SearchResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SearchResponse
}

// NewSearchResponseBuilder() creates a SearchResponseBuilder
func NewSearchResponseBuilder() SearchResponseBuilder {
	return &_SearchResponseBuilder{_SearchResponse: new(_SearchResponse)}
}

type _SearchResponseBuilder struct {
	*_SearchResponse

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (SearchResponseBuilder) = (*_SearchResponseBuilder)(nil)

func (b *_SearchResponseBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._SearchResponse
}

func (b *_SearchResponseBuilder) WithMandatoryFields(hpaiControlEndpoint HPAIControlEndpoint, dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) SearchResponseBuilder {
	return b.WithHpaiControlEndpoint(hpaiControlEndpoint).WithDibDeviceInfo(dibDeviceInfo).WithDibSuppSvcFamilies(dibSuppSvcFamilies)
}

func (b *_SearchResponseBuilder) WithHpaiControlEndpoint(hpaiControlEndpoint HPAIControlEndpoint) SearchResponseBuilder {
	b.HpaiControlEndpoint = hpaiControlEndpoint
	return b
}

func (b *_SearchResponseBuilder) WithHpaiControlEndpointBuilder(builderSupplier func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) SearchResponseBuilder {
	builder := builderSupplier(b.HpaiControlEndpoint.CreateHPAIControlEndpointBuilder())
	var err error
	b.HpaiControlEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIControlEndpointBuilder failed"))
	}
	return b
}

func (b *_SearchResponseBuilder) WithDibDeviceInfo(dibDeviceInfo DIBDeviceInfo) SearchResponseBuilder {
	b.DibDeviceInfo = dibDeviceInfo
	return b
}

func (b *_SearchResponseBuilder) WithDibDeviceInfoBuilder(builderSupplier func(DIBDeviceInfoBuilder) DIBDeviceInfoBuilder) SearchResponseBuilder {
	builder := builderSupplier(b.DibDeviceInfo.CreateDIBDeviceInfoBuilder())
	var err error
	b.DibDeviceInfo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DIBDeviceInfoBuilder failed"))
	}
	return b
}

func (b *_SearchResponseBuilder) WithDibSuppSvcFamilies(dibSuppSvcFamilies DIBSuppSvcFamilies) SearchResponseBuilder {
	b.DibSuppSvcFamilies = dibSuppSvcFamilies
	return b
}

func (b *_SearchResponseBuilder) WithDibSuppSvcFamiliesBuilder(builderSupplier func(DIBSuppSvcFamiliesBuilder) DIBSuppSvcFamiliesBuilder) SearchResponseBuilder {
	builder := builderSupplier(b.DibSuppSvcFamilies.CreateDIBSuppSvcFamiliesBuilder())
	var err error
	b.DibSuppSvcFamilies, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DIBSuppSvcFamiliesBuilder failed"))
	}
	return b
}

func (b *_SearchResponseBuilder) Build() (SearchResponse, error) {
	if b.HpaiControlEndpoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hpaiControlEndpoint' not set"))
	}
	if b.DibDeviceInfo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dibDeviceInfo' not set"))
	}
	if b.DibSuppSvcFamilies == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dibSuppSvcFamilies' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SearchResponse.deepCopy(), nil
}

func (b *_SearchResponseBuilder) MustBuild() SearchResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SearchResponseBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_SearchResponseBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_SearchResponseBuilder) DeepCopy() any {
	_copy := b.CreateSearchResponseBuilder().(*_SearchResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSearchResponseBuilder creates a SearchResponseBuilder
func (b *_SearchResponse) CreateSearchResponseBuilder() SearchResponseBuilder {
	if b == nil {
		return NewSearchResponseBuilder()
	}
	return &_SearchResponseBuilder{_SearchResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SearchResponse) GetMsgType() uint16 {
	return 0x0202
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SearchResponse) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SearchResponse) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

func (m *_SearchResponse) GetDibDeviceInfo() DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *_SearchResponse) GetDibSuppSvcFamilies() DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSearchResponse(structType any) SearchResponse {
	if casted, ok := structType.(SearchResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SearchResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SearchResponse) GetTypeName() string {
	return "SearchResponse"
}

func (m *_SearchResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits(ctx)

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SearchResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SearchResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__searchResponse SearchResponse, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SearchResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SearchResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hpaiControlEndpoint, err := ReadSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", ReadComplex[HPAIControlEndpoint](HPAIControlEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiControlEndpoint' field"))
	}
	m.HpaiControlEndpoint = hpaiControlEndpoint

	dibDeviceInfo, err := ReadSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", ReadComplex[DIBDeviceInfo](DIBDeviceInfoParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibDeviceInfo' field"))
	}
	m.DibDeviceInfo = dibDeviceInfo

	dibSuppSvcFamilies, err := ReadSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", ReadComplex[DIBSuppSvcFamilies](DIBSuppSvcFamiliesParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dibSuppSvcFamilies' field"))
	}
	m.DibSuppSvcFamilies = dibSuppSvcFamilies

	if closeErr := readBuffer.CloseContext("SearchResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SearchResponse")
	}

	return m, nil
}

func (m *_SearchResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SearchResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SearchResponse")
		}

		if err := WriteSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", m.GetHpaiControlEndpoint(), WriteComplex[HPAIControlEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiControlEndpoint' field")
		}

		if err := WriteSimpleField[DIBDeviceInfo](ctx, "dibDeviceInfo", m.GetDibDeviceInfo(), WriteComplex[DIBDeviceInfo](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibDeviceInfo' field")
		}

		if err := WriteSimpleField[DIBSuppSvcFamilies](ctx, "dibSuppSvcFamilies", m.GetDibSuppSvcFamilies(), WriteComplex[DIBSuppSvcFamilies](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("SearchResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SearchResponse")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SearchResponse) IsSearchResponse() {}

func (m *_SearchResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SearchResponse) deepCopy() *_SearchResponse {
	if m == nil {
		return nil
	}
	_SearchResponseCopy := &_SearchResponse{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopy[HPAIControlEndpoint](m.HpaiControlEndpoint),
		utils.DeepCopy[DIBDeviceInfo](m.DibDeviceInfo),
		utils.DeepCopy[DIBSuppSvcFamilies](m.DibSuppSvcFamilies),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _SearchResponseCopy
}

func (m *_SearchResponse) String() string {
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
