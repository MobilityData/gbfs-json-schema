package org.entur.gbfs.v2_2.system_pricing_plans;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class SystemPricingPlansTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/system_pricing_plans.json", GBFSSystemPricingPlans.class);
    }
}
