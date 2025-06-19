package xyz.lucidstack.repository;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.ApiKey;

@Repository
public interface ApiKeyRepository extends MongoRepository<ApiKey, String> {

    boolean existsByNameAndProjectIdAndOrganizationId(String name, String projectId, String organizationId);
}
