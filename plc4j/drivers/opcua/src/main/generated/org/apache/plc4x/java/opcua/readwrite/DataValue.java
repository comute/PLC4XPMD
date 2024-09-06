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
package org.apache.plc4x.java.opcua.readwrite;

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

public class DataValue implements Message {

  // Properties.
  protected final boolean serverPicosecondsSpecified;
  protected final boolean sourcePicosecondsSpecified;
  protected final boolean serverTimestampSpecified;
  protected final boolean sourceTimestampSpecified;
  protected final boolean statusCodeSpecified;
  protected final boolean valueSpecified;
  protected final Variant value;
  protected final StatusCode statusCode;
  protected final Long sourceTimestamp;
  protected final Integer sourcePicoseconds;
  protected final Long serverTimestamp;
  protected final Integer serverPicoseconds;

  public DataValue(
      boolean serverPicosecondsSpecified,
      boolean sourcePicosecondsSpecified,
      boolean serverTimestampSpecified,
      boolean sourceTimestampSpecified,
      boolean statusCodeSpecified,
      boolean valueSpecified,
      Variant value,
      StatusCode statusCode,
      Long sourceTimestamp,
      Integer sourcePicoseconds,
      Long serverTimestamp,
      Integer serverPicoseconds) {
    super();
    this.serverPicosecondsSpecified = serverPicosecondsSpecified;
    this.sourcePicosecondsSpecified = sourcePicosecondsSpecified;
    this.serverTimestampSpecified = serverTimestampSpecified;
    this.sourceTimestampSpecified = sourceTimestampSpecified;
    this.statusCodeSpecified = statusCodeSpecified;
    this.valueSpecified = valueSpecified;
    this.value = value;
    this.statusCode = statusCode;
    this.sourceTimestamp = sourceTimestamp;
    this.sourcePicoseconds = sourcePicoseconds;
    this.serverTimestamp = serverTimestamp;
    this.serverPicoseconds = serverPicoseconds;
  }

  public boolean getServerPicosecondsSpecified() {
    return serverPicosecondsSpecified;
  }

  public boolean getSourcePicosecondsSpecified() {
    return sourcePicosecondsSpecified;
  }

  public boolean getServerTimestampSpecified() {
    return serverTimestampSpecified;
  }

  public boolean getSourceTimestampSpecified() {
    return sourceTimestampSpecified;
  }

  public boolean getStatusCodeSpecified() {
    return statusCodeSpecified;
  }

  public boolean getValueSpecified() {
    return valueSpecified;
  }

  public Variant getValue() {
    return value;
  }

  public StatusCode getStatusCode() {
    return statusCode;
  }

  public Long getSourceTimestamp() {
    return sourceTimestamp;
  }

  public Integer getSourcePicoseconds() {
    return sourcePicoseconds;
  }

  public Long getServerTimestamp() {
    return serverTimestamp;
  }

  public Integer getServerPicoseconds() {
    return serverPicoseconds;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DataValue");

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (serverPicosecondsSpecified)
    writeSimpleField(
        "serverPicosecondsSpecified", serverPicosecondsSpecified, writeBoolean(writeBuffer));

    // Simple Field (sourcePicosecondsSpecified)
    writeSimpleField(
        "sourcePicosecondsSpecified", sourcePicosecondsSpecified, writeBoolean(writeBuffer));

    // Simple Field (serverTimestampSpecified)
    writeSimpleField(
        "serverTimestampSpecified", serverTimestampSpecified, writeBoolean(writeBuffer));

    // Simple Field (sourceTimestampSpecified)
    writeSimpleField(
        "sourceTimestampSpecified", sourceTimestampSpecified, writeBoolean(writeBuffer));

    // Simple Field (statusCodeSpecified)
    writeSimpleField("statusCodeSpecified", statusCodeSpecified, writeBoolean(writeBuffer));

    // Simple Field (valueSpecified)
    writeSimpleField("valueSpecified", valueSpecified, writeBoolean(writeBuffer));

    // Optional Field (value) (Can be skipped, if the value is null)
    writeOptionalField("value", value, writeComplex(writeBuffer));

    // Optional Field (statusCode) (Can be skipped, if the value is null)
    writeOptionalField("statusCode", statusCode, writeComplex(writeBuffer));

    // Optional Field (sourceTimestamp) (Can be skipped, if the value is null)
    writeOptionalField("sourceTimestamp", sourceTimestamp, writeSignedLong(writeBuffer, 64));

    // Optional Field (sourcePicoseconds) (Can be skipped, if the value is null)
    writeOptionalField("sourcePicoseconds", sourcePicoseconds, writeUnsignedInt(writeBuffer, 16));

    // Optional Field (serverTimestamp) (Can be skipped, if the value is null)
    writeOptionalField("serverTimestamp", serverTimestamp, writeSignedLong(writeBuffer, 64));

    // Optional Field (serverPicoseconds) (Can be skipped, if the value is null)
    writeOptionalField("serverPicoseconds", serverPicoseconds, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("DataValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DataValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (serverPicosecondsSpecified)
    lengthInBits += 1;

    // Simple field (sourcePicosecondsSpecified)
    lengthInBits += 1;

    // Simple field (serverTimestampSpecified)
    lengthInBits += 1;

    // Simple field (sourceTimestampSpecified)
    lengthInBits += 1;

    // Simple field (statusCodeSpecified)
    lengthInBits += 1;

    // Simple field (valueSpecified)
    lengthInBits += 1;

    // Optional Field (value)
    if (value != null) {
      lengthInBits += value.getLengthInBits();
    }

    // Optional Field (statusCode)
    if (statusCode != null) {
      lengthInBits += statusCode.getLengthInBits();
    }

    // Optional Field (sourceTimestamp)
    if (sourceTimestamp != null) {
      lengthInBits += 64;
    }

    // Optional Field (sourcePicoseconds)
    if (sourcePicoseconds != null) {
      lengthInBits += 16;
    }

    // Optional Field (serverTimestamp)
    if (serverTimestamp != null) {
      lengthInBits += 64;
    }

    // Optional Field (serverPicoseconds)
    if (serverPicoseconds != null) {
      lengthInBits += 16;
    }

    return lengthInBits;
  }

  public static DataValue staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DataValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 2), (byte) 0x00);

    boolean serverPicosecondsSpecified =
        readSimpleField("serverPicosecondsSpecified", readBoolean(readBuffer));

    boolean sourcePicosecondsSpecified =
        readSimpleField("sourcePicosecondsSpecified", readBoolean(readBuffer));

    boolean serverTimestampSpecified =
        readSimpleField("serverTimestampSpecified", readBoolean(readBuffer));

    boolean sourceTimestampSpecified =
        readSimpleField("sourceTimestampSpecified", readBoolean(readBuffer));

    boolean statusCodeSpecified = readSimpleField("statusCodeSpecified", readBoolean(readBuffer));

    boolean valueSpecified = readSimpleField("valueSpecified", readBoolean(readBuffer));

    Variant value =
        readOptionalField(
            "value",
            readComplex(() -> Variant.staticParse(readBuffer), readBuffer),
            valueSpecified);

    StatusCode statusCode =
        readOptionalField(
            "statusCode",
            readComplex(() -> StatusCode.staticParse(readBuffer), readBuffer),
            statusCodeSpecified);

    Long sourceTimestamp =
        readOptionalField(
            "sourceTimestamp", readSignedLong(readBuffer, 64), sourceTimestampSpecified);

    Integer sourcePicoseconds =
        readOptionalField(
            "sourcePicoseconds", readUnsignedInt(readBuffer, 16), sourcePicosecondsSpecified);

    Long serverTimestamp =
        readOptionalField(
            "serverTimestamp", readSignedLong(readBuffer, 64), serverTimestampSpecified);

    Integer serverPicoseconds =
        readOptionalField(
            "serverPicoseconds", readUnsignedInt(readBuffer, 16), serverPicosecondsSpecified);

    readBuffer.closeContext("DataValue");
    // Create the instance
    DataValue _dataValue;
    _dataValue =
        new DataValue(
            serverPicosecondsSpecified,
            sourcePicosecondsSpecified,
            serverTimestampSpecified,
            sourceTimestampSpecified,
            statusCodeSpecified,
            valueSpecified,
            value,
            statusCode,
            sourceTimestamp,
            sourcePicoseconds,
            serverTimestamp,
            serverPicoseconds);
    return _dataValue;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DataValue)) {
      return false;
    }
    DataValue that = (DataValue) o;
    return (getServerPicosecondsSpecified() == that.getServerPicosecondsSpecified())
        && (getSourcePicosecondsSpecified() == that.getSourcePicosecondsSpecified())
        && (getServerTimestampSpecified() == that.getServerTimestampSpecified())
        && (getSourceTimestampSpecified() == that.getSourceTimestampSpecified())
        && (getStatusCodeSpecified() == that.getStatusCodeSpecified())
        && (getValueSpecified() == that.getValueSpecified())
        && (getValue() == that.getValue())
        && (getStatusCode() == that.getStatusCode())
        && (getSourceTimestamp() == that.getSourceTimestamp())
        && (getSourcePicoseconds() == that.getSourcePicoseconds())
        && (getServerTimestamp() == that.getServerTimestamp())
        && (getServerPicoseconds() == that.getServerPicoseconds())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getServerPicosecondsSpecified(),
        getSourcePicosecondsSpecified(),
        getServerTimestampSpecified(),
        getSourceTimestampSpecified(),
        getStatusCodeSpecified(),
        getValueSpecified(),
        getValue(),
        getStatusCode(),
        getSourceTimestamp(),
        getSourcePicoseconds(),
        getServerTimestamp(),
        getServerPicoseconds());
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
