# PDF Service
This is a simple HTTP service that will convert HTML documents to PDF using Chrome
_headless_ mode.

Currently this solves an issue of printing HTML documents being a complete pain.

## Requirements
* Chromium (because of open source licence)
* Go for building

## Building
Just run
```shell
go build
```

## Usage
The port flag is optional, default is: __8765__.
```shell
./pdf-service --port=8765 &
```

To check that the service is running just do a simple GET request to /.

To upload a HTML document via `POST /upload`

### Todo
* Installation guide
* Unit tests
* Generating PDFs via URLs which will be useful for another usecase
* Garbage collection (removal of temporary files)
* Improve this documentation a little bit.
