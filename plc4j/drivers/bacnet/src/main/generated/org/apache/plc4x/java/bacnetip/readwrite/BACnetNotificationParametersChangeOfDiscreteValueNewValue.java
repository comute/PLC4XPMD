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

public abstract class BACnetNotificationParametersChangeOfDiscreteValueNewValue implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetTagHeader peekedTagHeader;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfDiscreteValueNewValue(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.peekedTagHeader = peekedTagHeader;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  public boolean getPeekedIsContextTag() {
    return (boolean) ((getPeekedTagHeader().getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS));
  }

  protected abstract void serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean peekedIsContextTag = getPeekedIsContextTag();
    writeBuffer.writeVirtual("peekedIsContextTag", peekedIsContextTag);

    // Switch field (Serialize the sub-type)
    serializeBACnetNotificationParametersChangeOfDiscreteValueNewValueChild(writeBuffer);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetNotificationParametersChangeOfDiscreteValueNewValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfDiscreteValueNewValue staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());
    boolean peekedIsContextTag =
        readVirtualField(
            "peekedIsContextTag",
            boolean.class,
            (peekedTagHeader.getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS));
    // Validation
    if (!(((!(peekedIsContextTag)))
        || ((((peekedIsContextTag) && ((peekedTagHeader.getLengthValueType()) != (0x6)))
            && ((peekedTagHeader.getLengthValueType()) != (0x7)))))) {
      throw new ParseValidationException("unexpected opening or closing tag");
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0x1)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x2)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x3)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x9)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x7)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x6)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xA)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xB)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xC)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) true)) {
      builder =
          BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
              .staticParseBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder(
                  readBuffer, tagNumber);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + " "
              + "peekedIsContextTag="
              + peekedIsContextTag
              + "]");
    }

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue");
    // Create the instance
    BACnetNotificationParametersChangeOfDiscreteValueNewValue
        _bACnetNotificationParametersChangeOfDiscreteValueNewValue =
            builder.build(openingTag, peekedTagHeader, closingTag, tagNumber);
    return _bACnetNotificationParametersChangeOfDiscreteValueNewValue;
  }

  public interface BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
    BACnetNotificationParametersChangeOfDiscreteValueNewValue build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfDiscreteValueNewValue)) {
      return false;
    }
    BACnetNotificationParametersChangeOfDiscreteValueNewValue that =
        (BACnetNotificationParametersChangeOfDiscreteValueNewValue) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getPeekedTagHeader() == that.getPeekedTagHeader())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getPeekedTagHeader(), getClosingTag());
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
