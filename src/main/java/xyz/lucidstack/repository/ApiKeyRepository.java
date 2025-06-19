package xyz.lucidstack.repository;

import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.ApiKey;

import java.util.List;

@Repository
public interface ApiKeyRepository extends MongoRepository<ApiKey, String> {

    boolean existsByNameAndProjectIdAndOrganizationId(String name, String projectId, String organizationId);

    List<ApiKey> findByProjectIdAndOrganizationId(String projectId, String organizationId, Pageable pageable);

    ApiKey findByIdAndProjectIdAndOrganizationId(String id, String projectId, String organizationId);

    boolean existsByIdNotAndNameAndProjectIdAndOrganizationId(String id, String name, String projectId, String organizationId);
}
