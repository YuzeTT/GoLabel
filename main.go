package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/gg"
	"github.com/gookit/color"
)

type image_size struct {
	width  int
	height int
}

type model struct {
	cursor int
	choice chan string
}

var choices = []string{"创建数字标签", "查询标签"}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			close(m.choice) // If we're quitting just close the channel.
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice <- choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("选择操作\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("> ")
		} else {
			s.WriteString("  ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(向下：j/↓ 向上：k/↑ 确定：enter 退出：ESC/q)\n")

	return s.String()
}

func createLabel(text string) {
	image_size := image_size{width: 480, height: 200}
	// dc := gg.NewContext(image_size.width, image_size.height)
	im, err := gg.LoadImage("./template/dark-4.png")
	if err != nil {
		panic(err)
	}
	image_size.width = im.Bounds().Size().X
	image_size.height = im.Bounds().Size().Y
	dc := gg.NewContext(image_size.width, image_size.height)
	dc.DrawImage(im, 0, 0)

	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("font/Alibaba-PuHuiTi-Bold.ttf", 190); err != nil {
		panic(err)
	}
	dc.DrawString(text, 80, float64(image_size.height)-75)

	dc.SavePNG("output/output.png")
}

func main() {
	// color.Style{color.FgBlack, color.BgBlue}.Println("         ")
	// color.Style{color.FgBlack, color.BgBlue}.Println(" GoTabel ")

	s := color.S256(231, 27)
	s.SetOpts(color.Opts{color.OpBold})
	s.Println(" GoTabel v.0.1.0 ")
	println("By: YuzeTT ")
	fmt.Println()
	// 生成条形码
	// cs, _ := code128.Encode("A1001")
	// file, _ := os.Create("qr3.png")
	// defer file.Close()
	// qrCode, _ := barcode.Scale(cs, 350, 70)
	// png.Encode(file, qrCode)

	// ====== 下方为测试代码，未清空推送dev分支 ======

	// Print out the final choice.
	fmt.Println("-----------------\n1. 创建四位数标签\n2. 查询标签(未完成)\nf. 进入批量创建(未完成)\nq. 退出程序")
	fmt.Printf("-----------------\n请选择：")
	var selectMode string
	fmt.Scanln(&selectMode)
	switch selectMode {
	case "1":
		fmt.Print("请输入要创建的识别码（1-6位英文/数字）：")
		var text string
		fmt.Scanln(&text)
		createLabel(text)
	case "2":
		fmt.Println("\n功能尚未完成")
	case "f":
		os.Exit(0)
	case "q":
		os.Exit(0)
	default:
		fmt.Print("未找到序号，请重试。\n\n")
	}
}
