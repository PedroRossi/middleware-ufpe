public class Car {
    private int id;
    private boolean direction;

    public Car(int id, boolean direction) {
        this.id = id;
        this.direction = direction;
    }

    public int getId() {
        return id;
    }

    public boolean getDirection() {
        return this.direction;
    }

    @Override
    public String toString() {
        return "Car{" +
                "id=" + id +
                ", direction=" + direction +
                '}';
    }
}
