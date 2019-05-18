package main

import (
	"fmt"
	"log"
	"net/http"

	"bytes"
	"io"
	"io/ioutil"
	"os"
	// "regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(i string) {
	// Request the HTML page.
	res, err := http.Get("http://www.zjdz.cc/products_detail/productId=" + i + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	// doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	band := s.Find("a").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	// })

	if doc.Find(".noinfo").Nodes != nil { //判断是否有产品
		return
	}
	// fmt.Println(doc.Find(".content").Nodes)

	model, _ := doc.Find("#v_id").Attr("value")               //产品编号
	title, _ := doc.Find("#v_nm").Attr("value")               //产品标题
	column, _ := doc.Find("input[name='v_ct']").Attr("value") //所属栏目

	img, _ := doc.Find(".pic-img").Attr("src") //产品缩略图
	imgpathname := upimg("productimg", img)
	content := `<p>`
	doc.Find(".detail img").Each(func(i int, s *goquery.Selection) {
		contentimg, _ := s.Attr("src") //产品详情
		pathname := upimg("productcontentimg", contentimg)
		content = content + `<img alt="" src="` + pathname + `">`
	})
	// fmt.Println(doc.Find("div[name='面包屑'] a").Length)
	doc.Find("div[name='面包屑'] a").Each(func(i int, s *goquery.Selection) {
		fmt.Println("---------------")
		fmt.Println(i)

		fmt.Println(s.Text())
		fmt.Println("---------------")
	})

	fmt.Println(model)
	fmt.Println(title)
	fmt.Println(column)
	fmt.Println(imgpathname)
	fmt.Println(content)
	// fmt.Println(doc.Find(".noinfo").Nodes == nil)

}

func upimg(where, imagPath string) string {
	// imagPath := "http://www.zjdz.cc/imageRepository/de7d5719-d8d7-449b-a6de-558f70aab6bb.jpg"
	//图片正则
	// reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	// name := reg.FindStringSubmatch(imagPath)[0]
	final := strings.Split(imagPath, ".")[1]
	timestamp := time.Now().UnixNano()
	timestamp2 := strconv.FormatInt(timestamp, 10)
	// timestamp11 := time.Now().Unix()
	// tm := time.Unix(timestamp11, 0)
	savepath, pathname, _, _ := Dir(where, timestamp2, final)
	fmt.Println(pathname)
	fmt.Println(savepath)
	//通过http请求获取图片的流文件
	imagPath = "http://www.zjdz.cc" + imagPath
	resp, _ := http.Get(imagPath)
	body, _ := ioutil.ReadAll(resp.Body)
	out, err := os.Create(savepath)
	// err := os.MkdirAll(dirpath, 0777)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(out, bytes.NewReader(body))
	return pathname

}
func Dir(style, timestamp2, final string) (savepath, pathname, path, dirpath string) {
	savepath = "./static/spider/" + style + "/" + timestamp2 + "." + final
	pathname = "/static/spider/" + style + "/" + timestamp2 + "." + final
	path = "/static/spider/" + style
	dirpath = "./static/spider/" + style
	return
}
func main() {
	for i := 75; i < 77; i++ {
		fmt.Println(i)
		ExampleScrape(strconv.Itoa(i))
	}
}
