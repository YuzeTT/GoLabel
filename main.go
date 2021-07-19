package main

import "github.com/fogleman/gg"

type image_size struct {
	width  int
	height int
}

func main() {
	// // 创建一个code128编码的 BarcodeIntCS
	// cs, _ := code128.Encode("A1001")
	// // 创建一个要输出数据的文件
	// file, _ := os.Create("qr3.png")
	// defer file.Close()

	// // 设置图片像素大小
	// qrCode, _ := barcode.Scale(cs, 350, 70)
	// // 将code128的条形码编码为png图片
	// png.Encode(file, qrCode)
	image_size := image_size{width: 480, height: 200}
	// dc := gg.NewContext(image_size.width, image_size.height)
	im, err := gg.LoadImage("./标签.png")
	if err != nil {
		panic(err)
	}
	image_size.width = im.Bounds().Size().X
	image_size.height = im.Bounds().Size().Y
	dc := gg.NewContext(image_size.width, image_size.height)
	dc.DrawImage(im, 0, 0)

	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("font/优设标题黑.ttf", 185); err != nil {
		panic(err)
	}
	dc.DrawString("101001", 88, float64(image_size.height)-83)

	dc.SavePNG("output/output.png")
}
