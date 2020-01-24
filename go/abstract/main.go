package main

import (
	"fmt"
	"strconv"
)

// RecordInterface :ファンクション定義
type RecordInterface interface {
	scraping()
}

// Tables :Baseのスライス
type Tables struct {
	Records []RecordInterface
}

// DogTable desu
type DogTable struct {
	id    int
	title string
}

func (d *DogTable) scraping() {
	fmt.Println(strconv.Itoa(d.id) + d.title + "ワンワンじゃ")
}

// CatTable dayo
type CatTable struct {
	id    int
	title string
}

func (c *CatTable) scraping() {
	fmt.Println(c.title + "にゃーにゃーじゃ")
}

func main() {
	// ベースとなる構造体を定義
	table := Tables{}

	// 犬の構造体と猫の構造体を生成
	dogRecord := DogTable{id: 100, title: "ジョン"}
	catRecord := CatTable{id: 101, title: "マイケル"}

	// それぞれ追加
	table.Records = append(table.Records, &dogRecord)
	table.Records = append(table.Records, &catRecord)

	fmt.Println(table.Records[0].(*DogTable).id)
	fmt.Println(table.Records[1].(*CatTable).id)

	// dogTableとcatTableのscrapingは挙動が違っても呼べる
	table.Records[0].scraping()
	table.Records[1].scraping()

}
