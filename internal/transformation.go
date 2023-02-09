package internal

import (
	"go/ast"
	"go/token"
	"strconv"
)

var (
	LoopCounter       = &ast.Ident{Name: "i"}
	StringsIdentifier = &ast.Ident{Name: "strings"}
	WrapIdentifier    = &ast.Ident{Name: "w"}
	WrapLength        = &ast.CallExpr{
		Fun: &ast.Ident{Name: "len"}, Args: []ast.Expr{
			&ast.Ident{Name: "w"},
		},
	}
)

func Transform(decl *ast.FuncDecl) *ast.FuncDecl {
	newParams, receiver := TransformParams(decl.Type.Params.List)
	newResults := Map(decl.Type.Results.List, TransformResult)
	flatResultCount := len(FlatMap(decl.Type.Results.List, FlattenField))

	var newBody []ast.Stmt
	newBody = append(newBody, MakeAssignments(newResults)...)
	newBody = append(newBody, MakeRange(decl.Name, receiver, decl.Type.Params.List, flatResultCount))
	newBody = append(newBody, MakeReturn(flatResultCount))

	newDecl := *decl
	newDecl.Type.Params = &ast.FieldList{List: newParams}
	newDecl.Recv = &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "Wrap"}}}}
	newDecl.Type.Results = &ast.FieldList{List: newResults}
	newDecl.Body = &ast.BlockStmt{List: newBody}

	return &newDecl
}

func TransformParams(params []*ast.Field) (newParams []*ast.Field, receiverIdentifier *ast.Ident) {
	flattenedParams := FlatMap(params, FlattenField)
	found := false
	for _, param := range flattenedParams {
		if !found && isBuiltinString(param.Type) {
			receiverIdentifier = param.Names[0]
			found = true
		} else {
			newParams = append(newParams, param)
		}
	}

	return newParams, receiverIdentifier
}

func TransformResult(field *ast.Field) *ast.Field {
	result := *field
	switch {
	case isBuiltinString(result.Type):
		result.Type = &ast.Ident{Name: "Wrap"}
	case isBuiltinStringSlice(result.Type):
		result.Type = &ast.ArrayType{Elt: &ast.Ident{Name: "Wrap"}}
	default:
		result.Type = &ast.ArrayType{Elt: result.Type}
	}

	return &result
}

func MakeAssignments(resultFields []*ast.Field) []ast.Stmt {
	flattenedResultFields := FlatMap(resultFields, FlattenField)
	resultNames := MakeResultNames(len(flattenedResultFields))
	assignments := make([]ast.Stmt, len(flattenedResultFields))
	for i, result := range flattenedResultFields {
		resultIdentifier := ast.Ident{Name: resultNames[i]}
		assignment := ast.AssignStmt{
			Lhs: []ast.Expr{&resultIdentifier},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{result.Type, WrapLength}}},
		}

		assignments[i] = &assignment
	}

	return assignments
}

func MakeRange(funcName, receiver *ast.Ident, params []*ast.Field, resultCount int) *ast.RangeStmt {
	var args []ast.Expr
	for _, param := range params {
		for _, name := range param.Names {
			args = append(args, name)
		}
	}

	lhs := make([]ast.Expr, resultCount)
	for i, name := range MakeResultNames(resultCount) {
		lhs[i] = &ast.IndexExpr{X: &ast.Ident{Name: name}, Index: LoopCounter}
	}

	rhs := ast.CallExpr{Fun: &ast.SelectorExpr{X: StringsIdentifier, Sel: funcName}, Args: args}
	return &ast.RangeStmt{
		Key:   LoopCounter,
		Value: receiver,
		X:     WrapIdentifier,
		Tok:   token.DEFINE,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: lhs,
					Rhs: []ast.Expr{&rhs},
					Tok: token.ASSIGN,
				},
			},
		},
	}
}

func MakeReturn(count int) *ast.ReturnStmt {
	results := make([]ast.Expr, count)
	for i, name := range MakeResultNames(count) {
		results[i] = &ast.Ident{Name: name}
	}

	return &ast.ReturnStmt{Results: results}
}

func MakeResultNames(count int) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i] = "result" + strconv.Itoa(i)
	}

	return result
}
