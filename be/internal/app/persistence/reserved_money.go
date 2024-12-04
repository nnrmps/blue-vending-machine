package persistence

type ReservedMoney struct {
	Coins1   int64
	Coins5   int64
	Coins10  int64
	Bank20   int64
	Bank50   int64
	Bank100  int64
	Bank500  int64
	Bank1000 int64
}

func (ReservedMoney) TableName() string {
	return "reserved_money"
}
