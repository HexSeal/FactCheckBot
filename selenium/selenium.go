package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"os"
	"os/exec"
	"strconv"
)

// Snippet from https://github.com/tebeka/selenium/issues/103

// ChromeTest is an example of a basic selenium headless call through Goolge Chrome to make a headless instance 
func ChromeTest() {
	// browserPath := GetBrowserPath("chromium")
	port, err := pickUnusedPort()

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService("chromedriver",
		port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")
	if err != nil {
		panic(err)
	}

	wd.Refresh()

	wd.Get("https://google.com")
	defer service.Stop()
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

// Don't need this, just for testing
func GetBrowserPath(browser string) string {
	if _, err := os.Stat(browser); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}