package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.model.Organization;
import xyz.lucidstack.repository.OrganizationRepository;
import xyz.lucidstack.request.OrganizationUpdateRequest;

@Service
@RequiredArgsConstructor
public class OrganizationService {

    private final OrganizationRepository organizationRepository;

    public Organization save(String name, String billingEmail, String creatorId) {
        Organization organization = Organization.builder()
                .name(name)
                .billingEmail(billingEmail)
                .creatorId(creatorId)
                .build();

        return organizationRepository.save(organization);
    }

    public Organization get(String organizationId) {
        return organizationRepository.findById(organizationId)
                .orElseThrow(() -> new NotFoundException("Organization not found"));
    }

    public Organization getByName(String name) {
        Organization organization = organizationRepository.findByName(name);

        if (organization == null) {
            throw new NotFoundException(String.format("Organization %s not found", name));
        }

        return organization;
    }

    public Organization update(String organizationId, OrganizationUpdateRequest request) {
        Organization organization = get(organizationId);
        organization.setBillingEmail(request.getBillingEmail());
        return organizationRepository.save(organization);
    }

    public Boolean nameExists(String name) {
        return organizationRepository.existsByName(name);
    }
}
