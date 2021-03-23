/*
Copyright 2020 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configmap

// RtspVisitorConfig is the modbus register configuration.
type RtspVisitorConfig struct {
	Path string `json:"path"`
}

// RtspProtocolConfig is the protocol configuration.
type RtspProtocolConfig struct {
	ProtocolName string     `json:"protocolName,omitempty"`
	ConfigData   ConfigData `json:"configData,omitempty"`
}

// RtspProtocolConfig is the protocol configuration.
type ConfigData struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	IP       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	Path     string `json:"path,omitempty"`
	FullPath string `json:"fullpath,omitempty"`
}

// RtspProtocolCommonConfig is the rtsp protocol configuration.
type RtspProtocolCommonConfig struct {
}

// CustomizedValue is the customized part for modbus protocol.
type CustomizedValue map[string]interface{}

// COMStruct is the serial configuration.
type COMStruct struct {
	SerialPort string `json:"serialPort"`
	BaudRate   int64  `json:"baudRate"`
	DataBits   int64  `json:"dataBits"`
	Parity     string `json:"parity"`
	StopBits   int64  `json:"stopBits"`
}

// TCPStruct is the TCP configuation.
type TCPStruct struct {
	IP   string `json:"ip"`
	Port int64  `json:"port"`
}
