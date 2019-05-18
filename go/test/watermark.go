package admin

import (
	"fmt"
	// "gycs/models"
	"errors"

	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func watermark(savepath, final, watermarktype string) error {
	//图片
	img_file, err := os.Open(savepath)
	if err != nil {
		fmt.Println("打开图片出错")
		fmt.Println(err)
		return err
	}
	defer img_file.Close()
	var img image.Image
	if final == "jpg" {
		img, err = jpeg.Decode(img_file)
		if err != nil {
			fmt.Println("把图片解码为结构体时出错")
			fmt.Println(err)
			return err
		}
	} else if final == "png" {
		img, err = png.Decode(img_file)
		if err != nil {
			fmt.Println("把图片解码为结构体时出错")
			fmt.Println(err)
			return err
		}
	} else {
		fmt.Println("图片格式不对")
		err1:=errors.New("图片格式不对")
		return err1
	}
	//水印
	wmb_file, err := os.Open("./static/watermark/watermark.png")
	if err != nil {
		fmt.Println("打开水印图片出错")
		fmt.Println(err)
		return err
	}
	defer wmb_file.Close()
	wmb_img, err := png.Decode(wmb_file)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		fmt.Println(err)
		return err
	}
	//把水印写在顶部中间
	// offset := image.Pt(img.Bounds().Dx()/2-wmb_img.Bounds().Dx()/2, img.Bounds().Dy()/8-wmb_img.Bounds().Dy()/2)
	//把水印写在底部中间
	// offset := image.Pt(img.Bounds().Dx()/2-wmb_img.Bounds().Dx()/2, img.Bounds().Dy()*7/8-wmb_img.Bounds().Dy()/2)

	//把水印写在右下角，并向0坐标偏移10个像素
	// offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	// 把水印写在中间
	offset := image.Pt(img.Bounds().Dx()/2-wmb_img.Bounds().Dx()/2, img.Bounds().Dy()/2-wmb_img.Bounds().Dy()/2)
	switch watermarktype {
	case "1":	// 把水印写在中间
	case "2":
		offset = image.Pt(img.Bounds().Dx()/2-wmb_img.Bounds().Dx()/2, img.Bounds().Dy()/8-wmb_img.Bounds().Dy()/2) //把水印写在顶部中间
	case "3":
		offset = image.Pt(img.Bounds().Dx()/2-wmb_img.Bounds().Dx()/2, img.Bounds().Dy()*7/8-wmb_img.Bounds().Dy()/2) //把水印写在底部中间
	default:
	}
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片,并设置图片质量
	imgw, err := os.Create(savepath)
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()

	// fmt.Println("添加水印图片结束请查看")
	return nil
}
