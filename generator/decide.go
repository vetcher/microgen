package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	mstrings "github.com/devimteam/generator/strings"
	"github.com/devimteam/microgen/generator/template"
	"github.com/devimteam/microgen/util"
	"github.com/vetcher/godecl/types"
)

const (
	TagMark         = template.TagMark
	MicrogenMainTag = template.MicrogenMainTag
	ProtobufTag     = "protobuf"
	GRPCRegAddr     = "grpc-addr"

	MiddlewareTag        = template.MiddlewareTag
	LoggingMiddlewareTag = template.LoggingMiddlewareTag
	RecoverMiddlewareTag = template.RecoverMiddlewareTag
	HttpTag              = template.HttpTag
	HttpServerTag        = template.HttpServerTag
	HttpClientTag        = template.HttpClientTag
	GrpcTag              = template.GrpcTag
	GrpcServerTag        = template.GrpcServerTag
	GrpcClientTag        = template.GrpcClientTag
	MainTag              = template.MainTag
)

func ListTemplatesForGen(iface *types.Interface, force bool, importPackageName, absOutPath, sourcePath string) (units []*generationUnit, err error) {
	importPackagePath, err := resolvePackagePath(absOutPath)
	if err != nil {
		return nil, err
	}
	absSourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}
	info := &template.GenerationInfo{
		ServiceImportPackageName: importPackageName,
		ServiceImportPath:        importPackagePath,
		Force:                    force,
		Iface:                    iface,
		AbsOutPath:               absOutPath,
		SourceFilePath:           absSourcePath,
		ProtobufPackage:          mstrings.FetchMetaInfo(TagMark+ProtobufTag, iface.Docs),
		GRPCRegAddr:              mstrings.FetchMetaInfo(TagMark+GRPCRegAddr, iface.Docs),
	}
	stubSvc, err := NewGenUnit(template.NewStubInterfaceTemplate(info), absOutPath)
	if err != nil {
		return nil, err
	}
	exch, err := NewGenUnit(template.NewExchangeTemplate(info), absOutPath)
	if err != nil {
		return nil, err
	}
	endp, err := NewGenUnit(template.NewEndpointsTemplate(info), absOutPath)
	if err != nil {
		return nil, err
	}
	units = append(units, stubSvc, exch, endp)

	genTags := util.FetchTags(iface.Docs, TagMark+MicrogenMainTag)
	fmt.Println("Tags:", strings.Join(genTags, ", "))
	for _, tag := range genTags {
		templates := tagToTemplate(tag, info)
		if templates == nil {
			fmt.Printf("Warning! unexpected tag %s\n", tag)
			continue
		}
		for _, t := range templates {
			unit, err := NewGenUnit(t, absOutPath)
			if err != nil {
				return nil, err
			}
			units = append(units, unit)
		}
	}
	return units, nil
}

func tagToTemplate(tag string, info *template.GenerationInfo) (tmpls []template.Template) {
	switch tag {
	case MiddlewareTag:
		return append(tmpls, template.NewMiddlewareTemplate(info))
	case LoggingMiddlewareTag:
		return append(tmpls, template.NewLoggingTemplate(info))
	case GrpcTag:
		return append(tmpls,
			template.NewGRPCClientTemplate(info),
			template.NewGRPCServerTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case GrpcClientTag:
		return append(tmpls,
			template.NewGRPCClientTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case GrpcServerTag:
		return append(tmpls,
			template.NewGRPCServerTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case HttpTag:
		return append(tmpls,
			template.NewHttpServerTemplate(info),
			template.NewHttpClientTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case HttpServerTag:
		return append(tmpls,
			template.NewHttpServerTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case HttpClientTag:
		return append(tmpls,
			template.NewHttpClientTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case RecoverMiddlewareTag:
		return append(tmpls, template.NewRecoverTemplate(info))
	case MainTag:
		return append(tmpls, template.NewMainTemplate(info))
	}
	return nil
}

func resolvePackagePath(outPath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("GOPATH is empty")
	}

	absOutPath, err := filepath.Abs(outPath)
	if err != nil {
		return "", err
	}

	gopathSrc := filepath.Join(gopath, "src")
	if !strings.HasPrefix(absOutPath, gopathSrc) {
		return "", fmt.Errorf("path not in GOPATH")
	}

	return absOutPath[len(gopathSrc)+1:], nil
}
