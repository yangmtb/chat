package e

// msgFlags .
var msgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "请求参数错误",

	ERROR_CAPTCHA_GET_FAIL:    "请求验证码失败",
	ERROR_CAPTCHA_VERIFY_FAIL: "验证验证码失败",

	ERROR_ACCOUNT_SIGN_IN_FAIL:   "用户名或密码错误",
	ERROR_ACCOUNT_SIGN_UP_FAIL:   "注册失败",
	ERROR_ACCOUNT_USERNAME_EXIST: "用户名已存在",
	ERROR_ACCOUNT_PHONE_EXIST:    "手机号已存在",
	ERROR_ACCOUNT_EMAIL_EXIST:    "邮箱已存在",

	ERROR_TOKEN_GET_TOKEN_FAIL: "获取token错误",
	ERROR_TOKEN_CHECK_FAIL:     "token检查失败",
	ERROR_TOKEN_TIMEOUT_FAIL:   "token超时",
}

// GetMsg get msg
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}
