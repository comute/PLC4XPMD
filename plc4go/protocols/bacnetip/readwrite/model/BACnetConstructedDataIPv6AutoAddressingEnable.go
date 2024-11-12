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

// BACnetConstructedDataIPv6AutoAddressingEnable is the corresponding interface of BACnetConstructedDataIPv6AutoAddressingEnable
type BACnetConstructedDataIPv6AutoAddressingEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAutoAddressingEnable returns AutoAddressingEnable (property field)
	GetAutoAddressingEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataIPv6AutoAddressingEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6AutoAddressingEnable()
	// CreateBuilder creates a BACnetConstructedDataIPv6AutoAddressingEnableBuilder
	CreateBACnetConstructedDataIPv6AutoAddressingEnableBuilder() BACnetConstructedDataIPv6AutoAddressingEnableBuilder
}

// _BACnetConstructedDataIPv6AutoAddressingEnable is the data-structure of this message
type _BACnetConstructedDataIPv6AutoAddressingEnable struct {
	BACnetConstructedDataContract
	AutoAddressingEnable BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataIPv6AutoAddressingEnable = (*_BACnetConstructedDataIPv6AutoAddressingEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6AutoAddressingEnable)(nil)

// NewBACnetConstructedDataIPv6AutoAddressingEnable factory function for _BACnetConstructedDataIPv6AutoAddressingEnable
func NewBACnetConstructedDataIPv6AutoAddressingEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, autoAddressingEnable BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6AutoAddressingEnable {
	if autoAddressingEnable == nil {
		panic("autoAddressingEnable of type BACnetApplicationTagBoolean for BACnetConstructedDataIPv6AutoAddressingEnable must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6AutoAddressingEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AutoAddressingEnable:          autoAddressingEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6AutoAddressingEnableBuilder is a builder for BACnetConstructedDataIPv6AutoAddressingEnable
type BACnetConstructedDataIPv6AutoAddressingEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(autoAddressingEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPv6AutoAddressingEnableBuilder
	// WithAutoAddressingEnable adds AutoAddressingEnable (property field)
	WithAutoAddressingEnable(BACnetApplicationTagBoolean) BACnetConstructedDataIPv6AutoAddressingEnableBuilder
	// WithAutoAddressingEnableBuilder adds AutoAddressingEnable (property field) which is build by the builder
	WithAutoAddressingEnableBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIPv6AutoAddressingEnableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPv6AutoAddressingEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6AutoAddressingEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6AutoAddressingEnable
}

// NewBACnetConstructedDataIPv6AutoAddressingEnableBuilder() creates a BACnetConstructedDataIPv6AutoAddressingEnableBuilder
func NewBACnetConstructedDataIPv6AutoAddressingEnableBuilder() BACnetConstructedDataIPv6AutoAddressingEnableBuilder {
	return &_BACnetConstructedDataIPv6AutoAddressingEnableBuilder{_BACnetConstructedDataIPv6AutoAddressingEnable: new(_BACnetConstructedDataIPv6AutoAddressingEnable)}
}

type _BACnetConstructedDataIPv6AutoAddressingEnableBuilder struct {
	*_BACnetConstructedDataIPv6AutoAddressingEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6AutoAddressingEnableBuilder) = (*_BACnetConstructedDataIPv6AutoAddressingEnableBuilder)(nil)

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPv6AutoAddressingEnable
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) WithMandatoryFields(autoAddressingEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPv6AutoAddressingEnableBuilder {
	return b.WithAutoAddressingEnable(autoAddressingEnable)
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) WithAutoAddressingEnable(autoAddressingEnable BACnetApplicationTagBoolean) BACnetConstructedDataIPv6AutoAddressingEnableBuilder {
	b.AutoAddressingEnable = autoAddressingEnable
	return b
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) WithAutoAddressingEnableBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIPv6AutoAddressingEnableBuilder {
	builder := builderSupplier(b.AutoAddressingEnable.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.AutoAddressingEnable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) Build() (BACnetConstructedDataIPv6AutoAddressingEnable, error) {
	if b.AutoAddressingEnable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'autoAddressingEnable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6AutoAddressingEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) MustBuild() BACnetConstructedDataIPv6AutoAddressingEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6AutoAddressingEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6AutoAddressingEnableBuilder().(*_BACnetConstructedDataIPv6AutoAddressingEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6AutoAddressingEnableBuilder creates a BACnetConstructedDataIPv6AutoAddressingEnableBuilder
func (b *_BACnetConstructedDataIPv6AutoAddressingEnable) CreateBACnetConstructedDataIPv6AutoAddressingEnableBuilder() BACnetConstructedDataIPv6AutoAddressingEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6AutoAddressingEnableBuilder()
	}
	return &_BACnetConstructedDataIPv6AutoAddressingEnableBuilder{_BACnetConstructedDataIPv6AutoAddressingEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_AUTO_ADDRESSING_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetAutoAddressingEnable() BACnetApplicationTagBoolean {
	return m.AutoAddressingEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetAutoAddressingEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6AutoAddressingEnable(structType any) BACnetConstructedDataIPv6AutoAddressingEnable {
	if casted, ok := structType.(BACnetConstructedDataIPv6AutoAddressingEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6AutoAddressingEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetTypeName() string {
	return "BACnetConstructedDataIPv6AutoAddressingEnable"
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (autoAddressingEnable)
	lengthInBits += m.AutoAddressingEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6AutoAddressingEnable BACnetConstructedDataIPv6AutoAddressingEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6AutoAddressingEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6AutoAddressingEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	autoAddressingEnable, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "autoAddressingEnable", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'autoAddressingEnable' field"))
	}
	m.AutoAddressingEnable = autoAddressingEnable

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), autoAddressingEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6AutoAddressingEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6AutoAddressingEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6AutoAddressingEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6AutoAddressingEnable")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "autoAddressingEnable", m.GetAutoAddressingEnable(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'autoAddressingEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6AutoAddressingEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6AutoAddressingEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) IsBACnetConstructedDataIPv6AutoAddressingEnable() {
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) deepCopy() *_BACnetConstructedDataIPv6AutoAddressingEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6AutoAddressingEnableCopy := &_BACnetConstructedDataIPv6AutoAddressingEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.AutoAddressingEnable),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6AutoAddressingEnableCopy
}

func (m *_BACnetConstructedDataIPv6AutoAddressingEnable) String() string {
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
