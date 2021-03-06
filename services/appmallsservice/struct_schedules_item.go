package appmallsservice

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

// SchedulesItem is a nested struct in appmallsservice response
type SchedulesItem struct {
	CinemaId     int64  `json:"CinemaId" xml:"CinemaId"`
	CloseTime    string `json:"CloseTime" xml:"CloseTime"`
	HallName     string `json:"HallName" xml:"HallName"`
	Id           int64  `json:"Id" xml:"Id"`
	IsExpired    bool   `json:"IsExpired" xml:"IsExpired"`
	MaxCanBuy    int64  `json:"MaxCanBuy" xml:"MaxCanBuy"`
	Price        int64  `json:"Price" xml:"Price"`
	ScheduleArea string `json:"ScheduleArea" xml:"ScheduleArea"`
	SectionId    string `json:"SectionId" xml:"SectionId"`
	ServiceFee   int64  `json:"ServiceFee" xml:"ServiceFee"`
	ShowDate     string `json:"ShowDate" xml:"ShowDate"`
	ShowId       int64  `json:"ShowId" xml:"ShowId"`
	ShowTime     string `json:"ShowTime" xml:"ShowTime"`
	ShowVersion  string `json:"ShowVersion" xml:"ShowVersion"`
	HallId       string `json:"HallId" xml:"HallId"`
}
