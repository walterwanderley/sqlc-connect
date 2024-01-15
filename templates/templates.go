package templates

import (
	"embed"
	"html/template"

	"github.com/walterwanderley/sqlc-connect/metadata"
	"github.com/walterwanderley/sqlc-grpc/converter"
)

//go:embed *
var Files embed.FS

var Funcs = template.FuncMap{
	"PascalCase":          converter.ToPascalCase,
	"SnakeCase":           converter.ToSnakeCase,
	"UpperFirstCharacter": converter.UpperFirstCharacter,
	"Input":               metadata.InputGrpc,
	"Output":              metadata.OutputGrpc,
}
