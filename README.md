# apixu-go

Go package for [Apixu Weather API](https://www.apixu.com/api.aspx)

## Install

go get github.com/andreiavrammsd/apixu-go

## Usage

See the [examples](./examples) and run them with: APIXUKEY=yourapikey go run examples/_filename_.go

## Full documentation

https://www.apixu.com/doc/


## Error handling

For more details of an API method error, assert it to the [apixu.Error](./error.go) type.

## Methods implementation status

| Method | Implemented
| :-   | :-
| Current weather | yes
| Forecast | yes
| Search or Autocomplete | yes
| History | yes
| Conditions list | yes

## Extending the package

See the [extend example](./examples/extend.go).

## Testing and QA tools for development

See [Makefile](./Makefile) file.
