package Exceptions;

public class UnknownMoveException extends Exception {

    public UnknownMoveException(String errorMessage) {
        super(errorMessage);
    }
}

