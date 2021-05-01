$('#nova-publicacao').on('submit', createPublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(() => {
        // alert("Publicação criada com sucesso!");
        //refresh
        window.location = "/home"
    }).fail(() => {
        alert("Erro ao criar publicação!");
    })
}