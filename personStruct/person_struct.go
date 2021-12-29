package personStruct

type Person struct {
	Name                           string
	Gender                         string
	Address                        string
	Country                        string
	Age                            int
	Married                        bool
	Phone                          string
	CreditCardNumber               string
	CreditCardExpirationDateString string
	CreditCardType                 string
	TotalSpendingGold              float64
	TotalSessionDuration           float64
	StartingDate                   int
	CurrentDate                    int
	StartingMonth                  int
	CurrentMonth                   int
	TotalClickEvent                int
	TotalSessionCount              int
	TotalScore                     float64
	TotalGold                      int
	TotalSkillCount                int
	SwipeRightCount                int
	SwipeLeftCount                 int
	SwipeDownCount                 int
	SwipeUpCount                   int
	StartCor                       float64
	FinishCor                      float64
	RemainLife                     float64
}

type PersonFilter struct {
	Name                          string
	Gender                        int
	Address                       int
	State                         string
	Age                           int
	Married                       int
	PhoneScore                    int
	CreditCardNumber              int
	CreditCardExpirationDateInt   int
	CreditScore                   float64
	CreditCardTypePoint           int
	CreditCardVisaType            int
	CreditCardMAsterCardType      int
	CreditCardAmericanExpressType int
	CreditCardOtherType           int
	TotalDateDay                  int
	AverageDailySpendingGold      float64
	AverageDailySessionDuration   float64
	DailyClickEvent               float64
	AverageDailySessionCount      float64
	DailyAverageScore             float64
	Swipescore                    float64
	DifferenceCor                 float64
	ClickingScore                 float64
	RemainGold                    float64
	SpendingLife                  float64
	FinallyGamingScore            float64
}
