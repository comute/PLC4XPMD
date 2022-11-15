#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License") you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from abc import ABC, abstractmethod
from dataclasses import dataclass



@dataclass
class ModbusPDUReadFileRecordRequest(PlcMessage,ModbusPDU):
            items: List<ModbusPDUReadFileRecordRequestItem>

    # Accessors for discriminator values.
    def Boolean getErrorFlag() {
        return (boolean) false
    }
    def Short getFunctionFlag() {
        return (short) 0x14
    }
    def Boolean getResponse() {
        return (boolean) false
    }


    def __post_init__(self):
super().__init__( )



    def getItems(self) -> List<ModbusPDUReadFileRecordRequestItem>:
        return items


    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusPDUReadFileRecordRequest")

                        # Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
                        short byteCount = (short) (ARRAY_SIZE_IN_BYTES(getItems()))
                        writeImplicitField("byteCount", byteCount, writeUnsignedShort(writeBuffer, 8))

                        # Array Field (items)
                        writeComplexTypeArrayField("items", items, writeBuffer)

            writeBuffer.popContext("ModbusPDUReadFileRecordRequest")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadFileRecordRequest = self

        # Implicit Field (byteCount)
        lengthInBits += 8

        # Array field
        if items is not None):
            for element in items:
                lengthInBits += element.getLengthInBits()



        return lengthInBits


    def  staticParseBuilder(readBuffer: ReadBuffer, Boolean response) -> ModbusPDUReadFileRecordRequestBuilder:
        readBuffer.pullContext("ModbusPDUReadFileRecordRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                byteCount: short = readImplicitField("byteCount", readUnsignedShort(readBuffer, 8))

                            items: List<ModbusPDUReadFileRecordRequestItem> = readLengthArrayField("items", new DataReaderComplexDefault<>(() -> ModbusPDUReadFileRecordRequestItem.staticParse(readBuffer), readBuffer), byteCount)

    readBuffer.closeContext("ModbusPDUReadFileRecordRequest")
    # Create the instance
        return ModbusPDUReadFileRecordRequestBuilder(
            items
        
        )

        class ModbusPDUReadFileRecordRequestBuilder(ModbusPDUModbusPDUBuilder {
        items: List<ModbusPDUReadFileRecordRequestItem>

        def ModbusPDUReadFileRecordRequestBuilder(
            List<ModbusPDUReadFileRecordRequestItem> items
        
        ):
            self.items = items


        def build(
        ) -> ModbusPDUReadFileRecordRequest:
            modbusPDUReadFileRecordRequest: ModbusPDUReadFileRecordRequest = ModbusPDUReadFileRecordRequest(
                items
)
            return modbusPDUReadFileRecordRequest


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusPDUReadFileRecordRequest):
            return False

        that: ModbusPDUReadFileRecordRequest = ModbusPDUReadFileRecordRequest(o)
        return
            (getItems() == that.getItems()) &&
            super().equals(that) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            super().hashCode(),
            getItems()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"

