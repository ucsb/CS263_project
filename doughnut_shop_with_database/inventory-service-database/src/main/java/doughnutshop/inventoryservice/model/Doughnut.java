package doughnutshop.inventoryservice.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "doughnuts")
public class Doughnut {
    @Id
    private String id;

    private String name;
    private double price;
    private int inventory;

    public Doughnut(){

    }

    public Doughnut(String id, String name, double price, int inventory) {
        super(); //calling parents constructor
        this.id = id;
        this.name = name;
        this.price = price;
        this.inventory = inventory;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }

    public int getInventory() {
        return inventory;
    }

    public void setInventory(int inventory) {
        this.inventory = inventory;
    }

    public String getId(){
        return id;
    }

    public void setId(String id){
        this.id = id;
    }

    @Override
    public String toString() {
        return "Doughnut{" +
                ", name = '" + getName() + '\'' +
                ", price = " + getPrice() +
                ", Inventory = '" + getInventory() +
                '}';
    }
}
