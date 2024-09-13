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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusPDUWriteFileRecordResponseItem implements Message {

  // Properties.
  protected final short referenceType;
  protected final int fileNumber;
  protected final int recordNumber;
  protected final byte[] recordData;

  public ModbusPDUWriteFileRecordResponseItem(
      short referenceType, int fileNumber, int recordNumber, byte[] recordData) {
    super();
    this.referenceType = referenceType;
    this.fileNumber = fileNumber;
    this.recordNumber = recordNumber;
    this.recordData = recordData;
  }

  public short getReferenceType() {
    return referenceType;
  }

  public int getFileNumber() {
    return fileNumber;
  }

  public int getRecordNumber() {
    return recordNumber;
  }

  public byte[] getRecordData() {
    return recordData;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ModbusPDUWriteFileRecordResponseItem");

    // Simple Field (referenceType)
    writeSimpleField("referenceType", referenceType, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (fileNumber)
    writeSimpleField("fileNumber", fileNumber, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (recordNumber)
    writeSimpleField("recordNumber", recordNumber, writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (recordLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int recordLength = (int) ((COUNT(getRecordData())) / (2));
    writeImplicitField("recordLength", recordLength, writeUnsignedInt(writeBuffer, 16));

    // Array Field (recordData)
    writeByteArrayField("recordData", recordData, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("ModbusPDUWriteFileRecordResponseItem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ModbusPDUWriteFileRecordResponseItem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (referenceType)
    lengthInBits += 8;

    // Simple field (fileNumber)
    lengthInBits += 16;

    // Simple field (recordNumber)
    lengthInBits += 16;

    // Implicit Field (recordLength)
    lengthInBits += 16;

    // Array field
    if (recordData != null) {
      lengthInBits += 8 * recordData.length;
    }

    return lengthInBits;
  }

  public static ModbusPDUWriteFileRecordResponseItem staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("ModbusPDUWriteFileRecordResponseItem");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short referenceType = readSimpleField("referenceType", readUnsignedShort(readBuffer, 8));

    int fileNumber = readSimpleField("fileNumber", readUnsignedInt(readBuffer, 16));

    int recordNumber = readSimpleField("recordNumber", readUnsignedInt(readBuffer, 16));

    int recordLength = readImplicitField("recordLength", readUnsignedInt(readBuffer, 16));

    byte[] recordData = readBuffer.readByteArray("recordData", Math.toIntExact(recordLength));

    readBuffer.closeContext("ModbusPDUWriteFileRecordResponseItem");
    // Create the instance
    ModbusPDUWriteFileRecordResponseItem _modbusPDUWriteFileRecordResponseItem;
    _modbusPDUWriteFileRecordResponseItem =
        new ModbusPDUWriteFileRecordResponseItem(
            referenceType, fileNumber, recordNumber, recordData);
    return _modbusPDUWriteFileRecordResponseItem;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUWriteFileRecordResponseItem)) {
      return false;
    }
    ModbusPDUWriteFileRecordResponseItem that = (ModbusPDUWriteFileRecordResponseItem) o;
    return (getReferenceType() == that.getReferenceType())
        && (getFileNumber() == that.getFileNumber())
        && (getRecordNumber() == that.getRecordNumber())
        && (getRecordData() == that.getRecordData())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getReferenceType(), getFileNumber(), getRecordNumber(), getRecordData());
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
