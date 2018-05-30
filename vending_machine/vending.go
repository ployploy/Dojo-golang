package vendingmachine

func BuyItem(insert []string, choosedItem string) []string {

	insertedMoney := getInsertedMoney(insert)
	itemPrice := getItemPrice(choosedItem)
	remainingMoney := insertedMoney - itemPrice
	returnCoin := getReturnCoin(remainingMoney)
	if choosedItem == "Coin Return" {
		return returnCoin
	} else if remainingMoney > 0 {
		return returnCoin
	}

	return []string{choosedItem}

}

func getReturnCoin(insertedMoney int) []string {
	returnCoin := []string{}
	remainingMoney := insertedMoney
	if (remainingMoney % 10) >= 0 {
		coins := remainingMoney / 10
		remainingMoney -= 10 * coins
		returnCoin = addReturnCoin(returnCoin, "T", coins)
	}
	if (remainingMoney % 5) >= 0 {
		coins := remainingMoney / 5
		remainingMoney -= 5 * coins
		returnCoin = addReturnCoin(returnCoin, "F", coins)
	}
	if (remainingMoney % 2) >= 0 {
		coins := remainingMoney / 2
		remainingMoney -= 2 * coins
		returnCoin = addReturnCoin(returnCoin, "TW", coins)
	}
	if (remainingMoney % 1) >= 0 {
		coins := remainingMoney / 1
		remainingMoney -= 1 * coins
		returnCoin = addReturnCoin(returnCoin, "O", coins)
	}

	return returnCoin
}

func addReturnCoin(returnCoin []string, symbol string, count int) []string {
	for index := 0; index < count; index++ {
		returnCoin = append(returnCoin, symbol)
	}

	return returnCoin
}

func getInsertedMoney(insert []string) int {
	var insertedMoney int
	for _, value := range insert {
		switch value {
		case "T":
			insertedMoney += 10
		case "F":
			insertedMoney += 5
		case "TW":
			insertedMoney += 2
		case "O":
			insertedMoney += 1
		}
	}

	return insertedMoney
}

func getItemPrice(item string) int {
	var itemPrice int

	switch item {
	case "SD":
		itemPrice += 18
	case "CC":
		itemPrice += 12
	case "DW":
		itemPrice += 7
	}

	return itemPrice
}
