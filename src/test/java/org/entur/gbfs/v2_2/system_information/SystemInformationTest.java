package org.entur.gbfs.v2_2.system_information;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v2_2.system_pricing_plans.GBFSSystemPricingPlans;
import org.junit.jupiter.api.Test;

class SystemInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_information.json", GBFSSystemPricingPlans.class);
    }
}
