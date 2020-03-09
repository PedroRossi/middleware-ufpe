public class Bridge {
    private Car current;

    public Bridge() {}

    public void setCurrent(Car current) {
        this.current = current;
    }

    public void removeCurrent() {
        this.current = null;
    }

    public Car getCurrent() {
        return this.current;
    }
}
