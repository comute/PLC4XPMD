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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class AccessUser implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final OptionalProperty description;
  protected final WritableProperty globalIdentifier;
  protected final ReadableProperty statusFlags;
  protected final ReadableProperty reliability;
  protected final ReadableProperty userType;
  protected final OptionalProperty userName;
  protected final OptionalProperty userExternalIdentifier;
  protected final OptionalProperty userInformationReference;
  protected final OptionalProperty members;
  protected final OptionalProperty memberOf;
  protected final ReadableProperty credentials;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final ReadableProperty propertyList;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public AccessUser(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      OptionalProperty description,
      WritableProperty globalIdentifier,
      ReadableProperty statusFlags,
      ReadableProperty reliability,
      ReadableProperty userType,
      OptionalProperty userName,
      OptionalProperty userExternalIdentifier,
      OptionalProperty userInformationReference,
      OptionalProperty members,
      OptionalProperty memberOf,
      ReadableProperty credentials,
      OptionalProperty reliabilityEvaluationInhibit,
      ReadableProperty propertyList,
      OptionalProperty tags,
      OptionalProperty profileLocation,
      OptionalProperty profileName) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.objectName = objectName;
    this.objectType = objectType;
    this.description = description;
    this.globalIdentifier = globalIdentifier;
    this.statusFlags = statusFlags;
    this.reliability = reliability;
    this.userType = userType;
    this.userName = userName;
    this.userExternalIdentifier = userExternalIdentifier;
    this.userInformationReference = userInformationReference;
    this.members = members;
    this.memberOf = memberOf;
    this.credentials = credentials;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.propertyList = propertyList;
    this.tags = tags;
    this.profileLocation = profileLocation;
    this.profileName = profileName;
  }

  public ReadableProperty getObjectIdentifier() {
    return objectIdentifier;
  }

  public ReadableProperty getObjectName() {
    return objectName;
  }

  public ReadableProperty getObjectType() {
    return objectType;
  }

  public OptionalProperty getDescription() {
    return description;
  }

  public WritableProperty getGlobalIdentifier() {
    return globalIdentifier;
  }

  public ReadableProperty getStatusFlags() {
    return statusFlags;
  }

  public ReadableProperty getReliability() {
    return reliability;
  }

  public ReadableProperty getUserType() {
    return userType;
  }

  public OptionalProperty getUserName() {
    return userName;
  }

  public OptionalProperty getUserExternalIdentifier() {
    return userExternalIdentifier;
  }

  public OptionalProperty getUserInformationReference() {
    return userInformationReference;
  }

  public OptionalProperty getMembers() {
    return members;
  }

  public OptionalProperty getMemberOf() {
    return memberOf;
  }

  public ReadableProperty getCredentials() {
    return credentials;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public OptionalProperty getTags() {
    return tags;
  }

  public OptionalProperty getProfileLocation() {
    return profileLocation;
  }

  public OptionalProperty getProfileName() {
    return profileName;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AccessUser");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (objectName)
    writeSimpleField("objectName", objectName, writeComplex(writeBuffer));

    // Simple Field (objectType)
    writeSimpleField("objectType", objectType, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (globalIdentifier)
    writeSimpleField("globalIdentifier", globalIdentifier, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (userType)
    writeSimpleField("userType", userType, writeComplex(writeBuffer));

    // Simple Field (userName)
    writeSimpleField("userName", userName, writeComplex(writeBuffer));

    // Simple Field (userExternalIdentifier)
    writeSimpleField("userExternalIdentifier", userExternalIdentifier, writeComplex(writeBuffer));

    // Simple Field (userInformationReference)
    writeSimpleField(
        "userInformationReference", userInformationReference, writeComplex(writeBuffer));

    // Simple Field (members)
    writeSimpleField("members", members, writeComplex(writeBuffer));

    // Simple Field (memberOf)
    writeSimpleField("memberOf", memberOf, writeComplex(writeBuffer));

    // Simple Field (credentials)
    writeSimpleField("credentials", credentials, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

    // Simple Field (tags)
    writeSimpleField("tags", tags, writeComplex(writeBuffer));

    // Simple Field (profileLocation)
    writeSimpleField("profileLocation", profileLocation, writeComplex(writeBuffer));

    // Simple Field (profileName)
    writeSimpleField("profileName", profileName, writeComplex(writeBuffer));

    writeBuffer.popContext("AccessUser");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AccessUser _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (objectName)
    lengthInBits += objectName.getLengthInBits();

    // Simple field (objectType)
    lengthInBits += objectType.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (globalIdentifier)
    lengthInBits += globalIdentifier.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (userType)
    lengthInBits += userType.getLengthInBits();

    // Simple field (userName)
    lengthInBits += userName.getLengthInBits();

    // Simple field (userExternalIdentifier)
    lengthInBits += userExternalIdentifier.getLengthInBits();

    // Simple field (userInformationReference)
    lengthInBits += userInformationReference.getLengthInBits();

    // Simple field (members)
    lengthInBits += members.getLengthInBits();

    // Simple field (memberOf)
    lengthInBits += memberOf.getLengthInBits();

    // Simple field (credentials)
    lengthInBits += credentials.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

    // Simple field (tags)
    lengthInBits += tags.getLengthInBits();

    // Simple field (profileLocation)
    lengthInBits += profileLocation.getLengthInBits();

    // Simple field (profileName)
    lengthInBits += profileName.getLengthInBits();

    return lengthInBits;
  }

  public static AccessUser staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AccessUser");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReadableProperty objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectIdentifier")),
                readBuffer));

    ReadableProperty objectName =
        readSimpleField(
            "objectName",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    ReadableProperty objectType =
        readSimpleField(
            "objectType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectType")),
                readBuffer));

    OptionalProperty description =
        readSimpleField(
            "description",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    WritableProperty globalIdentifier =
        readSimpleField(
            "globalIdentifier",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    ReadableProperty statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetStatusFlags")),
                readBuffer));

    ReadableProperty reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetReliability")),
                readBuffer));

    ReadableProperty userType =
        readSimpleField(
            "userType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetAccessUserType")),
                readBuffer));

    OptionalProperty userName =
        readSimpleField(
            "userName",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty userExternalIdentifier =
        readSimpleField(
            "userExternalIdentifier",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty userInformationReference =
        readSimpleField(
            "userInformationReference",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty members =
        readSimpleField(
            "members",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetLIST of BACnetDeviceObjectReference")),
                readBuffer));

    OptionalProperty memberOf =
        readSimpleField(
            "memberOf",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetLIST of BACnetDeviceObjectReference")),
                readBuffer));

    ReadableProperty credentials =
        readSimpleField(
            "credentials",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetLIST of BACnetDeviceObjectReference")),
                readBuffer));

    OptionalProperty reliabilityEvaluationInhibit =
        readSimpleField(
            "reliabilityEvaluationInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty propertyList =
        readSimpleField(
            "propertyList",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetPropertyIdentifier")),
                readBuffer));

    OptionalProperty tags =
        readSimpleField(
            "tags",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetNameValue")),
                readBuffer));

    OptionalProperty profileLocation =
        readSimpleField(
            "profileLocation",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty profileName =
        readSimpleField(
            "profileName",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    readBuffer.closeContext("AccessUser");
    // Create the instance
    AccessUser _accessUser;
    _accessUser =
        new AccessUser(
            objectIdentifier,
            objectName,
            objectType,
            description,
            globalIdentifier,
            statusFlags,
            reliability,
            userType,
            userName,
            userExternalIdentifier,
            userInformationReference,
            members,
            memberOf,
            credentials,
            reliabilityEvaluationInhibit,
            propertyList,
            tags,
            profileLocation,
            profileName);
    return _accessUser;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AccessUser)) {
      return false;
    }
    AccessUser that = (AccessUser) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getDescription() == that.getDescription())
        && (getGlobalIdentifier() == that.getGlobalIdentifier())
        && (getStatusFlags() == that.getStatusFlags())
        && (getReliability() == that.getReliability())
        && (getUserType() == that.getUserType())
        && (getUserName() == that.getUserName())
        && (getUserExternalIdentifier() == that.getUserExternalIdentifier())
        && (getUserInformationReference() == that.getUserInformationReference())
        && (getMembers() == that.getMembers())
        && (getMemberOf() == that.getMemberOf())
        && (getCredentials() == that.getCredentials())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getPropertyList() == that.getPropertyList())
        && (getTags() == that.getTags())
        && (getProfileLocation() == that.getProfileLocation())
        && (getProfileName() == that.getProfileName())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getObjectIdentifier(),
        getObjectName(),
        getObjectType(),
        getDescription(),
        getGlobalIdentifier(),
        getStatusFlags(),
        getReliability(),
        getUserType(),
        getUserName(),
        getUserExternalIdentifier(),
        getUserInformationReference(),
        getMembers(),
        getMemberOf(),
        getCredentials(),
        getReliabilityEvaluationInhibit(),
        getPropertyList(),
        getTags(),
        getProfileLocation(),
        getProfileName());
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