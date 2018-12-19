package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yeqown/gateway/config/rule"
	"github.com/yeqown/gateway/utils"
	"github.com/yeqown/server-common/code"
)

type proxycfgSrvsForm struct {
	Limit  int `form:"limit" valid:"gte=0,lte=10"`
	Offset int `form:"offset" valid:"gte=0"`
}

type proxycfgSrvsResp struct {
	code.CodeInfo
	Rules []rule.ServerRuler `json:"rules"`
	Total int                `json:"total"`
}

// ProxyConfigSrvsGET ...
func ProxyConfigSrvsGET(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var (
		form = new(proxycfgSrvsForm)
		resp = new(proxycfgSrvsResp)
	)

	if err := bind(form, req); err != nil {
		responseWithError(w, resp, err)
		return
	}

	if err := valid(form); err != nil {
		responseWithError(w, resp, err)
		return
	}

	if form.Limit == 0 {
		form.Limit = 10
	}

	resp.Rules = Global().ServerRulesPage(form.Offset, form.Offset+form.Limit)
	resp.Total = Global().ServerRulesCount()

	code.FillCodeInfo(resp, code.GetCodeInfo(code.CodeOk))
	utils.ResponseJSON(w, resp)
}

type proxycfgSrvResp struct {
	code.CodeInfo
	Rule rule.ServerRuler `json:"rule"`
}

// ProxyConfigSrvGET ...
func ProxyConfigSrvGET(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var (
		resp = new(proxycfgSrvResp)
	)

	id := param.ByName("id")
	resp.Rule = Global().ServerRuleByID(id)

	code.FillCodeInfo(resp, code.GetCodeInfo(code.CodeOk))
	utils.ResponseJSON(w, resp)
}

// ProxyConfigSrvPOST ...
func ProxyConfigSrvPOST(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var (
		form = new(formServerRuler)
		resp = new(commonResp)
	)
	if err := bind(form, req); err != nil {
		responseWithError(w, resp, err)
		return
	}

	if err := valid(form); err != nil {
		responseWithError(w, resp, err)
		return
	}

	if err := Global().NewServerRule(form); err != nil {
		responseWithError(w, resp, err)
		return
	}

	code.FillCodeInfo(resp, code.GetCodeInfo(code.CodeOk))
	utils.ResponseJSON(w, resp)
}

// ProxyConfigSrvPUT ...
func ProxyConfigSrvPUT(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var (
		form = new(formServerRuler)
		resp = new(commonResp)
	)
	if err := bind(form, req); err != nil {
		responseWithError(w, resp, err)
		return
	}

	if err := valid(form); err != nil {
		responseWithError(w, resp, err)
		return
	}
	id := param.ByName("id")
	if err := Global().UpdateServerRule(id, form); err != nil {
		responseWithError(w, resp, err)
		return
	}

	code.FillCodeInfo(resp, code.GetCodeInfo(code.CodeOk))
	utils.ResponseJSON(w, resp)
}

// ProxyConfigSrvDELETE ...
func ProxyConfigSrvDELETE(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var (
		resp = new(commonResp)
	)

	id := param.ByName("id")
	if err := Global().DelServerRule(id); err != nil {
		responseWithError(w, resp, err)
		return
	}

	code.FillCodeInfo(resp, code.GetCodeInfo(code.CodeOk))
	utils.ResponseJSON(w, resp)
}
