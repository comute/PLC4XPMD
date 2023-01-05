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

public class BACnetReadAccessResultListOfResults implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final List<BACnetReadAccessProperty> listOfReadAccessProperty;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetReadAccessResultListOfResults(
      BACnetOpeningTag openingTag,
      List<BACnetReadAccessProperty> listOfReadAccessProperty,
      BACnetClosingTag closingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super();
    this.openingTag = openingTag;
    this.listOfReadAccessProperty = listOfReadAccessProperty;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public List<BACnetReadAccessProperty> getListOfReadAccessProperty() {
    return listOfReadAccessProperty;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetReadAccessResultListOfResults");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Array Field (listOfReadAccessProperty)
    writeComplexTypeArrayField("listOfReadAccessProperty", listOfReadAccessProperty, writeBuffer);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetReadAccessResultListOfResults");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetReadAccessResultListOfResults _value = this;

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Array field
    if (listOfReadAccessProperty != null) {
      for (Message element : listOfReadAccessProperty) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetReadAccessResultListOfResults staticParse(
      ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    Short tagNumber;
    if (args[0] instanceof Short) {
      tagNumber = (Short) args[0];
    } else if (args[0] instanceof String) {
      tagNumber = Short.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Short or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    BACnetObjectType objectTypeArgument;
    if (args[1] instanceof BACnetObjectType) {
      objectTypeArgument = (BACnetObjectType) args[1];
    } else if (args[1] instanceof String) {
      objectTypeArgument = BACnetObjectType.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type BACnetObjectType or a string which is parseable but"
              + " was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, tagNumber, objectTypeArgument);
  }

  public static BACnetReadAccessResultListOfResults staticParse(
      ReadBuffer readBuffer, Short tagNumber, BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetReadAccessResultListOfResults");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    List<BACnetReadAccessProperty> listOfReadAccessProperty =
        readTerminatedArrayField(
            "listOfReadAccessProperty",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetReadAccessProperty.staticParse(
                        readBuffer, (BACnetObjectType) (objectTypeArgument)),
                readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetReadAccessResultListOfResults");
    // Create the instance
    BACnetReadAccessResultListOfResults _bACnetReadAccessResultListOfResults;
    _bACnetReadAccessResultListOfResults =
        new BACnetReadAccessResultListOfResults(
            openingTag, listOfReadAccessProperty, closingTag, tagNumber, objectTypeArgument);
    return _bACnetReadAccessResultListOfResults;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetReadAccessResultListOfResults)) {
      return false;
    }
    BACnetReadAccessResultListOfResults that = (BACnetReadAccessResultListOfResults) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getListOfReadAccessProperty() == that.getListOfReadAccessProperty())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getListOfReadAccessProperty(), getClosingTag());
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
