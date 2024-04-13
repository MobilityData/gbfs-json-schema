package org.entur.gbfs.v3_0.system_pricing_plans;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.system_pricing_plans.GBFSSystemPricingPlans;
import org.junit.jupiter.api.Test;

class SystemPricingPlansTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/system_pricing_plans.json", GBFSSystemPricingPlans.class);
    }
}
