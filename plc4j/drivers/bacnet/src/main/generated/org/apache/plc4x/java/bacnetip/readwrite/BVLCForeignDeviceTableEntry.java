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

public class BVLCForeignDeviceTableEntry implements Message {

  // Properties.
  protected final List<Short> ip;
  protected final int port;
  protected final int ttl;
  protected final int secondRemainingBeforePurge;

  public BVLCForeignDeviceTableEntry(
      List<Short> ip, int port, int ttl, int secondRemainingBeforePurge) {
    super();
    this.ip = ip;
    this.port = port;
    this.ttl = ttl;
    this.secondRemainingBeforePurge = secondRemainingBeforePurge;
  }

  public List<Short> getIp() {
    return ip;
  }

  public int getPort() {
    return port;
  }

  public int getTtl() {
    return ttl;
  }

  public int getSecondRemainingBeforePurge() {
    return secondRemainingBeforePurge;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BVLCForeignDeviceTableEntry");

    // Array Field (ip)
    writeSimpleTypeArrayField("ip", ip, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (port)
    writeSimpleField("port", port, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (ttl)
    writeSimpleField("ttl", ttl, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (secondRemainingBeforePurge)
    writeSimpleField(
        "secondRemainingBeforePurge",
        secondRemainingBeforePurge,
        writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("BVLCForeignDeviceTableEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BVLCForeignDeviceTableEntry _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (ip != null) {
      lengthInBits += 8 * ip.size();
    }

    // Simple field (port)
    lengthInBits += 16;

    // Simple field (ttl)
    lengthInBits += 16;

    // Simple field (secondRemainingBeforePurge)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static BVLCForeignDeviceTableEntry staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BVLCForeignDeviceTableEntry staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BVLCForeignDeviceTableEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    List<Short> ip = readCountArrayField("ip", readUnsignedShort(readBuffer, 8), 4);

    int port = readSimpleField("port", readUnsignedInt(readBuffer, 16));

    int ttl = readSimpleField("ttl", readUnsignedInt(readBuffer, 16));

    int secondRemainingBeforePurge =
        readSimpleField("secondRemainingBeforePurge", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("BVLCForeignDeviceTableEntry");
    // Create the instance
    BVLCForeignDeviceTableEntry _bVLCForeignDeviceTableEntry;
    _bVLCForeignDeviceTableEntry =
        new BVLCForeignDeviceTableEntry(ip, port, ttl, secondRemainingBeforePurge);
    return _bVLCForeignDeviceTableEntry;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BVLCForeignDeviceTableEntry)) {
      return false;
    }
    BVLCForeignDeviceTableEntry that = (BVLCForeignDeviceTableEntry) o;
    return (getIp() == that.getIp())
        && (getPort() == that.getPort())
        && (getTtl() == that.getTtl())
        && (getSecondRemainingBeforePurge() == that.getSecondRemainingBeforePurge())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getIp(), getPort(), getTtl(), getSecondRemainingBeforePurge());
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
