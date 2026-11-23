<?php

declare(strict_types=1);

// A dependency-free test: php mrz_test.php exits 1 on the first failure.

require __DIR__ . '/mrz.php';

function expect_same(int $expected, int $actual, string $name): void
{
    if ($expected !== $actual) {
        fwrite(STDERR, "FAIL {$name}: expected {$expected}, got {$actual}\n");
        exit(1);
    }
    echo "ok   {$name}\n";
}

// Fields of the ICAO Doc 9303 specimen passport line.
$line = 'L898902C36UTO7408122F1204159ZE184226B<<<<<10';

expect_same(6, mrz_check_digit('L898902C3'), 'document number');
expect_same(2, mrz_check_digit('740812'), 'birth date');
expect_same(9, mrz_check_digit('120415'), 'expiry date');
expect_same(1, mrz_check_digit('ZE184226B<<<<<'), 'personal number');
expect_same(0, mrz_check_digit(substr($line, 0, 10) . substr($line, 13, 7) . substr($line, 21, 22)), 'composite');
expect_same(0, mrz_check_digit('<<<<<<'), 'fillers count as zero');

if (mrz_check_digit('L898902C4') === 6) {
    fwrite(STDERR, "FAIL a misread document number still gave 6\n");
    exit(1);
}
echo "ok   one misread character changes the digit\n";

try {
    mrz_check_digit('l898902c3');
    fwrite(STDERR, "FAIL lower-case input was accepted\n");
    exit(1);
} catch (InvalidArgumentException $e) {
    echo "ok   refuses a character outside the MRZ alphabet\n";
}
