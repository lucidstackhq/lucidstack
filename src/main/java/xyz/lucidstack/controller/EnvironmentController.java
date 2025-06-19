package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.Environment;
import xyz.lucidstack.request.EnvironmentCreationRequest;
import xyz.lucidstack.service.EnvironmentService;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class EnvironmentController extends AuthenticatedUserController {

    private final EnvironmentService environmentService;
    private final org.springframework.core.env.Environment environment;

    @PostMapping("/environments")
    public Environment create(@Valid @RequestBody EnvironmentCreationRequest request) {
        return environmentService.create(request, getUser());
    }

    @GetMapping("/environments")
    public List<Environment> list(Pageable pageable) {
        return environmentService.list(getUser(), pageable);
    }

    @GetMapping("/environments/{environmentId}")
    public Environment get(@PathVariable String environmentId) {
        return environmentService.get(environmentId, getUser());
    }
}
