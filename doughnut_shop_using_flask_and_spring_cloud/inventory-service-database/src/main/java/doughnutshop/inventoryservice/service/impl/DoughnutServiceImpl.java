package doughnutshop.inventoryservice.service.impl;

import doughnutshop.inventoryservice.model.Doughnut;
import doughnutshop.inventoryservice.repository.DoughnutRepository;
import doughnutshop.inventoryservice.service.DoughnutService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class DoughnutServiceImpl implements DoughnutService {
    @Autowired
    private DoughnutRepository doughnutRepository;

    @Override
    public List<Doughnut> findAll() {
        return doughnutRepository.findAll();
    }

    @Override
    public Doughnut findByName(String name){
        return doughnutRepository.findByName(name);
    }

    @Override
    public List<Doughnut> findAllByOrderByPrice(){
        return doughnutRepository.findAllByOrderByPrice();
    }

    @Override
    public Doughnut saveOrUpdateDoughnut(Doughnut doughnut){
        return doughnutRepository.save(doughnut);
    }

    @Override
    public void deleteDoughnutById(String id){
        doughnutRepository.deleteById(id);
    }
}
