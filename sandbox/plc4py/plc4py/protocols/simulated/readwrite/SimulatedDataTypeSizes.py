#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from enum import IntEnum
    
class SimulatedDataTypeSizes(IntEnum):
    BOOL: int = (1 , int(1) )
    BYTE: int = (2 , int(1) )
    WORD: int = (3 , int(2) )
    DWORD: int = (4 , int(4) )
    LWORD: int = (5 , int(8) )
    SINT: int = (6 , int(1) )
    INT: int = (7 , int(2) )
    DINT: int = (8 , int(4) )
    LINT: int = (9 , int(8) )
    USINT: int = (10 , int(1) )
    UINT: int = (11 , int(2) )
    UDINT: int = (12 , int(4) )
    ULINT: int = (13 , int(8) )
    REAL: int = (14 , int(4) )
    LREAL: int = (15 , int(8) )
    TIME: int = (16 , int(8) )
    LTIME: int = (17 , int(8) )
    DATE: int = (18 , int(8) )
    LDATE: int = (19 , int(8) )
    TIME_OF_DAY: int = (20 , int(8) )
    LTIME_OF_DAY: int = (21 , int(8) )
    DATE_AND_TIME: int = (22 , int(8) )
    LDATE_AND_TIME: int = (23 , int(8) )
    CHAR: int = (24 , int(1) )
    WCHAR: int = (25 , int(2) )
    STRING: int = (26 , int(255) )
    WSTRING: int = (27 , int(127) )

    def __new__(cls, value, data_type_size ):
        obj = object.__new__(cls)
        obj._value_ = value
        obj.dataTypeSize = dataTypeSize
        return obj



