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

public enum CALCommandTypeContainer {
  CALCommandReset((short) 0x08, (short) 0, CALCommandType.RESET),
  CALCommandRecall((short) 0x1A, (short) 2, CALCommandType.RECALL),
  CALCommandIdentify((short) 0x21, (short) 1, CALCommandType.IDENTIFY),
  CALCommandGetStatus((short) 0x2A, (short) 2, CALCommandType.GET_STATUS),
  CALCommandAcknowledge((short) 0x32, (short) 2, CALCommandType.ACKNOWLEDGE),
  CALCommandReply_0Bytes((short) 0x80, (short) 0, CALCommandType.REPLY),
  CALCommandReply_1Bytes((short) 0x81, (short) 1, CALCommandType.REPLY),
  CALCommandReply_2Bytes((short) 0x82, (short) 2, CALCommandType.REPLY),
  CALCommandReply_3Bytes((short) 0x83, (short) 3, CALCommandType.REPLY),
  CALCommandReply_4Bytes((short) 0x84, (short) 4, CALCommandType.REPLY),
  CALCommandReply_5Bytes((short) 0x85, (short) 5, CALCommandType.REPLY),
  CALCommandReply_6Bytes((short) 0x86, (short) 6, CALCommandType.REPLY),
  CALCommandReply_7Bytes((short) 0x87, (short) 7, CALCommandType.REPLY),
  CALCommandReply_8Bytes((short) 0x88, (short) 8, CALCommandType.REPLY),
  CALCommandReply_9Bytes((short) 0x89, (short) 9, CALCommandType.REPLY),
  CALCommandReply_10Bytes((short) 0x8A, (short) 10, CALCommandType.REPLY),
  CALCommandReply_11Bytes((short) 0x8B, (short) 11, CALCommandType.REPLY),
  CALCommandReply_12Bytes((short) 0x8C, (short) 12, CALCommandType.REPLY),
  CALCommandReply_13Bytes((short) 0x8D, (short) 13, CALCommandType.REPLY),
  CALCommandReply_14Bytes((short) 0x8E, (short) 14, CALCommandType.REPLY),
  CALCommandReply_15Bytes((short) 0x8F, (short) 15, CALCommandType.REPLY),
  CALCommandReply_16Bytes((short) 0x90, (short) 16, CALCommandType.REPLY),
  CALCommandReply_17Bytes((short) 0x91, (short) 17, CALCommandType.REPLY),
  CALCommandReply_18Bytes((short) 0x92, (short) 18, CALCommandType.REPLY),
  CALCommandReply_19Bytes((short) 0x93, (short) 19, CALCommandType.REPLY),
  CALCommandReply_20Bytes((short) 0x94, (short) 20, CALCommandType.REPLY),
  CALCommandReply_21Bytes((short) 0x95, (short) 21, CALCommandType.REPLY),
  CALCommandReply_22Bytes((short) 0x96, (short) 22, CALCommandType.REPLY),
  CALCommandReply_23Bytes((short) 0x97, (short) 23, CALCommandType.REPLY),
  CALCommandReply_24Bytes((short) 0x98, (short) 24, CALCommandType.REPLY),
  CALCommandReply_25Bytes((short) 0x99, (short) 25, CALCommandType.REPLY),
  CALCommandReply_26Bytes((short) 0x9A, (short) 26, CALCommandType.REPLY),
  CALCommandReply_27Bytes((short) 0x9B, (short) 27, CALCommandType.REPLY),
  CALCommandReply_28Bytes((short) 0x9C, (short) 28, CALCommandType.REPLY),
  CALCommandReply_29Bytes((short) 0x9D, (short) 29, CALCommandType.REPLY),
  CALCommandReply_30Bytes((short) 0x9E, (short) 30, CALCommandType.REPLY),
  CALCommandReply_31Bytes((short) 0x9F, (short) 31, CALCommandType.REPLY),
  CALCommandWrite_0Bytes((short) 0xA0, (short) 0, CALCommandType.WRITE),
  CALCommandWrite_1Bytes((short) 0xA1, (short) 1, CALCommandType.WRITE),
  CALCommandWrite_2Bytes((short) 0xA2, (short) 2, CALCommandType.WRITE),
  CALCommandWrite_3Bytes((short) 0xA3, (short) 3, CALCommandType.WRITE),
  CALCommandWrite_4Bytes((short) 0xA4, (short) 4, CALCommandType.WRITE),
  CALCommandWrite_5Bytes((short) 0xA5, (short) 5, CALCommandType.WRITE),
  CALCommandWrite_6Bytes((short) 0xA6, (short) 6, CALCommandType.WRITE),
  CALCommandWrite_7Bytes((short) 0xA7, (short) 7, CALCommandType.WRITE),
  CALCommandWrite_8Bytes((short) 0xA8, (short) 8, CALCommandType.WRITE),
  CALCommandWrite_9Bytes((short) 0xA9, (short) 9, CALCommandType.WRITE),
  CALCommandWrite_10Bytes((short) 0xAA, (short) 10, CALCommandType.WRITE),
  CALCommandWrite_11Bytes((short) 0xAB, (short) 11, CALCommandType.WRITE),
  CALCommandWrite_12Bytes((short) 0xAC, (short) 12, CALCommandType.WRITE),
  CALCommandWrite_13Bytes((short) 0xAD, (short) 13, CALCommandType.WRITE),
  CALCommandWrite_14Bytes((short) 0xAE, (short) 14, CALCommandType.WRITE),
  CALCommandWrite_15Bytes((short) 0xAF, (short) 15, CALCommandType.WRITE),
  CALCommandStatus_0Bytes((short) 0xC0, (short) 0, CALCommandType.STATUS),
  CALCommandStatus_1Bytes((short) 0xC1, (short) 1, CALCommandType.STATUS),
  CALCommandStatus_2Bytes((short) 0xC2, (short) 2, CALCommandType.STATUS),
  CALCommandStatus_3Bytes((short) 0xC3, (short) 3, CALCommandType.STATUS),
  CALCommandStatus_4Bytes((short) 0xC4, (short) 4, CALCommandType.STATUS),
  CALCommandStatus_5Bytes((short) 0xC5, (short) 5, CALCommandType.STATUS),
  CALCommandStatus_6Bytes((short) 0xC6, (short) 6, CALCommandType.STATUS),
  CALCommandStatus_7Bytes((short) 0xC7, (short) 7, CALCommandType.STATUS),
  CALCommandStatus_8Bytes((short) 0xC8, (short) 8, CALCommandType.STATUS),
  CALCommandStatus_9Bytes((short) 0xC9, (short) 9, CALCommandType.STATUS),
  CALCommandStatus_10Bytes((short) 0xCA, (short) 10, CALCommandType.STATUS),
  CALCommandStatus_11Bytes((short) 0xCB, (short) 11, CALCommandType.STATUS),
  CALCommandStatus_12Bytes((short) 0xCC, (short) 12, CALCommandType.STATUS),
  CALCommandStatus_13Bytes((short) 0xCD, (short) 13, CALCommandType.STATUS),
  CALCommandStatus_14Bytes((short) 0xCE, (short) 14, CALCommandType.STATUS),
  CALCommandStatus_15Bytes((short) 0xCF, (short) 15, CALCommandType.STATUS),
  CALCommandStatus_16Bytes((short) 0xD0, (short) 16, CALCommandType.STATUS),
  CALCommandStatus_17Bytes((short) 0xD1, (short) 17, CALCommandType.STATUS),
  CALCommandStatus_18Bytes((short) 0xD2, (short) 18, CALCommandType.STATUS),
  CALCommandStatus_19Bytes((short) 0xD3, (short) 19, CALCommandType.STATUS),
  CALCommandStatus_20Bytes((short) 0xD4, (short) 20, CALCommandType.STATUS),
  CALCommandStatus_21Bytes((short) 0xD5, (short) 21, CALCommandType.STATUS),
  CALCommandStatus_22Bytes((short) 0xD6, (short) 22, CALCommandType.STATUS),
  CALCommandStatus_23Bytes((short) 0xD7, (short) 23, CALCommandType.STATUS),
  CALCommandStatus_24Bytes((short) 0xD8, (short) 24, CALCommandType.STATUS),
  CALCommandStatus_25Bytes((short) 0xD9, (short) 25, CALCommandType.STATUS),
  CALCommandStatus_26Bytes((short) 0xDA, (short) 26, CALCommandType.STATUS),
  CALCommandStatus_27Bytes((short) 0xDB, (short) 27, CALCommandType.STATUS),
  CALCommandStatus_28Bytes((short) 0xDC, (short) 28, CALCommandType.STATUS),
  CALCommandStatus_29Bytes((short) 0xDD, (short) 29, CALCommandType.STATUS),
  CALCommandStatus_30Bytes((short) 0xDE, (short) 30, CALCommandType.STATUS),
  CALCommandStatus_31Bytes((short) 0xDF, (short) 31, CALCommandType.STATUS),
  CALCommandStatusExtended_0Bytes((short) 0xE0, (short) 0, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_1Bytes((short) 0xE1, (short) 1, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_2Bytes((short) 0xE2, (short) 2, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_3Bytes((short) 0xE3, (short) 3, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_4Bytes((short) 0xE4, (short) 4, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_5Bytes((short) 0xE5, (short) 5, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_6Bytes((short) 0xE6, (short) 6, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_7Bytes((short) 0xE7, (short) 7, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_8Bytes((short) 0xE8, (short) 8, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_9Bytes((short) 0xE9, (short) 9, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_10Bytes((short) 0xEA, (short) 10, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_11Bytes((short) 0xEB, (short) 11, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_12Bytes((short) 0xEC, (short) 12, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_13Bytes((short) 0xED, (short) 13, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_14Bytes((short) 0xEE, (short) 14, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_15Bytes((short) 0xEF, (short) 15, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_16Bytes((short) 0xF0, (short) 16, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_17Bytes((short) 0xF1, (short) 17, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_18Bytes((short) 0xF2, (short) 18, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_19Bytes((short) 0xF3, (short) 19, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_20Bytes((short) 0xF4, (short) 20, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_21Bytes((short) 0xF5, (short) 21, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_22Bytes((short) 0xF6, (short) 22, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_23Bytes((short) 0xF7, (short) 23, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_24Bytes((short) 0xF8, (short) 24, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_25Bytes((short) 0xF9, (short) 25, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_26Bytes((short) 0xFA, (short) 26, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_27Bytes((short) 0xFB, (short) 27, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_28Bytes((short) 0xFC, (short) 28, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_29Bytes((short) 0xFD, (short) 29, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_30Bytes((short) 0xFE, (short) 30, CALCommandType.STATUS_EXTENDED),
  CALCommandStatusExtended_31Bytes((short) 0xFF, (short) 31, CALCommandType.STATUS_EXTENDED);
  private static final Map<Short, CALCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (CALCommandTypeContainer value : CALCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private short value;
  private short numBytes;
  private CALCommandType commandType;

  CALCommandTypeContainer(short value, short numBytes, CALCommandType commandType) {
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

  public static CALCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<CALCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<CALCommandTypeContainer> _values = new ArrayList();
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public CALCommandType getCommandType() {
    return commandType;
  }

  public static CALCommandTypeContainer firstEnumForFieldCommandType(CALCommandType fieldValue) {
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<CALCommandTypeContainer> enumsForFieldCommandType(CALCommandType fieldValue) {
    List<CALCommandTypeContainer> _values = new ArrayList();
    for (CALCommandTypeContainer _val : CALCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static CALCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
