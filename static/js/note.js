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
    switch (number) {
        case 8:
            x = 4;
            y = 2;
            break;
        case 16:
            x = 4;
            y = 4;
            break;
        case 32:
            x = 8;
            y = 4;
            break
        case 64:
            x = 8;
            y = 8;
            break;

    }


    columnWidth = Math.floor(100 / x);
    console.log("columnWidth = ", columnWidth);
    columnTmp = "repeat(" + x.toString() + "," + columnWidth.toString() + "%)";
    console.log("columnTmp = ", columnTmp);

    rowHeight = Math.floor(100 / y);
    console.log("rowHeight = ", rowHeight);
    rowTmp = "repeat(" + y.toString() + "," + rowHeight.toString() + "%)";
    console.log("rowTmp = ", rowTmp);

    parent = $('#parent');
    parent.css("grid-template-columns", columnTmp);
    parent.css("grid-template-rows", rowTmp);
    var num = number - 1;
    for (i = 0; i < num; i++) {
        var template = $('#baseGridTemplate').html();
        parent.append(template)
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

