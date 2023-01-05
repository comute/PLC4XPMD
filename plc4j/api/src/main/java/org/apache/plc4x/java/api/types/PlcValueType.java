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
package org.apache.plc4x.java.api.types;

import java.math.BigInteger;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum PlcValueType {
    NULL((short) 0x00, null),
    BOOL((short) 0x01, Boolean.class),
    BYTE((short) 0x02, Byte.class),
    WORD((short) 0x03, Short.class),
    DWORD((short) 0x04, Integer.class),
    LWORD((short) 0x05, Long.class),
    USINT((short) 0x11, Short.class),
    UINT((short) 0x12, Integer.class),
    UDINT((short) 0x13, Long.class),
    ULINT((short) 0x14, BigInteger.class),
    SINT((short) 0x21, Byte.class),
    INT((short) 0x22, Short.class),
    DINT((short) 0x23, Integer.class),
    LINT((short) 0x24, Long.class),
    REAL((short) 0x31, Float.class),
    LREAL((short) 0x32, Double.class),
    CHAR((short) 0x41, Character.class),
    WCHAR((short) 0x42, Short.class),
    STRING((short) 0x43, String.class),
    WSTRING((short) 0x44, String.class),
    TIME((short) 0x51, Duration.class),
    LTIME((short) 0x52, Duration.class),
    DATE((short) 0x53, LocalDate.class),
    LDATE((short) 0x54, LocalDate.class),
    TIME_OF_DAY((short) 0x55, LocalTime.class),
    LTIME_OF_DAY((short) 0x56, LocalTime.class),
    DATE_AND_TIME((short) 0x57, LocalDateTime.class),
    LDATE_AND_TIME((short) 0x58, LocalDateTime.class),
    Struct((short) 0x61, HashMap.class),
    List((short) 0x62, ArrayList.class),
    RAW_BYTE_ARRAY((short) 0x71, Byte.class);
    private static final Map<Short, PlcValueType> map;

    static {
        map = new HashMap<>();
        for (PlcValueType value : PlcValueType.values()) {
            map.put(value.getValue(), value);
        }
    }

    private short value;
    private Class<?> defaultJavaType;

    PlcValueType(short value, Class<?> defaultJavaType) {
        this.value = value;
        this.defaultJavaType = defaultJavaType;
    }

    public short getValue() {
        return value;
    }

    public Class<?> getDefaultJavaType() {
        return defaultJavaType;
    }

    public static PlcValueType enumForValue(short value) {
        return map.get(value);
    }

    public static Boolean isDefined(short value) {
        return map.containsKey(value);
    }
}
