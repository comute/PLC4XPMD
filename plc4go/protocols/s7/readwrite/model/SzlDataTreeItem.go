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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SzlDataTreeItem is the corresponding interface of SzlDataTreeItem
type SzlDataTreeItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetItemIndex returns ItemIndex (property field)
	GetItemIndex() uint16
	// GetMlfb returns Mlfb (property field)
	GetMlfb() []byte
	// GetModuleTypeId returns ModuleTypeId (property field)
	GetModuleTypeId() uint16
	// GetAusbg returns Ausbg (property field)
	GetAusbg() uint16
	// GetAusbe returns Ausbe (property field)
	GetAusbe() uint16
}

// SzlDataTreeItemExactly can be used when we want exactly this type and not a type which fulfills SzlDataTreeItem.
// This is useful for switch cases.
type SzlDataTreeItemExactly interface {
	SzlDataTreeItem
	isSzlDataTreeItem() bool
}

// _SzlDataTreeItem is the data-structure of this message
type _SzlDataTreeItem struct {
	ItemIndex    uint16
	Mlfb         []byte
	ModuleTypeId uint16
	Ausbg        uint16
	Ausbe        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SzlDataTreeItem) GetItemIndex() uint16 {
	return m.ItemIndex
}

func (m *_SzlDataTreeItem) GetMlfb() []byte {
	return m.Mlfb
}

func (m *_SzlDataTreeItem) GetModuleTypeId() uint16 {
	return m.ModuleTypeId
}

func (m *_SzlDataTreeItem) GetAusbg() uint16 {
	return m.Ausbg
}

func (m *_SzlDataTreeItem) GetAusbe() uint16 {
	return m.Ausbe
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSzlDataTreeItem factory function for _SzlDataTreeItem
func NewSzlDataTreeItem(itemIndex uint16, mlfb []byte, moduleTypeId uint16, ausbg uint16, ausbe uint16) *_SzlDataTreeItem {
	return &_SzlDataTreeItem{ItemIndex: itemIndex, Mlfb: mlfb, ModuleTypeId: moduleTypeId, Ausbg: ausbg, Ausbe: ausbe}
}

// Deprecated: use the interface for direct cast
func CastSzlDataTreeItem(structType any) SzlDataTreeItem {
	if casted, ok := structType.(SzlDataTreeItem); ok {
		return casted
	}
	if casted, ok := structType.(*SzlDataTreeItem); ok {
		return *casted
	}
	return nil
}

func (m *_SzlDataTreeItem) GetTypeName() string {
	return "SzlDataTreeItem"
}

func (m *_SzlDataTreeItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (itemIndex)
	lengthInBits += 16

	// Array field
	if len(m.Mlfb) > 0 {
		lengthInBits += 8 * uint16(len(m.Mlfb))
	}

	// Simple field (moduleTypeId)
	lengthInBits += 16

	// Simple field (ausbg)
	lengthInBits += 16

	// Simple field (ausbe)
	lengthInBits += 16

	return lengthInBits
}

func (m *_SzlDataTreeItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SzlDataTreeItemParse(ctx context.Context, theBytes []byte) (SzlDataTreeItem, error) {
	return SzlDataTreeItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SzlDataTreeItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
		return SzlDataTreeItemParseWithBuffer(ctx, readBuffer)
	}
}

func SzlDataTreeItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SzlDataTreeItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SzlDataTreeItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SzlDataTreeItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemIndex, err := ReadSimpleField(ctx, "itemIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemIndex' field"))
	}

	mlfb, err := readBuffer.ReadByteArray("mlfb", int(int32(20)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mlfb' field"))
	}

	moduleTypeId, err := ReadSimpleField(ctx, "moduleTypeId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moduleTypeId' field"))
	}

	ausbg, err := ReadSimpleField(ctx, "ausbg", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ausbg' field"))
	}

	ausbe, err := ReadSimpleField(ctx, "ausbe", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ausbe' field"))
	}

	if closeErr := readBuffer.CloseContext("SzlDataTreeItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SzlDataTreeItem")
	}

	// Create the instance
	return &_SzlDataTreeItem{
		ItemIndex:    itemIndex,
		Mlfb:         mlfb,
		ModuleTypeId: moduleTypeId,
		Ausbg:        ausbg,
		Ausbe:        ausbe,
	}, nil
}

func (m *_SzlDataTreeItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SzlDataTreeItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SzlDataTreeItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SzlDataTreeItem")
	}

	// Simple Field (itemIndex)
	itemIndex := uint16(m.GetItemIndex())
	_itemIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("itemIndex", 16, uint16((itemIndex)))
	if _itemIndexErr != nil {
		return errors.Wrap(_itemIndexErr, "Error serializing 'itemIndex' field")
	}

	if err := WriteByteArrayField(ctx, "mlfb", m.GetMlfb(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'mlfb' field")
	}

	// Simple Field (moduleTypeId)
	moduleTypeId := uint16(m.GetModuleTypeId())
	_moduleTypeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("moduleTypeId", 16, uint16((moduleTypeId)))
	if _moduleTypeIdErr != nil {
		return errors.Wrap(_moduleTypeIdErr, "Error serializing 'moduleTypeId' field")
	}

	// Simple Field (ausbg)
	ausbg := uint16(m.GetAusbg())
	_ausbgErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("ausbg", 16, uint16((ausbg)))
	if _ausbgErr != nil {
		return errors.Wrap(_ausbgErr, "Error serializing 'ausbg' field")
	}

	// Simple Field (ausbe)
	ausbe := uint16(m.GetAusbe())
	_ausbeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("ausbe", 16, uint16((ausbe)))
	if _ausbeErr != nil {
		return errors.Wrap(_ausbeErr, "Error serializing 'ausbe' field")
	}

	if popErr := writeBuffer.PopContext("SzlDataTreeItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SzlDataTreeItem")
	}
	return nil
}

func (m *_SzlDataTreeItem) isSzlDataTreeItem() bool {
	return true
}

func (m *_SzlDataTreeItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
