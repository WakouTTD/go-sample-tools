package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

func readLine(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func requestToInstagram(hashTagWord string) (string, error) {

	base, _ := url.Parse("https://www.instagram.com/")
	path := "/explore/tags/" + hashTagWord
	reference, _ := url.Parse(path)
	endpoint := base.ResolveReference(reference).String()
	fmt.Println("リクエスト:" + endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	//fmt.Println(q)
	//fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	// クローズ必要
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	page := string(body)
	//fmt.Println(page)
	return page, nil
}

func main() {

	fmt.Println("開始:" + time.Now().Format(time.RFC3339))

	wordListFile := "./input/word_list.csv"
	lines, err := readLine(wordListFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ハッシュタグページから件数を抜き出すための正規表現
	r := regexp.MustCompile("\"edge_hashtag_to_media\":{\"count\":([0-9]+)")

	outputCsvName := "./output/word_count_list.csv"
	outputCsvFile, err := os.Create(outputCsvName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO Sleep入れるが、ファイル開きっぱなし。全件取得後に書き込む使用だと途中でエラーが出た場合に途中経過が保存されない。
	outputCsvFile, err = os.OpenFile(outputCsvName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outputCsvFile.Close()
	w := bufio.NewWriter(outputCsvFile)

	for _, word := range lines {
		fmt.Println(word)
		page, err := requestToInstagram(word)
		if err != nil {
			//fmt.Println(os.Stderr, err)
			fmt.Println(err)
			os.Exit(1)
		}
		// [0]に'"edge_hashtag_to_media":{"count":13776'　[1]に'13776'のような展開
		//strCount := r.FindStringSubmatch(page)[1]
		matchString := r.FindStringSubmatch(page)

		strCount := "0"
		if len(matchString) == 2 {
			strCount = matchString[1]
		} else {
			fmt.Println(word+": 件数取得エラー:", matchString)
		}

		row := word + "," + strCount + "\n"
		// 出力ファイルに1行書き込み
		_, err = w.WriteString(row)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 1時間に200回までのアクセス制限 (1800*200=360000)
		time.Sleep(1801 * time.Millisecond)
	}

	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("終了:" + time.Now().Format(time.RFC3339))
}
