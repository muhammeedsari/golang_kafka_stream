package main

import (
	"fmt"
	kafka "kafka/kafka"
	"kafka/personStruct"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	read := make(chan *personStruct.PersonFilter)
	write := make( chan *[]byte)
	
	var obj = personStruct.Person{}
	var objFilter = personStruct.PersonFilter{}
	kafka.CreateTopic()
	fmt.Println("kafka started")
	go kafka.ProduceMesaage(&wg, write)
	go kafka.ReadKafka(&wg, read, &obj, &objFilter)
	write2 := <- write
	read2 := <- read
	fmt.Println("Produce    :   ", write2)
	fmt.Println("consume    :   ", read2)
	wg.Wait()

}
