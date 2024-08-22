package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	retrievedUnit, ok := units[unit]
	if ok {
		bill[item] += retrievedUnit
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	retrievedItemValue, itemOk := bill[item]
	retrievedUnit, unitOk := units[unit]
	if !itemOk || !unitOk {
		return false
	} else {
		quantity := retrievedItemValue - retrievedUnit
		if quantity < 0 {
			return false
		} else if quantity == 0 {
			delete(bill, item)
			return true
		} else {
			bill[item] -= retrievedUnit
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	retrievedItemValue, itemOk := bill[item]
	if !itemOk {
		return 0, false
	} else {
		return retrievedItemValue, true
	}
}
