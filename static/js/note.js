$(window).bind("load", function () {
    $('#button8').click(function () {
        genGrids(8)
    });
    $('#button16').click(function () {
        genGrids(16)
    });
    $('#button32').click(function () {
        genGrids(32)
    });
    $('#button64').click(function () {
        genGrids(64)
    });

    var number = 8;
    genGrids(number);

    var today = getCurrentDate();
    $('#noteDate').html(today);

    getNote();

    t = setInterval(function () {
        const {d, note} = updateNote();
    }, 1000)
});

const host = window.location.href;

function updateNote() {
    d = new Date($('#noteDate').html());
    note = {
        Date: d
    };
    fetch(host + "api/note/",
        {
            method: 'POST',
            body: JSON.stringify(note),
        }
    ).catch(function (err) {
        console.log(err)
    });
    return {d, note};
}
var t;
function getNote() {
    var note = {"Date": "0000-00-00"};
    fetch(host + '/api/note/', {method: 'GET'})
        .then(function (response) {
            return response.json();
        })
        .then(function (js) {
            note = js;
            if (note.Date !== null) {
                d = FormatDate(new Date(note.Date));
                $('#noteDate').html(d)
            }
        });
}

function genGrids(number) {
    $('.baseGrid').remove();

    [x, y] = getColumnRowNums(number);

    columnWidth = Math.floor(100 / x);
    columnTmp = "repeat(" + x.toString() + "," + columnWidth.toString() + "%)";

    rowHeight = Math.floor(100 / y);
    rowTmp = "repeat(" + y.toString() + "," + rowHeight.toString() + "%)";

    parent = $('#parent');
    parent.css("grid-template-columns", columnTmp);
    parent.css("grid-template-rows", rowTmp);
    var num = number - 1;
    for (i = 0; i < num; i++) {
        var template = $('#baseGridTemplate').html();
        parent.append(template)
    }
}

function getColumnRowNums(number) {
    switch (number) {
        case 8:
            return [4, 2];
        case 16:
            return [4, 4];
        case 32:
            return [8, 4];
        case 64:
            return [8, 8];
    }
}

function getCurrentDate() {
    var today = new Date();
    today = FormatDate(today);
    return today;
}

function FormatDate(date) {
    var d = [];
    d.push(date.getFullYear());
    d.push(String(date.getMonth() + 1).padStart(2, '0'));
    d.push(String(date.getDate()).padStart(2, '0'));
    return d.join("-");
}
