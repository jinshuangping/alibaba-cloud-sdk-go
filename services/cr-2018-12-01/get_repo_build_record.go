package cr_2018_12_01

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

// GetRepoBuildRecord invokes the cr.GetRepoBuildRecord API synchronously
// api document: https://help.aliyun.com/api/cr/getrepobuildrecord.html
func (client *Client) GetRepoBuildRecord(request *GetRepoBuildRecordRequest) (response *GetRepoBuildRecordResponse, err error) {
	response = CreateGetRepoBuildRecordResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoBuildRecordWithChan invokes the cr.GetRepoBuildRecord API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepobuildrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoBuildRecordWithChan(request *GetRepoBuildRecordRequest) (<-chan *GetRepoBuildRecordResponse, <-chan error) {
	responseChan := make(chan *GetRepoBuildRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoBuildRecord(request)
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

// GetRepoBuildRecordWithCallback invokes the cr.GetRepoBuildRecord API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepobuildrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoBuildRecordWithCallback(request *GetRepoBuildRecordRequest, callback func(response *GetRepoBuildRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoBuildRecordResponse
		var err error
		defer close(result)
		response, err = client.GetRepoBuildRecord(request)
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

// GetRepoBuildRecordRequest is the request struct for api GetRepoBuildRecord
type GetRepoBuildRecordRequest struct {
	*requests.RpcRequest
	BuildRecordId string `position:"Query" name:"BuildRecordId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// GetRepoBuildRecordResponse is the response struct for api GetRepoBuildRecord
type GetRepoBuildRecordResponse struct {
	*responses.BaseResponse
	GetRepoBuildRecordIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                        string `json:"Code" xml:"Code"`
	RequestId                   string `json:"RequestId" xml:"RequestId"`
	BuildRecordId               string `json:"BuildRecordId" xml:"BuildRecordId"`
	StartTime                   int64  `json:"StartTime" xml:"StartTime"`
	EndTime                     int64  `json:"EndTime" xml:"EndTime"`
	Status                      string `json:"Status" xml:"Status"`
	Image                       Image  `json:"Image" xml:"Image"`
}

// CreateGetRepoBuildRecordRequest creates a request to invoke GetRepoBuildRecord API
func CreateGetRepoBuildRecordRequest() (request *GetRepoBuildRecordRequest) {
	request = &GetRepoBuildRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetRepoBuildRecord", "cr", "openAPI")
	return
}

// CreateGetRepoBuildRecordResponse creates a response to parse from GetRepoBuildRecord response
func CreateGetRepoBuildRecordResponse() (response *GetRepoBuildRecordResponse) {
	response = &GetRepoBuildRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
