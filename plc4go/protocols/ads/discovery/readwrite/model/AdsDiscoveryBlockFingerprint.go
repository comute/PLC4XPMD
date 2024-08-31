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

// AdsDiscoveryBlockFingerprint is the corresponding interface of AdsDiscoveryBlockFingerprint
type AdsDiscoveryBlockFingerprint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsDiscoveryBlockFingerprintExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockFingerprint.
// This is useful for switch cases.
type AdsDiscoveryBlockFingerprintExactly interface {
	AdsDiscoveryBlockFingerprint
	isAdsDiscoveryBlockFingerprint() bool
}

// _AdsDiscoveryBlockFingerprint is the data-structure of this message
type _AdsDiscoveryBlockFingerprint struct {
	*_AdsDiscoveryBlock
	Data []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockFingerprint) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_FINGERPRINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockFingerprint) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockFingerprint) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockFingerprint) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockFingerprint factory function for _AdsDiscoveryBlockFingerprint
func NewAdsDiscoveryBlockFingerprint(data []byte) *_AdsDiscoveryBlockFingerprint {
	_result := &_AdsDiscoveryBlockFingerprint{
		Data:               data,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockFingerprint(structType any) AdsDiscoveryBlockFingerprint {
	if casted, ok := structType.(AdsDiscoveryBlockFingerprint); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockFingerprint); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockFingerprint) GetTypeName() string {
	return "AdsDiscoveryBlockFingerprint"
}

func (m *_AdsDiscoveryBlockFingerprint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (dataLen)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsDiscoveryBlockFingerprint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockFingerprintParse(ctx context.Context, theBytes []byte) (AdsDiscoveryBlockFingerprint, error) {
	return AdsDiscoveryBlockFingerprintParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockFingerprintParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockFingerprint, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockFingerprint, error) {
		return AdsDiscoveryBlockFingerprintParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDiscoveryBlockFingerprintParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockFingerprint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockFingerprint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockFingerprint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataLen, err := ReadImplicitField[uint16](ctx, "dataLen", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataLen' field"))
	}
	_ = dataLen

	data, err := readBuffer.ReadByteArray("data", int(dataLen))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockFingerprint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockFingerprint")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockFingerprint{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		Data:               data,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockFingerprint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockFingerprint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockFingerprint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockFingerprint")
		}

		// Implicit Field (dataLen) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		dataLen := uint16(uint16(len(m.GetData())))
		_dataLenErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("dataLen", 16, uint16((dataLen)))
		if _dataLenErr != nil {
			return errors.Wrap(_dataLenErr, "Error serializing 'dataLen' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockFingerprint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockFingerprint")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockFingerprint) isAdsDiscoveryBlockFingerprint() bool {
	return true
}

func (m *_AdsDiscoveryBlockFingerprint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
