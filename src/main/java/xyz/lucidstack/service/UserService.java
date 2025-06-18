package xyz.lucidstack.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import xyz.lucidstack.auth.AuthenticatedUser;
import xyz.lucidstack.auth.Jwt;
import xyz.lucidstack.exception.AuthenticationException;
import xyz.lucidstack.exception.ClientException;
import xyz.lucidstack.exception.NotFoundException;
import xyz.lucidstack.model.Organization;
import xyz.lucidstack.model.User;
import xyz.lucidstack.repository.UserRepository;
import xyz.lucidstack.request.*;
import xyz.lucidstack.response.UserTokenResponse;
import xyz.lucidstack.util.Password;
import xyz.lucidstack.util.Random;

import java.util.List;

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

    public User get(String userId, String organizationId) {
        User user = userRepository.findByIdAndOrganizationId(userId, organizationId);

        if (user == null) {
            throw new NotFoundException("User not found");
        }

        return user;
    }

    public User changePassword(String userId, UserPasswordChangeRequest request, String organizationId) {
        User user = get(userId, organizationId);
        user.setPassword(Password.hash(request.getPassword()));
        return userRepository.save(user);
    }

    public UserPasswordResponse add(UserAdditionRequest request, String creatorId, String organizationId) {
        if (userRepository.existsByUsernameAndOrganizationId(request.getUsername(), organizationId)) {
            throw new ClientException(String.format("Username %s already exists", request.getUsername()));
        }

        String password = Random.generateRandomString(16);

        User user = User.builder()
                .username(request.getUsername())
                .password(Password.hash(password))
                .admin(request.getAdmin())
                .creatorId(creatorId)
                .organizationId(organizationId)
                .build();

        user = userRepository.save(user);
        return UserPasswordResponse.builder().user(user).password(password).build();
    }

    public List<User> list(String organizationId, Pageable pageable) {
        return userRepository.findByOrganizationId(organizationId, pageable);
    }

    public User updateAdmin(String userId, UserAdminUpdateRequest request, String organizationId) {
        User user = get(userId, organizationId);
        user.setAdmin(request.getAdmin());
        return userRepository.save(user);
    }

    public UserPasswordResponse resetPassword(String userId, String organizationId) {
        User user = get(userId, organizationId);
        String password = Random.generateRandomString(16);
        user.setPassword(Password.hash(password));
        user = userRepository.save(user);
        return UserPasswordResponse.builder().user(user).password(password).build();
    }

    public User delete(String userId, String organizationId) {
        User user = get(userId, organizationId);
        userRepository.delete(user);
        return user;
    }
}
