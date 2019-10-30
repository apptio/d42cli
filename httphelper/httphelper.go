package httphelper

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// Username: d42 username
// Passwd: d42 password
// BaseURL: Device42 URL to hit for API. eg; https://device42.domain.com/api/1.0/
var (
	Username string
	Password string
	BaseURL  string
)

// DoRequest Extrapolate the baseURL to allow for testing
func DoRequest(requestType string, uri string, target string) string {
	// Get the BaseURL from viper config
	BaseURL = viper.GetString("BaseURL")

	return DoRequestImpl(requestType, BaseURL, uri, target)
}

// DoRequestImpl Run a request against device42
func DoRequestImpl(requestType string, BaseURL string, uri string, target string) string {
	// Build the URL
	requestURL := fmt.Sprintf("%s%s%s", BaseURL, uri, target)
	//fmt.Printf(requestURL)
	// Make an insecure request
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(requestType, requestURL, nil)

	// Prepare for auth
	Username = viper.GetString("Username")
	Password = viper.GetString("Password")
	req.SetBasicAuth(Username, Password)

	// Do the request
	rs, err := client.Do(req)

	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	return bodyString
}
