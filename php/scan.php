<?php

declare(strict_types=1);

// Send one document image to the doc.cheap API and print the main fields.
// PHP 8 with the curl and json extensions: php scan.php path/to/image.jpg

const API = 'https://api.doc.cheap/v1/scans';

if ($argc !== 2) {
    fwrite(STDERR, "usage: php scan.php <image.jpg|image.png>\n");
    exit(2);
}

// The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
$key = getenv('DOC_CHEAP_API_KEY') ?: 'sk_sandbox_public';
$raw = file_get_contents($argv[1]);
if ($raw === false) {
    exit(1);
}
$image = base64_encode($raw);

$ch = curl_init(API);
curl_setopt_array($ch, [
    CURLOPT_POST => true,
    CURLOPT_HTTPHEADER => ["Authorization: Bearer {$key}", 'Content-Type: application/json'],
    CURLOPT_POSTFIELDS => json_encode(['image' => $image]),
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_TIMEOUT => 60,
]);
$body = curl_exec($ch);
if ($body === false) {
    fwrite(STDERR, curl_error($ch) . "\n");
    exit(1);
}
$status = curl_getinfo($ch, CURLINFO_RESPONSE_CODE);
$result = json_decode($body, true);

if ($status !== 200) {
    // Every error body carries a stable code; branch on it, not on the status.
    $error = $result['error'];
    fwrite(STDERR, "{$status} {$error['code']}: {$error['message']} (request {$error['request_id']})\n");
    exit(1);
}

// A photo that holds no readable document is a 200 too: meta.status says so.
$document = $result['document'] ?? [];
$holder = $result['holder'] ?? [];
$show = static fn ($value): string => json_encode($value);
echo 'status:      ', $result['meta']['status'], "\n";
echo 'billed:      ', $show($result['meta']['billed']), "\n";
echo 'kind:        ', $show($document['kind'] ?? null), "\n";
echo 'country:     ', $show($document['country'] ?? null), "\n";
echo 'number:      ', $show($document['number'] ?? null), "\n";
echo 'expiry_date: ', $show($document['expiry_date'] ?? null), "\n";
echo 'surname:     ', $show($holder['surname'] ?? null), "\n";
echo 'given_names: ', $show($holder['given_names'] ?? null), "\n";
echo 'birth_date:  ', $show($holder['birth_date'] ?? null), "\n";
echo 'mrz:         ', $result['mrz']['status'], "\n";
