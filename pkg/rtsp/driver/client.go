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

package driver

import (
	"sync"
	"time"

	"k8s.io/klog/v2"

	"github.com/majoyz/gortsplib"
	"github.com/majoyz/gortsplib/pkg/base"
)

// Rtsp is the configurations of modbus TCP.
type Rtsp struct {
	User     string
	Password string
	IP       string
	Port     int
	Path     string
	Fullpath string
	Timeout  time.Duration
}

// RtspClient is the structure for modbus client.
type RtspClient struct {
	Client  Rtsp
	Handler interface{}
	Config  interface{}

	mu sync.Mutex
}

func NewClient(rtsp Rtsp) (*RtspClient, error) {
	rc := RtspClient{
		Client: rtsp,
	}
	return &rc, nil
}

/*
* In modbus RTU mode, devices could connect to one serial port on RS485. However,
* the serial port doesn't support paralleled visit, and for one tcp device, it also doesn't support
* paralleled visit, so we expect one client for one port.
 */
var clients map[string]*RtspClient

// GetStatus get device status.
// Now we could only get the connection status.
func (c *RtspClient) GetStatus() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, err := base.ParseURL(c.Client.Fullpath)
	// klog.Info("Getting status: ", c.Client.Fullpath)
	if err != nil {
		return DEVSTDISCONN
	}

	conn, err := gortsplib.Dial(u.Scheme, u.Host)
	if err != nil {
		return DEVSTDISCONN
	}
	defer conn.Close()

	_, err = conn.Options(u)
	if err != nil {
		return DEVSTDISCONN
	}

	// tracks, _, err := conn.Describe(u)
	_, _, err = conn.Describe(u)
	if err != nil {
		return DEVSTDISCONN
	}

	// klog.Info("available tracks: %v\n", tracks)
	if err == nil {
		return DEVSTOK
	}
	return DEVSTDISCONN
}

// func (c Rtsp) Connect() (err error) {
// 	_, err = newProgram(os.Args[1:])
// 	if err != nil {
// 		// log.Fatal("ERR: ", err)
// 		klog.Error(err)
// 	}

// 	// select {}
// 	return err
// }

// Get get register.
func (c *RtspClient) Get() (results []byte, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	results = []byte("OFF")
	klog.V(2).Info("Get result: ", results)
	return results, nil
}

// // Set set register.
// // func (c *RtspClient) Set(registerType string, addr uint16, value uint16) (results []byte, err error) {
// // 	c.mu.Lock()
// // 	defer c.mu.Unlock()

// // 	klog.V(1).Info("Set:", registerType, addr, value)

// // 	switch registerType {
// // 	case "CoilRegister":
// // 		var valueSet uint16
// // 		switch value {
// // 		case 0:
// // 			valueSet = 0x0000
// // 		case 1:
// // 			valueSet = 0xFF00
// // 		default:
// // 			return nil, errors.New("Wrong value")
// // 		}
// // 		results, err = c.Client.WriteSingleCoil(addr, valueSet)
// // 	case "HoldingRegister":
// // 		results, err = c.Client.WriteSingleRegister(addr, value)
// // 	default:
// // 		return nil, errors.New("Bad register type")
// // 	}
// // 	klog.V(1).Info("Set result:", err, results)
// // 	return results, err
// // }

// // parity convert into the format that modbus drvier requires.
// func parity(ori string) string {
// 	var p string
// 	switch ori {
// 	case "even":
// 		p = "E"
// 	case "odd":
// 		p = "O"
// 	default:
// 		p = "N"
// 	}
// 	return p
// }
