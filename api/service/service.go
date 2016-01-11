package service

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"sourcegraph.com/sourcegraph/go-selenium"
)

type Response struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type Query struct {
	Url      string
	Selector string
}

func ProcessQuery(query Query) (Response, error) {
	var webDriver selenium.WebDriver
	var err error
	var response Response

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	if webDriver, err = selenium.NewRemote(caps, "http://selenium:4444/wd/hub"); err != nil {
		return Response{}, errors.New("Failed to open session")
	}
	defer webDriver.Quit()

	driverWindow, err := webDriver.CurrentWindowHandle()
	if err != nil {
		return Response{}, errors.New("Failed to get driver window")
	}

	err = webDriver.ResizeWindow(driverWindow, selenium.Size{Height: 3000, Width: 1024})
	if err != nil {
		return Response{}, errors.New("Failed change driver window size")
	}

	err = webDriver.Get(query.Url)
	if err != nil {
		return Response{}, errors.New("Failed to load page")
	}

	sizeInterface, err := webDriver.ExecuteScript("return document.querySelector('"+query.Selector+"').getBoundingClientRect();", nil)
	var size selenium.Size
	if err == nil {
		sizeMap, _ := sizeInterface.(map[string]interface{})
		height, _ := sizeMap["height"].(float64)
		width, _ := sizeMap["width"].(float64)
		size = selenium.Size{Height: height, Width: width}
	} else {
		return Response{}, errors.New("Failed to get selected element size")
	}

	_, err = webDriver.ExecuteScript("document.body.innerHTML=document.querySelector('"+query.Selector+"').outerHTML;", nil)
	if err != nil {
		return Response{}, errors.New("Failed to isolate element")
	}

	err = webDriver.ResizeWindow(driverWindow, size)
	if err != nil {
		return Response{}, errors.New("Failed to change driver window size")
	}

	if image, err := webDriver.Screenshot(); err == nil {
		ioutil.WriteFile("image.jpg", image, 0644)
		response.Image = base64.StdEncoding.EncodeToString([]byte(image))
	} else {
		return Response{}, errors.New("Failed to capture screen shot")
	}

	var elem selenium.WebElement
	elem, err = webDriver.FindElement(selenium.ByCSSSelector, "body")
	if err != nil {
		return Response{}, errors.New("Failed to find element")
	}

	if response.Text, err = elem.Text(); err != nil {
		return Response{}, errors.New("Failed to get text of element")
	}

	contentInterface, err := webDriver.ExecuteScript("return document.querySelector('body').innerHTML;", nil)
	if err == nil {
		response.Content, _ = contentInterface.(string)
	} else {
		return Response{}, errors.New("Failed to get selected element size")
	}

	if response.Title, err = webDriver.Title(); err != nil {
		return Response{}, errors.New("Failed to get page title")
	}

	return response, nil
}
