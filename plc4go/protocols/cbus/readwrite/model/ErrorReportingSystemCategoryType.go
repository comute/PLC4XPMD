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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ErrorReportingSystemCategoryType is the corresponding interface of ErrorReportingSystemCategoryType
type ErrorReportingSystemCategoryType interface {
	ErrorReportingSystemCategoryTypeContract
	ErrorReportingSystemCategoryTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsErrorReportingSystemCategoryType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryType()
	// CreateBuilder creates a ErrorReportingSystemCategoryTypeBuilder
	CreateErrorReportingSystemCategoryTypeBuilder() ErrorReportingSystemCategoryTypeBuilder
}

// ErrorReportingSystemCategoryTypeContract provides a set of functions which can be overwritten by a sub struct
type ErrorReportingSystemCategoryTypeContract interface {
	// IsErrorReportingSystemCategoryType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryType()
	// CreateBuilder creates a ErrorReportingSystemCategoryTypeBuilder
	CreateErrorReportingSystemCategoryTypeBuilder() ErrorReportingSystemCategoryTypeBuilder
}

// ErrorReportingSystemCategoryTypeRequirements provides a set of functions which need to be implemented by a sub struct
type ErrorReportingSystemCategoryTypeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetErrorReportingSystemCategoryClass returns ErrorReportingSystemCategoryClass (discriminator field)
	GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass
}

// _ErrorReportingSystemCategoryType is the data-structure of this message
type _ErrorReportingSystemCategoryType struct {
	_SubType interface {
		ErrorReportingSystemCategoryTypeContract
		ErrorReportingSystemCategoryTypeRequirements
	}
}

var _ ErrorReportingSystemCategoryTypeContract = (*_ErrorReportingSystemCategoryType)(nil)

// NewErrorReportingSystemCategoryType factory function for _ErrorReportingSystemCategoryType
func NewErrorReportingSystemCategoryType() *_ErrorReportingSystemCategoryType {
	return &_ErrorReportingSystemCategoryType{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorReportingSystemCategoryTypeBuilder is a builder for ErrorReportingSystemCategoryType
type ErrorReportingSystemCategoryTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ErrorReportingSystemCategoryTypeBuilder
	// AsErrorReportingSystemCategoryTypeInputUnits converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeInputUnits() interface {
		ErrorReportingSystemCategoryTypeInputUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// AsErrorReportingSystemCategoryTypeSupportUnits converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeSupportUnits() interface {
		ErrorReportingSystemCategoryTypeSupportUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// AsErrorReportingSystemCategoryTypeBuildingManagementSystems converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeBuildingManagementSystems() interface {
		ErrorReportingSystemCategoryTypeBuildingManagementSystemsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// AsErrorReportingSystemCategoryTypeOutputUnits converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeOutputUnits() interface {
		ErrorReportingSystemCategoryTypeOutputUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// AsErrorReportingSystemCategoryTypeClimateControllers converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeClimateControllers() interface {
		ErrorReportingSystemCategoryTypeClimateControllersBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// AsErrorReportingSystemCategoryTypeReserved converts this build to a subType of ErrorReportingSystemCategoryType. It is always possible to return to current builder using Done()
	AsErrorReportingSystemCategoryTypeReserved() interface {
		ErrorReportingSystemCategoryTypeReservedBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}
	// Build builds the ErrorReportingSystemCategoryType or returns an error if something is wrong
	PartialBuild() (ErrorReportingSystemCategoryTypeContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ErrorReportingSystemCategoryTypeContract
	// Build builds the ErrorReportingSystemCategoryType or returns an error if something is wrong
	Build() (ErrorReportingSystemCategoryType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorReportingSystemCategoryType
}

// NewErrorReportingSystemCategoryTypeBuilder() creates a ErrorReportingSystemCategoryTypeBuilder
func NewErrorReportingSystemCategoryTypeBuilder() ErrorReportingSystemCategoryTypeBuilder {
	return &_ErrorReportingSystemCategoryTypeBuilder{_ErrorReportingSystemCategoryType: new(_ErrorReportingSystemCategoryType)}
}

type _ErrorReportingSystemCategoryTypeChildBuilder interface {
	utils.Copyable
	setParent(ErrorReportingSystemCategoryTypeContract)
	buildForErrorReportingSystemCategoryType() (ErrorReportingSystemCategoryType, error)
}

type _ErrorReportingSystemCategoryTypeBuilder struct {
	*_ErrorReportingSystemCategoryType

	childBuilder _ErrorReportingSystemCategoryTypeChildBuilder

	err *utils.MultiError
}

var _ (ErrorReportingSystemCategoryTypeBuilder) = (*_ErrorReportingSystemCategoryTypeBuilder)(nil)

func (b *_ErrorReportingSystemCategoryTypeBuilder) WithMandatoryFields() ErrorReportingSystemCategoryTypeBuilder {
	return b
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) PartialBuild() (ErrorReportingSystemCategoryTypeContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorReportingSystemCategoryType.deepCopy(), nil
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) PartialMustBuild() ErrorReportingSystemCategoryTypeContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeInputUnits() interface {
	ErrorReportingSystemCategoryTypeInputUnitsBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeInputUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeInputUnitsBuilder().(*_ErrorReportingSystemCategoryTypeInputUnitsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeSupportUnits() interface {
	ErrorReportingSystemCategoryTypeSupportUnitsBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeSupportUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeSupportUnitsBuilder().(*_ErrorReportingSystemCategoryTypeSupportUnitsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeBuildingManagementSystems() interface {
	ErrorReportingSystemCategoryTypeBuildingManagementSystemsBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeBuildingManagementSystemsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeBuildingManagementSystemsBuilder().(*_ErrorReportingSystemCategoryTypeBuildingManagementSystemsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeOutputUnits() interface {
	ErrorReportingSystemCategoryTypeOutputUnitsBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeOutputUnitsBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeOutputUnitsBuilder().(*_ErrorReportingSystemCategoryTypeOutputUnitsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeClimateControllers() interface {
	ErrorReportingSystemCategoryTypeClimateControllersBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeClimateControllersBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeClimateControllersBuilder().(*_ErrorReportingSystemCategoryTypeClimateControllersBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) AsErrorReportingSystemCategoryTypeReserved() interface {
	ErrorReportingSystemCategoryTypeReservedBuilder
	Done() ErrorReportingSystemCategoryTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ErrorReportingSystemCategoryTypeReservedBuilder
		Done() ErrorReportingSystemCategoryTypeBuilder
	}); ok {
		return cb
	}
	cb := NewErrorReportingSystemCategoryTypeReservedBuilder().(*_ErrorReportingSystemCategoryTypeReservedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) Build() (ErrorReportingSystemCategoryType, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForErrorReportingSystemCategoryType()
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) MustBuild() ErrorReportingSystemCategoryType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingSystemCategoryTypeBuilder) DeepCopy() any {
	_copy := b.CreateErrorReportingSystemCategoryTypeBuilder().(*_ErrorReportingSystemCategoryTypeBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ErrorReportingSystemCategoryTypeChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorReportingSystemCategoryTypeBuilder creates a ErrorReportingSystemCategoryTypeBuilder
func (b *_ErrorReportingSystemCategoryType) CreateErrorReportingSystemCategoryTypeBuilder() ErrorReportingSystemCategoryTypeBuilder {
	if b == nil {
		return NewErrorReportingSystemCategoryTypeBuilder()
	}
	return &_ErrorReportingSystemCategoryTypeBuilder{_ErrorReportingSystemCategoryType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryType(structType any) ErrorReportingSystemCategoryType {
	if casted, ok := structType.(ErrorReportingSystemCategoryType); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryType); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryType) GetTypeName() string {
	return "ErrorReportingSystemCategoryType"
}

func (m *_ErrorReportingSystemCategoryType) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeParse[T ErrorReportingSystemCategoryType](ctx context.Context, theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (T, error) {
	return ErrorReportingSystemCategoryTypeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeParseWithBufferProducer[T ErrorReportingSystemCategoryType](errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ErrorReportingSystemCategoryTypeParseWithBuffer[T](ctx, readBuffer, errorReportingSystemCategoryClass)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ErrorReportingSystemCategoryTypeParseWithBuffer[T ErrorReportingSystemCategoryType](ctx context.Context, readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (T, error) {
	v, err := (&_ErrorReportingSystemCategoryType{}).parse(ctx, readBuffer, errorReportingSystemCategoryClass)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ErrorReportingSystemCategoryType) parse(ctx context.Context, readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryType ErrorReportingSystemCategoryType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ErrorReportingSystemCategoryType
	switch {
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_INPUT_UNITS: // ErrorReportingSystemCategoryTypeInputUnits
		if _child, err = new(_ErrorReportingSystemCategoryTypeInputUnits).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeInputUnits for type-switch of ErrorReportingSystemCategoryType")
		}
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_SUPPORT_UNITS: // ErrorReportingSystemCategoryTypeSupportUnits
		if _child, err = new(_ErrorReportingSystemCategoryTypeSupportUnits).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeSupportUnits for type-switch of ErrorReportingSystemCategoryType")
		}
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS: // ErrorReportingSystemCategoryTypeBuildingManagementSystems
		if _child, err = new(_ErrorReportingSystemCategoryTypeBuildingManagementSystems).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeBuildingManagementSystems for type-switch of ErrorReportingSystemCategoryType")
		}
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_OUTPUT_UNITS: // ErrorReportingSystemCategoryTypeOutputUnits
		if _child, err = new(_ErrorReportingSystemCategoryTypeOutputUnits).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeOutputUnits for type-switch of ErrorReportingSystemCategoryType")
		}
	case errorReportingSystemCategoryClass == ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS: // ErrorReportingSystemCategoryTypeClimateControllers
		if _child, err = new(_ErrorReportingSystemCategoryTypeClimateControllers).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeClimateControllers for type-switch of ErrorReportingSystemCategoryType")
		}
	case 0 == 0: // ErrorReportingSystemCategoryTypeReserved
		if _child, err = new(_ErrorReportingSystemCategoryTypeReserved).parse(ctx, readBuffer, m, errorReportingSystemCategoryClass); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ErrorReportingSystemCategoryTypeReserved for type-switch of ErrorReportingSystemCategoryType")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [errorReportingSystemCategoryClass=%v]", errorReportingSystemCategoryClass)
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryType")
	}

	return _child, nil
}

func (pm *_ErrorReportingSystemCategoryType) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ErrorReportingSystemCategoryType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryType")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryType")
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryType) IsErrorReportingSystemCategoryType() {}

func (m *_ErrorReportingSystemCategoryType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorReportingSystemCategoryType) deepCopy() *_ErrorReportingSystemCategoryType {
	if m == nil {
		return nil
	}
	_ErrorReportingSystemCategoryTypeCopy := &_ErrorReportingSystemCategoryType{
		nil, // will be set by child
	}
	return _ErrorReportingSystemCategoryTypeCopy
}
