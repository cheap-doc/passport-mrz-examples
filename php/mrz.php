<?php

declare(strict_types=1);

// MRZ check digits as ICAO Doc 9303 defines them. No dependencies.

/** The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, "<" as 0. */
function mrz_char_value(string $c): int
{
    if (strlen($c) === 1) {
        $o = ord($c);
        if ($o >= 0x30 && $o <= 0x39) { // 0-9
            return $o - 0x30;
        }
        if ($o >= 0x41 && $o <= 0x5A) { // A-Z
            return $o - 0x41 + 10;
        }
        if ($c === '<') {
            return 0;
        }
    }
    throw new InvalidArgumentException('not an MRZ character: ' . var_export($c, true));
}

/** The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10. */
function mrz_check_digit(string $data): int
{
    $weights = [7, 3, 1];
    $sum = 0;
    for ($i = 0, $n = strlen($data); $i < $n; $i++) {
        $sum += mrz_char_value($data[$i]) * $weights[$i % 3];
    }
    return $sum % 10;
}
