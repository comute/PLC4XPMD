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

// AdsDiscoveryBlockVersion is the corresponding interface of AdsDiscoveryBlockVersion
type AdsDiscoveryBlockVersion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetVersionData returns VersionData (property field)
	GetVersionData() []byte
}

// AdsDiscoveryBlockVersionExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockVersion.
// This is useful for switch cases.
type AdsDiscoveryBlockVersionExactly interface {
	AdsDiscoveryBlockVersion
	isAdsDiscoveryBlockVersion() bool
}

// _AdsDiscoveryBlockVersion is the data-structure of this message
type _AdsDiscoveryBlockVersion struct {
	*_AdsDiscoveryBlock
	VersionData []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockVersion) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_VERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockVersion) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockVersion) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockVersion) GetVersionData() []byte {
	return m.VersionData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockVersion factory function for _AdsDiscoveryBlockVersion
func NewAdsDiscoveryBlockVersion(versionData []byte) *_AdsDiscoveryBlockVersion {
	_result := &_AdsDiscoveryBlockVersion{
		VersionData:        versionData,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockVersion(structType any) AdsDiscoveryBlockVersion {
	if casted, ok := structType.(AdsDiscoveryBlockVersion); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockVersion); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockVersion) GetTypeName() string {
	return "AdsDiscoveryBlockVersion"
}

func (m *_AdsDiscoveryBlockVersion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (versionDataLen)
	lengthInBits += 16

	// Array field
	if len(m.VersionData) > 0 {
		lengthInBits += 8 * uint16(len(m.VersionData))
	}

	return lengthInBits
}

func (m *_AdsDiscoveryBlockVersion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockVersionParse(ctx context.Context, theBytes []byte) (AdsDiscoveryBlockVersion, error) {
	return AdsDiscoveryBlockVersionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockVersionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockVersion, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockVersion, error) {
		return AdsDiscoveryBlockVersionParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDiscoveryBlockVersionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockVersion, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockVersion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	versionDataLen, err := ReadImplicitField[uint16](ctx, "versionDataLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'versionDataLen' field"))
	}
	_ = versionDataLen

	versionData, err := readBuffer.ReadByteArray("versionData", int(versionDataLen))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'versionData' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockVersion")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockVersion{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		VersionData:        versionData,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockVersion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockVersion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockVersion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockVersion")
		}

		// Implicit Field (versionDataLen) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		versionDataLen := uint16(uint16(len(m.GetVersionData())))
		_versionDataLenErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("versionDataLen", 16, uint16((versionDataLen)))
		if _versionDataLenErr != nil {
			return errors.Wrap(_versionDataLenErr, "Error serializing 'versionDataLen' field")
		}

		if err := WriteByteArrayField(ctx, "versionData", m.GetVersionData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'versionData' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockVersion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockVersion")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockVersion) isAdsDiscoveryBlockVersion() bool {
	return true
}

func (m *_AdsDiscoveryBlockVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
