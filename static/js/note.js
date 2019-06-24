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
    getNoteList();

    t1 = setInterval(function () {
        const {d, note} = updateNote();
    }, 1000);

    t2 = setInterval(function () {
        getNoteList();
    }, 2000);
});

function getHost() {
    u = new URL(window.location.href);
    return u.origin
}

function getNoteList() {
    fetch(getHost()+"/api/note/",
        {method:'GET'})
        .then(function (response) {
            return response.json()
        })
        .then(function (js) {
            var tmp = $('#noteItemTemplate').html();
            var count = 0;
            $('.noteItem').remove();
            for(item in js) {
                $('#noteListItems').append(tmp);
                count ++;
            }
            for(i=0;i<count;i++){
                $('.noteItem')[i].innerHTML = js[i].Title;
            }
        });
}

function getNoteId() {
    return $('#noteId').attr("value");
}
function updateNote() {
    d = new Date($('#noteDate').html());
    title = $('#noteTitle').html();

    var grids = [];
    $('.baseGrid').each(function () {
        var keyword = $(this).find(".keyword").html();
        var comment = $(this).find(".comment").html();
        var grid = {
            Keyword: keyword,
            Comment: comment
        };
        grids.push(grid)
    });

    note = {
        Date: d,
        Title: title,
        Grids: grids
    };
    fetch(getHost() + "/api/note/" + getNoteId(),
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
    fetch(getHost() + '/api/note/'+getNoteId(), {method: 'GET'})
        .then(function (response) {
            return response.json();
        })
        .then(function (js) {
            note = js;
            if (note.Date !== null) {
                d = FormatDate(new Date(note.Date));
                $('#noteDate').html(d);
                $('#noteTitle').html(note.Title);
                if (note.Grids !== undefined) {
                    for (i = 0; i < note.Grids.length; i++) {
                        let grid = note.Grids[i];
                        var keyword = grid.Keyword;
                        var comment = grid.Comment;
                        if (keyword !== undefined) {
                            $('.keyword')[i].innerHTML = keyword
                        }
                        if (comment !== undefined) {
                            $('.comment')[i].innerHTML = comment
                        }
                    }
                }

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

    let note = $('#note');
    note.css("grid-template-columns", columnTmp);
    note.css("grid-template-rows", rowTmp);
    var num = number - 1;
    for (i = 0; i < num; i++) {
        var template = $('#baseGridTemplate').html();
        note.append(template)
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
