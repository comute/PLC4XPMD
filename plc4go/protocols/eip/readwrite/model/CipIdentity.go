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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CipIdentity_ZEROES1 uint32 = 0x00000000
const CipIdentity_ZEROES2 uint32 = 0x00000000

// CipIdentity is the corresponding interface of CipIdentity
type CipIdentity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CommandSpecificDataItem
	// GetEncapsulationProtocolVersion returns EncapsulationProtocolVersion (property field)
	GetEncapsulationProtocolVersion() uint16
	// GetSocketAddressFamily returns SocketAddressFamily (property field)
	GetSocketAddressFamily() uint16
	// GetSocketAddressPort returns SocketAddressPort (property field)
	GetSocketAddressPort() uint16
	// GetSocketAddressAddress returns SocketAddressAddress (property field)
	GetSocketAddressAddress() []uint8
	// GetVendorId returns VendorId (property field)
	GetVendorId() uint16
	// GetDeviceType returns DeviceType (property field)
	GetDeviceType() uint16
	// GetProductCode returns ProductCode (property field)
	GetProductCode() uint16
	// GetRevisionMajor returns RevisionMajor (property field)
	GetRevisionMajor() uint8
	// GetRevisionMinor returns RevisionMinor (property field)
	GetRevisionMinor() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() uint32
	// GetProductName returns ProductName (property field)
	GetProductName() string
	// GetState returns State (property field)
	GetState() uint8
}

// CipIdentityExactly can be used when we want exactly this type and not a type which fulfills CipIdentity.
// This is useful for switch cases.
type CipIdentityExactly interface {
	CipIdentity
	isCipIdentity() bool
}

// _CipIdentity is the data-structure of this message
type _CipIdentity struct {
	*_CommandSpecificDataItem
	EncapsulationProtocolVersion uint16
	SocketAddressFamily          uint16
	SocketAddressPort            uint16
	SocketAddressAddress         []uint8
	VendorId                     uint16
	DeviceType                   uint16
	ProductCode                  uint16
	RevisionMajor                uint8
	RevisionMinor                uint8
	Status                       uint16
	SerialNumber                 uint32
	ProductName                  string
	State                        uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipIdentity) GetItemType() uint16 {
	return 0x000C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipIdentity) InitializeParent(parent CommandSpecificDataItem) {}

func (m *_CipIdentity) GetParent() CommandSpecificDataItem {
	return m._CommandSpecificDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipIdentity) GetEncapsulationProtocolVersion() uint16 {
	return m.EncapsulationProtocolVersion
}

func (m *_CipIdentity) GetSocketAddressFamily() uint16 {
	return m.SocketAddressFamily
}

func (m *_CipIdentity) GetSocketAddressPort() uint16 {
	return m.SocketAddressPort
}

func (m *_CipIdentity) GetSocketAddressAddress() []uint8 {
	return m.SocketAddressAddress
}

func (m *_CipIdentity) GetVendorId() uint16 {
	return m.VendorId
}

func (m *_CipIdentity) GetDeviceType() uint16 {
	return m.DeviceType
}

func (m *_CipIdentity) GetProductCode() uint16 {
	return m.ProductCode
}

func (m *_CipIdentity) GetRevisionMajor() uint8 {
	return m.RevisionMajor
}

func (m *_CipIdentity) GetRevisionMinor() uint8 {
	return m.RevisionMinor
}

func (m *_CipIdentity) GetStatus() uint16 {
	return m.Status
}

func (m *_CipIdentity) GetSerialNumber() uint32 {
	return m.SerialNumber
}

func (m *_CipIdentity) GetProductName() string {
	return m.ProductName
}

func (m *_CipIdentity) GetState() uint8 {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipIdentity) GetZeroes1() uint32 {
	return CipIdentity_ZEROES1
}

func (m *_CipIdentity) GetZeroes2() uint32 {
	return CipIdentity_ZEROES2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipIdentity factory function for _CipIdentity
func NewCipIdentity(encapsulationProtocolVersion uint16, socketAddressFamily uint16, socketAddressPort uint16, socketAddressAddress []uint8, vendorId uint16, deviceType uint16, productCode uint16, revisionMajor uint8, revisionMinor uint8, status uint16, serialNumber uint32, productName string, state uint8) *_CipIdentity {
	_result := &_CipIdentity{
		EncapsulationProtocolVersion: encapsulationProtocolVersion,
		SocketAddressFamily:          socketAddressFamily,
		SocketAddressPort:            socketAddressPort,
		SocketAddressAddress:         socketAddressAddress,
		VendorId:                     vendorId,
		DeviceType:                   deviceType,
		ProductCode:                  productCode,
		RevisionMajor:                revisionMajor,
		RevisionMinor:                revisionMinor,
		Status:                       status,
		SerialNumber:                 serialNumber,
		ProductName:                  productName,
		State:                        state,
		_CommandSpecificDataItem:     NewCommandSpecificDataItem(),
	}
	_result._CommandSpecificDataItem._CommandSpecificDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipIdentity(structType any) CipIdentity {
	if casted, ok := structType.(CipIdentity); ok {
		return casted
	}
	if casted, ok := structType.(*CipIdentity); ok {
		return *casted
	}
	return nil
}

func (m *_CipIdentity) GetTypeName() string {
	return "CipIdentity"
}

func (m *_CipIdentity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 16

	// Simple field (encapsulationProtocolVersion)
	lengthInBits += 16

	// Simple field (socketAddressFamily)
	lengthInBits += 16

	// Simple field (socketAddressPort)
	lengthInBits += 16

	// Array field
	if len(m.SocketAddressAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.SocketAddressAddress))
	}

	// Const Field (zeroes1)
	lengthInBits += 32

	// Const Field (zeroes2)
	lengthInBits += 32

	// Simple field (vendorId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (productCode)
	lengthInBits += 16

	// Simple field (revisionMajor)
	lengthInBits += 8

	// Simple field (revisionMinor)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (serialNumber)
	lengthInBits += 32

	// Implicit Field (productNameLength)
	lengthInBits += 8

	// Simple field (productName)
	lengthInBits += uint16(int32(uint8(len(m.GetProductName()))) * int32(int32(8)))

	// Simple field (state)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipIdentity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipIdentityParse(ctx context.Context, theBytes []byte) (CipIdentity, error) {
	return CipIdentityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CipIdentityParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (CipIdentity, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CipIdentity, error) {
		return CipIdentityParseWithBuffer(ctx, readBuffer)
	}
}

func CipIdentityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CipIdentity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipIdentity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipIdentity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint16](ctx, "itemLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	encapsulationProtocolVersion, err := ReadSimpleField(ctx, "encapsulationProtocolVersion", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encapsulationProtocolVersion' field"))
	}

	socketAddressFamily, err := ReadSimpleField(ctx, "socketAddressFamily", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressFamily' field"))
	}

	socketAddressPort, err := ReadSimpleField(ctx, "socketAddressPort", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressPort' field"))
	}

	socketAddressAddress, err := ReadCountArrayField[uint8](ctx, "socketAddressAddress", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressAddress' field"))
	}

	zeroes1, err := ReadConstField[uint32](ctx, "zeroes1", ReadUnsignedInt(readBuffer, uint8(32)), CipIdentity_ZEROES1)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zeroes1' field"))
	}
	_ = zeroes1

	zeroes2, err := ReadConstField[uint32](ctx, "zeroes2", ReadUnsignedInt(readBuffer, uint8(32)), CipIdentity_ZEROES2)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zeroes2' field"))
	}
	_ = zeroes2

	vendorId, err := ReadSimpleField(ctx, "vendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}

	deviceType, err := ReadSimpleField(ctx, "deviceType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceType' field"))
	}

	productCode, err := ReadSimpleField(ctx, "productCode", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productCode' field"))
	}

	revisionMajor, err := ReadSimpleField(ctx, "revisionMajor", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisionMajor' field"))
	}

	revisionMinor, err := ReadSimpleField(ctx, "revisionMinor", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisionMinor' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}

	serialNumber, err := ReadSimpleField(ctx, "serialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serialNumber' field"))
	}

	productNameLength, err := ReadImplicitField[uint8](ctx, "productNameLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productNameLength' field"))
	}
	_ = productNameLength

	productName, err := ReadSimpleField(ctx, "productName", ReadString(readBuffer, uint32(int32(productNameLength)*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productName' field"))
	}

	state, err := ReadSimpleField(ctx, "state", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'state' field"))
	}

	if closeErr := readBuffer.CloseContext("CipIdentity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipIdentity")
	}

	// Create a partially initialized instance
	_child := &_CipIdentity{
		_CommandSpecificDataItem:     &_CommandSpecificDataItem{},
		EncapsulationProtocolVersion: encapsulationProtocolVersion,
		SocketAddressFamily:          socketAddressFamily,
		SocketAddressPort:            socketAddressPort,
		SocketAddressAddress:         socketAddressAddress,
		VendorId:                     vendorId,
		DeviceType:                   deviceType,
		ProductCode:                  productCode,
		RevisionMajor:                revisionMajor,
		RevisionMinor:                revisionMinor,
		Status:                       status,
		SerialNumber:                 serialNumber,
		ProductName:                  productName,
		State:                        state,
	}
	_child._CommandSpecificDataItem._CommandSpecificDataItemChildRequirements = _child
	return _child, nil
}

func (m *_CipIdentity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipIdentity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipIdentity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipIdentity")
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint16(uint16(uint16(34)) + uint16(uint8(len(m.GetProductName()))))
		_itemLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("itemLength", 16, uint16((itemLength)))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (encapsulationProtocolVersion)
		encapsulationProtocolVersion := uint16(m.GetEncapsulationProtocolVersion())
		_encapsulationProtocolVersionErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("encapsulationProtocolVersion", 16, uint16((encapsulationProtocolVersion)))
		if _encapsulationProtocolVersionErr != nil {
			return errors.Wrap(_encapsulationProtocolVersionErr, "Error serializing 'encapsulationProtocolVersion' field")
		}

		// Simple Field (socketAddressFamily)
		socketAddressFamily := uint16(m.GetSocketAddressFamily())
		_socketAddressFamilyErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("socketAddressFamily", 16, uint16((socketAddressFamily)))
		if _socketAddressFamilyErr != nil {
			return errors.Wrap(_socketAddressFamilyErr, "Error serializing 'socketAddressFamily' field")
		}

		// Simple Field (socketAddressPort)
		socketAddressPort := uint16(m.GetSocketAddressPort())
		_socketAddressPortErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("socketAddressPort", 16, uint16((socketAddressPort)))
		if _socketAddressPortErr != nil {
			return errors.Wrap(_socketAddressPortErr, "Error serializing 'socketAddressPort' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "socketAddressAddress", m.GetSocketAddressAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'socketAddressAddress' field")
		}

		// Const Field (zeroes1)
		_zeroes1Err := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint32("zeroes1", 32, uint32(0x00000000))
		if _zeroes1Err != nil {
			return errors.Wrap(_zeroes1Err, "Error serializing 'zeroes1' field")
		}

		// Const Field (zeroes2)
		_zeroes2Err := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint32("zeroes2", 32, uint32(0x00000000))
		if _zeroes2Err != nil {
			return errors.Wrap(_zeroes2Err, "Error serializing 'zeroes2' field")
		}

		// Simple Field (vendorId)
		vendorId := uint16(m.GetVendorId())
		_vendorIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("vendorId", 16, uint16((vendorId)))
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (deviceType)
		deviceType := uint16(m.GetDeviceType())
		_deviceTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("deviceType", 16, uint16((deviceType)))
		if _deviceTypeErr != nil {
			return errors.Wrap(_deviceTypeErr, "Error serializing 'deviceType' field")
		}

		// Simple Field (productCode)
		productCode := uint16(m.GetProductCode())
		_productCodeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("productCode", 16, uint16((productCode)))
		if _productCodeErr != nil {
			return errors.Wrap(_productCodeErr, "Error serializing 'productCode' field")
		}

		// Simple Field (revisionMajor)
		revisionMajor := uint8(m.GetRevisionMajor())
		_revisionMajorErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("revisionMajor", 8, uint8((revisionMajor)))
		if _revisionMajorErr != nil {
			return errors.Wrap(_revisionMajorErr, "Error serializing 'revisionMajor' field")
		}

		// Simple Field (revisionMinor)
		revisionMinor := uint8(m.GetRevisionMinor())
		_revisionMinorErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("revisionMinor", 8, uint8((revisionMinor)))
		if _revisionMinorErr != nil {
			return errors.Wrap(_revisionMinorErr, "Error serializing 'revisionMinor' field")
		}

		// Simple Field (status)
		status := uint16(m.GetStatus())
		_statusErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("status", 16, uint16((status)))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (serialNumber)
		serialNumber := uint32(m.GetSerialNumber())
		_serialNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("serialNumber", 32, uint32((serialNumber)))
		if _serialNumberErr != nil {
			return errors.Wrap(_serialNumberErr, "Error serializing 'serialNumber' field")
		}

		// Implicit Field (productNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		productNameLength := uint8(uint8(len(m.GetProductName())))
		_productNameLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("productNameLength", 8, uint8((productNameLength)))
		if _productNameLengthErr != nil {
			return errors.Wrap(_productNameLengthErr, "Error serializing 'productNameLength' field")
		}

		// Simple Field (productName)
		productName := string(m.GetProductName())
		_productNameErr := /*TODO: migrate me*/ writeBuffer.WriteString("productName", uint32((uint8(len(m.GetProductName())))*(8)), (productName), utils.WithEncoding("UTF-8)"))
		if _productNameErr != nil {
			return errors.Wrap(_productNameErr, "Error serializing 'productName' field")
		}

		// Simple Field (state)
		state := uint8(m.GetState())
		_stateErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("state", 8, uint8((state)))
		if _stateErr != nil {
			return errors.Wrap(_stateErr, "Error serializing 'state' field")
		}

		if popErr := writeBuffer.PopContext("CipIdentity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipIdentity")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipIdentity) isCipIdentity() bool {
	return true
}

func (m *_CipIdentity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
