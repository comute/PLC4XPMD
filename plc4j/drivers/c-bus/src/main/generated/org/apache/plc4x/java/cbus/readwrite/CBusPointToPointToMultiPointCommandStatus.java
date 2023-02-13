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

public class CBusPointToPointToMultiPointCommandStatus extends CBusPointToPointToMultiPointCommand
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final StatusRequest statusRequest;

  // Arguments.
  protected final CBusOptions cBusOptions;
  // Reserved Fields
  private Byte reservedField0;

  public CBusPointToPointToMultiPointCommandStatus(
      BridgeAddress bridgeAddress,
      NetworkRoute networkRoute,
      byte peekedApplication,
      StatusRequest statusRequest,
      CBusOptions cBusOptions) {
    super(bridgeAddress, networkRoute, peekedApplication, cBusOptions);
    this.statusRequest = statusRequest;
    this.cBusOptions = cBusOptions;
  }

  public StatusRequest getStatusRequest() {
    return statusRequest;
  }

  @Override
  protected void serializeCBusPointToPointToMultiPointCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusPointToPointToMultiPointCommandStatus");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0xFF,
        writeByte(writeBuffer, 8));

    // Simple Field (statusRequest)
    writeSimpleField("statusRequest", statusRequest, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CBusPointToPointToMultiPointCommandStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusPointToPointToMultiPointCommandStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (statusRequest)
    lengthInBits += statusRequest.getLengthInBits();

    return lengthInBits;
  }

  public static CBusPointToPointToMultiPointCommandBuilder
      staticParseCBusPointToPointToMultiPointCommandBuilder(
          ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusPointToPointToMultiPointCommandStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0xFF);

    StatusRequest statusRequest =
        readSimpleField(
            "statusRequest",
            new DataReaderComplexDefault<>(
                () -> StatusRequest.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("CBusPointToPointToMultiPointCommandStatus");
    // Create the instance
    return new CBusPointToPointToMultiPointCommandStatusBuilderImpl(
        statusRequest, cBusOptions, reservedField0);
  }

  public static class CBusPointToPointToMultiPointCommandStatusBuilderImpl
      implements CBusPointToPointToMultiPointCommand.CBusPointToPointToMultiPointCommandBuilder {
    private final StatusRequest statusRequest;
    private final CBusOptions cBusOptions;
    private final Byte reservedField0;

    public CBusPointToPointToMultiPointCommandStatusBuilderImpl(
        StatusRequest statusRequest, CBusOptions cBusOptions, Byte reservedField0) {
      this.statusRequest = statusRequest;
      this.cBusOptions = cBusOptions;
      this.reservedField0 = reservedField0;
    }

    public CBusPointToPointToMultiPointCommandStatus build(
        BridgeAddress bridgeAddress,
        NetworkRoute networkRoute,
        byte peekedApplication,
        CBusOptions cBusOptions) {
      CBusPointToPointToMultiPointCommandStatus cBusPointToPointToMultiPointCommandStatus =
          new CBusPointToPointToMultiPointCommandStatus(
              bridgeAddress, networkRoute, peekedApplication, statusRequest, cBusOptions);
      cBusPointToPointToMultiPointCommandStatus.reservedField0 = reservedField0;
      return cBusPointToPointToMultiPointCommandStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusPointToPointToMultiPointCommandStatus)) {
      return false;
    }
    CBusPointToPointToMultiPointCommandStatus that = (CBusPointToPointToMultiPointCommandStatus) o;
    return (getStatusRequest() == that.getStatusRequest()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatusRequest());
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
