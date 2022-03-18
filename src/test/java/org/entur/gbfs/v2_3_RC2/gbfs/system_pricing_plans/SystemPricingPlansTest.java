package org.entur.gbfs.v2_3_RC2.gbfs.system_pricing_plans;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.entur.gbfs.v2_3_RC2.system_pricing_plans.GBFSSystemPricingPlans;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemPricingPlansTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("system_pricing_plans.json");
        objectMapper.readValue(resource, GBFSSystemPricingPlans.class);
    }
}
