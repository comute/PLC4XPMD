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

public class BACnetCalendarEntryDate extends BACnetCalendarEntry implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagDate dateValue;

  public BACnetCalendarEntryDate(BACnetTagHeader peekedTagHeader, BACnetContextTagDate dateValue) {
    super(peekedTagHeader);
    this.dateValue = dateValue;
  }

  public BACnetContextTagDate getDateValue() {
    return dateValue;
  }

  @Override
  protected void serializeBACnetCalendarEntryChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetCalendarEntryDate");

    // Simple Field (dateValue)
    writeSimpleField("dateValue", dateValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetCalendarEntryDate");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetCalendarEntryDate _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dateValue)
    lengthInBits += dateValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetCalendarEntryBuilder staticParseBACnetCalendarEntryBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetCalendarEntryDate");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagDate dateValue =
        readSimpleField(
            "dateValue",
            readComplex(
                () ->
                    (BACnetContextTagDate)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (0), (BACnetDataType) (BACnetDataType.DATE)),
                readBuffer));

    readBuffer.closeContext("BACnetCalendarEntryDate");
    // Create the instance
    return new BACnetCalendarEntryDateBuilderImpl(dateValue);
  }

  public static class BACnetCalendarEntryDateBuilderImpl
      implements BACnetCalendarEntry.BACnetCalendarEntryBuilder {
    private final BACnetContextTagDate dateValue;

    public BACnetCalendarEntryDateBuilderImpl(BACnetContextTagDate dateValue) {
      this.dateValue = dateValue;
    }

    public BACnetCalendarEntryDate build(BACnetTagHeader peekedTagHeader) {
      BACnetCalendarEntryDate bACnetCalendarEntryDate =
          new BACnetCalendarEntryDate(peekedTagHeader, dateValue);
      return bACnetCalendarEntryDate;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetCalendarEntryDate)) {
      return false;
    }
    BACnetCalendarEntryDate that = (BACnetCalendarEntryDate) o;
    return (getDateValue() == that.getDateValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDateValue());
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
