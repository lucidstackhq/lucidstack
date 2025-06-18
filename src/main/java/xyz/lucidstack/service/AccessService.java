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
import xyz.lucidstack.model.Access;
import xyz.lucidstack.repository.AccessRepository;

import java.time.Instant;
import java.util.List;
import java.util.Map;
import java.util.Set;

@Service
@RequiredArgsConstructor
public class AccessService {

    private final AccessRepository accessRepository;

    private final MongoTemplate mongoTemplate;

    public void addFullAccess(Resource resource, String userId, String organizationId) {
        Actor actor = Actor.builder().type(ActorType.USER).id(userId).build();

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

    public Boolean hasPermission(Resource resource, AuthenticatedUser user, String permission) {
        return accessRepository.existsByResourceAndActorAndOrganizationIdAndPermissionsIn(resource, Actor.builder()
                .type(ActorType.USER)
                .id(user.getId())
                .build(), user.getOrganizationId(), Set.of(permission, "*"));
    }

    public List<Resource> listResources(Map<String, Object> resourceFilter, String permission, AuthenticatedUser user, Pageable pageable) {
        Criteria criteria = Criteria
                .where("permissions").in("*", permission)
                .and("actor").is(Actor.builder().type(ActorType.USER).id(user.getId()).build())
                .and("organizationId").is(user.getOrganizationId());

        for (Map.Entry<String, Object> entry : resourceFilter.entrySet()) {
            criteria = criteria.and(String.format("resource.%s", entry.getKey())).is(entry.getValue());
        }

        Query query = Query.query(criteria).with(pageable);
        return mongoTemplate.find(query, Resource.class);
    }
}
