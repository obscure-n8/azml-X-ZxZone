package ext

import (
	"strings"
	"testing"

	"azml-go/types"
)

func TestBuildStatusMessageUsesLeechStatusLayout(t *testing.T) {
	task := NewUnifiedTask("8861026839_ccb0ca", "Resident.Evil.2026.1080p.AMZN.WEB-DL.Hindi-English.DDP5.1.H.264-KyoGo-4kHdHub.mkv", 10*1024*1024*1024, -100123, 8861026839, 1, nil)
	task.SetStatusDisplay("xyzasa44", "User", "#leech")
	task.SetProgress(3*1024*1024*1024, 10*1024*1024*1024, 69.7*1024*1024, "30s")
	taskManager := &TaskManager{tasks: map[string]types.DownloadTask{task.GID(): task}}

	message, buttons := taskManager.BuildStatusMessage(1)
	for _, expected := range []string{
		"1. Task Name:",
		"[⬤⬤⬤◔○○○○○○○○]",
		"Status:</b> <i>Downloading...</i>",
		"Done:</b> 3.00 GB",
		"Eta:</b> 30s",
		"Engine:</b> Aria2c",
		"User:</b> @xyzasa44",
		"In Mode:</b> #leech",
		"Out Mode:</b> telegram",
		"/cancel_8861026839_ccb0ca",
		"⌬ <b>ʙᴏᴛ sᴛᴀᴛs</b>\n<blockquote>╭ <b>Active Tasks:</b> 1",
		"<b>Ram:</b>",
		"<b>Cpu:</b>",
		"<b>Uptime:</b>",
		"<b>DL:</b>",
		"<b>Ul:</b>",
		"</blockquote>",
	} {
		if !strings.Contains(message, expected) {
			t.Errorf("status message does not contain %q:\n%s", expected, message)
		}
	}
	if buttons == nil {
		t.Fatal("expected a refresh button")
	}
}

func TestBuildStatusMessageEscapesTaskName(t *testing.T) {
	task := NewUnifiedTask("safe-name", "video<&>.mkv", 1024, -100123, 12, 1, nil)
	task.SetStatusDisplay("", "User & Name", "#leech")
	taskManager := &TaskManager{tasks: map[string]types.DownloadTask{task.GID(): task}}

	message, _ := taskManager.BuildStatusMessage(1)
	if !strings.Contains(message, "video-&amp;-.mkv") {
		t.Fatalf("expected escaped task name in status message, got:\n%s", message)
	}
	if !strings.Contains(message, "User &amp; Name") {
		t.Fatalf("expected escaped sender name in status message, got:\n%s", message)
	}
}
