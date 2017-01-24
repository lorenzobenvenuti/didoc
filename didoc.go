package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	image, _, err := cli.ImageInspectWithRaw(context.Background(), os.Args[1])
	if err != nil {
		panic(err)
	}

	var docType DocType = TEXT
	docTypeStr, ok := image.ContainerConfig.Labels["docs.type"]
	if ok {
		var err error
		docType, err = GetDocType(docTypeStr)
		if err != nil {
			panic(err)
		}
	}

	url, ok := image.ContainerConfig.Labels["docs.url"]
	if !ok {
		panic(errors.New("Cannot read doc url"))
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	renderer, err := NewRenderer(docType)
	if err != nil {
		panic(err)
	}

	doc, err := renderer.render(string(body))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", doc)

}
