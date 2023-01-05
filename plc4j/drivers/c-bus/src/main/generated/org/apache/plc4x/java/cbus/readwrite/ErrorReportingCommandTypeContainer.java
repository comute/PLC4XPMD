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

public enum ErrorReportingCommandTypeContainer {
  ErrorReportingCommandDeprecated((short) 0x05, (short) 5, ErrorReportingCommandType.DEPRECATED),
  ErrorReportingCommandErrorReport((short) 0x15, (short) 5, ErrorReportingCommandType.ERROR_REPORT),
  ErrorReportingCommandAcknowledge((short) 0x25, (short) 5, ErrorReportingCommandType.ACKNOWLEDGE),
  ErrorReportingCommandClearMostSevere(
      (short) 0x35, (short) 5, ErrorReportingCommandType.CLEAR_MOST_SEVERE);
  private static final Map<Short, ErrorReportingCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (ErrorReportingCommandTypeContainer value : ErrorReportingCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;
  private short numBytes;
  private ErrorReportingCommandType commandType;

  ErrorReportingCommandTypeContainer(
      short value, short numBytes, ErrorReportingCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static ErrorReportingCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (ErrorReportingCommandTypeContainer _val : ErrorReportingCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ErrorReportingCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<ErrorReportingCommandTypeContainer> _values = new ArrayList();
    for (ErrorReportingCommandTypeContainer _val : ErrorReportingCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public ErrorReportingCommandType getCommandType() {
    return commandType;
  }

  public static ErrorReportingCommandTypeContainer firstEnumForFieldCommandType(
      ErrorReportingCommandType fieldValue) {
    for (ErrorReportingCommandTypeContainer _val : ErrorReportingCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ErrorReportingCommandTypeContainer> enumsForFieldCommandType(
      ErrorReportingCommandType fieldValue) {
    List<ErrorReportingCommandTypeContainer> _values = new ArrayList();
    for (ErrorReportingCommandTypeContainer _val : ErrorReportingCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static ErrorReportingCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
