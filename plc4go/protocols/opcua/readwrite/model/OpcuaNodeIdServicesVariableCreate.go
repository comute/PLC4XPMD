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

// OpcuaNodeIdServicesVariableCreate is an enum
type OpcuaNodeIdServicesVariableCreate int32

type IOpcuaNodeIdServicesVariableCreate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_InputArguments  OpcuaNodeIdServicesVariableCreate = 12742
	OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_OutputArguments OpcuaNodeIdServicesVariableCreate = 12743
	OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_InputArguments       OpcuaNodeIdServicesVariableCreate = 13343
	OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_OutputArguments      OpcuaNodeIdServicesVariableCreate = 13344
	OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_InputArguments            OpcuaNodeIdServicesVariableCreate = 13346
	OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_OutputArguments           OpcuaNodeIdServicesVariableCreate = 13347
	OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_InputArguments      OpcuaNodeIdServicesVariableCreate = 15253
	OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_OutputArguments     OpcuaNodeIdServicesVariableCreate = 17495
)

var OpcuaNodeIdServicesVariableCreateValues []OpcuaNodeIdServicesVariableCreate

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableCreateValues = []OpcuaNodeIdServicesVariableCreate{
		OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_InputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_InputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_InputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_InputArguments,
		OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableCreateByValue(value int32) (enum OpcuaNodeIdServicesVariableCreate, ok bool) {
	switch value {
	case 12742:
		return OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_InputArguments, true
	case 12743:
		return OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_OutputArguments, true
	case 13343:
		return OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_InputArguments, true
	case 13344:
		return OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_OutputArguments, true
	case 13346:
		return OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_InputArguments, true
	case 13347:
		return OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_OutputArguments, true
	case 15253:
		return OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_InputArguments, true
	case 17495:
		return OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCreateByName(value string) (enum OpcuaNodeIdServicesVariableCreate, ok bool) {
	switch value {
	case "CreateSigningRequestMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_InputArguments, true
	case "CreateSigningRequestMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_OutputArguments, true
	case "CreateDirectoryMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_InputArguments, true
	case "CreateDirectoryMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_OutputArguments, true
	case "CreateFileMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_InputArguments, true
	case "CreateFileMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_OutputArguments, true
	case "CreateCredentialMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_InputArguments, true
	case "CreateCredentialMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCreateKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableCreateValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableCreate(structType any) OpcuaNodeIdServicesVariableCreate {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableCreate {
		if sOpcuaNodeIdServicesVariableCreate, ok := typ.(OpcuaNodeIdServicesVariableCreate); ok {
			return sOpcuaNodeIdServicesVariableCreate
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableCreate) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableCreate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableCreateParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableCreate, error) {
	return OpcuaNodeIdServicesVariableCreateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableCreateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableCreate, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableCreate", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableCreate")
	}
	if enum, ok := OpcuaNodeIdServicesVariableCreateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableCreate")
		return OpcuaNodeIdServicesVariableCreate(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableCreate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableCreate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableCreate", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableCreate) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableCreate) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_InputArguments:
		return "CreateSigningRequestMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateSigningRequestMethodType_OutputArguments:
		return "CreateSigningRequestMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_InputArguments:
		return "CreateDirectoryMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateDirectoryMethodType_OutputArguments:
		return "CreateDirectoryMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_InputArguments:
		return "CreateFileMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateFileMethodType_OutputArguments:
		return "CreateFileMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_InputArguments:
		return "CreateCredentialMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableCreate_CreateCredentialMethodType_OutputArguments:
		return "CreateCredentialMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableCreate) String() string {
	return e.PLC4XEnumName()
}
