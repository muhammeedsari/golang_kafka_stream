package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"kafka/createPerson"
	"net"
	"strconv"
	"sync"

	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("64.227.7.141:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CreateTopic() {
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", "64.227.7.141:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func ConvertStructToByteArray(v interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(v)
	return reqBodyBytes.Bytes()
}


func ProduceMesaage(wg *sync.WaitGroup, c chan *[]byte) {
	//defer wg.Done()
	writer := newKafkaWriter()
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for {
		wg.Add(1)
		person := createPerson.CreatePerson()

		ByteMessage := ConvertStructToByteArray(person)

		msg := kafka.Message{
			Key:   []byte("keyA"),
			Value: []byte(ByteMessage),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced   : ", string(msg.Value))
			c <- &msg.Value
		}
		wg.Done()
		//time.Sleep(1 * time.Second)
	}

}
