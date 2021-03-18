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
./pdf-service &
```

### Flags
| Flag    | Optional | Default         | Description                                |
| ------- | :------: | --------------- | ------------------------------------------ |
| port    | [ x ]      | 8765            | The port to run the server on              |
| tmp-dir | [ x ]      | /tmp/pdf-server | The directory to temporary store the files |

To check that the service is running just do a simple GET request to /.

To upload a HTML document via `POST /upload`

## API Routes
These are the API routes that are currently accepted. 

| Method  | Path    | Response | Params               | Descriptiuon |
| ------- | ------- | -------- | -------------------- | ------------ |
| GET     | /       | JSON     |                      | Just shows that the server is running. Useful for pinging to check its alive. |
| POST    | /upload | BINARY   | file=_document.html_ | The actual file upload. It will respond with a generated PDF. |

### Todo
* Installation guide
* Unit tests
* Generating PDFs via URLs which will be useful for another usecase
* ~~Garbage collection (removal of temporary files)~~
