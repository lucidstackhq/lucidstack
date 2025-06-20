package xyz.lucidstack.repository;

import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.User;

import java.util.Collection;
import java.util.List;

@Repository
public interface UserRepository extends MongoRepository<User, String> {

    User findByUsernameAndOrganizationId(String username, String organizationId);

    User findByIdAndOrganizationId(String id, String organizationId);

    boolean existsByUsernameAndOrganizationId(String username, String organizationId);

    List<User> findByOrganizationId(String organizationId, Pageable pageable);

    List<User> findByIdInAndOrganizationId(Collection<String> ids, String organizationId);

    Boolean existsByIdAndOrganizationId(String id, String organizationId);
}
