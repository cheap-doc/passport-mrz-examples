using PassportScan;
using Xunit;

public class CheckDigitTests
{
    // Fields of the ICAO Doc 9303 specimen passport line.
    private const string Line = "L898902C36UTO7408122F1204159ZE184226B<<<<<10";

    [Theory]
    [InlineData("L898902C3", 6)] // document number
    [InlineData("740812", 2)] // birth date
    [InlineData("120415", 9)] // expiry date
    [InlineData("ZE184226B<<<<<", 1)] // personal number
    [InlineData("<<<<<<", 0)] // fillers count as zero
    public void FieldsOfTheIcaoSpecimen(string data, int expected) =>
        Assert.Equal(expected, Mrz.CheckDigit(data));

    [Fact]
    public void CompositeOfTheIcaoSpecimen() =>
        Assert.Equal(0, Mrz.CheckDigit(Line[0..10] + Line[13..20] + Line[21..43]));

    [Fact]
    public void OneMisreadCharacterChangesTheDigit() =>
        Assert.NotEqual(6, Mrz.CheckDigit("L898902C4"));

    [Fact]
    public void RefusesACharacterOutsideTheAlphabet() =>
        Assert.Throws<ArgumentException>(() => Mrz.CheckDigit("l898902c3"));
}
