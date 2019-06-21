$(window).bind("load", function () {
    fetch(getHost()+"/api/note/", {method:"GET"})
        .then(function (response) {
            return response.json();
        })
        .then(function (js) {
            if (js !== null) {
                let noteCount = js.length;
                let noteTemplate = $("#noteTemplate").html();
                for ( i =0; i < noteCount; i++){
                    $("#noteList").append(noteTemplate)
                }

                for (i=0; i < noteCount; i++){
                    $('.noteTitle')[i].innerHTML = js[i].Title
                }
            }
        })
});

function getHost() {
    u = new URL(window.location.href);
    return u.origin
}