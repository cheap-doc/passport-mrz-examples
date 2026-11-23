package cheap.doc.examples;

/** MRZ check digits as ICAO Doc 9303 defines them. No dependencies. */
public final class Mrz {
    private static final int[] WEIGHTS = {7, 3, 1};

    private Mrz() {}

    /** The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, '&lt;' as 0. */
    public static int charValue(char c) {
        if (c >= '0' && c <= '9') return c - '0';
        if (c >= 'A' && c <= 'Z') return c - 'A' + 10;
        if (c == '<') return 0;
        throw new IllegalArgumentException("not an MRZ character: '" + c + "'");
    }

    /** The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10. */
    public static int checkDigit(String data) {
        int sum = 0;
        for (int i = 0; i < data.length(); i++) {
            sum += charValue(data.charAt(i)) * WEIGHTS[i % 3];
        }
        return sum % 10;
    }
}
