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

public class FieldTargetDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "14746";
  }

  // Properties.
  protected final GuidValue dataSetFieldId;
  protected final PascalString receiverIndexRange;
  protected final NodeId targetNodeId;
  protected final long attributeId;
  protected final PascalString writeIndexRange;
  protected final OverrideValueHandling overrideValueHandling;
  protected final Variant overrideValue;

  public FieldTargetDataType(
      GuidValue dataSetFieldId,
      PascalString receiverIndexRange,
      NodeId targetNodeId,
      long attributeId,
      PascalString writeIndexRange,
      OverrideValueHandling overrideValueHandling,
      Variant overrideValue) {
    super();
    this.dataSetFieldId = dataSetFieldId;
    this.receiverIndexRange = receiverIndexRange;
    this.targetNodeId = targetNodeId;
    this.attributeId = attributeId;
    this.writeIndexRange = writeIndexRange;
    this.overrideValueHandling = overrideValueHandling;
    this.overrideValue = overrideValue;
  }

  public GuidValue getDataSetFieldId() {
    return dataSetFieldId;
  }

  public PascalString getReceiverIndexRange() {
    return receiverIndexRange;
  }

  public NodeId getTargetNodeId() {
    return targetNodeId;
  }

  public long getAttributeId() {
    return attributeId;
  }

  public PascalString getWriteIndexRange() {
    return writeIndexRange;
  }

  public OverrideValueHandling getOverrideValueHandling() {
    return overrideValueHandling;
  }

  public Variant getOverrideValue() {
    return overrideValue;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("FieldTargetDataType");

    // Simple Field (dataSetFieldId)
    writeSimpleField("dataSetFieldId", dataSetFieldId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (receiverIndexRange)
    writeSimpleField(
        "receiverIndexRange", receiverIndexRange, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (targetNodeId)
    writeSimpleField("targetNodeId", targetNodeId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (attributeId)
    writeSimpleField("attributeId", attributeId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (writeIndexRange)
    writeSimpleField(
        "writeIndexRange", writeIndexRange, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (overrideValueHandling)
    writeSimpleEnumField(
        "overrideValueHandling",
        "OverrideValueHandling",
        overrideValueHandling,
        new DataWriterEnumDefault<>(
            OverrideValueHandling::getValue,
            OverrideValueHandling::name,
            writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (overrideValue)
    writeSimpleField("overrideValue", overrideValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("FieldTargetDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FieldTargetDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dataSetFieldId)
    lengthInBits += dataSetFieldId.getLengthInBits();

    // Simple field (receiverIndexRange)
    lengthInBits += receiverIndexRange.getLengthInBits();

    // Simple field (targetNodeId)
    lengthInBits += targetNodeId.getLengthInBits();

    // Simple field (attributeId)
    lengthInBits += 32;

    // Simple field (writeIndexRange)
    lengthInBits += writeIndexRange.getLengthInBits();

    // Simple field (overrideValueHandling)
    lengthInBits += 32;

    // Simple field (overrideValue)
    lengthInBits += overrideValue.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("FieldTargetDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    GuidValue dataSetFieldId =
        readSimpleField(
            "dataSetFieldId",
            new DataReaderComplexDefault<>(() -> GuidValue.staticParse(readBuffer), readBuffer));

    PascalString receiverIndexRange =
        readSimpleField(
            "receiverIndexRange",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    NodeId targetNodeId =
        readSimpleField(
            "targetNodeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    long attributeId = readSimpleField("attributeId", readUnsignedLong(readBuffer, 32));

    PascalString writeIndexRange =
        readSimpleField(
            "writeIndexRange",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    OverrideValueHandling overrideValueHandling =
        readEnumField(
            "overrideValueHandling",
            "OverrideValueHandling",
            new DataReaderEnumDefault<>(
                OverrideValueHandling::enumForValue, readUnsignedLong(readBuffer, 32)));

    Variant overrideValue =
        readSimpleField(
            "overrideValue",
            new DataReaderComplexDefault<>(() -> Variant.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("FieldTargetDataType");
    // Create the instance
    return new FieldTargetDataTypeBuilderImpl(
        dataSetFieldId,
        receiverIndexRange,
        targetNodeId,
        attributeId,
        writeIndexRange,
        overrideValueHandling,
        overrideValue);
  }

  public static class FieldTargetDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final GuidValue dataSetFieldId;
    private final PascalString receiverIndexRange;
    private final NodeId targetNodeId;
    private final long attributeId;
    private final PascalString writeIndexRange;
    private final OverrideValueHandling overrideValueHandling;
    private final Variant overrideValue;

    public FieldTargetDataTypeBuilderImpl(
        GuidValue dataSetFieldId,
        PascalString receiverIndexRange,
        NodeId targetNodeId,
        long attributeId,
        PascalString writeIndexRange,
        OverrideValueHandling overrideValueHandling,
        Variant overrideValue) {
      this.dataSetFieldId = dataSetFieldId;
      this.receiverIndexRange = receiverIndexRange;
      this.targetNodeId = targetNodeId;
      this.attributeId = attributeId;
      this.writeIndexRange = writeIndexRange;
      this.overrideValueHandling = overrideValueHandling;
      this.overrideValue = overrideValue;
    }

    public FieldTargetDataType build() {
      FieldTargetDataType fieldTargetDataType =
          new FieldTargetDataType(
              dataSetFieldId,
              receiverIndexRange,
              targetNodeId,
              attributeId,
              writeIndexRange,
              overrideValueHandling,
              overrideValue);
      return fieldTargetDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FieldTargetDataType)) {
      return false;
    }
    FieldTargetDataType that = (FieldTargetDataType) o;
    return (getDataSetFieldId() == that.getDataSetFieldId())
        && (getReceiverIndexRange() == that.getReceiverIndexRange())
        && (getTargetNodeId() == that.getTargetNodeId())
        && (getAttributeId() == that.getAttributeId())
        && (getWriteIndexRange() == that.getWriteIndexRange())
        && (getOverrideValueHandling() == that.getOverrideValueHandling())
        && (getOverrideValue() == that.getOverrideValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getDataSetFieldId(),
        getReceiverIndexRange(),
        getTargetNodeId(),
        getAttributeId(),
        getWriteIndexRange(),
        getOverrideValueHandling(),
        getOverrideValue());
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
