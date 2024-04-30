package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	zhipuai "zhipuai-go"
	"zhipuai-go/consts"
	"zhipuai-go/httpclient"

	"fmt"
)

func main() {
	// 从.env文件加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("无法加载.env文件")
	}

	// 读取环境变量的值
	apiKey := os.Getenv("API_KEY")
	consts.ApiKey = apiKey

	var msg string

	fmt.Scanln(&msg)

	Msg := httpclient.Message{
		Role:    "user",
		Content: msg,
	}

	/*
		//CogView 模型
		zhipuai.Glmctrl(consts.CogView,Msg)
		for line := range zhipuai.CogRespch {
			fmt.Println(line.Data[0].Url)
		}
	*/
	// Glm-4
	zhipuai.Glmctrl(consts.GLM4, Msg)

	for line := range zhipuai.GlmRespch {
		fmt.Printf(line.Choices[0].Delta.Content)
	}
	fmt.Println()

}
