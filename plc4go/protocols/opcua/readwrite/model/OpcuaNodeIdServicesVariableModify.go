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

// OpcuaNodeIdServicesVariableModify is an enum
type OpcuaNodeIdServicesVariableModify int32

type IOpcuaNodeIdServicesVariableModify interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_InputArguments     OpcuaNodeIdServicesVariableModify = 23730
	OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_OutputArguments    OpcuaNodeIdServicesVariableModify = 23731
	OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_InputArguments    OpcuaNodeIdServicesVariableModify = 23749
	OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_OutputArguments   OpcuaNodeIdServicesVariableModify = 23750
	OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_InputArguments    OpcuaNodeIdServicesVariableModify = 23771
	OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_OutputArguments   OpcuaNodeIdServicesVariableModify = 23772
	OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_InputArguments  OpcuaNodeIdServicesVariableModify = 23782
	OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_OutputArguments OpcuaNodeIdServicesVariableModify = 23783
	OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_InputArguments  OpcuaNodeIdServicesVariableModify = 23793
	OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_OutputArguments OpcuaNodeIdServicesVariableModify = 23794
	OpcuaNodeIdServicesVariableModify_ModifyUserMethodType_InputArguments           OpcuaNodeIdServicesVariableModify = 24285
)

var OpcuaNodeIdServicesVariableModifyValues []OpcuaNodeIdServicesVariableModify

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableModifyValues = []OpcuaNodeIdServicesVariableModify{
		OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_InputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableModify_ModifyUserMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableModifyByValue(value int32) (enum OpcuaNodeIdServicesVariableModify, ok bool) {
	switch value {
	case 23730:
		return OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_InputArguments, true
	case 23731:
		return OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_OutputArguments, true
	case 23749:
		return OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_InputArguments, true
	case 23750:
		return OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_OutputArguments, true
	case 23771:
		return OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_InputArguments, true
	case 23772:
		return OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_OutputArguments, true
	case 23782:
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_InputArguments, true
	case 23783:
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_OutputArguments, true
	case 23793:
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_InputArguments, true
	case 23794:
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_OutputArguments, true
	case 24285:
		return OpcuaNodeIdServicesVariableModify_ModifyUserMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableModifyByName(value string) (enum OpcuaNodeIdServicesVariableModify, ok bool) {
	switch value {
	case "ModifyConnectionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_InputArguments, true
	case "ModifyConnectionMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_OutputArguments, true
	case "ModifyWriterGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_InputArguments, true
	case "ModifyWriterGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_OutputArguments, true
	case "ModifyReaderGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_InputArguments, true
	case "ModifyReaderGroupMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_OutputArguments, true
	case "ModifyDataSetWriterMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_InputArguments, true
	case "ModifyDataSetWriterMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_OutputArguments, true
	case "ModifyDataSetReaderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_InputArguments, true
	case "ModifyDataSetReaderMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_OutputArguments, true
	case "ModifyUserMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableModify_ModifyUserMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableModifyKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableModifyValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableModify(structType any) OpcuaNodeIdServicesVariableModify {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableModify {
		if sOpcuaNodeIdServicesVariableModify, ok := typ.(OpcuaNodeIdServicesVariableModify); ok {
			return sOpcuaNodeIdServicesVariableModify
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableModify) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableModify) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableModifyParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableModify, error) {
	return OpcuaNodeIdServicesVariableModifyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableModifyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableModify, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableModify", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableModify")
	}
	if enum, ok := OpcuaNodeIdServicesVariableModifyByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableModify")
		return OpcuaNodeIdServicesVariableModify(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableModify) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableModify) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableModify", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableModify) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_InputArguments:
		return "ModifyConnectionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyConnectionMethodType_OutputArguments:
		return "ModifyConnectionMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_InputArguments:
		return "ModifyWriterGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyWriterGroupMethodType_OutputArguments:
		return "ModifyWriterGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_InputArguments:
		return "ModifyReaderGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyReaderGroupMethodType_OutputArguments:
		return "ModifyReaderGroupMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_InputArguments:
		return "ModifyDataSetWriterMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyDataSetWriterMethodType_OutputArguments:
		return "ModifyDataSetWriterMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_InputArguments:
		return "ModifyDataSetReaderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyDataSetReaderMethodType_OutputArguments:
		return "ModifyDataSetReaderMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableModify_ModifyUserMethodType_InputArguments:
		return "ModifyUserMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableModify) String() string {
	return e.PLC4XEnumName()
}
