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
Using the pre-build unmarshaller
```go
import(
    "net/http"
    "io"
    v30systemInformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_information"
)
const url = "https://berlin.example.tier-services.io/tier_paris/gbfs/3.0/system-information";
resp, err := http.Get(url)
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
systemInformationData, unmarshalError := v30systemInformation.UnmarshalSystemInformation(body)
if unmarshalError != nil {
    // NOTE: If any type mismatch occurs (ex: number instead of string) it will show up as an error
    fmt.Println("Error unmarshelling:", err)
}
// systemInformationData is now typed as SystemInformation
```
Or you can do it manually
```go
import(
    "net/http"
    "encoding/json"
    "io"
    v30systemInformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_information"
)
const url = "https://berlin.example.tier-services.io/tier_paris/gbfs/3.0/system-information";
resp, errHttp := http.Get(url)
defer resp.Body.Close()
body, errRead := io.ReadAll(resp.Body)

var systemInformationData v30systemInformation.SystemInformation
errMarshel := json.Unmarshal(body, &systemInformationData)
if errMarshel != nil {
    // NOTE: If any type mismatch occurs (ex: number instead of string) it will show up as an error
    fmt.Println("Error unmarshaling JSON:", err)
}

// systemInformationData is now typed as SystemInformation
```

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
