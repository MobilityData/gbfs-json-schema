# Golang GBFS Language Bindings

Golang types for parsing and working with General Bikeshare Feed Specification (GBFS) data, ensuring type safety and code consistency in Golang projects.

## Add the Dependency

To use the `gbfs-language-bindings` structs in your own project, you need to
first install this library with: 

```
go get github.com/MobilityData/gbfs-json-schema/models/golang 
```

## Versions
Currently only version 3.0 of GBFS is supported

## Example Code
See the [example code](https://github.com/MobilityData/gbfs-json-schema/models/golang/examples/unmarshal.go)

## Important to Note
When unmarshalling data Golang sets defaults if the required value does not exist
```
int, int8, int16, int32, int64: 0
uint, uint8 (byte), uint16, uint32, uint64: 0
float32, float64: 0.0
string: "" (empty string)
bool: false
Pointers, slices, maps, channels, functions, and interfaces: nil
```
