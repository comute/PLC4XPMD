//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type CEMIRawCon struct {
	CEMI
}

// The corresponding interface
type ICEMIRawCon interface {
	ICEMI
	Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m CEMIRawCon) MessageCode() uint8 {
	return 0x2F
}

func (m CEMIRawCon) initialize() spi.Message {
	return m
}

func NewCEMIRawCon() CEMIInitializer {
	return &CEMIRawCon{}
}

func CastICEMIRawCon(structType interface{}) ICEMIRawCon {
	castFunc := func(typ interface{}) ICEMIRawCon {
		if iCEMIRawCon, ok := typ.(ICEMIRawCon); ok {
			return iCEMIRawCon
		}
		return nil
	}
	return castFunc(structType)
}

func CastCEMIRawCon(structType interface{}) CEMIRawCon {
	castFunc := func(typ interface{}) CEMIRawCon {
		if sCEMIRawCon, ok := typ.(CEMIRawCon); ok {
			return sCEMIRawCon
		}
		return CEMIRawCon{}
	}
	return castFunc(structType)
}

func (m CEMIRawCon) LengthInBits() uint16 {
	var lengthInBits = m.CEMI.LengthInBits()

	return lengthInBits
}

func (m CEMIRawCon) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CEMIRawConParse(io *spi.ReadBuffer) (CEMIInitializer, error) {

	// Create the instance
	return NewCEMIRawCon(), nil
}

func (m CEMIRawCon) Serialize(io spi.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return CEMISerialize(io, m.CEMI, CastICEMI(m), ser)
}
