/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type DF1RequestProtectedTypedLogicalRead struct {
	*DF1RequestCommand
	ByteSize         uint8
	FileNumber       uint8
	FileType         uint8
	ElementNumber    uint8
	SubElementNumber uint8
}

// The corresponding interface
type IDF1RequestProtectedTypedLogicalRead interface {
	IDF1RequestCommand
	// GetByteSize returns ByteSize (property field)
	GetByteSize() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint8
	// GetFileType returns FileType (property field)
	GetFileType() uint8
	// GetElementNumber returns ElementNumber (property field)
	GetElementNumber() uint8
	// GetSubElementNumber returns SubElementNumber (property field)
	GetSubElementNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *DF1RequestProtectedTypedLogicalRead) GetFunctionCode() uint8 {
	return 0xA2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *DF1RequestProtectedTypedLogicalRead) InitializeParent(parent *DF1RequestCommand) {}

func (m *DF1RequestProtectedTypedLogicalRead) GetParent() *DF1RequestCommand {
	return m.DF1RequestCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *DF1RequestProtectedTypedLogicalRead) GetByteSize() uint8 {
	return m.ByteSize
}

func (m *DF1RequestProtectedTypedLogicalRead) GetFileNumber() uint8 {
	return m.FileNumber
}

func (m *DF1RequestProtectedTypedLogicalRead) GetFileType() uint8 {
	return m.FileType
}

func (m *DF1RequestProtectedTypedLogicalRead) GetElementNumber() uint8 {
	return m.ElementNumber
}

func (m *DF1RequestProtectedTypedLogicalRead) GetSubElementNumber() uint8 {
	return m.SubElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1RequestProtectedTypedLogicalRead factory function for DF1RequestProtectedTypedLogicalRead
func NewDF1RequestProtectedTypedLogicalRead(byteSize uint8, fileNumber uint8, fileType uint8, elementNumber uint8, subElementNumber uint8) *DF1RequestProtectedTypedLogicalRead {
	_result := &DF1RequestProtectedTypedLogicalRead{
		ByteSize:          byteSize,
		FileNumber:        fileNumber,
		FileType:          fileType,
		ElementNumber:     elementNumber,
		SubElementNumber:  subElementNumber,
		DF1RequestCommand: NewDF1RequestCommand(),
	}
	_result.Child = _result
	return _result
}

func CastDF1RequestProtectedTypedLogicalRead(structType interface{}) *DF1RequestProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1RequestProtectedTypedLogicalRead); ok {
		return &casted
	}
	if casted, ok := structType.(*DF1RequestProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(DF1RequestCommand); ok {
		return CastDF1RequestProtectedTypedLogicalRead(casted.Child)
	}
	if casted, ok := structType.(*DF1RequestCommand); ok {
		return CastDF1RequestProtectedTypedLogicalRead(casted.Child)
	}
	return nil
}

func (m *DF1RequestProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1RequestProtectedTypedLogicalRead"
}

func (m *DF1RequestProtectedTypedLogicalRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1RequestProtectedTypedLogicalRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (byteSize)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 8

	// Simple field (fileType)
	lengthInBits += 8

	// Simple field (elementNumber)
	lengthInBits += 8

	// Simple field (subElementNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *DF1RequestProtectedTypedLogicalRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1RequestProtectedTypedLogicalReadParse(readBuffer utils.ReadBuffer) (*DF1RequestProtectedTypedLogicalRead, error) {
	if pullErr := readBuffer.PullContext("DF1RequestProtectedTypedLogicalRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (byteSize)
	_byteSize, _byteSizeErr := readBuffer.ReadUint8("byteSize", 8)
	if _byteSizeErr != nil {
		return nil, errors.Wrap(_byteSizeErr, "Error parsing 'byteSize' field")
	}
	byteSize := _byteSize

	// Simple Field (fileNumber)
	_fileNumber, _fileNumberErr := readBuffer.ReadUint8("fileNumber", 8)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field")
	}
	fileNumber := _fileNumber

	// Simple Field (fileType)
	_fileType, _fileTypeErr := readBuffer.ReadUint8("fileType", 8)
	if _fileTypeErr != nil {
		return nil, errors.Wrap(_fileTypeErr, "Error parsing 'fileType' field")
	}
	fileType := _fileType

	// Simple Field (elementNumber)
	_elementNumber, _elementNumberErr := readBuffer.ReadUint8("elementNumber", 8)
	if _elementNumberErr != nil {
		return nil, errors.Wrap(_elementNumberErr, "Error parsing 'elementNumber' field")
	}
	elementNumber := _elementNumber

	// Simple Field (subElementNumber)
	_subElementNumber, _subElementNumberErr := readBuffer.ReadUint8("subElementNumber", 8)
	if _subElementNumberErr != nil {
		return nil, errors.Wrap(_subElementNumberErr, "Error parsing 'subElementNumber' field")
	}
	subElementNumber := _subElementNumber

	if closeErr := readBuffer.CloseContext("DF1RequestProtectedTypedLogicalRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1RequestProtectedTypedLogicalRead{
		ByteSize:          byteSize,
		FileNumber:        fileNumber,
		FileType:          fileType,
		ElementNumber:     elementNumber,
		SubElementNumber:  subElementNumber,
		DF1RequestCommand: &DF1RequestCommand{},
	}
	_child.DF1RequestCommand.Child = _child
	return _child, nil
}

func (m *DF1RequestProtectedTypedLogicalRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1RequestProtectedTypedLogicalRead"); pushErr != nil {
			return pushErr
		}

		// Simple Field (byteSize)
		byteSize := uint8(m.ByteSize)
		_byteSizeErr := writeBuffer.WriteUint8("byteSize", 8, (byteSize))
		if _byteSizeErr != nil {
			return errors.Wrap(_byteSizeErr, "Error serializing 'byteSize' field")
		}

		// Simple Field (fileNumber)
		fileNumber := uint8(m.FileNumber)
		_fileNumberErr := writeBuffer.WriteUint8("fileNumber", 8, (fileNumber))
		if _fileNumberErr != nil {
			return errors.Wrap(_fileNumberErr, "Error serializing 'fileNumber' field")
		}

		// Simple Field (fileType)
		fileType := uint8(m.FileType)
		_fileTypeErr := writeBuffer.WriteUint8("fileType", 8, (fileType))
		if _fileTypeErr != nil {
			return errors.Wrap(_fileTypeErr, "Error serializing 'fileType' field")
		}

		// Simple Field (elementNumber)
		elementNumber := uint8(m.ElementNumber)
		_elementNumberErr := writeBuffer.WriteUint8("elementNumber", 8, (elementNumber))
		if _elementNumberErr != nil {
			return errors.Wrap(_elementNumberErr, "Error serializing 'elementNumber' field")
		}

		// Simple Field (subElementNumber)
		subElementNumber := uint8(m.SubElementNumber)
		_subElementNumberErr := writeBuffer.WriteUint8("subElementNumber", 8, (subElementNumber))
		if _subElementNumberErr != nil {
			return errors.Wrap(_subElementNumberErr, "Error serializing 'subElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("DF1RequestProtectedTypedLogicalRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1RequestProtectedTypedLogicalRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
