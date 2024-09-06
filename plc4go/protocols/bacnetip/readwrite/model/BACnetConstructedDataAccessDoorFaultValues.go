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

// BACnetConstructedDataAccessDoorFaultValues is the corresponding interface of BACnetConstructedDataAccessDoorFaultValues
type BACnetConstructedDataAccessDoorFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []BACnetDoorAlarmStateTagged
	// IsBACnetConstructedDataAccessDoorFaultValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessDoorFaultValues()
}

// _BACnetConstructedDataAccessDoorFaultValues is the data-structure of this message
type _BACnetConstructedDataAccessDoorFaultValues struct {
	BACnetConstructedDataContract
	FaultValues []BACnetDoorAlarmStateTagged
}

var _ BACnetConstructedDataAccessDoorFaultValues = (*_BACnetConstructedDataAccessDoorFaultValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessDoorFaultValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetFaultValues() []BACnetDoorAlarmStateTagged {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessDoorFaultValues factory function for _BACnetConstructedDataAccessDoorFaultValues
func NewBACnetConstructedDataAccessDoorFaultValues(faultValues []BACnetDoorAlarmStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorFaultValues {
	_result := &_BACnetConstructedDataAccessDoorFaultValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultValues:                   faultValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorFaultValues(structType any) BACnetConstructedDataAccessDoorFaultValues {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorFaultValues"
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorFaultValues BACnetConstructedDataAccessDoorFaultValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultValues, err := ReadTerminatedArrayField[BACnetDoorAlarmStateTagged](ctx, "faultValues", ReadComplex[BACnetDoorAlarmStateTagged](BACnetDoorAlarmStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultValues' field"))
	}
	m.FaultValues = faultValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorFaultValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorFaultValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultValues", m.GetFaultValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorFaultValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) IsBACnetConstructedDataAccessDoorFaultValues() {
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
