
package rds

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

func (client *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
response = CreateDescribePriceResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) DescribePriceWithChan(request *DescribePriceRequest) (<-chan *DescribePriceResponse, <-chan error) {
responseChan := make(chan *DescribePriceResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.DescribePrice(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) DescribePriceWithCallback(request *DescribePriceRequest, callback func(response *DescribePriceResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *DescribePriceResponse
var err error
defer close(result)
response, err = client.DescribePrice(request)
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

type DescribePriceRequest struct {
*requests.RpcRequest
                InstanceUsedType  string `position:"Query" name:"InstanceUsedType"`
                UsedTime  string `position:"Query" name:"UsedTime"`
                ClientToken  string `position:"Query" name:"ClientToken"`
                ZoneId  string `position:"Query" name:"ZoneId"`
                Engine  string `position:"Query" name:"Engine"`
                DBInstanceClass  string `position:"Query" name:"DBInstanceClass"`
                DBInstanceStorage  string `position:"Query" name:"DBInstanceStorage"`
                OrderType  string `position:"Query" name:"OrderType"`
                OwnerId  string `position:"Query" name:"OwnerId"`
                CommodityCode  string `position:"Query" name:"CommodityCode"`
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                PayType  string `position:"Query" name:"PayType"`
                Quantity  string `position:"Query" name:"Quantity"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerAccount  string `position:"Query" name:"OwnerAccount"`
                EngineVersion  string `position:"Query" name:"EngineVersion"`
                TimeType  string `position:"Query" name:"TimeType"`
}


type DescribePriceResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            PriceInfo struct {
            Currency     string `json:"Currency" xml:"Currency"`
            OriginalPrice     float64 `json:"OriginalPrice" xml:"OriginalPrice"`
            TradePrice     float64 `json:"TradePrice" xml:"TradePrice"`
            DiscountPrice     float64 `json:"DiscountPrice" xml:"DiscountPrice"`
                RuleIds struct {
                RuleId []    string `json:"RuleId" xml:"RuleId"`
                } `json:"RuleIds" xml:"RuleIds"`
            ActivityInfo struct {
            CheckErrMsg     string `json:"CheckErrMsg" xml:"CheckErrMsg"`
            ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
            Success     string `json:"Success" xml:"Success"`
            }  `json:"ActivityInfo" xml:"ActivityInfo"`
                Coupons struct {
                    Coupon []struct {
            CouponNo     string `json:"CouponNo" xml:"CouponNo"`
            Name     string `json:"Name" xml:"Name"`
            Description     string `json:"Description" xml:"Description"`
            IsSelected     string `json:"IsSelected" xml:"IsSelected"`
                    }   `json:"Coupon" xml:"Coupon"`
                } `json:"Coupons" xml:"Coupons"`
            }  `json:"PriceInfo" xml:"PriceInfo"`
                Rules struct {
                    Rule []struct {
            RuleId     int64 `json:"RuleId" xml:"RuleId"`
            Name     string `json:"Name" xml:"Name"`
            Description     string `json:"Description" xml:"Description"`
                    }   `json:"Rule" xml:"Rule"`
                } `json:"Rules" xml:"Rules"`
}

func CreateDescribePriceRequest() (request *DescribePriceRequest) {
request = &DescribePriceRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Rds", "2014-08-15", "DescribePrice", "", "")
return
}

func CreateDescribePriceResponse() (response *DescribePriceResponse) {
response = &DescribePriceResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}
