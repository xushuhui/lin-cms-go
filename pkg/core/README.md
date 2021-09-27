
gin 验证
```go
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json tag的自定义方法
		//v.RegisterTagNameFunc(jsonTag)
		// 验证器注册翻译器
		e = zhtranslations.RegisterDefaultTranslations(v, trans)
		if e != nil {
			return
		}

		// 自定义验证方法
		e = v.RegisterValidation("mobile", mobileValidate)
		if e != nil {
			return
		}
		e = v.RegisterTranslation("mobile", trans, mobileRegister, mobileTranslation)
		if e != nil {
			return
		}

	}
```

自定义验证

```go
func mobileValidate(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)

}
func mobileRegister(ut ut.Translator) error {
	return ut.Add("mobile", "{0}填写不正确哦", true)
}
func mobileTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("mobile", fe.Field())
	return t
}
```