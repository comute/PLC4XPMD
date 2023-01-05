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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum LevelInformationNibblePair {
  Value_F((short) 0x55, (byte) 0xF),
  Value_E((short) 0x56, (byte) 0xE),
  Value_D((short) 0x59, (byte) 0xD),
  Value_C((short) 0x5A, (byte) 0xC),
  Value_B((short) 0x65, (byte) 0xB),
  Value_A((short) 0x66, (byte) 0xA),
  Value_9((short) 0x69, (byte) 0x9),
  Value_8((short) 0x6A, (byte) 0x8),
  Value_7((short) 0x95, (byte) 0x7),
  Value_6((short) 0x96, (byte) 0x6),
  Value_5((short) 0x99, (byte) 0x5),
  Value_4((short) 0x9A, (byte) 0x4),
  Value_3((short) 0xA5, (byte) 0x3),
  Value_2((short) 0xA6, (byte) 0x2),
  Value_1((short) 0xA9, (byte) 0x1),
  Value_0((short) 0xAA, (byte) 0x0);
  private static final Map<Short, LevelInformationNibblePair> map;

  static {
    map = new HashMap<>();
    for (LevelInformationNibblePair value : LevelInformationNibblePair.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;
  private byte nibbleValue;

  LevelInformationNibblePair(short value, byte nibbleValue) {
    this.value = value;
    this.nibbleValue = nibbleValue;
  }

  public short getValue() {
    return value;
  }

  public byte getNibbleValue() {
    return nibbleValue;
  }

  public static LevelInformationNibblePair firstEnumForFieldNibbleValue(byte fieldValue) {
    for (LevelInformationNibblePair _val : LevelInformationNibblePair.values()) {
      if (_val.getNibbleValue() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<LevelInformationNibblePair> enumsForFieldNibbleValue(byte fieldValue) {
    List<LevelInformationNibblePair> _values = new ArrayList();
    for (LevelInformationNibblePair _val : LevelInformationNibblePair.values()) {
      if (_val.getNibbleValue() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static LevelInformationNibblePair enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
