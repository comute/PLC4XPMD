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
package org.apache.plc4x.java.opcua.readwrite;

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

public class RepublishResponse extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 835;
  }

  // Properties.
  protected final ResponseHeader responseHeader;
  protected final NotificationMessage notificationMessage;

  public RepublishResponse(ResponseHeader responseHeader, NotificationMessage notificationMessage) {
    super();
    this.responseHeader = responseHeader;
    this.notificationMessage = notificationMessage;
  }

  public ResponseHeader getResponseHeader() {
    return responseHeader;
  }

  public NotificationMessage getNotificationMessage() {
    return notificationMessage;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RepublishResponse");

    // Simple Field (responseHeader)
    writeSimpleField("responseHeader", responseHeader, writeComplex(writeBuffer));

    // Simple Field (notificationMessage)
    writeSimpleField("notificationMessage", notificationMessage, writeComplex(writeBuffer));

    writeBuffer.popContext("RepublishResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RepublishResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (responseHeader)
    lengthInBits += responseHeader.getLengthInBits();

    // Simple field (notificationMessage)
    lengthInBits += notificationMessage.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("RepublishResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ResponseHeader responseHeader =
        readSimpleField(
            "responseHeader",
            readComplex(
                () ->
                    (ResponseHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (394)),
                readBuffer));

    NotificationMessage notificationMessage =
        readSimpleField(
            "notificationMessage",
            readComplex(
                () ->
                    (NotificationMessage)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (805)),
                readBuffer));

    readBuffer.closeContext("RepublishResponse");
    // Create the instance
    return new RepublishResponseBuilderImpl(responseHeader, notificationMessage);
  }

  public static class RepublishResponseBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ResponseHeader responseHeader;
    private final NotificationMessage notificationMessage;

    public RepublishResponseBuilderImpl(
        ResponseHeader responseHeader, NotificationMessage notificationMessage) {
      this.responseHeader = responseHeader;
      this.notificationMessage = notificationMessage;
    }

    public RepublishResponse build() {
      RepublishResponse republishResponse =
          new RepublishResponse(responseHeader, notificationMessage);
      return republishResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RepublishResponse)) {
      return false;
    }
    RepublishResponse that = (RepublishResponse) o;
    return (getResponseHeader() == that.getResponseHeader())
        && (getNotificationMessage() == that.getNotificationMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getResponseHeader(), getNotificationMessage());
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
