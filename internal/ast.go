package internal

import (
	"go/ast"
)

func ExtractExportedFuncDecls(decls []ast.Decl) []*ast.FuncDecl {
	var funcDecls []*ast.FuncDecl
	for _, decl := range decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || !funcDecl.Name.IsExported() {
			continue
		}
		funcDecls = append(funcDecls, funcDecl)
	}

	return funcDecls
}

func FlattenField(field *ast.Field) []*ast.Field {
	if len(field.Names) <= 1 {
		return []*ast.Field{field}
	}

	result := make([]*ast.Field, len(field.Names))
	for i, name := range field.Names {
		result[i] = &ast.Field{Names: []*ast.Ident{{Name: name.Name}}, Type: field.Type}
	}

	return result
}

func isBuiltinString(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == "string"
}

func isBuiltinStringSlice(expr ast.Expr) bool {
	ident, ok := expr.(*ast.ArrayType)
	return ok && isBuiltinString(ident.Elt)
}
