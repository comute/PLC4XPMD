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

// NLMNetworkNumberIs is the corresponding interface of NLMNetworkNumberIs
type NLMNetworkNumberIs interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() uint16
	// GetNetworkNumberConfigured returns NetworkNumberConfigured (property field)
	GetNetworkNumberConfigured() bool
}

// NLMNetworkNumberIsExactly can be used when we want exactly this type and not a type which fulfills NLMNetworkNumberIs.
// This is useful for switch cases.
type NLMNetworkNumberIsExactly interface {
	NLMNetworkNumberIs
	isNLMNetworkNumberIs() bool
}

// _NLMNetworkNumberIs is the data-structure of this message
type _NLMNetworkNumberIs struct {
	*_NLM
	NetworkNumber           uint16
	NetworkNumberConfigured bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMNetworkNumberIs) GetMessageType() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMNetworkNumberIs) InitializeParent(parent NLM) {}

func (m *_NLMNetworkNumberIs) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMNetworkNumberIs) GetNetworkNumber() uint16 {
	return m.NetworkNumber
}

func (m *_NLMNetworkNumberIs) GetNetworkNumberConfigured() bool {
	return m.NetworkNumberConfigured
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMNetworkNumberIs factory function for _NLMNetworkNumberIs
func NewNLMNetworkNumberIs(networkNumber uint16, networkNumberConfigured bool, apduLength uint16) *_NLMNetworkNumberIs {
	_result := &_NLMNetworkNumberIs{
		NetworkNumber:           networkNumber,
		NetworkNumberConfigured: networkNumberConfigured,
		_NLM:                    NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMNetworkNumberIs(structType any) NLMNetworkNumberIs {
	if casted, ok := structType.(NLMNetworkNumberIs); ok {
		return casted
	}
	if casted, ok := structType.(*NLMNetworkNumberIs); ok {
		return *casted
	}
	return nil
}

func (m *_NLMNetworkNumberIs) GetTypeName() string {
	return "NLMNetworkNumberIs"
}

func (m *_NLMNetworkNumberIs) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (networkNumber)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (networkNumberConfigured)
	lengthInBits += 1

	return lengthInBits
}

func (m *_NLMNetworkNumberIs) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMNetworkNumberIsParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMNetworkNumberIs, error) {
	return NLMNetworkNumberIsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMNetworkNumberIsParseWithBufferProducer(apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMNetworkNumberIs, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NLMNetworkNumberIs, error) {
		return NLMNetworkNumberIsParseWithBuffer(ctx, readBuffer, apduLength)
	}
}

func NLMNetworkNumberIsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMNetworkNumberIs, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMNetworkNumberIs"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMNetworkNumberIs")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumber, err := ReadSimpleField(ctx, "networkNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumber' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	networkNumberConfigured, err := ReadSimpleField(ctx, "networkNumberConfigured", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumberConfigured' field"))
	}

	if closeErr := readBuffer.CloseContext("NLMNetworkNumberIs"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMNetworkNumberIs")
	}

	// Create a partially initialized instance
	_child := &_NLMNetworkNumberIs{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		NetworkNumber:           networkNumber,
		NetworkNumberConfigured: networkNumberConfigured,
		reservedField0:          reservedField0,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMNetworkNumberIs) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMNetworkNumberIs) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMNetworkNumberIs"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMNetworkNumberIs")
		}

		// Simple Field (networkNumber)
		networkNumber := uint16(m.GetNetworkNumber())
		_networkNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("networkNumber", 16, uint16((networkNumber)))
		if _networkNumberErr != nil {
			return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (networkNumberConfigured)
		networkNumberConfigured := bool(m.GetNetworkNumberConfigured())
		_networkNumberConfiguredErr := /*TODO: migrate me*/ writeBuffer.WriteBit("networkNumberConfigured", (networkNumberConfigured))
		if _networkNumberConfiguredErr != nil {
			return errors.Wrap(_networkNumberConfiguredErr, "Error serializing 'networkNumberConfigured' field")
		}

		if popErr := writeBuffer.PopContext("NLMNetworkNumberIs"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMNetworkNumberIs")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMNetworkNumberIs) isNLMNetworkNumberIs() bool {
	return true
}

func (m *_NLMNetworkNumberIs) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
