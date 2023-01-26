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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageApplicationCommunicationStop extends OpenProtocolMessage
    implements Message {

  // Accessors for discriminator values.
  public Mid getMid() {
    return Mid.ApplicationCommunicationStop;
  }

  // Arguments.
  protected final OpenProtocolRevision connectionRevision;

  public OpenProtocolMessageApplicationCommunicationStop(
      OpenProtocolRevision selectedRevision,
      Short noAckFlag,
      Integer stationId,
      Integer spindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      OpenProtocolRevision connectionRevision) {
    super(
        selectedRevision,
        noAckFlag,
        stationId,
        spindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber,
        connectionRevision);
    this.connectionRevision = connectionRevision;
  }

  @Override
  protected void serializeOpenProtocolMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageApplicationCommunicationStop");

    writeBuffer.popContext("OpenProtocolMessageApplicationCommunicationStop");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageApplicationCommunicationStop _value = this;

    return lengthInBits;
  }

  public static OpenProtocolMessageApplicationCommunicationStopBuilder staticParseBuilder(
      ReadBuffer readBuffer, OpenProtocolRevision connectionRevision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageApplicationCommunicationStop");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    readBuffer.closeContext("OpenProtocolMessageApplicationCommunicationStop");
    // Create the instance
    return new OpenProtocolMessageApplicationCommunicationStopBuilder(connectionRevision);
  }

  public static class OpenProtocolMessageApplicationCommunicationStopBuilder
      implements OpenProtocolMessage.OpenProtocolMessageBuilder {
    private final OpenProtocolRevision connectionRevision;

    public OpenProtocolMessageApplicationCommunicationStopBuilder(
        OpenProtocolRevision connectionRevision) {

      this.connectionRevision = connectionRevision;
    }

    public OpenProtocolMessageApplicationCommunicationStop build(
        OpenProtocolRevision selectedRevision,
        Short noAckFlag,
        Integer stationId,
        Integer spindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber,
        OpenProtocolRevision connectionRevision) {
      OpenProtocolMessageApplicationCommunicationStop
          openProtocolMessageApplicationCommunicationStop =
              new OpenProtocolMessageApplicationCommunicationStop(
                  selectedRevision,
                  noAckFlag,
                  stationId,
                  spindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  connectionRevision);
      return openProtocolMessageApplicationCommunicationStop;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageApplicationCommunicationStop)) {
      return false;
    }
    OpenProtocolMessageApplicationCommunicationStop that =
        (OpenProtocolMessageApplicationCommunicationStop) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
