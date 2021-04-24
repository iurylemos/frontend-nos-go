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
    });
}