# gbfs-java-model

![Maven Central Version](https://img.shields.io/maven-central/v/org.mobilitydata/gbfs-java-model.svg)

Generates Java model from GBFS json schema using jsonschema2pojo with jackson2 annotations.

In order to avoid naming conflicts in generated classes, each feed gets its own package, and all
classes are prefixed with "GBFS".

Packages include GBFS versions to accommodate for future releases which may include older versions,
since some consumers may need to handle multiple non-compatible versions in the same application 
(i.e. you should be able to parse v2.1 feeds with the v2.2 package, but v3.0 package would not
be able to parse v2.2 feeds).

## Add the Dependency 

To use the `gbfs-java-model` classes in your own project, you need to add
an appropriate dependency.  We publish our module to the [Maven Central Repository](http://search.maven.org/)
so that it can be easily referenced by Java build tools like Maven, Ivy, and Gradle.

For [Maven](http://maven.apache.org/), add the following to your `pom.xml`
dependencies section:

```xml
<dependency>
  <groupId>org.mobilitydata</groupId>
  <artifactId>gbfs-java-model</artifactId>
  <version>1.0.5</version>
</dependency>
```

For [Gradle](https://www.gradle.org/), add the following to your `build.gradle`
dependecies section:

```
implementation group: 'org.mobilitydata', name: 'gbfs-java-model', version: '1.0.0'
```

Make sure the Maven central repository is referenced by your project.

## Example Code
```java
import org.mobilitydata.gbfs.v3_0.vehicle_status.GBFSVehicleStatus;

// Assume you use Jackson's ObjectMapper for serializing / deserializing,
// this is just one of many options
ObjectMapper objectMapper = new ObjectMapper();

// Assume some InputStream containing a serialized instance of vehicle_status.json
// Deserialize into corresponding java class
GBFSVehicleStatus vehicleStatus = objectMapper.readValue(inputStream, GBFSVehicleStatus.class);

// It can be serialized in a similar fashion,
objectMapper.writeValue(outputStream, vehicleStatus);
```

### Special cases: in v2.x

Note: This does not apply to v3.0

#### gbfs.json

There is no support in jsonschema2pojo to handle "patternProperties" sanely, it will just result
in a class without any properties. The `gbfs.json` feed uses "patternProperties" for the `data`
property to define an object per language, where the language code is the name of the property.

The generated Gbfs class is therefore extended by 
[gbfs-java-model/src/main/java/org/mobilitydata/gbfs/v2_2/gbfs/GBFS.java](gbfs-java-model/src/main/java/org/mobilitydata/gbfs/v2_2/gbfs/GBFS.java) 
to override the `data` property with a Map implementation. This extended class should be used
when unmarshalling the `gbfs.json` feed.

#### station_information.json

In v2.x the properties `vehicle_capacity` and `vehicle_type_capacity` in `station_information`, are objects
whose keys are vehicle type ids. The generated classes for these properties are `GBFSVehicleCapacity` and
`GBFSVehicleTypeCapacity` respectively. To mapped key-value pairs can be accessed via the property `additionalProperties`
of these generated classes.

## Maven central
This project is available in the central maven repository.
See https://search.maven.org/search?q=g:org.mobilitydata

## Contributions
We extend our gratitude to [Entur](https://www.entur.org/) for generously providing us with this repository. Their support and contribution have been instrumental in advancing our project and fostering collaboration within the community. Thank you, Entur, for your dedication to open source and for empowering developers worldwide.
