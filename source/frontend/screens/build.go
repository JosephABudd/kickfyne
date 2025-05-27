package screens

import (
	"fmt"

	_accordion_ "github.com/JosephABudd/kickfyne/source/frontend/screens/accordion"
	_simple_ "github.com/JosephABudd/kickfyne/source/frontend/screens/simple"
	_tabbar_ "github.com/JosephABudd/kickfyne/source/frontend/screens/tabbar"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// CreateFramework creates the HelloWorld simple screen so that the app will run.
func CreateFramework(
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.CreateFramework: %w", err)
		}
	}()

	docComment := `Package HelloWorld is a Simple screen package.
A Simple screen package displays only one of it's panels at a time.
It was added when you created this framework.
It is provided as an example.
See the code in the panels folder.
`
	err = BuildSimplePackage(
		"HelloWorld",
		[]string{"Hello", "HelloAgain"},
		docComment,
		importPrefix,
		folderPaths,
	)
	return
}

// BuildSimplePackage builds a type Simple screen package.
func BuildSimplePackage(
	packageName string,
	panelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _simple_.Build(
		panelNames,
		packageName,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}

// BuildSimplePackage builds a type Simple screen package.
func BuildDocTabsPackage(
	packageName string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _tabbar_.Build(
		allPanelNames, localPanelNames, remotePanelNames,
		packageName,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}

// BuildSimplePackage builds a type Simple screen package.
func BuildAccordionPackage(
	packageName string,
	allPanelNames, localPanelNames, remotePanelNames []string,
	packageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {
	return _accordion_.Build(
		allPanelNames, localPanelNames, remotePanelNames,
		packageName,
		packageDoc,
		importPrefix,
		folderPaths,
	)
}
