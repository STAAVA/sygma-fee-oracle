package api

import (
	"errors"

	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/gin-gonic/gin"
)

type ReturnErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type ReturnSuccessResponse interface{}

func newReturnErrorResp(fe *config.FeeOracleError, err error) *ReturnErrorResponse {
	if err == nil {
		err = errors.New("")
	}
	return &ReturnErrorResponse{
		Code:    fe.Code,
		Message: fe.Message,
		Detail:  fe.Wrap(err).Error(),
	}
}

func ginErrorReturn(c *gin.Context, httpCode int, resp *ReturnErrorResponse) {
	c.JSON(httpCode, gin.H{
		"error": resp,
	})
}

func ginSuccessReturn(c *gin.Context, httpCode int, resp ReturnSuccessResponse) {
	c.JSON(httpCode, gin.H{
		"response": resp,
	})
}

type FetchRateResp struct {
	BaseRate                 string       `json:"baseEffectiveRate"`
	TokenRate                string       `json:"tokenEffectiveRate"`
	DestinationChainGasPrice GasPriceResp `json:"dstGasPrice"`
	Signature                string       `json:"signature"`
	FromDomainID             int          `json:"fromDomainID"`
	ToDomainID               int          `json:"toDomainID"`
	ResourceID               int          `json:"resourceID"`
	Timestamp                int64        `json:"timestamp"`
}

type GasPriceResp struct {
	SafeGasPrice    string `json:"safeGasPrice" `
	ProposeGasPrice string `json:"proposeGasPrice"`
	FastGasPrice    string `json:"fastGasPrice" `
}
