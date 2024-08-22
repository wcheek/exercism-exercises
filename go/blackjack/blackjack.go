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

const (
	Stand   = "S"
	Hit     = "H"
	Split   = "P"
	AutoWin = "W"
)

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pairValue := ParseCard(card1) + ParseCard(card2)
	blackjack := pairValue == 21
	dealerFaceCard := ParseCard(dealerCard) == 10
	dealerAce := dealerCard == "ace"

	highValuePair := pairValue <= 20 && pairValue >= 17
	midValuePair := pairValue <= 16 && pairValue >= 12
	lowValuePair := pairValue <= 11

	switch {
	case card1 == "ace" && card2 == "ace":
		return Split
	case blackjack && !dealerFaceCard && !dealerAce:
		return AutoWin
	case blackjack && (dealerFaceCard || dealerAce):
		return Stand
	case highValuePair:
		return Stand
	case midValuePair:
		if ParseCard(dealerCard) >= 7 {
			return Hit
		} else {
			return Stand
		}
	case lowValuePair:
		return Hit
	}
	panic("Please implement the FirstTurn function")
}
