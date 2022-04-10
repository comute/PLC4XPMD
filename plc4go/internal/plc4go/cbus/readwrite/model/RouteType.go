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

type RouteType byte

type IRouteType interface {
	AdditionalBridges() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	RouteType_NoBridgeAtAll         RouteType = 0x00
	RouteType_NoAdditionalBridge    RouteType = 0x09
	RouteType_OneAdditionalBridge   RouteType = 0x12
	RouteType_TwoAdditionalBridge   RouteType = 0x1B
	RouteType_ThreeAdditionalBridge RouteType = 0x24
	RouteType_FourAdditionalBridge  RouteType = 0x2D
	RouteType_FiveAdditionalBridge  RouteType = 0x36
)

var RouteTypeValues []RouteType

func init() {
	_ = errors.New
	RouteTypeValues = []RouteType{
		RouteType_NoBridgeAtAll,
		RouteType_NoAdditionalBridge,
		RouteType_OneAdditionalBridge,
		RouteType_TwoAdditionalBridge,
		RouteType_ThreeAdditionalBridge,
		RouteType_FourAdditionalBridge,
		RouteType_FiveAdditionalBridge,
	}
}

func (e RouteType) AdditionalBridges() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0
		}
	case 0x09:
		{ /* '0x09' */
			return 0
		}
	case 0x12:
		{ /* '0x12' */
			return 1
		}
	case 0x1B:
		{ /* '0x1B' */
			return 2
		}
	case 0x24:
		{ /* '0x24' */
			return 3
		}
	case 0x2D:
		{ /* '0x2D' */
			return 4
		}
	case 0x36:
		{ /* '0x36' */
			return 5
		}
	default:
		{
			return 0
		}
	}
}

func RouteTypeFirstEnumForFieldAdditionalBridges(value uint8) (RouteType, error) {
	for _, sizeValue := range RouteTypeValues {
		if sizeValue.AdditionalBridges() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing AdditionalBridges not found", value)
}
func RouteTypeByValue(value byte) RouteType {
	switch value {
	case 0x00:
		return RouteType_NoBridgeAtAll
	case 0x09:
		return RouteType_NoAdditionalBridge
	case 0x12:
		return RouteType_OneAdditionalBridge
	case 0x1B:
		return RouteType_TwoAdditionalBridge
	case 0x24:
		return RouteType_ThreeAdditionalBridge
	case 0x2D:
		return RouteType_FourAdditionalBridge
	case 0x36:
		return RouteType_FiveAdditionalBridge
	}
	return 0
}

func RouteTypeByName(value string) RouteType {
	switch value {
	case "NoBridgeAtAll":
		return RouteType_NoBridgeAtAll
	case "NoAdditionalBridge":
		return RouteType_NoAdditionalBridge
	case "OneAdditionalBridge":
		return RouteType_OneAdditionalBridge
	case "TwoAdditionalBridge":
		return RouteType_TwoAdditionalBridge
	case "ThreeAdditionalBridge":
		return RouteType_ThreeAdditionalBridge
	case "FourAdditionalBridge":
		return RouteType_FourAdditionalBridge
	case "FiveAdditionalBridge":
		return RouteType_FiveAdditionalBridge
	}
	return 0
}

func RouteTypeKnows(value byte) bool {
	for _, typeValue := range RouteTypeValues {
		if byte(typeValue) == value {
			return true
		}
	}
	return false
}

func CastRouteType(structType interface{}) RouteType {
	castFunc := func(typ interface{}) RouteType {
		if sRouteType, ok := typ.(RouteType); ok {
			return sRouteType
		}
		return 0
	}
	return castFunc(structType)
}

func (m RouteType) GetLengthInBits() uint16 {
	return 8
}

func (m RouteType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RouteTypeParse(readBuffer utils.ReadBuffer) (RouteType, error) {
	val, err := readBuffer.ReadByte("RouteType")
	if err != nil {
		return 0, nil
	}
	return RouteTypeByValue(val), nil
}

func (e RouteType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteByte("RouteType", byte(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e RouteType) name() string {
	switch e {
	case RouteType_NoBridgeAtAll:
		return "NoBridgeAtAll"
	case RouteType_NoAdditionalBridge:
		return "NoAdditionalBridge"
	case RouteType_OneAdditionalBridge:
		return "OneAdditionalBridge"
	case RouteType_TwoAdditionalBridge:
		return "TwoAdditionalBridge"
	case RouteType_ThreeAdditionalBridge:
		return "ThreeAdditionalBridge"
	case RouteType_FourAdditionalBridge:
		return "FourAdditionalBridge"
	case RouteType_FiveAdditionalBridge:
		return "FiveAdditionalBridge"
	}
	return ""
}

func (e RouteType) String() string {
	return e.name()
}
