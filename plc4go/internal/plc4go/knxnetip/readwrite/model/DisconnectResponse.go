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
type DisconnectResponse struct {
	CommunicationChannelId uint8
	Status                 Status
	Parent                 *KnxNetIpMessage
}

// The corresponding interface
type IDisconnectResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DisconnectResponse) MsgType() uint16 {
	return 0x020A
}

func (m *DisconnectResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDisconnectResponse(communicationChannelId uint8, status Status) *KnxNetIpMessage {
	child := &DisconnectResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDisconnectResponse(structType interface{}) *DisconnectResponse {
	castFunc := func(typ interface{}) *DisconnectResponse {
		if casted, ok := typ.(DisconnectResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*DisconnectResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastDisconnectResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastDisconnectResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DisconnectResponse) GetTypeName() string {
	return "DisconnectResponse"
}

func (m *DisconnectResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DisconnectResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *DisconnectResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DisconnectResponseParse(io *utils.ReadBuffer) (*KnxNetIpMessage, error) {

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := io.ReadUint8(8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (status)
	status, _statusErr := StatusParse(io)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Create a partially initialized instance
	_child := &DisconnectResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DisconnectResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.CommunicationChannelId)
		_communicationChannelIdErr := io.WriteUint8(8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		_statusErr := m.Status.Serialize(io)
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *DisconnectResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "communicationChannelId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommunicationChannelId = data
			case "status":
				var data Status
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
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

func (m *DisconnectResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	return nil
}

func (m DisconnectResponse) String() string {
	return string(m.Box("DisconnectResponse", utils.DefaultWidth*2))
}

func (m DisconnectResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "DisconnectResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("CommunicationChannelId", m.CommunicationChannelId, width-2))
	boxes = append(boxes, utils.BoxAnything("Status", m.Status, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
