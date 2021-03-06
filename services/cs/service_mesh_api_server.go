package cs

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

// ServiceMeshApiServer invokes the cs.ServiceMeshApiServer API synchronously
// api document: https://help.aliyun.com/api/cs/servicemeshapiserver.html
func (client *Client) ServiceMeshApiServer(request *ServiceMeshApiServerRequest) (response *ServiceMeshApiServerResponse, err error) {
	response = CreateServiceMeshApiServerResponse()
	err = client.DoAction(request, response)
	return
}

// ServiceMeshApiServerWithChan invokes the cs.ServiceMeshApiServer API asynchronously
// api document: https://help.aliyun.com/api/cs/servicemeshapiserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ServiceMeshApiServerWithChan(request *ServiceMeshApiServerRequest) (<-chan *ServiceMeshApiServerResponse, <-chan error) {
	responseChan := make(chan *ServiceMeshApiServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ServiceMeshApiServer(request)
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

// ServiceMeshApiServerWithCallback invokes the cs.ServiceMeshApiServer API asynchronously
// api document: https://help.aliyun.com/api/cs/servicemeshapiserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ServiceMeshApiServerWithCallback(request *ServiceMeshApiServerRequest, callback func(response *ServiceMeshApiServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ServiceMeshApiServerResponse
		var err error
		defer close(result)
		response, err = client.ServiceMeshApiServer(request)
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

// ServiceMeshApiServerRequest is the request struct for api ServiceMeshApiServer
type ServiceMeshApiServerRequest struct {
	*requests.RoaRequest
	ServiceMeshId string `position:"Path" name:"ServiceMeshId"`
}

// ServiceMeshApiServerResponse is the response struct for api ServiceMeshApiServer
type ServiceMeshApiServerResponse struct {
	*responses.BaseResponse
}

// CreateServiceMeshApiServerRequest creates a request to invoke ServiceMeshApiServer API
func CreateServiceMeshApiServerRequest() (request *ServiceMeshApiServerRequest) {
	request = &ServiceMeshApiServerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ServiceMeshApiServer", "/servicemesh/[ServiceMeshId]/api_proxy", "csk", "openAPI")
	request.Method = requests.POST
	return
}

// CreateServiceMeshApiServerResponse creates a response to parse from ServiceMeshApiServer response
func CreateServiceMeshApiServerResponse() (response *ServiceMeshApiServerResponse) {
	response = &ServiceMeshApiServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
