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
class ModbusPDUReadCoilsRequest(PlcMessage,ModbusPDU):
            startingAddress: int
            quantity: int

    # Accessors for discriminator values.
    def Boolean getErrorFlag() {
        return (boolean) false
    }
    def Short getFunctionFlag() {
        return (short) 0x01
    }
    def Boolean getResponse() {
        return (boolean) false
    }


    def __post_init__(self):
super().__init__( )



    def getStartingAddress(self) -> int:
        return startingAddress

    def getQuantity(self) -> int:
        return quantity


    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusPDUReadCoilsRequest")

                        # Simple Field (startingAddress)
                            writeSimpleField("startingAddress", startingAddress, writeUnsignedInt(writeBuffer, 16))

                        # Simple Field (quantity)
                            writeSimpleField("quantity", quantity, writeUnsignedInt(writeBuffer, 16))

            writeBuffer.popContext("ModbusPDUReadCoilsRequest")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadCoilsRequest = self

        # Simple field (startingAddress)
        lengthInBits += 16

        # Simple field (quantity)
        lengthInBits += 16

        return lengthInBits


    def  staticParseBuilder(readBuffer: ReadBuffer, Boolean response) -> ModbusPDUReadCoilsRequestBuilder:
        readBuffer.pullContext("ModbusPDUReadCoilsRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                startingAddress: int = readSimpleField("startingAddress", readUnsignedInt(readBuffer, 16))

                quantity: int = readSimpleField("quantity", readUnsignedInt(readBuffer, 16))

    readBuffer.closeContext("ModbusPDUReadCoilsRequest")
    # Create the instance
        return ModbusPDUReadCoilsRequestBuilder(
            startingAddress, 
            quantity
        
        )

        class ModbusPDUReadCoilsRequestBuilder(ModbusPDUModbusPDUBuilder {
        startingAddress: int
        quantity: int

        def ModbusPDUReadCoilsRequestBuilder(
            int startingAddress, 
            int quantity
        
        ):
            self.startingAddress = startingAddress
            self.quantity = quantity


        def build(
        ) -> ModbusPDUReadCoilsRequest:
            modbusPDUReadCoilsRequest: ModbusPDUReadCoilsRequest = ModbusPDUReadCoilsRequest(
                startingAddress, 
                quantity
)
            return modbusPDUReadCoilsRequest


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusPDUReadCoilsRequest):
            return False

        that: ModbusPDUReadCoilsRequest = ModbusPDUReadCoilsRequest(o)
        return
            (getStartingAddress() == that.getStartingAddress()) &&
            (getQuantity() == that.getQuantity()) &&
            super().equals(that) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            super().hashCode(),
            getStartingAddress(),
            getQuantity()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"

