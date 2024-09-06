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

// BACnetConstructedDataAcceptedModes is the corresponding interface of BACnetConstructedDataAcceptedModes
type BACnetConstructedDataAcceptedModes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAcceptedModes returns AcceptedModes (property field)
	GetAcceptedModes() []BACnetLifeSafetyModeTagged
	// IsBACnetConstructedDataAcceptedModes is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAcceptedModes()
}

// _BACnetConstructedDataAcceptedModes is the data-structure of this message
type _BACnetConstructedDataAcceptedModes struct {
	BACnetConstructedDataContract
	AcceptedModes []BACnetLifeSafetyModeTagged
}

var _ BACnetConstructedDataAcceptedModes = (*_BACnetConstructedDataAcceptedModes)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAcceptedModes)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAcceptedModes) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAcceptedModes) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCEPTED_MODES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAcceptedModes) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAcceptedModes) GetAcceptedModes() []BACnetLifeSafetyModeTagged {
	return m.AcceptedModes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAcceptedModes factory function for _BACnetConstructedDataAcceptedModes
func NewBACnetConstructedDataAcceptedModes(acceptedModes []BACnetLifeSafetyModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAcceptedModes {
	_result := &_BACnetConstructedDataAcceptedModes{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AcceptedModes:                 acceptedModes,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAcceptedModes(structType any) BACnetConstructedDataAcceptedModes {
	if casted, ok := structType.(BACnetConstructedDataAcceptedModes); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAcceptedModes); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAcceptedModes) GetTypeName() string {
	return "BACnetConstructedDataAcceptedModes"
}

func (m *_BACnetConstructedDataAcceptedModes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AcceptedModes) > 0 {
		for _, element := range m.AcceptedModes {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAcceptedModes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAcceptedModes) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAcceptedModes BACnetConstructedDataAcceptedModes, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAcceptedModes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAcceptedModes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	acceptedModes, err := ReadTerminatedArrayField[BACnetLifeSafetyModeTagged](ctx, "acceptedModes", ReadComplex[BACnetLifeSafetyModeTagged](BACnetLifeSafetyModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acceptedModes' field"))
	}
	m.AcceptedModes = acceptedModes

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAcceptedModes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAcceptedModes")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAcceptedModes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAcceptedModes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAcceptedModes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAcceptedModes")
		}

		if err := WriteComplexTypeArrayField(ctx, "acceptedModes", m.GetAcceptedModes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'acceptedModes' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAcceptedModes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAcceptedModes")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAcceptedModes) IsBACnetConstructedDataAcceptedModes() {}

func (m *_BACnetConstructedDataAcceptedModes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
