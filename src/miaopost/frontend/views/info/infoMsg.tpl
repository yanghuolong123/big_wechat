<div class="comment ">
	<div class="row">
		<div class="col-md-8 col-xs-4"><h4 class="text-muted">留言区</h4></div>
		<div class="col-md-4 col-xs-8 text-right"><a href="javascript:;" class="msg-btn btn btn-success">我要留言</a></div>
	</div>
	
	<div id="commentlist">
		{{range .imvos}}
		<div class="msg">
			<div class="row">
				<div class="col-md-8 col-xs-3">
					<span><a href="javascript:;">{{.User.Nickname}}</a>{{if .Im.Pid}} @{{.Parent.User.Nickname}}{{end}}</span>
				</div>
				<div class="col-md-4 col-xs- 9 text-right">
					<span><a href="#" onclick="replyMsg({{.Im.Id}});return false;">回复</a></span>
					<span><a href="#" onclick="msgDelSuggest({{.Im.Id}}, this);return false;">建删</a></span>
					<span><a href="#" onclick="admire({{.Im.Id}}); return false;">赞赏</a></span>
					<span><a href="#" onclick="supportInfoMsg({{.Im.Id}}, this);return false;"><span class="glyphicon glyphicon-heart" aria-hidden="true"></span><span class="support_num">{{.Im.Support}}</span></a></span>
				</div>
			</div>
			<div class="row">
				<div class="col-md-12">{{str2html .Im.Content}}</div>
			</div>
		</div>
		{{end}}
	</div>
</div>

<div id="msgModal" class="modal fade" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
    <div class="modal-dialog" role="document">
      <div class="modal-content">   
          <div class="modal-header">
	        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h4 class="modal-title">我要留言</h4>
         </div>         
        <div class="modal-body row">
            <div class="col-md-12">	
		<input type="hidden" id="info_id" name="info_id" value="{{.info.Id}}">
		<input type="hidden" id="msg_pid" name="pid" value="0">
		<!--<textarea id="info_msg" class="form-control"></textarea>-->
		<form>
			<textarea id="info_msg" name="content" style="width:100%;height:200px;visibility:hidden;"></textarea>
		</form>
	</div>
        </div>
        <div class="modal-footer">
	        <button class="btn btn-success info-msg-btn">提交</button>
	        <button type="button" class="btn btn-info" data-dismiss="modal">取消</button> 
        </div>
      </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

<div id="admireModal" class="modal fade" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
    <div class="modal-dialog" role="document">
      <div class="modal-content">   
          <div class="modal-header">
	        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h5 class="modal-title">请选择您的打赏金额</h5>
         </div>         
        <div class="modal-body row">     
        	<input type="hidden" id="admire_msg_id" value="0">      
          	<div class="col-md-12 admire_pay text-center">
	          	<a href="#" onclick="admirePay(0.1);return false;" class="btn btn-success">0.1元</a>
	          	<a href="#" onclick="admirePay(0.5);return false;" class="btn btn-success">0.5元</a>
	          	<a href="#" onclick="admirePay(1);return false;" class="btn btn-success">1元</a>
	          	<a href="#" onclick="admirePay(2);return false;" class="btn btn-success">2元</a>
	          	<a href="#" onclick="admirePay(5);return false;" class="btn btn-success">5元</a>
	          	<a href="#" onclick="admirePay(10);return false;" class="btn btn-success">10元</a>
          	</div>           
        </div>
      </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

{{template "pay/qrcode.tpl" .}} 

<link rel="stylesheet" href="/static/plugin/kindeditor/themes/default/default.css" />
<script charset="utf-8" src="/static/plugin/kindeditor/kindeditor-all.modify.js"></script>
<script charset="utf-8" src="/static/plugin/kindeditor/lang/zh-CN.js"></script>
<script type="text/javascript">
	var editor;
	KindEditor.ready(function(K) {
		editor = K.create('textarea[name="content"]', {
			resizeType : 1,
			allowPreviewEmoticons : false,
			allowImageUpload : false,
			afterBlur: function () { this.sync(); },
			//afterFocus: function(){ this.html("");},
			items : [
				//'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold', 'italic', 'underline',
				//'removeformat', '|', 'justifyleft', 'justifycenter', 'justifyright', 'insertorderedlist',
				//'insertunorderedlist', '|', 'emoticons', 'image', 'link'
				'insertorderedlist','bold',  '|', 'emoticons', 'image', 'link'
				]
		});		
	});

	$(function(){
		$('#msgModal').off('shown.bs.modal').on('shown.bs.modal', function (e) {
		    $(document).off('focusin.modal');//解决编辑器弹出层文本框不能输入的问题
		});
	});
</script>
