package desktop

import (
	"os/exec"
)

func (manager *DesktopManagerCtx) HandleFileChooserDialog(uri string) error {
	mu.Lock()
	defer mu.Unlock()

	// TOOD: Use native API.
	cmd := exec.Command(
		"xdotool",
			"search", "--name", "Open File", "windowfocus",
			"sleep", "0.2",
			"key", "--clearmodifiers", "ctrl+l",
			"type", "--args", "1", uri + "//",
			"key", "--clearmodifiers", "Return",
			"sleep", "1",
			"key", "--clearmodifiers", "Down",
			"key", "--clearmodifiers", "ctrl+a",
			"key", "--clearmodifiers", "Return",
			"sleep", "0.3",
	)

	_, err := cmd.Output()
	return err
}

func (manager *DesktopManagerCtx) CloseFileChooserDialog() {
	for i := 0; i < 5; i++ {
		if !manager.IsFileChooserDialogOpened() {
			return
		}

		// TOOD: Use native API.
		mu.Lock()
		exec.Command(
			"xdotool",
				"search", "--name", "Open File", "windowfocus",
				"sleep", "0.2",
				"key", "--clearmodifiers", "alt+F4",
		).Output()
		mu.Unlock()
	}
}

func (manager *DesktopManagerCtx) IsFileChooserDialogOpened() bool {
	mu.Lock()
	defer mu.Unlock()

	// TOOD: Use native API.
	cmd := exec.Command(
		"xdotool",
			"search", "--name", "Open File",
	)

	_, err := cmd.Output()
	return err == nil
}