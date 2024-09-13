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
package org.apache.plc4x.java.cbus.readwrite;

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

public class RequestObsolete extends Request implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final CALData calData;
  protected final Alpha alpha;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public RequestObsolete(
      RequestType peekedByte,
      RequestType startingCR,
      RequestType resetMode,
      RequestType secondPeek,
      RequestTermination termination,
      CALData calData,
      Alpha alpha,
      CBusOptions cBusOptions) {
    super(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions);
    this.calData = calData;
    this.alpha = alpha;
    this.cBusOptions = cBusOptions;
  }

  public CALData getCalData() {
    return calData;
  }

  public Alpha getAlpha() {
    return alpha;
  }

  public CALData getCalDataDecoded() {
    return (CALData) (getCalData());
  }

  @Override
  protected void serializeRequestChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("RequestObsolete");

    // Manual Field (calData)
    writeManualField(
        "calData",
        () ->
            org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper.writeCALData(
                writeBuffer, calData),
        writeBuffer);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    CALData calDataDecoded = getCalDataDecoded();
    writeBuffer.writeVirtual("calDataDecoded", calDataDecoded);

    // Optional Field (alpha) (Can be skipped, if the value is null)
    writeOptionalField("alpha", alpha, writeComplex(writeBuffer));

    writeBuffer.popContext("RequestObsolete");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    RequestObsolete _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Manual Field (calData)
    lengthInBits += (((calData.getLengthInBytes()) * (2))) * (8);

    // A virtual field doesn't have any in- or output.

    // Optional Field (alpha)
    if (alpha != null) {
      lengthInBits += alpha.getLengthInBits();
    }

    return lengthInBits;
  }

  public static RequestBuilder staticParseRequestBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("RequestObsolete");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    CALData calData =
        readManualField(
            "calData",
            readBuffer,
            () ->
                (CALData)
                    (org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper.readCALData(
                        readBuffer)));
    CALData calDataDecoded = readVirtualField("calDataDecoded", CALData.class, calData);

    Alpha alpha =
        readOptionalField("alpha", readComplex(() -> Alpha.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("RequestObsolete");
    // Create the instance
    return new RequestObsoleteBuilderImpl(calData, alpha, cBusOptions);
  }

  public static class RequestObsoleteBuilderImpl implements Request.RequestBuilder {
    private final CALData calData;
    private final Alpha alpha;
    private final CBusOptions cBusOptions;

    public RequestObsoleteBuilderImpl(CALData calData, Alpha alpha, CBusOptions cBusOptions) {
      this.calData = calData;
      this.alpha = alpha;
      this.cBusOptions = cBusOptions;
    }

    public RequestObsolete build(
        RequestType peekedByte,
        RequestType startingCR,
        RequestType resetMode,
        RequestType secondPeek,
        RequestTermination termination,
        CBusOptions cBusOptions) {
      RequestObsolete requestObsolete =
          new RequestObsolete(
              peekedByte,
              startingCR,
              resetMode,
              secondPeek,
              termination,
              calData,
              alpha,
              cBusOptions);
      return requestObsolete;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof RequestObsolete)) {
      return false;
    }
    RequestObsolete that = (RequestObsolete) o;
    return (getCalData() == that.getCalData())
        && (getAlpha() == that.getAlpha())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCalData(), getAlpha());
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
