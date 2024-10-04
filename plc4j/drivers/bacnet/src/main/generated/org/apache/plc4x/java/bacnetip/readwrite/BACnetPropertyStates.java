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

public abstract class BACnetPropertyStates implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetTagHeader peekedTagHeader;

  public BACnetPropertyStates(BACnetTagHeader peekedTagHeader) {
    super();
    this.peekedTagHeader = peekedTagHeader;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  protected abstract void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyStates");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Switch field (Serialize the sub-type)
    serializeBACnetPropertyStatesChild(writeBuffer);

    writeBuffer.popContext("BACnetPropertyStates");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetPropertyStates _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetPropertyStates staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStates");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            readComplex(() -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetPropertyStatesBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0)) {
      builder =
          BACnetPropertyStatesBoolean.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 1)) {
      builder =
          BACnetPropertyStatesBinaryValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 2)) {
      builder =
          BACnetPropertyStatesEventType.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 3)) {
      builder =
          BACnetPropertyStatesPolarity.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 4)) {
      builder =
          BACnetPropertyStatesProgramChange.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 5)) {
      builder =
          BACnetPropertyStatesProgramChange.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 6)) {
      builder =
          BACnetPropertyStatesReasonForHalt.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 7)) {
      builder =
          BACnetPropertyStatesReliability.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 8)) {
      builder =
          BACnetPropertyStatesState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 9)) {
      builder =
          BACnetPropertyStatesSystemStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 10)) {
      builder =
          BACnetPropertyStatesUnits.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 11)) {
      builder =
          BACnetPropertyStatesExtendedValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 12)) {
      builder =
          BACnetPropertyStatesLifeSafetyMode.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 13)) {
      builder =
          BACnetPropertyStatesLifeSafetyState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 14)) {
      builder =
          BACnetPropertyStatesRestartReason.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 15)) {
      builder =
          BACnetPropertyStatesDoorAlarmState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 16)) {
      builder =
          BACnetPropertyStatesAction.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 17)) {
      builder =
          BACnetPropertyStatesDoorSecuredStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 18)) {
      builder =
          BACnetPropertyStatesDoorStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 19)) {
      builder =
          BACnetPropertyStatesDoorValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 20)) {
      builder =
          BACnetPropertyStatesFileAccessMethod.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 21)) {
      builder =
          BACnetPropertyStatesLockStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 22)) {
      builder =
          BACnetPropertyStatesLifeSafetyOperations.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 23)) {
      builder =
          BACnetPropertyStatesMaintenance.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 24)) {
      builder =
          BACnetPropertyStatesNodeType.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 25)) {
      builder =
          BACnetPropertyStatesNotifyType.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 26)) {
      builder =
          BACnetPropertyStatesSecurityLevel.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 27)) {
      builder =
          BACnetPropertyStatesShedState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 28)) {
      builder =
          BACnetPropertyStatesSilencedState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 30)) {
      builder =
          BACnetPropertyStatesAccessEvent.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 31)) {
      builder =
          BACnetPropertyStatesZoneOccupanyState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 32)) {
      builder =
          BACnetPropertyStatesAccessCredentialDisableReason.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 33)) {
      builder =
          BACnetPropertyStatesAccessCredentialDisable.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 34)) {
      builder =
          BACnetPropertyStatesAuthenticationStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 36)) {
      builder =
          BACnetPropertyStatesBackupState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 37)) {
      builder =
          BACnetPropertyStatesWriteStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 38)) {
      builder =
          BACnetPropertyStatesLightningInProgress.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 39)) {
      builder =
          BACnetPropertyStatesLightningOperation.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 40)) {
      builder =
          BACnetPropertyStatesLightningTransition.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 41)) {
      builder =
          BACnetPropertyStatesIntegerValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 42)) {
      builder =
          BACnetPropertyStatesBinaryLightningValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 43)) {
      builder =
          BACnetPropertyStatesTimerState.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 44)) {
      builder =
          BACnetPropertyStatesTimerTransition.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 45)) {
      builder =
          BACnetPropertyStatesBacnetIpMode.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 46)) {
      builder =
          BACnetPropertyStatesNetworkPortCommand.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 47)) {
      builder =
          BACnetPropertyStatesNetworkType.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 48)) {
      builder =
          BACnetPropertyStatesNetworkNumberQuality.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 49)) {
      builder =
          BACnetPropertyStatesEscalatorOperationDirection.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 50)) {
      builder =
          BACnetPropertyStatesEscalatorFault.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 51)) {
      builder =
          BACnetPropertyStatesEscalatorMode.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 52)) {
      builder =
          BACnetPropertyStatesLiftCarDirection.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 53)) {
      builder =
          BACnetPropertyStatesLiftCarDoorCommand.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 54)) {
      builder =
          BACnetPropertyStatesLiftCarDriveStatus.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 55)) {
      builder =
          BACnetPropertyStatesLiftCarMode.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 56)) {
      builder =
          BACnetPropertyStatesLiftGroupMode.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 57)) {
      builder =
          BACnetPropertyStatesLiftFault.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 58)) {
      builder =
          BACnetPropertyStatesProtocolLevel.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 63)) {
      builder =
          BACnetPropertyStatesExtendedValue.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    } else {
      builder =
          BACnetPropertyStateActionUnknown.staticParseBACnetPropertyStatesBuilder(
              readBuffer, peekedTagNumber);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + "]");
    }

    readBuffer.closeContext("BACnetPropertyStates");
    // Create the instance
    BACnetPropertyStates _bACnetPropertyStates = builder.build(peekedTagHeader);
    return _bACnetPropertyStates;
  }

  public interface BACnetPropertyStatesBuilder {
    BACnetPropertyStates build(BACnetTagHeader peekedTagHeader);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStates)) {
      return false;
    }
    BACnetPropertyStates that = (BACnetPropertyStates) o;
    return (getPeekedTagHeader() == that.getPeekedTagHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPeekedTagHeader());
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
