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

// StatusChangeNotification is the corresponding interface of StatusChangeNotification
type StatusChangeNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatus returns Status (property field)
	GetStatus() StatusCode
	// GetDiagnosticInfo returns DiagnosticInfo (property field)
	GetDiagnosticInfo() DiagnosticInfo
	// IsStatusChangeNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStatusChangeNotification()
	// CreateBuilder creates a StatusChangeNotificationBuilder
	CreateStatusChangeNotificationBuilder() StatusChangeNotificationBuilder
}

// _StatusChangeNotification is the data-structure of this message
type _StatusChangeNotification struct {
	ExtensionObjectDefinitionContract
	Status         StatusCode
	DiagnosticInfo DiagnosticInfo
}

var _ StatusChangeNotification = (*_StatusChangeNotification)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_StatusChangeNotification)(nil)

// NewStatusChangeNotification factory function for _StatusChangeNotification
func NewStatusChangeNotification(status StatusCode, diagnosticInfo DiagnosticInfo) *_StatusChangeNotification {
	if status == nil {
		panic("status of type StatusCode for StatusChangeNotification must not be nil")
	}
	if diagnosticInfo == nil {
		panic("diagnosticInfo of type DiagnosticInfo for StatusChangeNotification must not be nil")
	}
	_result := &_StatusChangeNotification{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Status:                            status,
		DiagnosticInfo:                    diagnosticInfo,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StatusChangeNotificationBuilder is a builder for StatusChangeNotification
type StatusChangeNotificationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(status StatusCode, diagnosticInfo DiagnosticInfo) StatusChangeNotificationBuilder
	// WithStatus adds Status (property field)
	WithStatus(StatusCode) StatusChangeNotificationBuilder
	// WithStatusBuilder adds Status (property field) which is build by the builder
	WithStatusBuilder(func(StatusCodeBuilder) StatusCodeBuilder) StatusChangeNotificationBuilder
	// WithDiagnosticInfo adds DiagnosticInfo (property field)
	WithDiagnosticInfo(DiagnosticInfo) StatusChangeNotificationBuilder
	// WithDiagnosticInfoBuilder adds DiagnosticInfo (property field) which is build by the builder
	WithDiagnosticInfoBuilder(func(DiagnosticInfoBuilder) DiagnosticInfoBuilder) StatusChangeNotificationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the StatusChangeNotification or returns an error if something is wrong
	Build() (StatusChangeNotification, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StatusChangeNotification
}

// NewStatusChangeNotificationBuilder() creates a StatusChangeNotificationBuilder
func NewStatusChangeNotificationBuilder() StatusChangeNotificationBuilder {
	return &_StatusChangeNotificationBuilder{_StatusChangeNotification: new(_StatusChangeNotification)}
}

type _StatusChangeNotificationBuilder struct {
	*_StatusChangeNotification

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (StatusChangeNotificationBuilder) = (*_StatusChangeNotificationBuilder)(nil)

func (b *_StatusChangeNotificationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._StatusChangeNotification
}

func (b *_StatusChangeNotificationBuilder) WithMandatoryFields(status StatusCode, diagnosticInfo DiagnosticInfo) StatusChangeNotificationBuilder {
	return b.WithStatus(status).WithDiagnosticInfo(diagnosticInfo)
}

func (b *_StatusChangeNotificationBuilder) WithStatus(status StatusCode) StatusChangeNotificationBuilder {
	b.Status = status
	return b
}

func (b *_StatusChangeNotificationBuilder) WithStatusBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) StatusChangeNotificationBuilder {
	builder := builderSupplier(b.Status.CreateStatusCodeBuilder())
	var err error
	b.Status, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_StatusChangeNotificationBuilder) WithDiagnosticInfo(diagnosticInfo DiagnosticInfo) StatusChangeNotificationBuilder {
	b.DiagnosticInfo = diagnosticInfo
	return b
}

func (b *_StatusChangeNotificationBuilder) WithDiagnosticInfoBuilder(builderSupplier func(DiagnosticInfoBuilder) DiagnosticInfoBuilder) StatusChangeNotificationBuilder {
	builder := builderSupplier(b.DiagnosticInfo.CreateDiagnosticInfoBuilder())
	var err error
	b.DiagnosticInfo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DiagnosticInfoBuilder failed"))
	}
	return b
}

func (b *_StatusChangeNotificationBuilder) Build() (StatusChangeNotification, error) {
	if b.Status == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'status' not set"))
	}
	if b.DiagnosticInfo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'diagnosticInfo' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StatusChangeNotification.deepCopy(), nil
}

func (b *_StatusChangeNotificationBuilder) MustBuild() StatusChangeNotification {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StatusChangeNotificationBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_StatusChangeNotificationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_StatusChangeNotificationBuilder) DeepCopy() any {
	_copy := b.CreateStatusChangeNotificationBuilder().(*_StatusChangeNotificationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStatusChangeNotificationBuilder creates a StatusChangeNotificationBuilder
func (b *_StatusChangeNotification) CreateStatusChangeNotificationBuilder() StatusChangeNotificationBuilder {
	if b == nil {
		return NewStatusChangeNotificationBuilder()
	}
	return &_StatusChangeNotificationBuilder{_StatusChangeNotification: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StatusChangeNotification) GetExtensionId() int32 {
	return int32(820)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusChangeNotification) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusChangeNotification) GetStatus() StatusCode {
	return m.Status
}

func (m *_StatusChangeNotification) GetDiagnosticInfo() DiagnosticInfo {
	return m.DiagnosticInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStatusChangeNotification(structType any) StatusChangeNotification {
	if casted, ok := structType.(StatusChangeNotification); ok {
		return casted
	}
	if casted, ok := structType.(*StatusChangeNotification); ok {
		return *casted
	}
	return nil
}

func (m *_StatusChangeNotification) GetTypeName() string {
	return "StatusChangeNotification"
}

func (m *_StatusChangeNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (status)
	lengthInBits += m.Status.GetLengthInBits(ctx)

	// Simple field (diagnosticInfo)
	lengthInBits += m.DiagnosticInfo.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StatusChangeNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StatusChangeNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__statusChangeNotification StatusChangeNotification, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusChangeNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusChangeNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	status, err := ReadSimpleField[StatusCode](ctx, "status", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	diagnosticInfo, err := ReadSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfo' field"))
	}
	m.DiagnosticInfo = diagnosticInfo

	if closeErr := readBuffer.CloseContext("StatusChangeNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusChangeNotification")
	}

	return m, nil
}

func (m *_StatusChangeNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusChangeNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusChangeNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusChangeNotification")
		}

		if err := WriteSimpleField[StatusCode](ctx, "status", m.GetStatus(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", m.GetDiagnosticInfo(), WriteComplex[DiagnosticInfo](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfo' field")
		}

		if popErr := writeBuffer.PopContext("StatusChangeNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusChangeNotification")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusChangeNotification) IsStatusChangeNotification() {}

func (m *_StatusChangeNotification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StatusChangeNotification) deepCopy() *_StatusChangeNotification {
	if m == nil {
		return nil
	}
	_StatusChangeNotificationCopy := &_StatusChangeNotification{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[StatusCode](m.Status),
		utils.DeepCopy[DiagnosticInfo](m.DiagnosticInfo),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _StatusChangeNotificationCopy
}

func (m *_StatusChangeNotification) String() string {
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
