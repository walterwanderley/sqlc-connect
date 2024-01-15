package metadata

import (
	"fmt"
	"strings"

	"github.com/walterwanderley/sqlc-grpc/converter"
	"github.com/walterwanderley/sqlc-grpc/metadata"
)

func InputGrpc(s *metadata.Service) []string {
	res := make([]string, 0)
	if s.EmptyInput() {
		return res
	}

	if s.HasCustomParams() {
		typ := s.InputTypes[0]
		in := s.InputNames[0]
		if strings.HasPrefix(typ, "*") {
			res = append(res, fmt.Sprintf("%s := new(%s)", in, typ[1:]))
		} else {
			res = append(res, fmt.Sprintf("var %s %s", in, typ))
		}
		m := s.Messages[converter.CanonicalName(typ)]
		for _, f := range m.Fields {
			attrName := converter.UpperFirstCharacter(f.Name)
			res = append(res, converter.BindToGo("req.Msg", fmt.Sprintf("%s.%s", in, attrName), attrName, f.Type, false)...)
		}
	} else {
		for i, n := range s.InputNames {
			res = append(res, converter.BindToGo("req.Msg", n, converter.UpperFirstCharacter(n), s.InputTypes[i], true)...)
		}
	}

	return res
}

func OutputGrpc(s *metadata.Service) []string {
	name := converter.UpperFirstCharacter(s.Name)
	res := make([]string, 0)
	if s.HasArrayOutput() {
		res = append(res, fmt.Sprintf("res := new(pb.%sResponse)", name))
		res = append(res, "for _, r := range result {")
		res = append(res, fmt.Sprintf("res.List = append(res.List, to%s(r))", converter.CanonicalName(s.Output)))
		res = append(res, "}")
		res = append(res, "return connect.NewResponse(res), nil")
		return res
	}

	if s.HasCustomOutput() {
		res = append(res, fmt.Sprintf("return connect.NewResponse(&pb.%sResponse{%s: to%s(result)}), nil", name, converter.CamelCaseProto(converter.CanonicalName(s.Output)), converter.CanonicalName(s.Output)))
		return res
	}
	if s.EmptyOutput() {
		res = append(res, fmt.Sprintf("return connect.NewResponse(&pb.%sResponse{}), nil", name))
	} else {
		if s.Output == "sql.Result" {
			res = append(res, fmt.Sprintf("return connect.NewResponse(&pb.%sResponse{Value: toExecResult(result)}), nil", name))
			return res
		}
		res = append(res, fmt.Sprintf("return connect.NewResponse(&pb.%sResponse{Value: result}), nil", name))
	}

	return res
}
