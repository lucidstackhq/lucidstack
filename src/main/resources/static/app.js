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

function updateUserAdmin() {

}

function resetUserPassword() {

}

function deleteUser() {

}