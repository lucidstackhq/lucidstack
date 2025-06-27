$.ajaxSetup({
    statusCode: {
        401: function(jqxhr, textStatus, errorThrown) {
            localStorage.clear()
            window.location.href = "/login"
        }
    }
})

function displayError(parent, message) {
    parent.html(`<div class="alert alert-danger">${message}</div>`);
}

function displaySuccess(parent, message) {
    parent.html(`<div class="alert alert-success">${message}</div>`);
}

function getHeaders() {
    return {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
    }
}

function userSignUp(username, password, organizationName, billingEmail, success, error) {
    $.ajax({
        method: "POST",
        url: "/api/v1/users/signup",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            username: username,
            password: password,
            organization_name: organizationName,
            billing_email: billingEmail,
        }),
        success: success,
        error: error,
    })
}

function getUserToken(username, password, organizationName, success, error) {
    $.ajax({
        method: "POST",
        url: "/api/v1/users/token",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            username: username,
            password: password,
            organization_name: organizationName,
        }),
        success: success,
        error: error,
    })
}

function getCurrentUser(success, error) {
    $.ajax({
        method: "GET",
        url: "/api/v1/users/me",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function changeCurrentUserPassword(password, success, error) {
    $.ajax({
        method: "PUT",
        url: "/api/v1/users/me/password",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            password: password,
        }),
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function getOrganization(success, error) {
    $.ajax({
        method: "GET",
        url: "/api/v1/organization",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateOrganization(billingEmail, success, error) {
    $.ajax({
        method: "PUT",
        url: "/api/v1/organization",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            billing_email: billingEmail,
        }),
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createUser(username, admin, success, error) {
    $.ajax({
        method: "POST",
        url: "/api/v1/users",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify({
            username: username,
            admin: admin,
        }),
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function listUsers(page, size, success, error) {
    $.ajax({
        method: "GET",
        url: "/api/v1/users",
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function getUser(userId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/users/${userId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function deleteUser(userId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/users/${userId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function resetUserPassword(userId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/users/${userId}/password`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateUserAdmin(userId, admin, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/users/${userId}/admin`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            admin: admin,
        }),
        success: success,
        error: error,
    })
}

function createApp(name, description, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/apps`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function listApps(page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/apps`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getApp(appId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/apps/${appId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateApp(appId, name, description, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/apps/${appId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteApp(appId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/apps/${appId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function getAppSecret(appId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/apps/${appId}/secret`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function resetAppSecret(appId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/apps/${appId}/secret`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createEnvironment(name, description, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/environments`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function listEnvironments(page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/environments`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getEnvironment(environmentId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/environments/${environmentId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateEnvironment(environmentId, name, description, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/environments/${environmentId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteEnvironment(environmentId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/environments/${environmentId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createModel(name, description, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/models`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function listModels(page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getModel(modelId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateModel(modelId, name, description, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/models/${modelId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteModel(modelId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/models/${modelId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error
    })
}

function createEntity(modelId, name, description, environmentId, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/entities`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
            environment_id: environmentId,
            model_id: modelId,
        }),
        success: success,
        error: error,
    })
}

function listEntities(modelId, environmentId, page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/environments/${environmentId}/entities`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function createProperty(modelId, name, description, dataSchema, defaultValue, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/models/${modelId}/properties`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
            data_schema: dataSchema,
            default_value: defaultValue,
        }),
        success: success,
        error: error,
    })
}

function listProperties(modelId, page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/properties`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getProperty(propertyId, modelId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/properties/${propertyId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateProperty(propertyId, name, description, modelId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/models/${modelId}/properties/${propertyId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteProperty(propertyId, modelId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/models/${modelId}/properties/${propertyId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updatePropertyDefaultValue(propertyId, defaultValue, modelId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/models/${modelId}/properties/${propertyId}/default-value`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            default_value: defaultValue,
        }),
        success: success,
        error: error,
    })
}

function createAction(modelId, name, description, inputSchema, outputSchema, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/models/${modelId}/actions`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
            input_schema: inputSchema,
            output_schema: outputSchema,
        }),
        success: success,
        error: error,
    })
}

function listActions(modelId, page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/actions`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getAction(actionId, modelId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/actions/${actionId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateAction(actionId, name, description, modelId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/models/${modelId}/actions/${actionId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteAction(actionId, modelId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/models/${modelId}/actions/${actionId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function createEvent(modelId, name, description, dataSchema, success, error) {
    $.ajax({
        method: "POST",
        url: `/api/v1/models/${modelId}/events`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
            data_schema: dataSchema,
        }),
        success: success,
        error: error,
    })
}

function listEvents(modelId, page, size, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/events`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: {
            page: page,
            size: size,
        },
        success: success,
        error: error,
    })
}

function getEvent(eventId, modelId, success, error) {
    $.ajax({
        method: "GET",
        url: `/api/v1/models/${modelId}/events/${eventId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}

function updateEvent(eventId, name, description, modelId, success, error) {
    $.ajax({
        method: "PUT",
        url: `/api/v1/models/${modelId}/events/${eventId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        data: JSON.stringify({
            name: name,
            description: description,
        }),
        success: success,
        error: error,
    })
}

function deleteEvent(eventId, modelId, success, error) {
    $.ajax({
        method: "DELETE",
        url: `/api/v1/models/${modelId}/events/${eventId}`,
        dataType: "json",
        contentType: "application/json",
        headers: getHeaders(),
        success: success,
        error: error,
    })
}
