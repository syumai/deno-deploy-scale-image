# deno-deploy-scale-image

* image scaling app built for Deno Deploy using Go WebAssembly.
  - this app fetches image from [syumai's GitHub Repo](https://github.com/syumai/images).

## Requirements

* Go 1.16
* Deno
* [deployctl](https://deno.com/deploy/docs/deployctl)

## Usage

### Run app

```
make run
```

### Scale image

show scaled image on browser

`http://0.0.0.0:8080/image?path=landscape.jpg&width=200`

available images

* syumai.png
* landscape.jpg
* ramen.jpg

## Status

* Runs locally
* Not working on Deno Deploy...

## Author

syumai

## License

MIT

