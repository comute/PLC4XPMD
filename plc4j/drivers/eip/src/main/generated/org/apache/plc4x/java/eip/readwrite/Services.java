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
package org.apache.plc4x.java.eip.readwrite;

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

public class Services implements Message {

  // Properties.
  protected final List<Integer> offsets;
  protected final List<CipService> services;

  public Services(List<Integer> offsets, List<CipService> services) {
    super();
    this.offsets = offsets;
    this.services = services;
  }

  public List<Integer> getOffsets() {
    return offsets;
  }

  public List<CipService> getServices() {
    return services;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Services");

    // Implicit Field (serviceNb) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int serviceNb = (int) (COUNT(getOffsets()));
    writeImplicitField("serviceNb", serviceNb, writeUnsignedInt(writeBuffer, 16));

    // Array Field (offsets)
    writeSimpleTypeArrayField("offsets", offsets, writeUnsignedInt(writeBuffer, 16));

    // Array Field (services)
    writeComplexTypeArrayField("services", services, writeBuffer);

    writeBuffer.popContext("Services");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Services _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (serviceNb)
    lengthInBits += 16;

    // Array field
    if (offsets != null) {
      lengthInBits += 16 * offsets.size();
    }

    // Array field
    if (services != null) {
      int i = 0;
      for (CipService element : services) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= services.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static Services staticParse(ReadBuffer readBuffer, Integer servicesLen)
      throws ParseException {
    readBuffer.pullContext("Services");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int serviceNb = readImplicitField("serviceNb", readUnsignedInt(readBuffer, 16));

    List<Integer> offsets =
        readCountArrayField("offsets", readUnsignedInt(readBuffer, 16), serviceNb);

    List<CipService> services =
        readCountArrayField(
            "services",
            readComplex(
                () ->
                    CipService.staticParse(
                        readBuffer, (boolean) (false), (int) ((servicesLen) / (serviceNb))),
                readBuffer),
            serviceNb);

    readBuffer.closeContext("Services");
    // Create the instance
    Services _services;
    _services = new Services(offsets, services);
    return _services;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Services)) {
      return false;
    }
    Services that = (Services) o;
    return (getOffsets() == that.getOffsets()) && (getServices() == that.getServices()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOffsets(), getServices());
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
