package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.embedded.Resource;
import xyz.lucidstack.embedded.resource.ProjectResource;
import xyz.lucidstack.embedded.resource.RootResource;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.exception.NotAllowedException;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.model.Project;
import xyz.lucidstack.repository.ProjectRepository;
import xyz.lucidstack.request.ProjectCreationRequest;
import xyz.lucidstack.request.ProjectUpdateRequest;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Service
@RequiredArgsConstructor
public class ProjectService {

    private final ProjectRepository projectRepository;

    private final AccessService accessService;

    public Project create(ProjectCreationRequest request, AuthenticatedUser creator) {
        if (!accessService.hasPermission(new RootResource(), creator, "create")) {
            throw new NotAllowedException();
        }

        if (projectRepository.existsByNameAndOrganizationId(request.getName(), creator.getOrganizationId())) {
            throw new ClientException(String.format("Project %s already exists", request.getName()));
        }

        Project project = Project.builder()
                .name(request.getName())
                .description(request.getDescription())
                .creatorId(creator.getId())
                .organizationId(creator.getOrganizationId())
                .build();

        project = projectRepository.save(project);
        accessService.addFullAccess(new ProjectResource(project.getId()), creator.getId(), creator.getOrganizationId());

        return project;
    }

    public List<Project> list(AuthenticatedUser requester, Pageable pageable) {
        if (requester.getAdmin()) {
            return projectRepository.findByOrganizationId(requester.getOrganizationId(), pageable);
        } else {
            List<String> projectIds = new ArrayList<>();
            List<Resource> resources = accessService.listResources(Map.of("type", "project"), "read", requester, pageable);
            for (Resource resource: resources) {
                ProjectResource projectResource = (ProjectResource) resource;
                projectIds.add(projectResource.getProjectId());
            }

            return projectRepository.findByIdInAndOrganizationId(projectIds, requester.getOrganizationId());
        }
    }

    public Project get(String projectId, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new ProjectResource(projectId), requester, "read")) {
            throw new NotAllowedException();
        }

        Project project = projectRepository.findByIdAndOrganizationId(projectId, requester.getOrganizationId());

        if (project == null) {
            throw new NotFoundException("Project not found");
        }

        return project;
    }

    public Project update(String projectId, ProjectUpdateRequest request, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new ProjectResource(projectId), requester, "update")) {
            throw new NotAllowedException();
        }

        Project project = projectRepository.findByIdAndOrganizationId(projectId, requester.getOrganizationId());

        if  (project == null) {
            throw new NotFoundException("Project not found");
        }

        if (StringUtils.hasText(request.getName())) {
            if (projectRepository.existsByIdNotAndNameAndOrganizationId(projectId, request.getName(),  requester.getOrganizationId())) {
                throw new ClientException(String.format("Project %s already exists", request.getName()));
            }

            project.setName(request.getName());
        }

        project.setDescription(request.getDescription());

        return projectRepository.save(project);
    }

    public Project delete(String projectId, AuthenticatedUser requester) {
        if (!accessService.hasPermission(new ProjectResource(projectId), requester, "delete")) {
            throw new NotAllowedException();
        }

        Project project = projectRepository.findByIdAndOrganizationId(projectId, requester.getOrganizationId());

        if (project == null) {
            throw new NotFoundException("Project not found");
        }

        projectRepository.delete(project);

        return project;
    }
}
