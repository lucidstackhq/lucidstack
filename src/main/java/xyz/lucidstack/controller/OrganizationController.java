package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.Organization;
import xyz.lucidstack.request.OrganizationUpdateRequest;
import xyz.lucidstack.service.OrganizationService;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class OrganizationController extends AuthenticatedUserController {

    private final OrganizationService organizationService;

    @GetMapping("/organization")
    public Organization get() {
        return organizationService.get(getOrganizationId());
    }

    @PutMapping("/organization")
    public Organization update(@Valid @RequestBody OrganizationUpdateRequest request) {
        checkAdmin();
        return organizationService.update(getOrganizationId(), request);
    }
}
