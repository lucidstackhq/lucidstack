package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Service;
import xyz.lucidstack.embedded.Actor;
import xyz.lucidstack.embedded.Resource;
import xyz.lucidstack.enums.ActorType;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.exception.ServerException;
import xyz.lucidstack.model.Access;
import xyz.lucidstack.model.ApiKey;
import xyz.lucidstack.model.User;
import xyz.lucidstack.response.AccessResponse;

import java.time.Instant;
import java.util.*;

@Service
@RequiredArgsConstructor
public class AccessManagementService {

    private final AccessService accessService;

    private final UserService userService;

    private final ApiKeyService apiKeyService;

    public void addPermission(Resource resource, String userId, Set<String> permissions, String organizationId) {
        if (!userService.exists(userId, organizationId)) {
            throw new NotFoundException("User not found");
        }

        accessService.addPermission(resource, userId, permissions, organizationId);
    }

    public void deletePermission(Resource resource, String userId, Set<String> permissions, String organizationId) {
        accessService.deletePermission(resource, userId, permissions, organizationId);
    }

    public List<AccessResponse> listForResource(Resource resource, String organizationId, Pageable pageable) {
        List<Access> accesses = accessService.listForResource(resource, organizationId, pageable);

        List<String> userIds = new ArrayList<>();
        List<String> apiKeyIds = new ArrayList<>();

        for (Access access : accesses) {
            switch (access.getActor().getType()) {
                case USER:
                    userIds.add(access.getActor().getReferenceId());
                    break;
                case API_KEY:
                    apiKeyIds.add(access.getActor().getReferenceId());
                    break;
                default:
                    throw new ServerException("Invalid actor type");
            }
        }

        Map<String, User> userMap = new HashMap<>();
        Map<String, ApiKey> apiKeyMap = new HashMap<>();

        if (!userIds.isEmpty()) {
            List<User> users = userService.get(userIds, organizationId);
            for (User user : users) {
                userMap.put(user.getId(), user);
            }
        }

        if (!apiKeyIds.isEmpty()) {
            List<ApiKey> apiKeys = apiKeyService.get(apiKeyIds, organizationId);
            for (ApiKey apiKey : apiKeys) {
                apiKeyMap.put(apiKey.getId(), apiKey);
            }
        }

        List<AccessResponse> accessesResponse = new ArrayList<>();

        for (Access access : accesses) {
            switch (access.getActor().getType()) {
                case USER:
                    User user = userMap.get(access.getActor().getReferenceId());
                    accessesResponse.add(AccessResponse.builder().access(access).actorData(user).build());
                    break;
                case API_KEY:
                    ApiKey apiKey = apiKeyMap.get(access.getActor().getReferenceId());
                    accessesResponse.add(AccessResponse.builder().access(access).actorData(apiKey).build());
                    break;
                default:
                    throw new ServerException("Invalid actor type");
            }
        }

        return accessesResponse;
    }
}
