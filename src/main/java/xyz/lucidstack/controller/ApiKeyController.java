package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.ApiKey;
import xyz.lucidstack.request.ApiKeyCreationRequest;
import xyz.lucidstack.service.ApiKeyService;

@RestController
@RequestMapping("/api/v1/projects/{projectId}")
@RequiredArgsConstructor
public class ApiKeyController extends AuthenticatedUserController {

    private final ApiKeyService apiKeyService;

    @PostMapping("/api-keys")
    public ApiKey create(@PathVariable String projectId, @Valid @RequestBody ApiKeyCreationRequest request) {
        return apiKeyService.create(projectId, request, getUser());
    }
}
