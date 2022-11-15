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
class ModbusPDUReadDeviceIdentificationResponse(PlcMessage,ModbusPDU):
            level: ModbusDeviceInformationLevel
            individualAccess: c_bool
            conformityLevel: ModbusDeviceInformationConformityLevel
            moreFollows: ModbusDeviceInformationMoreFollows
            nextObjectId: c_uint8
            objects: []ModbusDeviceInformationObject
            MEITYPE: c_uint8 = 0x0E

    # Accessors for discriminator values.
    def c_bool getErrorFlag() {
        return (c_bool) false
    }
    def c_uint8 getFunctionFlag() {
        return (c_uint8) 0x2B
    }
    def c_bool getResponse() {
        return (c_bool) true
    }


    def __post_init__(self):
super().__init__( )



    def getLevel(self) -> ModbusDeviceInformationLevel:
        return level

    def getIndividualAccess(self) -> c_bool:
        return individualAccess

    def getConformityLevel(self) -> ModbusDeviceInformationConformityLevel:
        return conformityLevel

    def getMoreFollows(self) -> ModbusDeviceInformationMoreFollows:
        return moreFollows

    def getNextObjectId(self) -> c_uint8:
        return nextObjectId

    def getObjects(self) -> []ModbusDeviceInformationObject:
        return objects

    def getMeiType(self) -> c_uint8:
        return MEITYPE


    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
            startPos: int = positionAware.getPos()
            writeBuffer.pushContext("ModbusPDUReadDeviceIdentificationResponse")

                        # Const Field (meiType)
                        writeConstField("meiType", MEITYPE, writeUnsignedShort(writeBuffer, 8))

                        # Simple Field (level)
                            writeSimpleEnumField("level", "ModbusDeviceInformationLevel", level, new DataWriterEnumDefault<>(ModbusDeviceInformationLevel::getValue, ModbusDeviceInformationLevel::name, writeUnsignedShort(writeBuffer, 8)))


                        # Simple Field (individualAccess)
                            writeSimpleField("individualAccess", individualAccess, writeBoolean(writeBuffer))

                        # Simple Field (conformityLevel)
                            writeSimpleEnumField("conformityLevel", "ModbusDeviceInformationConformityLevel", conformityLevel, new DataWriterEnumDefault<>(ModbusDeviceInformationConformityLevel::getValue, ModbusDeviceInformationConformityLevel::name, writeUnsignedShort(writeBuffer, 7)))


                        # Simple Field (moreFollows)
                            writeSimpleEnumField("moreFollows", "ModbusDeviceInformationMoreFollows", moreFollows, new DataWriterEnumDefault<>(ModbusDeviceInformationMoreFollows::getValue, ModbusDeviceInformationMoreFollows::name, writeUnsignedShort(writeBuffer, 8)))


                        # Simple Field (nextObjectId)
                            writeSimpleField("nextObjectId", nextObjectId, writeUnsignedShort(writeBuffer, 8))

                        # Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
                        c_uint8 numberOfObjects = (c_uint8) (COUNT(getObjects()))
                        writeImplicitField("numberOfObjects", numberOfObjects, writeUnsignedShort(writeBuffer, 8))

                        # Array Field (objects)
                        writeComplexTypeArrayField("objects", objects, writeBuffer)

            writeBuffer.popContext("ModbusPDUReadDeviceIdentificationResponse")


    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadDeviceIdentificationResponse = self

        # Const Field (meiType)
        lengthInBits += 8

        # Simple field (level)
        lengthInBits += 8

        # Simple field (individualAccess)
        lengthInBits += 1

        # Simple field (conformityLevel)
        lengthInBits += 7

        # Simple field (moreFollows)
        lengthInBits += 8

        # Simple field (nextObjectId)
        lengthInBits += 8

        # Implicit Field (numberOfObjects)
        lengthInBits += 8

        # Array field
        if objects is not None):
            i: int = 0
            for element in objects:
                last: boolean = ++i >= objects.size()
                lengthInBits += element.getLengthInBits()



        return lengthInBits


    def  staticParseBuilder(readBuffer: ReadBuffer, c_bool response) -> ModbusPDUReadDeviceIdentificationResponseBuilder:
        readBuffer.pullContext("ModbusPDUReadDeviceIdentificationResponse")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

                meiType: c_uint8 = readConstField("meiType", readUnsignedShort(readBuffer, 8), ModbusPDUReadDeviceIdentificationResponse.MEITYPE)

                level: ModbusDeviceInformationLevel = readEnumField("level", "ModbusDeviceInformationLevel", new DataReaderEnumDefault<>(ModbusDeviceInformationLevel::enumForValue, readUnsignedShort(readBuffer, 8)))

                individualAccess: c_bool = readSimpleField("individualAccess", readBoolean(readBuffer))

                conformityLevel: ModbusDeviceInformationConformityLevel = readEnumField("conformityLevel", "ModbusDeviceInformationConformityLevel", new DataReaderEnumDefault<>(ModbusDeviceInformationConformityLevel::enumForValue, readUnsignedShort(readBuffer, 7)))

                moreFollows: ModbusDeviceInformationMoreFollows = readEnumField("moreFollows", "ModbusDeviceInformationMoreFollows", new DataReaderEnumDefault<>(ModbusDeviceInformationMoreFollows::enumForValue, readUnsignedShort(readBuffer, 8)))

                nextObjectId: c_uint8 = readSimpleField("nextObjectId", readUnsignedShort(readBuffer, 8))

                numberOfObjects: c_uint8 = readImplicitField("numberOfObjects", readUnsignedShort(readBuffer, 8))

                        objects: []ModbusDeviceInformationObject = readCountArrayField("objects", new DataReaderComplexDefault<>(() -> ModbusDeviceInformationObject.staticParse(readBuffer), readBuffer), numberOfObjects)

    readBuffer.closeContext("ModbusPDUReadDeviceIdentificationResponse")
    # Create the instance
        return ModbusPDUReadDeviceIdentificationResponseBuilder(
            level, 
            individualAccess, 
            conformityLevel, 
            moreFollows, 
            nextObjectId, 
            objects
        
        )


    def equals(self, o: object) -> bool:
        if this == o:
            return True

        if not (instanceof(o, ModbusPDUReadDeviceIdentificationResponse):
            return False

        that: ModbusPDUReadDeviceIdentificationResponse = ModbusPDUReadDeviceIdentificationResponse(o)
        return
            (getLevel() == that.getLevel()) &&
            (getIndividualAccess() == that.getIndividualAccess()) &&
            (getConformityLevel() == that.getConformityLevel()) &&
            (getMoreFollows() == that.getMoreFollows()) &&
            (getNextObjectId() == that.getNextObjectId()) &&
            (getObjects() == that.getObjects()) &&
            super().equals(that) &&
            True

    def hashCode(self) -> int:
        return Objects.hash(
            super().hashCode(),
            getLevel(),
            getIndividualAccess(),
            getConformityLevel(),
            getMoreFollows(),
            getNextObjectId(),
            getObjects()
        )

    def toString(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(true, true)
        try:
            writeBufferBoxBased.writeSerializable(this)
        except SerializationException:
            raise RuntimeException(e)

        return "\n" + writeBufferBoxBased.getBox().toString()+ "\n"


class ModbusPDUReadDeviceIdentificationResponseBuilder(ModbusPDUModbusPDUBuilder: level: ModbusDeviceInformationLevel individualAccess: c_bool conformityLevel: ModbusDeviceInformationConformityLevel moreFollows: ModbusDeviceInformationMoreFollows nextObjectId: c_uint8 objects: []ModbusDeviceInformationObjectdef ModbusPDUReadDeviceIdentificationResponseBuilder( ModbusDeviceInformationLevel level, c_bool individualAccess, ModbusDeviceInformationConformityLevel conformityLevel, ModbusDeviceInformationMoreFollows moreFollows, c_uint8 nextObjectId, []ModbusDeviceInformationObject objects ):        self.level = level
        self.individualAccess = individualAccess
        self.conformityLevel = conformityLevel
        self.moreFollows = moreFollows
        self.nextObjectId = nextObjectId
        self.objects = objects


        def build(
        ) -> ModbusPDUReadDeviceIdentificationResponse:
        modbusPDUReadDeviceIdentificationResponse: ModbusPDUReadDeviceIdentificationResponse = ModbusPDUReadDeviceIdentificationResponse(
            level, 
            individualAccess, 
            conformityLevel, 
            moreFollows, 
            nextObjectId, 
            objects
)
        return modbusPDUReadDeviceIdentificationResponse


