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
				<div class="col-md-12">{{.Im.Content}}</div>
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
            <div class="col-md-7">	
		<input type="hidden" id="info_id" name="info_id" value="{{.info.Id}}">
		<input type="hidden" id="msg_pid" name="pid" value="0">
		<textarea id="info_msg" class="form-control"></textarea>
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
	        <h4 class="modal-title">如果你觉得留言不错，请随意打赏。您的支持将鼓励我继续创作！</h4>
         </div>         
        <div class="modal-body row">     
        	<input type="hidden" id="admire_msg_id" value="0">      
          	<div class="col-md-8 admire_pay">
	          	<a href="#" onclick="admirePay(0.1);return false;" class="btn btn-success">0.1元</a>
	          	<a href="#" onclick="admirePay(0.5);return false;" class="btn btn-success">0.5元</a>
	          	<a href="#" onclick="admirePay(1);return false;" class="btn btn-success">1元</a>
	          	<a href="#" onclick="admirePay(2);return false;" class="btn btn-success">2元</a>
	          	<a href="#" onclick="admirePay(5);return false;" class="btn btn-success">5元</a>
          	</div>           
        </div>
      </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div><!-- /.modal -->

<div id="qrPayModal" class="modal fade" tabindex="-1" role="dialog">
  <div class="modal-dialog modal-sm" role="document">
    <div class="modal-content">
      <div class="modal-header">      	
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
      </div>
      <div class="modal-body">
      	<div class="center-block">
      		<h4>微信扫码支付</h4>
	      	<p class="center-block">
	      		<img id="pay_qr_img" src="/static/img/loading.gif" />
	      	</p>
	        	<p class="center-block">支付金额 : <span class="pay_amount">￥2.00</span></p>
      	</div>
      	
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->