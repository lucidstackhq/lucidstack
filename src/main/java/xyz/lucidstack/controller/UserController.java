package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.User;
import xyz.lucidstack.request.UserPasswordChangeRequest;
import xyz.lucidstack.service.UserService;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class UserController extends AuthenticatedUserController {

    private final UserService userService;

    @GetMapping("/users/me")
    public User get() {
        return userService.get(getUserId(), getOrganizationId());
    }

    @PutMapping("/users/me/password")
    public User changePassword(@Valid @RequestBody UserPasswordChangeRequest request) {
        return userService.changePassword(getUserId(), request, getOrganizationId());
    }
}
