package constvalue

const (
	// ImageStringDpi image string dpi
	ImageStringDpi = 72.0
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz"
	//TxtSimpleCharaters simple numbers and alphabet
	TxtSimpleCharaters = "13467ertyiadfhjkxcvbnERTYADFGHJKXCVBN"

	//MimeTypeCaptchaAudio output base64 mine-type.
	MimeTypeCaptchaAudio = "audio/wav"
	//MimeTypeCaptchaImage output base64 mine-type.
	MimeTypeCaptchaImage = "image/png"

	//FileExtCaptchaAudio output file extension.
	FileExtCaptchaAudio = "wav"
	//FileExtCaptchaImage output file extension.
	FileExtCaptchaImage = "png"
)

const (
	//CaptchaComplexLower complex level lower.
	CaptchaComplexLower = iota
	//CaptchaComplexMedium complex level medium.
	CaptchaComplexMedium
	//CaptchaComplexHigh complex level high.
	CaptchaComplexHigh
)

const (
	//CaptchaModeNumber mode number.
	CaptchaModeNumber = iota
	//CaptchaModeAlphabet mode alphabet.
	CaptchaModeAlphabet
	//CaptchaModeArithmetic mode arithmetic.
	CaptchaModeArithmetic
	//CaptchaModeNumberAlphabet mode mix number and alphabet,this is also default mode.
	CaptchaModeNumberAlphabet
)

const (
	// DefaultLen Default number of digits in captcha solution.
	// 默认数字验证长度.
	DefaultLen = 6
	// MaxSkew max absolute skew factor of a single digit.
	// 图像验证码的最大干扰系数.
	MaxSkew = 0.7
	// DotCount Number of background circles.
	// 图像验证码干扰圆点的数量.
	DotCount = 20
)

const (
	// DigitFontWidth .
	DigitFontWidth = 11
	// DigitFontHeight .
	DigitFontHeight = 18
	// DigitFontBlackChar .
	DigitFontBlackChar = 1
)
