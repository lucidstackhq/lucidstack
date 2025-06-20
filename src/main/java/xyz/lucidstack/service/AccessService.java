package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.Actor;
import xyz.lucidstack.embedded.Resource;
import xyz.lucidstack.enums.ActorType;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.exception.ServerException;
import xyz.lucidstack.model.Access;
import xyz.lucidstack.model.ApiKey;
import xyz.lucidstack.model.User;
import xyz.lucidstack.repository.AccessRepository;
import xyz.lucidstack.response.AccessResponse;

import java.time.Instant;
import java.util.*;

@Service
@RequiredArgsConstructor
public class AccessService {

    private final AccessRepository accessRepository;

    private final MongoTemplate mongoTemplate;

    private final UserService userService;

    private final ApiKeyService apiKeyService;

    public void addFullAccess(Resource resource, String userId, String organizationId) {
        Actor actor = Actor.builder().type(ActorType.USER).referenceId(userId).build();

        mongoTemplate.upsert(Query.query(Criteria
                        .where("actor").is(actor)
                        .and("resource").is(resource)
                        .and("organizationId").is(organizationId)),
                new Update()
                        .setOnInsert("actor", actor)
                        .setOnInsert("resource", resource)
                        .setOnInsert("organizationId", organizationId)
                        .setOnInsert("createdAt", Instant.now())
                        .set("updatedAt", Instant.now())
                        .addToSet("permissions", "*"),
                Access.class);
    }

    public void addPermission(Resource resource, String userId, Set<String> permissions, String organizationId) {
        if (!userService.exists(userId, organizationId)) {
            throw new NotFoundException("User not found");
        }

        Actor actor = Actor.builder().type(ActorType.USER).referenceId(userId).build();

        mongoTemplate.upsert(Query.query(Criteria
                        .where("actor").is(actor)
                        .and("resource").is(resource)
                        .and("organizationId").is(organizationId)),
                new Update()
                        .setOnInsert("actor", actor)
                        .setOnInsert("resource", resource)
                        .setOnInsert("organizationId", organizationId)
                        .setOnInsert("createdAt", Instant.now())
                        .set("updatedAt", Instant.now())
                        .addToSet("permissions").each(permissions.toArray()),
                Access.class);
    }

    public void deletePermission(Resource resource, String userId, Set<String> permissions, String organizationId) {
        Actor actor = Actor.builder().type(ActorType.USER).referenceId(userId).build();

        mongoTemplate.updateFirst(Query.query(Criteria
                        .where("actor").is(actor)
                        .and("resource").is(resource)
                        .and("organizationId").is(organizationId)),
                new Update()
                        .set("updatedAt", Instant.now())
                        .pullAll("permissions", permissions.toArray()),
                Access.class);
    }

    public List<AccessResponse> listForResource(Resource resource, String organizationId, Pageable pageable) {
        List<Access> accesses = accessRepository.findByResourceAndOrganizationId(resource, organizationId, pageable);

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

    public Boolean hasPermission(Resource resource, AuthenticatedUser user, String permission) {
        if (user.getAdmin()) {
            return true;
        }

        return accessRepository.existsByResourceAndActorAndOrganizationIdAndPermissionsIn(resource, Actor.builder()
                .type(ActorType.USER)
                .referenceId(user.getId())
                .build(), user.getOrganizationId(), Set.of(permission, "*"));
    }

    public List<Resource> listResources(Map<String, Object> resourceFilter, String permission, AuthenticatedUser user, Pageable pageable) {
        Criteria criteria = Criteria
                .where("permissions").in("*", permission)
                .and("actor").is(Actor.builder().type(ActorType.USER).referenceId(user.getId()).build())
                .and("organizationId").is(user.getOrganizationId());

        for (Map.Entry<String, Object> entry : resourceFilter.entrySet()) {
            criteria = criteria.and(String.format("resource.%s", entry.getKey())).is(entry.getValue());
        }

        Query query = Query.query(criteria).with(pageable);
        List<Access> accesses = mongoTemplate.find(query, Access.class);
        return accesses.stream().map(Access::getResource).toList();
    }
}
