package org.mobilitydata.gbfs;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.TextNode;
import com.sun.codemodel.JPackage;
import com.sun.codemodel.JType;
import org.jsonschema2pojo.Schema;
import org.jsonschema2pojo.rules.ObjectRule;
import org.jsonschema2pojo.rules.Rule;
import org.jsonschema2pojo.rules.RuleFactory;
import org.jsonschema2pojo.util.ParcelableHelper;
import org.jsonschema2pojo.util.ReflectionHelper;

public class GeofencingRuleFactory extends RuleFactory {

    @Override
    public Rule<JPackage, JType> getObjectRule() {
        return new GeofencingObjectRule(this, new ParcelableHelper(), new ReflectionHelper(this));
    }

    static class GeofencingObjectRule extends ObjectRule {

        protected GeofencingObjectRule(
                RuleFactory ruleFactory,
                ParcelableHelper parcelableHelper,
                ReflectionHelper reflectionHelper
        ) {
            super(ruleFactory, parcelableHelper, reflectionHelper);
        }

        @Override
        public JType apply(
                String nodeName,
                JsonNode node,
                JsonNode parent,
                JPackage _package,
                Schema schema
        ) {
            if ((nodeName.equals("geometry") || nodeName.equals("station_area")) && node instanceof ObjectNode) {
                ObjectNode on = (ObjectNode) node;
                on.set("existingJavaType", new TextNode("org.geojson.MultiPolygon"));
            }
            return super.apply(nodeName, node, parent, _package, schema);
        }
    }
}
