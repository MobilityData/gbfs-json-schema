package org.entur.gbfs.v2_2.system_pricing_plans;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemPricingPlansTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v2_X/system_pricing_plans.json");
        objectMapper.readValue(resource, GBFSSystemPricingPlans.class);
    }
}
