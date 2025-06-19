package xyz.lucidstack.repository;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.Environment;

@Repository
public interface EnvironmentRepository extends MongoRepository<Environment, String> {

    boolean existsByNameAndOrganizationId(String name, String organizationId);
}
