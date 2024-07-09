$('#form-register').on('submit', createUser);

function createUser(event){
    event.preventDefault();
    if ($('#password').val() != $('#confirm-password').val()){
        Swal.fire("Ops...", "Passwords do not match!", "error");
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
    }).done(function(){
        Swal.fire("Success!", "User registered successfully!", "success")
        .then(function() {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val()
                }
            }).done(function() {
                window.location = "/home";
            }).fail(function() {
                Swal.fire("Ops...", "Error authenticating user!", "error");
            })
        })

    }).fail (function(error){
        console.log(error)
        Swal.fire("Ops...", "Error registering user!", "error");
    });
}
