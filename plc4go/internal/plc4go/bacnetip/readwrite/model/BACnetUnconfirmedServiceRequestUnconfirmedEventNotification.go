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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedEventNotification struct {
	Parent *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedEventNotification interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) ServiceChoice() uint8 {
	return 0x03
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestUnconfirmedEventNotification() *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		Parent: NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUnconfirmedEventNotification); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedEventNotification(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(io *utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedEventNotification{
		Parent: &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) String() string {
	return string(m.Box("BACnetUnconfirmedServiceRequestUnconfirmedEventNotification", utils.DefaultWidth*2))
}

func (m BACnetUnconfirmedServiceRequestUnconfirmedEventNotification) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetUnconfirmedServiceRequestUnconfirmedEventNotification"
	}
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
