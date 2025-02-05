/*
Copyright 2023 The KubeService-Stack Authors.

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

package nodestats

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetNodeStats(c Client, ip string, port string, path string) (map[string]float64, error) {
	urlString := "http://" + ip + ":" + port + "/" + strings.Trim(path, "/")
	reqURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(&http.Request{URL: reqURL})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("resp StatusCode is " + strconv.Itoa(resp.StatusCode))
	}
	da, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]float64
	err = json.Unmarshal(da, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
