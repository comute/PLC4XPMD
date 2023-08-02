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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpenFileMode is an enum
type OpenFileMode uint32

type IOpenFileMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpenFileMode_openFileModeRead          OpenFileMode = 1
	OpenFileMode_openFileModeWrite         OpenFileMode = 2
	OpenFileMode_openFileModeEraseExisting OpenFileMode = 4
	OpenFileMode_openFileModeAppend        OpenFileMode = 8
)

var OpenFileModeValues []OpenFileMode

func init() {
	_ = errors.New
	OpenFileModeValues = []OpenFileMode{
		OpenFileMode_openFileModeRead,
		OpenFileMode_openFileModeWrite,
		OpenFileMode_openFileModeEraseExisting,
		OpenFileMode_openFileModeAppend,
	}
}

func OpenFileModeByValue(value uint32) (enum OpenFileMode, ok bool) {
	switch value {
	case 1:
		return OpenFileMode_openFileModeRead, true
	case 2:
		return OpenFileMode_openFileModeWrite, true
	case 4:
		return OpenFileMode_openFileModeEraseExisting, true
	case 8:
		return OpenFileMode_openFileModeAppend, true
	}
	return 0, false
}

func OpenFileModeByName(value string) (enum OpenFileMode, ok bool) {
	switch value {
	case "openFileModeRead":
		return OpenFileMode_openFileModeRead, true
	case "openFileModeWrite":
		return OpenFileMode_openFileModeWrite, true
	case "openFileModeEraseExisting":
		return OpenFileMode_openFileModeEraseExisting, true
	case "openFileModeAppend":
		return OpenFileMode_openFileModeAppend, true
	}
	return 0, false
}

func OpenFileModeKnows(value uint32) bool {
	for _, typeValue := range OpenFileModeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpenFileMode(structType any) OpenFileMode {
	castFunc := func(typ any) OpenFileMode {
		if sOpenFileMode, ok := typ.(OpenFileMode); ok {
			return sOpenFileMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpenFileMode) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpenFileMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpenFileModeParse(ctx context.Context, theBytes []byte) (OpenFileMode, error) {
	return OpenFileModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpenFileModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpenFileMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("OpenFileMode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpenFileMode")
	}
	if enum, ok := OpenFileModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpenFileMode")
		return OpenFileMode(val), nil
	} else {
		return enum, nil
	}
}

func (e OpenFileMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpenFileMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("OpenFileMode", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpenFileMode) PLC4XEnumName() string {
	switch e {
	case OpenFileMode_openFileModeRead:
		return "openFileModeRead"
	case OpenFileMode_openFileModeWrite:
		return "openFileModeWrite"
	case OpenFileMode_openFileModeEraseExisting:
		return "openFileModeEraseExisting"
	case OpenFileMode_openFileModeAppend:
		return "openFileModeAppend"
	}
	return ""
}

func (e OpenFileMode) String() string {
	return e.PLC4XEnumName()
}
