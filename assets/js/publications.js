$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.unlike-publication', unlikePublication);

$('#update-publication').on('click', updatePublication);

$('.delete-publication').on('click', deletePublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $("#title").val(),
            content: $("#content").val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        Swal.fire("Ops...", "Error creating post!", "error");
    })

}

function likePublication(event) {
    event.preventDefault();
    const elementClicked = $(event.target);
    const publicationId = elementClicked.closest('div').data('publication-id');
 
    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/like`,
        method: "POST"
    }).done(function(event){
        const likesCount = elementClicked.next('span');
        const numberOfLikes = parseInt(likesCount.text());

        likesCount.text(numberOfLikes + 1);
        
        elementClicked.addClass('unlike-publication');
        elementClicked.addClass('text-danger');
        elementClicked.removeClass('like-publication');

    }).fail(function(event){
        Swal.fire("Ops...", "Error like post!", "error");
    }).always(function() {
        elementClicked.prop('disabled', false);
    });


}

function unlikePublication(event) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const publicationId = elementClicked.closest('div').data('publication-id');

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/unlike`,
        method: "POST"
    }).done(function() {
        const likesCount = elementClicked.next('span');
        const numberOfLikes = parseInt(likesCount.text());

        likesCount.text(numberOfLikes - 1);

        elementClicked.removeClass('unlike-publication');
        elementClicked.removeClass('text-danger');
        elementClicked.addClass('like-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Error unlike post!", "error");
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}

function updatePublication() {
    $(this).prop('disabled', true);

    const publicationId = $(this).data('publication-id');
    
    $.ajax({
        url: `/publications/${publicationId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        Swal.fire('Sucess!', 'Publication updated successfully!', 'success')
        .then(function() {
            window.location = "/home";
        })
    }).fail(function() {
    }).always(function() {
        $('#update-publication').prop('disabled', false);
    })
}

function deletePublication(evento) {
    evento.preventDefault();
    Swal.fire({
        title:"Attention!",
        text: "Are you sure you want to delete this post? This action is irreversible!",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation) {
        if (!confirmation.value) return;

        const elementClicked = $(evento.target);
        const publication = elementClicked.closest('div')
        const publicationId = publication.data('publication-id');
    
        elementClicked.prop('disabled', true);
    
        $.ajax({
            url: `/publications/${publicationId}`,
            method: "DELETE"
        }).done(function() {
            publication.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {

        });
    })
}