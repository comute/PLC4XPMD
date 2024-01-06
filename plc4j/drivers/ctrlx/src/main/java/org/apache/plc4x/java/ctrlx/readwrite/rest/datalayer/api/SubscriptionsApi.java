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

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.api;

import com.fasterxml.jackson.core.type.TypeReference;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiClient;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiException;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Configuration;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Pair;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.ChangeSubscriptionRequest;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.CreateSubscriptionRequest;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.SSEEvent;

import java.util.*;

@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class SubscriptionsApi {
  private ApiClient apiClient;

  public SubscriptionsApi() {
    this(Configuration.getDefaultApiClient());
  }

  public SubscriptionsApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Change node list of existing subscription.
   * Change node list of existing subscription. All opened SSE connections for this subscription are changed.
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @param changeSubscriptionRequest The data for the new subscription (optional)
   * @throws ApiException if fails to make API call
   */
  public void changeSubscription(String subscriptionName, ChangeSubscriptionRequest changeSubscriptionRequest) throws ApiException {
    Object localVarPostBody = changeSubscriptionRequest;
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling changeSubscription");
    }
    
    // create path and map variables
    String localVarPath = "/events/{SubscriptionName}"
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    apiClient.invokeAPI(
        localVarPath,
        "PUT",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        null
    );
  }
  /**
   * Create a new subscription as separate resource
   * Create a new subscription. You can later make GET request on this created subscription to start the subscription. If subscription is not used within one minute, subscription will be deleted automatically. You can start a created subscription multiple times. &lt;br&gt;&lt;br&gt; To create a subscription in this way allows to setup more complex settings within the subscription (queueing, data change filter, ...).
   * @param createSubscriptionRequest The data for the new subscription (optional)
   * @throws ApiException if fails to make API call
   */
  public void createSubscription(CreateSubscriptionRequest createSubscriptionRequest) throws ApiException {
    Object localVarPostBody = createSubscriptionRequest;
    
    // create path and map variables
    String localVarPath = "/events";

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    apiClient.invokeAPI(
        localVarPath,
        "POST",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        null
    );
  }
  /**
   * Deletes an existing subscription resource.
   * Deletes an existing subscription. A subscription can only be deleted if there are no open SSE connections for this SubscriptionName.
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @throws ApiException if fails to make API call
   */
  public void deleteSubscription(String subscriptionName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling deleteSubscription");
    }
    
    // create path and map variables
    String localVarPath = "/events/{SubscriptionName}"
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    apiClient.invokeAPI(
        localVarPath,
        "DELETE",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        null
    );
  }
  /**
   * Create a new subscription and starts server send event (SSE) stream.
   * Create a new subscription and starts server send event (SSE) stream. You can specify a bunch of paths to be observed for changes. &lt;br&gt; &lt;br&gt; Unfortunately, SSE is not compatible to OpenAPI specs and can therefore not be properly specified within this API description. An SSE event contains three parts separated by linebreaks: event, id and data. Different events are separated by empty lines. The event field can be one of these types: keepalive, update, browse, metadata, error. The data field is populated with the JSON object defined below. &lt;br&gt; &lt;br&gt; At startup you will get initial events for each node. Afterwards you will only get events if the value of the node changes.
   * @param nodes Comma sepearted list of nodes (required)
   * @param publishIntervalMs Publish interval for notifications. Internal poll interval. (optional, default to 1000)
   * @return List&lt;SSEEvent&gt;
   * @throws ApiException if fails to make API call
   */
  public List<SSEEvent> eventsGet(String nodes, Integer publishIntervalMs) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'nodes' is set
    if (nodes == null) {
      throw new ApiException(400, "Missing the required parameter 'nodes' when calling eventsGet");
    }
    
    // create path and map variables
    String localVarPath = "/events";

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    localVarQueryParams.addAll(apiClient.parameterToPair("nodes", nodes));
    localVarQueryParams.addAll(apiClient.parameterToPair("publishIntervalMs", publishIntervalMs));
    
    
    
    final String[] localVarAccepts = {
      "text/event-stream", "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<List<SSEEvent>> localVarReturnType = new TypeReference<List<SSEEvent>>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
  /**
   * Starts a server send event stream (SSE) on a given subscription
   * Starts a server send event stream (SSE) for a given subscription. You have to create this subscription before. &lt;br&gt; &lt;br&gt; Unfortunately, SSE is not compatible to OpenAPI specs and can therefore not be properly specified within this API description. An SSE event contains three parts separated by linebreaks: event, id and data. Different events are separated by empty lines. The event field can be one of these types: keepalive, update, browse, metadata, error. The data field is populated with the JSON object defined below. &lt;br&gt; &lt;br&gt; At startup you will get initial events for each node. Afterwards you will only get events if the value of the node changes.
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @return List&lt;SSEEvent&gt;
   * @throws ApiException if fails to make API call
   */
  public List<SSEEvent> getCreatedSubscription(String subscriptionName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling getCreatedSubscription");
    }
    
    // create path and map variables
    String localVarPath = "/events/{SubscriptionName}"
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "text/event-stream", "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<List<SSEEvent>> localVarReturnType = new TypeReference<List<SSEEvent>>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
}
