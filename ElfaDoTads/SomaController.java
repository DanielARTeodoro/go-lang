package br.edu.ifms.soma.controller;

import br.edu.ifms.soma.exception.NoNumeroException;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/sum")
public class SomaController {

    @GetMapping(value = "/{num1}/{num2}")
    public Integer getSum(@PathVariable String num1, @PathVariable String num2){
         if(num1.matches("[0-9]+") && num2.matches("[0-9]+")){
             return Integer.parseInt(num1) + Integer.parseInt(num2);
         } else {
             throw new NoNumeroException("Não pode usar letras em Soma!");
         }
    }
}
