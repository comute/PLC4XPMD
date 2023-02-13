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

public class MediaTransportControlDataSetSelection extends MediaTransportControlData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte selectionHi;
  protected final byte selectionLo;

  public MediaTransportControlDataSetSelection(
      MediaTransportControlCommandTypeContainer commandTypeContainer,
      byte mediaLinkGroup,
      byte selectionHi,
      byte selectionLo) {
    super(commandTypeContainer, mediaLinkGroup);
    this.selectionHi = selectionHi;
    this.selectionLo = selectionLo;
  }

  public byte getSelectionHi() {
    return selectionHi;
  }

  public byte getSelectionLo() {
    return selectionLo;
  }

  @Override
  protected void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MediaTransportControlDataSetSelection");

    // Simple Field (selectionHi)
    writeSimpleField("selectionHi", selectionHi, writeByte(writeBuffer, 8));

    // Simple Field (selectionLo)
    writeSimpleField("selectionLo", selectionLo, writeByte(writeBuffer, 8));

    writeBuffer.popContext("MediaTransportControlDataSetSelection");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MediaTransportControlDataSetSelection _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (selectionHi)
    lengthInBits += 8;

    // Simple field (selectionLo)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static MediaTransportControlDataBuilder staticParseMediaTransportControlDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MediaTransportControlDataSetSelection");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte selectionHi = readSimpleField("selectionHi", readByte(readBuffer, 8));

    byte selectionLo = readSimpleField("selectionLo", readByte(readBuffer, 8));

    readBuffer.closeContext("MediaTransportControlDataSetSelection");
    // Create the instance
    return new MediaTransportControlDataSetSelectionBuilderImpl(selectionHi, selectionLo);
  }

  public static class MediaTransportControlDataSetSelectionBuilderImpl
      implements MediaTransportControlData.MediaTransportControlDataBuilder {
    private final byte selectionHi;
    private final byte selectionLo;

    public MediaTransportControlDataSetSelectionBuilderImpl(byte selectionHi, byte selectionLo) {
      this.selectionHi = selectionHi;
      this.selectionLo = selectionLo;
    }

    public MediaTransportControlDataSetSelection build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
      MediaTransportControlDataSetSelection mediaTransportControlDataSetSelection =
          new MediaTransportControlDataSetSelection(
              commandTypeContainer, mediaLinkGroup, selectionHi, selectionLo);
      return mediaTransportControlDataSetSelection;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlDataSetSelection)) {
      return false;
    }
    MediaTransportControlDataSetSelection that = (MediaTransportControlDataSetSelection) o;
    return (getSelectionHi() == that.getSelectionHi())
        && (getSelectionLo() == that.getSelectionLo())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSelectionHi(), getSelectionLo());
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
