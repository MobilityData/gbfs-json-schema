package org.mobilitydata.gbfs.geojson;

import org.junit.jupiter.api.Test;
import tools.jackson.databind.JsonNode;
import tools.jackson.databind.ObjectMapper;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

class MultiPolygonTest {
    private static final ObjectMapper MAPPER = new ObjectMapper();

    private static final String JSON = """
            {"type":"MultiPolygon","coordinates":[[[[-122.5,45.5],[-122.6,45.6],[-122.7,45.7,12.5],[-122.5,45.5]]]]}""";

    @Test
    void readsPositionsWithAndWithoutAltitude() {
        MultiPolygon polygon = MAPPER.readValue(JSON, MultiPolygon.class);

        LngLatAlt first = polygon.getCoordinates().get(0).get(0).get(0);
        assertThat(first.getLongitude()).isEqualTo(-122.5);
        assertThat(first.getLatitude()).isEqualTo(45.5);
        assertThat(first.getAltitude()).isNull();
        assertThat(polygon.getCoordinates().get(0).get(0).get(2).getAltitude()).isEqualTo(12.5);
    }

    @Test
    void ignoresPositionElementsAfterTheThird() {
        MultiPolygon polygon = MAPPER.readValue(
                "{\"type\":\"MultiPolygon\",\"coordinates\":[[[[1,2,3,4,5]]]]}", MultiPolygon.class);

        assertThat(polygon.getCoordinates().get(0).get(0).get(0)).isEqualTo(new LngLatAlt(1, 2, 3.0));
    }

    @Test
    void ignoresUnknownProperties() {
        MultiPolygon polygon = MAPPER.readerFor(MultiPolygon.class)
                .with(tools.jackson.databind.DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES)
                .readValue("{\"type\":\"MultiPolygon\",\"bbox\":[0,0,1,1],\"coordinates\":[]}");

        assertThat(polygon.getCoordinates()).isEmpty();
    }

    @Test
    void roundTripsToEquivalentJson() {
        MultiPolygon polygon = MAPPER.readValue(JSON, MultiPolygon.class);

        JsonNode written = MAPPER.readTree(MAPPER.writeValueAsString(polygon));

        assertThat(written).isEqualTo(MAPPER.readTree(JSON));
        assertThat(MAPPER.readValue(MAPPER.writeValueAsString(polygon), MultiPolygon.class)).isEqualTo(polygon);
    }

    @Test
    void rejectsOtherGeometryTypes() {
        assertThatThrownBy(() -> MAPPER.readValue("{\"type\":\"Polygon\",\"coordinates\":[]}", MultiPolygon.class))
                .hasMessageContaining("MultiPolygon");
    }

    @Test
    void rejectsMalformedPositions() {
        for (String position : new String[] {"[1]", "[1,\"a\"]", "{\"lng\":1,\"lat\":2}", "1"}) {
            String json = "{\"type\":\"MultiPolygon\",\"coordinates\":[[[" + position + "]]]}";
            assertThatThrownBy(() -> MAPPER.readValue(json, MultiPolygon.class))
                    .as(position)
                    .isNotNull();
        }
    }
}
