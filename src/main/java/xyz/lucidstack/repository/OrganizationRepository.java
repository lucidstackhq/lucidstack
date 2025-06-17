package xyz.lucidstack.repository;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.Organization;

@Repository
public interface OrganizationRepository extends MongoRepository<Organization, String> {

    Boolean existsByName(String name);

    Organization findByName(String name);
}
