package Pokemon;

import Exceptions.UnknownMoveException;

import java.util.HashSet;
import java.util.Set;

/*Class representing Engine of the "game"*/
public class Pokemon {

    private static final String UNKNOWN_MOVE = "Move not supported. Available moves: [N,S,E,O]";

    private enum Move {
        N, S, E, O
    }

    private final Position currentPosition;
    private final Set<String> visitedSpots;
    private int pokemonCount;

    public Pokemon() {
        this.currentPosition = new Position(0, 0);  //Game starts at 0,0
        this.visitedSpots = new HashSet<>();              //HashSet for Unique Positions and O(1) Lookups
        this.pokemonCount = 1;                            //Assume pokemon count is minimum of 1

        visitedSpots.add(currentPosition.toString());     //Add initial Pokemon.Position
    }


    public int gottaCatchEmAll(String moves) throws UnknownMoveException {

        //Moves to char to iterate each move
        char[] charMoves = moves.toCharArray();

        //Iterate moves and apply action according to the given move
        for (char move : charMoves) {

            Move moveString = getEnum(move);

            switch (moveString) {
                case N:
                    currentPosition.goNorth();
                    break;
                case S:
                    currentPosition.goSouth();
                    break;
                case E:
                    currentPosition.goEast();
                    break;
                case O:
                    currentPosition.goWest();
                    break;
                default:
                    throw new UnknownMoveException(UNKNOWN_MOVE);

            }

            checkPosition();

        }
        return pokemonCount;
    }

    /*Check if position has been visited.
     * If it has, do nothing.
     * If it hasn't, add position to visitedSpots and increment pokemonCount by one*/
    private void checkPosition() {

        String positionKey = currentPosition.toString();

        if (!visitedSpots.contains(positionKey)) {

            visitedSpots.add(positionKey);
            pokemonCount++;
        }
    }

    /*Gets Move enum based on the char move given
     * Throws Exceptions.Exceptions.UnknownMoveException if move is not allowed*/
    private Move getEnum(char move) throws UnknownMoveException {

        try {

            return Move.valueOf(String.valueOf(move));

        } catch (IllegalArgumentException e) {
            throw new UnknownMoveException(UNKNOWN_MOVE);
        }
    }
}
