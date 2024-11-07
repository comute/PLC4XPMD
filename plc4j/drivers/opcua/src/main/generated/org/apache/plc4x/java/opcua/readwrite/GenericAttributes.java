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

public class GenericAttributes extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 17609;
  }

  // Properties.
  protected final long specifiedAttributes;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final long writeMask;
  protected final long userWriteMask;
  protected final List<GenericAttributeValue> attributeValues;

  public GenericAttributes(
      long specifiedAttributes,
      LocalizedText displayName,
      LocalizedText description,
      long writeMask,
      long userWriteMask,
      List<GenericAttributeValue> attributeValues) {
    super();
    this.specifiedAttributes = specifiedAttributes;
    this.displayName = displayName;
    this.description = description;
    this.writeMask = writeMask;
    this.userWriteMask = userWriteMask;
    this.attributeValues = attributeValues;
  }

  public long getSpecifiedAttributes() {
    return specifiedAttributes;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public long getWriteMask() {
    return writeMask;
  }

  public long getUserWriteMask() {
    return userWriteMask;
  }

  public List<GenericAttributeValue> getAttributeValues() {
    return attributeValues;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("GenericAttributes");

    // Simple Field (specifiedAttributes)
    writeSimpleField(
        "specifiedAttributes", specifiedAttributes, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (writeMask)
    writeSimpleField("writeMask", writeMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (userWriteMask)
    writeSimpleField("userWriteMask", userWriteMask, writeUnsignedLong(writeBuffer, 32));

    // Implicit Field (noOfAttributeValues) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfAttributeValues =
        (int) ((((getAttributeValues()) == (null)) ? -(1) : COUNT(getAttributeValues())));
    writeImplicitField("noOfAttributeValues", noOfAttributeValues, writeSignedInt(writeBuffer, 32));

    // Array Field (attributeValues)
    writeComplexTypeArrayField("attributeValues", attributeValues, writeBuffer);

    writeBuffer.popContext("GenericAttributes");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    GenericAttributes _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (specifiedAttributes)
    lengthInBits += 32;

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (writeMask)
    lengthInBits += 32;

    // Simple field (userWriteMask)
    lengthInBits += 32;

    // Implicit Field (noOfAttributeValues)
    lengthInBits += 32;

    // Array field
    if (attributeValues != null) {
      int i = 0;
      for (GenericAttributeValue element : attributeValues) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= attributeValues.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("GenericAttributes");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long specifiedAttributes =
        readSimpleField("specifiedAttributes", readUnsignedLong(readBuffer, 32));

    LocalizedText displayName =
        readSimpleField(
            "displayName", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    long writeMask = readSimpleField("writeMask", readUnsignedLong(readBuffer, 32));

    long userWriteMask = readSimpleField("userWriteMask", readUnsignedLong(readBuffer, 32));

    int noOfAttributeValues =
        readImplicitField("noOfAttributeValues", readSignedInt(readBuffer, 32));

    List<GenericAttributeValue> attributeValues =
        readCountArrayField(
            "attributeValues",
            readComplex(
                () ->
                    (GenericAttributeValue)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (17608)),
                readBuffer),
            noOfAttributeValues);

    readBuffer.closeContext("GenericAttributes");
    // Create the instance
    return new GenericAttributesBuilderImpl(
        specifiedAttributes, displayName, description, writeMask, userWriteMask, attributeValues);
  }

  public static class GenericAttributesBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long specifiedAttributes;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final long writeMask;
    private final long userWriteMask;
    private final List<GenericAttributeValue> attributeValues;

    public GenericAttributesBuilderImpl(
        long specifiedAttributes,
        LocalizedText displayName,
        LocalizedText description,
        long writeMask,
        long userWriteMask,
        List<GenericAttributeValue> attributeValues) {
      this.specifiedAttributes = specifiedAttributes;
      this.displayName = displayName;
      this.description = description;
      this.writeMask = writeMask;
      this.userWriteMask = userWriteMask;
      this.attributeValues = attributeValues;
    }

    public GenericAttributes build() {
      GenericAttributes genericAttributes =
          new GenericAttributes(
              specifiedAttributes,
              displayName,
              description,
              writeMask,
              userWriteMask,
              attributeValues);
      return genericAttributes;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof GenericAttributes)) {
      return false;
    }
    GenericAttributes that = (GenericAttributes) o;
    return (getSpecifiedAttributes() == that.getSpecifiedAttributes())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getWriteMask() == that.getWriteMask())
        && (getUserWriteMask() == that.getUserWriteMask())
        && (getAttributeValues() == that.getAttributeValues())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSpecifiedAttributes(),
        getDisplayName(),
        getDescription(),
        getWriteMask(),
        getUserWriteMask(),
        getAttributeValues());
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