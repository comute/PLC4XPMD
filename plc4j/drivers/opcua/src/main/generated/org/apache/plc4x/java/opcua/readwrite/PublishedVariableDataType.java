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

public class PublishedVariableDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 14275;
  }

  // Properties.
  protected final NodeId publishedVariable;
  protected final long attributeId;
  protected final double samplingIntervalHint;
  protected final long deadbandType;
  protected final double deadbandValue;
  protected final PascalString indexRange;
  protected final Variant substituteValue;
  protected final List<QualifiedName> metaDataProperties;

  public PublishedVariableDataType(
      NodeId publishedVariable,
      long attributeId,
      double samplingIntervalHint,
      long deadbandType,
      double deadbandValue,
      PascalString indexRange,
      Variant substituteValue,
      List<QualifiedName> metaDataProperties) {
    super();
    this.publishedVariable = publishedVariable;
    this.attributeId = attributeId;
    this.samplingIntervalHint = samplingIntervalHint;
    this.deadbandType = deadbandType;
    this.deadbandValue = deadbandValue;
    this.indexRange = indexRange;
    this.substituteValue = substituteValue;
    this.metaDataProperties = metaDataProperties;
  }

  public NodeId getPublishedVariable() {
    return publishedVariable;
  }

  public long getAttributeId() {
    return attributeId;
  }

  public double getSamplingIntervalHint() {
    return samplingIntervalHint;
  }

  public long getDeadbandType() {
    return deadbandType;
  }

  public double getDeadbandValue() {
    return deadbandValue;
  }

  public PascalString getIndexRange() {
    return indexRange;
  }

  public Variant getSubstituteValue() {
    return substituteValue;
  }

  public List<QualifiedName> getMetaDataProperties() {
    return metaDataProperties;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PublishedVariableDataType");

    // Simple Field (publishedVariable)
    writeSimpleField("publishedVariable", publishedVariable, writeComplex(writeBuffer));

    // Simple Field (attributeId)
    writeSimpleField("attributeId", attributeId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (samplingIntervalHint)
    writeSimpleField("samplingIntervalHint", samplingIntervalHint, writeDouble(writeBuffer, 64));

    // Simple Field (deadbandType)
    writeSimpleField("deadbandType", deadbandType, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (deadbandValue)
    writeSimpleField("deadbandValue", deadbandValue, writeDouble(writeBuffer, 64));

    // Simple Field (indexRange)
    writeSimpleField("indexRange", indexRange, writeComplex(writeBuffer));

    // Simple Field (substituteValue)
    writeSimpleField("substituteValue", substituteValue, writeComplex(writeBuffer));

    // Implicit Field (noOfMetaDataProperties) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfMetaDataProperties =
        (int) ((((getMetaDataProperties()) == (null)) ? -(1) : COUNT(getMetaDataProperties())));
    writeImplicitField(
        "noOfMetaDataProperties", noOfMetaDataProperties, writeSignedInt(writeBuffer, 32));

    // Array Field (metaDataProperties)
    writeComplexTypeArrayField("metaDataProperties", metaDataProperties, writeBuffer);

    writeBuffer.popContext("PublishedVariableDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PublishedVariableDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (publishedVariable)
    lengthInBits += publishedVariable.getLengthInBits();

    // Simple field (attributeId)
    lengthInBits += 32;

    // Simple field (samplingIntervalHint)
    lengthInBits += 64;

    // Simple field (deadbandType)
    lengthInBits += 32;

    // Simple field (deadbandValue)
    lengthInBits += 64;

    // Simple field (indexRange)
    lengthInBits += indexRange.getLengthInBits();

    // Simple field (substituteValue)
    lengthInBits += substituteValue.getLengthInBits();

    // Implicit Field (noOfMetaDataProperties)
    lengthInBits += 32;

    // Array field
    if (metaDataProperties != null) {
      int i = 0;
      for (QualifiedName element : metaDataProperties) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= metaDataProperties.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PublishedVariableDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId publishedVariable =
        readSimpleField(
            "publishedVariable", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    long attributeId = readSimpleField("attributeId", readUnsignedLong(readBuffer, 32));

    double samplingIntervalHint =
        readSimpleField("samplingIntervalHint", readDouble(readBuffer, 64));

    long deadbandType = readSimpleField("deadbandType", readUnsignedLong(readBuffer, 32));

    double deadbandValue = readSimpleField("deadbandValue", readDouble(readBuffer, 64));

    PascalString indexRange =
        readSimpleField(
            "indexRange", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    Variant substituteValue =
        readSimpleField(
            "substituteValue", readComplex(() -> Variant.staticParse(readBuffer), readBuffer));

    int noOfMetaDataProperties =
        readImplicitField("noOfMetaDataProperties", readSignedInt(readBuffer, 32));

    List<QualifiedName> metaDataProperties =
        readCountArrayField(
            "metaDataProperties",
            readComplex(() -> QualifiedName.staticParse(readBuffer), readBuffer),
            noOfMetaDataProperties);

    readBuffer.closeContext("PublishedVariableDataType");
    // Create the instance
    return new PublishedVariableDataTypeBuilderImpl(
        publishedVariable,
        attributeId,
        samplingIntervalHint,
        deadbandType,
        deadbandValue,
        indexRange,
        substituteValue,
        metaDataProperties);
  }

  public static class PublishedVariableDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId publishedVariable;
    private final long attributeId;
    private final double samplingIntervalHint;
    private final long deadbandType;
    private final double deadbandValue;
    private final PascalString indexRange;
    private final Variant substituteValue;
    private final List<QualifiedName> metaDataProperties;

    public PublishedVariableDataTypeBuilderImpl(
        NodeId publishedVariable,
        long attributeId,
        double samplingIntervalHint,
        long deadbandType,
        double deadbandValue,
        PascalString indexRange,
        Variant substituteValue,
        List<QualifiedName> metaDataProperties) {
      this.publishedVariable = publishedVariable;
      this.attributeId = attributeId;
      this.samplingIntervalHint = samplingIntervalHint;
      this.deadbandType = deadbandType;
      this.deadbandValue = deadbandValue;
      this.indexRange = indexRange;
      this.substituteValue = substituteValue;
      this.metaDataProperties = metaDataProperties;
    }

    public PublishedVariableDataType build() {
      PublishedVariableDataType publishedVariableDataType =
          new PublishedVariableDataType(
              publishedVariable,
              attributeId,
              samplingIntervalHint,
              deadbandType,
              deadbandValue,
              indexRange,
              substituteValue,
              metaDataProperties);
      return publishedVariableDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PublishedVariableDataType)) {
      return false;
    }
    PublishedVariableDataType that = (PublishedVariableDataType) o;
    return (getPublishedVariable() == that.getPublishedVariable())
        && (getAttributeId() == that.getAttributeId())
        && (getSamplingIntervalHint() == that.getSamplingIntervalHint())
        && (getDeadbandType() == that.getDeadbandType())
        && (getDeadbandValue() == that.getDeadbandValue())
        && (getIndexRange() == that.getIndexRange())
        && (getSubstituteValue() == that.getSubstituteValue())
        && (getMetaDataProperties() == that.getMetaDataProperties())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getPublishedVariable(),
        getAttributeId(),
        getSamplingIntervalHint(),
        getDeadbandType(),
        getDeadbandValue(),
        getIndexRange(),
        getSubstituteValue(),
        getMetaDataProperties());
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
