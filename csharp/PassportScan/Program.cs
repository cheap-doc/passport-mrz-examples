// Send one document image to the doc.cheap API and print the main fields.
// .NET 8, no packages: dotnet run --project PassportScan -- path/to/image.jpg
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text.Json;

const string Api = "https://api.doc.cheap/v1/scans";

if (args.Length != 1)
{
    Console.Error.WriteLine("usage: dotnet run --project PassportScan -- <image.jpg|image.png>");
    return 2;
}

// The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
var key = Environment.GetEnvironmentVariable("DOC_CHEAP_API_KEY") ?? "sk_sandbox_public";
var image = Convert.ToBase64String(await File.ReadAllBytesAsync(args[0]));

using var http = new HttpClient { Timeout = TimeSpan.FromSeconds(60) };
http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", key);
using var response = await http.PostAsJsonAsync(Api, new { image });
using var json = JsonDocument.Parse(await response.Content.ReadAsStringAsync());
var root = json.RootElement;

if (!response.IsSuccessStatusCode)
{
    // Every error body carries a stable code; branch on it, not on the status.
    var error = root.GetProperty("error");
    Console.Error.WriteLine(
        $"{(int)response.StatusCode} {error.GetProperty("code")}: {error.GetProperty("message")} " +
        $"(request {error.GetProperty("request_id")})");
    return 1;
}

// A field that is not known is null; a whole group is null when nothing was read.
static string Field(JsonElement group, string name) =>
    group.ValueKind == JsonValueKind.Object && group.GetProperty(name).ValueKind != JsonValueKind.Null
        ? group.GetProperty(name).ToString()
        : "null";

// A photo that holds no readable document is a 200 too: meta.status says so.
var meta = root.GetProperty("meta");
var document = root.GetProperty("document");
var holder = root.GetProperty("holder");
Console.WriteLine($"status:      {meta.GetProperty("status")}");
Console.WriteLine($"billed:      {meta.GetProperty("billed").GetBoolean()}");
Console.WriteLine($"kind:        {Field(document, "kind")}");
Console.WriteLine($"country:     {Field(document, "country")}");
Console.WriteLine($"number:      {Field(document, "number")}");
Console.WriteLine($"expiry_date: {Field(document, "expiry_date")}");
Console.WriteLine($"surname:     {Field(holder, "surname")}");
Console.WriteLine($"given_names: {Field(holder, "given_names")}");
Console.WriteLine($"birth_date:  {Field(holder, "birth_date")}");
Console.WriteLine($"mrz:         {root.GetProperty("mrz").GetProperty("status")}");
return 0;
