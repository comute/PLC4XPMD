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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObject_SET_POINT_COMMAND_SCALED_VALUE extends InformationObject
    implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.SET_POINT_COMMAND_SCALED_VALUE;
  }

  // Properties.
  protected final ScaledValue sva;
  protected final QualifierOfSetPointCommand qos;

  public InformationObject_SET_POINT_COMMAND_SCALED_VALUE(
      int address, ScaledValue sva, QualifierOfSetPointCommand qos) {
    super(address);
    this.sva = sva;
    this.qos = qos;
  }

  public ScaledValue getSva() {
    return sva;
  }

  public QualifierOfSetPointCommand getQos() {
    return qos;
  }

  @Override
  protected void serializeInformationObjectChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObject_SET_POINT_COMMAND_SCALED_VALUE");

    // Simple Field (sva)
    writeSimpleField(
        "sva",
        sva,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (qos)
    writeSimpleField(
        "qos",
        qos,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObject_SET_POINT_COMMAND_SCALED_VALUE");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObject_SET_POINT_COMMAND_SCALED_VALUE _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (sva)
    lengthInBits += sva.getLengthInBits();

    // Simple field (qos)
    lengthInBits += qos.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectBuilder staticParseInformationObjectBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification) throws ParseException {
    readBuffer.pullContext("InformationObject_SET_POINT_COMMAND_SCALED_VALUE");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ScaledValue sva =
        readSimpleField(
            "sva",
            new DataReaderComplexDefault<>(() -> ScaledValue.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    QualifierOfSetPointCommand qos =
        readSimpleField(
            "qos",
            new DataReaderComplexDefault<>(
                () -> QualifierOfSetPointCommand.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObject_SET_POINT_COMMAND_SCALED_VALUE");
    // Create the instance
    return new InformationObject_SET_POINT_COMMAND_SCALED_VALUEBuilderImpl(sva, qos);
  }

  public static class InformationObject_SET_POINT_COMMAND_SCALED_VALUEBuilderImpl
      implements InformationObject.InformationObjectBuilder {
    private final ScaledValue sva;
    private final QualifierOfSetPointCommand qos;

    public InformationObject_SET_POINT_COMMAND_SCALED_VALUEBuilderImpl(
        ScaledValue sva, QualifierOfSetPointCommand qos) {
      this.sva = sva;
      this.qos = qos;
    }

    public InformationObject_SET_POINT_COMMAND_SCALED_VALUE build(int address) {
      InformationObject_SET_POINT_COMMAND_SCALED_VALUE
          informationObject_SET_POINT_COMMAND_SCALED_VALUE =
              new InformationObject_SET_POINT_COMMAND_SCALED_VALUE(address, sva, qos);
      return informationObject_SET_POINT_COMMAND_SCALED_VALUE;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObject_SET_POINT_COMMAND_SCALED_VALUE)) {
      return false;
    }
    InformationObject_SET_POINT_COMMAND_SCALED_VALUE that =
        (InformationObject_SET_POINT_COMMAND_SCALED_VALUE) o;
    return (getSva() == that.getSva()) && (getQos() == that.getQos()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSva(), getQos());
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