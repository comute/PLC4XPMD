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

public class BACnetDateRange implements Message {

  // Properties.
  protected final BACnetApplicationTagDate startDate;
  protected final BACnetApplicationTagDate endDate;

  public BACnetDateRange(BACnetApplicationTagDate startDate, BACnetApplicationTagDate endDate) {
    super();
    this.startDate = startDate;
    this.endDate = endDate;
  }

  public BACnetApplicationTagDate getStartDate() {
    return startDate;
  }

  public BACnetApplicationTagDate getEndDate() {
    return endDate;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetDateRange");

    // Simple Field (startDate)
    writeSimpleField("startDate", startDate, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (endDate)
    writeSimpleField("endDate", endDate, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetDateRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetDateRange _value = this;

    // Simple field (startDate)
    lengthInBits += startDate.getLengthInBits();

    // Simple field (endDate)
    lengthInBits += endDate.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetDateRange staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetDateRange staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetDateRange");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagDate startDate =
        readSimpleField(
            "startDate",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagDate) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagDate endDate =
        readSimpleField(
            "endDate",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagDate) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetDateRange");
    // Create the instance
    BACnetDateRange _bACnetDateRange;
    _bACnetDateRange = new BACnetDateRange(startDate, endDate);
    return _bACnetDateRange;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetDateRange)) {
      return false;
    }
    BACnetDateRange that = (BACnetDateRange) o;
    return (getStartDate() == that.getStartDate()) && (getEndDate() == that.getEndDate()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getStartDate(), getEndDate());
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
