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

public class ProgramDiagnosticDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "896";
  }

  // Properties.
  protected final NodeId createSessionId;
  protected final PascalString createClientName;
  protected final long invocationCreationTime;
  protected final long lastTransitionTime;
  protected final PascalString lastMethodCall;
  protected final NodeId lastMethodSessionId;
  protected final int noOfLastMethodInputArguments;
  protected final List<ExtensionObjectDefinition> lastMethodInputArguments;
  protected final int noOfLastMethodOutputArguments;
  protected final List<ExtensionObjectDefinition> lastMethodOutputArguments;
  protected final long lastMethodCallTime;
  protected final ExtensionObjectDefinition lastMethodReturnStatus;

  public ProgramDiagnosticDataType(
      NodeId createSessionId,
      PascalString createClientName,
      long invocationCreationTime,
      long lastTransitionTime,
      PascalString lastMethodCall,
      NodeId lastMethodSessionId,
      int noOfLastMethodInputArguments,
      List<ExtensionObjectDefinition> lastMethodInputArguments,
      int noOfLastMethodOutputArguments,
      List<ExtensionObjectDefinition> lastMethodOutputArguments,
      long lastMethodCallTime,
      ExtensionObjectDefinition lastMethodReturnStatus) {
    super();
    this.createSessionId = createSessionId;
    this.createClientName = createClientName;
    this.invocationCreationTime = invocationCreationTime;
    this.lastTransitionTime = lastTransitionTime;
    this.lastMethodCall = lastMethodCall;
    this.lastMethodSessionId = lastMethodSessionId;
    this.noOfLastMethodInputArguments = noOfLastMethodInputArguments;
    this.lastMethodInputArguments = lastMethodInputArguments;
    this.noOfLastMethodOutputArguments = noOfLastMethodOutputArguments;
    this.lastMethodOutputArguments = lastMethodOutputArguments;
    this.lastMethodCallTime = lastMethodCallTime;
    this.lastMethodReturnStatus = lastMethodReturnStatus;
  }

  public NodeId getCreateSessionId() {
    return createSessionId;
  }

  public PascalString getCreateClientName() {
    return createClientName;
  }

  public long getInvocationCreationTime() {
    return invocationCreationTime;
  }

  public long getLastTransitionTime() {
    return lastTransitionTime;
  }

  public PascalString getLastMethodCall() {
    return lastMethodCall;
  }

  public NodeId getLastMethodSessionId() {
    return lastMethodSessionId;
  }

  public int getNoOfLastMethodInputArguments() {
    return noOfLastMethodInputArguments;
  }

  public List<ExtensionObjectDefinition> getLastMethodInputArguments() {
    return lastMethodInputArguments;
  }

  public int getNoOfLastMethodOutputArguments() {
    return noOfLastMethodOutputArguments;
  }

  public List<ExtensionObjectDefinition> getLastMethodOutputArguments() {
    return lastMethodOutputArguments;
  }

  public long getLastMethodCallTime() {
    return lastMethodCallTime;
  }

  public ExtensionObjectDefinition getLastMethodReturnStatus() {
    return lastMethodReturnStatus;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ProgramDiagnosticDataType");

    // Simple Field (createSessionId)
    writeSimpleField(
        "createSessionId", createSessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (createClientName)
    writeSimpleField(
        "createClientName", createClientName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (invocationCreationTime)
    writeSimpleField(
        "invocationCreationTime", invocationCreationTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (lastTransitionTime)
    writeSimpleField("lastTransitionTime", lastTransitionTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (lastMethodCall)
    writeSimpleField("lastMethodCall", lastMethodCall, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (lastMethodSessionId)
    writeSimpleField(
        "lastMethodSessionId", lastMethodSessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfLastMethodInputArguments)
    writeSimpleField(
        "noOfLastMethodInputArguments",
        noOfLastMethodInputArguments,
        writeSignedInt(writeBuffer, 32));

    // Array Field (lastMethodInputArguments)
    writeComplexTypeArrayField("lastMethodInputArguments", lastMethodInputArguments, writeBuffer);

    // Simple Field (noOfLastMethodOutputArguments)
    writeSimpleField(
        "noOfLastMethodOutputArguments",
        noOfLastMethodOutputArguments,
        writeSignedInt(writeBuffer, 32));

    // Array Field (lastMethodOutputArguments)
    writeComplexTypeArrayField("lastMethodOutputArguments", lastMethodOutputArguments, writeBuffer);

    // Simple Field (lastMethodCallTime)
    writeSimpleField("lastMethodCallTime", lastMethodCallTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (lastMethodReturnStatus)
    writeSimpleField(
        "lastMethodReturnStatus",
        lastMethodReturnStatus,
        new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ProgramDiagnosticDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ProgramDiagnosticDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (createSessionId)
    lengthInBits += createSessionId.getLengthInBits();

    // Simple field (createClientName)
    lengthInBits += createClientName.getLengthInBits();

    // Simple field (invocationCreationTime)
    lengthInBits += 64;

    // Simple field (lastTransitionTime)
    lengthInBits += 64;

    // Simple field (lastMethodCall)
    lengthInBits += lastMethodCall.getLengthInBits();

    // Simple field (lastMethodSessionId)
    lengthInBits += lastMethodSessionId.getLengthInBits();

    // Simple field (noOfLastMethodInputArguments)
    lengthInBits += 32;

    // Array field
    if (lastMethodInputArguments != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : lastMethodInputArguments) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= lastMethodInputArguments.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfLastMethodOutputArguments)
    lengthInBits += 32;

    // Array field
    if (lastMethodOutputArguments != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : lastMethodOutputArguments) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= lastMethodOutputArguments.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (lastMethodCallTime)
    lengthInBits += 64;

    // Simple field (lastMethodReturnStatus)
    lengthInBits += lastMethodReturnStatus.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ProgramDiagnosticDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId createSessionId =
        readSimpleField(
            "createSessionId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    PascalString createClientName =
        readSimpleField(
            "createClientName",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    long invocationCreationTime =
        readSimpleField("invocationCreationTime", readSignedLong(readBuffer, 64));

    long lastTransitionTime = readSimpleField("lastTransitionTime", readSignedLong(readBuffer, 64));

    PascalString lastMethodCall =
        readSimpleField(
            "lastMethodCall",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    NodeId lastMethodSessionId =
        readSimpleField(
            "lastMethodSessionId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    int noOfLastMethodInputArguments =
        readSimpleField("noOfLastMethodInputArguments", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> lastMethodInputArguments =
        readCountArrayField(
            "lastMethodInputArguments",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("298")),
                readBuffer),
            noOfLastMethodInputArguments);

    int noOfLastMethodOutputArguments =
        readSimpleField("noOfLastMethodOutputArguments", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> lastMethodOutputArguments =
        readCountArrayField(
            "lastMethodOutputArguments",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("298")),
                readBuffer),
            noOfLastMethodOutputArguments);

    long lastMethodCallTime = readSimpleField("lastMethodCallTime", readSignedLong(readBuffer, 64));

    ExtensionObjectDefinition lastMethodReturnStatus =
        readSimpleField(
            "lastMethodReturnStatus",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("301")),
                readBuffer));

    readBuffer.closeContext("ProgramDiagnosticDataType");
    // Create the instance
    return new ProgramDiagnosticDataTypeBuilderImpl(
        createSessionId,
        createClientName,
        invocationCreationTime,
        lastTransitionTime,
        lastMethodCall,
        lastMethodSessionId,
        noOfLastMethodInputArguments,
        lastMethodInputArguments,
        noOfLastMethodOutputArguments,
        lastMethodOutputArguments,
        lastMethodCallTime,
        lastMethodReturnStatus);
  }

  public static class ProgramDiagnosticDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId createSessionId;
    private final PascalString createClientName;
    private final long invocationCreationTime;
    private final long lastTransitionTime;
    private final PascalString lastMethodCall;
    private final NodeId lastMethodSessionId;
    private final int noOfLastMethodInputArguments;
    private final List<ExtensionObjectDefinition> lastMethodInputArguments;
    private final int noOfLastMethodOutputArguments;
    private final List<ExtensionObjectDefinition> lastMethodOutputArguments;
    private final long lastMethodCallTime;
    private final ExtensionObjectDefinition lastMethodReturnStatus;

    public ProgramDiagnosticDataTypeBuilderImpl(
        NodeId createSessionId,
        PascalString createClientName,
        long invocationCreationTime,
        long lastTransitionTime,
        PascalString lastMethodCall,
        NodeId lastMethodSessionId,
        int noOfLastMethodInputArguments,
        List<ExtensionObjectDefinition> lastMethodInputArguments,
        int noOfLastMethodOutputArguments,
        List<ExtensionObjectDefinition> lastMethodOutputArguments,
        long lastMethodCallTime,
        ExtensionObjectDefinition lastMethodReturnStatus) {
      this.createSessionId = createSessionId;
      this.createClientName = createClientName;
      this.invocationCreationTime = invocationCreationTime;
      this.lastTransitionTime = lastTransitionTime;
      this.lastMethodCall = lastMethodCall;
      this.lastMethodSessionId = lastMethodSessionId;
      this.noOfLastMethodInputArguments = noOfLastMethodInputArguments;
      this.lastMethodInputArguments = lastMethodInputArguments;
      this.noOfLastMethodOutputArguments = noOfLastMethodOutputArguments;
      this.lastMethodOutputArguments = lastMethodOutputArguments;
      this.lastMethodCallTime = lastMethodCallTime;
      this.lastMethodReturnStatus = lastMethodReturnStatus;
    }

    public ProgramDiagnosticDataType build() {
      ProgramDiagnosticDataType programDiagnosticDataType =
          new ProgramDiagnosticDataType(
              createSessionId,
              createClientName,
              invocationCreationTime,
              lastTransitionTime,
              lastMethodCall,
              lastMethodSessionId,
              noOfLastMethodInputArguments,
              lastMethodInputArguments,
              noOfLastMethodOutputArguments,
              lastMethodOutputArguments,
              lastMethodCallTime,
              lastMethodReturnStatus);
      return programDiagnosticDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ProgramDiagnosticDataType)) {
      return false;
    }
    ProgramDiagnosticDataType that = (ProgramDiagnosticDataType) o;
    return (getCreateSessionId() == that.getCreateSessionId())
        && (getCreateClientName() == that.getCreateClientName())
        && (getInvocationCreationTime() == that.getInvocationCreationTime())
        && (getLastTransitionTime() == that.getLastTransitionTime())
        && (getLastMethodCall() == that.getLastMethodCall())
        && (getLastMethodSessionId() == that.getLastMethodSessionId())
        && (getNoOfLastMethodInputArguments() == that.getNoOfLastMethodInputArguments())
        && (getLastMethodInputArguments() == that.getLastMethodInputArguments())
        && (getNoOfLastMethodOutputArguments() == that.getNoOfLastMethodOutputArguments())
        && (getLastMethodOutputArguments() == that.getLastMethodOutputArguments())
        && (getLastMethodCallTime() == that.getLastMethodCallTime())
        && (getLastMethodReturnStatus() == that.getLastMethodReturnStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getCreateSessionId(),
        getCreateClientName(),
        getInvocationCreationTime(),
        getLastTransitionTime(),
        getLastMethodCall(),
        getLastMethodSessionId(),
        getNoOfLastMethodInputArguments(),
        getLastMethodInputArguments(),
        getNoOfLastMethodOutputArguments(),
        getLastMethodOutputArguments(),
        getLastMethodCallTime(),
        getLastMethodReturnStatus());
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
