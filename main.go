package nosetter

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "notsetterlint",
	Doc:      "reports setter to domain model",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func init() {
	Analyzer.Flags.String("pkg", "", "will ensure pkg path with this prefix will err")
}

func run(pass *analysis.Pass) (interface{}, error) {
	pkg := pass.Analyzer.Flags.Lookup("pkg").Value.String()

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// We filter only assignments.
	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.AssignStmt)
		for _, token := range call.Lhs {
			switch sel := token.(type) {
			case *ast.SelectorExpr:
				typ, ok := pass.TypesInfo.Types[sel.X]
				if !ok {
					return
				}

				_, ok = typ.Type.Underlying().(*types.Struct)
				if !ok {
					return
				}

				t := typ.Type

				// If it is a pointer, get the element.
				if ptr, ok := t.(*types.Pointer); ok {
					t = ptr.Elem()
				}

				nTyp, ok := t.(*types.Named)
				if !ok {
					return
				}
				structName := nTyp.Obj().Name()    // Foo
				pkgPath := nTyp.Obj().Pkg().Path() // github.com/alextanhongpin/nosetter/foo

				if strings.HasPrefix(pkgPath, pkg) {
					err := fmt.Errorf("%s.%s", pkgPath, structName)
					pass.Reportf(call.TokPos, "Setter not allowed for struct: %v", err)
				}
			}
		}
	})

	return nil, nil
}
