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

// BACnetConstructedDataInProgress is the corresponding interface of BACnetConstructedDataInProgress
type BACnetConstructedDataInProgress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInProgress returns InProgress (property field)
	GetInProgress() BACnetLightingInProgressTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingInProgressTagged
	// IsBACnetConstructedDataInProgress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInProgress()
}

// _BACnetConstructedDataInProgress is the data-structure of this message
type _BACnetConstructedDataInProgress struct {
	BACnetConstructedDataContract
	InProgress BACnetLightingInProgressTagged
}

var _ BACnetConstructedDataInProgress = (*_BACnetConstructedDataInProgress)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInProgress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInProgress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IN_PROGRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInProgress) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetInProgress() BACnetLightingInProgressTagged {
	return m.InProgress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInProgress) GetActualValue() BACnetLightingInProgressTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLightingInProgressTagged(m.GetInProgress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInProgress factory function for _BACnetConstructedDataInProgress
func NewBACnetConstructedDataInProgress(inProgress BACnetLightingInProgressTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInProgress {
	if inProgress == nil {
		panic("inProgress of type BACnetLightingInProgressTagged for BACnetConstructedDataInProgress must not be nil")
	}
	_result := &_BACnetConstructedDataInProgress{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InProgress:                    inProgress,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInProgress(structType any) BACnetConstructedDataInProgress {
	if casted, ok := structType.(BACnetConstructedDataInProgress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInProgress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInProgress) GetTypeName() string {
	return "BACnetConstructedDataInProgress"
}

func (m *_BACnetConstructedDataInProgress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (inProgress)
	lengthInBits += m.InProgress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInProgress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInProgress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInProgress BACnetConstructedDataInProgress, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInProgress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	inProgress, err := ReadSimpleField[BACnetLightingInProgressTagged](ctx, "inProgress", ReadComplex[BACnetLightingInProgressTagged](BACnetLightingInProgressTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inProgress' field"))
	}
	m.InProgress = inProgress

	actualValue, err := ReadVirtualField[BACnetLightingInProgressTagged](ctx, "actualValue", (*BACnetLightingInProgressTagged)(nil), inProgress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInProgress")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInProgress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInProgress")
		}

		if err := WriteSimpleField[BACnetLightingInProgressTagged](ctx, "inProgress", m.GetInProgress(), WriteComplex[BACnetLightingInProgressTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'inProgress' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInProgress")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInProgress) IsBACnetConstructedDataInProgress() {}

func (m *_BACnetConstructedDataInProgress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
