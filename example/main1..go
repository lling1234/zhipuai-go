package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/joho/godotenv"
//	"log"
//	"net/http"
//	"os"
//	zhipuai "zhipuai-go"
//	"zhipuai-go/consts"
//	"zhipuai-go/httpclient"
//)
//
//type RequestData struct {
//	Role    string `json:"role"`
//	Content string `json:"content"`
//}
//
//func main() {
//	// 从.env文件加载环境变量
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("无法加载.env文件")
//	}
//
//	// 初始化Gin引擎
//	router := gin.Default()
//
//	// 读取环境变量的值
//	apiKey := os.Getenv("API_KEY")
//	consts.ApiKey = apiKey
//
//	// 设置路由处理程序
//	router.POST("/a", handleRequest)
//
//	// 启动服务器
//	log.Fatal(router.Run(":2080"))
//}
//
//func handleRequest(c *gin.Context) {
//	var msg string
//	log.Println("Received message:", msg)
//	var requestData RequestData
//	if err := c.ShouldBindJSON(&requestData); err != nil {
//		c.JSON(400, gin.H{"error": err.Error()})
//		return
//	}
//	log.Println("Received message:", requestData)
//
//	// 解析请求体
//	err := c.BindJSON(&msg)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	Msg := httpclient.Message{
//		Role:    "user",
//		Content: msg,
//	}
//
//	zhipuai.Glmctrl(consts.GLM4, Msg)
//	for line := range zhipuai.GlmRespch {
//		response := line.Choices[0].Delta.Content
//		c.JSON(http.StatusOK, gin.H{"response": response})
//	}
//}
