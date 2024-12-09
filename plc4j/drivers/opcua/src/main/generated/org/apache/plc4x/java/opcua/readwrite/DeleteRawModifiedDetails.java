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

public class DeleteRawModifiedDetails extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 688;
  }

  // Properties.
  protected final NodeId nodeId;
  protected final boolean isDeleteModified;
  protected final long startTime;
  protected final long endTime;

  public DeleteRawModifiedDetails(
      NodeId nodeId, boolean isDeleteModified, long startTime, long endTime) {
    super();
    this.nodeId = nodeId;
    this.isDeleteModified = isDeleteModified;
    this.startTime = startTime;
    this.endTime = endTime;
  }

  public NodeId getNodeId() {
    return nodeId;
  }

  public boolean getIsDeleteModified() {
    return isDeleteModified;
  }

  public long getStartTime() {
    return startTime;
  }

  public long getEndTime() {
    return endTime;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DeleteRawModifiedDetails");

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (isDeleteModified)
    writeSimpleField("isDeleteModified", isDeleteModified, writeBoolean(writeBuffer));

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (endTime)
    writeSimpleField("endTime", endTime, writeSignedLong(writeBuffer, 64));

    writeBuffer.popContext("DeleteRawModifiedDetails");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DeleteRawModifiedDetails _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (isDeleteModified)
    lengthInBits += 1;

    // Simple field (startTime)
    lengthInBits += 64;

    // Simple field (endTime)
    lengthInBits += 64;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DeleteRawModifiedDetails");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId nodeId =
        readSimpleField("nodeId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean isDeleteModified = readSimpleField("isDeleteModified", readBoolean(readBuffer));

    long startTime = readSimpleField("startTime", readSignedLong(readBuffer, 64));

    long endTime = readSimpleField("endTime", readSignedLong(readBuffer, 64));

    readBuffer.closeContext("DeleteRawModifiedDetails");
    // Create the instance
    return new DeleteRawModifiedDetailsBuilderImpl(nodeId, isDeleteModified, startTime, endTime);
  }

  public static class DeleteRawModifiedDetailsBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId nodeId;
    private final boolean isDeleteModified;
    private final long startTime;
    private final long endTime;

    public DeleteRawModifiedDetailsBuilderImpl(
        NodeId nodeId, boolean isDeleteModified, long startTime, long endTime) {
      this.nodeId = nodeId;
      this.isDeleteModified = isDeleteModified;
      this.startTime = startTime;
      this.endTime = endTime;
    }

    public DeleteRawModifiedDetails build() {
      DeleteRawModifiedDetails deleteRawModifiedDetails =
          new DeleteRawModifiedDetails(nodeId, isDeleteModified, startTime, endTime);
      return deleteRawModifiedDetails;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DeleteRawModifiedDetails)) {
      return false;
    }
    DeleteRawModifiedDetails that = (DeleteRawModifiedDetails) o;
    return (getNodeId() == that.getNodeId())
        && (getIsDeleteModified() == that.getIsDeleteModified())
        && (getStartTime() == that.getStartTime())
        && (getEndTime() == that.getEndTime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getNodeId(), getIsDeleteModified(), getStartTime(), getEndTime());
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