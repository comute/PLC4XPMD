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

public class TranslateBrowsePathsToNodeIdsRequest extends ExtensionObjectDefinition
    implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "554";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfBrowsePaths;
  protected final List<ExtensionObjectDefinition> browsePaths;

  public TranslateBrowsePathsToNodeIdsRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfBrowsePaths,
      List<ExtensionObjectDefinition> browsePaths) {
    super();
    this.requestHeader = requestHeader;
    this.noOfBrowsePaths = noOfBrowsePaths;
    this.browsePaths = browsePaths;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfBrowsePaths() {
    return noOfBrowsePaths;
  }

  public List<ExtensionObjectDefinition> getBrowsePaths() {
    return browsePaths;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TranslateBrowsePathsToNodeIdsRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfBrowsePaths)
    writeSimpleField("noOfBrowsePaths", noOfBrowsePaths, writeSignedInt(writeBuffer, 32));

    // Array Field (browsePaths)
    writeComplexTypeArrayField("browsePaths", browsePaths, writeBuffer);

    writeBuffer.popContext("TranslateBrowsePathsToNodeIdsRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TranslateBrowsePathsToNodeIdsRequest _value = this;

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfBrowsePaths)
    lengthInBits += 32;

    // Array field
    if (browsePaths != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : browsePaths) {
        boolean last = ++i >= browsePaths.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static TranslateBrowsePathsToNodeIdsRequestBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("TranslateBrowsePathsToNodeIdsRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfBrowsePaths = readSimpleField("noOfBrowsePaths", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> browsePaths =
        readCountArrayField(
            "browsePaths",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("545")),
                readBuffer),
            noOfBrowsePaths);

    readBuffer.closeContext("TranslateBrowsePathsToNodeIdsRequest");
    // Create the instance
    return new TranslateBrowsePathsToNodeIdsRequestBuilder(
        requestHeader, noOfBrowsePaths, browsePaths);
  }

  public static class TranslateBrowsePathsToNodeIdsRequestBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfBrowsePaths;
    private final List<ExtensionObjectDefinition> browsePaths;

    public TranslateBrowsePathsToNodeIdsRequestBuilder(
        ExtensionObjectDefinition requestHeader,
        int noOfBrowsePaths,
        List<ExtensionObjectDefinition> browsePaths) {

      this.requestHeader = requestHeader;
      this.noOfBrowsePaths = noOfBrowsePaths;
      this.browsePaths = browsePaths;
    }

    public TranslateBrowsePathsToNodeIdsRequest build() {
      TranslateBrowsePathsToNodeIdsRequest translateBrowsePathsToNodeIdsRequest =
          new TranslateBrowsePathsToNodeIdsRequest(requestHeader, noOfBrowsePaths, browsePaths);
      return translateBrowsePathsToNodeIdsRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TranslateBrowsePathsToNodeIdsRequest)) {
      return false;
    }
    TranslateBrowsePathsToNodeIdsRequest that = (TranslateBrowsePathsToNodeIdsRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfBrowsePaths() == that.getNoOfBrowsePaths())
        && (getBrowsePaths() == that.getBrowsePaths())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestHeader(), getNoOfBrowsePaths(), getBrowsePaths());
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
