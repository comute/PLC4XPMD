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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7VarRequestParameterItemAddress struct {
	Address *S7Address
	Parent  *S7VarRequestParameterItem
}

// The corresponding interface
type IS7VarRequestParameterItemAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7VarRequestParameterItemAddress) ItemType() uint8 {
	return 0x12
}

func (m *S7VarRequestParameterItemAddress) InitializeParent(parent *S7VarRequestParameterItem) {
}

func NewS7VarRequestParameterItemAddress(address *S7Address) *S7VarRequestParameterItem {
	child := &S7VarRequestParameterItemAddress{
		Address: address,
		Parent:  NewS7VarRequestParameterItem(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7VarRequestParameterItemAddress(structType interface{}) *S7VarRequestParameterItemAddress {
	castFunc := func(typ interface{}) *S7VarRequestParameterItemAddress {
		if casted, ok := typ.(S7VarRequestParameterItemAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*S7VarRequestParameterItemAddress); ok {
			return casted
		}
		if casted, ok := typ.(S7VarRequestParameterItem); ok {
			return CastS7VarRequestParameterItemAddress(casted.Child)
		}
		if casted, ok := typ.(*S7VarRequestParameterItem); ok {
			return CastS7VarRequestParameterItemAddress(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7VarRequestParameterItemAddress) GetTypeName() string {
	return "S7VarRequestParameterItemAddress"
}

func (m *S7VarRequestParameterItemAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7VarRequestParameterItemAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += m.Address.LengthInBits()

	return lengthInBits
}

func (m *S7VarRequestParameterItemAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7VarRequestParameterItemAddressParse(io *utils.ReadBuffer) (*S7VarRequestParameterItem, error) {

	// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := io.ReadUint8(8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field")
	}

	// Simple Field (address)
	address, _addressErr := S7AddressParse(io)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}

	// Create a partially initialized instance
	_child := &S7VarRequestParameterItemAddress{
		Address: address,
		Parent:  &S7VarRequestParameterItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7VarRequestParameterItemAddress) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(m.Address.LengthInBytes())
		_itemLengthErr := io.WriteUint8(8, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (address)
		_addressErr := m.Address.Serialize(io)
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7VarRequestParameterItemAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "address":
				var dt *S7Address
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Address = dt
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *S7VarRequestParameterItemAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
		return err
	}
	return nil
}

func (m S7VarRequestParameterItemAddress) String() string {
	return string(m.Box("S7VarRequestParameterItemAddress", utils.DefaultWidth*2))
}

func (m S7VarRequestParameterItemAddress) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7VarRequestParameterItemAddress"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Address", m.Address, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
