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

package org.apache.plc4x.java.profinet.device;

import org.apache.plc4x.java.api.messages.PlcDiscoveryItem;
import org.apache.plc4x.java.api.messages.PlcDiscoveryItemHandler;

import java.util.HashMap;

public class ProfinetDeviceMessageHandler implements PlcDiscoveryItemHandler {

    private HashMap<String, ProfinetDevice> configuredDevices;

    @Override
    public void handle(PlcDiscoveryItem discoveryItem) {
        String macAddress = discoveryItem.getOptions().get("macAddress").replace(":", "").toUpperCase();
        if (configuredDevices.containsKey(macAddress)) {
            configuredDevices.get(macAddress).handle(discoveryItem);
        }
    }

    public void setConfiguredDevices(HashMap<String, ProfinetDevice> configuredDevices) {
        this.configuredDevices = configuredDevices;
    }
}
