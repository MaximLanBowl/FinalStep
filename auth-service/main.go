package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

// func main() {
// 	conn, err := amqp.Dial("ampq://guest:guest@rabbitmq:5672/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"auth_queue",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	r := gin.Default()
// 	r.POST("/login", func(ctx *gin.Context) {
// 		username := ctx.PostForm("username")
// 		password := ctx.PostForm("password")
// 		body := fmt.Sprintf("Username: %s, Password: %s", username, password)
// 		err := ch.Publish(
// 			"",
// 			q.Name,
// 			false,
// 			false,
// 			amqp.Publishing{
// 				ContentType: "text/plain",
// 				Body: []byte(body),
// 			})
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Faild to publish message to RabbitMQ",
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "Logged is Successfully",
// 		})
// 	})
// 	r.Run()
// }



// func main() {
// 	conn, errConnectRabbit := amqp.Dial("ampq://guest:guest@rabbitmq:5672/")
// 	errConnectRabbit = fmt.Errorf("Failed to connect to Rabbitmq")
// 	if errConnectRabbit != nil {
// 		panic(errConnectRabbit)
// 	}
// 	defer conn.Close()

// 	ch, errChannelConnect := conn.Channel()
// 	errChannelConnect = fmt.Errorf("Channel connect failed")
// 	if errChannelConnect != nil {
// 		panic(errChannelConnect)
// 	}
// 	defer ch.Close()

// 	rout := gin.Default()
// 	rout.POST()
// }

func twoSum(nums []int, target int) []int {
    newmap := make(map[int]int)
    target = 0
    nums = newmap
    for i:=0; i < len(nums); i++ {
        if nums[i] == 
    }
    return []int int

}