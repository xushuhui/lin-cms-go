package core

import (
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidate() {
	validate = validator.New()
	// 中文翻译
	zh := zh2.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	e := zhtranslations.RegisterDefaultTranslations(validate, trans)
	if e != nil {
		panic(e)
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

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
