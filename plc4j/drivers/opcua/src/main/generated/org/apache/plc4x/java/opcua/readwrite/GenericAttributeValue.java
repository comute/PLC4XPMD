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

public class GenericAttributeValue extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "17608";
  }

  // Properties.
  protected final long attributeId;
  protected final Variant value;

  public GenericAttributeValue(long attributeId, Variant value) {
    super();
    this.attributeId = attributeId;
    this.value = value;
  }

  public long getAttributeId() {
    return attributeId;
  }

  public Variant getValue() {
    return value;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("GenericAttributeValue");

    // Simple Field (attributeId)
    writeSimpleField("attributeId", attributeId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (value)
    writeSimpleField("value", value, writeComplex(writeBuffer));

    writeBuffer.popContext("GenericAttributeValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    GenericAttributeValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (attributeId)
    lengthInBits += 32;

    // Simple field (value)
    lengthInBits += value.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("GenericAttributeValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long attributeId = readSimpleField("attributeId", readUnsignedLong(readBuffer, 32));

    Variant value =
        readSimpleField("value", readComplex(() -> Variant.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("GenericAttributeValue");
    // Create the instance
    return new GenericAttributeValueBuilderImpl(attributeId, value);
  }

  public static class GenericAttributeValueBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long attributeId;
    private final Variant value;

    public GenericAttributeValueBuilderImpl(long attributeId, Variant value) {
      this.attributeId = attributeId;
      this.value = value;
    }

    public GenericAttributeValue build() {
      GenericAttributeValue genericAttributeValue = new GenericAttributeValue(attributeId, value);
      return genericAttributeValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof GenericAttributeValue)) {
      return false;
    }
    GenericAttributeValue that = (GenericAttributeValue) o;
    return (getAttributeId() == that.getAttributeId())
        && (getValue() == that.getValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAttributeId(), getValue());
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
