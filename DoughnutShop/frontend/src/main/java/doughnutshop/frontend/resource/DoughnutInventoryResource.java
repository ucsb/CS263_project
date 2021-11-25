package doughnutshop.frontend.resource;

import doughnutshop.frontend.model.DoughnutInventory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
@RequestMapping("/doughnuts")
public class DoughnutInventoryResource {
    @Autowired
    private RestTemplate restTemplate;

    @RequestMapping("/")
    public DoughnutInventory getDoughnutInventory(){
        DoughnutInventory doughnutInventory = restTemplate.getForObject("http://inventory-service/doughnuts/inventory", DoughnutInventory.class);
        return doughnutInventory;
    }

}

