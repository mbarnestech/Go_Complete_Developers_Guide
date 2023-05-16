package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()
}
