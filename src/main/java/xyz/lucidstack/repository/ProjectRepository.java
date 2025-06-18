package xyz.lucidstack.repository;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;
import xyz.lucidstack.model.Project;

@Repository
public interface ProjectRepository extends MongoRepository<Project, String> {

}
