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

public class UadpWriterGroupMessageDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15647;
  }

  // Properties.
  protected final long groupVersion;
  protected final DataSetOrderingType dataSetOrdering;
  protected final UadpNetworkMessageContentMask networkMessageContentMask;
  protected final double samplingOffset;
  protected final List<Double> publishingOffset;

  public UadpWriterGroupMessageDataType(
      long groupVersion,
      DataSetOrderingType dataSetOrdering,
      UadpNetworkMessageContentMask networkMessageContentMask,
      double samplingOffset,
      List<Double> publishingOffset) {
    super();
    this.groupVersion = groupVersion;
    this.dataSetOrdering = dataSetOrdering;
    this.networkMessageContentMask = networkMessageContentMask;
    this.samplingOffset = samplingOffset;
    this.publishingOffset = publishingOffset;
  }

  public long getGroupVersion() {
    return groupVersion;
  }

  public DataSetOrderingType getDataSetOrdering() {
    return dataSetOrdering;
  }

  public UadpNetworkMessageContentMask getNetworkMessageContentMask() {
    return networkMessageContentMask;
  }

  public double getSamplingOffset() {
    return samplingOffset;
  }

  public List<Double> getPublishingOffset() {
    return publishingOffset;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("UadpWriterGroupMessageDataType");

    // Simple Field (groupVersion)
    writeSimpleField("groupVersion", groupVersion, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (dataSetOrdering)
    writeSimpleEnumField(
        "dataSetOrdering",
        "DataSetOrderingType",
        dataSetOrdering,
        writeEnum(
            DataSetOrderingType::getValue,
            DataSetOrderingType::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (networkMessageContentMask)
    writeSimpleEnumField(
        "networkMessageContentMask",
        "UadpNetworkMessageContentMask",
        networkMessageContentMask,
        writeEnum(
            UadpNetworkMessageContentMask::getValue,
            UadpNetworkMessageContentMask::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (samplingOffset)
    writeSimpleField("samplingOffset", samplingOffset, writeDouble(writeBuffer, 64));

    // Implicit Field (noOfPublishingOffset) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfPublishingOffset =
        (int) ((((getPublishingOffset()) == (null)) ? -(1) : COUNT(getPublishingOffset())));
    writeImplicitField(
        "noOfPublishingOffset", noOfPublishingOffset, writeSignedInt(writeBuffer, 32));

    // Array Field (publishingOffset)
    writeSimpleTypeArrayField("publishingOffset", publishingOffset, writeDouble(writeBuffer, 64));

    writeBuffer.popContext("UadpWriterGroupMessageDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    UadpWriterGroupMessageDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (groupVersion)
    lengthInBits += 32;

    // Simple field (dataSetOrdering)
    lengthInBits += 32;

    // Simple field (networkMessageContentMask)
    lengthInBits += 32;

    // Simple field (samplingOffset)
    lengthInBits += 64;

    // Implicit Field (noOfPublishingOffset)
    lengthInBits += 32;

    // Array field
    if (publishingOffset != null) {
      lengthInBits += 64 * publishingOffset.size();
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("UadpWriterGroupMessageDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long groupVersion = readSimpleField("groupVersion", readUnsignedLong(readBuffer, 32));

    DataSetOrderingType dataSetOrdering =
        readEnumField(
            "dataSetOrdering",
            "DataSetOrderingType",
            readEnum(DataSetOrderingType::enumForValue, readUnsignedLong(readBuffer, 32)));

    UadpNetworkMessageContentMask networkMessageContentMask =
        readEnumField(
            "networkMessageContentMask",
            "UadpNetworkMessageContentMask",
            readEnum(
                UadpNetworkMessageContentMask::enumForValue, readUnsignedLong(readBuffer, 32)));

    double samplingOffset = readSimpleField("samplingOffset", readDouble(readBuffer, 64));

    int noOfPublishingOffset =
        readImplicitField("noOfPublishingOffset", readSignedInt(readBuffer, 32));

    List<Double> publishingOffset =
        readCountArrayField("publishingOffset", readDouble(readBuffer, 64), noOfPublishingOffset);

    readBuffer.closeContext("UadpWriterGroupMessageDataType");
    // Create the instance
    return new UadpWriterGroupMessageDataTypeBuilderImpl(
        groupVersion, dataSetOrdering, networkMessageContentMask, samplingOffset, publishingOffset);
  }

  public static class UadpWriterGroupMessageDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long groupVersion;
    private final DataSetOrderingType dataSetOrdering;
    private final UadpNetworkMessageContentMask networkMessageContentMask;
    private final double samplingOffset;
    private final List<Double> publishingOffset;

    public UadpWriterGroupMessageDataTypeBuilderImpl(
        long groupVersion,
        DataSetOrderingType dataSetOrdering,
        UadpNetworkMessageContentMask networkMessageContentMask,
        double samplingOffset,
        List<Double> publishingOffset) {
      this.groupVersion = groupVersion;
      this.dataSetOrdering = dataSetOrdering;
      this.networkMessageContentMask = networkMessageContentMask;
      this.samplingOffset = samplingOffset;
      this.publishingOffset = publishingOffset;
    }

    public UadpWriterGroupMessageDataType build() {
      UadpWriterGroupMessageDataType uadpWriterGroupMessageDataType =
          new UadpWriterGroupMessageDataType(
              groupVersion,
              dataSetOrdering,
              networkMessageContentMask,
              samplingOffset,
              publishingOffset);
      return uadpWriterGroupMessageDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof UadpWriterGroupMessageDataType)) {
      return false;
    }
    UadpWriterGroupMessageDataType that = (UadpWriterGroupMessageDataType) o;
    return (getGroupVersion() == that.getGroupVersion())
        && (getDataSetOrdering() == that.getDataSetOrdering())
        && (getNetworkMessageContentMask() == that.getNetworkMessageContentMask())
        && (getSamplingOffset() == that.getSamplingOffset())
        && (getPublishingOffset() == that.getPublishingOffset())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getGroupVersion(),
        getDataSetOrdering(),
        getNetworkMessageContentMask(),
        getSamplingOffset(),
        getPublishingOffset());
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