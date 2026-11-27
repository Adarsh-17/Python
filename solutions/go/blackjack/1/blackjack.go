package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}


// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    value := ParseCard(card1) + ParseCard(card2)
	if ParseCard(card1) == 11 && ParseCard(card2) == 11{
        return "P"
    }else if value == 21 {
        if ParseCard(dealerCard) >= 10{   
    		return "S"
        }
        return "W"
    }else if value<=20 && value>=17 {
        return "S"
    }else if value<=16 && value >=12 {
        if ParseCard(dealerCard) >= 7{
        	return "H"
        }
        return "S"
    }else{
        return "H"
    }
}
