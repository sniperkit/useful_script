package main

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/guotie/gogb2312"
	"log"
	//"github.com/guotie/gogb2312"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	sp := NewSpider("News")
	news, err := sp.SpideHTML("http://www.163.com", "div.yaowen_news>div>ul>li")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range news {
		log.Println(s)
	}

	ycnews, err := sp.SpideHTML("http://sports.163.com/yc", "div.topnews>h2")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range ycnews {
		log.Println(s)
	}

	ycnews, err = sp.SpideHTML("http://sports.163.com/yc", "div.topnews>ul>li")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range ycnews {
		log.Println(s)
	}

	sinanews, err := sp.SpideHTML("http://www.sina.com.cn", "ul.news_top>li")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range sinanews {
		log.Println(s)
	}

	sinanews, err = sp.SpideHTML("http://www.sina.com.cn", "div.top_newslist>ul>li")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range sinanews {
		log.Println(s)
	}

	sinanews, err = sp.SpideHTML("http://sports.sina.com.cn/g/premierleague/", "ul.match_news_list>li")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, s := range sinanews {
		log.Println(s)
	}

	hupunews, err := sp.SpideHTML("https://soccer.hupu.com/england/", "ul.england-bignews>li>h2")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, h := range hupunews {
		log.Println(h)
	}

	hupunews, err = sp.SpideHTML("https://soccer.hupu.com/england/", "ul.england-bignews>li>h4")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, h := range hupunews {
		log.Println(h)
	}

	/*
		toutiaonews, err = sp.SpideHTML("http://www.toutiao.com/search/?keyword=英超", "ul.england-bignews>li>h4")
		if err != nil {
			log.Panic(err.Error())
		}
		for _, h := range hupunews {
			log.Println(h)
		}
	*/

}

type Spider struct {
	name   string
	client *http.Client
}

func NewSpider(name string) *Spider {
	return &Spider{
		name: name,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (s *Spider) SpideContent(page, rule string) ([]string, error) {

	resp, err := s.client.Get(page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctype := resp.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(ctype), "gbk") {
		body, err, _, _ = gogb2312.ConvertGB2312(body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	reader := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Cannot create Docuement by Response")
	}

	var res = make([]string, 0)
	doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		res = append(res, sl.Text())
	})
	return res, nil
}

func (s *Spider) SpideAttribute(page, rule, attr string) ([]string, error) {
	resp, err := s.client.Get(page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctype := resp.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(ctype), "gbk") {
		body, err, _, _ = gogb2312.ConvertGB2312(body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	reader := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Cannot create Docuement by Response")
	}

	var res = make([]string, 0)
	doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		attr, ok := sl.Attr(attr)
		if ok {
			res = append(res, attr)
		}
	})
	return res, nil
}

func (s *Spider) SpideHTML(page, rule string) ([]string, error) {
	resp, err := s.client.Get(page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctype := resp.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(ctype), "gbk") {
		body, err, _, _ = gogb2312.ConvertGB2312(body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	reader := bytes.NewReader(body)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("Cannot create Docuement by Response")
	}

	var res = make([]string, 0)
	doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		content, _ := sl.Html()
		res = append(res, content)
	})
	return res, nil
}
