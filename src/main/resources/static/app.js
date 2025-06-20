function displayError(parent, message) {
    parent.html(`<div class="alert alert-danger">${message}</div>`);
}

function displaySuccess(parent, message) {
    parent.html(`<div class="alert alert-success">${message}</div>`);
}
