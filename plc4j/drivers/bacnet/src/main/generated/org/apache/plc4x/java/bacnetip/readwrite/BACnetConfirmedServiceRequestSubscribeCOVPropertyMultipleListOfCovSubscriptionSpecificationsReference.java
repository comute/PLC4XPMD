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

public
class BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
    implements Message {

  // Properties.
  protected final BACnetPropertyReferenceEnclosed monitoredProperty;
  protected final BACnetContextTagReal covIncrement;
  protected final BACnetContextTagBoolean timestamped;

  public
  BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(
      BACnetPropertyReferenceEnclosed monitoredProperty,
      BACnetContextTagReal covIncrement,
      BACnetContextTagBoolean timestamped) {
    super();
    this.monitoredProperty = monitoredProperty;
    this.covIncrement = covIncrement;
    this.timestamped = timestamped;
  }

  public BACnetPropertyReferenceEnclosed getMonitoredProperty() {
    return monitoredProperty;
  }

  public BACnetContextTagReal getCovIncrement() {
    return covIncrement;
  }

  public BACnetContextTagBoolean getTimestamped() {
    return timestamped;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext(
        "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference");

    // Simple Field (monitoredProperty)
    writeSimpleField("monitoredProperty", monitoredProperty, writeComplex(writeBuffer));

    // Optional Field (covIncrement) (Can be skipped, if the value is null)
    writeOptionalField("covIncrement", covIncrement, writeComplex(writeBuffer));

    // Simple Field (timestamped)
    writeSimpleField("timestamped", timestamped, writeComplex(writeBuffer));

    writeBuffer.popContext(
        "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
        _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (monitoredProperty)
    lengthInBits += monitoredProperty.getLengthInBits();

    // Optional Field (covIncrement)
    if (covIncrement != null) {
      lengthInBits += covIncrement.getLengthInBits();
    }

    // Simple field (timestamped)
    lengthInBits += timestamped.getLengthInBits();

    return lengthInBits;
  }

  public static
  BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
      staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext(
        "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetPropertyReferenceEnclosed monitoredProperty =
        readSimpleField(
            "monitoredProperty",
            readComplex(
                () -> BACnetPropertyReferenceEnclosed.staticParse(readBuffer, (short) (1)),
                readBuffer));

    BACnetContextTagReal covIncrement =
        readOptionalField(
            "covIncrement",
            readComplex(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (1), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagBoolean timestamped =
        readSimpleField(
            "timestamped",
            readComplex(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    readBuffer.closeContext(
        "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference");
    // Create the instance
    BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
        _bACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference;
    _bACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference =
        new BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(
            monitoredProperty, covIncrement, timestamped);
    return _bACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o
        instanceof
        BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference)) {
      return false;
    }
    BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
        that =
            (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference)
                o;
    return (getMonitoredProperty() == that.getMonitoredProperty())
        && (getCovIncrement() == that.getCovIncrement())
        && (getTimestamped() == that.getTimestamped())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getMonitoredProperty(), getCovIncrement(), getTimestamped());
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
