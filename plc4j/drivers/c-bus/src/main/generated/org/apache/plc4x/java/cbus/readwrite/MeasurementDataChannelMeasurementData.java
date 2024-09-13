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
package org.apache.plc4x.java.cbus.readwrite;

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

public class MeasurementDataChannelMeasurementData extends MeasurementData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final short deviceId;
  protected final short channel;
  protected final MeasurementUnits units;
  protected final byte multiplier;
  protected final short msb;
  protected final short lsb;

  public MeasurementDataChannelMeasurementData(
      MeasurementCommandTypeContainer commandTypeContainer,
      short deviceId,
      short channel,
      MeasurementUnits units,
      byte multiplier,
      short msb,
      short lsb) {
    super(commandTypeContainer);
    this.deviceId = deviceId;
    this.channel = channel;
    this.units = units;
    this.multiplier = multiplier;
    this.msb = msb;
    this.lsb = lsb;
  }

  public short getDeviceId() {
    return deviceId;
  }

  public short getChannel() {
    return channel;
  }

  public MeasurementUnits getUnits() {
    return units;
  }

  public byte getMultiplier() {
    return multiplier;
  }

  public short getMsb() {
    return msb;
  }

  public short getLsb() {
    return lsb;
  }

  public int getRawValue() {
    return (int) (((getMsb()) << (8)) | (getLsb()));
  }

  public double getValue() {
    return (double) (((getRawValue()) * (getMultiplier())) * (10));
  }

  @Override
  protected void serializeMeasurementDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MeasurementDataChannelMeasurementData");

    // Simple Field (deviceId)
    writeSimpleField("deviceId", deviceId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (channel)
    writeSimpleField("channel", channel, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (units)
    writeSimpleEnumField(
        "units",
        "MeasurementUnits",
        units,
        writeEnum(
            MeasurementUnits::getValue,
            MeasurementUnits::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (multiplier)
    writeSimpleField("multiplier", multiplier, writeSignedByte(writeBuffer, 8));

    // Simple Field (msb)
    writeSimpleField("msb", msb, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (lsb)
    writeSimpleField("lsb", lsb, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    int rawValue = getRawValue();
    writeBuffer.writeVirtual("rawValue", rawValue);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    double value = getValue();
    writeBuffer.writeVirtual("value", value);

    writeBuffer.popContext("MeasurementDataChannelMeasurementData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MeasurementDataChannelMeasurementData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (deviceId)
    lengthInBits += 8;

    // Simple field (channel)
    lengthInBits += 8;

    // Simple field (units)
    lengthInBits += 8;

    // Simple field (multiplier)
    lengthInBits += 8;

    // Simple field (msb)
    lengthInBits += 8;

    // Simple field (lsb)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static MeasurementDataBuilder staticParseMeasurementDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("MeasurementDataChannelMeasurementData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short deviceId = readSimpleField("deviceId", readUnsignedShort(readBuffer, 8));

    short channel = readSimpleField("channel", readUnsignedShort(readBuffer, 8));

    MeasurementUnits units =
        readEnumField(
            "units",
            "MeasurementUnits",
            readEnum(MeasurementUnits::enumForValue, readUnsignedShort(readBuffer, 8)));

    byte multiplier = readSimpleField("multiplier", readSignedByte(readBuffer, 8));

    short msb = readSimpleField("msb", readUnsignedShort(readBuffer, 8));

    short lsb = readSimpleField("lsb", readUnsignedShort(readBuffer, 8));
    int rawValue = readVirtualField("rawValue", int.class, ((msb) << (8)) | (lsb));
    double value = readVirtualField("value", double.class, ((rawValue) * (multiplier)) * (10));

    readBuffer.closeContext("MeasurementDataChannelMeasurementData");
    // Create the instance
    return new MeasurementDataChannelMeasurementDataBuilderImpl(
        deviceId, channel, units, multiplier, msb, lsb);
  }

  public static class MeasurementDataChannelMeasurementDataBuilderImpl
      implements MeasurementData.MeasurementDataBuilder {
    private final short deviceId;
    private final short channel;
    private final MeasurementUnits units;
    private final byte multiplier;
    private final short msb;
    private final short lsb;

    public MeasurementDataChannelMeasurementDataBuilderImpl(
        short deviceId,
        short channel,
        MeasurementUnits units,
        byte multiplier,
        short msb,
        short lsb) {
      this.deviceId = deviceId;
      this.channel = channel;
      this.units = units;
      this.multiplier = multiplier;
      this.msb = msb;
      this.lsb = lsb;
    }

    public MeasurementDataChannelMeasurementData build(
        MeasurementCommandTypeContainer commandTypeContainer) {
      MeasurementDataChannelMeasurementData measurementDataChannelMeasurementData =
          new MeasurementDataChannelMeasurementData(
              commandTypeContainer, deviceId, channel, units, multiplier, msb, lsb);
      return measurementDataChannelMeasurementData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MeasurementDataChannelMeasurementData)) {
      return false;
    }
    MeasurementDataChannelMeasurementData that = (MeasurementDataChannelMeasurementData) o;
    return (getDeviceId() == that.getDeviceId())
        && (getChannel() == that.getChannel())
        && (getUnits() == that.getUnits())
        && (getMultiplier() == that.getMultiplier())
        && (getMsb() == that.getMsb())
        && (getLsb() == that.getLsb())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getDeviceId(),
        getChannel(),
        getUnits(),
        getMultiplier(),
        getMsb(),
        getLsb());
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
