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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ErrorReportingSystemCategoryClass {
  RESERVED_0((byte) 0x0),
  RESERVED_1((byte) 0x1),
  RESERVED_2((byte) 0x2),
  RESERVED_3((byte) 0x3),
  RESERVED_4((byte) 0x4),
  INPUT_UNITS((byte) 0x5),
  RESERVED_6((byte) 0x6),
  RESERVED_7((byte) 0x7),
  RESERVED_8((byte) 0x8),
  SUPPORT_UNITS((byte) 0x9),
  RESERVED_10((byte) 0xA),
  BUILDING_MANAGEMENT_SYSTEMS((byte) 0xB),
  RESERVED_12((byte) 0xC),
  OUTPUT_UNITS((byte) 0xD),
  RESERVED_14((byte) 0xE),
  CLIMATE_CONTROLLERS((byte) 0xF);
  private static final Map<Byte, ErrorReportingSystemCategoryClass> map;

  static {
    map = new HashMap<>();
    for (ErrorReportingSystemCategoryClass value : ErrorReportingSystemCategoryClass.values()) {
      map.put((byte) value.getValue(), value);
    }
  }

  private byte value;

  ErrorReportingSystemCategoryClass(byte value) {
    this.value = value;
  }

  public byte getValue() {
    return value;
  }

  public static ErrorReportingSystemCategoryClass enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}
