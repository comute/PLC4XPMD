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

public class PublishedDataSetDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15580;
  }

  // Properties.
  protected final PascalString name;
  protected final List<PascalString> dataSetFolder;
  protected final DataSetMetaDataType dataSetMetaData;
  protected final List<KeyValuePair> extensionFields;
  protected final ExtensionObject dataSetSource;

  public PublishedDataSetDataType(
      PascalString name,
      List<PascalString> dataSetFolder,
      DataSetMetaDataType dataSetMetaData,
      List<KeyValuePair> extensionFields,
      ExtensionObject dataSetSource) {
    super();
    this.name = name;
    this.dataSetFolder = dataSetFolder;
    this.dataSetMetaData = dataSetMetaData;
    this.extensionFields = extensionFields;
    this.dataSetSource = dataSetSource;
  }

  public PascalString getName() {
    return name;
  }

  public List<PascalString> getDataSetFolder() {
    return dataSetFolder;
  }

  public DataSetMetaDataType getDataSetMetaData() {
    return dataSetMetaData;
  }

  public List<KeyValuePair> getExtensionFields() {
    return extensionFields;
  }

  public ExtensionObject getDataSetSource() {
    return dataSetSource;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PublishedDataSetDataType");

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Implicit Field (noOfDataSetFolder) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfDataSetFolder =
        (int) ((((getDataSetFolder()) == (null)) ? -(1) : COUNT(getDataSetFolder())));
    writeImplicitField("noOfDataSetFolder", noOfDataSetFolder, writeSignedInt(writeBuffer, 32));

    // Array Field (dataSetFolder)
    writeComplexTypeArrayField("dataSetFolder", dataSetFolder, writeBuffer);

    // Simple Field (dataSetMetaData)
    writeSimpleField("dataSetMetaData", dataSetMetaData, writeComplex(writeBuffer));

    // Implicit Field (noOfExtensionFields) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfExtensionFields =
        (int) ((((getExtensionFields()) == (null)) ? -(1) : COUNT(getExtensionFields())));
    writeImplicitField("noOfExtensionFields", noOfExtensionFields, writeSignedInt(writeBuffer, 32));

    // Array Field (extensionFields)
    writeComplexTypeArrayField("extensionFields", extensionFields, writeBuffer);

    // Simple Field (dataSetSource)
    writeSimpleField("dataSetSource", dataSetSource, writeComplex(writeBuffer));

    writeBuffer.popContext("PublishedDataSetDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PublishedDataSetDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Implicit Field (noOfDataSetFolder)
    lengthInBits += 32;

    // Array field
    if (dataSetFolder != null) {
      int i = 0;
      for (PascalString element : dataSetFolder) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= dataSetFolder.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (dataSetMetaData)
    lengthInBits += dataSetMetaData.getLengthInBits();

    // Implicit Field (noOfExtensionFields)
    lengthInBits += 32;

    // Array field
    if (extensionFields != null) {
      int i = 0;
      for (KeyValuePair element : extensionFields) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= extensionFields.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (dataSetSource)
    lengthInBits += dataSetSource.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PublishedDataSetDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfDataSetFolder = readImplicitField("noOfDataSetFolder", readSignedInt(readBuffer, 32));

    List<PascalString> dataSetFolder =
        readCountArrayField(
            "dataSetFolder",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfDataSetFolder);

    DataSetMetaDataType dataSetMetaData =
        readSimpleField(
            "dataSetMetaData",
            readComplex(
                () ->
                    (DataSetMetaDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (14525)),
                readBuffer));

    int noOfExtensionFields =
        readImplicitField("noOfExtensionFields", readSignedInt(readBuffer, 32));

    List<KeyValuePair> extensionFields =
        readCountArrayField(
            "extensionFields",
            readComplex(
                () ->
                    (KeyValuePair) ExtensionObjectDefinition.staticParse(readBuffer, (int) (14535)),
                readBuffer),
            noOfExtensionFields);

    ExtensionObject dataSetSource =
        readSimpleField(
            "dataSetSource",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("PublishedDataSetDataType");
    // Create the instance
    return new PublishedDataSetDataTypeBuilderImpl(
        name, dataSetFolder, dataSetMetaData, extensionFields, dataSetSource);
  }

  public static class PublishedDataSetDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final List<PascalString> dataSetFolder;
    private final DataSetMetaDataType dataSetMetaData;
    private final List<KeyValuePair> extensionFields;
    private final ExtensionObject dataSetSource;

    public PublishedDataSetDataTypeBuilderImpl(
        PascalString name,
        List<PascalString> dataSetFolder,
        DataSetMetaDataType dataSetMetaData,
        List<KeyValuePair> extensionFields,
        ExtensionObject dataSetSource) {
      this.name = name;
      this.dataSetFolder = dataSetFolder;
      this.dataSetMetaData = dataSetMetaData;
      this.extensionFields = extensionFields;
      this.dataSetSource = dataSetSource;
    }

    public PublishedDataSetDataType build() {
      PublishedDataSetDataType publishedDataSetDataType =
          new PublishedDataSetDataType(
              name, dataSetFolder, dataSetMetaData, extensionFields, dataSetSource);
      return publishedDataSetDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PublishedDataSetDataType)) {
      return false;
    }
    PublishedDataSetDataType that = (PublishedDataSetDataType) o;
    return (getName() == that.getName())
        && (getDataSetFolder() == that.getDataSetFolder())
        && (getDataSetMetaData() == that.getDataSetMetaData())
        && (getExtensionFields() == that.getExtensionFields())
        && (getDataSetSource() == that.getDataSetSource())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getDataSetFolder(),
        getDataSetMetaData(),
        getExtensionFields(),
        getDataSetSource());
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
