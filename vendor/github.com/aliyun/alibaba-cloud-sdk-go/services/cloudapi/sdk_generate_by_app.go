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

// SdkGenerateByApp invokes the cloudapi.SdkGenerateByApp API synchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebyapp.html
func (client *Client) SdkGenerateByApp(request *SdkGenerateByAppRequest) (response *SdkGenerateByAppResponse, err error) {
	response = CreateSdkGenerateByAppResponse()
	err = client.DoAction(request, response)
	return
}

// SdkGenerateByAppWithChan invokes the cloudapi.SdkGenerateByApp API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebyapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SdkGenerateByAppWithChan(request *SdkGenerateByAppRequest) (<-chan *SdkGenerateByAppResponse, <-chan error) {
	responseChan := make(chan *SdkGenerateByAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SdkGenerateByApp(request)
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

// SdkGenerateByAppWithCallback invokes the cloudapi.SdkGenerateByApp API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/sdkgeneratebyapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SdkGenerateByAppWithCallback(request *SdkGenerateByAppRequest, callback func(response *SdkGenerateByAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SdkGenerateByAppResponse
		var err error
		defer close(result)
		response, err = client.SdkGenerateByApp(request)
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

// SdkGenerateByAppRequest is the request struct for api SdkGenerateByApp
type SdkGenerateByAppRequest struct {
	*requests.RpcRequest
	AppId    requests.Integer `position:"Query" name:"AppId"`
	Language string           `position:"Query" name:"Language"`
}

// SdkGenerateByAppResponse is the response struct for api SdkGenerateByApp
type SdkGenerateByAppResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DownloadLink string `json:"DownloadLink" xml:"DownloadLink"`
}

// CreateSdkGenerateByAppRequest creates a request to invoke SdkGenerateByApp API
func CreateSdkGenerateByAppRequest() (request *SdkGenerateByAppRequest) {
	request = &SdkGenerateByAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerateByApp", "apigateway", "openAPI")
	return
}

// CreateSdkGenerateByAppResponse creates a response to parse from SdkGenerateByApp response
func CreateSdkGenerateByAppResponse() (response *SdkGenerateByAppResponse) {
	response = &SdkGenerateByAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
