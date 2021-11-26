package doughnutshop.inventoryservice.resource;

import doughnutshop.inventoryservice.model.CheckoutDoughnut;
import doughnutshop.inventoryservice.model.Doughnut;
import doughnutshop.inventoryservice.model.DoughnutInventory;
import doughnutshop.inventoryservice.service.DoughnutService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.zip.CheckedInputStream;

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

    @GetMapping(value = "/checkout/{name}/{quantity}")
    public ResponseEntity<?> checkoutDoughnutByName(@PathVariable String name, @PathVariable int quantity) {
        //Doughnut doughnut = doughnutService.findByName(checkoutDoughnut.getName());
        Doughnut doughnut = doughnutService.findByName(name);
        int doughnutInventory = doughnut.getInventory();
        //int doughnutsToBuy = checkoutDoughnut.getQuantity();
        int doughnutsToBuy = quantity;
        if (doughnutInventory > doughnutsToBuy) {
            Doughnut updatedDoughnut = new Doughnut(doughnut.getId(), doughnut.getName(), doughnut.getPrice(), doughnutInventory - doughnutsToBuy);
            doughnut = doughnutService.saveOrUpdateDoughnut(updatedDoughnut);
            return new ResponseEntity("Doughnut checked out successfully", HttpStatus.OK);
            //return new ResponseEntity("Doughnut checked out successfully, doughnut quantity is now " + doughnut.getInventory(), HttpStatus.OK);
        } else if (doughnutInventory == doughnutsToBuy) {
            doughnutService.deleteDoughnutById(doughnut.getId());
            return new ResponseEntity("Doughnut checked out successfully", HttpStatus.OK);
            //return new ResponseEntity("You got the last of the "+ doughnut.getName() +"Doughnuts!", HttpStatus.OK);
        } else {
            return new ResponseEntity("Sorry we don't have that many doughnuts available or we are completely sold out!", HttpStatus.OK);
        }
    }
}
