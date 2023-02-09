package internal_test

import (
	"go/ast"
	"go/token"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"multistring/internal"
)

var _ = Describe("Transformation", func() {

	intType := &ast.Ident{Name: "int"}
	stringType := &ast.Ident{Name: "string"}

	Context("Transform", func() {

		It("should transform a function", func() {
			decl := ast.FuncDecl{
				Name: &ast.Ident{Name: "Cut"},
				Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "s"}, {Name: "sep"}}, Type: stringType}}},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "before"}, {Name: "after"}}, Type: stringType},
							{Names: []*ast.Ident{{Name: "found"}}, Type: &ast.Ident{Name: "bool"}},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.IfStmt{
							Init: &ast.AssignStmt{
								Lhs: []ast.Expr{&ast.Ident{Name: "i"}},
								Tok: token.DEFINE,
								Rhs: []ast.Expr{
									&ast.CallExpr{
										Fun:  &ast.Ident{Name: "Index"},
										Args: []ast.Expr{&ast.Ident{Name: "s"}, &ast.Ident{Name: "sep"}},
									},
								},
							},
							Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "i"}, Op: token.GEQ, Y: &ast.BasicLit{Kind: token.INT, Value: "0"}},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ReturnStmt{
										Results: []ast.Expr{
											&ast.SliceExpr{X: &ast.Ident{Name: "s"}, High: &ast.Ident{Name: "i"}},
											&ast.SliceExpr{
												X: &ast.Ident{Name: "s"}, Low: &ast.BinaryExpr{
													X:  &ast.Ident{Name: "i"},
													Op: token.ADD,
													Y: &ast.CallExpr{
														Fun:  &ast.Ident{Name: "len"},
														Args: []ast.Expr{&ast.Ident{Name: "sep"}},
													},
												},
											},
											&ast.Ident{Name: "true"},
										},
									},
								},
							},
						},
						&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "s"}, &ast.BasicLit{Kind: token.STRING, Value: `""`}, &ast.Ident{Name: "false"}}},
					},
				},
			}

			actual := internal.Transform(&decl)
			expected := &ast.FuncDecl{
				Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "w"}}, Type: &ast.Ident{Name: "Wrap"}}}},
				Name: &ast.Ident{Name: "Cut"},
				Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "sep"}}, Type: stringType}}},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "before"}, {Name: "after"}}, Type: &ast.Ident{Name: "Wrap"}},
							{Names: []*ast.Ident{{Name: "found"}}, Type: &ast.ArrayType{Elt: &ast.Ident{Name: "bool"}}},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: "result0"}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "make"},
									Args: []ast.Expr{
										&ast.Ident{Name: "Wrap"},
										&ast.CallExpr{Fun: &ast.Ident{Name: "len"}, Args: []ast.Expr{&ast.Ident{Name: "w"}}},
									},
								},
							},
						},
						&ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: "result1"}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "make"},
									Args: []ast.Expr{
										&ast.Ident{Name: "Wrap"},
										&ast.CallExpr{Fun: &ast.Ident{Name: "len"}, Args: []ast.Expr{&ast.Ident{Name: "w"}}},
									},
								},
							},
						},
						&ast.AssignStmt{
							Lhs: []ast.Expr{&ast.Ident{Name: "result2"}},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.Ident{Name: "make"},
									Args: []ast.Expr{
										&ast.ArrayType{Elt: &ast.Ident{Name: "bool"}},
										&ast.CallExpr{Fun: &ast.Ident{Name: "len"}, Args: []ast.Expr{&ast.Ident{Name: "w"}}},
									},
								},
							},
						},
						&ast.RangeStmt{
							Key:   &ast.Ident{Name: "i"},
							Value: &ast.Ident{Name: "s"},
							Tok:   token.DEFINE,
							X:     &ast.Ident{Name: "w"},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.AssignStmt{
										Lhs: []ast.Expr{
											&ast.IndexExpr{X: &ast.Ident{Name: "result0"}, Index: &ast.Ident{Name: "i"}},
											&ast.IndexExpr{X: &ast.Ident{Name: "result1"}, Index: &ast.Ident{Name: "i"}},
											&ast.IndexExpr{X: &ast.Ident{Name: "result2"}, Index: &ast.Ident{Name: "i"}},
										},
										Rhs: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   &ast.Ident{Name: "strings"},
													Sel: &ast.Ident{Name: "Cut"},
												},
												Args: []ast.Expr{
													&ast.Ident{Name: "s"},
													&ast.Ident{Name: "sep"},
												},
											},
										},
										Tok: token.ASSIGN,
									},
								},
							},
						},
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.Ident{Name: "result0"},
								&ast.Ident{Name: "result1"},
								&ast.Ident{Name: "result2"},
							},
						},
					},
				},
			}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("TransformParams", func() {

		It("should generate new params for a function with one parameter", func() {
			params := []*ast.Field{
				{Names: []*ast.Ident{{Name: "s"}}, Type: stringType},
			}

			actualParams, actualReceiver := internal.TransformParams(params)
			expectedReceiver := &ast.Ident{Name: "s"}

			Expect(actualParams).To(BeEmpty())
			Expect(actualReceiver).To(Equal(expectedReceiver))
		})

		It("should generate new params for a function with two parameters of different types", func() {
			params := []*ast.Field{
				{Names: []*ast.Ident{{Name: "s"}}, Type: stringType},
				{Names: []*ast.Ident{{Name: "i"}}, Type: intType},
			}

			actualParams, actualReceiver := internal.TransformParams(params)
			expectedParams := []*ast.Field{{Names: []*ast.Ident{{Name: "i"}}, Type: intType}}
			expectedReceiver := &ast.Ident{Name: "s"}

			Expect(actualParams).To(Equal(expectedParams))
			Expect(actualReceiver).To(Equal(expectedReceiver))
		})

		It("should generate new params for a function with two parameters of different types where the string is not first", func() {
			params := []*ast.Field{
				{Names: []*ast.Ident{{Name: "i"}}, Type: intType},
				{Names: []*ast.Ident{{Name: "s"}}, Type: stringType},
			}

			actualParams, actualReceiver := internal.TransformParams(params)
			expectedParams := []*ast.Field{{Names: []*ast.Ident{{Name: "i"}}, Type: intType}}
			expectedReceiver := &ast.Ident{Name: "s"}

			Expect(actualParams).To(Equal(expectedParams))
			Expect(actualReceiver).To(Equal(expectedReceiver))
		})

		It("should generate new params for a function with two parameters of the same type", func() {
			params := []*ast.Field{
				{Names: []*ast.Ident{{Name: "s1"}, {Name: "s2"}}, Type: stringType},
			}

			actualParams, actualReceiver := internal.TransformParams(params)
			expectedParams := []*ast.Field{{Names: []*ast.Ident{{Name: "s2"}}, Type: stringType}}
			expectedReceiver := &ast.Ident{Name: "s1"}

			Expect(actualParams).To(Equal(expectedParams))
			Expect(actualReceiver).To(Equal(expectedReceiver))
		})
	})

	Context("TransformResult", func() {

		It("should transform a result of type string to Wrap", func() {
			fields := ast.Field{Type: stringType}

			actual := internal.TransformResult(&fields)
			expected := &ast.Field{Type: &ast.Ident{Name: "Wrap"}}

			Expect(actual).To(Equal(expected))
		})

		It("should transform a result of type []string to []Wrap", func() {
			fields := ast.Field{Type: &ast.ArrayType{Elt: stringType}}

			actual := internal.TransformResult(&fields)
			expected := &ast.Field{Type: &ast.ArrayType{Elt: &ast.Ident{Name: "Wrap"}}}

			Expect(actual).To(Equal(expected))
		})

		It("should transform a result of an arbitrary type T to []T", func() {
			field := ast.Field{Type: intType}

			actual := internal.TransformResult(&field)
			expected := &ast.Field{Type: &ast.ArrayType{Elt: intType}}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("MakeAssignments", func() {

		It("should generate an assignment statement for a function with a single return value", func() {
			fields := []*ast.Field{
				{Type: stringType},
			}

			actual := internal.MakeAssignments(fields)
			expected := []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result0"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{stringType, internal.WrapLength},
						},
					},
				},
			}

			Expect(actual).To(Equal(expected))
		})

		It("should generate an assignment statement for a function with multiple return values", func() {
			fields := []*ast.Field{
				{Type: stringType},
				{Type: stringType},
				{Type: intType},
			}

			actual := internal.MakeAssignments(fields)
			expected := []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result0"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{stringType, internal.WrapLength},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result1"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{stringType, internal.WrapLength},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result2"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{intType, internal.WrapLength},
						},
					},
				},
			}

			Expect(actual).To(Equal(expected))
		})

		It("should generate an assignment statement for a function with multiple return values sharing a type", func() {
			fields := []*ast.Field{
				{Type: intType},
				{Names: []*ast.Ident{{Name: "a"}, {Name: "b"}}, Type: stringType},
			}

			actual := internal.MakeAssignments(fields)
			expected := []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result0"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{intType, internal.WrapLength},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result1"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{stringType, internal.WrapLength},
						},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "result2"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.Ident{Name: "make"}, Args: []ast.Expr{stringType, internal.WrapLength},
						},
					},
				},
			}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("MakeRange", func() {

		It("should generate a range statement for multiple results", func() {
			funcName := &ast.Ident{Name: "string_func"}
			receiver := &ast.Ident{Name: "receiver"}
			params := []*ast.Field{
				{Names: []*ast.Ident{{Name: "s1"}}},
				{Names: []*ast.Ident{{Name: "s2"}}},
			}
			actual := internal.MakeRange(funcName, receiver, params, 2)
			expected := &ast.RangeStmt{
				Key:   internal.LoopCounter,
				Value: receiver,
				X:     internal.WrapIdentifier,
				Tok:   token.DEFINE,
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.IndexExpr{X: &ast.Ident{Name: "result0"}, Index: internal.LoopCounter},
								&ast.IndexExpr{X: &ast.Ident{Name: "result1"}, Index: internal.LoopCounter},
							},
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun:  &ast.SelectorExpr{X: internal.StringsIdentifier, Sel: funcName},
									Args: []ast.Expr{&ast.Ident{Name: "s1"}, &ast.Ident{Name: "s2"}},
								},
							},
							Tok: token.ASSIGN,
						},
					},
				},
			}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("MakeReturn", func() {

		It("should make a return statement for a single result", func() {
			actual := internal.MakeReturn(1)
			expected := &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "result0"}}}

			Expect(actual).To(Equal(expected))
		})

		It("should make a return statement for multiple results", func() {
			actual := internal.MakeReturn(3)
			expected := &ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.Ident{Name: "result0"},
					&ast.Ident{Name: "result1"},
					&ast.Ident{Name: "result2"},
				},
			}

			Expect(actual).To(Equal(expected))
		})
	})

	Context("MakeResultNames", func() {

		It("should make a slice of result names with a given length", func() {
			actual := internal.MakeResultNames(3)
			expected := []string{"result0", "result1", "result2"}

			Expect(actual).To(Equal(expected))
		})
	})
})
