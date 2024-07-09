$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#update-user').on('submit', updateUser);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow(event) {
    event.preventDefault();
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Oops...", "Error unfollowing the user!", "error");
        $('#unfollow').prop('disabled', false);
    });
}

function follow(event) {
    event.preventDefault();
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Oops...", "Error following the user!", "error");
        $('#follow').prop('disabled', false);
    });
}

function updateUser(event) {
    event.preventDefault();

    $.ajax({
        url: "/update-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick_name: $('#nick_name').val(),
        }
    }).done(function() {
        Swal.fire("Success!", "User updated successfully!", "success")
            .then(function() {
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Oops...", "Error updating the user!", "error");
    });
}

function updatePassword(event) {
    event.preventDefault();

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire("Oops...", "Passwords do not match!", "warning");
        return;
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            current: $('#current-password').val(),
            new: $('#new-password').val()
        }
    }).done(function() {
        Swal.fire("Success!", "Password updated successfully!", "success")
            .then(function() {
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Oops...", "Error updating the password!", "error");
    });
}

function deleteUser() {
    Swal.fire({
        title: "Warning!",
        text: "Are you sure you want to delete your account? This action cannot be undone!",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation) {
        if (confirmation.value) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Success!", "Your account has been deleted successfully!", "success")
                    .then(function() {
                        window.location = "/logout";
                    });
            }).fail(function() {
                Swal.fire("Oops...", "Error deleting your account!", "error");
            });
        }
    });
}
