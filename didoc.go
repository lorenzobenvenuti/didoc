package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	urlLabel  = kingpin.Flag("url-label", "Label used to retrieve documentation url").Default("didoc.docs.url").String()
	typeLabel = kingpin.Flag("type-label", "Label used to retrieve documentation type").Default("didoc.docs.type").String()
	imageID   = kingpin.Arg("image", "Image name or id").Required().String()
)

func main() {

	kingpin.Version("0.1")
	kingpin.Parse()

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	image, _, err := cli.ImageInspectWithRaw(context.Background(), *imageID)
	if err != nil {
		panic(err)
	}

	var docType = TEXT
	docTypeStr, ok := image.ContainerConfig.Labels[*typeLabel]
	if ok {
		docType, err = GetDocType(docTypeStr)
		if err != nil {
			panic(err)
		}
	}

	url, ok := image.ContainerConfig.Labels[*urlLabel]
	if !ok {
		panic(errors.New("Cannot read doc url"))
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

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
