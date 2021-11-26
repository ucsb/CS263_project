package doughnutshop.frontend.resource;

import doughnutshop.frontend.model.Doughnut;
import doughnutshop.frontend.model.DoughnutInventory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.List;

@RestController
@RequestMapping("/doughnuts")
public class DoughnutInventoryResource {
    @Autowired
    private RestTemplate restTemplate;

    @RequestMapping("/")
    public String getDoughnutInventory(){
        DoughnutInventory doughnutInventory = restTemplate.getForObject("http://inventory-service-database/doughnuts/orderByPrice", DoughnutInventory.class);
        List<Doughnut> doughnuts = doughnutInventory.getDoughnutsInventory();
        String doughnutView = "";
        for(Doughnut doughnut : doughnuts){
            doughnutView += "Name: " + doughnut.getName() + "<br>"  + "Price: $" +
                    doughnut.getPrice() + "<br>"  + "Inventory: " +
                    doughnut.getInventory() + "<br>"  + "<br>" ;
        }
        return doughnutView;
    }

    @RequestMapping("/orders")
    public String getDoughnutOrders(){
        String doughnutOrders = restTemplate.getForObject("http://checkout-service/orders", String.class);
        return doughnutOrders;
    }


}

