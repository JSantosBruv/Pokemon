package Pokemon;

/*Class representing the Pokemon.Position of Ash in the game*/
public class Position {

    private int x; //East and West Moves for X axis
    private int y; //North and South Moves for Y axis

    public Position(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public void goNorth() {
        y += 1;
    }

    public void goSouth() {
        y -= 1;
    }

    public void goWest() {
        x -= 1;
    }

    public void goEast() {
        x += 1;
    }

    /*Custom toString to use in Set for unique Positions and O(1) lookup*/
    @Override
    public String toString() {
        return "Pokemon.Position{" + "x=" + x + ", y=" + y + '}';
    }
}
