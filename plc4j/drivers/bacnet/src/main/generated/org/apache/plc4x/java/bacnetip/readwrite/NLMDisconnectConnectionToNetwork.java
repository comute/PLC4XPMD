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

public class NLMDisconnectConnectionToNetwork extends NLM implements Message {

  // Accessors for discriminator values.
  public Short getMessageType() {
    return (short) 0x09;
  }

  // Properties.
  protected final int destinationNetworkAddress;

  // Arguments.
  protected final Integer apduLength;

  public NLMDisconnectConnectionToNetwork(int destinationNetworkAddress, Integer apduLength) {
    super(apduLength);
    this.destinationNetworkAddress = destinationNetworkAddress;
    this.apduLength = apduLength;
  }

  public int getDestinationNetworkAddress() {
    return destinationNetworkAddress;
  }

  @Override
  protected void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NLMDisconnectConnectionToNetwork");

    // Simple Field (destinationNetworkAddress)
    writeSimpleField(
        "destinationNetworkAddress", destinationNetworkAddress, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("NLMDisconnectConnectionToNetwork");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NLMDisconnectConnectionToNetwork _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (destinationNetworkAddress)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static NLMBuilder staticParseNLMBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("NLMDisconnectConnectionToNetwork");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int destinationNetworkAddress =
        readSimpleField("destinationNetworkAddress", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("NLMDisconnectConnectionToNetwork");
    // Create the instance
    return new NLMDisconnectConnectionToNetworkBuilderImpl(destinationNetworkAddress, apduLength);
  }

  public static class NLMDisconnectConnectionToNetworkBuilderImpl implements NLM.NLMBuilder {
    private final int destinationNetworkAddress;
    private final Integer apduLength;

    public NLMDisconnectConnectionToNetworkBuilderImpl(
        int destinationNetworkAddress, Integer apduLength) {
      this.destinationNetworkAddress = destinationNetworkAddress;
      this.apduLength = apduLength;
    }

    public NLMDisconnectConnectionToNetwork build(Integer apduLength) {

      NLMDisconnectConnectionToNetwork nLMDisconnectConnectionToNetwork =
          new NLMDisconnectConnectionToNetwork(destinationNetworkAddress, apduLength);
      return nLMDisconnectConnectionToNetwork;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMDisconnectConnectionToNetwork)) {
      return false;
    }
    NLMDisconnectConnectionToNetwork that = (NLMDisconnectConnectionToNetwork) o;
    return (getDestinationNetworkAddress() == that.getDestinationNetworkAddress())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDestinationNetworkAddress());
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
