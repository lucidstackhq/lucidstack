package xyz.lucidstack.repository;

import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.embedded.Actor;
import xyz.lucidstack.embedded.Resource;
import xyz.lucidstack.model.Access;

import java.util.List;
import java.util.Set;

@Repository
public interface AccessRepository extends MongoRepository<Access, String> {

    Boolean existsByResourceAndActorAndOrganizationIdAndPermissionsIn(Resource resource, Actor actor, String organizationId, Set<String> permissions);

    List<Access> findByResourceAndOrganizationId(Resource resource, String organizationId, Pageable pageable);
}
