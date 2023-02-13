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
package org.apache.plc4x.java.cbus.readwrite;

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

public class MonitoredSALReply extends EncodedReply implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final MonitoredSAL monitoredSAL;

  // Arguments.
  protected final CBusOptions cBusOptions;
  protected final RequestContext requestContext;

  public MonitoredSALReply(
      byte peekedByte,
      MonitoredSAL monitoredSAL,
      CBusOptions cBusOptions,
      RequestContext requestContext) {
    super(peekedByte, cBusOptions, requestContext);
    this.monitoredSAL = monitoredSAL;
    this.cBusOptions = cBusOptions;
    this.requestContext = requestContext;
  }

  public MonitoredSAL getMonitoredSAL() {
    return monitoredSAL;
  }

  @Override
  protected void serializeEncodedReplyChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MonitoredSALReply");

    // Simple Field (monitoredSAL)
    writeSimpleField("monitoredSAL", monitoredSAL, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("MonitoredSALReply");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoredSALReply _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (monitoredSAL)
    lengthInBits += monitoredSAL.getLengthInBits();

    return lengthInBits;
  }

  public static EncodedReplyBuilder staticParseEncodedReplyBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions, RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("MonitoredSALReply");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    MonitoredSAL monitoredSAL =
        readSimpleField(
            "monitoredSAL",
            new DataReaderComplexDefault<>(
                () -> MonitoredSAL.staticParse(readBuffer, (CBusOptions) (cBusOptions)),
                readBuffer));

    readBuffer.closeContext("MonitoredSALReply");
    // Create the instance
    return new MonitoredSALReplyBuilderImpl(monitoredSAL, cBusOptions, requestContext);
  }

  public static class MonitoredSALReplyBuilderImpl implements EncodedReply.EncodedReplyBuilder {
    private final MonitoredSAL monitoredSAL;
    private final CBusOptions cBusOptions;
    private final RequestContext requestContext;

    public MonitoredSALReplyBuilderImpl(
        MonitoredSAL monitoredSAL, CBusOptions cBusOptions, RequestContext requestContext) {
      this.monitoredSAL = monitoredSAL;
      this.cBusOptions = cBusOptions;
      this.requestContext = requestContext;
    }

    public MonitoredSALReply build(
        byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext) {
      MonitoredSALReply monitoredSALReply =
          new MonitoredSALReply(peekedByte, monitoredSAL, cBusOptions, requestContext);
      return monitoredSALReply;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredSALReply)) {
      return false;
    }
    MonitoredSALReply that = (MonitoredSALReply) o;
    return (getMonitoredSAL() == that.getMonitoredSAL()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMonitoredSAL());
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
