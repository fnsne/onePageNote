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

    var today = getCurrentDate();
    $('#noteDate').html(today);

    var number = 8;

    genGrids(number);
});


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
    var d = [];
    d.push(today.getFullYear());
    d.push(String(today.getMonth() + 1).padStart(2, '0'));
    d.push(String(today.getDate()).padStart(2, '0'));
    today = d.join("-");
    return today;
}

