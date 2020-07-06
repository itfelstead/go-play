package main

func main() {

	//var card string = "Ace of Spades"
	//card := newCard()
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	hand, remainingDeck := deal(cards, 3)

	hand.print()
	remainingDeck.print()

}

func newCard() string {
	return "Five of Diamonds"
}
