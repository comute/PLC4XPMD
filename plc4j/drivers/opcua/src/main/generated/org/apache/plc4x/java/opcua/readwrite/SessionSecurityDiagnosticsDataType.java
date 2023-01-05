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

public class SessionSecurityDiagnosticsDataType extends ExtensionObjectDefinition
    implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "870";
  }

  // Properties.
  protected final NodeId sessionId;
  protected final PascalString clientUserIdOfSession;
  protected final int noOfClientUserIdHistory;
  protected final List<PascalString> clientUserIdHistory;
  protected final PascalString authenticationMechanism;
  protected final PascalString encoding;
  protected final PascalString transportProtocol;
  protected final MessageSecurityMode securityMode;
  protected final PascalString securityPolicyUri;
  protected final PascalByteString clientCertificate;

  public SessionSecurityDiagnosticsDataType(
      NodeId sessionId,
      PascalString clientUserIdOfSession,
      int noOfClientUserIdHistory,
      List<PascalString> clientUserIdHistory,
      PascalString authenticationMechanism,
      PascalString encoding,
      PascalString transportProtocol,
      MessageSecurityMode securityMode,
      PascalString securityPolicyUri,
      PascalByteString clientCertificate) {
    super();
    this.sessionId = sessionId;
    this.clientUserIdOfSession = clientUserIdOfSession;
    this.noOfClientUserIdHistory = noOfClientUserIdHistory;
    this.clientUserIdHistory = clientUserIdHistory;
    this.authenticationMechanism = authenticationMechanism;
    this.encoding = encoding;
    this.transportProtocol = transportProtocol;
    this.securityMode = securityMode;
    this.securityPolicyUri = securityPolicyUri;
    this.clientCertificate = clientCertificate;
  }

  public NodeId getSessionId() {
    return sessionId;
  }

  public PascalString getClientUserIdOfSession() {
    return clientUserIdOfSession;
  }

  public int getNoOfClientUserIdHistory() {
    return noOfClientUserIdHistory;
  }

  public List<PascalString> getClientUserIdHistory() {
    return clientUserIdHistory;
  }

  public PascalString getAuthenticationMechanism() {
    return authenticationMechanism;
  }

  public PascalString getEncoding() {
    return encoding;
  }

  public PascalString getTransportProtocol() {
    return transportProtocol;
  }

  public MessageSecurityMode getSecurityMode() {
    return securityMode;
  }

  public PascalString getSecurityPolicyUri() {
    return securityPolicyUri;
  }

  public PascalByteString getClientCertificate() {
    return clientCertificate;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SessionSecurityDiagnosticsDataType");

    // Simple Field (sessionId)
    writeSimpleField("sessionId", sessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (clientUserIdOfSession)
    writeSimpleField(
        "clientUserIdOfSession",
        clientUserIdOfSession,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfClientUserIdHistory)
    writeSimpleField(
        "noOfClientUserIdHistory", noOfClientUserIdHistory, writeSignedInt(writeBuffer, 32));

    // Array Field (clientUserIdHistory)
    writeComplexTypeArrayField("clientUserIdHistory", clientUserIdHistory, writeBuffer);

    // Simple Field (authenticationMechanism)
    writeSimpleField(
        "authenticationMechanism",
        authenticationMechanism,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (encoding)
    writeSimpleField("encoding", encoding, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (transportProtocol)
    writeSimpleField(
        "transportProtocol", transportProtocol, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (securityMode)
    writeSimpleEnumField(
        "securityMode",
        "MessageSecurityMode",
        securityMode,
        new DataWriterEnumDefault<>(
            MessageSecurityMode::getValue,
            MessageSecurityMode::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (securityPolicyUri)
    writeSimpleField(
        "securityPolicyUri", securityPolicyUri, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (clientCertificate)
    writeSimpleField(
        "clientCertificate", clientCertificate, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SessionSecurityDiagnosticsDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SessionSecurityDiagnosticsDataType _value = this;

    // Simple field (sessionId)
    lengthInBits += sessionId.getLengthInBits();

    // Simple field (clientUserIdOfSession)
    lengthInBits += clientUserIdOfSession.getLengthInBits();

    // Simple field (noOfClientUserIdHistory)
    lengthInBits += 32;

    // Array field
    if (clientUserIdHistory != null) {
      int i = 0;
      for (PascalString element : clientUserIdHistory) {
        boolean last = ++i >= clientUserIdHistory.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (authenticationMechanism)
    lengthInBits += authenticationMechanism.getLengthInBits();

    // Simple field (encoding)
    lengthInBits += encoding.getLengthInBits();

    // Simple field (transportProtocol)
    lengthInBits += transportProtocol.getLengthInBits();

    // Simple field (securityMode)
    lengthInBits += 32;

    // Simple field (securityPolicyUri)
    lengthInBits += securityPolicyUri.getLengthInBits();

    // Simple field (clientCertificate)
    lengthInBits += clientCertificate.getLengthInBits();

    return lengthInBits;
  }

  public static SessionSecurityDiagnosticsDataTypeBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("SessionSecurityDiagnosticsDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    NodeId sessionId =
        readSimpleField(
            "sessionId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    PascalString clientUserIdOfSession =
        readSimpleField(
            "clientUserIdOfSession",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfClientUserIdHistory =
        readSimpleField("noOfClientUserIdHistory", readSignedInt(readBuffer, 32));

    List<PascalString> clientUserIdHistory =
        readCountArrayField(
            "clientUserIdHistory",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfClientUserIdHistory);

    PascalString authenticationMechanism =
        readSimpleField(
            "authenticationMechanism",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString encoding =
        readSimpleField(
            "encoding",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString transportProtocol =
        readSimpleField(
            "transportProtocol",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    MessageSecurityMode securityMode =
        readEnumField(
            "securityMode",
            "MessageSecurityMode",
            new DataReaderEnumDefault<>(
                MessageSecurityMode::enumForValue, readUnsignedLong(readBuffer, 32)));

    PascalString securityPolicyUri =
        readSimpleField(
            "securityPolicyUri",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalByteString clientCertificate =
        readSimpleField(
            "clientCertificate",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SessionSecurityDiagnosticsDataType");
    // Create the instance
    return new SessionSecurityDiagnosticsDataTypeBuilder(
        sessionId,
        clientUserIdOfSession,
        noOfClientUserIdHistory,
        clientUserIdHistory,
        authenticationMechanism,
        encoding,
        transportProtocol,
        securityMode,
        securityPolicyUri,
        clientCertificate);
  }

  public static class SessionSecurityDiagnosticsDataTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId sessionId;
    private final PascalString clientUserIdOfSession;
    private final int noOfClientUserIdHistory;
    private final List<PascalString> clientUserIdHistory;
    private final PascalString authenticationMechanism;
    private final PascalString encoding;
    private final PascalString transportProtocol;
    private final MessageSecurityMode securityMode;
    private final PascalString securityPolicyUri;
    private final PascalByteString clientCertificate;

    public SessionSecurityDiagnosticsDataTypeBuilder(
        NodeId sessionId,
        PascalString clientUserIdOfSession,
        int noOfClientUserIdHistory,
        List<PascalString> clientUserIdHistory,
        PascalString authenticationMechanism,
        PascalString encoding,
        PascalString transportProtocol,
        MessageSecurityMode securityMode,
        PascalString securityPolicyUri,
        PascalByteString clientCertificate) {

      this.sessionId = sessionId;
      this.clientUserIdOfSession = clientUserIdOfSession;
      this.noOfClientUserIdHistory = noOfClientUserIdHistory;
      this.clientUserIdHistory = clientUserIdHistory;
      this.authenticationMechanism = authenticationMechanism;
      this.encoding = encoding;
      this.transportProtocol = transportProtocol;
      this.securityMode = securityMode;
      this.securityPolicyUri = securityPolicyUri;
      this.clientCertificate = clientCertificate;
    }

    public SessionSecurityDiagnosticsDataType build() {
      SessionSecurityDiagnosticsDataType sessionSecurityDiagnosticsDataType =
          new SessionSecurityDiagnosticsDataType(
              sessionId,
              clientUserIdOfSession,
              noOfClientUserIdHistory,
              clientUserIdHistory,
              authenticationMechanism,
              encoding,
              transportProtocol,
              securityMode,
              securityPolicyUri,
              clientCertificate);
      return sessionSecurityDiagnosticsDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SessionSecurityDiagnosticsDataType)) {
      return false;
    }
    SessionSecurityDiagnosticsDataType that = (SessionSecurityDiagnosticsDataType) o;
    return (getSessionId() == that.getSessionId())
        && (getClientUserIdOfSession() == that.getClientUserIdOfSession())
        && (getNoOfClientUserIdHistory() == that.getNoOfClientUserIdHistory())
        && (getClientUserIdHistory() == that.getClientUserIdHistory())
        && (getAuthenticationMechanism() == that.getAuthenticationMechanism())
        && (getEncoding() == that.getEncoding())
        && (getTransportProtocol() == that.getTransportProtocol())
        && (getSecurityMode() == that.getSecurityMode())
        && (getSecurityPolicyUri() == that.getSecurityPolicyUri())
        && (getClientCertificate() == that.getClientCertificate())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSessionId(),
        getClientUserIdOfSession(),
        getNoOfClientUserIdHistory(),
        getClientUserIdHistory(),
        getAuthenticationMechanism(),
        getEncoding(),
        getTransportProtocol(),
        getSecurityMode(),
        getSecurityPolicyUri(),
        getClientCertificate());
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
