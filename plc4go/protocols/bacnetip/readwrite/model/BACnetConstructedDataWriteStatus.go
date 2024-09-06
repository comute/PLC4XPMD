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

// BACnetConstructedDataWriteStatus is the corresponding interface of BACnetConstructedDataWriteStatus
type BACnetConstructedDataWriteStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetWriteStatus returns WriteStatus (property field)
	GetWriteStatus() BACnetWriteStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetWriteStatusTagged
	// IsBACnetConstructedDataWriteStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataWriteStatus()
}

// _BACnetConstructedDataWriteStatus is the data-structure of this message
type _BACnetConstructedDataWriteStatus struct {
	BACnetConstructedDataContract
	WriteStatus BACnetWriteStatusTagged
}

var _ BACnetConstructedDataWriteStatus = (*_BACnetConstructedDataWriteStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataWriteStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWriteStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWriteStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WRITE_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWriteStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWriteStatus) GetWriteStatus() BACnetWriteStatusTagged {
	return m.WriteStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWriteStatus) GetActualValue() BACnetWriteStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetWriteStatusTagged(m.GetWriteStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataWriteStatus factory function for _BACnetConstructedDataWriteStatus
func NewBACnetConstructedDataWriteStatus(writeStatus BACnetWriteStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWriteStatus {
	if writeStatus == nil {
		panic("writeStatus of type BACnetWriteStatusTagged for BACnetConstructedDataWriteStatus must not be nil")
	}
	_result := &_BACnetConstructedDataWriteStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		WriteStatus:                   writeStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataWriteStatus(structType any) BACnetConstructedDataWriteStatus {
	if casted, ok := structType.(BACnetConstructedDataWriteStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWriteStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWriteStatus) GetTypeName() string {
	return "BACnetConstructedDataWriteStatus"
}

func (m *_BACnetConstructedDataWriteStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (writeStatus)
	lengthInBits += m.WriteStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataWriteStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataWriteStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataWriteStatus BACnetConstructedDataWriteStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWriteStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWriteStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	writeStatus, err := ReadSimpleField[BACnetWriteStatusTagged](ctx, "writeStatus", ReadComplex[BACnetWriteStatusTagged](BACnetWriteStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeStatus' field"))
	}
	m.WriteStatus = writeStatus

	actualValue, err := ReadVirtualField[BACnetWriteStatusTagged](ctx, "actualValue", (*BACnetWriteStatusTagged)(nil), writeStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWriteStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWriteStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataWriteStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataWriteStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWriteStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWriteStatus")
		}

		if err := WriteSimpleField[BACnetWriteStatusTagged](ctx, "writeStatus", m.GetWriteStatus(), WriteComplex[BACnetWriteStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWriteStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWriteStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWriteStatus) IsBACnetConstructedDataWriteStatus() {}

func (m *_BACnetConstructedDataWriteStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
