package qualitycheck

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

// UploadRuleForAnt invokes the qualitycheck.UploadRuleForAnt API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadruleforant.html
func (client *Client) UploadRuleForAnt(request *UploadRuleForAntRequest) (response *UploadRuleForAntResponse, err error) {
	response = CreateUploadRuleForAntResponse()
	err = client.DoAction(request, response)
	return
}

// UploadRuleForAntWithChan invokes the qualitycheck.UploadRuleForAnt API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadruleforant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadRuleForAntWithChan(request *UploadRuleForAntRequest) (<-chan *UploadRuleForAntResponse, <-chan error) {
	responseChan := make(chan *UploadRuleForAntResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadRuleForAnt(request)
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

// UploadRuleForAntWithCallback invokes the qualitycheck.UploadRuleForAnt API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadruleforant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadRuleForAntWithCallback(request *UploadRuleForAntRequest, callback func(response *UploadRuleForAntResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadRuleForAntResponse
		var err error
		defer close(result)
		response, err = client.UploadRuleForAnt(request)
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

// UploadRuleForAntRequest is the request struct for api UploadRuleForAnt
type UploadRuleForAntRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UploadRuleForAntResponse is the response struct for api UploadRuleForAnt
type UploadRuleForAntResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Success   bool                   `json:"Success" xml:"Success"`
	Code      string                 `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Data      DataInUploadRuleForAnt `json:"Data" xml:"Data"`
}

// CreateUploadRuleForAntRequest creates a request to invoke UploadRuleForAnt API
func CreateUploadRuleForAntRequest() (request *UploadRuleForAntRequest) {
	request = &UploadRuleForAntRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadRuleForAnt", "", "")
	return
}

// CreateUploadRuleForAntResponse creates a response to parse from UploadRuleForAnt response
func CreateUploadRuleForAntResponse() (response *UploadRuleForAntResponse) {
	response = &UploadRuleForAntResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}