package vpc

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

// DescribeSnatTableEntries invokes the vpc.DescribeSnatTableEntries API synchronously
// api document: https://help.aliyun.com/api/vpc/describesnattableentries.html
func (client *Client) DescribeSnatTableEntries(request *DescribeSnatTableEntriesRequest) (response *DescribeSnatTableEntriesResponse, err error) {
	response = CreateDescribeSnatTableEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSnatTableEntriesWithChan invokes the vpc.DescribeSnatTableEntries API asynchronously
// api document: https://help.aliyun.com/api/vpc/describesnattableentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnatTableEntriesWithChan(request *DescribeSnatTableEntriesRequest) (<-chan *DescribeSnatTableEntriesResponse, <-chan error) {
	responseChan := make(chan *DescribeSnatTableEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnatTableEntries(request)
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

// DescribeSnatTableEntriesWithCallback invokes the vpc.DescribeSnatTableEntries API asynchronously
// api document: https://help.aliyun.com/api/vpc/describesnattableentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnatTableEntriesWithCallback(request *DescribeSnatTableEntriesRequest, callback func(response *DescribeSnatTableEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnatTableEntriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnatTableEntries(request)
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

// DescribeSnatTableEntriesRequest is the request struct for api DescribeSnatTableEntries
type DescribeSnatTableEntriesRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SnatTableId          string           `position:"Query" name:"SnatTableId"`
	SnatEntryId          string           `position:"Query" name:"SnatEntryId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeSnatTableEntriesResponse is the response struct for api DescribeSnatTableEntries
type DescribeSnatTableEntriesResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int              `json:"PageNumber" xml:"PageNumber"`
	PageSize         int              `json:"PageSize" xml:"PageSize"`
	SnatTableEntries SnatTableEntries `json:"SnatTableEntries" xml:"SnatTableEntries"`
}

// CreateDescribeSnatTableEntriesRequest creates a request to invoke DescribeSnatTableEntries API
func CreateDescribeSnatTableEntriesRequest() (request *DescribeSnatTableEntriesRequest) {
	request = &DescribeSnatTableEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSnatTableEntries", "vpc", "openAPI")
	return
}

// CreateDescribeSnatTableEntriesResponse creates a response to parse from DescribeSnatTableEntries response
func CreateDescribeSnatTableEntriesResponse() (response *DescribeSnatTableEntriesResponse) {
	response = &DescribeSnatTableEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
