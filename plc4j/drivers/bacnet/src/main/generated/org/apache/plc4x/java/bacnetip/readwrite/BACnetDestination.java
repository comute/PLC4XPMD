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

public class BACnetDestination implements Message {

  // Properties.
  protected final BACnetDaysOfWeekTagged validDays;
  protected final BACnetApplicationTagTime fromTime;
  protected final BACnetApplicationTagTime toTime;
  protected final BACnetRecipient recipient;
  protected final BACnetApplicationTagUnsignedInteger processIdentifier;
  protected final BACnetApplicationTagBoolean issueConfirmedNotifications;
  protected final BACnetEventTransitionBitsTagged transitions;

  public BACnetDestination(
      BACnetDaysOfWeekTagged validDays,
      BACnetApplicationTagTime fromTime,
      BACnetApplicationTagTime toTime,
      BACnetRecipient recipient,
      BACnetApplicationTagUnsignedInteger processIdentifier,
      BACnetApplicationTagBoolean issueConfirmedNotifications,
      BACnetEventTransitionBitsTagged transitions) {
    super();
    this.validDays = validDays;
    this.fromTime = fromTime;
    this.toTime = toTime;
    this.recipient = recipient;
    this.processIdentifier = processIdentifier;
    this.issueConfirmedNotifications = issueConfirmedNotifications;
    this.transitions = transitions;
  }

  public BACnetDaysOfWeekTagged getValidDays() {
    return validDays;
  }

  public BACnetApplicationTagTime getFromTime() {
    return fromTime;
  }

  public BACnetApplicationTagTime getToTime() {
    return toTime;
  }

  public BACnetRecipient getRecipient() {
    return recipient;
  }

  public BACnetApplicationTagUnsignedInteger getProcessIdentifier() {
    return processIdentifier;
  }

  public BACnetApplicationTagBoolean getIssueConfirmedNotifications() {
    return issueConfirmedNotifications;
  }

  public BACnetEventTransitionBitsTagged getTransitions() {
    return transitions;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetDestination");

    // Simple Field (validDays)
    writeSimpleField("validDays", validDays, writeComplex(writeBuffer));

    // Simple Field (fromTime)
    writeSimpleField("fromTime", fromTime, writeComplex(writeBuffer));

    // Simple Field (toTime)
    writeSimpleField("toTime", toTime, writeComplex(writeBuffer));

    // Simple Field (recipient)
    writeSimpleField("recipient", recipient, writeComplex(writeBuffer));

    // Simple Field (processIdentifier)
    writeSimpleField("processIdentifier", processIdentifier, writeComplex(writeBuffer));

    // Simple Field (issueConfirmedNotifications)
    writeSimpleField(
        "issueConfirmedNotifications", issueConfirmedNotifications, writeComplex(writeBuffer));

    // Simple Field (transitions)
    writeSimpleField("transitions", transitions, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetDestination");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetDestination _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (validDays)
    lengthInBits += validDays.getLengthInBits();

    // Simple field (fromTime)
    lengthInBits += fromTime.getLengthInBits();

    // Simple field (toTime)
    lengthInBits += toTime.getLengthInBits();

    // Simple field (recipient)
    lengthInBits += recipient.getLengthInBits();

    // Simple field (processIdentifier)
    lengthInBits += processIdentifier.getLengthInBits();

    // Simple field (issueConfirmedNotifications)
    lengthInBits += issueConfirmedNotifications.getLengthInBits();

    // Simple field (transitions)
    lengthInBits += transitions.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetDestination staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetDestination");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetDaysOfWeekTagged validDays =
        readSimpleField(
            "validDays",
            readComplex(
                () ->
                    BACnetDaysOfWeekTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));

    BACnetApplicationTagTime fromTime =
        readSimpleField(
            "fromTime",
            readComplex(
                () -> (BACnetApplicationTagTime) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagTime toTime =
        readSimpleField(
            "toTime",
            readComplex(
                () -> (BACnetApplicationTagTime) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetRecipient recipient =
        readSimpleField(
            "recipient", readComplex(() -> BACnetRecipient.staticParse(readBuffer), readBuffer));

    BACnetApplicationTagUnsignedInteger processIdentifier =
        readSimpleField(
            "processIdentifier",
            readComplex(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetApplicationTagBoolean issueConfirmedNotifications =
        readSimpleField(
            "issueConfirmedNotifications",
            readComplex(
                () -> (BACnetApplicationTagBoolean) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    BACnetEventTransitionBitsTagged transitions =
        readSimpleField(
            "transitions",
            readComplex(
                () ->
                    BACnetEventTransitionBitsTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetDestination");
    // Create the instance
    BACnetDestination _bACnetDestination;
    _bACnetDestination =
        new BACnetDestination(
            validDays,
            fromTime,
            toTime,
            recipient,
            processIdentifier,
            issueConfirmedNotifications,
            transitions);
    return _bACnetDestination;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetDestination)) {
      return false;
    }
    BACnetDestination that = (BACnetDestination) o;
    return (getValidDays() == that.getValidDays())
        && (getFromTime() == that.getFromTime())
        && (getToTime() == that.getToTime())
        && (getRecipient() == that.getRecipient())
        && (getProcessIdentifier() == that.getProcessIdentifier())
        && (getIssueConfirmedNotifications() == that.getIssueConfirmedNotifications())
        && (getTransitions() == that.getTransitions())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getValidDays(),
        getFromTime(),
        getToTime(),
        getRecipient(),
        getProcessIdentifier(),
        getIssueConfirmedNotifications(),
        getTransitions());
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
