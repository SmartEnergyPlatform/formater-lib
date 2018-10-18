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
	"testing"
	"github.com/SmartEnergyPlatform/iot-device-repository/lib/model"
	"encoding/json"
	"reflect"
)

func TestJSON(t *testing.T) {

	vtStr := `{  
                  "id":"iot#b81351e3-28f1-4dc9-8471-61f7d930c305",
                  "name":"Generic-Reading",
                  "description":"Generic fields for measurements.",
                  "base_type":"http://www.sepl.wifa.uni-leipzig.de/ontlogies/device-repo#structure",
                  "fields":[  
                     {  
                        "id":"iot#1e47e2cb-4c36-4bcf-ab8f-73a5971da4b5",
                        "name":"time",
                        "type":{  
                           "id":"iot#c8c36810-c8e0-403e-b00f-187414a84ccd",
                           "name":"text",
                           "description":"text",
                           "base_type":"http://www.w3.org/2001/XMLSchema#string",
                           "fields":null,
                           "literal":""
                        }
                     },
                     {  
                        "id":"iot#4bbc743f-7d9f-4ac6-b26a-2d4f902c7baf",
                        "name":"value",
                        "type":{  
                           "id":"iot#cb0dc896-6d89-4e0c-ac59-33eceed512b0",
                           "name":"float",
                           "description":"float",
                           "base_type":"http://www.w3.org/2001/XMLSchema#decimal",
                           "fields":null,
                           "literal":""
                        }
                     },
                     {  
                        "id":"iot#76ff9d33-7b9d-414a-9e70-507c9cc43a16",
                        "name":"unit",
                        "type":{  
                           "id":"iot#c8c36810-c8e0-403e-b00f-187414a84ccd",
                           "name":"text",
                           "description":"text",
                           "base_type":"http://www.w3.org/2001/XMLSchema#string",
                           "fields":null,
                           "literal":""
                        }
                     }
                  ],
                  "literal":""
               }`

	vt := model.ValueType{}
	err := json.Unmarshal([]byte(vtStr), &vt)
	if err != nil {
		t.Fatal(err)
	}

	msg1 := `{
    	"time": "timeVal",
    	"unit": "unitVal",
    	"value": 42
	}`

	msg2 := `{
		"foo": "bar",
		"batz": [2, 3],
		"bla": {"name": "test", "something": 42},
		"list": [{"element":1}, {"element":2}],
		"n": null,
		"b": true,
		"time": "timeValue"
	}`

	result := jsonFormatHelper(t, msg1, vt)
	expected1 := `{
    "time": "timeVal",
    "unit": "unitVal",
    "value": 42
}`
	jsonEqualHelper(t,result, expected1)


	result = jsonFormatHelper(t, msg2, vt)
	expected2 := `{
		    "time": "timeValue"
		}`

	jsonEqualHelper(t,result, expected2)


}


func TestXML(t *testing.T) {

	vtStr := `{  
                  "id":"iot#b81351e3-28f1-4dc9-8471-61f7d930c305",
                  "name":"Generic-Reading",
                  "description":"Generic fields for measurements.",
                  "base_type":"http://www.sepl.wifa.uni-leipzig.de/ontlogies/device-repo#structure",
                  "fields":[  
                     {  
                        "id":"iot#1e47e2cb-4c36-4bcf-ab8f-73a5971da4b5",
                        "name":"time",
                        "type":{  
                           "id":"iot#c8c36810-c8e0-403e-b00f-187414a84ccd",
                           "name":"text",
                           "description":"text",
                           "base_type":"http://www.w3.org/2001/XMLSchema#string",
                           "fields":null,
                           "literal":""
                        }
                     },
                     {  
                        "id":"iot#4bbc743f-7d9f-4ac6-b26a-2d4f902c7baf",
                        "name":"value",
                        "type":{  
                           "id":"iot#cb0dc896-6d89-4e0c-ac59-33eceed512b0",
                           "name":"float",
                           "description":"float",
                           "base_type":"http://www.w3.org/2001/XMLSchema#decimal",
                           "fields":null,
                           "literal":""
                        }
                     },
                     {  
                        "id":"iot#76ff9d33-7b9d-414a-9e70-507c9cc43a16",
                        "name":"unit",
                        "type":{  
                           "id":"iot#c8c36810-c8e0-403e-b00f-187414a84ccd",
                           "name":"text",
                           "description":"text",
                           "base_type":"http://www.w3.org/2001/XMLSchema#string",
                           "fields":null,
                           "literal":""
                        }
                     }
                  ],
                  "literal":""
               }`

	vt := model.ValueType{}
	err := json.Unmarshal([]byte(vtStr), &vt)
	if err != nil {
		t.Fatal(err)
	}

	msg1 := `<root>
		<time>timeVal</time> 
    	<unit>unitVal</unit>
    	<value>42</value>
	</root>`

	msg2 := `<root>
		<foo>bar</foo>
		<bla name="test"><something>42</something></bla>
		<n>null</n>
		<b>true</b>
		<time>timeValue</time> 
	</root>`

	result := xmlFormatHelper(t, msg1, vt)
	expected1 := `{
    "time": "timeVal",
    "unit": "unitVal",
    "value": 42
}`
	jsonEqualHelper(t,result, expected1)


	result = xmlFormatHelper(t, msg2, vt)
	expected2 := `{
		    "time": "timeValue"
		}`

	jsonEqualHelper(t,result, expected2)


}


func jsonEqualHelper(t *testing.T, a string, b string){
	t.Helper()
	var vala interface{}
	var valb interface{}
	err := json.Unmarshal([]byte(a), &vala)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal([]byte(b), &valb)
	if err != nil {
		t.Fatal(err)
	}
	equal := reflect.DeepEqual(vala, valb)
	if !equal {
		t.Fatal(a, b)
	}
}

func jsonFormatHelper(t *testing.T, value string, vt model.ValueType)(result string){
	t.Helper()
	output, err := ParseFromJson(vt, value)
	if err != nil {
		t.Fatal(err)
	}
	result, err = FormatToJson([]model.ConfigField{}, output)
	if err != nil {
		t.Fatal(err)
	}
	return
}


func xmlFormatHelper(t *testing.T, value string, vt model.ValueType)(result string){
	t.Helper()
	output, err := ParseFromXml(vt, value, []model.AdditionalFormatInfo{})
	if err != nil {
		t.Fatal(err)
	}
	result, err = FormatToJson([]model.ConfigField{}, output)
	if err != nil {
		t.Fatal(err)
	}
	return
}
