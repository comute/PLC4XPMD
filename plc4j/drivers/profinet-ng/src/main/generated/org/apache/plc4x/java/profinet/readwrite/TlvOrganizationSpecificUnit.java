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

public abstract class TlvOrganizationSpecificUnit implements Message {

  // Abstract accessors for discriminator values.
  public abstract Integer getUniqueCode();

  public TlvOrganizationSpecificUnit() {
    super();
  }

  protected abstract void serializeTlvOrganizationSpecificUnitChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvOrganizationSpecificUnit");

    // Discriminator Field (uniqueCode) (Used as input to a switch field)
    writeDiscriminatorField("uniqueCode", getUniqueCode(), writeUnsignedInt(writeBuffer, 24));

    // Switch field (Serialize the sub-type)
    serializeTlvOrganizationSpecificUnitChild(writeBuffer);

    writeBuffer.popContext("TlvOrganizationSpecificUnit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    TlvOrganizationSpecificUnit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (uniqueCode)
    lengthInBits += 24;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static TlvOrganizationSpecificUnit staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TlvOrganizationSpecificUnit");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int uniqueCode = readDiscriminatorField("uniqueCode", readUnsignedInt(readBuffer, 24));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    TlvOrganizationSpecificUnitBuilder builder = null;
    if (EvaluationHelper.equals(uniqueCode, (int) 0x000ECF)) {
      builder = TlvOrgSpecificProfibus.staticParseTlvOrganizationSpecificUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(uniqueCode, (int) 0x00120F)) {
      builder = TlvOrgSpecificIeee8023.staticParseTlvOrganizationSpecificUnitBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "uniqueCode="
              + uniqueCode
              + "]");
    }

    readBuffer.closeContext("TlvOrganizationSpecificUnit");
    // Create the instance
    TlvOrganizationSpecificUnit _tlvOrganizationSpecificUnit = builder.build();
    return _tlvOrganizationSpecificUnit;
  }

  public interface TlvOrganizationSpecificUnitBuilder {
    TlvOrganizationSpecificUnit build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvOrganizationSpecificUnit)) {
      return false;
    }
    TlvOrganizationSpecificUnit that = (TlvOrganizationSpecificUnit) o;
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
