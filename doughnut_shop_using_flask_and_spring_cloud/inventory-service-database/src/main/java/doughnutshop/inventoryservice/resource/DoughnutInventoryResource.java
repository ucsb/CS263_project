package doughnutshop.inventoryservice.resource;

import doughnutshop.inventoryservice.model.Doughnut;
import doughnutshop.inventoryservice.model.DoughnutInventory;
import doughnutshop.inventoryservice.service.DoughnutService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/doughnuts")
public class DoughnutInventoryResource {

    @Autowired
    private DoughnutService doughnutService;

    @GetMapping(value = "/")
    public DoughnutInventory getAllDoughnuts() {
        List<Doughnut> doughnuts = doughnutService.findAll();
        DoughnutInventory doughnutInventory = new DoughnutInventory(doughnuts);
        return doughnutInventory;
    }

    @GetMapping(value = "/byDoughnutName/{doughnutName}")
    public Doughnut getDoughnutByDoughnutName(@PathVariable("doughnutName") String doughnutName) {
        return doughnutService.findByName(doughnutName);
    }

    @GetMapping(value = "/orderByPrice")
    public DoughnutInventory findAllByOrderByPrice() {
        List<Doughnut> doughnuts = doughnutService.findAllByOrderByPrice();
        DoughnutInventory doughnutInventory = new DoughnutInventory(doughnuts);
        return doughnutInventory;
    }

    @PostMapping(value = "/save")
    public ResponseEntity<?> saveOrUpdateDoughnut(@RequestBody Doughnut doughnut) {
        doughnutService.saveOrUpdateDoughnut(doughnut);
        return new ResponseEntity("Doughnut added successfully", HttpStatus.OK);
    }

    @DeleteMapping(value = "/delete/{name}")
    public ResponseEntity<?> deleteDoughnutByName(@PathVariable String name) {
        doughnutService.deleteDoughnutById(doughnutService.findByName(name).getId());
        return new ResponseEntity("Doughnut deleted successfully", HttpStatus.OK);
    }

    @PostMapping(value = "/checkout/{name}")
    public ResponseEntity<?> checkoutDoughnutByName(@PathVariable String name) {
        Doughnut doughnut = doughnutService.findByName(name);
        int doughnutInventory = doughnut.getInventory();
        if (doughnutInventory > 1) {
            Doughnut updatedDoughnut = new Doughnut(doughnut.getId(), doughnut.getName(), doughnut.getPrice(), doughnut.getInventory() - 1);
            doughnut = doughnutService.saveOrUpdateDoughnut(updatedDoughnut);
            return new ResponseEntity("Doughnut checked out successfully, doughnut quantity is now " + doughnut.getInventory(), HttpStatus.OK);
        } else if (doughnutInventory == 1) {
            doughnutService.deleteDoughnutById(doughnut.getId());
            return new ResponseEntity("You got the last "+ doughnut.getName() +"!", HttpStatus.OK);
        } else {
            return new ResponseEntity("Doughnut not available", HttpStatus.OK);
        }
    }
}
