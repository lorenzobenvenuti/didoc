# didoc (Docker Image DOCumentation)

Quick and dirty tool to attach documentation to Docker images. Supports text, markdown and html rendering.

Documentation is linked to an image declaring two `LABEL`s in the `Dockerfile`:

* `didoc.docs.url`: the url used to retrieve the documentation.
* `didoc.docs.type`: specifies the documentation type. Supported values are `txt`, `md`, `html` (default: `txt`)

Label names can be different; you can specify them using command line options `--url-label` and `--type-label`.

## Usage

```bash
$ ./didoc --help
usage: didoc [<flags>] <image>

Flags:
  --help                        Show context-sensitive help (also try
                                --help-long and --help-man).
  --url-label="didoc.docs.url"  Label used to retrieve documentation url
  --type-label="didoc.docs.type"  
                                Label used to retrieve documentation type
  --version                     Show application version.

Args:
  <image>  Image name or id
```

```bash
$ didoc my-image:latest
My Image
--------

My Image is just a test image. You can use the following environment variables:

* ENV_VAR_1: sets an option
* ENV_VAR_2: sets another option

...
```

## TODO

* Better error handling
* Specify different label names in the arguments
* Colors in output?
