package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hqwangningbo/nbgo/global"
	"github.com/hqwangningbo/nbgo/utils"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (baseApi *BaseApi) BuildRequest(bro BuildRequestOption) *BaseApi {
	var errResult error

	//绑定请求上下文

	baseApi.Ctx = bro.Ctx

	if bro.DTO != nil {
		if bro.BindParamsFromUri {
			errResult = baseApi.Ctx.ShouldBindUri(bro.DTO)
		} else {
			errResult = baseApi.Ctx.ShouldBind(bro.DTO)
		}

		if errResult != nil {
			errResult = baseApi.ParseValidateErrors(errResult.(validator.ValidationErrors), bro.DTO)
			baseApi.AddError(errResult)
			Fail(baseApi.Ctx, ResponseJson{
				Msg: baseApi.GetError().Error(),
			})
		}
	}
	return baseApi
}

func (baseApi *BaseApi) AddError(newErr error) {
	baseApi.Errors = utils.AppendError(baseApi.Errors, newErr)
}

func (baseApi *BaseApi) GetError() error {
	return baseApi.Errors
}

// 解析错误
func (baseApi *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取指定元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessagetag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessagetag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

func (baseApi *BaseApi) Fail(resp ResponseJson) {
	Fail(baseApi.Ctx, resp)
}
func (baseApi *BaseApi) OK(resp ResponseJson) {
	OK(baseApi.Ctx, resp)
}
func (baseApi *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(baseApi.Ctx, resp)
}
