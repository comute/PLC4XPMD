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

// MPropReadCon is the corresponding interface of MPropReadCon
type MPropReadCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
	// GetData returns Data (property field)
	GetData() uint16
	// IsMPropReadCon is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMPropReadCon()
	// CreateBuilder creates a MPropReadConBuilder
	CreateMPropReadConBuilder() MPropReadConBuilder
}

// _MPropReadCon is the data-structure of this message
type _MPropReadCon struct {
	CEMIContract
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
	Data                uint16
}

var _ MPropReadCon = (*_MPropReadCon)(nil)
var _ CEMIRequirements = (*_MPropReadCon)(nil)

// NewMPropReadCon factory function for _MPropReadCon
func NewMPropReadCon(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, data uint16, size uint16) *_MPropReadCon {
	_result := &_MPropReadCon{
		CEMIContract:        NewCEMI(size),
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Data:                data,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MPropReadConBuilder is a builder for MPropReadCon
type MPropReadConBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, data uint16) MPropReadConBuilder
	// WithInterfaceObjectType adds InterfaceObjectType (property field)
	WithInterfaceObjectType(uint16) MPropReadConBuilder
	// WithObjectInstance adds ObjectInstance (property field)
	WithObjectInstance(uint8) MPropReadConBuilder
	// WithPropertyId adds PropertyId (property field)
	WithPropertyId(uint8) MPropReadConBuilder
	// WithNumberOfElements adds NumberOfElements (property field)
	WithNumberOfElements(uint8) MPropReadConBuilder
	// WithStartIndex adds StartIndex (property field)
	WithStartIndex(uint16) MPropReadConBuilder
	// WithData adds Data (property field)
	WithData(uint16) MPropReadConBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the MPropReadCon or returns an error if something is wrong
	Build() (MPropReadCon, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MPropReadCon
}

// NewMPropReadConBuilder() creates a MPropReadConBuilder
func NewMPropReadConBuilder() MPropReadConBuilder {
	return &_MPropReadConBuilder{_MPropReadCon: new(_MPropReadCon)}
}

type _MPropReadConBuilder struct {
	*_MPropReadCon

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MPropReadConBuilder) = (*_MPropReadConBuilder)(nil)

func (b *_MPropReadConBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._MPropReadCon
}

func (b *_MPropReadConBuilder) WithMandatoryFields(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, data uint16) MPropReadConBuilder {
	return b.WithInterfaceObjectType(interfaceObjectType).WithObjectInstance(objectInstance).WithPropertyId(propertyId).WithNumberOfElements(numberOfElements).WithStartIndex(startIndex).WithData(data)
}

func (b *_MPropReadConBuilder) WithInterfaceObjectType(interfaceObjectType uint16) MPropReadConBuilder {
	b.InterfaceObjectType = interfaceObjectType
	return b
}

func (b *_MPropReadConBuilder) WithObjectInstance(objectInstance uint8) MPropReadConBuilder {
	b.ObjectInstance = objectInstance
	return b
}

func (b *_MPropReadConBuilder) WithPropertyId(propertyId uint8) MPropReadConBuilder {
	b.PropertyId = propertyId
	return b
}

func (b *_MPropReadConBuilder) WithNumberOfElements(numberOfElements uint8) MPropReadConBuilder {
	b.NumberOfElements = numberOfElements
	return b
}

func (b *_MPropReadConBuilder) WithStartIndex(startIndex uint16) MPropReadConBuilder {
	b.StartIndex = startIndex
	return b
}

func (b *_MPropReadConBuilder) WithData(data uint16) MPropReadConBuilder {
	b.Data = data
	return b
}

func (b *_MPropReadConBuilder) Build() (MPropReadCon, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MPropReadCon.deepCopy(), nil
}

func (b *_MPropReadConBuilder) MustBuild() MPropReadCon {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MPropReadConBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_MPropReadConBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MPropReadConBuilder) DeepCopy() any {
	_copy := b.CreateMPropReadConBuilder().(*_MPropReadConBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMPropReadConBuilder creates a MPropReadConBuilder
func (b *_MPropReadCon) CreateMPropReadConBuilder() MPropReadConBuilder {
	if b == nil {
		return NewMPropReadConBuilder()
	}
	return &_MPropReadConBuilder{_MPropReadCon: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropReadCon) GetMessageCode() uint8 {
	return 0xFB
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropReadCon) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MPropReadCon) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *_MPropReadCon) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *_MPropReadCon) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_MPropReadCon) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *_MPropReadCon) GetStartIndex() uint16 {
	return m.StartIndex
}

func (m *_MPropReadCon) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMPropReadCon(structType any) MPropReadCon {
	if casted, ok := structType.(MPropReadCon); ok {
		return casted
	}
	if casted, ok := structType.(*MPropReadCon); ok {
		return *casted
	}
	return nil
}

func (m *_MPropReadCon) GetTypeName() string {
	return "MPropReadCon"
}

func (m *_MPropReadCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_MPropReadCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropReadCon) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropReadCon MPropReadCon, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropReadCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropReadCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceObjectType, err := ReadSimpleField(ctx, "interfaceObjectType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceObjectType' field"))
	}
	m.InterfaceObjectType = interfaceObjectType

	objectInstance, err := ReadSimpleField(ctx, "objectInstance", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectInstance' field"))
	}
	m.ObjectInstance = objectInstance

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	numberOfElements, err := ReadSimpleField(ctx, "numberOfElements", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfElements' field"))
	}
	m.NumberOfElements = numberOfElements

	startIndex, err := ReadSimpleField(ctx, "startIndex", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startIndex' field"))
	}
	m.StartIndex = startIndex

	data, err := ReadSimpleField(ctx, "data", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("MPropReadCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropReadCon")
	}

	return m, nil
}

func (m *_MPropReadCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropReadCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropReadCon")
		}

		if err := WriteSimpleField[uint16](ctx, "interfaceObjectType", m.GetInterfaceObjectType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceObjectType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "objectInstance", m.GetObjectInstance(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectInstance' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "numberOfElements", m.GetNumberOfElements(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfElements' field")
		}

		if err := WriteSimpleField[uint16](ctx, "startIndex", m.GetStartIndex(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'startIndex' field")
		}

		if err := WriteSimpleField[uint16](ctx, "data", m.GetData(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropReadCon")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropReadCon) IsMPropReadCon() {}

func (m *_MPropReadCon) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MPropReadCon) deepCopy() *_MPropReadCon {
	if m == nil {
		return nil
	}
	_MPropReadConCopy := &_MPropReadCon{
		m.CEMIContract.(*_CEMI).deepCopy(),
		m.InterfaceObjectType,
		m.ObjectInstance,
		m.PropertyId,
		m.NumberOfElements,
		m.StartIndex,
		m.Data,
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MPropReadConCopy
}

func (m *_MPropReadCon) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
