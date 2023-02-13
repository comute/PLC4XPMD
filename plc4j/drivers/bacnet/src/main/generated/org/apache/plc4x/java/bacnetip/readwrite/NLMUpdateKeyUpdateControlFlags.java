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

public class NLMUpdateKeyUpdateControlFlags implements Message {

  // Properties.
  protected final boolean set1KeyRevisionActivationTimeExpirationTimePresent;
  protected final boolean set1KeyCountKeyParametersPresent;
  protected final boolean set1ShouldBeCleared;
  protected final boolean set2KeyRevisionActivationTimeExpirationTimePresent;
  protected final boolean set2KeyCountKeyParametersPresent;
  protected final boolean set2ShouldBeCleared;
  protected final boolean moreMessagesToBeExpected;
  protected final boolean removeAllKeys;

  public NLMUpdateKeyUpdateControlFlags(
      boolean set1KeyRevisionActivationTimeExpirationTimePresent,
      boolean set1KeyCountKeyParametersPresent,
      boolean set1ShouldBeCleared,
      boolean set2KeyRevisionActivationTimeExpirationTimePresent,
      boolean set2KeyCountKeyParametersPresent,
      boolean set2ShouldBeCleared,
      boolean moreMessagesToBeExpected,
      boolean removeAllKeys) {
    super();
    this.set1KeyRevisionActivationTimeExpirationTimePresent =
        set1KeyRevisionActivationTimeExpirationTimePresent;
    this.set1KeyCountKeyParametersPresent = set1KeyCountKeyParametersPresent;
    this.set1ShouldBeCleared = set1ShouldBeCleared;
    this.set2KeyRevisionActivationTimeExpirationTimePresent =
        set2KeyRevisionActivationTimeExpirationTimePresent;
    this.set2KeyCountKeyParametersPresent = set2KeyCountKeyParametersPresent;
    this.set2ShouldBeCleared = set2ShouldBeCleared;
    this.moreMessagesToBeExpected = moreMessagesToBeExpected;
    this.removeAllKeys = removeAllKeys;
  }

  public boolean getSet1KeyRevisionActivationTimeExpirationTimePresent() {
    return set1KeyRevisionActivationTimeExpirationTimePresent;
  }

  public boolean getSet1KeyCountKeyParametersPresent() {
    return set1KeyCountKeyParametersPresent;
  }

  public boolean getSet1ShouldBeCleared() {
    return set1ShouldBeCleared;
  }

  public boolean getSet2KeyRevisionActivationTimeExpirationTimePresent() {
    return set2KeyRevisionActivationTimeExpirationTimePresent;
  }

  public boolean getSet2KeyCountKeyParametersPresent() {
    return set2KeyCountKeyParametersPresent;
  }

  public boolean getSet2ShouldBeCleared() {
    return set2ShouldBeCleared;
  }

  public boolean getMoreMessagesToBeExpected() {
    return moreMessagesToBeExpected;
  }

  public boolean getRemoveAllKeys() {
    return removeAllKeys;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NLMUpdateKeyUpdateControlFlags");

    // Simple Field (set1KeyRevisionActivationTimeExpirationTimePresent)
    writeSimpleField(
        "set1KeyRevisionActivationTimeExpirationTimePresent",
        set1KeyRevisionActivationTimeExpirationTimePresent,
        writeBoolean(writeBuffer));

    // Simple Field (set1KeyCountKeyParametersPresent)
    writeSimpleField(
        "set1KeyCountKeyParametersPresent",
        set1KeyCountKeyParametersPresent,
        writeBoolean(writeBuffer));

    // Simple Field (set1ShouldBeCleared)
    writeSimpleField("set1ShouldBeCleared", set1ShouldBeCleared, writeBoolean(writeBuffer));

    // Simple Field (set2KeyRevisionActivationTimeExpirationTimePresent)
    writeSimpleField(
        "set2KeyRevisionActivationTimeExpirationTimePresent",
        set2KeyRevisionActivationTimeExpirationTimePresent,
        writeBoolean(writeBuffer));

    // Simple Field (set2KeyCountKeyParametersPresent)
    writeSimpleField(
        "set2KeyCountKeyParametersPresent",
        set2KeyCountKeyParametersPresent,
        writeBoolean(writeBuffer));

    // Simple Field (set2ShouldBeCleared)
    writeSimpleField("set2ShouldBeCleared", set2ShouldBeCleared, writeBoolean(writeBuffer));

    // Simple Field (moreMessagesToBeExpected)
    writeSimpleField(
        "moreMessagesToBeExpected", moreMessagesToBeExpected, writeBoolean(writeBuffer));

    // Simple Field (removeAllKeys)
    writeSimpleField("removeAllKeys", removeAllKeys, writeBoolean(writeBuffer));

    writeBuffer.popContext("NLMUpdateKeyUpdateControlFlags");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    NLMUpdateKeyUpdateControlFlags _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (set1KeyRevisionActivationTimeExpirationTimePresent)
    lengthInBits += 1;

    // Simple field (set1KeyCountKeyParametersPresent)
    lengthInBits += 1;

    // Simple field (set1ShouldBeCleared)
    lengthInBits += 1;

    // Simple field (set2KeyRevisionActivationTimeExpirationTimePresent)
    lengthInBits += 1;

    // Simple field (set2KeyCountKeyParametersPresent)
    lengthInBits += 1;

    // Simple field (set2ShouldBeCleared)
    lengthInBits += 1;

    // Simple field (moreMessagesToBeExpected)
    lengthInBits += 1;

    // Simple field (removeAllKeys)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static NLMUpdateKeyUpdateControlFlags staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static NLMUpdateKeyUpdateControlFlags staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("NLMUpdateKeyUpdateControlFlags");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean set1KeyRevisionActivationTimeExpirationTimePresent =
        readSimpleField(
            "set1KeyRevisionActivationTimeExpirationTimePresent", readBoolean(readBuffer));

    boolean set1KeyCountKeyParametersPresent =
        readSimpleField("set1KeyCountKeyParametersPresent", readBoolean(readBuffer));

    boolean set1ShouldBeCleared = readSimpleField("set1ShouldBeCleared", readBoolean(readBuffer));

    boolean set2KeyRevisionActivationTimeExpirationTimePresent =
        readSimpleField(
            "set2KeyRevisionActivationTimeExpirationTimePresent", readBoolean(readBuffer));

    boolean set2KeyCountKeyParametersPresent =
        readSimpleField("set2KeyCountKeyParametersPresent", readBoolean(readBuffer));

    boolean set2ShouldBeCleared = readSimpleField("set2ShouldBeCleared", readBoolean(readBuffer));

    boolean moreMessagesToBeExpected =
        readSimpleField("moreMessagesToBeExpected", readBoolean(readBuffer));

    boolean removeAllKeys = readSimpleField("removeAllKeys", readBoolean(readBuffer));

    readBuffer.closeContext("NLMUpdateKeyUpdateControlFlags");
    // Create the instance
    NLMUpdateKeyUpdateControlFlags _nLMUpdateKeyUpdateControlFlags;
    _nLMUpdateKeyUpdateControlFlags =
        new NLMUpdateKeyUpdateControlFlags(
            set1KeyRevisionActivationTimeExpirationTimePresent,
            set1KeyCountKeyParametersPresent,
            set1ShouldBeCleared,
            set2KeyRevisionActivationTimeExpirationTimePresent,
            set2KeyCountKeyParametersPresent,
            set2ShouldBeCleared,
            moreMessagesToBeExpected,
            removeAllKeys);
    return _nLMUpdateKeyUpdateControlFlags;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMUpdateKeyUpdateControlFlags)) {
      return false;
    }
    NLMUpdateKeyUpdateControlFlags that = (NLMUpdateKeyUpdateControlFlags) o;
    return (getSet1KeyRevisionActivationTimeExpirationTimePresent()
            == that.getSet1KeyRevisionActivationTimeExpirationTimePresent())
        && (getSet1KeyCountKeyParametersPresent() == that.getSet1KeyCountKeyParametersPresent())
        && (getSet1ShouldBeCleared() == that.getSet1ShouldBeCleared())
        && (getSet2KeyRevisionActivationTimeExpirationTimePresent()
            == that.getSet2KeyRevisionActivationTimeExpirationTimePresent())
        && (getSet2KeyCountKeyParametersPresent() == that.getSet2KeyCountKeyParametersPresent())
        && (getSet2ShouldBeCleared() == that.getSet2ShouldBeCleared())
        && (getMoreMessagesToBeExpected() == that.getMoreMessagesToBeExpected())
        && (getRemoveAllKeys() == that.getRemoveAllKeys())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getSet1KeyRevisionActivationTimeExpirationTimePresent(),
        getSet1KeyCountKeyParametersPresent(),
        getSet1ShouldBeCleared(),
        getSet2KeyRevisionActivationTimeExpirationTimePresent(),
        getSet2KeyCountKeyParametersPresent(),
        getSet2ShouldBeCleared(),
        getMoreMessagesToBeExpected(),
        getRemoveAllKeys());
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
