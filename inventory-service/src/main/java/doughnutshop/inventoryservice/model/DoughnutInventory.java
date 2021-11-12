package doughnutshop.inventoryservice.model;

import java.util.List;

public class DoughnutInventory {
    private List<Doughnut> doughnutsInventory;

    public DoughnutInventory(List<Doughnut> doughnutsInventory) {
        this.doughnutsInventory = doughnutsInventory;
    }

    public List<Doughnut> getDoughnutsInventory() {
        return doughnutsInventory;
    }

    public void setDoughnutsInventory(List<Doughnut> doughnutsInventory) {
        this.doughnutsInventory = doughnutsInventory;
    }
}
