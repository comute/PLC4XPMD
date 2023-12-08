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

// OpcuaNodeIdServicesVariableNamespaces is an enum
type OpcuaNodeIdServicesVariableNamespaces int32

type IOpcuaNodeIdServicesVariableNamespaces interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri                              OpcuaNodeIdServicesVariableNamespaces = 11647
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion                          OpcuaNodeIdServicesVariableNamespaces = 11648
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate                  OpcuaNodeIdServicesVariableNamespaces = 11649
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset                         OpcuaNodeIdServicesVariableNamespaces = 11650
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes                         OpcuaNodeIdServicesVariableNamespaces = 11651
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange                  OpcuaNodeIdServicesVariableNamespaces = 11652
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern                 OpcuaNodeIdServicesVariableNamespaces = 11653
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size                        OpcuaNodeIdServicesVariableNamespaces = 11655
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount                   OpcuaNodeIdServicesVariableNamespaces = 11658
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments         OpcuaNodeIdServicesVariableNamespaces = 11660
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments        OpcuaNodeIdServicesVariableNamespaces = 11661
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments        OpcuaNodeIdServicesVariableNamespaces = 11663
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments         OpcuaNodeIdServicesVariableNamespaces = 11665
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments        OpcuaNodeIdServicesVariableNamespaces = 11666
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments        OpcuaNodeIdServicesVariableNamespaces = 11668
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments  OpcuaNodeIdServicesVariableNamespaces = 11670
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments OpcuaNodeIdServicesVariableNamespaces = 11671
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments  OpcuaNodeIdServicesVariableNamespaces = 11673
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable                    OpcuaNodeIdServicesVariableNamespaces = 12692
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable                OpcuaNodeIdServicesVariableNamespaces = 12693
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType                    OpcuaNodeIdServicesVariableNamespaces = 13400
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions                    OpcuaNodeIdServicesVariableNamespaces = 16140
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions                OpcuaNodeIdServicesVariableNamespaces = 16141
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions                 OpcuaNodeIdServicesVariableNamespaces = 16142
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength         OpcuaNodeIdServicesVariableNamespaces = 24247
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime            OpcuaNodeIdServicesVariableNamespaces = 25203
	OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion                      OpcuaNodeIdServicesVariableNamespaces = 25268
)

var OpcuaNodeIdServicesVariableNamespacesValues []OpcuaNodeIdServicesVariableNamespaces

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNamespacesValues = []OpcuaNodeIdServicesVariableNamespaces{
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime,
		OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion,
	}
}

func OpcuaNodeIdServicesVariableNamespacesByValue(value int32) (enum OpcuaNodeIdServicesVariableNamespaces, ok bool) {
	switch value {
	case 11647:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri, true
	case 11648:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion, true
	case 11649:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate, true
	case 11650:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset, true
	case 11651:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes, true
	case 11652:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange, true
	case 11653:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern, true
	case 11655:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size, true
	case 11658:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount, true
	case 11660:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments, true
	case 11661:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments, true
	case 11663:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments, true
	case 11665:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments, true
	case 11666:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments, true
	case 11668:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments, true
	case 11670:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments, true
	case 11671:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments, true
	case 11673:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments, true
	case 12692:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable, true
	case 12693:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable, true
	case 13400:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType, true
	case 16140:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions, true
	case 16141:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions, true
	case 16142:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions, true
	case 24247:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength, true
	case 25203:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime, true
	case 25268:
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNamespacesByName(value string) (enum OpcuaNodeIdServicesVariableNamespaces, ok bool) {
	switch value {
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime, true
	case "NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion":
		return OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNamespacesKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNamespacesValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableNamespaces(structType any) OpcuaNodeIdServicesVariableNamespaces {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableNamespaces {
		if sOpcuaNodeIdServicesVariableNamespaces, ok := typ.(OpcuaNodeIdServicesVariableNamespaces); ok {
			return sOpcuaNodeIdServicesVariableNamespaces
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableNamespaces) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableNamespaces) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNamespacesParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableNamespaces, error) {
	return OpcuaNodeIdServicesVariableNamespacesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNamespacesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableNamespaces, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableNamespaces", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableNamespaces")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNamespacesByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableNamespaces")
		return OpcuaNodeIdServicesVariableNamespaces(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableNamespaces) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableNamespaces) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableNamespaces", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableNamespaces) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceUri"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceVersion"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespacePublicationDate"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset:
		return "NamespacesType_NamespaceIdentifier_Placeholder_IsNamespaceSubset"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes:
		return "NamespacesType_NamespaceIdentifier_Placeholder_StaticNodeIdTypes"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange:
		return "NamespacesType_NamespaceIdentifier_Placeholder_StaticNumericNodeIdRange"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern:
		return "NamespacesType_NamespaceIdentifier_Placeholder_StaticStringNodeIdPattern"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Size"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_OpenCount"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Open_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Close_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Read_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Write_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_GetPosition_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_SetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_Writable"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_UserWritable"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MimeType"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions:
		return "NamespacesType_NamespaceIdentifier_Placeholder_DefaultRolePermissions"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions:
		return "NamespacesType_NamespaceIdentifier_Placeholder_DefaultUserRolePermissions"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions:
		return "NamespacesType_NamespaceIdentifier_Placeholder_DefaultAccessRestrictions"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_MaxByteStringLength"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime:
		return "NamespacesType_NamespaceIdentifier_Placeholder_NamespaceFile_LastModifiedTime"
	case OpcuaNodeIdServicesVariableNamespaces_NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion:
		return "NamespacesType_NamespaceIdentifier_Placeholder_ConfigurationVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableNamespaces) String() string {
	return e.PLC4XEnumName()
}
