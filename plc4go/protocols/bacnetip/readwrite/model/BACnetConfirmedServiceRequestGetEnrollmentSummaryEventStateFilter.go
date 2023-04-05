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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter is an enum
type BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter uint8

type IBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter interface {
	fmt.Stringer
	utils.Serializable
}

const (
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter = 0
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_FAULT     BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter = 1
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_NORMAL    BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter = 2
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ALL       BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter = 3
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ACTIVE    BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter = 4
)

var BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterValues []BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterValues = []BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter{
		BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL,
		BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_FAULT,
		BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_NORMAL,
		BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ALL,
		BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ACTIVE,
	}
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterByValue(value uint8) (enum BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, ok bool) {
	switch value {
	case 0:
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL, true
	case 1:
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_FAULT, true
	case 2:
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_NORMAL, true
	case 3:
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ALL, true
	case 4:
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ACTIVE, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterByName(value string) (enum BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, ok bool) {
	switch value {
	case "OFFNORMAL":
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL, true
	case "FAULT":
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_FAULT, true
	case "NORMAL":
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_NORMAL, true
	case "ALL":
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ALL, true
	case "ACTIVE":
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ACTIVE, true
	}
	return 0, false
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter(structType interface{}) BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter {
		if sBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, ok := typ.(BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter); ok {
			return sBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterParse(ctx context.Context, theBytes []byte) (BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter")
	}
	if enum, ok := BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) PLC4XEnumName() string {
	switch e {
	case BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL:
		return "OFFNORMAL"
	case BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_FAULT:
		return "FAULT"
	case BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_NORMAL:
		return "NORMAL"
	case BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ALL:
		return "ALL"
	case BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_ACTIVE:
		return "ACTIVE"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter) String() string {
	return e.PLC4XEnumName()
}
