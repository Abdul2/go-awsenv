package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
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

var a Awstoken

func main() {

	handleRequests()

}

func handleRequests() {
	http.HandleFunc("/", Env)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Env(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile(os.Getenv("myjsonfile")) //file in ~/.aws/cli/cache
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(content, &a)

	os.Setenv("AWS_ACCESS_KEY_ID", a.Credentials.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", a.Credentials.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", a.Credentials.SessionToken)

	//f, err := os.OpenFile("/", os.O_APPEND, 0666)

	file := os.Getenv("myprofilefile")

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	n, err := f.WriteString("\n export AWS_ACCESS_KEY_ID=" + a.Credentials.AccessKeyId)
	m, err := f.WriteString("\n export AWS_SECRET_ACCESS_KEY=" + a.Credentials.SecretAccessKey)
	h, err := f.WriteString("\n export AWS_SESSION_TOKEN=" + a.Credentials.SessionToken)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	f.Close()

	fmt.Fprintln(os.Stdout, n, m, h)

	fmt.Fprintf(w, " values are set \n AccessKeyId %s\n SecretAccessKey %s\n SessionToken %s\n", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN"))

	exec.Command("source", file)
	//replace go shell with new shell with env values - note go run exists

	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

	fmt.Fprintf(w, " values are set \n AccessKeyId %s\n SecretAccessKey %s\n SessionToken %s\n", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN"))

	//json.NewEncoder(w).Encode(a)

}
