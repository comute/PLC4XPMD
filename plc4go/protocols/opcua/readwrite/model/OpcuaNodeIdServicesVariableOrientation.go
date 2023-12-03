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

// OpcuaNodeIdServicesVariableOrientation is an enum
type OpcuaNodeIdServicesVariableOrientation int32

type IOpcuaNodeIdServicesVariableOrientation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableOrientation_OrientationType_AngleUnit OpcuaNodeIdServicesVariableOrientation = 18780
)

var OpcuaNodeIdServicesVariableOrientationValues []OpcuaNodeIdServicesVariableOrientation

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOrientationValues = []OpcuaNodeIdServicesVariableOrientation{
		OpcuaNodeIdServicesVariableOrientation_OrientationType_AngleUnit,
	}
}

func OpcuaNodeIdServicesVariableOrientationByValue(value int32) (enum OpcuaNodeIdServicesVariableOrientation, ok bool) {
	switch value {
	case 18780:
		return OpcuaNodeIdServicesVariableOrientation_OrientationType_AngleUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOrientationByName(value string) (enum OpcuaNodeIdServicesVariableOrientation, ok bool) {
	switch value {
	case "OrientationType_AngleUnit":
		return OpcuaNodeIdServicesVariableOrientation_OrientationType_AngleUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOrientationKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOrientationValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableOrientation(structType any) OpcuaNodeIdServicesVariableOrientation {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOrientation {
		if sOpcuaNodeIdServicesVariableOrientation, ok := typ.(OpcuaNodeIdServicesVariableOrientation); ok {
			return sOpcuaNodeIdServicesVariableOrientation
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOrientation) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOrientation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOrientationParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOrientation, error) {
	return OpcuaNodeIdServicesVariableOrientationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOrientationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOrientation, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOrientation", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOrientation")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOrientationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOrientation")
		return OpcuaNodeIdServicesVariableOrientation(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOrientation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOrientation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOrientation", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOrientation) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOrientation_OrientationType_AngleUnit:
		return "OrientationType_AngleUnit"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOrientation) String() string {
	return e.PLC4XEnumName()
}
