package xyz.lucidstack.repository;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.User;

@Repository
public interface UserRepository extends MongoRepository<User, String> {

    User findByUsernameAndOrganizationId(String username, String organizationId);

    User findByIdAndOrganizationId(String id, String organizationId);
}
