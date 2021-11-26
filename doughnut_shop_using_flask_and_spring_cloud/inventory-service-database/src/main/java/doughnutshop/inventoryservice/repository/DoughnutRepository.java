package doughnutshop.inventoryservice.repository;

import doughnutshop.inventoryservice.model.Doughnut;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import java.util.List;

public interface DoughnutRepository extends MongoRepository<Doughnut, String> {

    Doughnut findByName(String name);

    List<Doughnut> findAllByOrderByPrice();

}
