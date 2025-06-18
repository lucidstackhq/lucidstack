package xyz.lucidstack.repository;

import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.Project;

import java.util.Collection;
import java.util.List;

@Repository
public interface ProjectRepository extends MongoRepository<Project, String> {

    boolean existsByNameAndOrganizationId(String name, String organizationId);

    List<Project> findByOrganizationId(String organizationId, Pageable pageable);

    List<Project> findByIdInAndOrganizationId(Collection<String> ids, String organizationId);

    Project findByIdAndOrganizationId(String id, String organizationId);
}
