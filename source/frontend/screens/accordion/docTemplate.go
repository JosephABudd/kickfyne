package accordion

import (
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

type docTemplateData struct {
	PackageName string
	PackageDoc  string
	Funcs       _utils_.Funcs
}

const (
	docFileName = _utils_.DocFileName

	docTemplate = `{{ call .Funcs.Comment .PackageDoc }}
package {{ .PackageName }}
`
)
