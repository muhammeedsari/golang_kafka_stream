package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka/personStruct"
	"strconv"
	"strings"
	s "strings"
	"sync"

	"github.com/segmentio/kafka-go"
)

func ReadKafka(wg *sync.WaitGroup, c chan *personStruct.PersonFilter,
	obj *personStruct.Person, objFilter *personStruct.PersonFilter) {
	//defer wg.Done()
	conf := kafka.ReaderConfig{
		Brokers:  []string{"64.227.7.141:9092"},
		Topic:    "my-topic",
		GroupID:  "g1",
		MaxBytes: 1000,
	}

	reader := kafka.NewReader(conf)

	for {
		wg.Add(1)
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error occured", err)
			continue
		}
		//fmt.Println("message is : ", string(m.Value))
		//d <- string(m.Value)
		if err := json.Unmarshal(m.Value, obj); err != nil {
			panic(err)
		}
		//fmt.Println(obj)

		objFilter.Name = obj.Name

		if obj.Gender == "Female" {
			objFilter.Gender = 1
		} else {
			objFilter.Gender = 2
		}

		arr := s.Split(obj.Address, " ")
		objFilter.Address, _ = strconv.Atoi(arr[0])

		objFilter.State = arr[len(arr)-2]

		objFilter.Age = obj.Age

		if obj.Married == true {
			objFilter.Married = 1
		} else {
			objFilter.Married = 2
		}

		phonePoint := obj.Phone
		phonePoint = phonePoint[len(phonePoint)-3:]
		objFilter.PhoneScore, _ = strconv.Atoi(phonePoint)

		creditNumberStringArray := strings.Split(obj.CreditCardNumber, "")
		var creditCardIntArray = []int{}
		for _, i := range creditNumberStringArray {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			creditCardIntArray = append(creditCardIntArray, j)
		}
		toplam := 0
		for _, b := range creditCardIntArray {
			toplam = toplam + b
		}
		objFilter.CreditCardNumber = toplam

		creditCardED := strings.Split(obj.CreditCardExpirationDateString, "/")
		cced, _ := strconv.Atoi(creditCardED[1])
		objFilter.CreditCardExpirationDateInt = cced

		if obj.CreditCardType == "Visa" {
			objFilter.CreditCardTypePoint = 30
			objFilter.CreditCardVisaType = 1
			//return
		}
		if obj.CreditCardType == "MasterCard" {
			objFilter.CreditCardTypePoint = 25
			objFilter.CreditCardMAsterCardType = 1
			//return
		}
		if obj.CreditCardType == "American Express" {
			objFilter.CreditCardTypePoint = 20
			objFilter.CreditCardAmericanExpressType = 1
			//return
		} else {
			objFilter.CreditCardTypePoint = 10
			objFilter.CreditCardOtherType = 1
		}

		objFilter.CreditScore = (float64(objFilter.CreditCardTypePoint) + float64(objFilter.Age) + float64(objFilter.Married) + float64(objFilter.CreditCardExpirationDateInt) + float64(objFilter.Address) - float64(objFilter.PhoneScore)) / float64(objFilter.CreditCardNumber)

		objFilter.AverageDailySessionDuration = obj.TotalSessionDuration / float64(objFilter.TotalDateDay)

		objFilter.TotalDateDay = (obj.CurrentMonth-obj.StartingMonth)*30 + (obj.CurrentDate - obj.StartingDate)

		objFilter.AverageDailySpendingGold = float64(obj.TotalSpendingGold) / float64(objFilter.TotalDateDay)

		objFilter.DailyClickEvent = float64(obj.TotalClickEvent) / float64(objFilter.TotalDateDay)

		objFilter.AverageDailySessionCount = float64(obj.TotalSessionCount) / float64(objFilter.TotalDateDay)

		objFilter.DailyAverageScore = float64(obj.TotalScore) + float64(obj.TotalSkillCount)/float64(objFilter.TotalDateDay)

		objFilter.Swipescore = float64(obj.SwipeRightCount+obj.SwipeLeftCount) / float64(obj.SwipeDownCount+obj.SwipeUpCount)

		objFilter.DifferenceCor = obj.StartCor - obj.FinishCor

		objFilter.ClickingScore = objFilter.Swipescore/objFilter.DifferenceCor + objFilter.DailyClickEvent + float64(objFilter.Gender)

		objFilter.RemainGold = float64(obj.TotalGold) - float64(obj.TotalSpendingGold)

		objFilter.SpendingLife = 100 - obj.RemainLife

		objFilter.FinallyGamingScore = ((objFilter.SpendingLife+objFilter.RemainGold)/100 + objFilter.Swipescore + (5 * objFilter.DailyAverageScore) + (2 * objFilter.CreditScore)) / 10

		//fmt.Println(objFilter)

		c <- objFilter

		empJSON, _ := json.Marshal(objFilter)

		fmt.Printf("Marshal consume     %s\n: ", string(empJSON))
		wg.Done()
	}

}
