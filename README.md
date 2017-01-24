# didoc (Docker Image DOCumentation)

Quick and dirty tool to attach documentation to Docker images. Supports text, markdown and html rendering.
It uses:

* [Blackfriday markdown processor](https://github.com/russross/blackfriday)
* [Bluemonday HTML sanitizer](https://github.com/microcosm-cc/bluemonday)
* [html2text](https://github.com/jaytaylor/html2text)

## Usage

```bash
$ didoc <my-image-id>
My Image
--------

My Image is just a test image. You can use the following environment variables:

* ENV_VAR_1: sets an option
* ENV_VAR_2: sets another option

...
```

## Image configuration

In your `Dockerfile` you can use the `LABEL` directive to define two labels:

* `dido.docs.url`: the url used to retrieve the documentation.
* `dido.docs.type`: specifies the documentation type. Supported values are `txt`, `md`, `html` (default: `txt`)


## TODO

* Better arguments handling
* Better error handling
* Specify different label names in the arguments
* Colors in output?
