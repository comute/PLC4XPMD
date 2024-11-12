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

// RolePermissionType is the corresponding interface of RolePermissionType
type RolePermissionType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRoleId returns RoleId (property field)
	GetRoleId() NodeId
	// GetPermissions returns Permissions (property field)
	GetPermissions() PermissionType
	// IsRolePermissionType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRolePermissionType()
	// CreateBuilder creates a RolePermissionTypeBuilder
	CreateRolePermissionTypeBuilder() RolePermissionTypeBuilder
}

// _RolePermissionType is the data-structure of this message
type _RolePermissionType struct {
	ExtensionObjectDefinitionContract
	RoleId      NodeId
	Permissions PermissionType
}

var _ RolePermissionType = (*_RolePermissionType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RolePermissionType)(nil)

// NewRolePermissionType factory function for _RolePermissionType
func NewRolePermissionType(roleId NodeId, permissions PermissionType) *_RolePermissionType {
	if roleId == nil {
		panic("roleId of type NodeId for RolePermissionType must not be nil")
	}
	_result := &_RolePermissionType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RoleId:                            roleId,
		Permissions:                       permissions,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RolePermissionTypeBuilder is a builder for RolePermissionType
type RolePermissionTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(roleId NodeId, permissions PermissionType) RolePermissionTypeBuilder
	// WithRoleId adds RoleId (property field)
	WithRoleId(NodeId) RolePermissionTypeBuilder
	// WithRoleIdBuilder adds RoleId (property field) which is build by the builder
	WithRoleIdBuilder(func(NodeIdBuilder) NodeIdBuilder) RolePermissionTypeBuilder
	// WithPermissions adds Permissions (property field)
	WithPermissions(PermissionType) RolePermissionTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the RolePermissionType or returns an error if something is wrong
	Build() (RolePermissionType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RolePermissionType
}

// NewRolePermissionTypeBuilder() creates a RolePermissionTypeBuilder
func NewRolePermissionTypeBuilder() RolePermissionTypeBuilder {
	return &_RolePermissionTypeBuilder{_RolePermissionType: new(_RolePermissionType)}
}

type _RolePermissionTypeBuilder struct {
	*_RolePermissionType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RolePermissionTypeBuilder) = (*_RolePermissionTypeBuilder)(nil)

func (b *_RolePermissionTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._RolePermissionType
}

func (b *_RolePermissionTypeBuilder) WithMandatoryFields(roleId NodeId, permissions PermissionType) RolePermissionTypeBuilder {
	return b.WithRoleId(roleId).WithPermissions(permissions)
}

func (b *_RolePermissionTypeBuilder) WithRoleId(roleId NodeId) RolePermissionTypeBuilder {
	b.RoleId = roleId
	return b
}

func (b *_RolePermissionTypeBuilder) WithRoleIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) RolePermissionTypeBuilder {
	builder := builderSupplier(b.RoleId.CreateNodeIdBuilder())
	var err error
	b.RoleId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_RolePermissionTypeBuilder) WithPermissions(permissions PermissionType) RolePermissionTypeBuilder {
	b.Permissions = permissions
	return b
}

func (b *_RolePermissionTypeBuilder) Build() (RolePermissionType, error) {
	if b.RoleId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'roleId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RolePermissionType.deepCopy(), nil
}

func (b *_RolePermissionTypeBuilder) MustBuild() RolePermissionType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RolePermissionTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_RolePermissionTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RolePermissionTypeBuilder) DeepCopy() any {
	_copy := b.CreateRolePermissionTypeBuilder().(*_RolePermissionTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRolePermissionTypeBuilder creates a RolePermissionTypeBuilder
func (b *_RolePermissionType) CreateRolePermissionTypeBuilder() RolePermissionTypeBuilder {
	if b == nil {
		return NewRolePermissionTypeBuilder()
	}
	return &_RolePermissionTypeBuilder{_RolePermissionType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RolePermissionType) GetExtensionId() int32 {
	return int32(98)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RolePermissionType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RolePermissionType) GetRoleId() NodeId {
	return m.RoleId
}

func (m *_RolePermissionType) GetPermissions() PermissionType {
	return m.Permissions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRolePermissionType(structType any) RolePermissionType {
	if casted, ok := structType.(RolePermissionType); ok {
		return casted
	}
	if casted, ok := structType.(*RolePermissionType); ok {
		return *casted
	}
	return nil
}

func (m *_RolePermissionType) GetTypeName() string {
	return "RolePermissionType"
}

func (m *_RolePermissionType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (roleId)
	lengthInBits += m.RoleId.GetLengthInBits(ctx)

	// Simple field (permissions)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RolePermissionType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RolePermissionType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__rolePermissionType RolePermissionType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RolePermissionType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RolePermissionType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	roleId, err := ReadSimpleField[NodeId](ctx, "roleId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'roleId' field"))
	}
	m.RoleId = roleId

	permissions, err := ReadEnumField[PermissionType](ctx, "permissions", "PermissionType", ReadEnum(PermissionTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'permissions' field"))
	}
	m.Permissions = permissions

	if closeErr := readBuffer.CloseContext("RolePermissionType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RolePermissionType")
	}

	return m, nil
}

func (m *_RolePermissionType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RolePermissionType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RolePermissionType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RolePermissionType")
		}

		if err := WriteSimpleField[NodeId](ctx, "roleId", m.GetRoleId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'roleId' field")
		}

		if err := WriteSimpleEnumField[PermissionType](ctx, "permissions", "PermissionType", m.GetPermissions(), WriteEnum[PermissionType, uint32](PermissionType.GetValue, PermissionType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'permissions' field")
		}

		if popErr := writeBuffer.PopContext("RolePermissionType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RolePermissionType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RolePermissionType) IsRolePermissionType() {}

func (m *_RolePermissionType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RolePermissionType) deepCopy() *_RolePermissionType {
	if m == nil {
		return nil
	}
	_RolePermissionTypeCopy := &_RolePermissionType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.RoleId),
		m.Permissions,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RolePermissionTypeCopy
}

func (m *_RolePermissionType) String() string {
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
