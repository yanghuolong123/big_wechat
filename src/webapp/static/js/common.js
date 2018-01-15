var displayTime,  $dom;
$(function(){
	var tpl = 	'<div class="prompt-container container">'+
					'<div class="prompt-content"></div>'+
				'</div>',
		domId = '#prompt-'+Math.round(Math.random()*100000);
	$dom = $(tpl).attr('id',domId),
		displayTime = 2500;
	if($(domId).length == 0) {
		$('body').append($dom.addClass('hide'));
	}
});

var prompt = function(option){
		var msg;  
		displayTime = option['displayTime'] ? option['displayTime'] : 2500; 
		if( typeof option == 'string' ) {
			msg = option;
		} else if (typeof option == 'object') {
			msg = option['msg'];	
		}
		$dom.find('.prompt-content').html(msg);
		$dom.removeClass('hide').addClass('show');
		if(typeof option == 'string' || !option['static']) {
			setTimeout(function(){
				$dom.removeClass('show').addClass('hide');
			},displayTime);
		}
	};