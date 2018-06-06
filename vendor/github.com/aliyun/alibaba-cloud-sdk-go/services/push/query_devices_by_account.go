package push

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

// QueryDevicesByAccount invokes the push.QueryDevicesByAccount API synchronously
// api document: https://help.aliyun.com/api/push/querydevicesbyaccount.html
func (client *Client) QueryDevicesByAccount(request *QueryDevicesByAccountRequest) (response *QueryDevicesByAccountResponse, err error) {
	response = CreateQueryDevicesByAccountResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDevicesByAccountWithChan invokes the push.QueryDevicesByAccount API asynchronously
// api document: https://help.aliyun.com/api/push/querydevicesbyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicesByAccountWithChan(request *QueryDevicesByAccountRequest) (<-chan *QueryDevicesByAccountResponse, <-chan error) {
	responseChan := make(chan *QueryDevicesByAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDevicesByAccount(request)
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

// QueryDevicesByAccountWithCallback invokes the push.QueryDevicesByAccount API asynchronously
// api document: https://help.aliyun.com/api/push/querydevicesbyaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicesByAccountWithCallback(request *QueryDevicesByAccountRequest, callback func(response *QueryDevicesByAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDevicesByAccountResponse
		var err error
		defer close(result)
		response, err = client.QueryDevicesByAccount(request)
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

// QueryDevicesByAccountRequest is the request struct for api QueryDevicesByAccount
type QueryDevicesByAccountRequest struct {
	*requests.RpcRequest
	AppKey  requests.Integer `position:"Query" name:"AppKey"`
	Account string           `position:"Query" name:"Account"`
}

// QueryDevicesByAccountResponse is the response struct for api QueryDevicesByAccount
type QueryDevicesByAccountResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	DeviceIds DeviceIdsInQueryDevicesByAccount `json:"DeviceIds" xml:"DeviceIds"`
}

// CreateQueryDevicesByAccountRequest creates a request to invoke QueryDevicesByAccount API
func CreateQueryDevicesByAccountRequest() (request *QueryDevicesByAccountRequest) {
	request = &QueryDevicesByAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryDevicesByAccount", "", "")
	return
}

// CreateQueryDevicesByAccountResponse creates a response to parse from QueryDevicesByAccount response
func CreateQueryDevicesByAccountResponse() (response *QueryDevicesByAccountResponse) {
	response = &QueryDevicesByAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
