package main

func main() {

	//var card string = "Ace of Spades"
	//card := newCard()
	cards := newDeck()

	cards = append(cards, "Six of Spades")

	cards.shuffle()

	// fmt.Println(cards.toString())
	cards.print()

	// cards.saveToFile("mycards")
	// newDeck := newDeckFromFile("mycards")
	// fmt.Println("Loaded :" + newDeck.toString())

	// hand, remainingDeck := deal(cards, 3)
	// hand.print()
	// remainingDeck.print()

}

func newCard() string {
	return "Five of Diamonds"
}
