import 'jquery';
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/main.css';
import '../css/game.css';

$('.board-cell').on('click', function() {
    var move = $(this).html();
    var url = '/move?move=' + move.trim();
    $.get(url, function(data) {
        console.log(data);
    });
});