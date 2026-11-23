package cheap.doc.examples;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Duration;
import java.util.Base64;
import java.util.Map;

/**
 * Send one document image to the doc.cheap API and print the main fields.
 * mvn -q compile exec:java -Dexec.args="path/to/image.jpg"
 */
public final class Scan {
    private static final String API = "https://api.doc.cheap/v1/scans";

    public static void main(String[] args) throws Exception {
        if (args.length != 1) {
            System.err.println("usage: mvn -q compile exec:java -Dexec.args=\"<image.jpg|image.png>\"");
            System.exit(2);
        }
        // The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
        String key = System.getenv().getOrDefault("DOC_CHEAP_API_KEY", "sk_sandbox_public");
        String image = Base64.getEncoder().encodeToString(Files.readAllBytes(Path.of(args[0])));

        ObjectMapper json = new ObjectMapper();
        HttpRequest request = HttpRequest.newBuilder(URI.create(API))
                .timeout(Duration.ofSeconds(60))
                .header("Authorization", "Bearer " + key)
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(json.writeValueAsString(Map.of("image", image))))
                .build();
        HttpResponse<String> response =
                HttpClient.newHttpClient().send(request, HttpResponse.BodyHandlers.ofString());
        JsonNode result = json.readTree(response.body());

        if (response.statusCode() != 200) {
            // Every error body carries a stable code; branch on it, not on the status.
            JsonNode error = result.path("error");
            System.err.printf("%d %s: %s (request %s)%n", response.statusCode(),
                    error.path("code").asText(), error.path("message").asText(),
                    error.path("request_id").asText());
            System.exit(1);
        }

        // A photo that holds no readable document is a 200 too: meta.status says so.
        // path() turns a null group or field into a missing node, printed as "null".
        JsonNode meta = result.path("meta");
        JsonNode document = result.path("document");
        JsonNode holder = result.path("holder");
        System.out.println("status:      " + meta.path("status").asText());
        System.out.println("billed:      " + meta.path("billed").asBoolean());
        System.out.println("kind:        " + text(document, "kind"));
        System.out.println("country:     " + text(document, "country"));
        System.out.println("number:      " + text(document, "number"));
        System.out.println("expiry_date: " + text(document, "expiry_date"));
        System.out.println("surname:     " + text(holder, "surname"));
        System.out.println("given_names: " + text(holder, "given_names"));
        System.out.println("birth_date:  " + text(holder, "birth_date"));
        System.out.println("mrz:         " + result.path("mrz").path("status").asText());
    }

    private static String text(JsonNode group, String name) {
        JsonNode value = group.path(name);
        return value.isMissingNode() || value.isNull() ? "null" : value.asText();
    }
}
