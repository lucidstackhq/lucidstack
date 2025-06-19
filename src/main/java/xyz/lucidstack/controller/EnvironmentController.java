package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.Environment;
import xyz.lucidstack.request.EnvironmentCreationRequest;
import xyz.lucidstack.service.EnvironmentService;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class EnvironmentController extends AuthenticatedUserController {

    private final EnvironmentService environmentService;

    @PostMapping("/environments")
    public Environment create(@Valid @RequestBody EnvironmentCreationRequest request) {
        return environmentService.create(request, getUser());
    }
}
