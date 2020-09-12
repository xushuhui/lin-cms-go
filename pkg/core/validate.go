package core

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

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

var trans ut.Translator

func initValidate() (e error) {
	// 中文翻译
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json tag的自定义方法
		//v.RegisterTagNameFunc(jsonTag)
		// 验证器注册翻译器
		e = zhtranslations.RegisterDefaultTranslations(v, trans)
		if e != nil {
			return
		}
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
		// 自定义验证方法
		e = v.RegisterValidation("mobile", mobileValidate)
		if e != nil {
			return
		}
		e = v.RegisterTranslation("mobile", trans, mobileRegister, mobileTranslation)
		if e != nil {
			return
		}
		//v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		//	return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("required", fe.Field())
		//
		//	return t
		//})
	}
	return
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string

	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}
