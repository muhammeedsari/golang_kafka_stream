package createPerson

import (
	"kafka/personStruct"

	faker "github.com/jaswdr/faker"
)

var f = faker.New()
var p = personStruct.Person{}

func CreatePerson() personStruct.Person {

	p.Name = f.Person().Name()
	p.Gender = f.Person().Gender()
	p.Address = f.Address().Address()
	p.Country = f.Address().Country()
	p.Age = f.IntBetween(0, 100)
	p.Married = f.Bool()
	p.Phone = f.Phone().Number()
	p.CreditCardNumber = f.Payment().CreditCardNumber()
	p.CreditCardExpirationDateString = f.Payment().CreditCardExpirationDateString()
	p.CreditCardType = f.Payment().CreditCardType()
	p.TotalSpendingGold = f.Float64(1000, 1240, 999990)
	p.TotalSessionDuration = f.Float64(1000, 120, 999990)
	p.StartingDate = f.IntBetween(1, 14)
	p.CurrentDate = f.IntBetween(14, 31)
	p.StartingMonth = f.IntBetween(10,12)
	p.CurrentMonth = f.IntBetween(10,12)
	p.TotalClickEvent = f.IntBetween(100, 100000)
	p.TotalSessionCount = f.IntBetween(10, 1500)
	p.TotalScore = f.RandomFloat(10000, 1000, 10000)
	p.TotalGold = f.IntBetween(1000, 95000)
	p.TotalSkillCount = f.IntBetween(1000, 9500)
	p.SwipeLeftCount = f.IntBetween(100, 500)
	p.SwipeRightCount = f.IntBetween(100, 500)
	p.SwipeDownCount = f.IntBetween(100, 500)
	p.SwipeUpCount = f.IntBetween(100, 500)
	p.StartCor = f.RandomFloat(10000, 100, 1000)
	p.FinishCor = f.RandomFloat(10000, 100, 1000)
	p.RemainLife = f.RandomFloat(10000, -20, 100)
	//time.Sleep(8 * time.Second)
	return p
}

// func ConvertGenderToInt(obj personStruct.Person ,objFilter personStruct.PersonFilter) int {

// 	if obj.Gender == "Female" {
// 		return objFilter.Gender = 1
// 	}
// 	return objFilter.Gender = 0
// }