package main

import (
	"big2/cards"
	"fmt"
)

// 牌型 同花順> 鐵支> 葫蘆 > 順子> 三條> 對子> 單張
// 數字大小 2>A>K>Q>J>10>9>8>7>6>5>4>3

// 先比牌型，再比點數，最後花色

// 23456 > 10JQKA > 910JQK > 8910JQ > 78910J > 678910 > 56789 > 45678 > 34567 > A2345
// 1、23456為最大順，以2的花色作為比大小的依據
// 2、A2345為最小順，以5的花色作為比大小的依據
// 3、無JQKA2、QKA23、KA234這種順

func RemoveIndex(s []cards.Card, i int) []cards.Card {
	s[i] = s[len(s)-1]         // Copy last element to index i.
	s[len(s)-1] = cards.Card{} // Erase last element (write zero value).
	s = s[:len(s)-1]           // Truncate slice.
	return s
}

func main() {
	card := cards.Card{}
	// 1. 產生player牌組
	playerHands, _ := card.NewDeck()

	for i, hand := range playerHands {
		fmt.Printf("玩家 %d 的手牌: %v\n", i+1, hand)
	}

	// 判斷牌型
	fmt.Println("--------------cards1-----------------")
	cards1 := []cards.Card{
		playerHands[0][2],
		playerHands[0][3],
		playerHands[0][4],
		playerHands[0][5],
		playerHands[0][6],
	}

	for _, card := range cards1 {
		fmt.Println("card:", cards.Suits[card.Suit], card.Value)
	}
	cardType, highCard, err := card.AnalyzeCards(cards1)
	if err != nil {
		fmt.Println("判斷牌型失敗:", err)
		return
	}
	// fmt.Println("cards1:", cards2)
	fmt.Printf("牌型: %s, 最高牌: %s %d\n", cards.CardTypeStringSlice[cardType], cards.Suits[highCard.Suit], highCard.Value)
	fmt.Println("--------------cards2-----------------")

	cards2 := []cards.Card{
		playerHands[1][10],
		playerHands[1][9],
		playerHands[1][5],
		playerHands[1][3],
		playerHands[1][1],
	}

	for _, card := range cards2 {
		fmt.Println("card:", cards.Suits[card.Suit], card.Value)
	}

	cardType, highCard, err = card.AnalyzeCards(cards2)
	if err != nil {
		fmt.Println("判斷牌型失敗:", err)
		return
	}
	// fmt.Println("cards1:", cards2)
	fmt.Printf("牌型: %s, 最高牌: %s %d\n", cards.CardTypeStringSlice[cardType], cards.Suits[highCard.Suit], highCard.Value)
	fmt.Println("-------------------------------------")
	// 比較牌型
	result, err := card.CompareCard(cards1, cards2)
	if err != nil {
		fmt.Println("比較牌型失敗:", err)
		return
	}

	fmt.Println("比較結果:", result)

	fmt.Println("--------------free test card-----------------")
	cards3 := []cards.Card{
		{Suit: cards.Spades, Value: cards.Six},
		{Suit: cards.Spades, Value: cards.Eight},
		{Suit: cards.Spades, Value: cards.Nine},
		{Suit: cards.Spades, Value: cards.Seven},
		{Suit: cards.Spades, Value: cards.Five},
	}

	cards4 := []cards.Card{
		{Suit: cards.Plum, Value: cards.Five},
		{Suit: cards.Heart, Value: cards.Five},
		{Suit: cards.Heart, Value: cards.Five},
		{Suit: cards.Block, Value: cards.Five},
		{Suit: cards.Plum, Value: cards.Six},
	}

	cards5 := []cards.Card{
		{Suit: cards.Plum, Value: cards.Four},
		{Suit: cards.Heart, Value: cards.Eight},
		{Suit: cards.Plum, Value: cards.Eight},
		{Suit: cards.Spades, Value: cards.Eight},
		{Suit: cards.Spades, Value: cards.Four},
	}

	cards6 := []cards.Card{
		{Suit: cards.Plum, Value: cards.Six},
		{Suit: cards.Heart, Value: cards.Six},
		{Suit: cards.Plum, Value: cards.Seven},
		{Suit: cards.Heart, Value: cards.Seven},
		{Suit: cards.Block, Value: cards.Seven},
	}

	// 比較牌型
	result, err = card.CompareCard(cards3, cards4)
	if err != nil {
		fmt.Println("比較牌型失敗2:", err)
		return
	}

	fmt.Println("比較結果2:", result)

	// 比較牌型
	result, err = card.CompareCard(cards5, cards6)
	if err != nil {
		fmt.Println("比較牌型失敗3:", err)
		return
	}

	fmt.Println("比較結果3:", result)
}
