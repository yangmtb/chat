package captcha

import (
	"bytes"
	"chat/pkg/constvalue"
	"chat/pkg/fonts"
	"chat/pkg/random"
	"chat/pkg/rng"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
)

// ConfigDigit digit captcha config
type ConfigDigit struct {
	Height     int
	Width      int
	CaptchaLen int
	MaxSkew    float64
	DotCount   int
}

// ImageDigit digit captcha struct
type ImageDigit struct {
	Item
	*image.Paletted
	dotSize int
	rg      rng.Siprng
}

func (captcha *ImageDigit) getRandomPalette() color.Palette {
	p := make([]color.Color, constvalue.DotCount+1)
	// transparent color
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// primary color
	prim := color.RGBA{
		uint8(captcha.rg.Intn(129)),
		uint8(captcha.rg.Intn(129)),
		uint8(captcha.rg.Intn(129)),
		0xFF,
	}
	p[1] = prim
	for i := 2; i <= constvalue.DotCount; i++ {
		p[i] = captcha.randomBrightness(prim, 255)
	}
	return p
}

func (captcha *ImageDigit) calculateSizes(width, height, n int) {
	// goal: fit all digits inside the image
	var border int
	if width > height {
		border = height / 4
	} else {
		border = width / 4
	}
	// convert everything to floats for calculations
	w := float64(width - border*2)
	h := float64(height - border*2)
	// fw takes into account 1-dot spacing between digits
	fw := float64(constvalue.DigitFontWidth + 1)
	fh := float64(constvalue.DigitFontHeight)
	nc := float64(n)
	// calculate the width of a single digit taking into account only the width of the image
	nw := w / nc
	// calculate the height of a digit from this width
	nh := nw * fh / fw
	// digit too high?
	if nh > h {
		nh = h
		nw = fw / fh * nh
	}
	// calculate dot size
	captcha.dotSize = int(nh / fh)
	if captcha.dotSize < 1 {
		captcha.dotSize = 1
	}
	// save everything, making the actual width smaller by 1 dot to account for spacing between digits.
	captcha.ImageWidth = int(nw) - captcha.dotSize
	captcha.ImageHeight = int(nh)
}

func (captcha *ImageDigit) drawHorizLine(fromX, toX, y int, colorIdx uint8) {
	for x := fromX; x <= toX; x++ {
		captcha.SetColorIndex(x, y, colorIdx)
	}
}

func (captcha *ImageDigit) drawCircle(x, y, radius int, colorIdx uint8) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	x0 := 0
	y0 := radius
	captcha.SetColorIndex(x, y+radius, colorIdx)
	captcha.SetColorIndex(x, y-radius, colorIdx)
	captcha.drawHorizLine(x-radius, x+radius, y, colorIdx)
	for x0 < y0 {
		if f >= 0 {
			y0--
			dfy += 2
			f += dfy
		}
		x0++
		dfx += 2
		f += dfx
		captcha.drawHorizLine(x-x0, x+x0, y+y0, colorIdx)
		captcha.drawHorizLine(x-x0, x+x0, y-y0, colorIdx)
		captcha.drawHorizLine(x-y0, x+y0, y+x0, colorIdx)
		captcha.drawHorizLine(x-y0, x+y0, y-x0, colorIdx)
	}
}

func (captcha *ImageDigit) fillWithCircles(n, maxradius int) {
	maxx := captcha.Bounds().Max.X
	maxy := captcha.Bounds().Max.Y
	for i := 0; i < n; i++ {
		colorIdx := uint8(captcha.rg.Int(1, constvalue.DotCount-1))
		r := captcha.rg.Int(1, maxradius)
		captcha.drawCircle(captcha.rg.Int(r, maxx-r), captcha.rg.Int(r, maxy-r), r, colorIdx)
	}
}

func (captcha *ImageDigit) strikeThrough() {
	maxx := captcha.Bounds().Max.X
	maxy := captcha.Bounds().Max.Y
	y := captcha.rg.Int(maxy/3, maxy-maxy/3)
	amplitude := captcha.rg.Float(5, 20)
	period := captcha.rg.Float(80, 180)
	dx := 2.0 * math.Pi / period
	for x := 0; x < maxx; x++ {
		x0 := amplitude * math.Cos(float64(y)*dx)
		y0 := amplitude * math.Sin(float64(x)*dx)
		for yn := 0; yn < captcha.dotSize; yn++ {
			r := captcha.rg.Int(0, captcha.dotSize)
			captcha.drawCircle(x+int(x0), y+int(y0)+(yn*captcha.dotSize), r/2, 1)
		}
	}
}

func (captcha *ImageDigit) drawDigit(digit []byte, x, y int) {
	skf := captcha.rg.Float(-constvalue.MaxSkew, constvalue.MaxSkew)
	xs := float64(x)
	r := captcha.dotSize / 2
	y += captcha.rg.Int(-r, r)
	for y0 := 0; y0 < constvalue.DigitFontHeight; y0++ {
		for x0 := 0; x0 < constvalue.DigitFontWidth; x0++ {
			if digit[y0*constvalue.DigitFontWidth+x0] != constvalue.DigitFontBlackChar {
				continue
			}
			captcha.drawCircle(x+x0*captcha.dotSize, y+y0*captcha.dotSize, r, 1)
		}
		xs += skf
		x = int(xs)
	}
}

func (captcha *ImageDigit) distort(amplude, period float64) {
	w := captcha.Bounds().Max.X
	h := captcha.Bounds().Max.Y
	oldm := captcha.Paletted
	newm := image.NewPaletted(image.Rect(0, 0, w, h), oldm.Palette)
	dx := 2.0 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			x0 := amplude * math.Sin(float64(y)*dx)
			y0 := amplude * math.Cos(float64(x)*dx)
			newm.SetColorIndex(x, y, oldm.ColorIndexAt(x+int(x0), y+int(y0)))
		}
	}
	captcha.Paletted = newm
}

func (captcha *ImageDigit) randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return c
	}
	n := captcha.rg.Intn(int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}

// BinaryEncodeing encode
func (captcha *ImageDigit) BinaryEncodeing() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, captcha.Paletted); nil != err {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo .
func (captcha *ImageDigit) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(captcha.BinaryEncodeing())
	return int64(n), err
}

// CreateDigitsEngine create digits engine
func CreateDigitsEngine(id string, config ConfigDigit) (img *ImageDigit) {
	digits := random.RandomDigits(config.CaptchaLen)
	// initialize PRNG
	img = new(ImageDigit)
	img.VerifyValue = random.ParseDigitsToString(digits)
	img.rg.Seed(random.DeriveSeed(random.ImageSeedPurpose, id, digits))
	img.Paletted = image.NewPaletted(image.Rect(0, 0, config.Width, config.Height), img.getRandomPalette())
	img.calculateSizes(config.Width, config.Height, len(digits))
	maxx := config.Width - (img.ImageWidth+img.dotSize)*len(digits) - img.dotSize
	maxy := config.Height - img.ImageHeight - img.dotSize*2
	var border int
	if config.Width > config.Height {
		border = config.Height / 5
	} else {
		border = config.Width / 5
	}
	x := img.rg.Int(border, maxx-border)
	y := img.rg.Int(border, maxy-border)
	for _, n := range digits {
		img.drawDigit(fonts.DigitFontData[n], x, y)
		x += img.ImageWidth + img.dotSize
	}
	img.strikeThrough()
	img.distort(img.rg.Float(5, 10), img.rg.Float(100, 200))
	img.fillWithCircles(constvalue.DotCount, img.dotSize)
	return
}
