package configmap

import (
	"encoding/json"
	"testing"

	"github.com/kubeedge/mappers-go/pkg/common"
	"github.com/kubeedge/mappers-go/pkg/rtsp/globals"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var devices map[string]*globals.RtspDev
	var models map[string]common.DeviceModel
	var protocols map[string]common.Protocol

	devices = make(map[string]*globals.RtspDev)
	models = make(map[string]common.DeviceModel)
	protocols = make(map[string]common.Protocol)

	assert.Nil(t, Parse("./configmap_test.json", devices, models, protocols))
	for _, device := range devices {
		var pc RtspProtocolConfig
		assert.Nil(t, json.Unmarshal([]byte(device.Instance.PProtocol.ProtocolCommonConfig), &pc))
		assert.Equal(t, "192.168.123.211", pc.ConfigData.IP)
	}
}

func TestParseNeg(t *testing.T) {
	var devices map[string]*globals.RtspDev
	var models map[string]common.DeviceModel
	var protocols map[string]common.Protocol

	devices = make(map[string]*globals.RtspDev)
	models = make(map[string]common.DeviceModel)
	protocols = make(map[string]common.Protocol)

	assert.NotNil(t, Parse("./configmap_negtest.json", devices, models, protocols))
}
