/*

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

package v13

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/apache/trafficcontrol/lib/go-tc"
	"github.com/apache/trafficcontrol/lib/go-tc/v13"
)

const (
	API_v13_CacheGroups = "/api/1.3/cachegroups"
)

// Create a CacheGroup
func (to *Session) CreateCacheGroupNullable(cachegroup v13.CacheGroupNullable) (*v13.CacheGroupDetailResponse, ReqInf, error) {

	var remoteAddr net.Addr
	reqBody, err := json.Marshal(cachegroup)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	resp, remoteAddr, err := to.request(http.MethodPost, API_v13_CacheGroups, reqBody)
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()
	var cachegroupResp v13.CacheGroupDetailResponse
	if err = json.NewDecoder(resp.Body).Decode(&cachegroupResp); err != nil {
		return nil, reqInf, err
	}
	return &cachegroupResp, reqInf, nil
}

// Create a CacheGroup
// Deprecated: Use CreateCacheGroupNullable
func (to *Session) CreateCacheGroup(cachegroup v13.CacheGroup) (tc.Alerts, ReqInf, error) {

	var remoteAddr net.Addr
	reqBody, err := json.Marshal(cachegroup)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return tc.Alerts{}, reqInf, err
	}
	resp, remoteAddr, err := to.request(http.MethodPost, API_v13_CacheGroups, reqBody)
	if err != nil {
		return tc.Alerts{}, reqInf, err
	}
	defer resp.Body.Close()
	var alerts tc.Alerts
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	return alerts, reqInf, nil
}

// Update a CacheGroup by ID
func (to *Session) UpdateCacheGroupNullableByID(id int, cachegroup v13.CacheGroupNullable) (*v13.CacheGroupDetailResponse, ReqInf, error) {

	var remoteAddr net.Addr
	reqBody, err := json.Marshal(cachegroup)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	route := fmt.Sprintf("%s/%d", API_v13_CacheGroups, id)
	resp, remoteAddr, err := to.request(http.MethodPut, route, reqBody)
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()
	var cachegroupResp v13.CacheGroupDetailResponse
	if err = json.NewDecoder(resp.Body).Decode(&cachegroupResp); err != nil {
		return nil, reqInf, err
	}
	return &cachegroupResp, reqInf, nil
}

// Update a CacheGroup by ID
// Deprecated: use UpdateCachegroupNullableByID
func (to *Session) UpdateCacheGroupByID(id int, cachegroup v13.CacheGroup) (tc.Alerts, ReqInf, error) {

	var remoteAddr net.Addr
	reqBody, err := json.Marshal(cachegroup)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return tc.Alerts{}, reqInf, err
	}
	route := fmt.Sprintf("%s/%d", API_v13_CacheGroups, id)
	resp, remoteAddr, err := to.request(http.MethodPut, route, reqBody)
	if err != nil {
		return tc.Alerts{}, reqInf, err
	}
	defer resp.Body.Close()
	var alerts tc.Alerts
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	return alerts, reqInf, nil
}

// Returns a list of CacheGroups
func (to *Session) GetCacheGroupsNullable() ([]v13.CacheGroupNullable, ReqInf, error) {
	resp, remoteAddr, err := to.request(http.MethodGet, API_v13_CacheGroups, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsNullableResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, reqInf, err
	}
	return data.Response, reqInf, nil
}

// Returns a list of CacheGroups
// Deprecated: use GetCacheGroupsNullable
func (to *Session) GetCacheGroups() ([]v13.CacheGroup, ReqInf, error) {
	resp, remoteAddr, err := to.request(http.MethodGet, API_v13_CacheGroups, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data.Response, reqInf, nil
}

// GET a CacheGroup by the CacheGroup id
func (to *Session) GetCacheGroupNullableByID(id int) ([]v13.CacheGroupNullable, ReqInf, error) {
	route := fmt.Sprintf("%s/%d", API_v13_CacheGroups, id)
	resp, remoteAddr, err := to.request(http.MethodGet, route, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsNullableResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, reqInf, err
	}

	return data.Response, reqInf, nil
}

// GET a CacheGroup by the CacheGroup id
// Deprecated: use GetCacheGroupNullableByID
func (to *Session) GetCacheGroupByID(id int) ([]v13.CacheGroup, ReqInf, error) {
	route := fmt.Sprintf("%s/%d", API_v13_CacheGroups, id)
	resp, remoteAddr, err := to.request(http.MethodGet, route, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, reqInf, err
	}

	return data.Response, reqInf, nil
}

// GET a CacheGroup by the CacheGroup name
func (to *Session) GetCacheGroupNullableByName(name string) ([]v13.CacheGroupNullable, ReqInf, error) {
	url := fmt.Sprintf("%s?name=%s", API_v13_CacheGroups, name)
	resp, remoteAddr, err := to.request(http.MethodGet, url, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsNullableResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, reqInf, err
	}

	return data.Response, reqInf, nil
}

// GET a CacheGroup by the CacheGroup name
// Deprecated: use GetCachegroupNullableByName
func (to *Session) GetCacheGroupByName(name string) ([]v13.CacheGroup, ReqInf, error) {
	url := fmt.Sprintf("%s?name=%s", API_v13_CacheGroups, name)
	resp, remoteAddr, err := to.request(http.MethodGet, url, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return nil, reqInf, err
	}
	defer resp.Body.Close()

	var data v13.CacheGroupsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, reqInf, err
	}

	return data.Response, reqInf, nil
}

// DELETE a CacheGroup by ID
func (to *Session) DeleteCacheGroupByID(id int) (tc.Alerts, ReqInf, error) {
	route := fmt.Sprintf("%s/%d", API_v13_CacheGroups, id)
	resp, remoteAddr, err := to.request(http.MethodDelete, route, nil)
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	if err != nil {
		return tc.Alerts{}, reqInf, err
	}
	defer resp.Body.Close()
	var alerts tc.Alerts
	if err = json.NewDecoder(resp.Body).Decode(&alerts); err != nil {
		return tc.Alerts{}, reqInf, err
	}
	return alerts, reqInf, nil
}
