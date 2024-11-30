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

// ReplyOrConfirmation is the corresponding interface of ReplyOrConfirmation
type ReplyOrConfirmation interface {
	ReplyOrConfirmationContract
	ReplyOrConfirmationRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsReplyOrConfirmation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReplyOrConfirmation()
	// CreateBuilder creates a ReplyOrConfirmationBuilder
	CreateReplyOrConfirmationBuilder() ReplyOrConfirmationBuilder
}

// ReplyOrConfirmationContract provides a set of functions which can be overwritten by a sub struct
type ReplyOrConfirmationContract interface {
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() byte
	// GetIsAlpha returns IsAlpha (virtual field)
	GetIsAlpha() bool
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// IsReplyOrConfirmation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReplyOrConfirmation()
	// CreateBuilder creates a ReplyOrConfirmationBuilder
	CreateReplyOrConfirmationBuilder() ReplyOrConfirmationBuilder
}

// ReplyOrConfirmationRequirements provides a set of functions which need to be implemented by a sub struct
type ReplyOrConfirmationRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsAlpha returns IsAlpha (discriminator field)
	GetIsAlpha() bool
	// GetPeekedByte returns PeekedByte (discriminator field)
	GetPeekedByte() byte
}

// _ReplyOrConfirmation is the data-structure of this message
type _ReplyOrConfirmation struct {
	_SubType interface {
		ReplyOrConfirmationContract
		ReplyOrConfirmationRequirements
	}
	PeekedByte byte

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

var _ ReplyOrConfirmationContract = (*_ReplyOrConfirmation)(nil)

// NewReplyOrConfirmation factory function for _ReplyOrConfirmation
func NewReplyOrConfirmation(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmation {
	return &_ReplyOrConfirmation{PeekedByte: peekedByte, CBusOptions: cBusOptions, RequestContext: requestContext}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReplyOrConfirmationBuilder is a builder for ReplyOrConfirmation
type ReplyOrConfirmationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedByte byte) ReplyOrConfirmationBuilder
	// WithPeekedByte adds PeekedByte (property field)
	WithPeekedByte(byte) ReplyOrConfirmationBuilder
	// AsServerErrorReply converts this build to a subType of ReplyOrConfirmation. It is always possible to return to current builder using Done()
	AsServerErrorReply() interface {
		ServerErrorReplyBuilder
		Done() ReplyOrConfirmationBuilder
	}
	// AsReplyOrConfirmationConfirmation converts this build to a subType of ReplyOrConfirmation. It is always possible to return to current builder using Done()
	AsReplyOrConfirmationConfirmation() interface {
		ReplyOrConfirmationConfirmationBuilder
		Done() ReplyOrConfirmationBuilder
	}
	// AsReplyOrConfirmationReply converts this build to a subType of ReplyOrConfirmation. It is always possible to return to current builder using Done()
	AsReplyOrConfirmationReply() interface {
		ReplyOrConfirmationReplyBuilder
		Done() ReplyOrConfirmationBuilder
	}
	// Build builds the ReplyOrConfirmation or returns an error if something is wrong
	PartialBuild() (ReplyOrConfirmationContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ReplyOrConfirmationContract
	// Build builds the ReplyOrConfirmation or returns an error if something is wrong
	Build() (ReplyOrConfirmation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReplyOrConfirmation
}

// NewReplyOrConfirmationBuilder() creates a ReplyOrConfirmationBuilder
func NewReplyOrConfirmationBuilder() ReplyOrConfirmationBuilder {
	return &_ReplyOrConfirmationBuilder{_ReplyOrConfirmation: new(_ReplyOrConfirmation)}
}

type _ReplyOrConfirmationChildBuilder interface {
	utils.Copyable
	setParent(ReplyOrConfirmationContract)
	buildForReplyOrConfirmation() (ReplyOrConfirmation, error)
}

type _ReplyOrConfirmationBuilder struct {
	*_ReplyOrConfirmation

	childBuilder _ReplyOrConfirmationChildBuilder

	err *utils.MultiError
}

var _ (ReplyOrConfirmationBuilder) = (*_ReplyOrConfirmationBuilder)(nil)

func (b *_ReplyOrConfirmationBuilder) WithMandatoryFields(peekedByte byte) ReplyOrConfirmationBuilder {
	return b.WithPeekedByte(peekedByte)
}

func (b *_ReplyOrConfirmationBuilder) WithPeekedByte(peekedByte byte) ReplyOrConfirmationBuilder {
	b.PeekedByte = peekedByte
	return b
}

func (b *_ReplyOrConfirmationBuilder) PartialBuild() (ReplyOrConfirmationContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReplyOrConfirmation.deepCopy(), nil
}

func (b *_ReplyOrConfirmationBuilder) PartialMustBuild() ReplyOrConfirmationContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReplyOrConfirmationBuilder) AsServerErrorReply() interface {
	ServerErrorReplyBuilder
	Done() ReplyOrConfirmationBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ServerErrorReplyBuilder
		Done() ReplyOrConfirmationBuilder
	}); ok {
		return cb
	}
	cb := NewServerErrorReplyBuilder().(*_ServerErrorReplyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ReplyOrConfirmationBuilder) AsReplyOrConfirmationConfirmation() interface {
	ReplyOrConfirmationConfirmationBuilder
	Done() ReplyOrConfirmationBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ReplyOrConfirmationConfirmationBuilder
		Done() ReplyOrConfirmationBuilder
	}); ok {
		return cb
	}
	cb := NewReplyOrConfirmationConfirmationBuilder().(*_ReplyOrConfirmationConfirmationBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ReplyOrConfirmationBuilder) AsReplyOrConfirmationReply() interface {
	ReplyOrConfirmationReplyBuilder
	Done() ReplyOrConfirmationBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ReplyOrConfirmationReplyBuilder
		Done() ReplyOrConfirmationBuilder
	}); ok {
		return cb
	}
	cb := NewReplyOrConfirmationReplyBuilder().(*_ReplyOrConfirmationReplyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ReplyOrConfirmationBuilder) Build() (ReplyOrConfirmation, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForReplyOrConfirmation()
}

func (b *_ReplyOrConfirmationBuilder) MustBuild() ReplyOrConfirmation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReplyOrConfirmationBuilder) DeepCopy() any {
	_copy := b.CreateReplyOrConfirmationBuilder().(*_ReplyOrConfirmationBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ReplyOrConfirmationChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReplyOrConfirmationBuilder creates a ReplyOrConfirmationBuilder
func (b *_ReplyOrConfirmation) CreateReplyOrConfirmationBuilder() ReplyOrConfirmationBuilder {
	if b == nil {
		return NewReplyOrConfirmationBuilder()
	}
	return &_ReplyOrConfirmationBuilder{_ReplyOrConfirmation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmation) GetPeekedByte() byte {
	return m.PeekedByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_ReplyOrConfirmation) GetIsAlpha() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((bool((m.GetPeekedByte()) >= (0x67)))) && bool((bool((m.GetPeekedByte()) <= (0x7A)))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmation(structType any) ReplyOrConfirmation {
	if casted, ok := structType.(ReplyOrConfirmation); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmation); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmation) GetTypeName() string {
	return "ReplyOrConfirmation"
}

func (m *_ReplyOrConfirmation) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ReplyOrConfirmation) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ReplyOrConfirmationParse[T ReplyOrConfirmation](ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	return ReplyOrConfirmationParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ReplyOrConfirmationParseWithBufferProducer[T ReplyOrConfirmation](cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ReplyOrConfirmationParseWithBuffer[T](ctx, readBuffer, cBusOptions, requestContext)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ReplyOrConfirmationParseWithBuffer[T ReplyOrConfirmation](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	v, err := (&_ReplyOrConfirmation{CBusOptions: cBusOptions, RequestContext: requestContext}).parse(ctx, readBuffer, cBusOptions, requestContext)
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

func (m *_ReplyOrConfirmation) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (__replyOrConfirmation ReplyOrConfirmation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedByte, err := ReadPeekField[byte](ctx, "peekedByte", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedByte' field"))
	}
	m.PeekedByte = peekedByte

	isAlpha, err := ReadVirtualField[bool](ctx, "isAlpha", (*bool)(nil), bool((bool((peekedByte) >= (0x67)))) && bool((bool((peekedByte) <= (0x7A)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAlpha' field"))
	}
	_ = isAlpha

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ReplyOrConfirmation
	switch {
	case isAlpha == bool(false) && peekedByte == 0x21: // ServerErrorReply
		if _child, err = new(_ServerErrorReply).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ServerErrorReply for type-switch of ReplyOrConfirmation")
		}
	case isAlpha == bool(true): // ReplyOrConfirmationConfirmation
		if _child, err = new(_ReplyOrConfirmationConfirmation).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ReplyOrConfirmationConfirmation for type-switch of ReplyOrConfirmation")
		}
	case isAlpha == bool(false): // ReplyOrConfirmationReply
		if _child, err = new(_ReplyOrConfirmationReply).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ReplyOrConfirmationReply for type-switch of ReplyOrConfirmation")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isAlpha=%v, peekedByte=%v]", isAlpha, peekedByte)
	}

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmation")
	}

	return _child, nil
}

func (pm *_ReplyOrConfirmation) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ReplyOrConfirmation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ReplyOrConfirmation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmation")
	}
	// Virtual field
	isAlpha := m.GetIsAlpha()
	_ = isAlpha
	if _isAlphaErr := writeBuffer.WriteVirtual(ctx, "isAlpha", m.GetIsAlpha()); _isAlphaErr != nil {
		return errors.Wrap(_isAlphaErr, "Error serializing 'isAlpha' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ReplyOrConfirmation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ReplyOrConfirmation")
	}
	return nil
}

////
// Arguments Getter

func (m *_ReplyOrConfirmation) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_ReplyOrConfirmation) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_ReplyOrConfirmation) IsReplyOrConfirmation() {}

func (m *_ReplyOrConfirmation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReplyOrConfirmation) deepCopy() *_ReplyOrConfirmation {
	if m == nil {
		return nil
	}
	_ReplyOrConfirmationCopy := &_ReplyOrConfirmation{
		nil, // will be set by child
		m.PeekedByte,
		m.CBusOptions,
		m.RequestContext,
	}
	return _ReplyOrConfirmationCopy
}
