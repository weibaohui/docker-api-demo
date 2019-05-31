package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type loadResult struct {
	ErrorDetail errorDetail `json:"errorDetail"`
	Stream      string      `json:"stream"`
	Error       string      `json:"error"`
}
type errorDetail struct {
	Message string `json:"message"`
}

func dockerCli() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return cli
}
func Push(tarPath string) {
	cli := dockerCli()

	ctx := context.Background()
	cli.NegotiateAPIVersion(ctx)

	//导入镜像tar包
	fr, err := os.Open(tarPath)
	response, err := cli.ImageLoad(ctx, fr, false)
	bytes, err := ioutil.ReadAll(response.Body)
	result := &loadResult{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		fmt.Println(err.Error())
	}
	if result.Error != "" {
		fmt.Println(result.ErrorDetail.Message)
		return
	}

	//获取导入的镜像名称
	//Loaded image: golang:alpine
	fmt.Println(result.Stream)
	imageName := strings.TrimSpace(result.Stream)
	imageNameOrigin := strings.ReplaceAll(imageName, "Loaded image: ", "")
	imageName = strings.ReplaceAll(imageNameOrigin, "/", "-")

	//镜像全称
	imageNameWithRepo := imageName
	if !strings.HasPrefix(imageName, "weibh/") {
		imageNameWithRepo = "weibh/" + imageName
	}
	//打 tag
	fmt.Println("tag", imageNameOrigin, imageNameWithRepo)

	err = cli.ImageTag(ctx, imageNameOrigin, imageNameWithRepo)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("push", imageNameWithRepo)

	//push 镜像
	authConfig := auth()
	reader, err := cli.ImagePush(ctx, imageNameWithRepo, types.ImagePushOptions{
		RegistryAuth: authConfig,
	})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, reader)

}
func auth() string {
	m := types.AuthConfig{
		Username:      "weibh",
		Password:      "xxxxxx",
		ServerAddress: "docker.io",
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	encodeStr := base64.StdEncoding.EncodeToString(b)
	return encodeStr
}
