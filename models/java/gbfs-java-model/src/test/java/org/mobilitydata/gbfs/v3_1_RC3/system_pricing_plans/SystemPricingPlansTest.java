package org.mobilitydata.gbfs.v3_1_RC3.system_pricing_plans;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemPricingPlansTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_1_RC3/system_pricing_plans.json", GBFSSystemPricingPlans.class);
    }
}
