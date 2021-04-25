$('#login').on('submit', login)

function login(event) {
    event.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            senha: $('#password').val(),
        }
    }).done(() => {
        window.location = '/home';
    }).fail((error) => {
        alert('Usuário ou senha invalidos')
    });
}