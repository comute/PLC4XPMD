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
class ModbusTcpADU(PlcMessage,ModbusADU):
            transactionIdentifier: c_uint16
            unitIdentifier: c_uint8
            pdu: ModbusPDU
        # Arguments.
            response: c_bool
            PROTOCOLIDENTIFIER: c_uint16 = 0x0000

    # Accessors for discriminator values.
    def DriverType getDriverType() {
        return DriverType.MODBUS_TCP
    }


    def __post_init__(self):
super().__init__( self.response )



    def getTransactionIdentifier(self) -> c_uint16:
        return transactionIdentifier

    def getUnitIdentifier(self) -> c_uint8:
        return unitIdentifier

    def getPdu(self) -> ModbusPDU:
        return pdu

    def getProtocolIdentifier(self) -> c_uint16:
        return PROTOCOLIDENTIFIER


    def serializeModbusADUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusTcpADU")

                        # Simple Field (transactionIdentifier)
                            writeSimpleField("transactionIdentifier", transactionIdentifier, writeUnsignedInt(writeBuffer, 16), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                        # Const Field (protocolIdentifier)
                        writeConstField("protocolIdentifier", PROTOCOLIDENTIFIER, writeUnsignedInt(writeBuffer, 16))

                        # Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
                        c_uint16 length = (c_uint16) ((getPdu().getLengthInBytes()) + (1))
                        writeImplicitField("length", length, writeUnsignedInt(writeBuffer, 16))

                        # Simple Field (unitIdentifier)
                            writeSimpleField("unitIdentifier", unitIdentifier, writeUnsignedShort(writeBuffer, 8), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                        # Simple Field (pdu)
                            writeSimpleField("pdu", pdu, new DataWriterComplexDefault<>(writeBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

            writeBuffer.popContext("ModbusTcpADU")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusTcpADU = self

        # Simple field (transactionIdentifier)
        lengthInBits += 16

        # Const Field (protocolIdentifier)
        lengthInBits += 16

        # Implicit Field (length)
        lengthInBits += 16

        # Simple field (unitIdentifier)
        lengthInBits += 8

        # Simple field (pdu)
        lengthInBits += pdu.getLengthInBits()

        return lengthInBits


    def  staticParseBuilder(readBuffer: ReadBuffer, DriverType driverType, c_bool response) -> ModbusTcpADUBuilder:
        readBuffer.pullContext("ModbusTcpADU")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                transactionIdentifier: c_uint16 = readSimpleField("transactionIdentifier", readUnsignedInt(readBuffer, 16), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                protocolIdentifier: c_uint16 = readConstField("protocolIdentifier", readUnsignedInt(readBuffer, 16), ModbusTcpADU.PROTOCOLIDENTIFIER, WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                length: c_uint16 = readImplicitField("length", readUnsignedInt(readBuffer, 16), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                unitIdentifier: c_uint8 = readSimpleField("unitIdentifier", readUnsignedShort(readBuffer, 8), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

                pdu: ModbusPDU = readSimpleField("pdu", new DataReaderComplexDefault<>(() -> ModbusPDU.staticParse(readBuffer, (c_bool) (response)), readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN))

    readBuffer.closeContext("ModbusTcpADU")
    # Create the instance
        return ModbusTcpADUBuilder(
            transactionIdentifier, 
            unitIdentifier, 
            pdu
            , 
            response
        
        )


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusTcpADU):
            return False

        that: ModbusTcpADU = ModbusTcpADU(o)
        return
            (getTransactionIdentifier() == that.getTransactionIdentifier()) &&
            (getUnitIdentifier() == that.getUnitIdentifier()) &&
            (getPdu() == that.getPdu()) &&
            super().equals(that) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            super().hashCode(),
            getTransactionIdentifier(),
            getUnitIdentifier(),
            getPdu()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"


class ModbusTcpADUBuilder(ModbusADUModbusADUBuilder: transactionIdentifier: c_uint16 unitIdentifier: c_uint8 pdu: ModbusPDU response: c_booldef ModbusTcpADUBuilder( c_uint16 transactionIdentifier, c_uint8 unitIdentifier, ModbusPDU pdu , c_bool response ):        self.transactionIdentifier = transactionIdentifier
        self.unitIdentifier = unitIdentifier
        self.pdu = pdu
            self.response = response


        def build(
            
                c_bool response
        ) -> ModbusTcpADU:
        modbusTcpADU: ModbusTcpADU = ModbusTcpADU(
            transactionIdentifier, 
            unitIdentifier, 
            pdu
            , 
                response
        )
        return modbusTcpADU


