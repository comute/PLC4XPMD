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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class QueryDataSet extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 579;
  }

  // Properties.
  protected final ExpandedNodeId nodeId;
  protected final ExpandedNodeId typeDefinitionNode;
  protected final List<Variant> values;

  public QueryDataSet(
      ExpandedNodeId nodeId, ExpandedNodeId typeDefinitionNode, List<Variant> values) {
    super();
    this.nodeId = nodeId;
    this.typeDefinitionNode = typeDefinitionNode;
    this.values = values;
  }

  public ExpandedNodeId getNodeId() {
    return nodeId;
  }

  public ExpandedNodeId getTypeDefinitionNode() {
    return typeDefinitionNode;
  }

  public List<Variant> getValues() {
    return values;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QueryDataSet");

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, writeComplex(writeBuffer));

    // Simple Field (typeDefinitionNode)
    writeSimpleField("typeDefinitionNode", typeDefinitionNode, writeComplex(writeBuffer));

    // Implicit Field (noOfValues) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int noOfValues = (int) ((((getValues()) == (null)) ? -(1) : COUNT(getValues())));
    writeImplicitField("noOfValues", noOfValues, writeSignedInt(writeBuffer, 32));

    // Array Field (values)
    writeComplexTypeArrayField("values", values, writeBuffer);

    writeBuffer.popContext("QueryDataSet");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    QueryDataSet _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Simple field (typeDefinitionNode)
    lengthInBits += typeDefinitionNode.getLengthInBits();

    // Implicit Field (noOfValues)
    lengthInBits += 32;

    // Array field
    if (values != null) {
      int i = 0;
      for (Variant element : values) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= values.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("QueryDataSet");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExpandedNodeId nodeId =
        readSimpleField(
            "nodeId", readComplex(() -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    ExpandedNodeId typeDefinitionNode =
        readSimpleField(
            "typeDefinitionNode",
            readComplex(() -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    int noOfValues = readImplicitField("noOfValues", readSignedInt(readBuffer, 32));

    List<Variant> values =
        readCountArrayField(
            "values", readComplex(() -> Variant.staticParse(readBuffer), readBuffer), noOfValues);

    readBuffer.closeContext("QueryDataSet");
    // Create the instance
    return new QueryDataSetBuilderImpl(nodeId, typeDefinitionNode, values);
  }

  public static class QueryDataSetBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExpandedNodeId nodeId;
    private final ExpandedNodeId typeDefinitionNode;
    private final List<Variant> values;

    public QueryDataSetBuilderImpl(
        ExpandedNodeId nodeId, ExpandedNodeId typeDefinitionNode, List<Variant> values) {
      this.nodeId = nodeId;
      this.typeDefinitionNode = typeDefinitionNode;
      this.values = values;
    }

    public QueryDataSet build() {
      QueryDataSet queryDataSet = new QueryDataSet(nodeId, typeDefinitionNode, values);
      return queryDataSet;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QueryDataSet)) {
      return false;
    }
    QueryDataSet that = (QueryDataSet) o;
    return (getNodeId() == that.getNodeId())
        && (getTypeDefinitionNode() == that.getTypeDefinitionNode())
        && (getValues() == that.getValues())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNodeId(), getTypeDefinitionNode(), getValues());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
