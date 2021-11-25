package doughnutshop.inventoryservice.resource;

import doughnutshop.inventoryservice.model.Doughnut;
import doughnutshop.inventoryservice.model.DoughnutInventory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/doughnuts")
public class DoughnutInventoryResource {

    @RequestMapping("/inventory")
    public DoughnutInventory getDoughnutInventory(){
        List<Doughnut> doughnuts = new ArrayList<>();
        doughnuts.add(new Doughnut("Chocolate", 2.00, 400));
        doughnuts.add(new Doughnut("Sugar", 2.00, 350));
        doughnuts.add(new Doughnut("Glazed", 2.00, 700));

        DoughnutInventory doughnutInventory = new DoughnutInventory();
        doughnutInventory.setDoughnutsInventory(doughnuts);
        return doughnutInventory;
    }

}