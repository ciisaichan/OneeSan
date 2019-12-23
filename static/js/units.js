function BSAlert(msgs, types) {
    var ids =  'msgboxs-' + Math.round(Math.random() * 25565);
    var html_type = 'alert-info';
    if (types == 1) {
        html_type = 'alert-success';
    } else if (types == 2) {
        html_type = 'alert-warning';
    } else if (types == 3) {
        html_type = 'alert-danger';
    }

    html_result = '<div id="' + ids + '" class="alert ' + html_type + ' alert-dismissible"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a>' + msgs + '</div>';
    $('.alerts').prepend(html_result);

    setTimeout(function () {
        $('#' + ids).alert('close');
    }, 5000);
}