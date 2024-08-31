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

// APDUComplexAck is the corresponding interface of APDUComplexAck
type APDUComplexAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceAck returns ServiceAck (property field)
	GetServiceAck() BACnetServiceAck
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *BACnetConfirmedServiceChoice
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
}

// APDUComplexAckExactly can be used when we want exactly this type and not a type which fulfills APDUComplexAck.
// This is useful for switch cases.
type APDUComplexAckExactly interface {
	APDUComplexAck
	isAPDUComplexAck() bool
}

// _APDUComplexAck is the data-structure of this message
type _APDUComplexAck struct {
	*_APDU
	SegmentedMessage     bool
	MoreFollows          bool
	OriginalInvokeId     uint8
	SequenceNumber       *uint8
	ProposedWindowSize   *uint8
	ServiceAck           BACnetServiceAck
	SegmentServiceChoice *BACnetConfirmedServiceChoice
	Segment              []byte
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUComplexAck) GetApduType() ApduType {
	return ApduType_COMPLEX_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUComplexAck) InitializeParent(parent APDU) {}

func (m *_APDUComplexAck) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUComplexAck) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *_APDUComplexAck) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *_APDUComplexAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUComplexAck) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *_APDUComplexAck) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *_APDUComplexAck) GetServiceAck() BACnetServiceAck {
	return m.ServiceAck
}

func (m *_APDUComplexAck) GetSegmentServiceChoice() *BACnetConfirmedServiceChoice {
	return m.SegmentServiceChoice
}

func (m *_APDUComplexAck) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_APDUComplexAck) GetApduHeaderReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceAck := m.ServiceAck
	_ = serviceAck
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(uint16(uint16(2)) + uint16((utils.InlineIf(m.GetSegmentedMessage(), func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
}

func (m *_APDUComplexAck) GetSegmentReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceAck := m.ServiceAck
	_ = serviceAck
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(utils.InlineIf((bool((m.GetSegmentServiceChoice()) != (nil))), func() any { return uint16((uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() any { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUComplexAck factory function for _APDUComplexAck
func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck BACnetServiceAck, segmentServiceChoice *BACnetConfirmedServiceChoice, segment []byte, apduLength uint16) *_APDUComplexAck {
	_result := &_APDUComplexAck{
		SegmentedMessage:     segmentedMessage,
		MoreFollows:          moreFollows,
		OriginalInvokeId:     originalInvokeId,
		SequenceNumber:       sequenceNumber,
		ProposedWindowSize:   proposedWindowSize,
		ServiceAck:           serviceAck,
		SegmentServiceChoice: segmentServiceChoice,
		Segment:              segment,
		_APDU:                NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUComplexAck(structType any) APDUComplexAck {
	if casted, ok := structType.(APDUComplexAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUComplexAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUComplexAck) GetTypeName() string {
	return "APDUComplexAck"
}

func (m *_APDUComplexAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceAck)
	if m.ServiceAck != nil {
		lengthInBits += m.ServiceAck.GetLengthInBits(ctx)
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *_APDUComplexAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDUComplexAckParse(ctx context.Context, theBytes []byte, apduLength uint16) (APDUComplexAck, error) {
	return APDUComplexAckParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDUComplexAckParseWithBufferProducer(apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUComplexAck, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (APDUComplexAck, error) {
		return APDUComplexAckParseWithBuffer(ctx, readBuffer, apduLength)
	}
}

func APDUComplexAckParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUComplexAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUComplexAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUComplexAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentedMessage, err := ReadSimpleField(ctx, "segmentedMessage", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedMessage' field"))
	}

	moreFollows, err := ReadSimpleField(ctx, "moreFollows", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreFollows' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(2)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}

	sequenceNumber, err := ReadOptionalField[uint8](ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}

	proposedWindowSize, err := ReadOptionalField[uint8](ctx, "proposedWindowSize", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proposedWindowSize' field"))
	}

	apduHeaderReduction, err := ReadVirtualField[uint16](ctx, "apduHeaderReduction", (*uint16)(nil), uint16(uint16(2))+uint16((utils.InlineIf(segmentedMessage, func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduHeaderReduction' field"))
	}
	_ = apduHeaderReduction

	_serviceAck, err := ReadOptionalField[BACnetServiceAck](ctx, "serviceAck", ReadComplex[BACnetServiceAck](BACnetServiceAckParseWithBufferProducer[BACnetServiceAck]((uint32)(uint32(apduLength)-uint32(apduHeaderReduction))), readBuffer), !(segmentedMessage))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceAck' field"))
	}
	var serviceAck BACnetServiceAck
	if _serviceAck != nil {
		serviceAck = *_serviceAck
	}

	// Validation
	if !(bool((bool(!(segmentedMessage)) && bool(bool((serviceAck) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "service ack should be set"})
	}

	segmentServiceChoice, err := ReadOptionalField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentServiceChoice' field"))
	}

	segmentReduction, err := ReadVirtualField[uint16](ctx, "segmentReduction", (*uint16)(nil), utils.InlineIf((bool((segmentServiceChoice) != (nil))), func() any { return uint16((uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() any { return uint16(apduHeaderReduction) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentReduction' field"))
	}
	_ = segmentReduction

	segment, err := readBuffer.ReadByteArray("segment", int(utils.InlineIf(segmentedMessage, func() any {
		return int32((utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(segmentReduction))) }, func() any { return int32(int32(0)) }).(int32)))
	}, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segment' field"))
	}

	if closeErr := readBuffer.CloseContext("APDUComplexAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUComplexAck")
	}

	// Create a partially initialized instance
	_child := &_APDUComplexAck{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		SegmentedMessage:     segmentedMessage,
		MoreFollows:          moreFollows,
		OriginalInvokeId:     originalInvokeId,
		SequenceNumber:       sequenceNumber,
		ProposedWindowSize:   proposedWindowSize,
		ServiceAck:           serviceAck,
		SegmentServiceChoice: segmentServiceChoice,
		Segment:              segment,
		reservedField0:       reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUComplexAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUComplexAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUComplexAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUComplexAck")
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.GetSegmentedMessage())
		_segmentedMessageErr := /*TODO: migrate me*/ writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.GetMoreFollows())
		_moreFollowsErr := /*TODO: migrate me*/ writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
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
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 2, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("originalInvokeId", 8, uint8((originalInvokeId)))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.GetSequenceNumber() != nil {
			sequenceNumber = m.GetSequenceNumber()
			_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("sequenceNumber", 8, uint8(*(sequenceNumber)))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.GetProposedWindowSize() != nil {
			proposedWindowSize = m.GetProposedWindowSize()
			_proposedWindowSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("proposedWindowSize", 8, uint8(*(proposedWindowSize)))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}
		// Virtual field
		apduHeaderReduction := m.GetApduHeaderReduction()
		_ = apduHeaderReduction
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual(ctx, "apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		// Optional Field (serviceAck) (Can be skipped, if the value is null)
		var serviceAck BACnetServiceAck = nil
		if m.GetServiceAck() != nil {
			if pushErr := writeBuffer.PushContext("serviceAck"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceAck")
			}
			serviceAck = m.GetServiceAck()
			_serviceAckErr := writeBuffer.WriteSerializable(ctx, serviceAck)
			if popErr := writeBuffer.PopContext("serviceAck"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceAck")
			}
			if _serviceAckErr != nil {
				return errors.Wrap(_serviceAckErr, "Error serializing 'serviceAck' field")
			}
		}

		// Optional Field (segmentServiceChoice) (Can be skipped, if the value is null)
		var segmentServiceChoice *BACnetConfirmedServiceChoice = nil
		if m.GetSegmentServiceChoice() != nil {
			if pushErr := writeBuffer.PushContext("segmentServiceChoice"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for segmentServiceChoice")
			}
			segmentServiceChoice = m.GetSegmentServiceChoice()
			_segmentServiceChoiceErr := writeBuffer.WriteSerializable(ctx, segmentServiceChoice)
			if popErr := writeBuffer.PopContext("segmentServiceChoice"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for segmentServiceChoice")
			}
			if _segmentServiceChoiceErr != nil {
				return errors.Wrap(_segmentServiceChoiceErr, "Error serializing 'segmentServiceChoice' field")
			}
		}
		// Virtual field
		segmentReduction := m.GetSegmentReduction()
		_ = segmentReduction
		if _segmentReductionErr := writeBuffer.WriteVirtual(ctx, "segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		if err := WriteByteArrayField(ctx, "segment", m.GetSegment(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'segment' field")
		}

		if popErr := writeBuffer.PopContext("APDUComplexAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUComplexAck")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUComplexAck) isAPDUComplexAck() bool {
	return true
}

func (m *_APDUComplexAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
