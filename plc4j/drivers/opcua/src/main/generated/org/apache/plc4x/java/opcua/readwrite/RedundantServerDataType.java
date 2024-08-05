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

public class RedundantServerDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 855;
  }

  // Properties.
  protected final PascalString serverId;
  protected final short serviceLevel;
  protected final ServerState serverState;

  public RedundantServerDataType(
      PascalString serverId, short serviceLevel, ServerState serverState) {
    super();
    this.serverId = serverId;
    this.serviceLevel = serviceLevel;
    this.serverState = serverState;
  }

  public PascalString getServerId() {
    return serverId;
  }

  public short getServiceLevel() {
    return serviceLevel;
  }

  public ServerState getServerState() {
    return serverState;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RedundantServerDataType");

    // Simple Field (serverId)
    writeSimpleField("serverId", serverId, writeComplex(writeBuffer));

    // Simple Field (serviceLevel)
    writeSimpleField("serviceLevel", serviceLevel, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (serverState)
    writeSimpleEnumField(
        "serverState",
        "ServerState",
        serverState,
        writeEnum(ServerState::getValue, ServerState::name, writeUnsignedLong(writeBuffer, 32)));

    writeBuffer.popContext("RedundantServerDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RedundantServerDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (serverId)
    lengthInBits += serverId.getLengthInBits();

    // Simple field (serviceLevel)
    lengthInBits += 8;

    // Simple field (serverState)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("RedundantServerDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString serverId =
        readSimpleField(
            "serverId", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    short serviceLevel = readSimpleField("serviceLevel", readUnsignedShort(readBuffer, 8));

    ServerState serverState =
        readEnumField(
            "serverState",
            "ServerState",
            readEnum(ServerState::enumForValue, readUnsignedLong(readBuffer, 32)));

    readBuffer.closeContext("RedundantServerDataType");
    // Create the instance
    return new RedundantServerDataTypeBuilderImpl(serverId, serviceLevel, serverState);
  }

  public static class RedundantServerDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString serverId;
    private final short serviceLevel;
    private final ServerState serverState;

    public RedundantServerDataTypeBuilderImpl(
        PascalString serverId, short serviceLevel, ServerState serverState) {
      this.serverId = serverId;
      this.serviceLevel = serviceLevel;
      this.serverState = serverState;
    }

    public RedundantServerDataType build() {
      RedundantServerDataType redundantServerDataType =
          new RedundantServerDataType(serverId, serviceLevel, serverState);
      return redundantServerDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RedundantServerDataType)) {
      return false;
    }
    RedundantServerDataType that = (RedundantServerDataType) o;
    return (getServerId() == that.getServerId())
        && (getServiceLevel() == that.getServiceLevel())
        && (getServerState() == that.getServerState())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getServerId(), getServiceLevel(), getServerState());
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
