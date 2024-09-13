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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetErrorGeneral extends BACnetError implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getErrorChoice() {
    return null;
  }

  // Properties.
  protected final Error error;

  public BACnetErrorGeneral(Error error) {
    super();
    this.error = error;
  }

  public Error getError() {
    return error;
  }

  @Override
  protected void serializeBACnetErrorChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetErrorGeneral");

    // Simple Field (error)
    writeSimpleField("error", error, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetErrorGeneral");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetErrorGeneral _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (error)
    lengthInBits += error.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetErrorBuilder staticParseBACnetErrorBuilder(
      ReadBuffer readBuffer, BACnetConfirmedServiceChoice errorChoice) throws ParseException {
    readBuffer.pullContext("BACnetErrorGeneral");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Error error =
        readSimpleField("error", readComplex(() -> Error.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("BACnetErrorGeneral");
    // Create the instance
    return new BACnetErrorGeneralBuilderImpl(error);
  }

  public static class BACnetErrorGeneralBuilderImpl implements BACnetError.BACnetErrorBuilder {
    private final Error error;

    public BACnetErrorGeneralBuilderImpl(Error error) {
      this.error = error;
    }

    public BACnetErrorGeneral build() {
      BACnetErrorGeneral bACnetErrorGeneral = new BACnetErrorGeneral(error);
      return bACnetErrorGeneral;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetErrorGeneral)) {
      return false;
    }
    BACnetErrorGeneral that = (BACnetErrorGeneral) o;
    return (getError() == that.getError()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getError());
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
