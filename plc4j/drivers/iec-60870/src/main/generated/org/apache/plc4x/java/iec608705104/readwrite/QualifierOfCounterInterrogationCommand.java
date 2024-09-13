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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class QualifierOfCounterInterrogationCommand implements Message {

  // Properties.
  protected final byte freeze;
  protected final byte request;

  public QualifierOfCounterInterrogationCommand(byte freeze, byte request) {
    super();
    this.freeze = freeze;
    this.request = request;
  }

  public byte getFreeze() {
    return freeze;
  }

  public byte getRequest() {
    return request;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QualifierOfCounterInterrogationCommand");

    // Simple Field (freeze)
    writeSimpleField(
        "freeze",
        freeze,
        writeUnsignedByte(writeBuffer, 2),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (request)
    writeSimpleField(
        "request",
        request,
        writeUnsignedByte(writeBuffer, 6),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("QualifierOfCounterInterrogationCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    QualifierOfCounterInterrogationCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (freeze)
    lengthInBits += 2;

    // Simple field (request)
    lengthInBits += 6;

    return lengthInBits;
  }

  public static QualifierOfCounterInterrogationCommand staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("QualifierOfCounterInterrogationCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte freeze =
        readSimpleField(
            "freeze",
            readUnsignedByte(readBuffer, 2),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte request =
        readSimpleField(
            "request",
            readUnsignedByte(readBuffer, 6),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("QualifierOfCounterInterrogationCommand");
    // Create the instance
    QualifierOfCounterInterrogationCommand _qualifierOfCounterInterrogationCommand;
    _qualifierOfCounterInterrogationCommand =
        new QualifierOfCounterInterrogationCommand(freeze, request);
    return _qualifierOfCounterInterrogationCommand;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QualifierOfCounterInterrogationCommand)) {
      return false;
    }
    QualifierOfCounterInterrogationCommand that = (QualifierOfCounterInterrogationCommand) o;
    return (getFreeze() == that.getFreeze()) && (getRequest() == that.getRequest()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getFreeze(), getRequest());
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
