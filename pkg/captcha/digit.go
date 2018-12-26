package captcha

import (
	"chat/pkg/constvalue"
	"chat/pkg/random"
	"chat/pkg/rng"
	"image"
	"image/color"
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
		uint8(m.rg.Intn(129)),
		uint8(m.rg.Intn(129)),
		uint8(m.rg.Intn(129)),
		0xFF,
	}
	p[1] = prim
	for i := 2; i <= constvalue.DotCount; i++ {
		p[i] = m.randomBrightness(prim, 255)
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
	// fw takes into account 1-dot spacing between
}

// CreateDigitsEngine create digits engine
func CreateDigitsEngine(id string, config ConfigDigit) (m *ImageDigit) {
	digits := random.RandomDigits(config.CaptchaLen)
	// initialize PRNG
	m = new(ImageDigit)
	m.VerifyValue = random.ParseDigitsToString(digits)
	m.rg.Seed(random.DeriveSeed(random.ImageSeedPurpose, id, digits))
	m.Paletted = image.NewPaletted(image.Rect(0, 0, config.Width, config.Height), m.getRandomPalette())

	return
}
