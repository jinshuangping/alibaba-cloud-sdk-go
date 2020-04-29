package lrg

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

// AddFailMachineToBlackList invokes the lrg.AddFailMachineToBlackList API synchronously
// api document: https://help.aliyun.com/api/lrg/addfailmachinetoblacklist.html
func (client *Client) AddFailMachineToBlackList(request *AddFailMachineToBlackListRequest) (response *AddFailMachineToBlackListResponse, err error) {
	response = CreateAddFailMachineToBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// AddFailMachineToBlackListWithChan invokes the lrg.AddFailMachineToBlackList API asynchronously
// api document: https://help.aliyun.com/api/lrg/addfailmachinetoblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddFailMachineToBlackListWithChan(request *AddFailMachineToBlackListRequest) (<-chan *AddFailMachineToBlackListResponse, <-chan error) {
	responseChan := make(chan *AddFailMachineToBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFailMachineToBlackList(request)
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

// AddFailMachineToBlackListWithCallback invokes the lrg.AddFailMachineToBlackList API asynchronously
// api document: https://help.aliyun.com/api/lrg/addfailmachinetoblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddFailMachineToBlackListWithCallback(request *AddFailMachineToBlackListRequest, callback func(response *AddFailMachineToBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFailMachineToBlackListResponse
		var err error
		defer close(result)
		response, err = client.AddFailMachineToBlackList(request)
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

// AddFailMachineToBlackListRequest is the request struct for api AddFailMachineToBlackList
type AddFailMachineToBlackListRequest struct {
	*requests.RoaRequest
	ProcessId requests.Integer `position:"Body" name:"ProcessId"`
	IpList    string           `position:"Body" name:"IpList"`
}

// AddFailMachineToBlackListResponse is the response struct for api AddFailMachineToBlackList
type AddFailMachineToBlackListResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Status    string                   `json:"Status" xml:"Status"`
	Details   string                   `json:"Details" xml:"Details"`
	Reason    string                   `json:"Reason" xml:"Reason"`
	ErrorCode string                   `json:"ErrorCode" xml:"ErrorCode"`
	Data      []map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateAddFailMachineToBlackListRequest creates a request to invoke AddFailMachineToBlackList API
func CreateAddFailMachineToBlackListRequest() (request *AddFailMachineToBlackListRequest) {
	request = &AddFailMachineToBlackListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "AddFailMachineToBlackList", "/api/v2/addfailmachinetoblacklist", "", "")
	request.Method = requests.POST
	return
}

// CreateAddFailMachineToBlackListResponse creates a response to parse from AddFailMachineToBlackList response
func CreateAddFailMachineToBlackListResponse() (response *AddFailMachineToBlackListResponse) {
	response = &AddFailMachineToBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}