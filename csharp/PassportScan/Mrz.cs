namespace PassportScan;

/// <summary>MRZ check digits as ICAO Doc 9303 defines them. No dependencies.</summary>
public static class Mrz
{
    private static readonly int[] Weights = { 7, 3, 1 };

    /// <summary>The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, '&lt;' as 0.</summary>
    public static int CharValue(char c) => c switch
    {
        >= '0' and <= '9' => c - '0',
        >= 'A' and <= 'Z' => c - 'A' + 10,
        '<' => 0,
        _ => throw new ArgumentException($"not an MRZ character: '{c}'"),
    };

    /// <summary>The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10.</summary>
    public static int CheckDigit(string data)
    {
        var sum = 0;
        for (var i = 0; i < data.Length; i++)
        {
            sum += CharValue(data[i]) * Weights[i % 3];
        }
        return sum % 10;
    }
}
