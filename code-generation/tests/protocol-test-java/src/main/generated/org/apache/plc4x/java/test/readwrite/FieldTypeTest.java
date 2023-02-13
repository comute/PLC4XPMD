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
package org.apache.plc4x.java.test.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class FieldTypeTest implements Message {

  // Constant values.
  public static final Short CONSTFIELD = 5;

  // Properties.
  protected final short simpleField;
  protected final List<Short> arrayFieldCount;
  protected final List<Short> arrayFieldLength;
  protected final List<Short> arrayFieldTerminated;
  protected final short assertField;
  protected final EnumTypeParameters enumField;
  protected final short manualField;
  protected final Short optionalField;
  protected final short peekField;

  public FieldTypeTest(
      short simpleField,
      List<Short> arrayFieldCount,
      List<Short> arrayFieldLength,
      List<Short> arrayFieldTerminated,
      short assertField,
      EnumTypeParameters enumField,
      short manualField,
      Short optionalField,
      short peekField) {
    super();
    this.simpleField = simpleField;
    this.arrayFieldCount = arrayFieldCount;
    this.arrayFieldLength = arrayFieldLength;
    this.arrayFieldTerminated = arrayFieldTerminated;
    this.assertField = assertField;
    this.enumField = enumField;
    this.manualField = manualField;
    this.optionalField = optionalField;
    this.peekField = peekField;
  }

  public short getSimpleField() {
    return simpleField;
  }

  public List<Short> getArrayFieldCount() {
    return arrayFieldCount;
  }

  public List<Short> getArrayFieldLength() {
    return arrayFieldLength;
  }

  public List<Short> getArrayFieldTerminated() {
    return arrayFieldTerminated;
  }

  public short getAssertField() {
    return assertField;
  }

  public EnumTypeParameters getEnumField() {
    return enumField;
  }

  public short getManualField() {
    return manualField;
  }

  public Short getOptionalField() {
    return optionalField;
  }

  public short getPeekField() {
    return peekField;
  }

  public short getVirtualField() {
    return (short) ((getConstField()) + (10));
  }

  public short getConstField() {
    return CONSTFIELD;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    throw new SerializationException("Unknown field not serializable");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    FieldTypeTest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (simpleField)
    lengthInBits += 8;

    // Array field
    if (arrayFieldCount != null) {
      lengthInBits += 8 * arrayFieldCount.size();
    }

    // Array field
    if (arrayFieldLength != null) {
      lengthInBits += 8 * arrayFieldLength.size();
    }

    // Array field
    if (arrayFieldTerminated != null) {
      lengthInBits += 8 * arrayFieldTerminated.size();
    }

    // Checksum Field (checksum)
    lengthInBits += 8;

    // Const Field (constField)
    lengthInBits += 8;

    // Enum Field (enumField)
    lengthInBits += 8;

    // Implicit Field (implicitField)
    lengthInBits += 8;

    // Manual Field (manualField)
    lengthInBits += (simpleField) * (8);

    // Optional Field (optionalField)
    if (optionalField != null) {
      lengthInBits += 8;
    }

    // Padding Field (padding)
    int _timesPadding = (int) (simpleField);
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Unknown field
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static FieldTypeTest staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static FieldTypeTest staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("FieldTypeTest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short simpleField = readSimpleField("simpleField", readUnsignedShort(readBuffer, 8));

    List<Short> arrayFieldCount =
        readCountArrayField("arrayFieldCount", readUnsignedShort(readBuffer, 8), 5);

    List<Short> arrayFieldLength =
        readLengthArrayField("arrayFieldLength", readUnsignedShort(readBuffer, 8), 5);

    List<Short> arrayFieldTerminated =
        readTerminatedArrayField(
            "arrayFieldTerminated", readUnsignedShort(readBuffer, 8), () -> ((boolean) (true)));

    short assertField =
        readAssertField("assertField", readUnsignedShort(readBuffer, 8), (short) (42));

    short checksumField =
        readChecksumField(
            "checksumField",
            readUnsignedShort(readBuffer, 8),
            (short) (org.apache.plc4x.java.test.readwrite.utils.StaticHelper.crcUint8(-(23))));

    short constField =
        readConstField("constField", readUnsignedShort(readBuffer, 8), FieldTypeTest.CONSTFIELD);

    EnumTypeParameters enumField =
        readEnumField(
            "enumField",
            "EnumTypeParameters",
            readEnum(EnumTypeParameters::firstEnumForFieldIntType, readSignedByte(readBuffer, 8)));

    short implicitField = readImplicitField("implicitField", readUnsignedShort(readBuffer, 8));

    short manualField =
        readManualField(
            "manualField",
            readBuffer,
            () ->
                (short)
                    (org.apache.plc4x.java.test.readwrite.utils.StaticHelper.readManualField(
                        readBuffer, simpleField)));

    Short optionalField =
        readOptionalField("optionalField", readUnsignedShort(readBuffer, 8), (simpleField) == (5));

    readPaddingField(readUnsignedShort(readBuffer, 8), (int) (simpleField));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    readUnknownField("unknown", readUnsignedShort(readBuffer, 8));
    short virtualField = readVirtualField("virtualField", short.class, (constField) + (10));
    // Validation
    if (!(false)) {
      throw new ParseValidationException("Validation Message");
    }

    short peekField = readPeekField("peekField", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("FieldTypeTest");
    // Create the instance
    FieldTypeTest _fieldTypeTest;
    _fieldTypeTest =
        new FieldTypeTest(
            simpleField,
            arrayFieldCount,
            arrayFieldLength,
            arrayFieldTerminated,
            assertField,
            enumField,
            manualField,
            optionalField,
            peekField);
    return _fieldTypeTest;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FieldTypeTest)) {
      return false;
    }
    FieldTypeTest that = (FieldTypeTest) o;
    return (getSimpleField() == that.getSimpleField())
        && (getArrayFieldCount() == that.getArrayFieldCount())
        && (getArrayFieldLength() == that.getArrayFieldLength())
        && (getArrayFieldTerminated() == that.getArrayFieldTerminated())
        && (getAssertField() == that.getAssertField())
        && (getEnumField() == that.getEnumField())
        && (getManualField() == that.getManualField())
        && (getOptionalField() == that.getOptionalField())
        && (getPeekField() == that.getPeekField())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getSimpleField(),
        getArrayFieldCount(),
        getArrayFieldLength(),
        getArrayFieldTerminated(),
        getAssertField(),
        getEnumField(),
        getManualField(),
        getOptionalField(),
        getPeekField());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
