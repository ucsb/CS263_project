package doughnutshop.inventoryservice.service;

import doughnutshop.inventoryservice.model.Doughnut;

import java.util.List;

public interface DoughnutService {
    List<Doughnut> findAll();

    Doughnut findByName(String name);

    List<Doughnut> findAllByOrderByPrice();

    Doughnut saveOrUpdateDoughnut(Doughnut doughnut);

    void deleteDoughnutById(String id);
}
