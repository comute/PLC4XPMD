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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetAccessUserTypeTagged implements Message {

  // Properties.
  protected final BACnetTagHeader header;
  protected final BACnetAccessUserType value;
  protected final long proprietaryValue;

  // Arguments.
  protected final Short tagNumber;
  protected final TagClass tagClass;

  public BACnetAccessUserTypeTagged(
      BACnetTagHeader header,
      BACnetAccessUserType value,
      long proprietaryValue,
      Short tagNumber,
      TagClass tagClass) {
    super();
    this.header = header;
    this.value = value;
    this.proprietaryValue = proprietaryValue;
    this.tagNumber = tagNumber;
    this.tagClass = tagClass;
  }

  public BACnetTagHeader getHeader() {
    return header;
  }

  public BACnetAccessUserType getValue() {
    return value;
  }

  public long getProprietaryValue() {
    return proprietaryValue;
  }

  public boolean getIsProprietary() {
    return (boolean) ((getValue()) == (BACnetAccessUserType.VENDOR_PROPRIETARY_VALUE));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetAccessUserTypeTagged");

    // Simple Field (header)
    writeSimpleField("header", header, new DataWriterComplexDefault<>(writeBuffer));

    // Manual Field (value)
    writeManualField(
        "value",
        () ->
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.writeEnumGeneric(
                writeBuffer, value),
        writeBuffer);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isProprietary = getIsProprietary();
    writeBuffer.writeVirtual("isProprietary", isProprietary);

    // Manual Field (proprietaryValue)
    writeManualField(
        "proprietaryValue",
        () ->
            org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.writeProprietaryEnumGeneric(
                writeBuffer, proprietaryValue, isProprietary),
        writeBuffer);

    writeBuffer.popContext("BACnetAccessUserTypeTagged");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetAccessUserTypeTagged _value = this;

    // Simple field (header)
    lengthInBits += header.getLengthInBits();

    // Manual Field (value)
    lengthInBits += ((_value.getIsProprietary()) ? 0 : ((header.getActualLength()) * (8)));

    // A virtual field doesn't have any in- or output.

    // Manual Field (proprietaryValue)
    lengthInBits += ((_value.getIsProprietary()) ? ((header.getActualLength()) * (8)) : 0);

    return lengthInBits;
  }

  public static BACnetAccessUserTypeTagged staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    Short tagNumber;
    if (args[0] instanceof Short) {
      tagNumber = (Short) args[0];
    } else if (args[0] instanceof String) {
      tagNumber = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    TagClass tagClass;
    if (args[1] instanceof TagClass) {
      tagClass = (TagClass) args[1];
    } else if (args[1] instanceof String) {
      tagClass = TagClass.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type TagClass or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, tagNumber, tagClass);
  }

  public static BACnetAccessUserTypeTagged staticParse(
      ReadBuffer readBuffer, Short tagNumber, TagClass tagClass) throws ParseException {
    readBuffer.pullContext("BACnetAccessUserTypeTagged");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetTagHeader header =
        readSimpleField(
            "header",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    // Validation
    if (!((header.getTagClass()) == (tagClass))) {
      throw new ParseValidationException("tag class doesn't match");
    }
    // Validation
    if (!((((header.getTagClass()) == (TagClass.APPLICATION_TAGS)))
        || (((header.getActualTagNumber()) == (tagNumber))))) {
      throw new ParseAssertException("tagnumber doesn't match");
    }

    BACnetAccessUserType value =
        readManualField(
            "value",
            readBuffer,
            () ->
                (BACnetAccessUserType)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper.readEnumGeneric(
                        readBuffer,
                        header.getActualLength(),
                        BACnetAccessUserType.VENDOR_PROPRIETARY_VALUE)));
    boolean isProprietary =
        readVirtualField(
            "isProprietary",
            boolean.class,
            (value) == (BACnetAccessUserType.VENDOR_PROPRIETARY_VALUE));

    long proprietaryValue =
        readManualField(
            "proprietaryValue",
            readBuffer,
            () ->
                (long)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .readProprietaryEnumGeneric(
                            readBuffer, header.getActualLength(), isProprietary)));

    readBuffer.closeContext("BACnetAccessUserTypeTagged");
    // Create the instance
    BACnetAccessUserTypeTagged _bACnetAccessUserTypeTagged;
    _bACnetAccessUserTypeTagged =
        new BACnetAccessUserTypeTagged(header, value, proprietaryValue, tagNumber, tagClass);
    return _bACnetAccessUserTypeTagged;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetAccessUserTypeTagged)) {
      return false;
    }
    BACnetAccessUserTypeTagged that = (BACnetAccessUserTypeTagged) o;
    return (getHeader() == that.getHeader())
        && (getValue() == that.getValue())
        && (getProprietaryValue() == that.getProprietaryValue())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getHeader(), getValue(), getProprietaryValue());
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
