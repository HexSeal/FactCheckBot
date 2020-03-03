package factCheck

import (
	"fmt"
	"os"
	"time"
	"strings"
	"github.com/tebeka/selenium"
)

// Start a Selenium WebDriver server instance (if one is not already
// running).
const (
    // These paths will be different on your system.
    seleniumPath    = "/selenium-server-standalone-3.4.jar"
    geckoDriverPath = "/geckodriver-v0.18.0-linux64"
    port            = 8080
)

// SnopesCheck runs selenium on snopes.com and returns the answer to the query // CURRENTLY TESTING WITH GOPLAYGROUND
func SnopesCheck() {
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver("/maxfinn/dev/dependencies/chromedriver"), 		// Specify the path to chromedriver. MIGHT BE CAUSING EOF ERROR
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to snopes
	if err := wd.Get("snopes.com"); err != nil {
		panic(err)
	}

	// Get a reference to the search bar
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#header-search")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate "Search Snopes.com" already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter the user's query, soon to be from Slack
	err = elem.SendKeys("Did Trump Fire the US Pandemic Response Team?")
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	// This simply returns if the query is true or false. The selector for the explanation itself is:
	// body > div.theme-content > div > div > main > article > div.content-wrapper.card > div.content
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "body > div.theme-content > div > div > main > article > div.rating-wrapper.card > div > div > div > h5")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}


	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))
}
