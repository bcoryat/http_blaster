package request_generators

import (
	"github.com/v3io/http_blaster/httpblaster/config"
	"github.com/valyala/fasthttp"
)

type Generator interface {
	UseCommon(c RequestCommon)
	GenerateRequests(workload config.Workload, tls_mode bool, host string) chan *fasthttp.Request
}
