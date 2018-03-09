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

function isWeiXin() {
	var ua = window.navigator.userAgent.toLowerCase();
	//console.log(ua);//mozilla/5.0 (iphone; cpu iphone os 9_1 like mac os x) applewebkit/601.1.46 (khtml, like gecko)version/9.0 mobile/13b143 safari/601.1
	if (ua.match(/MicroMessenger/i) == 'micromessenger') {
		return true;
	} else {
		return false;
	}
}

function sleep(milliSeconds){
    var startTime = new Date().getTime(); // get the current time
    while (new Date().getTime() < startTime + milliSeconds); // hog cpu
}