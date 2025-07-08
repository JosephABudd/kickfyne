package frontend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_txrx_ "github.com/JosephABudd/kickfyne/source/backend/txrx"
	_frontend_ "github.com/JosephABudd/kickfyne/source/frontend"
	_screens_ "github.com/JosephABudd/kickfyne/source/frontend/screens"
	_message_ "github.com/JosephABudd/kickfyne/source/shared/message"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

// handleScreen passes control to the correct handlers.
func handleScreen(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreen: %w", err)
		}
	}()

	var folderPaths *_utils_.FolderPaths
	if folderPaths, err = _utils_.BuildFolderPaths(pathWD); err != nil {
		return
	}

	if !isBuilt {
		fmt.Println("The app must be initailized before the front end panels can be added or removed.")
		return
	}
	if len(args) == 1 {
		fmt.Println(UsageScreen())
		return
	}
	// args[0] is "screen"
	// args[1] is the verb
	switch args[1] {
	case verbList:
		// args[0] is "screen"
		// args[1] is "list"
		if len(args) != 2 {
			fmt.Println(UsageScreen())
			return
		}
		err = handleScreenList(folderPaths)
	case verbAddSimple:
		// args[0] is "screen"
		// args[1] is "add-simple"
		// args[2] is the «screen-package-name»
		// args[3..] is «panel-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		var isValid bool
		var failureMessage string
		if isValid, failureMessage, err = _utils_.ValidateNewScreenPackageName(args[2], folderPaths); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		fmt.Println("VALID SCREEN NAME")
		// Validate panel names.
		if isValid, failureMessage = _utils_.ValidatePanelNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		fmt.Println("VALID PANEL NAMES")
		screenDocComment := fmt.Sprintf("Package %s is a Simple screen package.\nA simple screen displays only one panel at a time.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddSimple(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		return
	case verbAddAccordion:
		// args[0] is "screen"
		// args[1] is "add-accordion"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]accordion-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		var isValid bool
		var failureMessage string
		if isValid, failureMessage, err = _utils_.ValidateNewScreenPackageName(args[2], folderPaths); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		// Validate panel names.
		if isValid, failureMessage = _utils_.ValidateAccordionItemNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is an Accordion screen package.\nAn accordion screen displays each panel as a titled item that can be opened.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddAccordion(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		return
	case verbAddAppTabs:
		// args[0] is "screen"
		// args[1] is "add-apptabs"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		var isValid bool
		var failureMessage string
		if isValid, failureMessage, err = _utils_.ValidateNewScreenPackageName(args[2], folderPaths); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		// Validate panel names.
		if isValid, failureMessage = _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is an AppTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nThe user is not able to close a tab but your app can.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddAppTabs(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		return
	case verbAddDocTabs:
		// args[0] is "screen"
		// args[1] is "add-doctabs"
		// args[2] is the «screen-package-name»
		// args[3..] is «[*]tab-item-name, ...».
		if len(args) < 4 {
			fmt.Println(UsageScreen())
			return
		}
		// Validate the screen package name.
		var isValid bool
		var failureMessage string
		if isValid, failureMessage, err = _utils_.ValidateNewScreenPackageName(args[2], folderPaths); !isValid || err != nil {
			fmt.Println(failureMessage)
			return
		}
		// Validate panel names.
		if isValid, failureMessage = _utils_.ValidateTabNames(args[3:]); !isValid {
			fmt.Println(failureMessage)
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is a DocTabs screen package.\nA tabbar screen where a tab displays it's panel or another screen.\nKICKFYNE TODO: Complete this package doc commment.", args[2])
		if err = handleScreenAddDocTabs(args[2], args[3:], screenDocComment, importPrefix, folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		return
	case verbRemove:
		// args[0] is "screen"
		// args[1] is "remove"
		// args[2] is the «screen-package-name»
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		if err = handleScreenRemove(args[2], folderPaths); err != nil {
			return
		}
		err = _frontend_.RebuildFrontendGo(importPrefix, folderPaths)
		return
	case subCmdHelp:
		// args[0] is "screen"
		// args[1] is "help"
		fmt.Println(UsageScreen())
	default:
		// args[0] is "screen"
		fmt.Println(UsageScreen())
	}
	return
}

// handleScreenAddSimple handles adding a screen package.
func handleScreenAddSimple(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAdd: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Create the package folder with no panels.
	if err = _screens_.BuildSimplePackage(
		screenPackageName,
		panelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessageScreenAddSimple(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddAccordion handles adding a Accordion screen package.
func handleScreenAddAccordion(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAccordion: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildAccordionPackage(
		screenPackageName,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	// Add the messages, tabs, messenger id storage.
	if err = _message_.AddAccordion(screenPackageName, allPanelNames, localPanelNames, importPrefix, folderPaths); err != nil {
		return
	}
	// Add the back-end message handlers.
	if err = _txrx_.AddAccordion(screenPackageName, importPrefix, folderPaths); err != nil {
		return
	}

	successMessage = successMessageScreenAddAccordion(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddAppTabs handles adding a AppTabs screen package.
func handleScreenAddAppTabs(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddAppTabs: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildAppTabsPackage(
		screenPackageName,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	// Add the messages, tabs, messenger id storage.
	if err = _message_.AddAppTabs(screenPackageName, allPanelNames, localPanelNames, importPrefix, folderPaths); err != nil {
		return
	}
	// Add the back-end message handlers.
	if err = _txrx_.AddAppTab(screenPackageName, importPrefix, folderPaths); err != nil {
		return
	}

	successMessage = successMessageScreenAddAppTabs(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenAddDocTabs handles adding a DocTabs screen package.
func handleScreenAddDocTabs(
	screenPackageName string,
	panelNames []string,
	screenPackageDoc string,
	importPrefix string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAddDocTabs: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Add the screen package.
	allPanelNames, localPanelNames, remotePanelNames := separatePanelNames(panelNames)
	if err = _screens_.BuildDocTabsPackage(
		screenPackageName,
		allPanelNames, localPanelNames, remotePanelNames,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	// Add the messages, tabs, messenger id storage.
	if err = _message_.AddDocTabs(screenPackageName, allPanelNames, localPanelNames, importPrefix, folderPaths); err != nil {
		return
	}
	// Add the back-end message handlers.
	if err = _txrx_.AddDocTab(screenPackageName, importPrefix, folderPaths); err != nil {
		return
	}

	successMessage = successMessageScreenAddDocTabs(screenPackageName, panelNames, folderPaths)
	return
}

// handleScreenRemove handles the removal of a screen package.
func handleScreenRemove(
	screenPackageName string,
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	var builder = strings.Builder{}
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemove: %w", err)
			return
		}
		successMessage = builder.String()
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Validate the screen package name.
	var isValid bool
	var msg string
	if isValid, msg, err = _utils_.ValidateCurrentScreenPackageName(screenPackageName, folderPaths); !isValid || err != nil {
		if !isValid {
			failureMessage = msg
		}
		return
	}
	packageFolderPath := filepath.Join(folderPaths.FrontendScreens, screenPackageName)
	if _, err = os.Stat(packageFolderPath); err != nil {
		if os.IsNotExist(err) {
			// The folder does not exist.
			err = nil
			builder.WriteString(fmt.Sprintf("The screen package %q was previously removed for some reason.", screenPackageName))
		}
		if err != nil {
			return
		}
	}
	if err = os.RemoveAll(packageFolderPath); err != nil {
		return
	}
	// Remove the frontend/screen/«screenPackageName» folder.
	builder.WriteString(fmt.Sprintf("Removed the %[1]s screen's folder at frontend/screens/%[1]s.\n", screenPackageName))
	// Remove the shared/message/«screenPackageName» folder if this is a DocTabs screen.
	sharedMessageFolderPath := _utils_.DocTabSharedMessageFolderPath(screenPackageName, folderPaths)
	if err = os.RemoveAll(sharedMessageFolderPath); err != nil {
		if os.IsNotExist(err) {
			// The folder does not exist.
			err = nil
		}
		if err != nil {
			return
		}
	} else {
		builder.WriteString(fmt.Sprintf("Removed the screen's DocTabs messages at shared/message/%s.\n", screenPackageName))
	}
	// Remove the backend/txrx/«screenPackageName» folder if this is a DocTabs screen.
	backendMessageFolderPath := _utils_.DocTabBackendMessengerFolderPath(screenPackageName, folderPaths)
	if err = os.RemoveAll(backendMessageFolderPath); err != nil {
		if os.IsNotExist(err) {
			// The folder does not exist.
			err = nil
		}
		if err != nil {
			return
		}
	} else {
		builder.WriteString(fmt.Sprintf("Removed the screen's DocTabs back-end message handlers at backend/txrs/%s.\n", screenPackageName))
	}

	return
}

// handleScreenList handles the listing of the screen packages.
func handleScreenList(
	folderPaths *_utils_.FolderPaths,
) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenList: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	// Get the screen names.
	var screenNames []string
	if screenNames, err = _utils_.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	// Get the panel names.
	screenPanelFolderNames := make(map[string][]string)
	for _, screenName := range screenNames {
		screenFolderPath := filepath.Join(folderPaths.FrontendScreens, screenName)
		if screenPanelFolderNames[screenName], err = _utils_.PanelNames(screenFolderPath); err != nil {
			return
		}
	}
	// Display the list.
	fmt.Printf("List of %d screen packages.\n", len(screenNames))
	for i, screenName := range screenNames {
		panelFolderNames := screenPanelFolderNames[screenName]
		// Display the panels
		lengthPanel := len("Panel")
		fmt.Printf("% d. %s:\n", i+1, screenName)
		for j, panelFolderName := range panelFolderNames {
			length := len(panelFolderName) - lengthPanel
			panelName := panelFolderName[0:length]
			contentPath := _utils_.PanelContentFilePath(screenName, panelName, folderPaths)
			statePath := _utils_.PanelStateFilePath(screenName, panelName, folderPaths)
			messengerPath := _utils_.PanelMessengerFilePath(screenName, panelName, folderPaths)
			fmt.Printf("  %d: %s Panel\n", j+1, panelName)
			fmt.Printf("    Content:   %s.\n", _utils_.Clickable(contentPath))
			fmt.Printf("    State:     %s.\n", _utils_.Clickable(statePath))
			fmt.Printf("    Messenger: %s.\n", _utils_.Clickable(messengerPath))
		}
	}

	return
}

func successMessageScreenAddSimple(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the Simple screen package named %q.\n", screenPackageName))
	builder.WriteString("The package's docs and each panel's content and state require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		messengerPath := _utils_.PanelMessengerFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Messenger: %s.\n", _utils_.Clickable(messengerPath)))
	}
	successMessage = builder.String()
	return
}

func successMessageScreenAddAppTabs(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the AppTabs screen package named %q.\n", screenPackageName))
	builder.WriteString("The package's docs and each panel's content and state require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		messengerPath := _utils_.PanelMessengerFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + " Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Messenger: %s.\n", _utils_.Clickable(messengerPath)))
	}

	successMessage = builder.String()
	return
}

func successMessageScreenAddDocTabs(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the DocTabs screen package named %q.\n", screenPackageName))
	builder.WriteString("The package's docs and each panel's content and state require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		messengerPath := _utils_.PanelMessengerFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + " Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Messenger: %s.\n", _utils_.Clickable(messengerPath)))
	}
	// Shared Messages
	messageFolderPath := _utils_.DocTabSharedMessageFolderPath(screenPackageName, folderPaths)
	messageNameFileName := _utils_.DocTabMessageNameFileName(screenPackageName)
	builder.WriteString("Shared messages:\n")
	for messageName, messageFileName := range messageNameFileName {
		messagePath := filepath.Join(messageFolderPath, messageFileName)
		builder.WriteString(fmt.Sprintf("  %s: %s\n", messageName, _utils_.Clickable(messagePath)))
	}
	// Back-end Messengers
	messengerFolderPath := _utils_.DocTabBackendMessengerFolderPath(screenPackageName, folderPaths)
	builder.WriteString("Back-end messengers:\n")
	for messageName, messageFileName := range messageNameFileName {
		messengerPath := filepath.Join(messengerFolderPath, messageFileName)
		builder.WriteString(fmt.Sprintf("  %s: %s\n", messageName, _utils_.Clickable(messengerPath)))
	}

	successMessage = builder.String()
	return
}

func successMessageScreenAddAccordion(
	screenPackageName string,
	localPanelNames []string,
	folderPaths *_utils_.FolderPaths,
) (successMessage string) {
	docFilePath := _utils_.ScreenDocFilePath(screenPackageName, folderPaths)
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Added the Accordion screen package named %q.\n", screenPackageName))
	builder.WriteString("The package's docs and each panel's content and state require editing.\n")
	builder.WriteString("Package docs: " + _utils_.Clickable(docFilePath) + "\n")
	for _, panelName := range localPanelNames {
		contentPath := _utils_.PanelContentFilePath(screenPackageName, panelName, folderPaths)
		statePath := _utils_.PanelStateFilePath(screenPackageName, panelName, folderPaths)
		messengerPath := _utils_.PanelMessengerFilePath(screenPackageName, panelName, folderPaths)
		builder.WriteString(panelName + "Panel\n")
		builder.WriteString(fmt.Sprintf("  Content:   %s.\n", _utils_.Clickable(contentPath)))
		builder.WriteString(fmt.Sprintf("  State:     %s.\n", _utils_.Clickable(statePath)))
		builder.WriteString(fmt.Sprintf("  Messenger: %s.\n", _utils_.Clickable(messengerPath)))
	}
	successMessage = builder.String()
	return
}

func separatePanelNames(panelNames []string) (all []string, local []string, remote []string) {
	all = make([]string, 0, len(panelNames))
	local = make([]string, 0, len(panelNames))
	remote = make([]string, 0, len(panelNames))
	for _, panelName := range panelNames {
		if panelName[0] == '*' {
			fixed := panelName[1:]
			remote = append(remote, fixed)
			all = append(all, fixed)
		} else {
			local = append(local, panelName)
			all = append(all, panelName)
		}
	}
	return
}
