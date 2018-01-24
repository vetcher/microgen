package template

import (
	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/devimteam/microgen/util"
)

type httpClientTemplate struct {
	Info *GenerationInfo
}

func NewHttpClientTemplate(info *GenerationInfo) Template {
	return &httpClientTemplate{
		Info: info.Copy(),
	}
}

func (t *httpClientTemplate) DefaultPath() string {
	return "./transport/http/client.go"
}

func (t *httpClientTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	if err := util.StatFile(t.Info.AbsOutPath, t.DefaultPath()); !t.Info.Force && err == nil {
		return nil, nil
	}
	return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}

func (t *httpClientTemplate) Prepare() error {
	tags := util.FetchTags(t.Info.Iface.Docs, TagMark+ForceTag)
	if util.IsInStringSlice("http", tags) || util.IsInStringSlice("http-client", tags) {
		t.Info.Force = true
	}
	return nil
}

// Render http client.
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package transporthttp
//
//		import (
//			svc "github.com/devimteam/microgen/example/svc"
//			http1 "github.com/devimteam/microgen/example/svc/transport/converter/http"
//			http "github.com/go-kit/kit/transport/http"
//			url "net/url"
//			strings "strings"
//		)
//
//		func NewHTTPClient(addr string, opts ...http.ClientOption) (svc.StringService, error) {
//			if !strings.HasPrefix(addr, "http") {
//				addr = "http://" + addr
//			}
//			u, err := url.Parse(addr)
//			if err != nil {
//				return nil, err
//			}
//			return &svc.Endpoints{
//				EmptyReqEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPEmptyReqRequest,
//					http1.DecodeHTTPEmptyReqResponse,
//					opts...,
//				).Endpoint(),
//				EmptyRespEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPEmptyRespRequest,
//					http1.DecodeHTTPEmptyRespResponse,
//					opts...,
//				).Endpoint(),
//				TestCaseEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPTestCaseRequest,
//					http1.DecodeHTTPTestCaseResponse,
//					opts...,
//				).Endpoint(),
//			}, nil
//		}
//
func (t *httpClientTemplate) Render() write_strategy.Renderer {
	f := NewFile("transporthttp")
	f.PackageComment(t.Info.FileHeader)
	f.PackageComment(`Please, do not edit.`)

	f.Func().Id("NewHTTPClient").Params(
		Id("addr").Id("string"),
		Id("opts").Op("...").Qual(PackagePathGoKitTransportHTTP, "ClientOption"),
	).Params(
		Qual(t.Info.ServiceImportPath, t.Info.Iface.Name),
		Error(),
	).Block(
		t.clientBody(),
	)

	return f
}

// Render client body.
//		if !strings.HasPrefix(addr, "http") {
//			addr = "http://" + addr
//		}
//		u, err := url.Parse(addr)
//		if err != nil {
//			return nil, err
//		}
//		return &svc.Endpoints{
//			EmptyReqEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPEmptyReqRequest,
//				http1.DecodeHTTPEmptyReqResponse,
//				opts...,
//			).Endpoint(),
//			EmptyRespEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPEmptyRespRequest,
//				http1.DecodeHTTPEmptyRespResponse,
//				opts...,
//			).Endpoint(),
//			TestCaseEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPTestCaseRequest,
//				http1.DecodeHTTPTestCaseResponse,
//				opts...,
//			).Endpoint(),
//		}, nil
//
func (t *httpClientTemplate) clientBody() *Statement {
	return If(
		Op("!").Qual(PackagePathStrings, "HasPrefix").Call(Id("addr"), Lit("http")),
	).Block(
		Id("addr").Op("=").Lit("http://").Op("+").Id("addr"),
	).
		Line().List(Id("u"), Err()).Op(":=").Qual(PackagePathUrl, "Parse").Call(Id("addr")).
		Line().If(Err().Op("!=").Nil()).
		Block(
			Return(Nil(), Err()),
		).
		Line().Return(Op("&").Qual(t.Info.ServiceImportPath, "Endpoints").Values(DictFunc(
		func(d Dict) {
			for _, fn := range t.Info.Iface.Methods {
				method := fetchMethodTag(fn.Docs)
				d[Id(endpointStructName(fn.Name))] = Qual(PackagePathGoKitTransportHTTP, "NewClient").Call(
					Line().Lit(method),
					Line().Id("u"),
					Line().Qual(pathToHttpConverter(t.Info.ServiceImportPath), httpEncodeRequestName(fn)),
					Line().Qual(pathToHttpConverter(t.Info.ServiceImportPath), httpDecodeResponseName(fn)),
					Line().Id("opts").Op("...").Line(),
				).Dot("Endpoint").Call()
			}
		},
	)), Nil())
}
