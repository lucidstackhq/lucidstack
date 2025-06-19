package xyz.lucidstack.repository;

import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.Environment;

import java.util.Collection;
import java.util.List;

@Repository
public interface EnvironmentRepository extends MongoRepository<Environment, String> {

    boolean existsByNameAndOrganizationId(String name, String organizationId);

    List<Environment> findByOrganizationId(String organizationId, Pageable pageable);

    List<Environment> findByIdInAndOrganizationId(Collection<String> ids, String organizationId);

    Environment findByIdAndOrganizationId(String id, String organizationId);

    boolean existsByIdNotAndNameAndOrganizationId(String id, String name, String organizationId);
}
