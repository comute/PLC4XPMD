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

public class SessionlessInvokeRequestType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15903;
  }

  // Properties.
  protected final long urisVersion;
  protected final List<PascalString> namespaceUris;
  protected final List<PascalString> serverUris;
  protected final List<PascalString> localeIds;
  protected final long serviceId;

  public SessionlessInvokeRequestType(
      long urisVersion,
      List<PascalString> namespaceUris,
      List<PascalString> serverUris,
      List<PascalString> localeIds,
      long serviceId) {
    super();
    this.urisVersion = urisVersion;
    this.namespaceUris = namespaceUris;
    this.serverUris = serverUris;
    this.localeIds = localeIds;
    this.serviceId = serviceId;
  }

  public long getUrisVersion() {
    return urisVersion;
  }

  public List<PascalString> getNamespaceUris() {
    return namespaceUris;
  }

  public List<PascalString> getServerUris() {
    return serverUris;
  }

  public List<PascalString> getLocaleIds() {
    return localeIds;
  }

  public long getServiceId() {
    return serviceId;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SessionlessInvokeRequestType");

    // Simple Field (urisVersion)
    writeSimpleField("urisVersion", urisVersion, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (noOfNamespaceUris) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfNamespaceUris =
        (int) ((((getNamespaceUris()) == (null)) ? -(1) : COUNT(getNamespaceUris())));
    writeImplicitField("noOfNamespaceUris", noOfNamespaceUris, writeSignedInt(writeBuffer, 32));

    // Array Field (namespaceUris)
    writeComplexTypeArrayField("namespaceUris", namespaceUris, writeBuffer);

    // Implicit Field (noOfServerUris) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfServerUris = (int) ((((getServerUris()) == (null)) ? -(1) : COUNT(getServerUris())));
    writeImplicitField("noOfServerUris", noOfServerUris, writeSignedInt(writeBuffer, 32));

    // Array Field (serverUris)
    writeComplexTypeArrayField("serverUris", serverUris, writeBuffer);

    // Implicit Field (noOfLocaleIds) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfLocaleIds = (int) ((((getLocaleIds()) == (null)) ? -(1) : COUNT(getLocaleIds())));
    writeImplicitField("noOfLocaleIds", noOfLocaleIds, writeSignedInt(writeBuffer, 32));

    // Array Field (localeIds)
    writeComplexTypeArrayField("localeIds", localeIds, writeBuffer);

    // Simple Field (serviceId)
    writeSimpleField("serviceId", serviceId, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("SessionlessInvokeRequestType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SessionlessInvokeRequestType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (urisVersion)
    lengthInBits += 32;

    // Implicit Field (noOfNamespaceUris)
    lengthInBits += 32;

    // Array field
    if (namespaceUris != null) {
      int i = 0;
      for (PascalString element : namespaceUris) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= namespaceUris.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfServerUris)
    lengthInBits += 32;

    // Array field
    if (serverUris != null) {
      int i = 0;
      for (PascalString element : serverUris) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverUris.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfLocaleIds)
    lengthInBits += 32;

    // Array field
    if (localeIds != null) {
      int i = 0;
      for (PascalString element : localeIds) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= localeIds.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (serviceId)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("SessionlessInvokeRequestType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long urisVersion = readSimpleField("urisVersion", readUnsignedLong(readBuffer, 32));

    int noOfNamespaceUris = readImplicitField("noOfNamespaceUris", readSignedInt(readBuffer, 32));

    List<PascalString> namespaceUris =
        readCountArrayField(
            "namespaceUris",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfNamespaceUris);

    int noOfServerUris = readImplicitField("noOfServerUris", readSignedInt(readBuffer, 32));

    List<PascalString> serverUris =
        readCountArrayField(
            "serverUris",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerUris);

    int noOfLocaleIds = readImplicitField("noOfLocaleIds", readSignedInt(readBuffer, 32));

    List<PascalString> localeIds =
        readCountArrayField(
            "localeIds",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfLocaleIds);

    long serviceId = readSimpleField("serviceId", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("SessionlessInvokeRequestType");
    // Create the instance
    return new SessionlessInvokeRequestTypeBuilderImpl(
        urisVersion, namespaceUris, serverUris, localeIds, serviceId);
  }

  public static class SessionlessInvokeRequestTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long urisVersion;
    private final List<PascalString> namespaceUris;
    private final List<PascalString> serverUris;
    private final List<PascalString> localeIds;
    private final long serviceId;

    public SessionlessInvokeRequestTypeBuilderImpl(
        long urisVersion,
        List<PascalString> namespaceUris,
        List<PascalString> serverUris,
        List<PascalString> localeIds,
        long serviceId) {
      this.urisVersion = urisVersion;
      this.namespaceUris = namespaceUris;
      this.serverUris = serverUris;
      this.localeIds = localeIds;
      this.serviceId = serviceId;
    }

    public SessionlessInvokeRequestType build() {
      SessionlessInvokeRequestType sessionlessInvokeRequestType =
          new SessionlessInvokeRequestType(
              urisVersion, namespaceUris, serverUris, localeIds, serviceId);
      return sessionlessInvokeRequestType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SessionlessInvokeRequestType)) {
      return false;
    }
    SessionlessInvokeRequestType that = (SessionlessInvokeRequestType) o;
    return (getUrisVersion() == that.getUrisVersion())
        && (getNamespaceUris() == that.getNamespaceUris())
        && (getServerUris() == that.getServerUris())
        && (getLocaleIds() == that.getLocaleIds())
        && (getServiceId() == that.getServiceId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getUrisVersion(),
        getNamespaceUris(),
        getServerUris(),
        getLocaleIds(),
        getServiceId());
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
