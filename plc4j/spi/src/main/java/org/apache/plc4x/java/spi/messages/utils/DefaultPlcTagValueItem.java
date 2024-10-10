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
package org.apache.plc4x.java.spi.messages.utils;

import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.utils.Serializable;

public class DefaultPlcTagValueItem<T extends PlcTag> extends DefaultPlcTagItem<T> implements PlcTagValueItem<T>, Serializable {

    private final PlcValue value;

    public DefaultPlcTagValueItem(T tag, PlcValue value) {
        super(tag);
        this.value = value;
    }

    public PlcValue getValue() {
        return value;
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        writeBuffer.pushContext("PlcTagValueItem");
        if(getTag() instanceof Serializable) {
            writeBuffer.pushContext("tag");
            writeBuffer.writeSerializable((Serializable) getTag());
            writeBuffer.popContext("tag");
        }
        if(value instanceof Serializable) {
            writeBuffer.pushContext("value");
            writeBuffer.writeSerializable((Serializable) value);
            writeBuffer.popContext("value");
        }
        writeBuffer.popContext("PlcTagValueItem");
    }

}
