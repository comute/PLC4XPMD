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

public class BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
    extends BACnetNotificationParametersChangeOfDiscreteValueNewValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagCharacterString characterStringValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagCharacterString characterStringValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.characterStringValue = characterStringValue;
    this.tagNumber = tagNumber;
  }

  public BACnetApplicationTagCharacterString getCharacterStringValue() {
    return characterStringValue;
  }

  @Override
  protected void serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext(
        "BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString");

    // Simple Field (characterStringValue)
    writeSimpleField(
        "characterStringValue", characterStringValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext(
        "BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString _value = this;

    // Simple field (characterStringValue)
    lengthInBits += characterStringValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
      staticParseBuilder(ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext(
        "BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagCharacterString characterStringValue =
        readSimpleField(
            "characterStringValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagCharacterString)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext(
        "BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString");
    // Create the instance
    return new BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder(
        characterStringValue, tagNumber);
  }

  public static
  class BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder
      implements BACnetNotificationParametersChangeOfDiscreteValueNewValue
          .BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
    private final BACnetApplicationTagCharacterString characterStringValue;
    private final Short tagNumber;

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringBuilder(
        BACnetApplicationTagCharacterString characterStringValue, Short tagNumber) {

      this.characterStringValue = characterStringValue;
      this.tagNumber = tagNumber;
    }

    public BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
          bACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString =
              new BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString(
                  openingTag, peekedTagHeader, closingTag, characterStringValue, tagNumber);
      return bACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString)) {
      return false;
    }
    BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString that =
        (BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString) o;
    return (getCharacterStringValue() == that.getCharacterStringValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCharacterStringValue());
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
