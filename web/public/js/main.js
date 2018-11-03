import 'jquery';
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/main.css';
import '../css/game.css';

function updateBoard(res, elem) {
    if (res.currentPlayer === 0) {
        elem.addClass("player-1-move");
        $('.player-1').removeClass('in-play');
        $('.player-2').addClass('in-play');
    } else {
        elem.addClass("player-2-move");
        $('.player-2').removeClass('in-play');
        $('.player-1').addClass('in-play');
    }

    if(res.winner !== -1) {
        $("#winnerMsg").html(`Player ${res.winner} is the winner!`);
        $("#winnerMsg").removeClass("hidden");
    }
}

$('.board-cell').on('click', function() {
    var elem = $(this)
    var move = elem.find('span').html();
    var url = '/move?move=' + move.trim();
    $.get(url, function(data) {
        updateBoard(data, elem);
    });
});