package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}


// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}
	bill[item] += unitValue // add to existing quantity
	return true
}


// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemExists := bill[item]
	unitValue, unitExists := units[unit]

	// Check if item exists in bill and unit exists in units map
	if !itemExists || !unitExists {
		return false
	}

	newValue := itemValue - unitValue

	if newValue < 0 {
		// Can't have negative quantity
		return false
	} else if newValue == 0 {
		// Remove item from bill
		delete(bill, item)
		return true
	} else {
		// Reduce item quantity
		bill[item] = newValue
		return true
	}
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	}
	return quantity, true
}

