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

// BACnetConstructedDataDateList is the corresponding interface of BACnetConstructedDataDateList
type BACnetConstructedDataDateList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDateList returns DateList (property field)
	GetDateList() []BACnetCalendarEntry
	// IsBACnetConstructedDataDateList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDateList()
	// CreateBuilder creates a BACnetConstructedDataDateListBuilder
	CreateBACnetConstructedDataDateListBuilder() BACnetConstructedDataDateListBuilder
}

// _BACnetConstructedDataDateList is the data-structure of this message
type _BACnetConstructedDataDateList struct {
	BACnetConstructedDataContract
	DateList []BACnetCalendarEntry
}

var _ BACnetConstructedDataDateList = (*_BACnetConstructedDataDateList)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDateList)(nil)

// NewBACnetConstructedDataDateList factory function for _BACnetConstructedDataDateList
func NewBACnetConstructedDataDateList(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, dateList []BACnetCalendarEntry, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateList {
	_result := &_BACnetConstructedDataDateList{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DateList:                      dateList,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDateListBuilder is a builder for BACnetConstructedDataDateList
type BACnetConstructedDataDateListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateList []BACnetCalendarEntry) BACnetConstructedDataDateListBuilder
	// WithDateList adds DateList (property field)
	WithDateList(...BACnetCalendarEntry) BACnetConstructedDataDateListBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDateList or returns an error if something is wrong
	Build() (BACnetConstructedDataDateList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDateList
}

// NewBACnetConstructedDataDateListBuilder() creates a BACnetConstructedDataDateListBuilder
func NewBACnetConstructedDataDateListBuilder() BACnetConstructedDataDateListBuilder {
	return &_BACnetConstructedDataDateListBuilder{_BACnetConstructedDataDateList: new(_BACnetConstructedDataDateList)}
}

type _BACnetConstructedDataDateListBuilder struct {
	*_BACnetConstructedDataDateList

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDateListBuilder) = (*_BACnetConstructedDataDateListBuilder)(nil)

func (b *_BACnetConstructedDataDateListBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDateList
}

func (b *_BACnetConstructedDataDateListBuilder) WithMandatoryFields(dateList []BACnetCalendarEntry) BACnetConstructedDataDateListBuilder {
	return b.WithDateList(dateList...)
}

func (b *_BACnetConstructedDataDateListBuilder) WithDateList(dateList ...BACnetCalendarEntry) BACnetConstructedDataDateListBuilder {
	b.DateList = dateList
	return b
}

func (b *_BACnetConstructedDataDateListBuilder) Build() (BACnetConstructedDataDateList, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDateList.deepCopy(), nil
}

func (b *_BACnetConstructedDataDateListBuilder) MustBuild() BACnetConstructedDataDateList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDateListBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDateListBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDateListBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDateListBuilder().(*_BACnetConstructedDataDateListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDateListBuilder creates a BACnetConstructedDataDateListBuilder
func (b *_BACnetConstructedDataDateList) CreateBACnetConstructedDataDateListBuilder() BACnetConstructedDataDateListBuilder {
	if b == nil {
		return NewBACnetConstructedDataDateListBuilder()
	}
	return &_BACnetConstructedDataDateListBuilder{_BACnetConstructedDataDateList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDateList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DATE_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateList) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateList) GetDateList() []BACnetCalendarEntry {
	return m.DateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateList(structType any) BACnetConstructedDataDateList {
	if casted, ok := structType.(BACnetConstructedDataDateList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateList) GetTypeName() string {
	return "BACnetConstructedDataDateList"
}

func (m *_BACnetConstructedDataDateList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.DateList) > 0 {
		for _, element := range m.DateList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataDateList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDateList BACnetConstructedDataDateList, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateList, err := ReadTerminatedArrayField[BACnetCalendarEntry](ctx, "dateList", ReadComplex[BACnetCalendarEntry](BACnetCalendarEntryParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateList' field"))
	}
	m.DateList = dateList

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateList")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateList")
		}

		if err := WriteComplexTypeArrayField(ctx, "dateList", m.GetDateList(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dateList' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateList")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateList) IsBACnetConstructedDataDateList() {}

func (m *_BACnetConstructedDataDateList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDateList) deepCopy() *_BACnetConstructedDataDateList {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDateListCopy := &_BACnetConstructedDataDateList{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetCalendarEntry, BACnetCalendarEntry](m.DateList),
	}
	_BACnetConstructedDataDateListCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDateListCopy
}

func (m *_BACnetConstructedDataDateList) String() string {
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
