package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.User;
import xyz.lucidstack.request.UserAdditionRequest;
import xyz.lucidstack.request.UserAdminUpdateRequest;
import xyz.lucidstack.request.UserPasswordChangeRequest;
import xyz.lucidstack.service.UserPasswordResponse;
import xyz.lucidstack.service.UserService;

import java.util.List;

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

    @PostMapping("/users")
    public UserPasswordResponse add(@Valid @RequestBody UserAdditionRequest request) {
        checkAdmin();
        return userService.add(request, getUserId(), getOrganizationId());
    }

    @GetMapping("/users")
    public List<User> list(Pageable pageable) {
        return userService.list(getOrganizationId(), pageable);
    }

    @GetMapping("/users/{userId}")
    public User get(@PathVariable String userId) {
        return userService.get(userId, getOrganizationId());
    }

    @PutMapping("/users/{userId}/admin")
    public User updateAdmin(@PathVariable String userId, @Valid @RequestBody UserAdminUpdateRequest request) {
        checkAdmin();
        return userService.updateAdmin(userId, request, getOrganizationId());
    }

    @PutMapping("/users/{userId}/password")
    public UserPasswordResponse resetPassword(@PathVariable String userId) {
        checkAdmin();
        return userService.resetPassword(userId, getOrganizationId());
    }

    @DeleteMapping("/users/{userId}")
    public User delete(@PathVariable String userId) {
        checkAdmin();
        return userService.delete(userId, getOrganizationId());
    }
}
