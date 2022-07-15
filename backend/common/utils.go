// This file contains useful generic functions that could be used
// in any other packages.

package common

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	CUSTOM_NEG_INF  = -999999999999
	DATETIME_FORMAT = "2006-01-02 15:04:05"
	TIME_FORMAT     = "15:04:05"
)

// Checks if an element (key) exists within a sorted list
// Returns if the element is found and index of the key
// in the array if found. If the key is not found, the index
// will be custom -inf for
func BinarySearch(list []int, left int, right int, key int) (bool, int) {
	fmt.Println("Binary Search", left, right, key, list)
	if right < left {
		return false, CUSTOM_NEG_INF
	}

	mid := left + (right-left)/2

	if list[mid] == key {
		return true, mid
	}
	if key < list[mid] {
		return BinarySearch(list, left, mid-1, key)
	} else {
		return BinarySearch(list, mid+1, right, key)
	}
}

func HttpGetWithJWT(url, jwt string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("HttpGetWithJWT ERROR %s", err.Error())
	}

	bearer := "Bearer " + jwt
	req.Header.Set("Authorization", bearer)
	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HttpGetWithJWT ERROR %s", err.Error())
	}

	return response, nil
}

func HttpPostWithJWT(url, jwt string, body string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("HttpPostWithJWT ERROR %s", err.Error())
	}

	bearer := "Bearer " + jwt
	req.Header.Set("x-authorization", bearer)
	req.Header.Set("content-type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HttpPostWithJWT ERROR %s", err.Error())
	}

	return response, nil
}
