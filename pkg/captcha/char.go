package captcha

import (
	"bytes"
	"chat/pkg/constvalue"
	"chat/pkg/fonts"
	"chat/pkg/random"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

// ImageChar captcha engine char return type
type ImageChar struct {
	Item
	nrgba   *image.NRGBA
	Complex int
}

// ConfigCharacter captcha config
type ConfigCharacter struct {
	Height             int
	Width              int
	Mode               int
	ComplexOfNoiseText int
	ComplexOfNoiseDot  int
	IsUseSimpleFont    bool
	IsShowNoiseText    bool
	IsShowNoiseDot     bool
	IsShowSlimeLine    bool
	IsShowHollowLine   bool
	IsShowSineLine     bool
	CaptchaLen         int
}

type point struct {
	X int
	Y int
}

func newCaptchaImage(w, h int, bg color.RGBA) (img *ImageChar, err error) {
	m := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)
	img = &ImageChar{}
	img.nrgba = m
	img.ImageHeight = h
	img.ImageWidth = w
	return
}

// drawHollowLine draw strong and bold white line
func (captcha *ImageChar) drawHollowLine() *ImageChar {
	first := captcha.ImageWidth / 20
	end := first * 19
	lineColor := color.RGBA{R: 245, G: 250, B: 252, A: 255}
	x1 := float64(rand.Intn(first))
	x2 := float64(rand.Intn(first) + end)
	multiple := float64(rand.Intn(5)+3) / float64(5)
	if 0 == int(multiple*10)%3 {
		multiple = multiple * -1.0
	}
	w := captcha.ImageHeight / 20
	for ; x1 < x2; x1++ {
		y := math.Sin(x1*math.Pi*multiple/float64(captcha.ImageWidth)) * float64(captcha.ImageHeight/3)
		if multiple < 0 {
			y = y + float64(captcha.ImageHeight/2)
		}
		captcha.nrgba.Set(int(x1), int(y), lineColor)
		for i := 0; i <= w; i++ {
			captcha.nrgba.Set(int(x1), int(y)+i, lineColor)
		}
	}
	return captcha
}

// drawSineLine draw a sine line
func (captcha *ImageChar) drawSineLine() *ImageChar {
	// 振幅
	a := rand.Intn(captcha.ImageHeight / 2)
	// Y 轴方向的偏移量
	b := random.Random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))
	// X 轴方向的偏移量
	f := random.Random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))
	// 周期
	var t float64
	if captcha.ImageHeight > captcha.ImageWidth/2 {
		t = random.Random(int64(captcha.ImageWidth/2), int64(captcha.ImageHeight))
	} else if captcha.ImageHeight == captcha.ImageWidth/2 {
		t = float64(captcha.ImageHeight)
	} else {
		t = random.Random(int64(captcha.ImageHeight), int64(captcha.ImageWidth/2))
	}
	w := float64((2 * math.Pi) / t)
	// 曲线横坐标起始位置
	px1 := 0
	px2 := int(random.Random(int64(float64(captcha.ImageWidth)*0.8), int64(captcha.ImageWidth)))
	// 颜色
	c := color.RGBA{R: uint8(rand.Intn(150)), G: uint8(rand.Intn(150)), B: uint8(rand.Intn(151)), A: uint8(255)}

	var py float64
	for px := px1; px < px2; px++ {
		if 0 != w {
			py = float64(a)*math.Sin(w*float64(px)+f) + b + (float64(captcha.ImageWidth) / float64(5))
			i := captcha.ImageHeight / 5
			for i > 0 {
				captcha.nrgba.Set(px+i, int(py), c)
				i--
			}
		}
	}
	return captcha
}

// drawSlimLine draw n slim random color lines
func (captcha *ImageChar) drawSlimLine(num int) *ImageChar {
	first := captcha.ImageWidth / 10
	end := first * 9
	y := captcha.ImageHeight / 3
	for i := 0; i < num; i++ {
		point1 := point{X: rand.Intn(first), Y: rand.Intn(y)}
		point2 := point{X: rand.Intn(first) + end, Y: rand.Intn(y)}
		if 0 == i%2 {
			point1.Y = rand.Intn(y) + y*2
			point2.Y = rand.Intn(y)
		} else {
			point1.Y = rand.Intn(y) + y*(i%2)
			point2.Y = rand.Intn(y) + y*2
		}
		captcha.drawBeeLine(point1, point2, random.RandDeepColor())
	}
	return captcha
}

func (captcha *ImageChar) drawBeeLine(point1, point2 point, lineColor color.RGBA) {
	dx := math.Abs(float64(point1.X - point2.X))
	dy := math.Abs(float64(point1.Y - point2.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		captcha.nrgba.Set(point1.X, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+2, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
}

// drawDotNoise draw noise dots
func (captcha *ImageChar) drawDotNoise(complex int) *ImageChar {
	density := 16
	if constvalue.CaptchaComplexLower == complex {
		density = 32
	} else if constvalue.CaptchaComplexMedium == complex {
		density = 16
	} else if constvalue.CaptchaComplexHigh == complex {
		density = 8
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density
	for i := 0; i < maxSize; i++ {
		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)
		captcha.nrgba.Set(rw, rh, random.RandColor())
		sz := rand.Intn(maxSize)
		if 0 == sz%3 {
			captcha.nrgba.Set(rw+1, rh+1, random.RandColor())
		}
	}
	return captcha
}

// drawTextNoise draw noise text
func (captcha *ImageChar) drawTextNoise(complex int, isSimpleFont bool) {
	density := 1500
	if constvalue.CaptchaComplexLower == complex {
		density = 2000
	} else if constvalue.CaptchaComplexMedium == complex {
		density = 1500
	} else if constvalue.CaptchaComplexHigh == complex {
		density = 1000
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density
	c := freetype.NewContext()
	c.SetDPI(constvalue.ImageStringDpi)
	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)
	rawFontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(7))/float64(10))
	for i := 0; i < maxSize; i++ {
		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)
		text := random.RandText(1, constvalue.TxtNumbers+constvalue.TxtAlphabet)
		fontSize := rawFontSize/2 + float64(rand.Intn(5))
		c.SetSrc(image.NewUniform(random.RandLightColor()))
		c.SetFontSize(fontSize)
		if isSimpleFont {
			c.SetFont(fonts.TrueTypeFontFamilys[0])
		} else {
			c.SetFont(fonts.RandFontFamily())
		}
		pt := freetype.Pt(rw, rh)
		if _, err := c.DrawString(text, pt); nil != err {
			log.Println(err)
		}
	}
	return
}

// drawText draw captcha string to image
func (captcha *ImageChar) drawText(text string, isSimpleFont bool) {
	c := freetype.NewContext()
	c.SetDPI(constvalue.ImageStringDpi)
	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)
	fontWidth := captcha.ImageWidth / len(text)
	for i, s := range text {
		fontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(7))/float64(9))
		c.SetSrc(image.NewUniform(random.RandDeepColor()))
		c.SetFontSize(fontSize)
		if isSimpleFont {
			c.SetFont(fonts.TrueTypeFontFamilys[0])
		} else {
			c.SetFont(fonts.RandFontFamily())
		}
		x := int(fontWidth)*i + int(fontWidth)/int(fontSize)
		y := 5 + rand.Intn(captcha.ImageHeight/2) + int(fontSize/2)
		pt := freetype.Pt(x, y)
		if _, err := c.DrawString(string(s), pt); nil != err {
			log.Println(err)
		}
	}
	return
}

// BinaryEncodeing encode image to binary
func (captcha *ImageChar) BinaryEncodeing() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, captcha.nrgba); nil != err {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo writes image in png format into the given writer
func (captcha *ImageChar) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(captcha.BinaryEncodeing())
	return int64(n), err
}

// CreateCharEngine create captcha with config struct
func CreateCharEngine(config ConfigCharacter) *ImageChar {
	img, err := newCaptchaImage(config.Width, config.Height, random.RandLightColor())
	if nil != err {
		log.Println(err)
	}
	if config.IsShowNoiseDot {
		img.drawDotNoise(config.ComplexOfNoiseDot)
	}
	if config.IsShowHollowLine {
		img.drawHollowLine()
	}
	if config.IsShowNoiseText {
		img.drawTextNoise(config.ComplexOfNoiseText, config.IsUseSimpleFont)
	}
	if config.IsShowSlimeLine {
		img.drawSlimLine(4)
	}
	if config.IsShowSineLine {
		img.drawSineLine()
	}
	var content string
	switch config.Mode {
	case constvalue.CaptchaModeAlphabet:
		content = random.RandText(config.CaptchaLen, constvalue.TxtAlphabet)
		img.VerifyValue = content
	case constvalue.CaptchaModeArithmetic:
		content, img.VerifyValue = random.RandArithmetic()
	case constvalue.CaptchaModeNumber:
		content = random.RandText(config.CaptchaLen, constvalue.TxtNumbers)
		img.VerifyValue = content
	default:
		content = random.RandText(config.CaptchaLen, constvalue.TxtNumbers)
		img.VerifyValue = content
	}
	img.drawText(content, config.IsUseSimpleFont)
	img.Content = content
	return img
}
