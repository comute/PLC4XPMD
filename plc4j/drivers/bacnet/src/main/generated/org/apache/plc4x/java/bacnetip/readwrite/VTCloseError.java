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

public class VTCloseError extends BACnetError implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getErrorChoice() {
    return BACnetConfirmedServiceChoice.VT_CLOSE;
  }

  // Properties.
  protected final ErrorEnclosed errorType;
  protected final VTCloseErrorListOfVTSessionIdentifiers listOfVtSessionIdentifiers;

  public VTCloseError(
      ErrorEnclosed errorType, VTCloseErrorListOfVTSessionIdentifiers listOfVtSessionIdentifiers) {
    super();
    this.errorType = errorType;
    this.listOfVtSessionIdentifiers = listOfVtSessionIdentifiers;
  }

  public ErrorEnclosed getErrorType() {
    return errorType;
  }

  public VTCloseErrorListOfVTSessionIdentifiers getListOfVtSessionIdentifiers() {
    return listOfVtSessionIdentifiers;
  }

  @Override
  protected void serializeBACnetErrorChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("VTCloseError");

    // Simple Field (errorType)
    writeSimpleField("errorType", errorType, writeComplex(writeBuffer));

    // Optional Field (listOfVtSessionIdentifiers) (Can be skipped, if the value is null)
    writeOptionalField(
        "listOfVtSessionIdentifiers", listOfVtSessionIdentifiers, writeComplex(writeBuffer));

    writeBuffer.popContext("VTCloseError");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    VTCloseError _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (errorType)
    lengthInBits += errorType.getLengthInBits();

    // Optional Field (listOfVtSessionIdentifiers)
    if (listOfVtSessionIdentifiers != null) {
      lengthInBits += listOfVtSessionIdentifiers.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetErrorBuilder staticParseBACnetErrorBuilder(
      ReadBuffer readBuffer, BACnetConfirmedServiceChoice errorChoice) throws ParseException {
    readBuffer.pullContext("VTCloseError");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ErrorEnclosed errorType =
        readSimpleField(
            "errorType",
            readComplex(() -> ErrorEnclosed.staticParse(readBuffer, (short) (0)), readBuffer));

    VTCloseErrorListOfVTSessionIdentifiers listOfVtSessionIdentifiers =
        readOptionalField(
            "listOfVtSessionIdentifiers",
            readComplex(
                () -> VTCloseErrorListOfVTSessionIdentifiers.staticParse(readBuffer, (short) (1)),
                readBuffer));

    readBuffer.closeContext("VTCloseError");
    // Create the instance
    return new VTCloseErrorBuilderImpl(errorType, listOfVtSessionIdentifiers);
  }

  public static class VTCloseErrorBuilderImpl implements BACnetError.BACnetErrorBuilder {
    private final ErrorEnclosed errorType;
    private final VTCloseErrorListOfVTSessionIdentifiers listOfVtSessionIdentifiers;

    public VTCloseErrorBuilderImpl(
        ErrorEnclosed errorType,
        VTCloseErrorListOfVTSessionIdentifiers listOfVtSessionIdentifiers) {
      this.errorType = errorType;
      this.listOfVtSessionIdentifiers = listOfVtSessionIdentifiers;
    }

    public VTCloseError build() {
      VTCloseError vTCloseError = new VTCloseError(errorType, listOfVtSessionIdentifiers);
      return vTCloseError;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof VTCloseError)) {
      return false;
    }
    VTCloseError that = (VTCloseError) o;
    return (getErrorType() == that.getErrorType())
        && (getListOfVtSessionIdentifiers() == that.getListOfVtSessionIdentifiers())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getErrorType(), getListOfVtSessionIdentifiers());
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
