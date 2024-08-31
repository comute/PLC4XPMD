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

// HistoryData is the corresponding interface of HistoryData
type HistoryData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfDataValues returns NoOfDataValues (property field)
	GetNoOfDataValues() int32
	// GetDataValues returns DataValues (property field)
	GetDataValues() []DataValue
}

// HistoryDataExactly can be used when we want exactly this type and not a type which fulfills HistoryData.
// This is useful for switch cases.
type HistoryDataExactly interface {
	HistoryData
	isHistoryData() bool
}

// _HistoryData is the data-structure of this message
type _HistoryData struct {
	*_ExtensionObjectDefinition
	NoOfDataValues int32
	DataValues     []DataValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryData) GetIdentifier() string {
	return "658"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryData) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryData) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryData) GetNoOfDataValues() int32 {
	return m.NoOfDataValues
}

func (m *_HistoryData) GetDataValues() []DataValue {
	return m.DataValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryData factory function for _HistoryData
func NewHistoryData(noOfDataValues int32, dataValues []DataValue) *_HistoryData {
	_result := &_HistoryData{
		NoOfDataValues:             noOfDataValues,
		DataValues:                 dataValues,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryData(structType any) HistoryData {
	if casted, ok := structType.(HistoryData); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryData); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryData) GetTypeName() string {
	return "HistoryData"
}

func (m *_HistoryData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfDataValues)
	lengthInBits += 32

	// Array field
	if len(m.DataValues) > 0 {
		for _curItem, element := range m.DataValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryDataParse(ctx context.Context, theBytes []byte, identifier string) (HistoryData, error) {
	return HistoryDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryDataParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (HistoryData, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HistoryData, error) {
		return HistoryDataParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func HistoryDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfDataValues, err := ReadSimpleField(ctx, "noOfDataValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataValues' field"))
	}

	dataValues, err := ReadCountArrayField[DataValue](ctx, "dataValues", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(noOfDataValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataValues' field"))
	}

	if closeErr := readBuffer.CloseContext("HistoryData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryData")
	}

	// Create a partially initialized instance
	_child := &_HistoryData{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfDataValues:             noOfDataValues,
		DataValues:                 dataValues,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryData")
		}

		// Simple Field (noOfDataValues)
		noOfDataValues := int32(m.GetNoOfDataValues())
		_noOfDataValuesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDataValues", 32, int32((noOfDataValues)))
		if _noOfDataValuesErr != nil {
			return errors.Wrap(_noOfDataValuesErr, "Error serializing 'noOfDataValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataValues", m.GetDataValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataValues' field")
		}

		if popErr := writeBuffer.PopContext("HistoryData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryData) isHistoryData() bool {
	return true
}

func (m *_HistoryData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
