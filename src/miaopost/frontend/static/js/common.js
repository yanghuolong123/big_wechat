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

/*
 * 操作确认 modal
 */
var actionConfirm = function(settings) {
    var modalId = 'fcjz-confirm-modal';
    var modalTplHtml = '<div id="' + modalId + '" class="modal modal-default fade "> <div class="modal-dialog"> <div class="modal-content"> <div class="modal-header"> <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button> <h4 class="modal-title">提示</h4> </div> <div class="modal-body"> <p>One fine body…</p> </div> <div class="modal-footer"> <button type="button" class="btn btn-default pull-left" data-dismiss="modal">取消</button> <button type="button" class="btn btn-warning confirm-btn" data-dismiss="modal">确定</button> </div> </div></div></div>';
    var modalEle = null;
    if ($('#' + modalId).length == 0) {
        $('body').append($(modalTplHtml));
        modalEle = $('#' + modalId);
        modalEle.modal({
            backdrop: 'static',
            keyboard: false,
            show: false
        });
    } else {
        modalEle = $('#' + modalId);
    }
    $('.modal-body', modalEle).html(settings.msg);
    $('.confirm-btn', modalEle).off('click').on('click', function() {
        settings['confirm']();
        modalEle.modal('hide');
    });
    modalEle.modal('show');
}

/*
 * 操作成功提示 modal
 */
var greeting = function(settings) {
    var modalId = 'fcjz-greeting-modal';
    var modalTplHtml = '<div id="' + modalId + '" class="modal modal-default fade "> <div class="modal-dialog"> <div class="modal-content"> <div class="modal-header"> <h4 class="modal-title">操作成功</h4> </div> <div class="modal-body"> <p>One fine body…</p> </div> <div class="modal-footer">  <button type="button" class="btn btn-success confirm-btn" data-dismiss="modal">知道了</button> </div> </div></div></div>';
    var modalEle = null;
    if ($('#' + modalId).length == 0) {
        $('body').append($(modalTplHtml));
        modalEle = $('#' + modalId);
        modalEle.modal({
            backdrop: 'static',
            keyboard: false,
            show: false
        });
    } else {
        modalEle = $('#' + modalId);
    }
    $('.modal-body', modalEle).html(settings.msg);
    $('.modal-title', modalEle).html(settings['title'] ? settings['title'] : '操作成功');
    if (settings['confirm']) {
        $('.confirm-btn', modalEle).off('click').on('click', function() {
            settings['confirm']();
            modalEle.modal('hide');
        });
    }
    modalEle.modal('show');
}

/*
 * 操作警告
 */
var alerting = function(settings) {
    var modalId = 'fcjz-alerting-modal';
    var modalTplHtml = '<div id="' + modalId + '" class="modal modal-default fade "> <div class="modal-dialog"> <div class="modal-content"> <div class="modal-header  alert-warning"> <h4 class="modal-title">警告</h4> </div> <div class="modal-body"> <p>One fine body…</p> </div> <div class="modal-footer">  <button type="button" class="btn btn-warning confirm-btn" data-dismiss="modal">知道了</button> </div> </div></div></div>';
    var modalEle = null;
    if ($('#' + modalId).length == 0) {
        $('body').append($(modalTplHtml));
        modalEle = $('#' + modalId);
        modalEle.modal({
            backdrop: 'static',
            keyboard: false,
            show: false
        });
    } else {
        modalEle = $('#' + modalId);
    }
    $('.modal-body', modalEle).html(settings.msg);
    $('.modal-title', modalEle).html(settings['title'] ? settings['title'] : '警告');
    if (settings['confirm']) {
        $('.confirm-btn', modalEle).off('click').on('click', function() {
            settings['confirm']();
            modalEle.modal('hide');
        });
    }
    modalEle.modal('show');
}