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

public class SecurityDataDisplayMessage extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final String message;

  public SecurityDataDisplayMessage(
      SecurityCommandTypeContainer commandTypeContainer, byte argument, String message) {
    super(commandTypeContainer, argument);
    this.message = message;
  }

  public String getMessage() {
    return message;
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SecurityDataDisplayMessage");

    // Simple Field (message)
    writeSimpleField(
        "message",
        message,
        writeString(writeBuffer, (((commandTypeContainer.getNumBytes()) - (1))) * (8)));

    writeBuffer.popContext("SecurityDataDisplayMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataDisplayMessage _value = this;

    // Simple field (message)
    lengthInBits += (((commandTypeContainer.getNumBytes()) - (1))) * (8);

    return lengthInBits;
  }

  public static SecurityDataDisplayMessageBuilder staticParseBuilder(
      ReadBuffer readBuffer, SecurityCommandTypeContainer commandTypeContainer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataDisplayMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    String message =
        readSimpleField(
            "message",
            readString(readBuffer, (((commandTypeContainer.getNumBytes()) - (1))) * (8)));

    readBuffer.closeContext("SecurityDataDisplayMessage");
    // Create the instance
    return new SecurityDataDisplayMessageBuilder(message);
  }

  public static class SecurityDataDisplayMessageBuilder
      implements SecurityData.SecurityDataBuilder {
    private final String message;

    public SecurityDataDisplayMessageBuilder(String message) {

      this.message = message;
    }

    public SecurityDataDisplayMessage build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataDisplayMessage securityDataDisplayMessage =
          new SecurityDataDisplayMessage(commandTypeContainer, argument, message);
      return securityDataDisplayMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataDisplayMessage)) {
      return false;
    }
    SecurityDataDisplayMessage that = (SecurityDataDisplayMessage) o;
    return (getMessage() == that.getMessage()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMessage());
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
