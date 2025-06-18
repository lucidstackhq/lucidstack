package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.auth.Jwt;
import xyz.lucidstack.exception.AuthenticationException;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.model.Organization;
import xyz.lucidstack.model.User;
import xyz.lucidstack.repository.UserRepository;
import xyz.lucidstack.request.UserSignUpRequest;
import xyz.lucidstack.request.UserTokenRequest;
import xyz.lucidstack.response.UserTokenResponse;
import xyz.lucidstack.util.Password;

@Service
@RequiredArgsConstructor
public class UserService {

    private final UserRepository userRepository;

    private final OrganizationService organizationService;

    private final Jwt jwt;

    public User signUp(UserSignUpRequest request) {
        if (organizationService.nameExists(request.getOrganizationName())) {
            throw new ClientException(String.format("Organization %s already exists", request.getOrganizationName()));
        }

        User user = User.builder()
                .username(request.getUsername())
                .password(Password.hash(request.getPassword()))
                .admin(true)
                .build();

        user = userRepository.save(user);

        Organization organization = organizationService.save(request.getOrganizationName(), request.getBillingEmail(), user.getId());
        user.setOrganizationId(organization.getId());

        return userRepository.save(user);
    }

    public UserTokenResponse getToken(UserTokenRequest request) {
        Organization organization = organizationService.getByName(request.getOrganizationName());

        User user = userRepository.findByUsernameAndOrganizationId(request.getUsername(), organization.getId());

        if (user == null || !Password.verify(request.getPassword(), user.getPassword())) {
            throw new AuthenticationException();
        }

        String token = jwt.getUserToken(AuthenticatedUser.builder()
                        .id(user.getId())
                        .organizationId(organization.getId())
                        .admin(user.getAdmin())
                .build());

        return UserTokenResponse.builder().token(token).build();
    }
}
