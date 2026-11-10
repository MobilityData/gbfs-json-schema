package org.mobilitydata.gbfs;

import org.junit.jupiter.api.Test;
import org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeedName;

import java.io.IOException;
import java.lang.module.ModuleDescriptor;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import java.util.stream.StreamSupport;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assumptions.assumeTrue;

/**
 * Fails when a package is added (a new GBFS version or feed) without an {@code exports} line in module-info.java.
 * <p>
 * The expected packages are read from the compiled main classes rather than from the module, because surefire
 * patches test classes and test resource directories into the module as well.
 * Only meaningful on the module path, which is how surefire runs the tests when a module descriptor is present.
 */
class ModuleDescriptorTest {

    @Test
    void everyPackageIsExported() throws IOException, URISyntaxException {
        Module module = GBFSFeedName.class.getModule();
        assumeTrue(module.isNamed(), "tests are not running on the module path");
        Path mainClasses = Path.of(GBFSFeedName.class.getProtectionDomain().getCodeSource().getLocation().toURI());
        assumeTrue(Files.isDirectory(mainClasses), "main classes are not in a directory");

        Set<String> expected;
        try (Stream<Path> files = Files.walk(mainClasses)) {
            expected = files
                    .filter(f -> f.getFileName().toString().endsWith(".class"))
                    .filter(f -> !f.getFileName().toString().equals("module-info.class"))
                    .map(f -> StreamSupport.stream(mainClasses.relativize(f.getParent()).spliterator(), false)
                            .map(Path::toString)
                            .collect(Collectors.joining(".")))
                    .collect(Collectors.toSet());
        }
        Set<String> exported = module.getDescriptor().exports().stream()
                .map(ModuleDescriptor.Exports::source)
                .collect(Collectors.toSet());

        assertThat(expected).isNotEmpty();
        assertThat(exported).isEqualTo(expected);
    }
}
