package ecs

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

// Instance is a nested struct in ecs response
type Instance struct {
	InstanceId              string                               `json:"InstanceId" xml:"InstanceId"`
	InstanceName            string                               `json:"InstanceName" xml:"InstanceName"`
	Description             string                               `json:"Description" xml:"Description"`
	ImageId                 string                               `json:"ImageId" xml:"ImageId"`
	OSName                  string                               `json:"OSName" xml:"OSName"`
	OSType                  string                               `json:"OSType" xml:"OSType"`
	RegionId                string                               `json:"RegionId" xml:"RegionId"`
	ZoneId                  string                               `json:"ZoneId" xml:"ZoneId"`
	ClusterId               string                               `json:"ClusterId" xml:"ClusterId"`
	InstanceType            string                               `json:"InstanceType" xml:"InstanceType"`
	Cpu                     int                                  `json:"Cpu" xml:"Cpu"`
	Memory                  int                                  `json:"Memory" xml:"Memory"`
	HostName                string                               `json:"HostName" xml:"HostName"`
	Status                  string                               `json:"Status" xml:"Status"`
	SerialNumber            string                               `json:"SerialNumber" xml:"SerialNumber"`
	InternetChargeType      string                               `json:"InternetChargeType" xml:"InternetChargeType"`
	InternetMaxBandwidthIn  int                                  `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int                                  `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	VlanId                  string                               `json:"VlanId" xml:"VlanId"`
	CreationTime            string                               `json:"CreationTime" xml:"CreationTime"`
	StartTime               string                               `json:"StartTime" xml:"StartTime"`
	InstanceNetworkType     string                               `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	InstanceChargeType      string                               `json:"InstanceChargeType" xml:"InstanceChargeType"`
	SaleCycle               string                               `json:"SaleCycle" xml:"SaleCycle"`
	ExpiredTime             string                               `json:"ExpiredTime" xml:"ExpiredTime"`
	AutoReleaseTime         string                               `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
	IoOptimized             bool                                 `json:"IoOptimized" xml:"IoOptimized"`
	DeviceAvailable         bool                                 `json:"DeviceAvailable" xml:"DeviceAvailable"`
	InstanceTypeFamily      string                               `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	LocalStorageCapacity    int                                  `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
	LocalStorageAmount      int                                  `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
	GPUAmount               int                                  `json:"GPUAmount" xml:"GPUAmount"`
	GPUSpec                 string                               `json:"GPUSpec" xml:"GPUSpec"`
	SpotStrategy            string                               `json:"SpotStrategy" xml:"SpotStrategy"`
	SpotPriceLimit          float64                              `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	ResourceGroupId         string                               `json:"ResourceGroupId" xml:"ResourceGroupId"`
	KeyPairName             string                               `json:"KeyPairName" xml:"KeyPairName"`
	Recyclable              bool                                 `json:"Recyclable" xml:"Recyclable"`
	HpcClusterId            string                               `json:"HpcClusterId" xml:"HpcClusterId"`
	StoppedMode             string                               `json:"StoppedMode" xml:"StoppedMode"`
	SecurityGroupIds        SecurityGroupIdsInDescribeInstances  `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	PublicIpAddress         PublicIpAddressInDescribeInstances   `json:"PublicIpAddress" xml:"PublicIpAddress"`
	InnerIpAddress          InnerIpAddressInDescribeInstances    `json:"InnerIpAddress" xml:"InnerIpAddress"`
	RdmaIpAddress           RdmaIpAddress                        `json:"RdmaIpAddress" xml:"RdmaIpAddress"`
	VpcAttributes           VpcAttributes                        `json:"VpcAttributes" xml:"VpcAttributes"`
	EipAddress              EipAddress                           `json:"EipAddress" xml:"EipAddress"`
	DedicatedHostAttribute  DedicatedHostAttribute               `json:"DedicatedHostAttribute" xml:"DedicatedHostAttribute"`
	NetworkInterfaces       NetworkInterfacesInDescribeInstances `json:"NetworkInterfaces" xml:"NetworkInterfaces"`
	OperationLocks          OperationLocksInDescribeInstances    `json:"OperationLocks" xml:"OperationLocks"`
	Tags                    TagsInDescribeInstances              `json:"Tags" xml:"Tags"`
}
