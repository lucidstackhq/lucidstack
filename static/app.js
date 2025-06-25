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
