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
	myCard1 := ParseCard(card1)
	myCard2 := ParseCard(card2)
	dealCard := ParseCard(dealerCard)
	sumCards := myCard1 + myCard2
	var action rune

	switch {
	case sumCards == 21 && dealCard != 11 && dealCard != 10:
		action = 'W'
	case sumCards >= 17 && sumCards <= 20 ||
		sumCards >= 12 && sumCards <= 16 && dealCard < 7:
		action = 'S'
	case sumCards >= 12 && sumCards <= 16 && dealCard >= 7 ||
		sumCards <= 11:
		action = 'H'
	case sumCards == 22:
		action = 'P'
	default:
		action = 'S'
	}

	return string(action)

}
