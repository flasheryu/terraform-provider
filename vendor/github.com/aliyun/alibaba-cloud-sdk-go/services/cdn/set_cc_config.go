package cdn

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

// SetCcConfig invokes the cdn.SetCcConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setccconfig.html
func (client *Client) SetCcConfig(request *SetCcConfigRequest) (response *SetCcConfigResponse, err error) {
	response = CreateSetCcConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetCcConfigWithChan invokes the cdn.SetCcConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setccconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetCcConfigWithChan(request *SetCcConfigRequest) (<-chan *SetCcConfigResponse, <-chan error) {
	responseChan := make(chan *SetCcConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCcConfig(request)
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

// SetCcConfigWithCallback invokes the cdn.SetCcConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setccconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetCcConfigWithCallback(request *SetCcConfigRequest, callback func(response *SetCcConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCcConfigResponse
		var err error
		defer close(result)
		response, err = client.SetCcConfig(request)
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

// SetCcConfigRequest is the request struct for api SetCcConfig
type SetCcConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AllowIps      string           `position:"Query" name:"AllowIps"`
	BlockIps      string           `position:"Query" name:"BlockIps"`
}

// SetCcConfigResponse is the response struct for api SetCcConfig
type SetCcConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCcConfigRequest creates a request to invoke SetCcConfig API
func CreateSetCcConfigRequest() (request *SetCcConfigRequest) {
	request = &SetCcConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetCcConfig", "", "")
	return
}

// CreateSetCcConfigResponse creates a response to parse from SetCcConfig response
func CreateSetCcConfigResponse() (response *SetCcConfigResponse) {
	response = &SetCcConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
