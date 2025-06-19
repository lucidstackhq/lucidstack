package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.resource.RootResource;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.exception.NotAllowedException;
import xyz.lucidstack.model.Environment;
import xyz.lucidstack.repository.EnvironmentRepository;
import xyz.lucidstack.request.EnvironmentCreationRequest;

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
}
