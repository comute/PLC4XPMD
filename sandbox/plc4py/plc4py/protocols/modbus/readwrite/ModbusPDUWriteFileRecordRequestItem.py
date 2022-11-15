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
class ModbusPDUWriteFileRecordRequestItem(PlcMessage):
            referenceType: short
            fileNumber: int
            recordNumber: int
            recordData: byte[]



    def __post_init__(self):
super().__init__( )



    def getReferenceType(self) -> short:
        return referenceType

    def getFileNumber(self) -> int:
        return fileNumber

    def getRecordNumber(self) -> int:
        return recordNumber

    def getRecordData(self) -> byte[]:
        return recordData


    def serialize(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusPDUWriteFileRecordRequestItem")

                        # Simple Field (referenceType)
                            writeSimpleField("referenceType", referenceType, writeUnsignedShort(writeBuffer, 8))

                        # Simple Field (fileNumber)
                            writeSimpleField("fileNumber", fileNumber, writeUnsignedInt(writeBuffer, 16))

                        # Simple Field (recordNumber)
                            writeSimpleField("recordNumber", recordNumber, writeUnsignedInt(writeBuffer, 16))

                        # Implicit Field (recordLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
                        int recordLength = (int) ((COUNT(getRecordData())) / (2))
                        writeImplicitField("recordLength", recordLength, writeUnsignedInt(writeBuffer, 16))

                        # Array Field (recordData)
                        writeByteArrayField("recordData", recordData, writeByteArray(writeBuffer, 8))

            writeBuffer.popContext("ModbusPDUWriteFileRecordRequestItem")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = 0
        _value: ModbusPDUWriteFileRecordRequestItem = self

        # Simple field (referenceType)
        lengthInBits += 8

        # Simple field (fileNumber)
        lengthInBits += 16

        # Simple field (recordNumber)
        lengthInBits += 16

        # Implicit Field (recordLength)
        lengthInBits += 16

        # Array field
        if recordData is not None):
            lengthInBits += 8 * recordData.length


        return lengthInBits


    def staticParse(readBuffer: ReadBuffer , args) -> ModbusPDUWriteFileRecordRequestItem:
        positionAware: PositionAware = readBuffer
        return staticParse(readBuffer)
    }

    def  staticParse(readBuffer: ReadBuffer) -> ModbusPDUWriteFileRecordRequestItem:
        readBuffer.pullContext("ModbusPDUWriteFileRecordRequestItem")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                referenceType: short = readSimpleField("referenceType", readUnsignedShort(readBuffer, 8))

                fileNumber: int = readSimpleField("fileNumber", readUnsignedInt(readBuffer, 16))

                recordNumber: int = readSimpleField("recordNumber", readUnsignedInt(readBuffer, 16))

                recordLength: int = readImplicitField("recordLength", readUnsignedInt(readBuffer, 16))

                    recordData: byte[] = readBuffer.readByteArray("recordData", Math.toIntExact((recordLength) * (2)))

    readBuffer.closeContext("ModbusPDUWriteFileRecordRequestItem")
    # Create the instance
        _modbusPDUWriteFileRecordRequestItem: ModbusPDUWriteFileRecordRequestItem = ModbusPDUWriteFileRecordRequestItem(
            referenceType, 
            fileNumber, 
            recordNumber, 
            recordData
        )
        return _modbusPDUWriteFileRecordRequestItem


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusPDUWriteFileRecordRequestItem):
            return False

        that: ModbusPDUWriteFileRecordRequestItem = ModbusPDUWriteFileRecordRequestItem(o)
        return
            (getReferenceType() == that.getReferenceType()) &&
            (getFileNumber() == that.getFileNumber()) &&
            (getRecordNumber() == that.getRecordNumber()) &&
            (getRecordData() == that.getRecordData()) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            getReferenceType(),
            getFileNumber(),
            getRecordNumber(),
            getRecordData()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"

