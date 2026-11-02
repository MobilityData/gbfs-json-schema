package org.mobilitydata.gbfs;

import com.fasterxml.jackson.databind.JsonNode;
import com.sun.codemodel.JCodeModel;
import com.sun.codemodel.JDefinedClass;
import com.sun.codemodel.JFieldVar;
import org.jsonschema2pojo.AbstractAnnotator;

/**
 * Makes {@link java.util.Date} properties generated from {@code "format": "date-time"} follow RFC 3339
 * when written and read, independent of how the consumer has configured their mapper.
 * <p>
 * The (de)serializers live in the generated module, so they are referenced by name only.
 */
public class Rfc3339DateTimeAnnotator extends AbstractAnnotator {

    private static final String JSON_SERIALIZE = "tools.jackson.databind.annotation.JsonSerialize";
    private static final String JSON_DESERIALIZE = "tools.jackson.databind.annotation.JsonDeserialize";
    private static final String SERIALIZER = "org.mobilitydata.gbfs.jackson.Rfc3339DateSerializer";
    private static final String DESERIALIZER = "org.mobilitydata.gbfs.jackson.Rfc3339DateDeserializer";

    @Override
    public void dateTimeField(JFieldVar field, JDefinedClass clazz, JsonNode propertyNode) {
        if (!"java.util.Date".equals(field.type().fullName())) {
            return;
        }
        JCodeModel model = clazz.owner();
        field.annotate(model.directClass(JSON_SERIALIZE)).param("using", model.directClass(SERIALIZER));
        field.annotate(model.directClass(JSON_DESERIALIZE)).param("using", model.directClass(DESERIALIZER));
    }
}
