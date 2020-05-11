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

// AddNodeDnResource invokes the lrg.AddNodeDnResource API synchronously
// api document: https://help.aliyun.com/api/lrg/addnodednresource.html
func (client *Client) AddNodeDnResource(request *AddNodeDnResourceRequest) (response *AddNodeDnResourceResponse, err error) {
	response = CreateAddNodeDnResourceResponse()
	err = client.DoAction(request, response)
	return
}

// AddNodeDnResourceWithChan invokes the lrg.AddNodeDnResource API asynchronously
// api document: https://help.aliyun.com/api/lrg/addnodednresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddNodeDnResourceWithChan(request *AddNodeDnResourceRequest) (<-chan *AddNodeDnResourceResponse, <-chan error) {
	responseChan := make(chan *AddNodeDnResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddNodeDnResource(request)
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

// AddNodeDnResourceWithCallback invokes the lrg.AddNodeDnResource API asynchronously
// api document: https://help.aliyun.com/api/lrg/addnodednresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddNodeDnResourceWithCallback(request *AddNodeDnResourceRequest, callback func(response *AddNodeDnResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddNodeDnResourceResponse
		var err error
		defer close(result)
		response, err = client.AddNodeDnResource(request)
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

// AddNodeDnResourceRequest is the request struct for api AddNodeDnResource
type AddNodeDnResourceRequest struct {
	*requests.RoaRequest
	Id requests.Integer `position:"Path" name:"id"`
}

// AddNodeDnResourceResponse is the response struct for api AddNodeDnResource
type AddNodeDnResourceResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Data    map[string]interface{} `json:"data" xml:"data"`
	Message string                 `json:"message" xml:"message"`
	Success bool                   `json:"success" xml:"success"`
}

// CreateAddNodeDnResourceRequest creates a request to invoke AddNodeDnResource API
func CreateAddNodeDnResourceRequest() (request *AddNodeDnResourceRequest) {
	request = &AddNodeDnResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "AddNodeDnResource", "/api/v2/tianji/node/[id]/dn", "", "")
	request.Method = requests.POST
	return
}

// CreateAddNodeDnResourceResponse creates a response to parse from AddNodeDnResource response
func CreateAddNodeDnResourceResponse() (response *AddNodeDnResourceResponse) {
	response = &AddNodeDnResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}