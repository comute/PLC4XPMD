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

public abstract class BACnetServiceAck implements Message {

  // Abstract accessors for discriminator values.
  public abstract BACnetConfirmedServiceChoice getServiceChoice();

  // Arguments.
  protected final Long serviceAckLength;

  public BACnetServiceAck(Long serviceAckLength) {
    super();
    this.serviceAckLength = serviceAckLength;
  }

  public long getServiceAckPayloadLength() {
    return (long) (((((serviceAckLength) > (0))) ? ((serviceAckLength) - (1L)) : 0L));
  }

  protected abstract void serializeBACnetServiceAckChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetServiceAck");

    // Discriminator Field (serviceChoice) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "serviceChoice",
        "BACnetConfirmedServiceChoice",
        getServiceChoice(),
        new DataWriterEnumDefault<>(
            BACnetConfirmedServiceChoice::getValue,
            BACnetConfirmedServiceChoice::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    long serviceAckPayloadLength = getServiceAckPayloadLength();
    writeBuffer.writeVirtual("serviceAckPayloadLength", serviceAckPayloadLength);

    // Switch field (Serialize the sub-type)
    serializeBACnetServiceAckChild(writeBuffer);

    writeBuffer.popContext("BACnetServiceAck");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetServiceAck _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (serviceChoice)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetServiceAck staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Long serviceAckLength;
    if (args[0] instanceof Long) {
      serviceAckLength = (Long) args[0];
    } else if (args[0] instanceof String) {
      serviceAckLength = Long.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Long or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, serviceAckLength);
  }

  public static BACnetServiceAck staticParse(ReadBuffer readBuffer, Long serviceAckLength)
      throws ParseException {
    readBuffer.pullContext("BACnetServiceAck");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetConfirmedServiceChoice serviceChoice =
        readDiscriminatorEnumField(
            "serviceChoice",
            "BACnetConfirmedServiceChoice",
            new DataReaderEnumDefault<>(
                BACnetConfirmedServiceChoice::enumForValue, readUnsignedShort(readBuffer, 8)));
    long serviceAckPayloadLength =
        readVirtualField(
            "serviceAckPayloadLength",
            long.class,
            ((((serviceAckLength) > (0))) ? ((serviceAckLength) - (1L)) : 0L));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetServiceAckBuilder builder = null;
    if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.GET_ALARM_SUMMARY)) {
      builder =
          BACnetServiceAckGetAlarmSummary.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.GET_ENROLLMENT_SUMMARY)) {
      builder =
          BACnetServiceAckGetEnrollmentSummary.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.GET_EVENT_INFORMATION)) {
      builder =
          BACnetServiceAckGetEventInformation.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.ATOMIC_READ_FILE)) {
      builder =
          BACnetServiceAckAtomicReadFile.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.ATOMIC_WRITE_FILE)) {
      builder =
          BACnetServiceAckAtomicWriteFile.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.CREATE_OBJECT)) {
      builder =
          BACnetServiceAckCreateObject.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.READ_PROPERTY)) {
      builder =
          BACnetServiceAckReadProperty.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.READ_PROPERTY_MULTIPLE)) {
      builder =
          BACnetServiceAckReadPropertyMultiple.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckPayloadLength, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.READ_RANGE)) {
      builder =
          BACnetServiceAckReadRange.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.CONFIRMED_PRIVATE_TRANSFER)) {
      builder =
          BACnetServiceAckConfirmedPrivateTransfer.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.VT_OPEN)) {
      builder =
          BACnetServiceAckVTOpen.staticParseBACnetServiceAckBuilder(readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.VT_DATA)) {
      builder =
          BACnetServiceAckVTData.staticParseBACnetServiceAckBuilder(readBuffer, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.AUTHENTICATE)) {
      builder =
          BACnetServiceAckAuthenticate.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckPayloadLength, serviceAckLength);
    } else if (EvaluationHelper.equals(serviceChoice, BACnetConfirmedServiceChoice.REQUEST_KEY)) {
      builder =
          BACnetServiceAckRequestKey.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckPayloadLength, serviceAckLength);
    } else if (EvaluationHelper.equals(
        serviceChoice, BACnetConfirmedServiceChoice.READ_PROPERTY_CONDITIONAL)) {
      builder =
          BACnetServiceAckReadPropertyConditional.staticParseBACnetServiceAckBuilder(
              readBuffer, serviceAckPayloadLength, serviceAckLength);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "serviceChoice="
              + serviceChoice
              + "]");
    }

    readBuffer.closeContext("BACnetServiceAck");
    // Create the instance
    BACnetServiceAck _bACnetServiceAck = builder.build(serviceAckLength);

    return _bACnetServiceAck;
  }

  public interface BACnetServiceAckBuilder {
    BACnetServiceAck build(Long serviceAckLength);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetServiceAck)) {
      return false;
    }
    BACnetServiceAck that = (BACnetServiceAck) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
