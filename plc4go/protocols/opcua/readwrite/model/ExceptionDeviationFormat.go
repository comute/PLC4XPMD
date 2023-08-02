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

// ExceptionDeviationFormat is an enum
type ExceptionDeviationFormat uint32

type IExceptionDeviationFormat interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ExceptionDeviationFormat_exceptionDeviationFormatAbsoluteValue    ExceptionDeviationFormat = 0
	ExceptionDeviationFormat_exceptionDeviationFormatPercentOfValue   ExceptionDeviationFormat = 1
	ExceptionDeviationFormat_exceptionDeviationFormatPercentOfRange   ExceptionDeviationFormat = 2
	ExceptionDeviationFormat_exceptionDeviationFormatPercentOfEURange ExceptionDeviationFormat = 3
	ExceptionDeviationFormat_exceptionDeviationFormatUnknown          ExceptionDeviationFormat = 4
)

var ExceptionDeviationFormatValues []ExceptionDeviationFormat

func init() {
	_ = errors.New
	ExceptionDeviationFormatValues = []ExceptionDeviationFormat{
		ExceptionDeviationFormat_exceptionDeviationFormatAbsoluteValue,
		ExceptionDeviationFormat_exceptionDeviationFormatPercentOfValue,
		ExceptionDeviationFormat_exceptionDeviationFormatPercentOfRange,
		ExceptionDeviationFormat_exceptionDeviationFormatPercentOfEURange,
		ExceptionDeviationFormat_exceptionDeviationFormatUnknown,
	}
}

func ExceptionDeviationFormatByValue(value uint32) (enum ExceptionDeviationFormat, ok bool) {
	switch value {
	case 0:
		return ExceptionDeviationFormat_exceptionDeviationFormatAbsoluteValue, true
	case 1:
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfValue, true
	case 2:
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfRange, true
	case 3:
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfEURange, true
	case 4:
		return ExceptionDeviationFormat_exceptionDeviationFormatUnknown, true
	}
	return 0, false
}

func ExceptionDeviationFormatByName(value string) (enum ExceptionDeviationFormat, ok bool) {
	switch value {
	case "exceptionDeviationFormatAbsoluteValue":
		return ExceptionDeviationFormat_exceptionDeviationFormatAbsoluteValue, true
	case "exceptionDeviationFormatPercentOfValue":
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfValue, true
	case "exceptionDeviationFormatPercentOfRange":
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfRange, true
	case "exceptionDeviationFormatPercentOfEURange":
		return ExceptionDeviationFormat_exceptionDeviationFormatPercentOfEURange, true
	case "exceptionDeviationFormatUnknown":
		return ExceptionDeviationFormat_exceptionDeviationFormatUnknown, true
	}
	return 0, false
}

func ExceptionDeviationFormatKnows(value uint32) bool {
	for _, typeValue := range ExceptionDeviationFormatValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastExceptionDeviationFormat(structType any) ExceptionDeviationFormat {
	castFunc := func(typ any) ExceptionDeviationFormat {
		if sExceptionDeviationFormat, ok := typ.(ExceptionDeviationFormat); ok {
			return sExceptionDeviationFormat
		}
		return 0
	}
	return castFunc(structType)
}

func (m ExceptionDeviationFormat) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m ExceptionDeviationFormat) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ExceptionDeviationFormatParse(ctx context.Context, theBytes []byte) (ExceptionDeviationFormat, error) {
	return ExceptionDeviationFormatParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ExceptionDeviationFormatParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ExceptionDeviationFormat, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("ExceptionDeviationFormat", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ExceptionDeviationFormat")
	}
	if enum, ok := ExceptionDeviationFormatByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ExceptionDeviationFormat")
		return ExceptionDeviationFormat(val), nil
	} else {
		return enum, nil
	}
}

func (e ExceptionDeviationFormat) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ExceptionDeviationFormat) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("ExceptionDeviationFormat", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ExceptionDeviationFormat) PLC4XEnumName() string {
	switch e {
	case ExceptionDeviationFormat_exceptionDeviationFormatAbsoluteValue:
		return "exceptionDeviationFormatAbsoluteValue"
	case ExceptionDeviationFormat_exceptionDeviationFormatPercentOfValue:
		return "exceptionDeviationFormatPercentOfValue"
	case ExceptionDeviationFormat_exceptionDeviationFormatPercentOfRange:
		return "exceptionDeviationFormatPercentOfRange"
	case ExceptionDeviationFormat_exceptionDeviationFormatPercentOfEURange:
		return "exceptionDeviationFormatPercentOfEURange"
	case ExceptionDeviationFormat_exceptionDeviationFormatUnknown:
		return "exceptionDeviationFormatUnknown"
	}
	return ""
}

func (e ExceptionDeviationFormat) String() string {
	return e.PLC4XEnumName()
}
