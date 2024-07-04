$('#form-register').on('submit', createUser);

function createUser(event){
    event.preventDefault();
    if ($('#password').val() != $('#confirm-password').val()){
        alert("The passwords do not match!");
    }

    $.ajax({
        url: "/users",
        method:"POST",
        data:{
            name: $('#name').val(),
            email: $('#email').val(),
            nick_name: $('#nick-name').val(),
            password: $('#password').val(),
        }
    })
}