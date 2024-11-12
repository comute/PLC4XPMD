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

// AdsDiscoveryBlockOsData is the corresponding interface of AdsDiscoveryBlockOsData
type AdsDiscoveryBlockOsData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetOsData returns OsData (property field)
	GetOsData() []byte
	// IsAdsDiscoveryBlockOsData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockOsData()
	// CreateBuilder creates a AdsDiscoveryBlockOsDataBuilder
	CreateAdsDiscoveryBlockOsDataBuilder() AdsDiscoveryBlockOsDataBuilder
}

// _AdsDiscoveryBlockOsData is the data-structure of this message
type _AdsDiscoveryBlockOsData struct {
	AdsDiscoveryBlockContract
	OsData []byte
}

var _ AdsDiscoveryBlockOsData = (*_AdsDiscoveryBlockOsData)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockOsData)(nil)

// NewAdsDiscoveryBlockOsData factory function for _AdsDiscoveryBlockOsData
func NewAdsDiscoveryBlockOsData(osData []byte) *_AdsDiscoveryBlockOsData {
	_result := &_AdsDiscoveryBlockOsData{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		OsData:                    osData,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDiscoveryBlockOsDataBuilder is a builder for AdsDiscoveryBlockOsData
type AdsDiscoveryBlockOsDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(osData []byte) AdsDiscoveryBlockOsDataBuilder
	// WithOsData adds OsData (property field)
	WithOsData(...byte) AdsDiscoveryBlockOsDataBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AdsDiscoveryBlockBuilder
	// Build builds the AdsDiscoveryBlockOsData or returns an error if something is wrong
	Build() (AdsDiscoveryBlockOsData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDiscoveryBlockOsData
}

// NewAdsDiscoveryBlockOsDataBuilder() creates a AdsDiscoveryBlockOsDataBuilder
func NewAdsDiscoveryBlockOsDataBuilder() AdsDiscoveryBlockOsDataBuilder {
	return &_AdsDiscoveryBlockOsDataBuilder{_AdsDiscoveryBlockOsData: new(_AdsDiscoveryBlockOsData)}
}

type _AdsDiscoveryBlockOsDataBuilder struct {
	*_AdsDiscoveryBlockOsData

	parentBuilder *_AdsDiscoveryBlockBuilder

	err *utils.MultiError
}

var _ (AdsDiscoveryBlockOsDataBuilder) = (*_AdsDiscoveryBlockOsDataBuilder)(nil)

func (b *_AdsDiscoveryBlockOsDataBuilder) setParent(contract AdsDiscoveryBlockContract) {
	b.AdsDiscoveryBlockContract = contract
	contract.(*_AdsDiscoveryBlock)._SubType = b._AdsDiscoveryBlockOsData
}

func (b *_AdsDiscoveryBlockOsDataBuilder) WithMandatoryFields(osData []byte) AdsDiscoveryBlockOsDataBuilder {
	return b.WithOsData(osData...)
}

func (b *_AdsDiscoveryBlockOsDataBuilder) WithOsData(osData ...byte) AdsDiscoveryBlockOsDataBuilder {
	b.OsData = osData
	return b
}

func (b *_AdsDiscoveryBlockOsDataBuilder) Build() (AdsDiscoveryBlockOsData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDiscoveryBlockOsData.deepCopy(), nil
}

func (b *_AdsDiscoveryBlockOsDataBuilder) MustBuild() AdsDiscoveryBlockOsData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsDiscoveryBlockOsDataBuilder) Done() AdsDiscoveryBlockBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAdsDiscoveryBlockBuilder().(*_AdsDiscoveryBlockBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsDiscoveryBlockOsDataBuilder) buildForAdsDiscoveryBlock() (AdsDiscoveryBlock, error) {
	return b.Build()
}

func (b *_AdsDiscoveryBlockOsDataBuilder) DeepCopy() any {
	_copy := b.CreateAdsDiscoveryBlockOsDataBuilder().(*_AdsDiscoveryBlockOsDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDiscoveryBlockOsDataBuilder creates a AdsDiscoveryBlockOsDataBuilder
func (b *_AdsDiscoveryBlockOsData) CreateAdsDiscoveryBlockOsDataBuilder() AdsDiscoveryBlockOsDataBuilder {
	if b == nil {
		return NewAdsDiscoveryBlockOsDataBuilder()
	}
	return &_AdsDiscoveryBlockOsDataBuilder{_AdsDiscoveryBlockOsData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockOsData) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_OS_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockOsData) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockOsData) GetOsData() []byte {
	return m.OsData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockOsData(structType any) AdsDiscoveryBlockOsData {
	if casted, ok := structType.(AdsDiscoveryBlockOsData); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockOsData); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockOsData) GetTypeName() string {
	return "AdsDiscoveryBlockOsData"
}

func (m *_AdsDiscoveryBlockOsData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Implicit Field (osDataLen)
	lengthInBits += 16

	// Array field
	if len(m.OsData) > 0 {
		lengthInBits += 8 * uint16(len(m.OsData))
	}

	return lengthInBits
}

func (m *_AdsDiscoveryBlockOsData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockOsData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockOsData AdsDiscoveryBlockOsData, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockOsData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockOsData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	osDataLen, err := ReadImplicitField[uint16](ctx, "osDataLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'osDataLen' field"))
	}
	_ = osDataLen

	osData, err := readBuffer.ReadByteArray("osData", int(osDataLen))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'osData' field"))
	}
	m.OsData = osData

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockOsData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockOsData")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockOsData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockOsData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockOsData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockOsData")
		}
		osDataLen := uint16(uint16(len(m.GetOsData())))
		if err := WriteImplicitField(ctx, "osDataLen", osDataLen, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'osDataLen' field")
		}

		if err := WriteByteArrayField(ctx, "osData", m.GetOsData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'osData' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockOsData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockOsData")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockOsData) IsAdsDiscoveryBlockOsData() {}

func (m *_AdsDiscoveryBlockOsData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockOsData) deepCopy() *_AdsDiscoveryBlockOsData {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockOsDataCopy := &_AdsDiscoveryBlockOsData{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.OsData),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockOsDataCopy
}

func (m *_AdsDiscoveryBlockOsData) String() string {
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
