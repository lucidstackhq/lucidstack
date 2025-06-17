package xyz.lucidstack.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import xyz.lucidstack.response.HealthCheckResponse;

@RestController
public class HealthCheckController {

    @GetMapping("/health")
    public HealthCheckResponse healthCheck() {
        return HealthCheckResponse.ok();
    }
}
