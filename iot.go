/*
 * Copyright 2018 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package formatter_lib

import (
	"encoding/json"
	"net/url"
	"net/http"
	"github.com/SmartEnergyPlatform/iot-device-repository/lib/model"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

func (this *EventTransformer) GetDevice(client HttpClient, deviceId string) (result model.DeviceInstance, err error) {
	resp, err := client.Get(this.IotRepoUrl + "/deviceInstance/" + url.QueryEscape(deviceId))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (this *EventTransformer) GetDeviceType(client HttpClient, id string) (result model.DeviceType, err error) {
	resp, err := client.Get(this.IotRepoUrl + "/deviceType/" + url.QueryEscape(id))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
