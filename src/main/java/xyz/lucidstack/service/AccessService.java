package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.Resource;

@Service
@RequiredArgsConstructor
public class AccessService {

    public void addFullAccess(Resource resource, String userId, String organizationId) {

    }

    public Boolean hasPermission(Resource resource, AuthenticatedUser user, String permission) {
        return true;
    }
}
