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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataReliabilityEvaluationInhibit is the corresponding interface of BACnetConstructedDataReliabilityEvaluationInhibit
type BACnetConstructedDataReliabilityEvaluationInhibit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReliabilityEvaluationInhibit returns ReliabilityEvaluationInhibit (property field)
	GetReliabilityEvaluationInhibit() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataReliabilityEvaluationInhibitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataReliabilityEvaluationInhibit.
// This is useful for switch cases.
type BACnetConstructedDataReliabilityEvaluationInhibitExactly interface {
	BACnetConstructedDataReliabilityEvaluationInhibit
	isBACnetConstructedDataReliabilityEvaluationInhibit() bool
}

// _BACnetConstructedDataReliabilityEvaluationInhibit is the data-structure of this message
type _BACnetConstructedDataReliabilityEvaluationInhibit struct {
	*_BACnetConstructedData
	ReliabilityEvaluationInhibit BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELIABILITY_EVALUATION_INHIBIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetReliabilityEvaluationInhibit() BACnetApplicationTagBoolean {
	return m.ReliabilityEvaluationInhibit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetReliabilityEvaluationInhibit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReliabilityEvaluationInhibit factory function for _BACnetConstructedDataReliabilityEvaluationInhibit
func NewBACnetConstructedDataReliabilityEvaluationInhibit(reliabilityEvaluationInhibit BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReliabilityEvaluationInhibit {
	_result := &_BACnetConstructedDataReliabilityEvaluationInhibit{
		ReliabilityEvaluationInhibit: reliabilityEvaluationInhibit,
		_BACnetConstructedData:       NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReliabilityEvaluationInhibit(structType interface{}) BACnetConstructedDataReliabilityEvaluationInhibit {
	if casted, ok := structType.(BACnetConstructedDataReliabilityEvaluationInhibit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliabilityEvaluationInhibit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetTypeName() string {
	return "BACnetConstructedDataReliabilityEvaluationInhibit"
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reliabilityEvaluationInhibit)
	lengthInBits += m.ReliabilityEvaluationInhibit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataReliabilityEvaluationInhibitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReliabilityEvaluationInhibit, error) {
	return BACnetConstructedDataReliabilityEvaluationInhibitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataReliabilityEvaluationInhibitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReliabilityEvaluationInhibit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliabilityEvaluationInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReliabilityEvaluationInhibit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reliabilityEvaluationInhibit)
	if pullErr := readBuffer.PullContext("reliabilityEvaluationInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reliabilityEvaluationInhibit")
	}
	_reliabilityEvaluationInhibit, _reliabilityEvaluationInhibitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _reliabilityEvaluationInhibitErr != nil {
		return nil, errors.Wrap(_reliabilityEvaluationInhibitErr, "Error parsing 'reliabilityEvaluationInhibit' field of BACnetConstructedDataReliabilityEvaluationInhibit")
	}
	reliabilityEvaluationInhibit := _reliabilityEvaluationInhibit.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("reliabilityEvaluationInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reliabilityEvaluationInhibit")
	}

	// Virtual field
	_actualValue := reliabilityEvaluationInhibit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliabilityEvaluationInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReliabilityEvaluationInhibit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataReliabilityEvaluationInhibit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ReliabilityEvaluationInhibit: reliabilityEvaluationInhibit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliabilityEvaluationInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReliabilityEvaluationInhibit")
		}

		// Simple Field (reliabilityEvaluationInhibit)
		if pushErr := writeBuffer.PushContext("reliabilityEvaluationInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reliabilityEvaluationInhibit")
		}
		_reliabilityEvaluationInhibitErr := writeBuffer.WriteSerializable(ctx, m.GetReliabilityEvaluationInhibit())
		if popErr := writeBuffer.PopContext("reliabilityEvaluationInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reliabilityEvaluationInhibit")
		}
		if _reliabilityEvaluationInhibitErr != nil {
			return errors.Wrap(_reliabilityEvaluationInhibitErr, "Error serializing 'reliabilityEvaluationInhibit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReliabilityEvaluationInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReliabilityEvaluationInhibit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) isBACnetConstructedDataReliabilityEvaluationInhibit() bool {
	return true
}

func (m *_BACnetConstructedDataReliabilityEvaluationInhibit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
