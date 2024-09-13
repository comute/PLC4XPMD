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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ConnectionStateRequest extends KnxNetIpMessage implements Message {

  // Accessors for discriminator values.
  public Integer getMsgType() {
    return (int) 0x0207;
  }

  // Properties.
  protected final short communicationChannelId;
  protected final HPAIControlEndpoint hpaiControlEndpoint;

  public ConnectionStateRequest(
      short communicationChannelId, HPAIControlEndpoint hpaiControlEndpoint) {
    super();
    this.communicationChannelId = communicationChannelId;
    this.hpaiControlEndpoint = hpaiControlEndpoint;
  }

  public short getCommunicationChannelId() {
    return communicationChannelId;
  }

  public HPAIControlEndpoint getHpaiControlEndpoint() {
    return hpaiControlEndpoint;
  }

  @Override
  protected void serializeKnxNetIpMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ConnectionStateRequest");

    // Simple Field (communicationChannelId)
    writeSimpleField(
        "communicationChannelId",
        communicationChannelId,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (hpaiControlEndpoint)
    writeSimpleField(
        "hpaiControlEndpoint",
        hpaiControlEndpoint,
        writeComplex(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("ConnectionStateRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConnectionStateRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (communicationChannelId)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (hpaiControlEndpoint)
    lengthInBits += hpaiControlEndpoint.getLengthInBits();

    return lengthInBits;
  }

  public static KnxNetIpMessageBuilder staticParseKnxNetIpMessageBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("ConnectionStateRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short communicationChannelId =
        readSimpleField(
            "communicationChannelId",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Short reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedShort(readBuffer, 8),
            (short) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    HPAIControlEndpoint hpaiControlEndpoint =
        readSimpleField(
            "hpaiControlEndpoint",
            readComplex(() -> HPAIControlEndpoint.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("ConnectionStateRequest");
    // Create the instance
    return new ConnectionStateRequestBuilderImpl(communicationChannelId, hpaiControlEndpoint);
  }

  public static class ConnectionStateRequestBuilderImpl
      implements KnxNetIpMessage.KnxNetIpMessageBuilder {
    private final short communicationChannelId;
    private final HPAIControlEndpoint hpaiControlEndpoint;

    public ConnectionStateRequestBuilderImpl(
        short communicationChannelId, HPAIControlEndpoint hpaiControlEndpoint) {
      this.communicationChannelId = communicationChannelId;
      this.hpaiControlEndpoint = hpaiControlEndpoint;
    }

    public ConnectionStateRequest build() {
      ConnectionStateRequest connectionStateRequest =
          new ConnectionStateRequest(communicationChannelId, hpaiControlEndpoint);
      return connectionStateRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConnectionStateRequest)) {
      return false;
    }
    ConnectionStateRequest that = (ConnectionStateRequest) o;
    return (getCommunicationChannelId() == that.getCommunicationChannelId())
        && (getHpaiControlEndpoint() == that.getHpaiControlEndpoint())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCommunicationChannelId(), getHpaiControlEndpoint());
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
