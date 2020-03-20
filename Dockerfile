#  Licensed to the Apache Software Foundation (ASF) under one or more
#  contributor license agreements.  See the NOTICE file distributed with
#  this work for additional information regarding copyright ownership.
#  The ASF licenses this file to You under the Apache License, Version 2.0
#  (the "License"); you may not use this file except in compliance with
#  the License.  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

##########################################################################################
# Build PLC4X
##########################################################################################

# This is the image we'll use to execute the build (and give it the name 'build').
# (This image is based on Ubuntu)
# Fixed version of this in order to have a fixed JDK version
FROM openjdk:8 as build

# Install some stuff we need to run the build
RUN apt update -y

# Install general purpose tools
RUN apt-get install libpcap-dev -y

# Required for the general build
RUN apt install -y git

# Copy the project into the docker container
COPY . /ws/

# Change the working directory (where commands are executed) into the new "ws" directory
WORKDIR /ws

# Tell Maven to fetch all needed dependencies first, so they can get cached
# (Tried a patched version of the plugin to allow exclusion of inner artifacts.
# See https://issues.apache.org/jira/browse/MDEP-568 for details)
#RUN ./mvnw -P with-sandbox,with-cpp,with-boost,with-dotnet,with-python,with-proxies,with-logstash com.offbytwo.maven.plugins:maven-dependency-plugin:3.1.1.MDEP568:go-offline -DexcludeGroupIds=org.apache.plc4x,org.apache.plc4x.examples,org.apache.plc4x.sandbox
RUN ./mvnw -P with-sandbox,with-boost,with-dotnet,with-python,with-proxies,with-logstash com.offbytwo.maven.plugins:maven-dependency-plugin:3.1.1.MDEP568:go-offline -DexcludeGroupIds=org.apache.plc4x,org.apache.plc4x.examples,org.apache.plc4x.sandbox

# Build everything with all tests
#RUN ./mvnw -P with-sandbox,with-cpp,with-boost,with-dotnet,with-python,with-proxies,with-logstash install
RUN ./mvnw install -DskipTests

# Get the version of the project and save it in a local file on the container
RUN ./mvnw org.apache.maven.plugins:maven-help-plugin:3.2.0:evaluate -Dexpression=project.version -DforceStdout -q -pl . > project_version

##########################################################################################
# Build a demo container
##########################################################################################

# Move the file to a place we can reference it from without a version
RUN PROJECT_VERSION=`cat project_version`; mv plc4j/examples/hello-storage-elasticsearch/target/plc4j-hello-storage-elasticsearch-$PROJECT_VERSION-uber-jar.jar plc4xdemo.jar

# Build a highly optimized JRE
FROM alpine:3.10 as packager

# Install regular JDK
RUN apk update
RUN apk --no-cache add openjdk11-jdk openjdk11-jmods

# build minimal JRE
ENV JAVA_MINIMAL="/opt/java-minimal"
RUN /usr/lib/jvm/java-11-openjdk/bin/jlink \
    --verbose \
    --add-modules \
        java.base,java.sql,java.naming,java.desktop,java.management,java.security.jgss,java.instrument \
    --compress 2 --strip-debug --no-header-files --no-man-pages \
    --release-info="add:IMPLEMENTOR=radistao:IMPLEMENTOR_VERSION=radistao_JRE" \
    --output "$JAVA_MINIMAL"

# Now create an actuall deployment container
FROM alpine:3.10

# Install our optimized JRE
ENV JAVA_HOME=/opt/java-minimal
ENV PATH="$PATH:$JAVA_HOME/bin"
COPY --from=packager "$JAVA_HOME" "$JAVA_HOME"

# Prepare the demo by copying the example artifact from the 'build' container into this new one.
COPY --from=build /ws/plc4xdemo.jar /plc4xdemo.jar

# Let runtime know which ports we will be listening on
EXPOSE 9200 9300

# Allow for extra options to be passed to the jar using PLC4X_OPTIONS env variable
ENV PLC4X_OPTIONS ""

# This will be executed as soon as the container is started.
ENTRYPOINT ["sh", "-c", "[ -f /run/plc4xdemo.env ] && . /run/plc4xdemo.env ; java -jar /plc4xdemo.jar $PLC4X_OPTIONS"]
