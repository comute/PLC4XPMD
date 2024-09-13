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

public class BACnetConstructedDataDefaultSubordinateRelationship extends BACnetConstructedData
    implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.DEFAULT_SUBORDINATE_RELATIONSHIP;
  }

  // Properties.
  protected final BACnetRelationshipTagged defaultSubordinateRelationship;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataDefaultSubordinateRelationship(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetRelationshipTagged defaultSubordinateRelationship,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.defaultSubordinateRelationship = defaultSubordinateRelationship;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetRelationshipTagged getDefaultSubordinateRelationship() {
    return defaultSubordinateRelationship;
  }

  public BACnetRelationshipTagged getActualValue() {
    return (BACnetRelationshipTagged) (getDefaultSubordinateRelationship());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataDefaultSubordinateRelationship");

    // Simple Field (defaultSubordinateRelationship)
    writeSimpleField(
        "defaultSubordinateRelationship",
        defaultSubordinateRelationship,
        writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetRelationshipTagged actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataDefaultSubordinateRelationship");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataDefaultSubordinateRelationship _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (defaultSubordinateRelationship)
    lengthInBits += defaultSubordinateRelationship.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataDefaultSubordinateRelationship");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetRelationshipTagged defaultSubordinateRelationship =
        readSimpleField(
            "defaultSubordinateRelationship",
            readComplex(
                () ->
                    BACnetRelationshipTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));
    BACnetRelationshipTagged actualValue =
        readVirtualField(
            "actualValue", BACnetRelationshipTagged.class, defaultSubordinateRelationship);

    readBuffer.closeContext("BACnetConstructedDataDefaultSubordinateRelationship");
    // Create the instance
    return new BACnetConstructedDataDefaultSubordinateRelationshipBuilderImpl(
        defaultSubordinateRelationship, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataDefaultSubordinateRelationshipBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetRelationshipTagged defaultSubordinateRelationship;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataDefaultSubordinateRelationshipBuilderImpl(
        BACnetRelationshipTagged defaultSubordinateRelationship,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.defaultSubordinateRelationship = defaultSubordinateRelationship;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataDefaultSubordinateRelationship build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataDefaultSubordinateRelationship
          bACnetConstructedDataDefaultSubordinateRelationship =
              new BACnetConstructedDataDefaultSubordinateRelationship(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  defaultSubordinateRelationship,
                  tagNumber,
                  arrayIndexArgument);
      return bACnetConstructedDataDefaultSubordinateRelationship;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataDefaultSubordinateRelationship)) {
      return false;
    }
    BACnetConstructedDataDefaultSubordinateRelationship that =
        (BACnetConstructedDataDefaultSubordinateRelationship) o;
    return (getDefaultSubordinateRelationship() == that.getDefaultSubordinateRelationship())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDefaultSubordinateRelationship());
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
