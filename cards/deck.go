package cards

import "math/rand"

type Deck struct {
	Cards []Card
}

func (d *Deck) newDeck() [][]Card {
	// 產生牌組
	generateDeck := d.generateDeck()
	// 洗牌
	d.shuffleDeck(generateDeck)
	// 發牌
	return d.dealDeck(generateDeck)
}

// 產生牌組
func (d *Deck) generateDeck() []Card {
	deck := []Card{}
	suits := []Suit{Plum, Block, Heart, Spades}
	values := []int{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, J, Q, K}
	for _, suit := range suits {
		for _, val := range values {
			deck = append(deck, Card{Suit: suit, Value: val})
		}
	}
	d.Cards = deck
	return d.Cards
}

// 洗牌
func (d *Deck) shuffleDeck(deck []Card) {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

// 發牌
func (d *Deck) dealDeck(deck []Card) [][]Card {
	// 計算每個玩家可以獲得的牌數
	numPlayers := 4
	cardsPerPlayer := len(deck) / numPlayers

	// 儲存每個玩家的手牌
	playerHands := make([][]Card, 4)

	// 發牌
	for i := 0; i < numPlayers; i++ {
		// 直接從牌組中取出指定数量的牌，作為玩家的手牌
		playerHands[i] = deck[i*cardsPerPlayer : (i+1)*cardsPerPlayer]
	}

	return playerHands
}
