package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.Resource;
import xyz.lucidstack.embedded.resource.EnvironmentResource;
import xyz.lucidstack.embedded.resource.RootResource;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.exception.NotAllowedException;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.model.Environment;
import xyz.lucidstack.repository.EnvironmentRepository;
import xyz.lucidstack.request.EnvironmentCreationRequest;
import xyz.lucidstack.request.EnvironmentUpdateRequest;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Service
@RequiredArgsConstructor
public class EnvironmentService {

    private final EnvironmentRepository environmentRepository;

    private final AccessService accessService;

    public Environment create(EnvironmentCreationRequest request, AuthenticatedUser creator) {
        if (!accessService.hasPermission(new RootResource(), creator, "create_environment")) {
            throw new NotAllowedException();
        }

        if (environmentRepository.existsByNameAndOrganizationId(request.getName(), creator.getOrganizationId())) {
            throw new ClientException(String.format("Environment %s already exists", request.getName()));
        }

        Environment environment = Environment.builder()
                .name(request.getName())
                .description(request.getDescription())
                .creatorId(creator.getId())
                .organizationId(creator.getOrganizationId())
                .build();

        return environmentRepository.save(environment);
    }

    public List<Environment> list(AuthenticatedUser requester, Pageable pageable) {
        if (requester.getAdmin()) {
            return environmentRepository.findByOrganizationId(requester.getOrganizationId(), pageable);
        } else {
            List<String> environmentIds = new ArrayList<>();
            List<Resource> resources = accessService.listResources(Map.of("type", "environment"), "read", requester, pageable);
            for (Resource resource : resources) {
                EnvironmentResource environmentResource = (EnvironmentResource) resource;
                environmentIds.add(environmentResource.getEnvironmentId());
            }

            return environmentRepository.findByIdInAndOrganizationId(environmentIds, requester.getOrganizationId());
        }
    }

    public Environment get(String environmentId, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new EnvironmentResource(environmentId), requester, "read")) {
            throw new NotAllowedException();
        }

        Environment environment = environmentRepository.findByIdAndOrganizationId(environmentId, requester.getOrganizationId());

        if (environment == null) {
            throw new NotFoundException("Environment not found");
        }

        return environment;
    }

    public Environment update(String environmentId, EnvironmentUpdateRequest request, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new EnvironmentResource(environmentId), requester, "update")) {
            throw new NotAllowedException();
        }

        Environment environment = environmentRepository.findByIdAndOrganizationId(environmentId, requester.getOrganizationId());

        if (environment == null) {
            throw new NotFoundException("Environment not found");
        }

        if (StringUtils.hasText(request.getName())) {
            if (environmentRepository.existsByIdNotAndNameAndOrganizationId(environmentId, request.getName(), requester.getOrganizationId())) {
                throw new ClientException(String.format("Environment %s already exists", request.getName()));
            }

            environment.setName(request.getName());
        }

        environment.setDescription(request.getDescription());

        return environmentRepository.save(environment);
    }

    public Environment delete(String environmentId, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new EnvironmentResource(environmentId), requester, "delete")) {
            throw new NotAllowedException();
        }

        Environment environment = environmentRepository.findByIdAndOrganizationId(environmentId, requester.getOrganizationId());

        if (environment == null) {
            throw new NotFoundException("Environment not found");
        }

        environmentRepository.delete(environment);
        return environment;
    }
}
