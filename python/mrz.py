"""MRZ check digits as ICAO Doc 9303 defines them. No dependencies."""

WEIGHTS = (7, 3, 1)


def char_value(c):
    """The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, '<' as 0."""
    if "0" <= c <= "9":
        return ord(c) - ord("0")
    if "A" <= c <= "Z":
        return ord(c) - ord("A") + 10
    if c == "<":
        return 0
    raise ValueError(f"not an MRZ character: {c!r}")


def check_digit(data):
    """The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10."""
    return sum(char_value(c) * WEIGHTS[i % 3] for i, c in enumerate(data)) % 10
