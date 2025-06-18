package xyz.lucidstack.embedded.resource;

import lombok.Getter;
import lombok.Setter;
import xyz.lucidstack.embedded.Resource;

@Getter
@Setter
public class ProjectResource extends Resource {

    private String projectId;

    public ProjectResource(String projectId) {
        this.setType("project");
        this.projectId = projectId;
    }
}
