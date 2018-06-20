package db

import (
	"crawler"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"stock"
	"strings"
	"util"
)

type DB struct {
	Stocks map[string]*stock.Stock
}

var shenzhen = "http://quote.eastmoney.com/stocklist.html#sz"
var stockRe = regexp.MustCompile(`(?P<name>[[\p{Han}[:word:]\*]+)\((?P<code>[[:word:]]+)\)`)

var All map[string]string

func GetStockList() (map[string]string, error) {
	stocks, err := crawler.Crawl(shenzhen, "div.quotebody>div#quotesearch>ul>li")
	if err != nil {
		log.Println("Cannot get db from network, get from local")
		return GetStockListFromLocal()
	}
	db := make(map[string]string, len(stocks))

	for _, s := range stocks {
		match := stockRe.FindStringSubmatch(s)
		if len(match) == 3 {
			db[match[2]] = match[1]
			util.AppendToFile("asset/stocklist.txt", []byte(fmt.Sprintf("%s:%s", match[2], match[1])))
		}
	}

	return db, nil
}

func GetStockListFromLocal() (map[string]string, error) {
	content, err := ioutil.ReadFile("asset/stocklist.txt")
	if err != nil {
		return nil, fmt.Errorf("Cannot get stock list from local with: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	db := make(map[string]string, len(lines))
	for _, line := range lines {
		fields := strings.Split(line, ":")
		db[fields[0]] = fields[1]
	}

	return db, nil
}

func init() {
	all, err := GetStockList()
	if err != nil {
		panic(err)
	}

	All = all
}

func New() (*DB, error) {
	db := &DB{
		Stocks: make(map[string]*stock.Stock, len(All)),
	}

	for c, n := range All {
		s, err := stock.New(c, n)
		if err != nil {
			log.Printf("Cannot create new stock %s with %s\n", c, err)
			continue
		}

		db.Stocks[c] = s
	}

	return db, nil
}

func (db *DB) GetStock(code string) (*stock.Stock, error) {
	if s, ok := db.Stocks[code]; ok {
		return s, nil
	}

	return nil, fmt.Errorf("Stock %s does not exist", code)
}
