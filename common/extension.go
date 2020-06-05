package common

import (
	"fmt"
	"github.com/cargod-bj/b2c-common/logger"
	"github.com/cargod-bj/b2c-common/utils"
)
import "github.com/cargod-bj/b2c-common/resp"

// 对Response对象做的扩展方法，便于数据处理
type ResponseExtension interface {

	// 初始化当前response的code和msg为成功
	InitSuccess()

	// 解析response中的data到指定out中
	// 有可能报错，参见：resp.FAILED_DTO_DATA_NIL、resp.FAILED_DTO_DECODE
	// 返回结果(否有错误)：true 处理失败，false 处理成功
	ParseData2Dto(out interface{}) bool

	// 解析微服务入参中的dto到po中
	// 有可能报错，参见：resp.FAILED_DTO_DECODE
	// 返回结果(否有错误)：true 处理失败，false 处理成功
	ParseParams2Dto(in, out interface{}) bool

	HoldError(err error) bool

	GetFormatMsg() string

	// 验证当前response是否已经发生错误
	IsError() bool
}

func (x *Response) ParseData2Dto(out interface{}) bool {
	data := x.Data
	if data == nil {
		x.Code = resp.FAILED_DTO_DATA_NIL
		x.Msg = resp.FAILED_DTO_DATA_NIL_MSG
		return true
	}
	if err := utils.DecodeDto(data, out); err != nil {
		x.Code = resp.FAILED_DTO_DECODE
		x.Msg = resp.FAILED_DTO_DECODE_MSG
		logger.Info("解析数据错误", x, err)
		return true
	}
	return false
}

func (x *Response) ParseParams2Dto(in, out interface{}) bool {
	data := in
	if data == nil {
		x.Code = resp.FAILED_DTO_DECODE
		x.Msg = resp.FAILED_DTO_DECODE_MSG
		return true
	}
	if err := utils.DecodeDto(data, out); err != nil {
		x.Code = resp.FAILED_DTO_DECODE
		x.Msg = resp.FAILED_DTO_DECODE_MSG
		logger.Info("解析数据错误", x, err)
		return true
	}
	return false
}

func (x *Response) HoldError(err error) bool {
	if err != nil {
		x.Code = resp.FAILED_UNKNOWN
		x.Msg = resp.FAILED_UNKNOWN_MSG
		logger.Info("数据处理错误", err)
		return true
	}
	return false
}

func (x *Response) IsError() bool {
	return x.Code != "" && x.Code != resp.SUCCESS
}

func (x *Response) InitSuccess() {
	x.Code = resp.SUCCESS
	x.Msg = resp.SUCCESS_MSG
}

func (x *Response) GetFormatMsg() string {
	return fmt.Sprintf("%s(%s)", x.Msg, x.Code)
}
