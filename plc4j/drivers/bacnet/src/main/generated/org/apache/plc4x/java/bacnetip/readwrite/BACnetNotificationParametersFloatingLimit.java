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

public class BACnetNotificationParametersFloatingLimit extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetContextTagReal referenceValue;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetContextTagReal setPointValue;
  protected final BACnetContextTagReal errorLimit;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersFloatingLimit(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetContextTagReal referenceValue,
      BACnetStatusFlagsTagged statusFlags,
      BACnetContextTagReal setPointValue,
      BACnetContextTagReal errorLimit,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.referenceValue = referenceValue;
    this.statusFlags = statusFlags;
    this.setPointValue = setPointValue;
    this.errorLimit = errorLimit;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetContextTagReal getReferenceValue() {
    return referenceValue;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetContextTagReal getSetPointValue() {
    return setPointValue;
  }

  public BACnetContextTagReal getErrorLimit() {
    return errorLimit;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetNotificationParametersFloatingLimit");

    // Simple Field (innerOpeningTag)
    writeSimpleField(
        "innerOpeningTag", innerOpeningTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (referenceValue)
    writeSimpleField("referenceValue", referenceValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (setPointValue)
    writeSimpleField("setPointValue", setPointValue, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (errorLimit)
    writeSimpleField("errorLimit", errorLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField(
        "innerClosingTag", innerClosingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersFloatingLimit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersFloatingLimit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (referenceValue)
    lengthInBits += referenceValue.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (setPointValue)
    lengthInBits += setPointValue.getLengthInBits();

    // Simple field (errorLimit)
    lengthInBits += errorLimit.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersBuilder staticParseBACnetNotificationParametersBuilder(
      ReadBuffer readBuffer,
      Short peekedTagNumber,
      Short tagNumber,
      BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersFloatingLimit");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetContextTagReal referenceValue =
        readSimpleField(
            "referenceValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (0), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagReal setPointValue =
        readSimpleField(
            "setPointValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (2), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetContextTagReal errorLimit =
        readSimpleField(
            "errorLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (3), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersFloatingLimit");
    // Create the instance
    return new BACnetNotificationParametersFloatingLimitBuilderImpl(
        innerOpeningTag,
        referenceValue,
        statusFlags,
        setPointValue,
        errorLimit,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersFloatingLimitBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetContextTagReal referenceValue;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetContextTagReal setPointValue;
    private final BACnetContextTagReal errorLimit;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersFloatingLimitBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetContextTagReal referenceValue,
        BACnetStatusFlagsTagged statusFlags,
        BACnetContextTagReal setPointValue,
        BACnetContextTagReal errorLimit,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.referenceValue = referenceValue;
      this.statusFlags = statusFlags;
      this.setPointValue = setPointValue;
      this.errorLimit = errorLimit;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersFloatingLimit build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersFloatingLimit bACnetNotificationParametersFloatingLimit =
          new BACnetNotificationParametersFloatingLimit(
              openingTag,
              peekedTagHeader,
              closingTag,
              innerOpeningTag,
              referenceValue,
              statusFlags,
              setPointValue,
              errorLimit,
              innerClosingTag,
              tagNumber,
              objectTypeArgument);
      return bACnetNotificationParametersFloatingLimit;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersFloatingLimit)) {
      return false;
    }
    BACnetNotificationParametersFloatingLimit that = (BACnetNotificationParametersFloatingLimit) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getReferenceValue() == that.getReferenceValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getSetPointValue() == that.getSetPointValue())
        && (getErrorLimit() == that.getErrorLimit())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getReferenceValue(),
        getStatusFlags(),
        getSetPointValue(),
        getErrorLimit(),
        getInnerClosingTag());
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
