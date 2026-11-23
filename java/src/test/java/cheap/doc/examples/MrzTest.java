package cheap.doc.examples;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

class MrzTest {
    // Fields of the ICAO Doc 9303 specimen passport line.
    private static final String LINE = "L898902C36UTO7408122F1204159ZE184226B<<<<<10";

    @Test
    void fieldsOfTheIcaoSpecimen() {
        assertEquals(6, Mrz.checkDigit("L898902C3")); // document number
        assertEquals(2, Mrz.checkDigit("740812")); // birth date
        assertEquals(9, Mrz.checkDigit("120415")); // expiry date
        assertEquals(1, Mrz.checkDigit("ZE184226B<<<<<")); // personal number
    }

    @Test
    void compositeOfTheIcaoSpecimen() {
        String data = LINE.substring(0, 10) + LINE.substring(13, 20) + LINE.substring(21, 43);
        assertEquals(0, Mrz.checkDigit(data));
    }

    @Test
    void fillersCountAsZero() {
        assertEquals(0, Mrz.checkDigit("<<<<<<"));
    }

    @Test
    void oneMisreadCharacterChangesTheDigit() {
        assertNotEquals(6, Mrz.checkDigit("L898902C4"));
    }

    @Test
    void refusesACharacterOutsideTheAlphabet() {
        assertThrows(IllegalArgumentException.class, () -> Mrz.checkDigit("l898902c3"));
    }
}
