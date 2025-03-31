package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnits := make(map[string]int)

	grossUnits["quarter_of_a_dozen"] = 3
	grossUnits["half_of_a_dozen"] = 6
	grossUnits["dozen"] = 12
	grossUnits["small_gross"] = 120
	grossUnits["gross"] =144
	grossUnits["great_gross"] = 1728
	
	
	return grossUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if value, ok := units[unit]; ok {
		bill[item] += value
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// reduce item
	if value, ok := units[unit]; ok{
		if quanity, ok := bill[item]; ok{
			updatedValue := quanity - value
			
			switch {
			case  updatedValue < 0:
				return false
			case updatedValue == 0:
				delete(bill, item)
				return true
			case updatedValue > 0:
				bill[item] = updatedValue
				return true
			}
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if qty, ok  := bill[item]; ok {
		return qty, true
	} 
	return 0, false
}
