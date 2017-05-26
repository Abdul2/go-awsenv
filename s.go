package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Awstoken struct to hold content of system generated json file in ~/.aws/cli/cashe/<filename.json>
type Awstoken struct {
	AssumedRoleUser  AssumedRoleUser
	Credentials      Credentials
	ResponseMetadata ResponseMetadata
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId"`
	Arn           string `json:"Arn"`
}

type Credentials struct {
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
	AccessKeyId     string `json:"AccessKeyId"`
}

type ResponseMetadata struct {
	RetryAttempts  string `json:"RetryAttempts"`
	HTTPStatusCode int    `json:"HTTPStatusCode"`
	RequestId      string `json:"RequestId"`
	HTTPHeaders    HTTPHeaders
}

type HTTPHeaders struct {
	xamznrequestid string `json:"x-amzn-requestid"`
	date           string `json:"date"`
	contentlength  string `json:"content-length"`
	contenttype    string `json:"content-type"`
}

func main() {

	var a Awstoken

	content, err := ioutil.ReadFile(os.Getenv("myjsonfile")) //file in ~/.aws/cli/cache
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(content, &a)

	os.Setenv("AWS_ACCESS_KEY_ID", a.Credentials.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", a.Credentials.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", a.Credentials.SessionToken)

	fmt.Fprintf(os.Stdout, " values are set \n AccessKeyId %s\n SecretAccessKey %s\n SessionToken %s\n", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN"))

}
