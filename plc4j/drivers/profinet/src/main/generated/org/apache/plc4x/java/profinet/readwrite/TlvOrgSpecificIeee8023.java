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
package org.apache.plc4x.java.profinet.readwrite;

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

public class TlvOrgSpecificIeee8023 extends TlvOrganizationSpecificUnit implements Message {

  // Accessors for discriminator values.
  public Long getUniqueCode() {
    return (long) 0x00120F;
  }

  // Properties.
  protected final TlvOrgSpecificIeee8023Unit specificUnit;

  public TlvOrgSpecificIeee8023(TlvOrgSpecificIeee8023Unit specificUnit) {
    super();
    this.specificUnit = specificUnit;
  }

  public TlvOrgSpecificIeee8023Unit getSpecificUnit() {
    return specificUnit;
  }

  @Override
  protected void serializeTlvOrganizationSpecificUnitChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TlvOrgSpecificIeee8023");

    // Simple Field (specificUnit)
    writeSimpleField("specificUnit", specificUnit, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("TlvOrgSpecificIeee8023");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvOrgSpecificIeee8023 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (specificUnit)
    lengthInBits += specificUnit.getLengthInBits();

    return lengthInBits;
  }

  public static TlvOrganizationSpecificUnitBuilder staticParseTlvOrganizationSpecificUnitBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("TlvOrgSpecificIeee8023");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TlvOrgSpecificIeee8023Unit specificUnit =
        readSimpleField(
            "specificUnit",
            new DataReaderComplexDefault<>(
                () -> TlvOrgSpecificIeee8023Unit.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("TlvOrgSpecificIeee8023");
    // Create the instance
    return new TlvOrgSpecificIeee8023BuilderImpl(specificUnit);
  }

  public static class TlvOrgSpecificIeee8023BuilderImpl
      implements TlvOrganizationSpecificUnit.TlvOrganizationSpecificUnitBuilder {
    private final TlvOrgSpecificIeee8023Unit specificUnit;

    public TlvOrgSpecificIeee8023BuilderImpl(TlvOrgSpecificIeee8023Unit specificUnit) {
      this.specificUnit = specificUnit;
    }

    public TlvOrgSpecificIeee8023 build() {
      TlvOrgSpecificIeee8023 tlvOrgSpecificIeee8023 = new TlvOrgSpecificIeee8023(specificUnit);
      return tlvOrgSpecificIeee8023;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvOrgSpecificIeee8023)) {
      return false;
    }
    TlvOrgSpecificIeee8023 that = (TlvOrgSpecificIeee8023) o;
    return (getSpecificUnit() == that.getSpecificUnit()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSpecificUnit());
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
