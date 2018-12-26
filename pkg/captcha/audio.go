package captcha

import (
	"bytes"
	"chat/pkg/random"
	"chat/pkg/rng"
	"encoding/binary"
	"io"
	"math"
)

// ConfigAudio config
type ConfigAudio struct {
	CaptchaLen int
	Language   string
}

const (
	sampleRate = 8000 // Hz
)

var (
	endingBeepSound []byte
)

func init() {
	endingBeepSound = changeSpeed(beepSound, 1.4)
}

// Audio ..
type Audio struct {
	Item
	body        *bytes.Buffer
	digitSounds [][]byte
	rg          rng.Siprng
}

// CreateAudioEngine .
func CreateAudioEngine(id string, config ConfigAudio) *Audio {
	digits := random.RandomDigits(config.CaptchaLen)
	audio := newAudio(id, digits, config.Language)
	audio.VerifyValue = random.ParseDigitsToString(digits)
	return audio
}

// newAudio ....
func newAudio(id string, digits []byte, lang string) (a *Audio) {
	a = new(Audio)
	a.rg.Seed(random.DeriveSeed(random.AudioSeedPurpose, id, digits))
	if sounds, ok := digitSounds[lang]; ok {
		a.digitSounds = sounds
	} else {
		a.digitSounds = digitSounds["en"]
	}
	numsnd := make([][]byte, len(digits))
	for i, n := range digits {
		snd := a.randomizedDigitSound(n)
		numsnd[i] = snd
	}
	intervals := make([]int, len(digits)+1)
	intdur := 0
	for i := range intervals {
		dur := a.rg.Int(sampleRate, sampleRate*3)
		intdur += dur
		intervals[i] = dur
	}
	bg := a.makeBackgroundSound(a.longestDigitSndLen()*len(digits) + intdur)
	sil := makeSilence(sampleRate / 5)
	bufcap := 3*len(beepSound) + 2*len(sil) + len(bg) + len(endingBeepSound)
	a.body = bytes.NewBuffer(make([]byte, 0, bufcap))
	a.body.Write(beepSound)
	a.body.Write(sil)
	a.body.Write(beepSound)
	a.body.Write(sil)
	a.body.Write(beepSound)
	pos := intervals[0]
	for i, v := range numsnd {
		mixSound(bg[pos:], v)
		pos += len(v) + intervals[i+1]
	}
	a.body.Write(bg)
	a.body.Write(endingBeepSound)
	return
}

func (a *Audio) makeBackgroundSound(length int) []byte {
	b := a.makeWhiteNoise(length, 4)
	for i := 0; i < length/(sampleRate/10); i++ {
		snd := reversedSound(a.digitSounds[a.rg.Intn(10)])
		snd = changeSpeed(snd, a.rg.Float(0.8, 1.4))
		place := a.rg.Intn(len(b) - len(snd))
		setSoundLevel(snd, a.rg.Float(0.2, 0.5))
		mixSound(b[place:], snd)
	}
	return b
}

func (a *Audio) randomizedDigitSound(n byte) []byte {
	s := a.randomSpeed(a.digitSounds[n])
	setSoundLevel(s, a.rg.Float(0.75, 1.2))
	return s
}

func (a *Audio) longestDigitSndLen() int {
	n := 0
	for _, v := range a.digitSounds {
		if n < len(v) {
			n = len(v)
		}
	}
	return n
}

func (a *Audio) randomSpeed(b []byte) []byte {
	pitch := a.rg.Float(0.9, 1.2)
	return changeSpeed(b, pitch)
}

func (a *Audio) makeWhiteNoise(length int, level uint8) []byte {
	noise := a.rg.Bytes(length)
	adj := 128 - level/2
	for i, v := range noise {
		v %= level
		v += adj
		noise[i] = v
	}
	return noise
}

func mixSound(dst, src []byte) {
	for i, v := range src {
		av := int(v)
		bv := int(dst[i])
		if av < 128 && bv < 128 {
			dst[i] = byte(av * bv / 128)
		} else {
			dst[i] = byte(2*(av+bv) - av*bv/128 - 256)
		}
	}
}

func setSoundLevel(a []byte, level float64) {
	for i, v := range a {
		av := float64(v)
		switch {
		case av > 128:
			if av = (av-128)*level + 128; av < 128 {
				av = 128
			}
		case av < 128:
			if av = 128 - (128-av)*level; av > 128 {
				av = 128
			}
		default:
			continue
		}
		a[i] = byte(av)
	}
}

func changeSpeed(a []byte, speed float64) []byte {
	b := make([]byte, int(math.Floor(float64(len(a))*speed)))
	var p float64
	for _, v := range a {
		for i := int(p); i < int(p+speed); i++ {
			b[i] = v
		}
		p += speed
	}
	return b
}

func makeSilence(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = 128
	}
	return b
}

func reversedSound(a []byte) []byte {
	n := len(a)
	b := make([]byte, n)
	for i, v := range a {
		b[n-1-i] = v
	}
	return b
}

// WriteTo .
func (a *Audio) WriteTo(w io.Writer) (n int64, err error) {
	bodyLen := uint32(a.body.Len())
	paddedBodyLen := bodyLen
	if 0 != bodyLen%2 {
		paddedBodyLen++
	}
	totalLen := uint32(len(waveHeader)) - 4 + paddedBodyLen
	header := make([]byte, len(waveHeader)+4)
	copy(header, waveHeader)
	binary.LittleEndian.PutUint32(header[4:], totalLen)
	binary.LittleEndian.PutUint32(header[len(waveHeader):], bodyLen)
	nn, err := w.Write(header)
	n = int64(nn)
	if nil != err {
		return
	}
	n, err = a.body.WriteTo(w)
	n += int64(nn)
	if nil != err {
		return
	}
	if bodyLen != paddedBodyLen {
		w.Write([]byte{0})
		n++
	}
	return
}

// BinaryEncodeing .
func (a *Audio) BinaryEncodeing() []byte {
	var buf bytes.Buffer
	if _, err := a.WriteTo(&buf); nil != err {
		panic(err.Error())
	}
	return buf.Bytes()
}
