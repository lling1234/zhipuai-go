package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	zhipuai "zhipuai-go"
	"zhipuai-go/consts"
	"zhipuai-go/httpclient"
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

	// 循环以接收用户输入
	for {
		var input string
		fmt.Print("请输入你的问题，或者按下回车键结束对话：")
		fmt.Scanln(&input)

		// 如果用户没有输入任何内容，那么退出循环
		if strings.TrimSpace(input) == "" {
			break
		}

		// 创建消息结构体
		msg := httpclient.Message{
			Role:    "user",
			Content: input,
		}

		// 调用Zhipuai模型处理用户的问题
		zhipuai.Glmctrl(consts.GLM4, msg)

		// 打印回答
		for line := range zhipuai.GlmRespch {
			fmt.Printf(line.Choices[0].Delta.Content)
		}

		// 输出回答后换行
		fmt.Println()
	}

}
