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

    public List<Access> listForResource(Resource resource, String organizationId, Pageable pageable) {
        return accessRepository.findByResourceAndOrganizationId(resource, organizationId, pageable);
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
