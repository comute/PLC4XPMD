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

package readwrite

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

type BacnetipXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings
func init() {
	_ = strconv.Atoi
	_ = strings.Join
	_ = utils.Dump
	_ = xml.NewDecoder
}

func (m BacnetipXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "APDU":
		return nil, errors.New("APDU unmappable")
	case "BACnetTag":
		return nil, errors.New("BACnetTag unmappable")
	case "BACnetTagWithContent":
		return model.BACnetTagWithContentParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetError":
		return nil, errors.New("BACnetError unmappable")
	case "NLM":
		return nil, errors.New("NLM unmappable")
	case "BACnetConfirmedServiceRequest":
		return nil, errors.New("BACnetConfirmedServiceRequest unmappable")
	case "BACnetAddress":
		return model.BACnetAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConfirmedServiceACK":
		return nil, errors.New("BACnetConfirmedServiceACK unmappable")
	case "BACnetUnconfirmedServiceRequest":
		return nil, errors.New("BACnetUnconfirmedServiceRequest unmappable")
	case "BACnetServiceAck":
		return nil, errors.New("BACnetServiceAck unmappable")
	case "BVLC":
		return nil, errors.New("BVLC unmappable")
	case "NPDU":
		//var npduLength uint16
		atoi, err := strconv.Atoi(parserArguments[0])
		if err != nil {
			return nil, err
		}
		npduLength := uint16(atoi)
		return model.NPDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), npduLength)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}

// Deprecated: will be removed in favor of Parse soon
func (m BacnetipXmlParserHelper) ParseOld(typeName string, xmlString string) (interface{}, error) {
	switch typeName {
	case "APDU":
		var obj *model.APDU
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetTag":
		var obj *model.BACnetTag
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetTagWithContent":
		var obj *model.BACnetTagWithContent
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetError":
		var obj *model.BACnetError
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "NLM":
		var obj *model.NLM
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetConfirmedServiceRequest":
		var obj *model.BACnetConfirmedServiceRequest
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetAddress":
		var obj *model.BACnetAddress
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetConfirmedServiceACK":
		var obj *model.BACnetConfirmedServiceACK
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetUnconfirmedServiceRequest":
		var obj *model.BACnetUnconfirmedServiceRequest
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BACnetServiceAck":
		var obj *model.BACnetServiceAck
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "BVLC":
		var obj *model.BVLC
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "NPDU":
		var obj *model.NPDU
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
