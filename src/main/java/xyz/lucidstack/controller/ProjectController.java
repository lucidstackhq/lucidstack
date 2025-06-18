package xyz.lucidstack.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.web.bind.annotation.*;
import xyz.lucidstack.auth.AuthenticatedUserController;
import xyz.lucidstack.model.Project;
import xyz.lucidstack.request.ProjectCreationRequest;
import xyz.lucidstack.service.ProjectService;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class ProjectController extends AuthenticatedUserController {

    private final ProjectService projectService;

    @PostMapping("/projects")
    public Project create(@Valid @RequestBody ProjectCreationRequest request) {
        return projectService.create(request, getUser());
    }

    @GetMapping("/projects")
    public List<Project> list(Pageable pageable) {
        return projectService.list(getUser(), pageable);
    }

    @GetMapping("/projects/{projectId}")
    public Project get(@PathVariable String projectId) {
        return projectService.get(projectId, getUser());
    }
}
