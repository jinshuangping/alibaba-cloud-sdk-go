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

// GetRepoTagManifest invokes the cr.GetRepoTagManifest API synchronously
// api document: https://help.aliyun.com/api/cr/getrepotagmanifest.html
func (client *Client) GetRepoTagManifest(request *GetRepoTagManifestRequest) (response *GetRepoTagManifestResponse, err error) {
	response = CreateGetRepoTagManifestResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoTagManifestWithChan invokes the cr.GetRepoTagManifest API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagmanifest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagManifestWithChan(request *GetRepoTagManifestRequest) (<-chan *GetRepoTagManifestResponse, <-chan error) {
	responseChan := make(chan *GetRepoTagManifestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoTagManifest(request)
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

// GetRepoTagManifestWithCallback invokes the cr.GetRepoTagManifest API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagmanifest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagManifestWithCallback(request *GetRepoTagManifestRequest, callback func(response *GetRepoTagManifestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoTagManifestResponse
		var err error
		defer close(result)
		response, err = client.GetRepoTagManifest(request)
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

// GetRepoTagManifestRequest is the request struct for api GetRepoTagManifest
type GetRepoTagManifestRequest struct {
	*requests.RpcRequest
	RepoId        string           `position:"Query" name:"RepoId"`
	SchemaVersion requests.Integer `position:"Query" name:"SchemaVersion"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Tag           string           `position:"Query" name:"Tag"`
}

// GetRepoTagManifestResponse is the response struct for api GetRepoTagManifest
type GetRepoTagManifestResponse struct {
	*responses.BaseResponse
	GetRepoTagManifestIsSuccess bool     `json:"IsSuccess" xml:"IsSuccess"`
	Code                        string   `json:"Code" xml:"Code"`
	RequestId                   string   `json:"RequestId" xml:"RequestId"`
	Manifest                    Manifest `json:"Manifest" xml:"Manifest"`
}

// CreateGetRepoTagManifestRequest creates a request to invoke GetRepoTagManifest API
func CreateGetRepoTagManifestRequest() (request *GetRepoTagManifestRequest) {
	request = &GetRepoTagManifestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetRepoTagManifest", "cr", "openAPI")
	return
}

// CreateGetRepoTagManifestResponse creates a response to parse from GetRepoTagManifest response
func CreateGetRepoTagManifestResponse() (response *GetRepoTagManifestResponse) {
	response = &GetRepoTagManifestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
