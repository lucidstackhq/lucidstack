package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import xyz.lucidstack.model.User;
import xyz.lucidstack.request.UserSignUpRequest;
import xyz.lucidstack.request.UserTokenRequest;
import xyz.lucidstack.response.UserTokenResponse;
import xyz.lucidstack.service.UserService;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class AccountController {

    private final UserService userService;

    @PostMapping("/users/signup")
    public User signUp(@Valid @RequestBody UserSignUpRequest request) {
        return userService.signUp(request);
    }

    @PostMapping("/users/token")
    public UserTokenResponse getToken(@Valid @RequestBody UserTokenRequest request) {
        return userService.getToken(request);
    }
}
