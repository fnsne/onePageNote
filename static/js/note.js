$(window).bind("load", function () {
    var today = getCurrentDate();
    $('#noteDate').html(today)
});

function getCurrentDate() {
    var today = new Date();
    var d = [];
    d.push(today.getFullYear());
    d.push(String(today.getMonth() + 1).padStart(2, '0'));
    d.push(String(today.getDate()).padStart(2, '0'));
    today = d.join("-");
    return today;
}
