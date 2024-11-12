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

// AirConditioningDataZoneHvacPlantStatus is the corresponding interface of AirConditioningDataZoneHvacPlantStatus
type AirConditioningDataZoneHvacPlantStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetHvacStatus returns HvacStatus (property field)
	GetHvacStatus() HVACStatusFlags
	// GetHvacErrorCode returns HvacErrorCode (property field)
	GetHvacErrorCode() HVACError
	// IsAirConditioningDataZoneHvacPlantStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataZoneHvacPlantStatus()
	// CreateBuilder creates a AirConditioningDataZoneHvacPlantStatusBuilder
	CreateAirConditioningDataZoneHvacPlantStatusBuilder() AirConditioningDataZoneHvacPlantStatusBuilder
}

// _AirConditioningDataZoneHvacPlantStatus is the data-structure of this message
type _AirConditioningDataZoneHvacPlantStatus struct {
	AirConditioningDataContract
	ZoneGroup     byte
	ZoneList      HVACZoneList
	HvacType      HVACType
	HvacStatus    HVACStatusFlags
	HvacErrorCode HVACError
}

var _ AirConditioningDataZoneHvacPlantStatus = (*_AirConditioningDataZoneHvacPlantStatus)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneHvacPlantStatus)(nil)

// NewAirConditioningDataZoneHvacPlantStatus factory function for _AirConditioningDataZoneHvacPlantStatus
func NewAirConditioningDataZoneHvacPlantStatus(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, hvacType HVACType, hvacStatus HVACStatusFlags, hvacErrorCode HVACError) *_AirConditioningDataZoneHvacPlantStatus {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataZoneHvacPlantStatus must not be nil")
	}
	if hvacStatus == nil {
		panic("hvacStatus of type HVACStatusFlags for AirConditioningDataZoneHvacPlantStatus must not be nil")
	}
	_result := &_AirConditioningDataZoneHvacPlantStatus{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HvacType:                    hvacType,
		HvacStatus:                  hvacStatus,
		HvacErrorCode:               hvacErrorCode,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataZoneHvacPlantStatusBuilder is a builder for AirConditioningDataZoneHvacPlantStatus
type AirConditioningDataZoneHvacPlantStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, hvacType HVACType, hvacStatus HVACStatusFlags, hvacErrorCode HVACError) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithHvacType adds HvacType (property field)
	WithHvacType(HVACType) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithHvacStatus adds HvacStatus (property field)
	WithHvacStatus(HVACStatusFlags) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithHvacStatusBuilder adds HvacStatus (property field) which is build by the builder
	WithHvacStatusBuilder(func(HVACStatusFlagsBuilder) HVACStatusFlagsBuilder) AirConditioningDataZoneHvacPlantStatusBuilder
	// WithHvacErrorCode adds HvacErrorCode (property field)
	WithHvacErrorCode(HVACError) AirConditioningDataZoneHvacPlantStatusBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataZoneHvacPlantStatus or returns an error if something is wrong
	Build() (AirConditioningDataZoneHvacPlantStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataZoneHvacPlantStatus
}

// NewAirConditioningDataZoneHvacPlantStatusBuilder() creates a AirConditioningDataZoneHvacPlantStatusBuilder
func NewAirConditioningDataZoneHvacPlantStatusBuilder() AirConditioningDataZoneHvacPlantStatusBuilder {
	return &_AirConditioningDataZoneHvacPlantStatusBuilder{_AirConditioningDataZoneHvacPlantStatus: new(_AirConditioningDataZoneHvacPlantStatus)}
}

type _AirConditioningDataZoneHvacPlantStatusBuilder struct {
	*_AirConditioningDataZoneHvacPlantStatus

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataZoneHvacPlantStatusBuilder) = (*_AirConditioningDataZoneHvacPlantStatusBuilder)(nil)

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
	contract.(*_AirConditioningData)._SubType = b._AirConditioningDataZoneHvacPlantStatus
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, hvacType HVACType, hvacStatus HVACStatusFlags, hvacErrorCode HVACError) AirConditioningDataZoneHvacPlantStatusBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHvacType(hvacType).WithHvacStatus(hvacStatus).WithHvacErrorCode(hvacErrorCode)
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataZoneHvacPlantStatusBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataZoneHvacPlantStatusBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataZoneHvacPlantStatusBuilder {
	builder := builderSupplier(b.ZoneList.CreateHVACZoneListBuilder())
	var err error
	b.ZoneList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithHvacType(hvacType HVACType) AirConditioningDataZoneHvacPlantStatusBuilder {
	b.HvacType = hvacType
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithHvacStatus(hvacStatus HVACStatusFlags) AirConditioningDataZoneHvacPlantStatusBuilder {
	b.HvacStatus = hvacStatus
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithHvacStatusBuilder(builderSupplier func(HVACStatusFlagsBuilder) HVACStatusFlagsBuilder) AirConditioningDataZoneHvacPlantStatusBuilder {
	builder := builderSupplier(b.HvacStatus.CreateHVACStatusFlagsBuilder())
	var err error
	b.HvacStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACStatusFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) WithHvacErrorCode(hvacErrorCode HVACError) AirConditioningDataZoneHvacPlantStatusBuilder {
	b.HvacErrorCode = hvacErrorCode
	return b
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) Build() (AirConditioningDataZoneHvacPlantStatus, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.HvacStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hvacStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataZoneHvacPlantStatus.deepCopy(), nil
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) MustBuild() AirConditioningDataZoneHvacPlantStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataZoneHvacPlantStatusBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataZoneHvacPlantStatusBuilder().(*_AirConditioningDataZoneHvacPlantStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataZoneHvacPlantStatusBuilder creates a AirConditioningDataZoneHvacPlantStatusBuilder
func (b *_AirConditioningDataZoneHvacPlantStatus) CreateAirConditioningDataZoneHvacPlantStatusBuilder() AirConditioningDataZoneHvacPlantStatusBuilder {
	if b == nil {
		return NewAirConditioningDataZoneHvacPlantStatusBuilder()
	}
	return &_AirConditioningDataZoneHvacPlantStatusBuilder{_AirConditioningDataZoneHvacPlantStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacStatus() HVACStatusFlags {
	return m.HvacStatus
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacErrorCode() HVACError {
	return m.HvacErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHvacPlantStatus(structType any) AirConditioningDataZoneHvacPlantStatus {
	if casted, ok := structType.(AirConditioningDataZoneHvacPlantStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHvacPlantStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetTypeName() string {
	return "AirConditioningDataZoneHvacPlantStatus"
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (hvacType)
	lengthInBits += 8

	// Simple field (hvacStatus)
	lengthInBits += m.HvacStatus.GetLengthInBits(ctx)

	// Simple field (hvacErrorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneHvacPlantStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneHvacPlantStatus AirConditioningDataZoneHvacPlantStatus, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHvacPlantStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHvacPlantStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	hvacType, err := ReadEnumField[HVACType](ctx, "hvacType", "HVACType", ReadEnum(HVACTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacType' field"))
	}
	m.HvacType = hvacType

	hvacStatus, err := ReadSimpleField[HVACStatusFlags](ctx, "hvacStatus", ReadComplex[HVACStatusFlags](HVACStatusFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacStatus' field"))
	}
	m.HvacStatus = hvacStatus

	hvacErrorCode, err := ReadEnumField[HVACError](ctx, "hvacErrorCode", "HVACError", ReadEnum(HVACErrorByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacErrorCode' field"))
	}
	m.HvacErrorCode = hvacErrorCode

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHvacPlantStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHvacPlantStatus")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHvacPlantStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHvacPlantStatus")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleEnumField[HVACType](ctx, "hvacType", "HVACType", m.GetHvacType(), WriteEnum[HVACType, uint8](HVACType.GetValue, HVACType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacType' field")
		}

		if err := WriteSimpleField[HVACStatusFlags](ctx, "hvacStatus", m.GetHvacStatus(), WriteComplex[HVACStatusFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacStatus' field")
		}

		if err := WriteSimpleEnumField[HVACError](ctx, "hvacErrorCode", "HVACError", m.GetHvacErrorCode(), WriteEnum[HVACError, uint8](HVACError.GetValue, HVACError.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacErrorCode' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHvacPlantStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHvacPlantStatus")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHvacPlantStatus) IsAirConditioningDataZoneHvacPlantStatus() {}

func (m *_AirConditioningDataZoneHvacPlantStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataZoneHvacPlantStatus) deepCopy() *_AirConditioningDataZoneHvacPlantStatus {
	if m == nil {
		return nil
	}
	_AirConditioningDataZoneHvacPlantStatusCopy := &_AirConditioningDataZoneHvacPlantStatus{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		m.HvacType,
		utils.DeepCopy[HVACStatusFlags](m.HvacStatus),
		m.HvacErrorCode,
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataZoneHvacPlantStatusCopy
}

func (m *_AirConditioningDataZoneHvacPlantStatus) String() string {
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
