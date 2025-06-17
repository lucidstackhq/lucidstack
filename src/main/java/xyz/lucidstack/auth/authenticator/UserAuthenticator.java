package xyz.lucidstack.auth.authenticator;

import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.auth.Jwt;
import xyz.lucidstack.auth.ThreadLocalWrapper;
import xyz.lucidstack.exception.InvalidTokenException;
import org.springframework.stereotype.Component;

import java.util.Objects;

@Component
@RequiredArgsConstructor
public class UserAuthenticator implements Authenticator {

    private final Jwt jwt;

    @Override
    public void authenticate(HttpServletRequest request) {
        String jwtToken = AuthUtils.extractToken(request);

        AuthenticatedUser user = jwt.getUser(jwtToken);

        if (Objects.isNull(user)) {
            throw new InvalidTokenException();
        }

        ThreadLocalWrapper.setUser(user);
    }

    @Override
    public Class<?> getType() {
        return AuthenticatedUserController.class;
    }
}