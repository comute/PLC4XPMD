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

public class BACnetLightingCommand implements Message {

  // Properties.
  protected final BACnetLightingOperationTagged lightningOperation;
  protected final BACnetContextTagReal targetLevel;
  protected final BACnetContextTagReal rampRate;
  protected final BACnetContextTagReal stepIncrement;
  protected final BACnetContextTagUnsignedInteger fadeTime;
  protected final BACnetContextTagUnsignedInteger priority;

  public BACnetLightingCommand(
      BACnetLightingOperationTagged lightningOperation,
      BACnetContextTagReal targetLevel,
      BACnetContextTagReal rampRate,
      BACnetContextTagReal stepIncrement,
      BACnetContextTagUnsignedInteger fadeTime,
      BACnetContextTagUnsignedInteger priority) {
    super();
    this.lightningOperation = lightningOperation;
    this.targetLevel = targetLevel;
    this.rampRate = rampRate;
    this.stepIncrement = stepIncrement;
    this.fadeTime = fadeTime;
    this.priority = priority;
  }

  public BACnetLightingOperationTagged getLightningOperation() {
    return lightningOperation;
  }

  public BACnetContextTagReal getTargetLevel() {
    return targetLevel;
  }

  public BACnetContextTagReal getRampRate() {
    return rampRate;
  }

  public BACnetContextTagReal getStepIncrement() {
    return stepIncrement;
  }

  public BACnetContextTagUnsignedInteger getFadeTime() {
    return fadeTime;
  }

  public BACnetContextTagUnsignedInteger getPriority() {
    return priority;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLightingCommand");

    // Simple Field (lightningOperation)
    writeSimpleField("lightningOperation", lightningOperation, writeComplex(writeBuffer));

    // Optional Field (targetLevel) (Can be skipped, if the value is null)
    writeOptionalField("targetLevel", targetLevel, writeComplex(writeBuffer));

    // Optional Field (rampRate) (Can be skipped, if the value is null)
    writeOptionalField("rampRate", rampRate, writeComplex(writeBuffer));

    // Optional Field (stepIncrement) (Can be skipped, if the value is null)
    writeOptionalField("stepIncrement", stepIncrement, writeComplex(writeBuffer));

    // Optional Field (fadeTime) (Can be skipped, if the value is null)
    writeOptionalField("fadeTime", fadeTime, writeComplex(writeBuffer));

    // Optional Field (priority) (Can be skipped, if the value is null)
    writeOptionalField("priority", priority, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetLightingCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetLightingCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (lightningOperation)
    lengthInBits += lightningOperation.getLengthInBits();

    // Optional Field (targetLevel)
    if (targetLevel != null) {
      lengthInBits += targetLevel.getLengthInBits();
    }

    // Optional Field (rampRate)
    if (rampRate != null) {
      lengthInBits += rampRate.getLengthInBits();
    }

    // Optional Field (stepIncrement)
    if (stepIncrement != null) {
      lengthInBits += stepIncrement.getLengthInBits();
    }

    // Optional Field (fadeTime)
    if (fadeTime != null) {
      lengthInBits += fadeTime.getLengthInBits();
    }

    // Optional Field (priority)
    if (priority != null) {
      lengthInBits += priority.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetLightingCommand staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLightingCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetLightingOperationTagged lightningOperation =
        readSimpleField(
            "lightningOperation",
            readComplex(
                () ->
                    BACnetLightingOperationTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagReal targetLevel =
        readOptionalField(
            "targetLevel",
            readComplex(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (1), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagReal rampRate =
        readOptionalField(
            "rampRate",
            readComplex(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagReal stepIncrement =
        readOptionalField(
            "stepIncrement",
            readComplex(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (3), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagUnsignedInteger fadeTime =
        readOptionalField(
            "fadeTime",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (4),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger priority =
        readOptionalField(
            "priority",
            readComplex(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (5),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetLightingCommand");
    // Create the instance
    BACnetLightingCommand _bACnetLightingCommand;
    _bACnetLightingCommand =
        new BACnetLightingCommand(
            lightningOperation, targetLevel, rampRate, stepIncrement, fadeTime, priority);
    return _bACnetLightingCommand;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLightingCommand)) {
      return false;
    }
    BACnetLightingCommand that = (BACnetLightingCommand) o;
    return (getLightningOperation() == that.getLightningOperation())
        && (getTargetLevel() == that.getTargetLevel())
        && (getRampRate() == that.getRampRate())
        && (getStepIncrement() == that.getStepIncrement())
        && (getFadeTime() == that.getFadeTime())
        && (getPriority() == that.getPriority())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getLightningOperation(),
        getTargetLevel(),
        getRampRate(),
        getStepIncrement(),
        getFadeTime(),
        getPriority());
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
