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
					<span><a href="javascript:;">赞赏</a></span>
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