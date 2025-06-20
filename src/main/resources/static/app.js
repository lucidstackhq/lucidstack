function displayError(parent, message) {
    parent.html(`<div class="alert alert-danger">${message}</div>`);
}

function displaySuccess(parent, message) {
    parent.html(`<div class="alert alert-success">${message}</div>`);
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
