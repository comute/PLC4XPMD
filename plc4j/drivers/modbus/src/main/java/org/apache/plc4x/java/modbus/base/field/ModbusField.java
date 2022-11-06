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
package org.apache.plc4x.java.modbus.base.field;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.spi.utils.Serializable;

public abstract class ModbusField implements PlcField, Serializable {

    public static ModbusField of(String addressString) {
        if (ModbusFieldCoil.matches(addressString)) {
            return ModbusFieldCoil.of(addressString);
        }
        if (ModbusFieldDiscreteInput.matches(addressString)) {
            return ModbusFieldDiscreteInput.of(addressString);
        }
        if (ModbusFieldHoldingRegister.matches(addressString)) {
            return ModbusFieldHoldingRegister.of(addressString);
        }
        if (ModbusFieldInputRegister.matches(addressString)) {
            return ModbusFieldInputRegister.of(addressString);
        }
        if (ModbusExtendedRegister.matches(addressString)) {
            return ModbusExtendedRegister.of(addressString);
        }
        throw new PlcInvalidFieldException("Unable to parse address: " + addressString);
    }

}
