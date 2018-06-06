package cloudphoto

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

// CreateEvent invokes the cloudphoto.CreateEvent API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/createevent.html
func (client *Client) CreateEvent(request *CreateEventRequest) (response *CreateEventResponse, err error) {
	response = CreateCreateEventResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEventWithChan invokes the cloudphoto.CreateEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEventWithChan(request *CreateEventRequest) (<-chan *CreateEventResponse, <-chan error) {
	responseChan := make(chan *CreateEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEvent(request)
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

// CreateEventWithCallback invokes the cloudphoto.CreateEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEventWithCallback(request *CreateEventRequest, callback func(response *CreateEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEventResponse
		var err error
		defer close(result)
		response, err = client.CreateEvent(request)
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

// CreateEventRequest is the request struct for api CreateEvent
type CreateEventRequest struct {
	*requests.RpcRequest
	StoreName        string           `position:"Query" name:"StoreName"`
	LibraryId        string           `position:"Query" name:"LibraryId"`
	Title            string           `position:"Query" name:"Title"`
	StartAt          requests.Integer `position:"Query" name:"StartAt"`
	EndAt            requests.Integer `position:"Query" name:"EndAt"`
	Identity         string           `position:"Query" name:"Identity"`
	WeixinTitle      string           `position:"Query" name:"WeixinTitle"`
	SplashPhotoId    string           `position:"Query" name:"SplashPhotoId"`
	BannerPhotoId    string           `position:"Query" name:"BannerPhotoId"`
	WatermarkPhotoId string           `position:"Query" name:"WatermarkPhotoId"`
	Remark           string           `position:"Query" name:"Remark"`
}

// CreateEventResponse is the response struct for api CreateEvent
type CreateEventResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Event     Event  `json:"Event" xml:"Event"`
}

// CreateCreateEventRequest creates a request to invoke CreateEvent API
func CreateCreateEventRequest() (request *CreateEventRequest) {
	request = &CreateEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateEvent", "cloudphoto", "openAPI")
	return
}

// CreateCreateEventResponse creates a response to parse from CreateEvent response
func CreateCreateEventResponse() (response *CreateEventResponse) {
	response = &CreateEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
