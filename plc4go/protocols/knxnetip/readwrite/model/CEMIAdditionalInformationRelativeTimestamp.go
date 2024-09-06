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

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = uint8(2)

// CEMIAdditionalInformationRelativeTimestamp is the corresponding interface of CEMIAdditionalInformationRelativeTimestamp
type CEMIAdditionalInformationRelativeTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMIAdditionalInformation
	// GetRelativeTimestamp returns RelativeTimestamp (property field)
	GetRelativeTimestamp() RelativeTimestamp
	// IsCEMIAdditionalInformationRelativeTimestamp is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMIAdditionalInformationRelativeTimestamp()
}

// _CEMIAdditionalInformationRelativeTimestamp is the data-structure of this message
type _CEMIAdditionalInformationRelativeTimestamp struct {
	CEMIAdditionalInformationContract
	RelativeTimestamp RelativeTimestamp
}

var _ CEMIAdditionalInformationRelativeTimestamp = (*_CEMIAdditionalInformationRelativeTimestamp)(nil)
var _ CEMIAdditionalInformationRequirements = (*_CEMIAdditionalInformationRelativeTimestamp)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetAdditionalInformationType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetParent() CEMIAdditionalInformationContract {
	return m.CEMIAdditionalInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetRelativeTimestamp() RelativeTimestamp {
	return m.RelativeTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLen() uint8 {
	return CEMIAdditionalInformationRelativeTimestamp_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationRelativeTimestamp factory function for _CEMIAdditionalInformationRelativeTimestamp
func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp RelativeTimestamp) *_CEMIAdditionalInformationRelativeTimestamp {
	if relativeTimestamp == nil {
		panic("relativeTimestamp of type RelativeTimestamp for CEMIAdditionalInformationRelativeTimestamp must not be nil")
	}
	_result := &_CEMIAdditionalInformationRelativeTimestamp{
		CEMIAdditionalInformationContract: NewCEMIAdditionalInformation(),
		RelativeTimestamp:                 relativeTimestamp,
	}
	_result.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformationRelativeTimestamp(structType any) CEMIAdditionalInformationRelativeTimestamp {
	if casted, ok := structType.(CEMIAdditionalInformationRelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationRelativeTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
	return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).getLengthInBits(ctx))

	// Const Field (len)
	lengthInBits += 8

	// Simple field (relativeTimestamp)
	lengthInBits += m.RelativeTimestamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMIAdditionalInformation) (__cEMIAdditionalInformationRelativeTimestamp CEMIAdditionalInformationRelativeTimestamp, err error) {
	m.CEMIAdditionalInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationRelativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformationRelativeTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	len, err := ReadConstField[uint8](ctx, "len", ReadUnsignedByte(readBuffer, uint8(8)), CEMIAdditionalInformationRelativeTimestamp_LEN)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'len' field"))
	}
	_ = len

	relativeTimestamp, err := ReadSimpleField[RelativeTimestamp](ctx, "relativeTimestamp", ReadComplex[RelativeTimestamp](RelativeTimestampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relativeTimestamp' field"))
	}
	m.RelativeTimestamp = relativeTimestamp

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationRelativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformationRelativeTimestamp")
	}

	return m, nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationRelativeTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformationRelativeTimestamp")
		}

		if err := WriteConstField(ctx, "len", CEMIAdditionalInformationRelativeTimestamp_LEN, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'len' field")
		}

		if err := WriteSimpleField[RelativeTimestamp](ctx, "relativeTimestamp", m.GetRelativeTimestamp(), WriteComplex[RelativeTimestamp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relativeTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationRelativeTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformationRelativeTimestamp")
		}
		return nil
	}
	return m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) IsCEMIAdditionalInformationRelativeTimestamp() {
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
