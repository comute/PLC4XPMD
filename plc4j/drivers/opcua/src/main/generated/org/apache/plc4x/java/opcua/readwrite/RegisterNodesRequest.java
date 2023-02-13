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

public class RegisterNodesRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "560";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfNodesToRegister;
  protected final List<NodeId> nodesToRegister;

  public RegisterNodesRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfNodesToRegister,
      List<NodeId> nodesToRegister) {
    super();
    this.requestHeader = requestHeader;
    this.noOfNodesToRegister = noOfNodesToRegister;
    this.nodesToRegister = nodesToRegister;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfNodesToRegister() {
    return noOfNodesToRegister;
  }

  public List<NodeId> getNodesToRegister() {
    return nodesToRegister;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("RegisterNodesRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfNodesToRegister)
    writeSimpleField("noOfNodesToRegister", noOfNodesToRegister, writeSignedInt(writeBuffer, 32));

    // Array Field (nodesToRegister)
    writeComplexTypeArrayField("nodesToRegister", nodesToRegister, writeBuffer);

    writeBuffer.popContext("RegisterNodesRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RegisterNodesRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfNodesToRegister)
    lengthInBits += 32;

    // Array field
    if (nodesToRegister != null) {
      int i = 0;
      for (NodeId element : nodesToRegister) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= nodesToRegister.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("RegisterNodesRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfNodesToRegister = readSimpleField("noOfNodesToRegister", readSignedInt(readBuffer, 32));

    List<NodeId> nodesToRegister =
        readCountArrayField(
            "nodesToRegister",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer),
            noOfNodesToRegister);

    readBuffer.closeContext("RegisterNodesRequest");
    // Create the instance
    return new RegisterNodesRequestBuilderImpl(requestHeader, noOfNodesToRegister, nodesToRegister);
  }

  public static class RegisterNodesRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfNodesToRegister;
    private final List<NodeId> nodesToRegister;

    public RegisterNodesRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        int noOfNodesToRegister,
        List<NodeId> nodesToRegister) {
      this.requestHeader = requestHeader;
      this.noOfNodesToRegister = noOfNodesToRegister;
      this.nodesToRegister = nodesToRegister;
    }

    public RegisterNodesRequest build() {
      RegisterNodesRequest registerNodesRequest =
          new RegisterNodesRequest(requestHeader, noOfNodesToRegister, nodesToRegister);
      return registerNodesRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RegisterNodesRequest)) {
      return false;
    }
    RegisterNodesRequest that = (RegisterNodesRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfNodesToRegister() == that.getNoOfNodesToRegister())
        && (getNodesToRegister() == that.getNodesToRegister())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestHeader(), getNoOfNodesToRegister(), getNodesToRegister());
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
