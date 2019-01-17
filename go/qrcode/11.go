package main

import qrcode "github.com/skip2/go-qrcode"
import "fmt"
import "encoding/base64"

func main() {
	// err := qrcode.WriteFile("http://blog.csdn.net/wangshubo1989", qrcode.Medium, 256, "../got/11/qr.png")
	// // qr, err := qrcode.New("http://www.flysnow.org/", qrcode.Medium)
	// if err != nil {
	// 	fmt.Println("write error")
	// }
	byqr, _ := qrcode.Encode("http://www.flysnow.org/", qrcode.Medium, 256)
	fmt.Println(byqr)
	Base64Str := base64.StdEncoding.EncodeToString(byqr)
	fmt.Println(Base64Str)
}
