package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.resource.ProjectResource;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.exception.NotAllowedException;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.model.ApiKey;
import xyz.lucidstack.repository.ApiKeyRepository;
import xyz.lucidstack.request.ApiKeyCreationRequest;
import xyz.lucidstack.util.Random;

import java.util.List;

@Service
@RequiredArgsConstructor
public class ApiKeyService {

    private final ApiKeyRepository apiKeyRepository;

    private final ProjectService projectService;

    private final AccessService accessService;

    public ApiKey create(String projectId, ApiKeyCreationRequest request, AuthenticatedUser creator) {
        if (!projectService.exists(projectId, creator.getOrganizationId())) {
            throw new NotFoundException("Project not found");
        }

        if (!accessService.hasPermission(new ProjectResource(projectId), creator, "manage_api_keys")) {
            throw new NotAllowedException();
        }

        if (apiKeyRepository.existsByNameAndProjectIdAndOrganizationId(request.getName(), projectId, creator.getOrganizationId())) {
            throw new ClientException(String.format("Api key %s already exists", request.getName()));
        }

        ApiKey apiKey = ApiKey.builder()
                .name(request.getName())
                .description(request.getDescription())
                .projectId(projectId)
                .secret(Random.generateRandomString(128))
                .creatorId(creator.getId())
                .organizationId(creator.getOrganizationId())
                .build();

        return apiKeyRepository.save(apiKey);
    }

    public List<ApiKey> list(String projectId, AuthenticatedUser requester, Pageable pageable) {
        if (!accessService.hasPermission(new ProjectResource(projectId), requester, "manage_api_keys")) {
            throw new NotAllowedException();
        }

        return apiKeyRepository.findByProjectIdAndOrganizationId(projectId, requester.getOrganizationId(), pageable);
    }
}
