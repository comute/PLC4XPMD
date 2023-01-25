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

public class BVLCDeleteForeignDeviceTableEntry extends BVLC implements Message {

  // Accessors for discriminator values.
  public Short getBvlcFunction() {
    return (short) 0x08;
  }

  // Properties.
  protected final List<Short> ip;
  protected final int port;

  public BVLCDeleteForeignDeviceTableEntry(List<Short> ip, int port) {
    super();
    this.ip = ip;
    this.port = port;
  }

  public List<Short> getIp() {
    return ip;
  }

  public int getPort() {
    return port;
  }

  @Override
  protected void serializeBVLCChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BVLCDeleteForeignDeviceTableEntry");

    // Array Field (ip)
    writeSimpleTypeArrayField(
        "ip",
        ip,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (port)
    writeSimpleField(
        "port",
        port,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("BVLCDeleteForeignDeviceTableEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BVLCDeleteForeignDeviceTableEntry _value = this;

    // Array field
    if (ip != null) {
      lengthInBits += 8 * ip.size();
    }

    // Simple field (port)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static BVLCDeleteForeignDeviceTableEntryBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BVLCDeleteForeignDeviceTableEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    List<Short> ip =
        readCountArrayField(
            "ip",
            readUnsignedShort(readBuffer, 8),
            4,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int port =
        readSimpleField(
            "port",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("BVLCDeleteForeignDeviceTableEntry");
    // Create the instance
    return new BVLCDeleteForeignDeviceTableEntryBuilder(ip, port);
  }

  public static class BVLCDeleteForeignDeviceTableEntryBuilder implements BVLC.BVLCBuilder {
    private final List<Short> ip;
    private final int port;

    public BVLCDeleteForeignDeviceTableEntryBuilder(List<Short> ip, int port) {

      this.ip = ip;
      this.port = port;
    }

    public BVLCDeleteForeignDeviceTableEntry build() {
      BVLCDeleteForeignDeviceTableEntry bVLCDeleteForeignDeviceTableEntry =
          new BVLCDeleteForeignDeviceTableEntry(ip, port);
      return bVLCDeleteForeignDeviceTableEntry;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BVLCDeleteForeignDeviceTableEntry)) {
      return false;
    }
    BVLCDeleteForeignDeviceTableEntry that = (BVLCDeleteForeignDeviceTableEntry) o;
    return (getIp() == that.getIp()) && (getPort() == that.getPort()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIp(), getPort());
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
