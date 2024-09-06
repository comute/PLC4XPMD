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

public abstract class TlvOrgSpecificProfibusUnit implements Message {

  // Abstract accessors for discriminator values.
  public abstract TlvProfibusSubType getSubType();

  public TlvOrgSpecificProfibusUnit() {
    super();
  }

  protected abstract void serializeTlvOrgSpecificProfibusUnitChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvOrgSpecificProfibusUnit");

    // Discriminator Field (subType) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "subType",
        "TlvProfibusSubType",
        getSubType(),
        writeEnum(
            TlvProfibusSubType::getValue,
            TlvProfibusSubType::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Switch field (Serialize the sub-type)
    serializeTlvOrgSpecificProfibusUnitChild(writeBuffer);

    writeBuffer.popContext("TlvOrgSpecificProfibusUnit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    TlvOrgSpecificProfibusUnit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (subType)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static TlvOrgSpecificProfibusUnit staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TlvOrgSpecificProfibusUnit");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TlvProfibusSubType subType =
        readDiscriminatorEnumField(
            "subType",
            "TlvProfibusSubType",
            readEnum(TlvProfibusSubType::enumForValue, readUnsignedShort(readBuffer, 8)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    TlvOrgSpecificProfibusUnitBuilder builder = null;
    if (EvaluationHelper.equals(subType, TlvProfibusSubType.MEASURED_DELAY)) {
      builder =
          TlvProfibusSubTypeMeasuredDelay.staticParseTlvOrgSpecificProfibusUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(subType, TlvProfibusSubType.PORT_STATUS)) {
      builder =
          TlvProfibusSubTypePortStatus.staticParseTlvOrgSpecificProfibusUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(subType, TlvProfibusSubType.MRP_PORT_STATUS)) {
      builder =
          TlvProfibusSubTypeMrpPortStatus.staticParseTlvOrgSpecificProfibusUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(subType, TlvProfibusSubType.CHASSIS_MAC)) {
      builder =
          TlvProfibusSubTypeChassisMac.staticParseTlvOrgSpecificProfibusUnitBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "subType=" + subType + "]");
    }

    readBuffer.closeContext("TlvOrgSpecificProfibusUnit");
    // Create the instance
    TlvOrgSpecificProfibusUnit _tlvOrgSpecificProfibusUnit = builder.build();
    return _tlvOrgSpecificProfibusUnit;
  }

  public interface TlvOrgSpecificProfibusUnitBuilder {
    TlvOrgSpecificProfibusUnit build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvOrgSpecificProfibusUnit)) {
      return false;
    }
    TlvOrgSpecificProfibusUnit that = (TlvOrgSpecificProfibusUnit) o;
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
