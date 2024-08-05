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

public class FieldMetaData extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 14526;
  }

  // Properties.
  protected final PascalString name;
  protected final LocalizedText description;
  protected final DataSetFieldFlags fieldFlags;
  protected final short builtInType;
  protected final NodeId dataType;
  protected final int valueRank;
  protected final List<Long> arrayDimensions;
  protected final long maxStringLength;
  protected final GuidValue dataSetFieldId;
  protected final List<KeyValuePair> properties;

  public FieldMetaData(
      PascalString name,
      LocalizedText description,
      DataSetFieldFlags fieldFlags,
      short builtInType,
      NodeId dataType,
      int valueRank,
      List<Long> arrayDimensions,
      long maxStringLength,
      GuidValue dataSetFieldId,
      List<KeyValuePair> properties) {
    super();
    this.name = name;
    this.description = description;
    this.fieldFlags = fieldFlags;
    this.builtInType = builtInType;
    this.dataType = dataType;
    this.valueRank = valueRank;
    this.arrayDimensions = arrayDimensions;
    this.maxStringLength = maxStringLength;
    this.dataSetFieldId = dataSetFieldId;
    this.properties = properties;
  }

  public PascalString getName() {
    return name;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public DataSetFieldFlags getFieldFlags() {
    return fieldFlags;
  }

  public short getBuiltInType() {
    return builtInType;
  }

  public NodeId getDataType() {
    return dataType;
  }

  public int getValueRank() {
    return valueRank;
  }

  public List<Long> getArrayDimensions() {
    return arrayDimensions;
  }

  public long getMaxStringLength() {
    return maxStringLength;
  }

  public GuidValue getDataSetFieldId() {
    return dataSetFieldId;
  }

  public List<KeyValuePair> getProperties() {
    return properties;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("FieldMetaData");

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (fieldFlags)
    writeSimpleEnumField(
        "fieldFlags",
        "DataSetFieldFlags",
        fieldFlags,
        writeEnum(
            DataSetFieldFlags::getValue,
            DataSetFieldFlags::name,
            writeUnsignedInt(writeBuffer, 16)));

    // Simple Field (builtInType)
    writeSimpleField("builtInType", builtInType, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (dataType)
    writeSimpleField("dataType", dataType, writeComplex(writeBuffer));

    // Simple Field (valueRank)
    writeSimpleField("valueRank", valueRank, writeSignedInt(writeBuffer, 32));

    // Implicit Field (noOfArrayDimensions) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfArrayDimensions =
        (int) ((((getArrayDimensions()) == (null)) ? -(1) : COUNT(getArrayDimensions())));
    writeImplicitField("noOfArrayDimensions", noOfArrayDimensions, writeSignedInt(writeBuffer, 32));

    // Array Field (arrayDimensions)
    writeSimpleTypeArrayField(
        "arrayDimensions", arrayDimensions, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxStringLength)
    writeSimpleField("maxStringLength", maxStringLength, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (dataSetFieldId)
    writeSimpleField("dataSetFieldId", dataSetFieldId, writeComplex(writeBuffer));

    // Implicit Field (noOfProperties) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfProperties = (int) ((((getProperties()) == (null)) ? -(1) : COUNT(getProperties())));
    writeImplicitField("noOfProperties", noOfProperties, writeSignedInt(writeBuffer, 32));

    // Array Field (properties)
    writeComplexTypeArrayField("properties", properties, writeBuffer);

    writeBuffer.popContext("FieldMetaData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FieldMetaData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (fieldFlags)
    lengthInBits += 16;

    // Simple field (builtInType)
    lengthInBits += 8;

    // Simple field (dataType)
    lengthInBits += dataType.getLengthInBits();

    // Simple field (valueRank)
    lengthInBits += 32;

    // Implicit Field (noOfArrayDimensions)
    lengthInBits += 32;

    // Array field
    if (arrayDimensions != null) {
      lengthInBits += 32 * arrayDimensions.size();
    }

    // Simple field (maxStringLength)
    lengthInBits += 32;

    // Simple field (dataSetFieldId)
    lengthInBits += dataSetFieldId.getLengthInBits();

    // Implicit Field (noOfProperties)
    lengthInBits += 32;

    // Array field
    if (properties != null) {
      int i = 0;
      for (KeyValuePair element : properties) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= properties.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("FieldMetaData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    DataSetFieldFlags fieldFlags =
        readEnumField(
            "fieldFlags",
            "DataSetFieldFlags",
            readEnum(DataSetFieldFlags::enumForValue, readUnsignedInt(readBuffer, 16)));

    short builtInType = readSimpleField("builtInType", readUnsignedShort(readBuffer, 8));

    NodeId dataType =
        readSimpleField("dataType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int valueRank = readSimpleField("valueRank", readSignedInt(readBuffer, 32));

    int noOfArrayDimensions =
        readImplicitField("noOfArrayDimensions", readSignedInt(readBuffer, 32));

    List<Long> arrayDimensions =
        readCountArrayField(
            "arrayDimensions", readUnsignedLong(readBuffer, 32), noOfArrayDimensions);

    long maxStringLength = readSimpleField("maxStringLength", readUnsignedLong(readBuffer, 32));

    GuidValue dataSetFieldId =
        readSimpleField(
            "dataSetFieldId", readComplex(() -> GuidValue.staticParse(readBuffer), readBuffer));

    int noOfProperties = readImplicitField("noOfProperties", readSignedInt(readBuffer, 32));

    List<KeyValuePair> properties =
        readCountArrayField(
            "properties",
            readComplex(
                () ->
                    (KeyValuePair) ExtensionObjectDefinition.staticParse(readBuffer, (int) (14535)),
                readBuffer),
            noOfProperties);

    readBuffer.closeContext("FieldMetaData");
    // Create the instance
    return new FieldMetaDataBuilderImpl(
        name,
        description,
        fieldFlags,
        builtInType,
        dataType,
        valueRank,
        arrayDimensions,
        maxStringLength,
        dataSetFieldId,
        properties);
  }

  public static class FieldMetaDataBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final LocalizedText description;
    private final DataSetFieldFlags fieldFlags;
    private final short builtInType;
    private final NodeId dataType;
    private final int valueRank;
    private final List<Long> arrayDimensions;
    private final long maxStringLength;
    private final GuidValue dataSetFieldId;
    private final List<KeyValuePair> properties;

    public FieldMetaDataBuilderImpl(
        PascalString name,
        LocalizedText description,
        DataSetFieldFlags fieldFlags,
        short builtInType,
        NodeId dataType,
        int valueRank,
        List<Long> arrayDimensions,
        long maxStringLength,
        GuidValue dataSetFieldId,
        List<KeyValuePair> properties) {
      this.name = name;
      this.description = description;
      this.fieldFlags = fieldFlags;
      this.builtInType = builtInType;
      this.dataType = dataType;
      this.valueRank = valueRank;
      this.arrayDimensions = arrayDimensions;
      this.maxStringLength = maxStringLength;
      this.dataSetFieldId = dataSetFieldId;
      this.properties = properties;
    }

    public FieldMetaData build() {
      FieldMetaData fieldMetaData =
          new FieldMetaData(
              name,
              description,
              fieldFlags,
              builtInType,
              dataType,
              valueRank,
              arrayDimensions,
              maxStringLength,
              dataSetFieldId,
              properties);
      return fieldMetaData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FieldMetaData)) {
      return false;
    }
    FieldMetaData that = (FieldMetaData) o;
    return (getName() == that.getName())
        && (getDescription() == that.getDescription())
        && (getFieldFlags() == that.getFieldFlags())
        && (getBuiltInType() == that.getBuiltInType())
        && (getDataType() == that.getDataType())
        && (getValueRank() == that.getValueRank())
        && (getArrayDimensions() == that.getArrayDimensions())
        && (getMaxStringLength() == that.getMaxStringLength())
        && (getDataSetFieldId() == that.getDataSetFieldId())
        && (getProperties() == that.getProperties())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getDescription(),
        getFieldFlags(),
        getBuiltInType(),
        getDataType(),
        getValueRank(),
        getArrayDimensions(),
        getMaxStringLength(),
        getDataSetFieldId(),
        getProperties());
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
