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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableAddress {
  AddressSpaceFileType_Size((int) 11596L),
  AddressSpaceFileType_OpenCount((int) 11599L),
  AddressSpaceFileType_Open_InputArguments((int) 11601L),
  AddressSpaceFileType_Open_OutputArguments((int) 11602L),
  AddressSpaceFileType_Close_InputArguments((int) 11604L),
  AddressSpaceFileType_Read_InputArguments((int) 11606L),
  AddressSpaceFileType_Read_OutputArguments((int) 11607L),
  AddressSpaceFileType_Write_InputArguments((int) 11609L),
  AddressSpaceFileType_GetPosition_InputArguments((int) 11611L),
  AddressSpaceFileType_GetPosition_OutputArguments((int) 11612L),
  AddressSpaceFileType_SetPosition_InputArguments((int) 11614L),
  AddressSpaceFileType_Writable((int) 12688L),
  AddressSpaceFileType_UserWritable((int) 12689L),
  AddressSpaceFileType_MimeType((int) 13398L),
  AddressSpaceFileType_MaxByteStringLength((int) 24245L),
  AddressSpaceFileType_LastModifiedTime((int) 25201L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableAddress> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableAddress value : OpcuaNodeIdServicesVariableAddress.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableAddress(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableAddress enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}