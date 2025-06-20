function displayError(parent, message) {
    parent.html(`<div class="alert alert-danger">${message}</div>`);
}

function displaySuccess(parent, message) {
    parent.html(`<div class="alert alert-success">${message}</div>`);
}

function getHeaders() {
    return {
        Authorization: 'Bearer ' + localStorage.getItem('token'),
    }
}

function userSignUp(username, password, organizationName, billingEmail, success, error) {
    $.ajax({
        url: "/api/v1/users/signup",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            username,
            password,
            organizationName,
            billingEmail,
        }),
        success: success,
        error: error,
    })
}

function getUserToken(username, password, organizationName, success, error) {
    $.ajax({
        url: "/api/v1/users/token",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            username,
            password,
            organizationName,
        }),
        success: success,
        error: error,
    })
}

function getCurrentUser(success, error) {
    $.ajax({
        url: "/api/v1/users/me",
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function changeCurrentUserPassword(password, success, error) {
    $.ajax({
        url: "/api/v1/users/me/password",
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            password,
        }),
        success: success,
        error: error,
    })
}

function getOrganization(success, error) {
    $.ajax({
        url: "/api/v1/organization",
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateOrganization(billingEmail, success, error) {
    $.ajax({
        url: "/api/v1/organization",
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            billingEmail,
        }),
        success: success,
        error: error,
    })
}

function addUser(username, admin, success, error) {
    $.ajax({
        url: "/api/v1/users",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            username,
            admin,
        }),
        success: success,
        error: error,
    })
}

function listUsers(page, size, success, error) {
    $.ajax({
        url: "/api/v1/users",
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page,
            size,
        },
        success: success,
        error: error,
    })
}

function getUser(userId, success, error) {
    $.ajax({
        url: `/api/v1/users/${userId}`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateUserAdmin(userId, admin, success, error) {
    $.ajax({
        url: `/api/v1/users/${userId}/admin`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            admin,
        }),
        success: success,
        error: error,
    })
}

function resetUserPassword(userId, success, error) {
    $.ajax({
        url: `/api/v1/users/${userId}/password`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function deleteUser(userId, success, error) {
    $.ajax({
        url: `/api/v1/users/${userId}`,
        method: "DELETE",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createEnvironment(name, description, success, error) {
    $.ajax({
        url: `/api/v1/environments`,
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function listEnvironments(page, size, success, error) {
    $.ajax({
        url: "/api/v1/environments",
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page,
            size,
        },
        success: success,
        error: error,
    })
}

function getEnvironment(environmentId, success, error) {
    $.ajax({
        url: `/api/v1/environments/${environmentId}`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateEnvironment(environmentId, name, description, success, error) {
    $.ajax({
        url: `/api/v1/environments/${environmentId}`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function deleteEnvironment(environmentId, success, error) {
    $.ajax({
        url: `/api/v1/environments/${environmentId}`,
        method: "DELETE",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createProject(name, description, success, error) {
    $.ajax({
        url: `/api/v1/projects`,
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function listProjects(page, size, success, error) {
    $.ajax({
        url: "/api/v1/projects",
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page,
            size,
        },
        success: success,
        error: error,
    })
}

function getProject(projectId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateProject(projectId, name, description, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function deleteProject(projectId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}`,
        method: "DELETE",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createApiKey(projectId, name, description, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys`,
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function listApiKeys(projectId, page, size, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page,
            size,
        },
        success: success,
        error: error,
    })
}

function getApiKey(projectId, apiKeyId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys/${apiKeyId}`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateApiKey(projectId, apiKeyId, name, description, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys/${apiKeyId}`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name,
            description,
        }),
        success: success,
        error: error,
    })
}

function deleteApiKey(projectId, apiKeyId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys/${apiKeyId}`,
        method: "DELETE",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error
    })
}

function getApiKeySecret(projectId, apiKeyId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys/${apiKeyId}/secret`,
        method: "GET",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function resetApiKeySecret(projectId, apiKeyId, success, error) {
    $.ajax({
        url: `/api/v1/projects/${projectId}/api-keys/${apiKeyId}/secret`,
        method: "PUT",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error
    })
}
