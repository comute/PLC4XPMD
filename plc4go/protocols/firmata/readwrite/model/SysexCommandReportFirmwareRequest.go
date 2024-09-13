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

// SysexCommandReportFirmwareRequest is the corresponding interface of SysexCommandReportFirmwareRequest
type SysexCommandReportFirmwareRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// IsSysexCommandReportFirmwareRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandReportFirmwareRequest()
}

// _SysexCommandReportFirmwareRequest is the data-structure of this message
type _SysexCommandReportFirmwareRequest struct {
	SysexCommandContract
}

var _ SysexCommandReportFirmwareRequest = (*_SysexCommandReportFirmwareRequest)(nil)
var _ SysexCommandRequirements = (*_SysexCommandReportFirmwareRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandReportFirmwareRequest) GetCommandType() uint8 {
	return 0x79
}

func (m *_SysexCommandReportFirmwareRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandReportFirmwareRequest) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// NewSysexCommandReportFirmwareRequest factory function for _SysexCommandReportFirmwareRequest
func NewSysexCommandReportFirmwareRequest() *_SysexCommandReportFirmwareRequest {
	_result := &_SysexCommandReportFirmwareRequest{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandReportFirmwareRequest(structType any) SysexCommandReportFirmwareRequest {
	if casted, ok := structType.(SysexCommandReportFirmwareRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandReportFirmwareRequest) GetTypeName() string {
	return "SysexCommandReportFirmwareRequest"
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandReportFirmwareRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandReportFirmwareRequest SysexCommandReportFirmwareRequest, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandReportFirmwareRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandReportFirmwareRequest")
	}

	return m, nil
}

func (m *_SysexCommandReportFirmwareRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandReportFirmwareRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandReportFirmwareRequest")
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandReportFirmwareRequest")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandReportFirmwareRequest) IsSysexCommandReportFirmwareRequest() {}

func (m *_SysexCommandReportFirmwareRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
