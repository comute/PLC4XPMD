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

// ReadRequest is the corresponding interface of ReadRequest
type ReadRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetMaxAge returns MaxAge (property field)
	GetMaxAge() float64
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetNoOfNodesToRead returns NoOfNodesToRead (property field)
	GetNoOfNodesToRead() int32
	// GetNodesToRead returns NodesToRead (property field)
	GetNodesToRead() []ExtensionObjectDefinition
}

// ReadRequestExactly can be used when we want exactly this type and not a type which fulfills ReadRequest.
// This is useful for switch cases.
type ReadRequestExactly interface {
	ReadRequest
	isReadRequest() bool
}

// _ReadRequest is the data-structure of this message
type _ReadRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader      ExtensionObjectDefinition
	MaxAge             float64
	TimestampsToReturn TimestampsToReturn
	NoOfNodesToRead    int32
	NodesToRead        []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadRequest) GetIdentifier() string {
	return "631"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ReadRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ReadRequest) GetMaxAge() float64 {
	return m.MaxAge
}

func (m *_ReadRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_ReadRequest) GetNoOfNodesToRead() int32 {
	return m.NoOfNodesToRead
}

func (m *_ReadRequest) GetNodesToRead() []ExtensionObjectDefinition {
	return m.NodesToRead
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReadRequest factory function for _ReadRequest
func NewReadRequest(requestHeader ExtensionObjectDefinition, maxAge float64, timestampsToReturn TimestampsToReturn, noOfNodesToRead int32, nodesToRead []ExtensionObjectDefinition) *_ReadRequest {
	_result := &_ReadRequest{
		RequestHeader:              requestHeader,
		MaxAge:                     maxAge,
		TimestampsToReturn:         timestampsToReturn,
		NoOfNodesToRead:            noOfNodesToRead,
		NodesToRead:                nodesToRead,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReadRequest(structType any) ReadRequest {
	if casted, ok := structType.(ReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ReadRequest) GetTypeName() string {
	return "ReadRequest"
}

func (m *_ReadRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (maxAge)
	lengthInBits += 64

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Simple field (noOfNodesToRead)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRead) > 0 {
		for _curItem, element := range m.NodesToRead {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRead), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ReadRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReadRequestParse(ctx context.Context, theBytes []byte, identifier string) (ReadRequest, error) {
	return ReadRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ReadRequestParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ReadRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ReadRequest, error) {
		return ReadRequestParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ReadRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ReadRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}

	maxAge, err := ReadSimpleField(ctx, "maxAge", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxAge' field"))
	}

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}

	noOfNodesToRead, err := ReadSimpleField(ctx, "noOfNodesToRead", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToRead' field"))
	}

	nodesToRead, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToRead", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("628")), readBuffer), uint64(noOfNodesToRead))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToRead' field"))
	}

	if closeErr := readBuffer.CloseContext("ReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadRequest")
	}

	// Create a partially initialized instance
	_child := &_ReadRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		MaxAge:                     maxAge,
		TimestampsToReturn:         timestampsToReturn,
		NoOfNodesToRead:            noOfNodesToRead,
		NodesToRead:                nodesToRead,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ReadRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (maxAge)
		maxAge := float64(m.GetMaxAge())
		_maxAgeErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("maxAge", 64, (maxAge))
		if _maxAgeErr != nil {
			return errors.Wrap(_maxAgeErr, "Error serializing 'maxAge' field")
		}

		// Simple Field (timestampsToReturn)
		if pushErr := writeBuffer.PushContext("timestampsToReturn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestampsToReturn")
		}
		_timestampsToReturnErr := writeBuffer.WriteSerializable(ctx, m.GetTimestampsToReturn())
		if popErr := writeBuffer.PopContext("timestampsToReturn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestampsToReturn")
		}
		if _timestampsToReturnErr != nil {
			return errors.Wrap(_timestampsToReturnErr, "Error serializing 'timestampsToReturn' field")
		}

		// Simple Field (noOfNodesToRead)
		noOfNodesToRead := int32(m.GetNoOfNodesToRead())
		_noOfNodesToReadErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfNodesToRead", 32, int32((noOfNodesToRead)))
		if _noOfNodesToReadErr != nil {
			return errors.Wrap(_noOfNodesToReadErr, "Error serializing 'noOfNodesToRead' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToRead", m.GetNodesToRead(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToRead' field")
		}

		if popErr := writeBuffer.PopContext("ReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadRequest) isReadRequest() bool {
	return true
}

func (m *_ReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
