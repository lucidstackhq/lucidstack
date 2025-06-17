package xyz.lucidstack.auth.authenticator;

import jakarta.servlet.http.HttpServletRequest;

public interface Authenticator {

    void authenticate(HttpServletRequest request);

    Class<?> getType();
}