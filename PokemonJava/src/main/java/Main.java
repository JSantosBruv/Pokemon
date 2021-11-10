import Exceptions.UnknownMoveException;
import Pokemon.Pokemon;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;


public class Main {

    public static void main(String[] args) throws IOException, UnknownMoveException {

        //Read Input Once
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

        String moves = reader.readLine();

        //Initialize Game
        Pokemon game = new Pokemon();

        // Run game with provided moves
        int count = game.gottaCatchEmAll(moves);

        //Print result
        System.out.println(count);
    }
}
