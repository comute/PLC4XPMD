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

public class CBusCommandPointToPoint extends CBusCommand implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final CBusPointToPointCommand command;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public CBusCommandPointToPoint(
      CBusHeader header, CBusPointToPointCommand command, CBusOptions cBusOptions) {
    super(header, cBusOptions);
    this.command = command;
    this.cBusOptions = cBusOptions;
  }

  public CBusPointToPointCommand getCommand() {
    return command;
  }

  @Override
  protected void serializeCBusCommandChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusCommandPointToPoint");

    // Simple Field (command)
    writeSimpleField("command", command, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CBusCommandPointToPoint");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusCommandPointToPoint _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (command)
    lengthInBits += command.getLengthInBits();

    return lengthInBits;
  }

  public static CBusCommandBuilder staticParseCBusCommandBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusCommandPointToPoint");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    CBusPointToPointCommand command =
        readSimpleField(
            "command",
            new DataReaderComplexDefault<>(
                () -> CBusPointToPointCommand.staticParse(readBuffer, (CBusOptions) (cBusOptions)),
                readBuffer));

    readBuffer.closeContext("CBusCommandPointToPoint");
    // Create the instance
    return new CBusCommandPointToPointBuilderImpl(command, cBusOptions);
  }

  public static class CBusCommandPointToPointBuilderImpl implements CBusCommand.CBusCommandBuilder {
    private final CBusPointToPointCommand command;
    private final CBusOptions cBusOptions;

    public CBusCommandPointToPointBuilderImpl(
        CBusPointToPointCommand command, CBusOptions cBusOptions) {
      this.command = command;
      this.cBusOptions = cBusOptions;
    }

    public CBusCommandPointToPoint build(CBusHeader header, CBusOptions cBusOptions) {
      CBusCommandPointToPoint cBusCommandPointToPoint =
          new CBusCommandPointToPoint(header, command, cBusOptions);
      return cBusCommandPointToPoint;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusCommandPointToPoint)) {
      return false;
    }
    CBusCommandPointToPoint that = (CBusCommandPointToPoint) o;
    return (getCommand() == that.getCommand()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCommand());
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
