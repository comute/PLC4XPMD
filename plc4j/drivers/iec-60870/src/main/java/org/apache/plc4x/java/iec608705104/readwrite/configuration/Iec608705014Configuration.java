/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.iec608705104.readwrite.configuration;

import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.configuration.annotations.ComplexConfigurationParameter;
import org.apache.plc4x.java.spi.configuration.annotations.ConfigurationParameter;
import org.apache.plc4x.java.spi.configuration.annotations.defaults.IntDefaultValue;
import org.apache.plc4x.java.spi.transport.TransportConfiguration;
import org.apache.plc4x.java.spi.transport.TransportConfigurationProvider;

public class Iec608705014Configuration implements Configuration, TransportConfigurationProvider {

    @ConfigurationParameter("timeout-request")
    @IntDefaultValue(4000)
    protected int timeoutRequest;

    @ComplexConfigurationParameter(prefix = "tcp", defaultOverrides = {}, requiredOverrides = {})
    private Iec608705014TcpTransportConfiguration tcpTransportConfiguration;

    public int getTimeoutRequest() {
        return timeoutRequest;
    }

    public void setTimeoutRequest(int timeoutRequest) {
        this.timeoutRequest = timeoutRequest;
    }

    public Iec608705014TcpTransportConfiguration getTcpTransportConfiguration() {
        return tcpTransportConfiguration;
    }

    public void setTcpTransportConfiguration(Iec608705014TcpTransportConfiguration tcpTransportConfiguration) {
        this.tcpTransportConfiguration = tcpTransportConfiguration;
    }

    @Override
    public TransportConfiguration getTransportConfiguration(String transportCode) {
        switch (transportCode) {
            case "tcp":
                return tcpTransportConfiguration;
        }
        return null;
    }

}
