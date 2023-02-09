package internal_test

import (
	"go/ast"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"multistring/internal"
)

var _ = Describe("Ast", func() {

	Context("ExtractExportedFuncDecls", func() {

		exportedFuncDecl := &ast.FuncDecl{Name: &ast.Ident{Name: "Exported"}}
		nonFuncDecl := &ast.GenDecl{}
		unexportedFuncDecl := &ast.FuncDecl{Name: &ast.Ident{Name: "unexported"}}

		decls := []ast.Decl{exportedFuncDecl, nonFuncDecl, unexportedFuncDecl}

		It("should retain exported function declarations", func() {
			actual := internal.ExtractExportedFuncDecls(decls)
			Expect(actual).To(ContainElement(exportedFuncDecl))
		})

		It("should filter out non-function declarations", func() {
			actual := internal.ExtractExportedFuncDecls(decls)
			Expect(actual).NotTo(ContainElement(nonFuncDecl))
		})

		It("should filter out non-exported function declarations", func() {
			actual := internal.ExtractExportedFuncDecls(decls)
			Expect(actual).NotTo(ContainElement(unexportedFuncDecl))
		})
	})

	Context("FlattenField", func() {

		It("should not change an anonymous field", func() {
			field := ast.Field{}
			actual := internal.FlattenField(&field)

			Expect(actual).To(ConsistOf(&field))
		})

		It("should not change a field with one name", func() {
			field := ast.Field{Names: []*ast.Ident{{Name: "field"}}}
			actual := internal.FlattenField(&field)

			Expect(actual).To(ConsistOf(&field))
		})

		It("should return one field for each name a multi-name field has", func() {
			field := ast.Field{Names: []*ast.Ident{{Name: "field_1"}, {Name: "field_2"}}}
			actual := internal.FlattenField(&field)

			expected := []*ast.Field{
				{Names: []*ast.Ident{{Name: "field_1"}}},
				{Names: []*ast.Ident{{Name: "field_2"}}},
			}

			Expect(actual).To(Equal(expected))
		})
	})
})
