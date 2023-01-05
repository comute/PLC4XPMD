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

import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.EvaluationHelper;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataItem {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataItem.class);

  public static PlcValue staticParse(
      ReadBuffer readBuffer, ModbusDataType dataType, Integer numberOfValues)
      throws ParseException {
    if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BOOL

      // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
      {
        int reserved = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 15);
        if (reserved != (int) 0x0000) {
          LOGGER.info(
              "Expected constant value "
                  + 0x0000
                  + " but got "
                  + reserved
                  + " for reserved field.");
        }
      }

      // Simple Field (value)
      Boolean value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("");

      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcBOOL(
                  (Boolean) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("")));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BYTE

      // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
      {
        short reserved = /*TODO: migrate me*/ /*TODO: migrate me*/
            readBuffer.readUnsignedShort("", 8);
        if (reserved != (short) 0x00) {
          LOGGER.info(
              "Expected constant value " + 0x00 + " but got " + reserved + " for reserved field.");
        }
      }

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcBYTE(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)) { // List
      // Array field (value)
      // Count array
      if ((numberOfValues) * (8) > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + ((numberOfValues) * (8))
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) (numberOfValues) * (8);
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcBOOL(
                  (Boolean) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("")));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WORD)) { // WORD

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcWORD(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DWORD)) { // DWORD

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 32);

      return new PlcDWORD(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LWORD)) { // LWORD

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return new PlcLWORD(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // SINT

      // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
      {
        short reserved = /*TODO: migrate me*/ /*TODO: migrate me*/
            readBuffer.readUnsignedShort("", 8);
        if (reserved != (short) 0x00) {
          LOGGER.info(
              "Expected constant value " + 0x00 + " but got " + reserved + " for reserved field.");
        }
      }

      // Simple Field (value)
      Byte value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readSignedByte("", 8);

      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcSINT(
                  (Byte) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readSignedByte("", 8)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // INT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readShort("", 16);

      return new PlcINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcINT(
                  (Short) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readShort("", 16)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // DINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readInt("", 32);

      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcDINT(
                  (Integer) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readInt("", 32)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 64);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcLINT(
                  (Long) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 64)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // USINT

      // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
      {
        short reserved = /*TODO: migrate me*/ /*TODO: migrate me*/
            readBuffer.readUnsignedShort("", 8);
        if (reserved != (short) 0x00) {
          LOGGER.info(
              "Expected constant value " + 0x00 + " but got " + reserved + " for reserved field.");
        }
      }

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcUINT(
                  (Short) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readUnsignedShort("", 8)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcUDINT(
                  (Integer) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readUnsignedInt("", 16)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UDINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 32);

      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcULINT(
                  (Long) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readUnsignedLong("", 32)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcLINT(
                  (BigInteger) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readUnsignedBigInteger("", 64)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // REAL

      // Simple Field (value)
      Float value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readFloat("", 32);

      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcREAL(
                  (Float) /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readFloat("", 32)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LREAL

      // Simple Field (value)
      Double value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readDouble("", 64);

      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcLREAL(
                  (Double) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readDouble("", 64)));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // CHAR

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", 8, "UTF-8");

      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcSTRING(
                  (String) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readString("", 8, "UTF-8")));
        }
      }

      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // WCHAR

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", 16, "UTF-16");

      return new PlcWCHAR(value);
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)) { // List
      // Array field (value)
      // Count array
      if (numberOfValues > Integer.MAX_VALUE) {
        throw new ParseException(
            "Array count of "
                + (numberOfValues)
                + " exceeds the maximum allowed count of "
                + Integer.MAX_VALUE);
      }
      List<PlcValue> value;
      {
        int itemCount = (int) numberOfValues;
        value = new LinkedList<>();
        for (int curItem = 0; curItem < itemCount; curItem++) {
          value.add(
              new PlcSTRING(
                  (String) /*TODO: migrate me*/ /*TODO: migrate me*/
                      readBuffer.readString("", 16, "UTF-16")));
        }
      }

      return new PlcList(value);
    }
    return null;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, ModbusDataType dataType, Integer numberOfValues)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, dataType, numberOfValues, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      ModbusDataType dataType,
      Integer numberOfValues,
      ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BOOL
      // Reserved Field
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 15, ((Number) (int) 0x0000).intValue());
      // Simple Field (value)
      boolean value = (boolean) _value.getBoolean();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeBit("", (boolean) (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Boolean value = (Boolean) val.getBoolean();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeBit("", (boolean) (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BYTE
      // Reserved Field
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort(
          "", 8, ((Number) (short) 0x00).shortValue());
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Boolean value = (Boolean) val.getBoolean();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeBit("", (boolean) (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WORD)) { // WORD
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DWORD)) { // DWORD
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LWORD)) { // LWORD
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // SINT
      // Reserved Field
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort(
          "", 8, ((Number) (short) 0x00).shortValue());
      // Simple Field (value)
      byte value = (byte) _value.getByte();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeSignedByte("", 8, ((Number) (value)).byteValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Byte value = (Byte) val.getByte();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeSignedByte("", 8, ((Number) (value)).byteValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // INT
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeShort("", 16, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Short value = (Short) val.getShort();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeShort("", 16, ((Number) (value)).shortValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // DINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeInt("", 32, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Integer value = (Integer) val.getInteger();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeInt("", 32, ((Number) (value)).intValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 64, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Long value = (Long) val.getLong();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeLong("", 64, ((Number) (value)).longValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // USINT
      // Reserved Field
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort(
          "", 8, ((Number) (short) 0x00).shortValue());
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Short value = (Short) val.getShort();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Integer value = (Integer) val.getInteger();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UDINT
      // Simple Field (value)
      long value = (long) _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Long value = (Long) val.getLong();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // ULINT
      // Simple Field (value)
      BigInteger value = (BigInteger) _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, (BigInteger) (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        BigInteger value = (BigInteger) val.getBigInteger();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, (BigInteger) (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // REAL
      // Simple Field (value)
      float value = (float) _value.getFloat();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeFloat("", 32, (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Float value = (Float) val.getFloat();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeFloat("", 32, (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LREAL
      // Simple Field (value)
      double value = (double) _value.getDouble();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeDouble("", 64, (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        Double value = (Double) val.getDouble();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeDouble("", 64, (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // CHAR
      // Simple Field (value)
      String value = (String) _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString("", 8, "UTF-8", (String) (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        String value = (String) val.getString();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeString("", 8, "UTF-8", (String) (value));
      }

    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // WCHAR
      // Simple Field (value)
      String value = (String) _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString("", 16, "UTF-16", (String) (value));
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)) { // List
      PlcList values = (PlcList) _value;

      for (PlcValue val : ((List<PlcValue>) values.getList())) {
        String value = (String) val.getString();
        /*TODO: migrate me*/
        /*TODO: migrate me*/ writeBuffer.writeString("", 16, "UTF-16", (String) (value));
      }
    }
  }

  public static int getLengthInBytes(
      PlcValue _value, ModbusDataType dataType, Integer numberOfValues) {
    return (int) Math.ceil((float) getLengthInBits(_value, dataType, numberOfValues) / 8.0);
  }

  public static int getLengthInBits(
      PlcValue _value, ModbusDataType dataType, Integer numberOfValues) {
    int sizeInBits = 0;
    if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BOOL
      // Reserved Field
      sizeInBits += 15;
      // Simple Field (value)
      sizeInBits += 1;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BOOL)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 1;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)
        && EvaluationHelper.equals(numberOfValues, 1)) { // BYTE
      // Reserved Field
      sizeInBits += 8;
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.BYTE)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 1;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WORD)) { // WORD
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DWORD)) { // DWORD
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LWORD)) { // LWORD
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // SINT
      // Reserved Field
      sizeInBits += 8;
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.SINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // INT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.INT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // DINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.DINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // USINT
      // Reserved Field
      sizeInBits += 8;
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.USINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UINT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // UDINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.UDINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)
        && EvaluationHelper.equals(numberOfValues, 1)) { // ULINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.ULINT)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // REAL
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.REAL)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 32;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)
        && EvaluationHelper.equals(numberOfValues, 1)) { // LREAL
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.LREAL)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 64;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // CHAR
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.CHAR)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 8;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)
        && EvaluationHelper.equals(numberOfValues, 1)) { // WCHAR
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataType, ModbusDataType.WCHAR)) { // List
      PlcList values = (PlcList) _value;
      sizeInBits += values.getList().size() * 16;
    }
    return sizeInBits;
  }
}
