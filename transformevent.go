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
	"errors"
	"log"
	"strings"
	"github.com/SmartEnergyPlatform/iot-device-repository/lib/model"
)

type EventTransformer struct {
	Service      model.Service
	deviceConfig []model.ConfigField
	IotRepoUrl	 string
}

func NewTransformerFromRouting(iotRepo string, httpclient HttpClient, routing string) (result EventTransformer, err error) {
	parts := strings.Split(routing, ".")
	if len(parts) < 2 {
		err = errors.New("ERROR: routing is to short to identify service and device")
		log.Println("ERROR: routing is to short to identify service and device")
	}
	serviceid := parts[1]
	deviceid := parts[0]
	return NewTransformer(iotRepo, httpclient, deviceid, serviceid)
}

func NewTransformer(iotRepo string, httpclient HttpClient, deviceid string, serviceid string) (transformer EventTransformer, err error) {
	transformer.IotRepoUrl = iotRepo
	device, err := transformer.GetDevice(httpclient, deviceid)
	if err != nil {
		return transformer, err
	}
	dt, err := transformer.GetDeviceType(httpclient, device.DeviceType)
	if err != nil {
		return transformer, err
	}
	service, err := getService(dt.Services, serviceid)
	if err != nil {
		return transformer, err
	}
	transformer.Service = service
	transformer.deviceConfig = device.Config
	return
}

func getService(services []model.Service, serviceId string) (service model.Service, err error) {
	for _, service := range services {
		if service.Id == serviceId {
			return service, err
		}
	}
	return service, errors.New("unknown service")
}

func (this *EventTransformer) Transform(event EventMsg) (result string, err error) {
	formated, err := this.toInternOutput(event)
	if err != nil {
		return result, err
	}
	byteResult, err := json.Marshal(formated)
	if err != nil {
		return result, err
	}
	result = string(byteResult)
	return
}

func (this *EventTransformer) toInternOutput(eventMsg EventMsg) (result FormatedOutput, err error) {
	result = map[string]interface{}{}
	for _, output := range eventMsg {
		for _, serviceOutput := range this.Service.Output {
			if serviceOutput.MsgSegment.Name == output.Name {
				parsedOutput, err := ParseFormat(serviceOutput.Type, serviceOutput.Format, output.Value, serviceOutput.AdditionalFormatinfo)
				if err != nil {
					log.Println("error on parsing")
					return result, err
				}
				outputInterface, err := FormatToJsonStruct([]model.ConfigField{}, parsedOutput)
				if err != nil {
					return result, err
				}
				parsedOutput.Name = serviceOutput.Name
				result[serviceOutput.Name] = outputInterface
			}
		}
	}
	return
}
