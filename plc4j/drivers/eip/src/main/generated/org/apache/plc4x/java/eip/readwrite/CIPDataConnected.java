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
package org.apache.plc4x.java.eip.readwrite;

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

public class CIPDataConnected implements Message {

  // Properties.
  protected final long value;
  protected final int tagStatus;

  public CIPDataConnected(long value, int tagStatus) {
    super();
    this.value = value;
    this.tagStatus = tagStatus;
  }

  public long getValue() {
    return value;
  }

  public int getTagStatus() {
    return tagStatus;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CIPDataConnected");

    // Simple Field (value)
    writeSimpleField("value", value, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (tagStatus)
    writeSimpleField("tagStatus", tagStatus, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("CIPDataConnected");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CIPDataConnected _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (value)
    lengthInBits += 32;

    // Simple field (tagStatus)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static CIPDataConnected staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CIPDataConnected");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));

    int tagStatus = readSimpleField("tagStatus", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("CIPDataConnected");
    // Create the instance
    CIPDataConnected _cIPDataConnected;
    _cIPDataConnected = new CIPDataConnected(value, tagStatus);
    return _cIPDataConnected;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CIPDataConnected)) {
      return false;
    }
    CIPDataConnected that = (CIPDataConnected) o;
    return (getValue() == that.getValue()) && (getTagStatus() == that.getTagStatus()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getValue(), getTagStatus());
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
