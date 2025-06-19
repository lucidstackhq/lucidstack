package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.ApiKey;
import xyz.lucidstack.request.ApiKeyCreationRequest;
import xyz.lucidstack.request.ApiKeyUpdateRequest;
import xyz.lucidstack.response.ApiKeySecretResponse;
import xyz.lucidstack.service.ApiKeyService;

import java.util.List;

@RestController
@RequestMapping("/api/v1/projects/{projectId}")
@RequiredArgsConstructor
public class ApiKeyController extends AuthenticatedUserController {

    private final ApiKeyService apiKeyService;

    @PostMapping("/api-keys")
    public ApiKey create(@PathVariable String projectId, @Valid @RequestBody ApiKeyCreationRequest request) {
        return apiKeyService.create(projectId, request, getUser());
    }

    @GetMapping("/api-keys")
    public List<ApiKey> list(@PathVariable String projectId, Pageable pageable) {
        return apiKeyService.list(projectId, getUser(), pageable);
    }

    @GetMapping("/api-keys/{apiKeyId}")
    public ApiKey get(@PathVariable String  projectId, @PathVariable String apiKeyId) {
        return apiKeyService.get(apiKeyId, projectId, getUser());
    }

    @PutMapping("/api-keys/{apiKeyId}")
    public ApiKey update(@PathVariable String projectId, @PathVariable String apiKeyId, @Valid @RequestBody ApiKeyUpdateRequest request) {
        return apiKeyService.update(apiKeyId, request, projectId, getUser());
    }

    @DeleteMapping("/api-keys/{apiKeyId}")
    public ApiKey delete(@PathVariable String apiKeyId, @PathVariable String projectId) {
        return apiKeyService.delete(apiKeyId, projectId, getUser());
    }

    @GetMapping("/api-keys/{apiKeyId}/secret")
    public ApiKeySecretResponse getSecret(@PathVariable String  projectId, @PathVariable String apiKeyId) {
        return apiKeyService.getSecret(apiKeyId, projectId, getUser());
    }

    @PutMapping("/api-keys/{apiKeyId}/secret")
    public ApiKeySecretResponse resetSecret(@PathVariable String projectId, @PathVariable String apiKeyId) {
        return apiKeyService.resetSecret(apiKeyId, projectId, getUser());
    }
}
