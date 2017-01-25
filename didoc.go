package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/docker/docker/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	urlLabel  = kingpin.Flag("url-label", "Label used to retrieve documentation url").Default("didoc.docs.url").String()
	typeLabel = kingpin.Flag("type-label", "Label used to retrieve documentation type").Default("didoc.docs.type").String()
	imageID   = kingpin.Arg("image", "Image name or id").Required().String()
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	kingpin.Version("0.2.0")
	kingpin.Parse()

	cli, err := client.NewEnvClient()
	handleErr(err)

	image, _, err := cli.ImageInspectWithRaw(context.Background(), *imageID)
	handleErr(err)

	var docType = TEXT
	docTypeStr, ok := image.ContainerConfig.Labels[*typeLabel]
	if ok {
		docType, err = GetDocType(docTypeStr)
		handleErr(err)
	}

	url, ok := image.ContainerConfig.Labels[*urlLabel]
	if !ok {
		handleErr(fmt.Errorf("Cannot read label %s", *urlLabel))
	}

	resp, err := http.Get(url)
	handleErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	renderer, err := NewRenderer(docType)
	handleErr(err)

	doc, err := renderer.render(string(body))

	handleErr(err)

	fmt.Printf("\n%s\n\n", doc)

}
