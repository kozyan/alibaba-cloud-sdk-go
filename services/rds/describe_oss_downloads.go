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

func (client *Client) DescribeOssDownloads(request *DescribeOssDownloadsRequest) (response *DescribeOssDownloadsResponse, err error) {
	response = CreateDescribeOssDownloadsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeOssDownloadsWithChan(request *DescribeOssDownloadsRequest) (<-chan *DescribeOssDownloadsResponse, <-chan error) {
	responseChan := make(chan *DescribeOssDownloadsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssDownloads(request)
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

func (client *Client) DescribeOssDownloadsWithCallback(request *DescribeOssDownloadsRequest, callback func(response *DescribeOssDownloadsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssDownloadsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssDownloads(request)
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

type DescribeOssDownloadsRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MigrateTaskId        string           `position:"Query" name:"MigrateTaskId"`
}

type DescribeOssDownloadsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	MigrateIaskId  string `json:"MigrateIaskId" xml:"MigrateIaskId"`
	Items          struct {
		OssDownload []struct {
			FileName   string `json:"FileName" xml:"FileName"`
			CreateTime string `json:"CreateTime" xml:"CreateTime"`
			BakType    string `json:"BakType" xml:"BakType"`
			FileSize   string `json:"FileSize" xml:"FileSize"`
			Status     string `json:"Status" xml:"Status"`
			IsAvail    string `json:"IsAvail" xml:"IsAvail"`
			Desc       string `json:"Desc" xml:"Desc"`
		} `json:"OssDownload" xml:"OssDownload"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeOssDownloadsRequest() (request *DescribeOssDownloadsRequest) {
	request = &DescribeOssDownloadsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeOssDownloads", "", "")
	return
}

func CreateDescribeOssDownloadsResponse() (response *DescribeOssDownloadsResponse) {
	response = &DescribeOssDownloadsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}