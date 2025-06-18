package xyz.lucidstack.embedded.resource;

import lombok.Getter;
import lombok.Setter;
import xyz.lucidstack.embedded.Resource;

@Getter
@Setter
public class EnvironmentResource extends Resource {

    private String environmentId;

    public EnvironmentResource(String environmentId) {
        this.environmentId = environmentId;
        this.setType("environment");
    }
}
