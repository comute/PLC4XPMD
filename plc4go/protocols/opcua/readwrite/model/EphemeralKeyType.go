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

// EphemeralKeyType is the corresponding interface of EphemeralKeyType
type EphemeralKeyType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetPublicKey returns PublicKey (property field)
	GetPublicKey() PascalByteString
	// GetSignature returns Signature (property field)
	GetSignature() PascalByteString
}

// EphemeralKeyTypeExactly can be used when we want exactly this type and not a type which fulfills EphemeralKeyType.
// This is useful for switch cases.
type EphemeralKeyTypeExactly interface {
	EphemeralKeyType
	isEphemeralKeyType() bool
}

// _EphemeralKeyType is the data-structure of this message
type _EphemeralKeyType struct {
	*_ExtensionObjectDefinition
	PublicKey PascalByteString
	Signature PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EphemeralKeyType) GetIdentifier() string {
	return "17550"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EphemeralKeyType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_EphemeralKeyType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EphemeralKeyType) GetPublicKey() PascalByteString {
	return m.PublicKey
}

func (m *_EphemeralKeyType) GetSignature() PascalByteString {
	return m.Signature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEphemeralKeyType factory function for _EphemeralKeyType
func NewEphemeralKeyType(publicKey PascalByteString, signature PascalByteString) *_EphemeralKeyType {
	_result := &_EphemeralKeyType{
		PublicKey:                  publicKey,
		Signature:                  signature,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEphemeralKeyType(structType any) EphemeralKeyType {
	if casted, ok := structType.(EphemeralKeyType); ok {
		return casted
	}
	if casted, ok := structType.(*EphemeralKeyType); ok {
		return *casted
	}
	return nil
}

func (m *_EphemeralKeyType) GetTypeName() string {
	return "EphemeralKeyType"
}

func (m *_EphemeralKeyType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (publicKey)
	lengthInBits += m.PublicKey.GetLengthInBits(ctx)

	// Simple field (signature)
	lengthInBits += m.Signature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EphemeralKeyType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EphemeralKeyTypeParse(ctx context.Context, theBytes []byte, identifier string) (EphemeralKeyType, error) {
	return EphemeralKeyTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func EphemeralKeyTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (EphemeralKeyType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EphemeralKeyType, error) {
		return EphemeralKeyTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func EphemeralKeyTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (EphemeralKeyType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EphemeralKeyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EphemeralKeyType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	publicKey, err := ReadSimpleField[PascalByteString](ctx, "publicKey", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publicKey' field"))
	}

	signature, err := ReadSimpleField[PascalByteString](ctx, "signature", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'signature' field"))
	}

	if closeErr := readBuffer.CloseContext("EphemeralKeyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EphemeralKeyType")
	}

	// Create a partially initialized instance
	_child := &_EphemeralKeyType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		PublicKey:                  publicKey,
		Signature:                  signature,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_EphemeralKeyType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EphemeralKeyType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EphemeralKeyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EphemeralKeyType")
		}

		// Simple Field (publicKey)
		if pushErr := writeBuffer.PushContext("publicKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for publicKey")
		}
		_publicKeyErr := writeBuffer.WriteSerializable(ctx, m.GetPublicKey())
		if popErr := writeBuffer.PopContext("publicKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for publicKey")
		}
		if _publicKeyErr != nil {
			return errors.Wrap(_publicKeyErr, "Error serializing 'publicKey' field")
		}

		// Simple Field (signature)
		if pushErr := writeBuffer.PushContext("signature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for signature")
		}
		_signatureErr := writeBuffer.WriteSerializable(ctx, m.GetSignature())
		if popErr := writeBuffer.PopContext("signature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for signature")
		}
		if _signatureErr != nil {
			return errors.Wrap(_signatureErr, "Error serializing 'signature' field")
		}

		if popErr := writeBuffer.PopContext("EphemeralKeyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EphemeralKeyType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EphemeralKeyType) isEphemeralKeyType() bool {
	return true
}

func (m *_EphemeralKeyType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
