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

public class BACnetLogRecordLogDatumAnyValue extends BACnetLogRecordLogDatum implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetConstructedData anyValue;

  // Arguments.
  protected final Short tagNumber;

  public BACnetLogRecordLogDatumAnyValue(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetConstructedData anyValue,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.anyValue = anyValue;
    this.tagNumber = tagNumber;
  }

  public BACnetConstructedData getAnyValue() {
    return anyValue;
  }

  @Override
  protected void serializeBACnetLogRecordLogDatumChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogRecordLogDatumAnyValue");

    // Optional Field (anyValue) (Can be skipped, if the value is null)
    writeOptionalField("anyValue", anyValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetLogRecordLogDatumAnyValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogRecordLogDatumAnyValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (anyValue)
    if (anyValue != null) {
      lengthInBits += anyValue.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetLogRecordLogDatumBuilder staticParseBACnetLogRecordLogDatumBuilder(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetLogRecordLogDatumAnyValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetConstructedData anyValue =
        readOptionalField(
            "anyValue",
            readComplex(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (10),
                        (BACnetObjectType) (BACnetObjectType.VENDOR_PROPRIETARY_VALUE),
                        (BACnetPropertyIdentifier)
                            (BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE),
                        (BACnetTagPayloadUnsignedInteger) (null)),
                readBuffer));

    readBuffer.closeContext("BACnetLogRecordLogDatumAnyValue");
    // Create the instance
    return new BACnetLogRecordLogDatumAnyValueBuilderImpl(anyValue, tagNumber);
  }

  public static class BACnetLogRecordLogDatumAnyValueBuilderImpl
      implements BACnetLogRecordLogDatum.BACnetLogRecordLogDatumBuilder {
    private final BACnetConstructedData anyValue;
    private final Short tagNumber;

    public BACnetLogRecordLogDatumAnyValueBuilderImpl(
        BACnetConstructedData anyValue, Short tagNumber) {
      this.anyValue = anyValue;
      this.tagNumber = tagNumber;
    }

    public BACnetLogRecordLogDatumAnyValue build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetLogRecordLogDatumAnyValue bACnetLogRecordLogDatumAnyValue =
          new BACnetLogRecordLogDatumAnyValue(
              openingTag, peekedTagHeader, closingTag, anyValue, tagNumber);
      return bACnetLogRecordLogDatumAnyValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogRecordLogDatumAnyValue)) {
      return false;
    }
    BACnetLogRecordLogDatumAnyValue that = (BACnetLogRecordLogDatumAnyValue) o;
    return (getAnyValue() == that.getAnyValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAnyValue());
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
