$('#formulario-registro').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    // comparing passwords
    if ($('#password').val() != $('#confirm-password').val()) {
        alert("Senhas não coincidem");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#password').val(),
        }

    }).done(() => { // 201 (Created) 200 (OK) 204 (No Content)
        alert("Usuário cadastrado com sucesso!")
    }).fail((error) => { // 400 404 401 403 500
        console.log(error)
        alert("Error ao cadastrar o usuário!", error)
    });


}