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
class ModbusPDUError(PlcMessage,ModbusPDU):
            exceptionCode: ModbusErrorCode

    # Accessors for discriminator values.
    def c_bool getErrorFlag() {
        return (c_bool) true
    }
    def c_uint8 getFunctionFlag() {
        return 0
    }
    def c_bool getResponse() {
        return False
    }


    def __post_init__(self):
super().__init__( )



    def getExceptionCode(self) -> ModbusErrorCode:
        return exceptionCode


    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusPDUError")

                        # Simple Field (exceptionCode)
                            writeSimpleEnumField("exceptionCode", "ModbusErrorCode", exceptionCode, new DataWriterEnumDefault<>(ModbusErrorCode::getValue, ModbusErrorCode::name, writeUnsignedShort(writeBuffer, 8)))


            writeBuffer.popContext("ModbusPDUError")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUError = self

        # Simple field (exceptionCode)
        lengthInBits += 8

        return lengthInBits


    def  staticParseBuilder(readBuffer: ReadBuffer, c_bool response) -> ModbusPDUErrorBuilder:
        readBuffer.pullContext("ModbusPDUError")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                exceptionCode: ModbusErrorCode = readEnumField("exceptionCode", "ModbusErrorCode", new DataReaderEnumDefault<>(ModbusErrorCode::enumForValue, readUnsignedShort(readBuffer, 8)))

    readBuffer.closeContext("ModbusPDUError")
    # Create the instance
        return ModbusPDUErrorBuilder(
            exceptionCode
        
        )


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusPDUError):
            return False

        that: ModbusPDUError = ModbusPDUError(o)
        return
            (getExceptionCode() == that.getExceptionCode()) &&
            super().equals(that) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            super().hashCode(),
            getExceptionCode()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"


class ModbusPDUErrorBuilder(ModbusPDUModbusPDUBuilder: exceptionCode: ModbusErrorCodedef ModbusPDUErrorBuilder( ModbusErrorCode exceptionCode ):        self.exceptionCode = exceptionCode


        def build(
        ) -> ModbusPDUError:
        modbusPDUError: ModbusPDUError = ModbusPDUError(
            exceptionCode
)
        return modbusPDUError


