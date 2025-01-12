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

// BACnetConstructedDataBackupFailureTimeout is the corresponding interface of BACnetConstructedDataBackupFailureTimeout
type BACnetConstructedDataBackupFailureTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBackupFailureTimeout returns BackupFailureTimeout (property field)
	GetBackupFailureTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataBackupFailureTimeout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBackupFailureTimeout()
	// CreateBuilder creates a BACnetConstructedDataBackupFailureTimeoutBuilder
	CreateBACnetConstructedDataBackupFailureTimeoutBuilder() BACnetConstructedDataBackupFailureTimeoutBuilder
}

// _BACnetConstructedDataBackupFailureTimeout is the data-structure of this message
type _BACnetConstructedDataBackupFailureTimeout struct {
	BACnetConstructedDataContract
	BackupFailureTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataBackupFailureTimeout = (*_BACnetConstructedDataBackupFailureTimeout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBackupFailureTimeout)(nil)

// NewBACnetConstructedDataBackupFailureTimeout factory function for _BACnetConstructedDataBackupFailureTimeout
func NewBACnetConstructedDataBackupFailureTimeout(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, backupFailureTimeout BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBackupFailureTimeout {
	if backupFailureTimeout == nil {
		panic("backupFailureTimeout of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataBackupFailureTimeout must not be nil")
	}
	_result := &_BACnetConstructedDataBackupFailureTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BackupFailureTimeout:          backupFailureTimeout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBackupFailureTimeoutBuilder is a builder for BACnetConstructedDataBackupFailureTimeout
type BACnetConstructedDataBackupFailureTimeoutBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(backupFailureTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupFailureTimeoutBuilder
	// WithBackupFailureTimeout adds BackupFailureTimeout (property field)
	WithBackupFailureTimeout(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupFailureTimeoutBuilder
	// WithBackupFailureTimeoutBuilder adds BackupFailureTimeout (property field) which is build by the builder
	WithBackupFailureTimeoutBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBackupFailureTimeoutBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBackupFailureTimeout or returns an error if something is wrong
	Build() (BACnetConstructedDataBackupFailureTimeout, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBackupFailureTimeout
}

// NewBACnetConstructedDataBackupFailureTimeoutBuilder() creates a BACnetConstructedDataBackupFailureTimeoutBuilder
func NewBACnetConstructedDataBackupFailureTimeoutBuilder() BACnetConstructedDataBackupFailureTimeoutBuilder {
	return &_BACnetConstructedDataBackupFailureTimeoutBuilder{_BACnetConstructedDataBackupFailureTimeout: new(_BACnetConstructedDataBackupFailureTimeout)}
}

type _BACnetConstructedDataBackupFailureTimeoutBuilder struct {
	*_BACnetConstructedDataBackupFailureTimeout

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBackupFailureTimeoutBuilder) = (*_BACnetConstructedDataBackupFailureTimeoutBuilder)(nil)

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBackupFailureTimeout
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) WithMandatoryFields(backupFailureTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupFailureTimeoutBuilder {
	return b.WithBackupFailureTimeout(backupFailureTimeout)
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) WithBackupFailureTimeout(backupFailureTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBackupFailureTimeoutBuilder {
	b.BackupFailureTimeout = backupFailureTimeout
	return b
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) WithBackupFailureTimeoutBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBackupFailureTimeoutBuilder {
	builder := builderSupplier(b.BackupFailureTimeout.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.BackupFailureTimeout, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) Build() (BACnetConstructedDataBackupFailureTimeout, error) {
	if b.BackupFailureTimeout == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'backupFailureTimeout' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBackupFailureTimeout.deepCopy(), nil
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) MustBuild() BACnetConstructedDataBackupFailureTimeout {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBackupFailureTimeoutBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBackupFailureTimeoutBuilder().(*_BACnetConstructedDataBackupFailureTimeoutBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBackupFailureTimeoutBuilder creates a BACnetConstructedDataBackupFailureTimeoutBuilder
func (b *_BACnetConstructedDataBackupFailureTimeout) CreateBACnetConstructedDataBackupFailureTimeoutBuilder() BACnetConstructedDataBackupFailureTimeoutBuilder {
	if b == nil {
		return NewBACnetConstructedDataBackupFailureTimeoutBuilder()
	}
	return &_BACnetConstructedDataBackupFailureTimeoutBuilder{_BACnetConstructedDataBackupFailureTimeout: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_FAILURE_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetBackupFailureTimeout() BACnetApplicationTagUnsignedInteger {
	return m.BackupFailureTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetBackupFailureTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBackupFailureTimeout(structType any) BACnetConstructedDataBackupFailureTimeout {
	if casted, ok := structType.(BACnetConstructedDataBackupFailureTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupFailureTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetTypeName() string {
	return "BACnetConstructedDataBackupFailureTimeout"
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (backupFailureTimeout)
	lengthInBits += m.BackupFailureTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBackupFailureTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBackupFailureTimeout BACnetConstructedDataBackupFailureTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupFailureTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBackupFailureTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	backupFailureTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "backupFailureTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'backupFailureTimeout' field"))
	}
	m.BackupFailureTimeout = backupFailureTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), backupFailureTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupFailureTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBackupFailureTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupFailureTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBackupFailureTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "backupFailureTimeout", m.GetBackupFailureTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'backupFailureTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupFailureTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBackupFailureTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBackupFailureTimeout) IsBACnetConstructedDataBackupFailureTimeout() {}

func (m *_BACnetConstructedDataBackupFailureTimeout) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBackupFailureTimeout) deepCopy() *_BACnetConstructedDataBackupFailureTimeout {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBackupFailureTimeoutCopy := &_BACnetConstructedDataBackupFailureTimeout{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.BackupFailureTimeout),
	}
	_BACnetConstructedDataBackupFailureTimeoutCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBackupFailureTimeoutCopy
}

func (m *_BACnetConstructedDataBackupFailureTimeout) String() string {
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
