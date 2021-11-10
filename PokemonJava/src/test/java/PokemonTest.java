import Exceptions.UnknownMoveException;
import Pokemon.Pokemon;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class PokemonTest {

    private Pokemon engine;

    @BeforeEach
    public void initEngine() {
        this.engine = new Pokemon();
    }

    @Test
    public void TestGivenCases() throws UnknownMoveException {
        String moves = "E";

        int expectedCount = 2;

        int result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);

        initEngine();

        moves = "NESO";

        expectedCount = 4;

        result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);

        initEngine();

        moves = "NSNSNSNSNS";

        expectedCount = 2;

        result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);
    }

    @Test
    public void TestCustom() throws UnknownMoveException {
        String moves = "NSSN";

        int expectedCount = 3;

        int result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);

        initEngine();

        moves = "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN";

        expectedCount = 200;

        result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);

    }

    @Test
    public void TestEmpty() throws UnknownMoveException {
        String moves = "";

        int expectedCount = 1;

        int result = engine.gottaCatchEmAll(moves);

        Assertions.assertEquals(expectedCount, result);


    }

    @Test()
    public void TestForUnknownMove() {

        String moves = "Isto deve dar erro (hopefully)";


        Assertions.assertThrows(UnknownMoveException.class, () -> engine.gottaCatchEmAll(moves));

    }
}
