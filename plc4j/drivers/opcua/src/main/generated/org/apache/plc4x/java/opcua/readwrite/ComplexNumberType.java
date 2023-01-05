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

public class ComplexNumberType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "12173";
  }

  // Properties.
  protected final float real;
  protected final float imaginary;

  public ComplexNumberType(float real, float imaginary) {
    super();
    this.real = real;
    this.imaginary = imaginary;
  }

  public float getReal() {
    return real;
  }

  public float getImaginary() {
    return imaginary;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ComplexNumberType");

    // Simple Field (real)
    writeSimpleField("real", real, writeFloat(writeBuffer, 32));

    // Simple Field (imaginary)
    writeSimpleField("imaginary", imaginary, writeFloat(writeBuffer, 32));

    writeBuffer.popContext("ComplexNumberType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ComplexNumberType _value = this;

    // Simple field (real)
    lengthInBits += 32;

    // Simple field (imaginary)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ComplexNumberTypeBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ComplexNumberType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    float real = readSimpleField("real", readFloat(readBuffer, 32));

    float imaginary = readSimpleField("imaginary", readFloat(readBuffer, 32));

    readBuffer.closeContext("ComplexNumberType");
    // Create the instance
    return new ComplexNumberTypeBuilder(real, imaginary);
  }

  public static class ComplexNumberTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final float real;
    private final float imaginary;

    public ComplexNumberTypeBuilder(float real, float imaginary) {

      this.real = real;
      this.imaginary = imaginary;
    }

    public ComplexNumberType build() {
      ComplexNumberType complexNumberType = new ComplexNumberType(real, imaginary);
      return complexNumberType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ComplexNumberType)) {
      return false;
    }
    ComplexNumberType that = (ComplexNumberType) o;
    return (getReal() == that.getReal())
        && (getImaginary() == that.getImaginary())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReal(), getImaginary());
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
