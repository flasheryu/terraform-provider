package cloudapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteLogConfig invokes the cloudapi.DeleteLogConfig API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deletelogconfig.html
func (client *Client) DeleteLogConfig(request *DeleteLogConfigRequest) (response *DeleteLogConfigResponse, err error) {
	response = CreateDeleteLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLogConfigWithChan invokes the cloudapi.DeleteLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletelogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogConfigWithChan(request *DeleteLogConfigRequest) (<-chan *DeleteLogConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLogConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteLogConfigWithCallback invokes the cloudapi.DeleteLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletelogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogConfigWithCallback(request *DeleteLogConfigRequest, callback func(response *DeleteLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLogConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLogConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteLogConfigRequest is the request struct for api DeleteLogConfig
type DeleteLogConfigRequest struct {
	*requests.RpcRequest
	LogType string `position:"Query" name:"LogType"`
}

// DeleteLogConfigResponse is the response struct for api DeleteLogConfig
type DeleteLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLogConfigRequest creates a request to invoke DeleteLogConfig API
func CreateDeleteLogConfigRequest() (request *DeleteLogConfigRequest) {
	request = &DeleteLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteLogConfig", "apigateway", "openAPI")
	return
}

// CreateDeleteLogConfigResponse creates a response to parse from DeleteLogConfig response
func CreateDeleteLogConfigResponse() (response *DeleteLogConfigResponse) {
	response = &DeleteLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
